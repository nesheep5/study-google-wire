package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	ctx := context.Background()
	baz, err := initializeBaz(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("baz.X = %+v\n", baz.X)
}
