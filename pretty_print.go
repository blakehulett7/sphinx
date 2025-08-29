package main

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(text string, data interface{}) {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return
	}

	fmt.Print("\n-----\n\n")
	fmt.Printf("%s: %s", text, string(b))
	fmt.Print("\n\n-----\n")
}
