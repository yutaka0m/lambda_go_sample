package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Name string `json:"name"`
}

func handler(ctx context.Context, name Event) (string, error) {
	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
	lambda.Start(handler)
}
