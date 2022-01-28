package main

import (
	"fmt"
	"mydict/mydict"
)
func main() {
	dictionary := mydict.Dictionary{}

	// ===== Search =====
	// definition, err := dictionary.Search("first")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }

	// ===== Add =====
	// word := "hello"
	// definition := "Greeting"
	// err := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// hello, _ := dictionary.Search("hello")
	// fmt.Println(dictionary, hello)

	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	// ===== Update =====
	// baseWord := "hello"
	// dictionary.Add(baseWord, "First")
	// err := dictionary.Update(baseWord, "Second")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// word, _ := dictionary.Search(baseWord)
	// fmt.Println(word)

	baseword := "hello"
	dictionary.Add(baseword, "First")
	dictionary.Search(baseword)
	check := dictionary.Delete(baseword)
	fmt.Println(check)
	word, err := dictionary.Search(baseword)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)
}