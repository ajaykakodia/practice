package main

import (
	"context"
	"fmt"
	"time"
)

type cmd struct {
	op   string
	key  string
	resp chan string // or error/ result type
}

var inbox = make(chan cmd, 1024)

func runOwner(ctx context.Context, inbox <-chan cmd) {
	store := make(map[string]string)
	for {
		select {
		case <-ctx.Done():
			return
		case c := <-inbox:
			switch c.op {
			case "get":
				c.resp <- store[c.key]
			case "set":
				store[c.key] = <-c.resp
				close(c.resp) // ack
			}
		}
	}
}

func Get(k string) string {
	r := make(chan string, 1)
	inbox <- cmd{op: "get", key: k, resp: r}
	return <-r
}

func Set(k, v string) {
	r := make(chan string)
	inbox <- cmd{op: "set", key: k, resp: r}
	r <- v
	if _, ok := <-r; !ok {
		fmt.Println("Value is set")
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go runOwner(ctx, inbox)

	Set("Ajay", "Yadav")
	Set("Ram", "Ji")

	val := Get("Ram")
	fmt.Println(val)

	time.AfterFunc(1*time.Second, func() {
		cancel()
		close(inbox)
	})
}
