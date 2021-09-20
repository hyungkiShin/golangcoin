package main

import "fmt"

func main() {
	foods := []string{"potato", "pizza", "pasta"}
	fmt.Printf("%v\n", foods)

	foodsNew := append(foods, "tomato") // 복사본에 할당
	fmt.Printf("%v\n", foodsNew)

	fmt.Println(len(foodsNew)) // 길이.
}
