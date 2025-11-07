package main

import (
	"fmt"

	"github.com/fatih/color"
)

type Rating struct {
	ProductId string
	Rating    int
	Comment   []string
	User      string
}

func main() {

	color.Red("Welcome to Product Review System\n")
	// fmt.Println(hello)

	// var (
	// 	productId string
	// 	rating    int
	// 	comment   []string
	// 	user      string
	// )

	// productId = "bottle"
	// rating = 1
	// comment = append(comment, "Good product")
	// comment = append(comment, "Wrong color bottle")
	// user = "Hitesh"

	// fmt.Printf("The rating for product %v given by %v is %v with comments %v \n", productId, user, rating, comment)

	// map1 := make(map[string]int)
	// map1["apple"] = 10
	// map1["banana"] = 20
	// map1["orange"] = 30
	// fmt.Println(map1)

	// v, exist := map1["apple"]
	// if exist {
	// 	fmt.Println("vaue is", v)
	// }
	// v1, exist1 := map1["banana"]
	// if exist1 {
	// 	fmt.Println("vaue is", v1)
	// } else {
	// 	fmt.Println("value is not found")
	// }

	// for key, value := range map1 {
	// 	fmt.Println("key is", key, "value is", value)
	// }

	rating1 := Rating{
		ProductId: "bottle",
		Rating:    1,
		Comment:   []string{"Good product", "Wrong color bottle"},
		User:      "Hitesh",
	}

	rating1.Comment = append(rating1.Comment, "Low price product")

	printProductReview(rating1)

	rating2 := Rating{
		ProductId: "Pen",
		Rating:    4,
		Comment:   []string{"Good product", "Wrong color bottle"},
		User:      "Hitesh",
	}

	printProductReview(rating2)

}

func printMessage(rating Rating) {
	if rating.Rating > 3 {
		fmt.Println("Thank you for your feedback")
	} else {
		fmt.Println("We are sorry to hear that you are not satisfied with our product")
	}
}

func convertRatingToStars(rating int) string {

	stars := ""
	for i := 0; i < rating; i++ {
		stars += "*"
	}
	return stars
}

func printProductReview(rating Rating) {
	fmt.Printf("---------------------------------\n")
	fmt.Printf("The rating for product: %v\n", rating.ProductId)
	fmt.Printf("Rating: %v \n", convertRatingToStars(rating.Rating))
	fmt.Printf("Given By: %v \n", rating.User)
	fmt.Printf("Comments: %v \n", rating.Comment)
	
}
