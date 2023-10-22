package gocontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
)

func TestContext(t *testing.T)  {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")
	contextF := context.WithValue(contextC, "f", "F")


	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))
}

func CreateCounter() chan int  {
	destination := make(chan int)

	go func()  {
		defer close(destination)
		counter := 1

		for {
			destination <- counter
			counter++
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T)  {
	fmt.Println("Total goroutine", runtime.NumGoroutine())	

	destination := CreateCounter()

	for n := range destination{
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Total goroutine", runtime.NumGoroutine())	

}