package main

import (
	"fmt"
	"flag"
	"gochapter/places"
)

func main() {
	flag.Parse()

	matrix, err := places.GetMappedDistanceMatrix(flag.Args())

	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	for origin, distances := range matrix {
		fmt.Println(origin)

		for destination, distance := range distances {
			fmt.Printf("  %s: %s\n", destination, distance.HumanReadable)
		}
	}
}