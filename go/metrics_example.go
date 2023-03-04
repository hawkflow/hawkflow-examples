package main

import (
	"fmt"
	"github.com/hawkflow/hawkflow-go"
)

func main() {
	// Authenticate with your API key
	hf := hawkflow.New("YOUR_API_KEY")

	items := map[string]float64{
		"users_count":     5465,
		"model_accuracy":  0.85,
		"rows_entered_db": 4654465,
	}

	fmt.Println("Sending metrics data to hawkflow")
	// Processing the response is optional
	err := hf.Metrics("hawkflow_examples", "your_meta_data", items)
	if err != nil {
		fmt.Println(err)
	}
}
