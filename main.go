package main

import "fmt"

func main() {
	fmt.Println("Dominus Iesus Christus")
	fmt.Println("----------------------")
	fmt.Println()

	key := "cache_key"
	best_fusion := WithCache(key, func() any {

	})
}
