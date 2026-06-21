package main

import (
	"context"
	"fmt"
	"log"

	temporalclient "go.temporal.io/sdk/client"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	client, err := temporalclient.Dial(temporalclient.Options{})
	if err != nil {
		return err
	}
	fmt.Println(client)
	return nil
}
