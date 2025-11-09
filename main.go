package main

import (
	"fmt"

	"github.com/fatih/color"
)

type Review struct {
	Name    string
	Rating  int
	Comment string
}

func getUserInput() (string, int, string) {
	var rating int
	var fristName string
	var comment string
	fmt.Println("Welcome to the Rating App")
	fmt.Println("What is your name?")
	fmt.Scanln(&fristName)
	fmt.Println("Provide a rating between 1 and 5:")
	fmt.Scanln(&rating)
	fmt.Println("Any comments?")
	fmt.Scanln(&comment)

	return fristName, rating, comment
}

func displayRatingStars(rating int) {
	for i := 0; i < rating; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func displayAcknowledgementMessage(fristName string, rating int, comment string) {
	if rating >= 3 && rating <= 5 {
		color.Green("%v, Thank you for the positive feedback! \n", fristName)
		color.Green("We are glad you had a %v experience", comment)
		displayRatingStars(rating)
	} else if rating >= 1 && rating <= 2 {
		color.Red("We appreciate your feedback %v, sorry for bad exprience !", fristName)
		displayRatingStars(rating)
	} else {
		color.Yellow("Invalid rating. Please provide a rating between 1 and 5.")
	}

	fmt.Println("----- Review Summary -----")

	var review Review
	review.Name = fristName
	review.Rating = rating
	review.Comment = comment

	fmt.Printf("Name: %s\n", review.Name)
	fmt.Printf("Rating: %d\n", review.Rating)
	fmt.Printf("Comment: %s\n", review.Comment)

}

func main() {
	// 	fmt.Println("Hello, World!")
	// 	color.Red("Hello, World!")
	// 	color.Green("Hello, World!")

	// 	var test string = "Line01 \nLine02 \nLine03"
	// 	fmt.Println(test)

	//	var newLine string = `Line11
	//
	// Line12
	// Line13
	// Blackslash is \n literal here.`
	//
	//	fmt.Println(newLine)

	fristName, rating, comment := getUserInput()
	displayAcknowledgementMessage(fristName, rating, comment)

	// switch rating {
	// case 1:
	// 	fmt.Println("*")
	// case 2:
	// 	fmt.Println("**")
	// case 3:
	// 	fmt.Println("***")
	// case 4:
	// 	fmt.Println("****")
	// case 5:
	// 	fmt.Println("*****")
	// default:
	// 	fmt.Println("Invalid rating stars to display.")
	// }

	// array example//

	// var num [3]int
	// num[0] = 10
	// num[1] = 20
	// num[2] = 30
	// fmt.Println(num)
	// fmt.Println(len(num))
	// fmt.Println(cap(num))

	// slice example//

	// slice1 := make([]int, 4)
	// slice1[0] = 10
	// slice1[1] = 20
	// slice1[2] = 30
	// slice1 = append(slice1, 40)
	// fmt.Println(slice1)
	// fmt.Println(len(slice1))
	// fmt.Println(cap(slice1))

	// var (
	// 	name     []string
	// 	rating   []int
	// 	comments []string
	// )

	// fmt.Println("Rating Screen")
	// name = append(name, "Jack")
	// comments = append(comments, "Good")
	// rating = append(rating, 5)

	// name = append(name, "Tack")
	// comments = append(comments, "Bad")
	// rating = append(rating, 1)

	// name = append(name, "Mack")
	// comments = append(comments, "Great")
	// rating = append(rating, 4)

	// fmt.Println("Name - Rating - Comments")
	// fmt.Printf("%s %v %s", name, rating, comments)

	// map example//

	// map1 := make(map[string]int)
	// map1["Jack"] = 5
	// map1["Tack"] = 1
	// map1["Mack"] = 4

	// fmt.Println("Map Data")
	// fmt.Println(map1)

	// name, exists := map1["Tack"]
	// if exists {
	// 	fmt.Println("Name:", name)
	// } else {
	// 	fmt.Println("Name not found")
	// }

	// for key, value := range map1 {
	// 	fmt.Printf("Key: %s, Value: %d\n", key, value)
	// }

	// structs example //

	// var reviews []Review

	// reviews = append(reviews, Review{Name: "Jack", Rating: 5, Comment: "Good"})
	// reviews = append(reviews, Review{Name: "Tack", Rating: 1, Comment: "Bad"})
	// reviews = append(reviews, Review{Name: "Mack", Rating: 4, Comment: "Great"})

	// for _, review := range reviews {
	// 	fmt.Printf("Name: %s, Rating: %d, Comment: %s\n", review.Name, review.Rating, review.Comment)
	// }

}
