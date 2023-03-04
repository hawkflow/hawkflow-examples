package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hawkflow/hawkflow-go"
	"time"
)

func main() {
	// Authenticate with your API key
	hf := hawkflow.New("YOUR_API_KEY", hawkflow.OptionDebug(true))

	// Pass a unique string if you are sending through process and meta that
	// are not unique. this could be for example if you are timing part of a distributed architecture
	// where several instances of the same code is running on many machines or pat of a web app
	// where many users will be hitting the same code at the same time
	uid := uuid.New()

	// Start timing your code - pass through process (required) and meta (optional) parameters
	fmt.Println("Sending timing start data to hawkflow")
	// Processing the err is optional
	err := hf.Start("your_process_name_uid", "your_meta_data", uid.String())
	if err != nil {
		fmt.Println(err)
	}

	// Do some work
	fmt.Println("Sleeping for 5 seconds...")
	time.Sleep(5 * time.Second)

	// End timing this piece of code - process (required) and meta (optional) parameters should match the start
	fmt.Println("Sending timing end data to hawkflow")
	// Processing the err is optional
	err = hf.End("your_process_name_uid", "your_meta_data", uid.String())
	if err != nil {
		fmt.Println(err)
	}
}
