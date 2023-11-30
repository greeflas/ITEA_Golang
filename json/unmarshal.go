package main

import (
	"encoding/json"
	"fmt"
)

type refund struct {
	Id     string `json:"id"`
	Amount int    `json:"amount"`
	Status string `json:"operation_status"`
}

func main() {
	refundJsonBytes := []byte(`{
	"id": "a8cbd2be-126d-4ad7-9749-143f0ffa09ec",
	"amount": 500,
	"operation_status": "refunded"
}`)

	// unknown JSON structure
	var data map[string]any

	if err := json.Unmarshal(refundJsonBytes, &data); err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(data)

	status, ok := data["status"].(string)
	if !ok {
		fmt.Println("Status is not string")
	}

	fmt.Println(status)

	// known JSON structure
	r := new(refund)
	if err := json.Unmarshal(refundJsonBytes, r); err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(r)
}
