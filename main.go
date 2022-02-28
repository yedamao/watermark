package main

import (
	"context"
	"fmt"
	"log"
	"os"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/kelseyhightower/envconfig"
)

type envConfig struct {
	S3_BUCKET            string `envconfig:"S3_BUCKET" default:"us-east-1"`
	S3_REGION            string `envconfig:"S3_REGION" default:"damao-workspace"`
	S3_ACCESS_KEY_ID     string `envconfig:"S3_ACCESS_KEY_ID"`
	S3_SECRET_ACCESS_KEY string `envconfig:"S3_SECRET_ACCESS_KEY"`
}

var (
	env envConfig
)

func init() {
	if err := envconfig.Process("", &env); err != nil {
		log.Printf("[ERROR] Failed to process env var: %s", err)
		os.Exit(1)
	}
}

type Input struct {
	Url     string `json:"url"`
	Comment string `json:"comment"`
}

func receive(event cloudevents.Event) {
	input := &Input{}
	if err := event.DataAs(input); err != nil {
		fmt.Println("Input parse error: ", err)
		return
	}
	log.Printf("CloudEvent:\n%s", event)

	fmt.Printf("input %s\n", input)

	data, err := downloadImage(input.Url)
	if err != nil {
		log.Printf("download image error: %v\n", err)
		return
	}

	processedBlob, err := addComment(data, input.Comment)
	if err != nil {
		log.Printf("add comment error: %v\n", err)
		return
	}

	if err := uploadImage(processedBlob); err != nil {
		log.Printf("upload image: %v\n", err)
		return
	}

	fmt.Printf("%s processed successfully\n", input.Url)
}

func main() {
	// The default client is HTTP.
	c, err := cloudevents.NewClientHTTP()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}

	log.Printf("listening on port %d\n", 8080)
	log.Fatal(c.StartReceiver(context.Background(), receive))
}
