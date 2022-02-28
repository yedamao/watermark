// Port of http://members.shaw.ca/el.supremo/MagickWand/resize.htm to Go
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func downloadImage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	blob, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return blob, nil
}

func uploadImage() {
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
		log.Printf("download image error: \n", err)
		return
	}

	processed, err := addComment(data, input.Comment)
	if err != nil {
		log.Printf("add comment error: \n", err)
		return
	}

	_ = processed

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
