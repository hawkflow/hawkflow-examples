package main

import (
	"errors"
	"fmt"
	"github.com/hawkflow/hawkflow-go"
)

func main() {
	// Authenticate with your API key
	hf := hawkflow.New("YOUR_API_KEY", hawkflow.OptionDebug(true))

	err := fail()
	if err != nil {
		// catch and send errors through to hawkflow
		fmt.Println("Sending error data to hawkflow")
		err := hf.Exception("hawkflow_examples", "your_meta_data", err.Error())
		if err != nil {
			fmt.Println(err)
		}
	}
}

func fail() error {
	return errors.New("your error message")
}
