package change

import "fmt"

func ExampleChange() {
	coins := []int{1, 5, 10}
	target := 25
	fmt.Println(Change(coins, target))
	// Output:
	// [5 10 10] <nil>
}

func ExampleChange40() {
	coins := []int{1, 5, 10, 25, 100}
	target := 40
	fmt.Println(Change(coins, target))
	// Output:
	// [5 10 25] <nil>
}
