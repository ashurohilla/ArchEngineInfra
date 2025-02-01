package main

import "fmt"

func divider(numerator int, denominator int) (int, string) {

	var result = numerator / denominator

	var ashish = "ashish"

	return result, ashish

}

func main() {

	var result, ashish = divider(12, 2)

	fmt.Println(result, ashish)

}
