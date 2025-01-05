package main

import "fmt"

func main() {
	fmt.Println("###### Welcome to our Todolist APP! ######")

	var shortGolang = "Watch Go crash course"
	var fullGolang = "Watch Nana's Golang Full Course"
	var rewardDessert = "Reward myself with a cheesecake"

	var taskItems = []string{shortGolang, fullGolang, rewardDessert}

	printTasks(taskItems)
}

func printTasks(taskItems []string) {
	fmt.Println("List of my Todos")

	for index, task := range taskItems {
		fmt.Printf("%d. %s\n", index+1, task)
	}
}
