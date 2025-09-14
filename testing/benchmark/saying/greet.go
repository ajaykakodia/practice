package saying

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("Welcome my dear %s", name)
}

func GreetT[T string | int](name T) string {
	return fmt.Sprintf("Welcome my dear %v", name)
}
