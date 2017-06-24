package main

import (
	"gochapter/config"
	"gochapter/places"
	"fmt"
)

func main() {
	cmdConfig, _ := config.Parse()

	matrix, err := places.GetMappedDistanceMatrix(cmdConfig.Arguments)

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