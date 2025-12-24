package main

import "fmt"

func main() {
	var score int32 = 60
	if score >= 90 {
		fmt.Printf("score value is %d,so it's A", score)
	} else if score >= 80 {
		fmt.Printf("score value is %d,so it's B", score)
	} else {
		fmt.Printf("score value is %d,so it's C", score)
	}
}
