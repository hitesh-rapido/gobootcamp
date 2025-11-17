package main

import (
	"fmt"
	"hello_world/rating"

	"github.com/fatih/color"
)

type CheckResult struct {
	url    string
	status string
}

func main() {

	color.Red("Welcome to Product Review System\n")

	productRatings := make(map[string]*rating.Rating)

	for {
		fmt.Println("Select the option: 1. Add Rating 2. View Rating 3. Exit")
		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
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
			r := productRatings[productName]
			if r == nil {
				r = &rating.Rating{}
			}
			e := r.AddRating(userName, ratingValue, comment)
			if e != nil {
				fmt.Println(e)
			}
			r.AverageRating()
			productRatings[productName] = r

		case 2:
			fmt.Println("Enter the product name:")
			var productName string
			fmt.Scanln(&productName)
			r := productRatings[productName]
			if r == nil {
				fmt.Println("Product not found")

			} else {
				fmt.Println(r)
			}

		case 3:
			return
		default:
			fmt.Println("Invalid option")
		}
	}

}
