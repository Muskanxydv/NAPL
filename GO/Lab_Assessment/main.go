package main

import "fmt"

type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    int
}

func (i *Item) TotalValue() int {
	return (i.Quantity) * i.Price
}

func (i *Item) Itemcount() int {
	return i.Quantity
}

func (i *Item) UpdateQuantity(newQuantity int) error {
	if newQuantity < 0 {
		return fmt.Errorf("Quantity of items cannot be negative!! Enter a positive value")
	}
	i.Quantity = newQuantity
	return nil
}

func (i *Item) RemoveItem(items []Item) ([]Item, error) {
	for index, item := range items {
		if item.ID == i.ID {
			return append(items[:index], items[index+1:]...), nil
		}
	}
	return items, fmt.Errorf("item cannot be found in the available items!! ENter a valid Item")
}

func (i *Item) AddPriceofallgivenitems(newPrice int) error {
	if newPrice <= 0 {
		return fmt.Errorf("Price of items cannot be zero or negative!!! Enter a positive value")
	}
	i.Price = newPrice
	return nil
}

func main() {
	item := Item{ID: 1, Name: "iPhone 17", Quantity: 3, Price: 130000}
	fmt.Println("Total Value:", item.TotalValue())
	fmt.Println("Item Count:", item.Itemcount())
	err := item.UpdateQuantity(5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Updated Quantity:", item.Quantity)
	}
}
