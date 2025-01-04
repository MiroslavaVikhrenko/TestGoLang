package main

import "fmt"

func main() {
	var shortGolang = "Watch Go crash course"
	var fullGolang = "Watch Nana's Golang Full Course"
	var rewardDessert = "Reward myself with a cheesecake"

	taskItems := []string{shortGolang, fullGolang, rewardDessert}

	fmt.Println("###### Welcome to our Todolist APP! ######")

	fmt.Println("List of my Todos")
	fmt.Println("Tasks:", taskItems)
}
