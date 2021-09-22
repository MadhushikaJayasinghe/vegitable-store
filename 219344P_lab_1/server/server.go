package main

import (
	"log"
	"net"
	"net/http"
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

var database []Item

type API int

func (a *API) GetItemList(req string, reply *[]string) error {
	var vegList []string
	for idx := range database {
		vegList = append(vegList, database[idx].Name)
	}
	*reply = vegList
	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	var isAlreadyExist = false
	for _, val := range database {
		if item.Name == val.Name{
			isAlreadyExist = true
		}
	}
	if !isAlreadyExist {
		database = append(database, item)
		*reply = item
	}
	return nil
}

func (a *API) UpdateItem(item Item, reply *Item) error {
	var updatedItem Item
	for idx, val := range database {
		if val.Name == item.Name {
			database[idx] = item
			updatedItem = item
		}
	}
	*reply = updatedItem
	return nil
}

func (a *API) GetPriceByName(name string, reply *ItemNamePrice) error {
	var itemNamePrice ItemNamePrice

	for idx, val := range database {
		if val.Name == name {
			itemNamePrice.Price = database[idx].Price
			itemNamePrice.Name = name
		}
	}

	*reply = itemNamePrice
	return nil
}

func (a *API) GetAmountByName(name string, reply *ItemNameAmount) error {
	var itemNameAmount ItemNameAmount

	for idx, val := range database {
		if val.Name == name {
			itemNameAmount.Amount = database[idx].Amount
			itemNameAmount.Name = name
		}
	}

	*reply = itemNameAmount
	return nil
}

func main() {
	var api = new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API ", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("error listening API ", err)
	}
	log.Printf("serving rpc on port %d", 4040)

	err = http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving: ", err)
	}

}
