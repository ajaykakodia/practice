package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Strategy to implement
type RetryStrategy interface {
	NextBackoff(retryCount int) time.Duration
}

// concrete strategy

// Fixed Delay Strategy
type FixedDelay struct {
	Delay time.Duration
}

func (f *FixedDelay) NextBackoff(retryCount int) time.Duration {
	return f.Delay
}

// Exponential backoff strategy
type ExponentialBackoff struct {
	BaseDelay time.Duration
	MaxDelay  time.Duration
}

func (e *ExponentialBackoff) NextBackoff(retryCount int) time.Duration {
	delay := e.BaseDelay * (1 << retryCount) //exponential growth
	if delay > e.MaxDelay {
		delay = e.MaxDelay
	}

	return delay
}

type JitterBackoff struct {
	BaseDelay time.Duration
	MaxDelay  time.Duration
}

func (j *JitterBackoff) NextBackoff(retryCount int) time.Duration {
	expDelay := j.BaseDelay * (1 << retryCount)
	if expDelay > j.MaxDelay {
		expDelay = j.MaxDelay
	}
	// add jitter (0.5x - 1.5x range)
	jitter := time.Duration(rand.Int63n(int64(expDelay)))
	return expDelay/2 + jitter
}

type Retrier struct {
	strategy   RetryStrategy
	maxRetries int
}

func NewRetrier(strategy RetryStrategy, maxRetries int) *Retrier {
	return &Retrier{
		strategy:   strategy,
		maxRetries: maxRetries,
	}
}

func (r *Retrier) Execute(operation func() error) error {
	var err error

	for attempt := 0; attempt <= r.maxRetries; attempt++ {
		err := operation()
		if err == nil {
			return nil
		}

		if attempt < r.maxRetries {
			backoff := r.strategy.NextBackoff(attempt)
			fmt.Printf("Attempt %d failed: %v. Retrying in %v... \n", attempt+1, err, backoff)
			time.Sleep(backoff)
		}
	}

	return fmt.Errorf("operation failed after %d retries: %w", r.maxRetries, err)
}

// RetryHTTPClient wraps http.Client with retry logic
type RetryHTTPClient struct {
	client  *http.Client
	retrier *Retrier
}

func NewRetryHttpClient(strategy RetryStrategy, maxRetries int) *RetryHTTPClient {
	return &RetryHTTPClient{
		client:  &http.Client{Timeout: 10 * time.Second},
		retrier: NewRetrier(strategy, maxRetries),
	}
}

func (r *RetryHTTPClient) Get(url string) ([]byte, error) {
	var responseBody []byte

	err := r.retrier.Execute(func() error {
		resp, err := r.client.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode >= 500 {
			// retry on 5xx errors
			return fmt.Errorf("server error: %d", resp.StatusCode)
		}

		responseBody, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return nil
	})
	return responseBody, err
}

// Middleware to inject retry-aware HTTP client into Gin context

func RetryClientMiddleware(strategy RetryStrategy, maxRetries int) gin.HandlerFunc {
	return func(c *gin.Context) {
		client := NewRetryHttpClient(strategy, maxRetries)

		// Store client in Gin context
		c.Set("retryClient", client)
		c.Next()
	}
}

func main() {
	// Mock API call that randomly fails
	mockAPICall := func() error {
		if rand.Intn(3) == 0 {
			fmt.Println(" - API call succeeded!")
			return nil
		}
		return fmt.Errorf(" # temporary network error")
	}

	// Try different strategies
	fmt.Println("========== Fixed Delay =========")

	retrier := NewRetrier(&FixedDelay{Delay: 1 * time.Second}, 3)
	_ = retrier.Execute(mockAPICall)

	fmt.Println("============= Exponential Backoff ==========")
	retrier = NewRetrier(&ExponentialBackoff{BaseDelay: 500 * time.Millisecond, MaxDelay: 5 * time.Second}, 3)
	_ = retrier.Execute(mockAPICall)

	fmt.Println("========== Jitter Backoff ============")
	retrier = NewRetrier(&JitterBackoff{BaseDelay: 500 * time.Millisecond, MaxDelay: 5 * time.Second}, 3)
	_ = retrier.Execute((mockAPICall))

	r := gin.Default()

	// Attach middleware with Jitter Backoff strategy
	r.Use(RetryClientMiddleware(&JitterBackoff{
		BaseDelay: 500 * time.Millisecond,
		MaxDelay:  5 * time.Second,
	}, 3))

	// Route: Calls an external API using retry logic
	r.GET("/call-api", func(c *gin.Context) {
		client, exists := c.Get("retryClient")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "retry client not found"})
			return
		}

		retryClient := client.(*RetryHTTPClient)

		// Example API (replace with real microservice endpoint)
		body, err := retryClient.Get("http://httpbin.org/status/500,200")
		// ^ httpbin returns 500 sometimes, 200 sometimes (simulates flaky service)

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "API call succeeded with retry strategy",
			"body":    string(body),
		})
	})

	r.Run(":8080")
}
