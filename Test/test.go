package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	person := []*Person{
		// {Name: "a"},
		// {Name: "b"},
		// {Name: "h"},
		// {Name: "d"},
		// {Name: "e"},
	}

	// Convert the object to an array
	arr := []string{}
	for _, item := range person {
		arr = append(arr, item.Name)
	}

	fmt.Println(isUniqueArray(arr))
}

func isUniqueArray(arr []string) bool {
	m := make(map[string]struct{})

	for _, v := range arr {
		m[v] = struct{}{}
	}

	return len(m) == len(arr)
}
