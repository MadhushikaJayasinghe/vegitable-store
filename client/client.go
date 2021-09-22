package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Name   string
	Price  float32
	Amount int
}

type ItemNamePrice struct {
	Name  string
	Price float32
}

type ItemNameAmount struct {
	Name   string
	Amount int
}


func main() {
	var reply Item
	var replyItemPrice ItemNamePrice
	var replyItemAmount ItemNameAmount
	var db []string

	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("connection error : ", err)
	}

	item1 := Item{"Carrot", 75.6, 10}
	item2 := Item{"Cucumber", 120, 8}
	item3 := Item{"Tomato", 45, 6}

	client.Call("API.AddItem", item1, &reply)

	fmt.Println("added item : ", reply)

	client.Call("API.AddItem", item2, &reply)

	fmt.Println("added item : ", reply)

	client.Call("API.AddItem", item3, &reply)

	fmt.Println("added item : ", reply)

	client.Call("API.GetItemList" ,"" , &db)

	fmt.Println("vegetable list : ", db)

	client.Call("API.GetPriceByName", "Tomato", &replyItemPrice)

	fmt.Println("price of ", replyItemPrice.Name , " is Rs.", replyItemPrice.Price)

	client.Call("API.GetAmountByName", "Tomato", &replyItemAmount)

	fmt.Println("available amount of ", replyItemAmount.Name , " is ", replyItemAmount.Amount , "kg")

	item4 := Item{"Tomato", 80, 3}

	client.Call("API.UpdateItem",  item4, &reply)

	fmt.Println("updated item : ", reply)

	client.Call("API.GetPriceByName", "Tomato", &replyItemPrice)

	fmt.Println("price of ", replyItemPrice.Name , " is Rs.", replyItemPrice.Price)

	client.Call("API.GetAmountByName", "Tomato", &replyItemAmount)

	fmt.Println("available amount of ", replyItemAmount.Name , " is ", replyItemAmount.Amount , "kg")

}