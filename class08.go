package main

import (
	"fmt"
)

func main() {
		//quiz01()
		//switchExamples()
		//destination()
		//movieSwitch()
		//foodSwitch()
		musicGenres()
}


 func quiz01(){
	
	// Get input from the user or set them programmatically
	var budget int
	var flowerChoice string

	fmt.Print("Enter your budget: ")
	fmt.Scan(&budget)

	fmt.Print("Enter your flower choice: ")
	fmt.Scan(&flowerChoice)

	// Conditions 
	if budget < 1000 {
		fmt.Println("Your budget is too low for the mentioned types of flowers. Consider opting for more budget-friendly options.")
	} else if budget >= 1000 && budget < 5000 {
		switch flowerChoice {
		case "roses", "tulips", "lilies":
			fmt.Println("With your budget, you can afford beautiful flowers. Enjoy your choice!")
		case "exotic orchids":
			fmt.Println("With your budget, you can go for exotic orchids. They'll make a unique and impressive gift!")
		default:
			fmt.Println("You have enough for a wonderful choice of flowers, but the type is unspecified. Choose your flowers and enjoy!")
		}
	} else if budget >= 5000 && flowerChoice == "luxury flower arrangement" {
		fmt.Println("You can afford a luxury flower arrangement for a truly special occasion. Make it memorable!")
	} else {
		fmt.Println("Invalid input. Please check your budget and flower choice.")
	}
}

 
func switchExamples() {
	var iceCream string

	fmt.Println("Please insert your prefered ice cream: ")
	fmt.Scan(&iceCream)

	switch iceCream {
	case "vanilla":
	 fmt.Println("You selected vanilla ice cream.")
	case "chocolate":
	 fmt.Println("You selected chocolate ice cream.")
	case "strawberry":
	 fmt.Println("You selected strawberry ice cream.")
	case "mint":
	 fmt.Println("You selected mint ice cream.")
	default:
	 fmt.Println("We don't have that flavor. Please choose another.")
 }
}

func destination() {

	destination := "beach"
 
	switch destination {
	case "beach":
	 fmt.Println("You chose a beach vacation")
	case "mountain":
	 fmt.Println("You chose a mountain vacation")
	case "city":
	 fmt.Println("You chose a city vacation")
	case "countryside":
	 fmt.Println("You chose a countryside")
	default:
	 fmt.Println("That destination is not available. Please pick another")
	}
 
 }

 func movieSwitch() {

	genre := "action"
 
	switch genre {
	case "action":
	 fmt.Println("You selected an action movie")
	case "comedy":
	 fmt.Println("You selected a comedy movie")
	case "drama":
	 fmt.Println("You selected a dramatic movie")
	case "sci-fi":
	 fmt.Println("You selected a scifi movie")
	default:
	 fmt.Println("We don't have movies of that genre. Please choose another")
 
	}
 
 }


 func petSwitch() {

	pet := "dog"
 
	switch pet {
 
	case "dog":
	 fmt.Println("Sorry to hear that you selected a dog as your pet")
	case "cat":
	 fmt.Println("Congratulations, you selected a cat as your pet")
	case "fish":
	 fmt.Println("You selected a fish as your pet")
	case "hamster":
	 fmt.Println("You selected a hamster as your pet")
	default:
	 fmt.Println("WE DON'T EVEN HAVE THAT CREATURE ON EARTHHH!!!")
 
	}
 }

 func foodSwitch() {

	food := "brazilian bbq"
 
	switch food {
 
	case "pizza":
	 fmt.Println("You chose pizza as your meal")
	case "burger":
	 fmt.Println("You chose burger as your meal")
	case "sushi":
	 fmt.Println("You chose sushi as your meal")
	case "brazilian bbq":
	 fmt.Println("You chose brazilian bbq as your meal")
	default:
 
	 fmt.Println("You ate chips")
 
	}
 
 }

 func musicGenres(){
	var musicGenre string
	
	
	fmt.Println("Please insert your prefered music among these: pop, hip-hop, rock, jazz")
	fmt.Scan(&musicGenre)

	switch musicGenre {
		
	case "pop":
		fmt.Println("You seleceted pop music")		
	case "hip-hop":
		fmt.Println("You seleceted hip-hop music")
	case "rock":
		fmt.Println("You seleceted rock music")
	case "jazz":
		fmt.Println("You seleceted jazz music")
	default:
		fmt.Println("we don't have music of that genre available. Please choose another genre.")

	}

 }