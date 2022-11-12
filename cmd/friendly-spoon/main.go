package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	log.Println("starting app...")

	<-ctx.Done()
	log.Println("closing app")
}
