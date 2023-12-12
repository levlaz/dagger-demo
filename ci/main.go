package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()

	// init dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// use golang container, get version, execute
	golang := client.Container().
		From("golang:latest").
		WithExec([]string{"go", "version"})

	version, err := golang.Stdout(ctx)
	if err != nil {
		panic(err)
	}

	// print output
	fmt.Println("Hello from Dagger and " + version)
}
