package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func calculate(n int64) int64 {
	var (
		a   int64 = 1
		b   int64 = 0
		tmp int64 = 0
	)

	for n > 0 {
		tmp = a
		a = a + b
		b = tmp
		n--
	}
	return b
}

func handleRequest(ctx context.Context, n int64) (int64, error) {

	return calculate(n), nil
}

func main() {
	lambda.Start(handleRequest)
}
