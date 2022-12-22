package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Outside a goroutine")

	go func() {
		fmt.Println("Inside a goroutine")
	}()

	fmt.Println("Outside again")

	defer func() {
		fmt.Print("adfsdfdasfds")
	}()
	<-quit
}
