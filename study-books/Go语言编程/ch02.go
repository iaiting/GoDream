package main

import "fmt"

func c0201() {
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var mySlice []int = myArray[:5]
	for i, v := range myArray {
		fmt.Println(i, ":", v)
	}

	fmt.Println("--------------------")
	for i, v := range mySlice {
		fmt.Println(i, ":", v)
	}
}

func c0202() {

	mySlice := make([]int, 5, 10)
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))

}

func c0203() {
	type PersonInfo struct {
		ID      string
		Name    string
		Address string
	}
	// var personDB map[string]PersonInfo
	personDB := make(map[string]PersonInfo)
	personDB["123"] = PersonInfo{"123", "tom", "mas"}
	personDB["1"] = PersonInfo{"1", "Jack", "nj"}

	person, ok := personDB["123"]
	if ok {
		fmt.Println("found ok", person)
	} else {
		fmt.Println("found not ok")
	}
}

////////////////////////////////////////////////////////////////////////////////
func main() {
	c0201()
	c0202()
	c0203()
}
