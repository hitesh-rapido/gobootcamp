package main

import (
	"fmt"
	"hello_world/rating"

	"github.com/fatih/color"
)

func main() {

	color.Red("Welcome to Product Review System\n")

	productRatings := make(map[string]*rating.Rating)
	//mapRating := &rating.Rating{}

	// Initialize a rating aggregate and add entries

	// e := mapRating.AddRating("Hitesh", 1, "Good product")

	// e = mapRating.AddRating("Gaurav", 3, "Good product")
	// if e != nil {
	// 	fmt.Println(e)
	// }
	// mapRating.AverageRating()
	// fmt.Println(mapRating)

	for {
		fmt.Println("Select the option: 1. Add Rating 2. View Rating 3. Exit")
		var option int
		fmt.Scanln(&option)
		if option == 1 {

			fmt.Println("Enter the product name:")
			var productName string
			fmt.Scanln(&productName)

			fmt.Println("Enter the user name:")
			var userName string
			fmt.Scanln(&userName)

			fmt.Println("Enter the rating:")
			var ratingValue int
			fmt.Scanln(&ratingValue)

			fmt.Println("Enter the comment:")
			var comment string
			fmt.Scanln(&comment)

			if productRatings[productName] == nil {
				productRatings[productName] = &rating.Rating{}
			}
			e := productRatings[productName].AddRating(userName, ratingValue, comment)
			if e != nil {
				fmt.Println(e)
			}
		}
		if option == 2 {
			fmt.Println("Enter the product name:")
			var productName string
			fmt.Scanln(&productName)
			if productRatings[productName] == nil {
				fmt.Println("Product not found")

			} else {
				fmt.Println(productRatings[productName])
			}
		}
		if option == 3 {
			break
		}
	}

}
