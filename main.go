package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

type groceryList map[string][]string


func (gl groceryList) AddItem(category, item string) {
	category = strings.ToLower(category)
	item = strings.Title(item)
	gl[category] = append(gl[category], item)
	sort.Strings(gl[category]) 
	fmt.Printf("Added '%s' to category '%s'.\n", item, category)
}

func (gl groceryList) ListItems(category string) error {
	category = strings.ToLower(category)
	items, exists := gl[category]
	if !exists || len(items) == 0 {
		return errors.New("no items found in this category")
	}

	fmt.Printf("Items in category '%s':\n", category)
	for _, item := range items {
		fmt.Printf("- %s\n", item)
	}
	return nil
}

func (gl groceryList) DeleteItem(category, item string) error {
	category = strings.ToLower(category)
	item = strings.Title(item)

	items, exists := gl[category]
	if !exists {
		return errors.New("category does not exist")
	}

	index := -1
	for i, v := range items {
		if v == item {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("item not found in this category")
	}

	gl[category] = append(items[:index], items[index+1:]...)
	fmt.Printf("Deleted '%s' from category '%s'.\n", item, category)
	return nil
}

func (gl groceryList) ListAllCategories() {
	fmt.Println("Grocery List:")
	for category, items := range gl {
		fmt.Printf("Category: %s\n", category)
		for _, item := range items {
			fmt.Printf("- %s\n", item)
		}
	}
}

func menu() {
	fmt.Println("\nGrocery List Organizer")
	fmt.Println("1. Add Item")
	fmt.Println("2. List Items by Category")
	fmt.Println("3. Delete Item")
	fmt.Println("4. List All Categories")
	fmt.Println("5. Exit")
	fmt.Print("Choose an option: ")
}

func main() {
	gl := groceryList{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		menu()
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter category: ")
			scanner.Scan()
			category := scanner.Text()

			fmt.Print("Enter item name: ")
			scanner.Scan()
			item := scanner.Text()

			gl.AddItem(category, item)

		case "2":
			fmt.Print("Enter category to list: ")
			scanner.Scan()
			category := scanner.Text()

			if err := gl.ListItems(category); err != nil {
				fmt.Println("Error:", err)
			}

		case "3":
			fmt.Print("Enter category: ")
			scanner.Scan()
			category := scanner.Text()

			fmt.Print("Enter item name to delete: ")
			scanner.Scan()
			item := scanner.Text()

			if err := gl.DeleteItem(category, item); err != nil {
				fmt.Println("Error:", err)
			}

		case "4":
			gl.ListAllCategories()

		case "5":
			fmt.Println("Exiting Grocery List Organizer.")
			return

		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
