package main

import (
	"fmt"
	"log"

	"github.com/markstgodard/intro/circle"
)

func main() {
	area, err := circle.Area(-3.0) // HL
	if err != nil {                // HL
		log.Fatalf("error: %s\n", err)
	} // HL

	fmt.Printf("area of circle: %f\n", area)
}
