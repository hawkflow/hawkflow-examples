package main

import (
	"fmt"
	"github.com/hawkflow/hawkflow-go"
	"time"
)

func main() {
	// Authenticate with your API key
	hf := hawkflow.New("YOUR_API_KEY")

	// Start timing your code - pass through process (required) and meta (optional) parameters
	fmt.Println("Sending timing start data to hawkflow")
	err := hf.Start("hawkflow_examples", "your_meta_data", "")
	if err != nil {
		fmt.Println(err)
	}

	// Do some work
	fmt.Println("Sleeping for 5 seconds...")
	time.Sleep(5 * time.Second)

	// End timing this piece of code - process (required) and meta (optional) parameters should match the start
	fmt.Println("Sending timing end data to hawkflow")
	err = hf.End("hawkflow_examples", "your_meta_data", "")
	if err != nil {
		fmt.Println(err)
	}
}
