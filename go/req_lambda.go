package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

func req() string {

	var url = "http://google.com"

	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := netClient.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	return string(body)

}

func handleRequest(ctx context.Context) (string, error) {

	return req(), nil
}

func main() {
	lambda.Start(handleRequest)
}
