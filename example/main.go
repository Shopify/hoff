package main

import (
	"fmt"

	"github.com/Shopify/hoff"
)

func main() {
	filtered := hoff.Filter([]int{1, 2, 3}, func(i int) bool { return i < 3 })
	fmt.Println(filtered)
}
