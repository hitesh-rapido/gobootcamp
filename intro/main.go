package main

import (
	"fmt"
	"hello_world/rating"
	"github.com/fatih/color"
)

func main() {

	color.Red("Welcome to Product Review System\n")

	// Initialize a rating aggregate and add entries
	mapRating := &rating.Rating{}
	e :=mapRating.AddRating("Hitesh", 1, "Good product")
	if(e!=nil){
		fmt.Println(e)
	}
	e = mapRating.AddRating("Gaurav", 3, "Good product")
	if(e!=nil){
		fmt.Println(e)
	}

	fmt.Println(mapRating)

}


