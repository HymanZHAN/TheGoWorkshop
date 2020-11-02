package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zipcode int    `json:"zipcode"`
}

type item struct {
	Name        string `json:"itemname"`
	Description string `json:"desc,omitempty"`
	Quantity    int    `json:"qty"`
	Price       int    `json:"price"`
}

type order struct {
	TotalPrice  int  `json:"total"`
	IsPaid      bool `json:"paid,omitempty"`
	Fragile     bool
	OrderDetail []item `json:"orderdetail"`
}

type customer struct {
	UserName      string  `json:"username"`
	Password      string  `json:"-"`
	Token         string  `json:"-"`
	ShipTo        address `json:"shipto"`
	PurchaseOrder order   `json:"order"`
}

func main() {
	jsonData := []byte(` 
	{ 
	  "username" :"blackhat", 
	  "shipto": 
	  { 
		"street": "Sulphur Springs Rd", 
		"city": "Park City", 
		"state": "VA", 
		"zipcode": 12345 
	  }, 
	  "order": 
	  { 
		"paid":false, 
		"orderdetail" : 
		[{ 
		  "itemname":"A Guide to the World of zeros and ones", 
		  "desc": "book", 
		  "qty": 3, 
		  "price": 50 
		}] 
	  } 
	} 
	`)

	if !json.Valid(jsonData) {
		fmt.Printf("json is not valid: %s", jsonData)
		os.Exit(1)
	}

	var c customer
	err := json.Unmarshal(jsonData, &c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	newItem1 := item{
		Name:        "Final Fantasy The Zodiac Age",
		Description: "Nintendo Switch Game",
		Quantity:    1,
		Price:       50,
	}

	newItem2 := item{
		Name:     "Crystal Drinking Glass",
		Quantity: 11,
		Price:    25,
	}

	c.PurchaseOrder.OrderDetail = append(c.PurchaseOrder.OrderDetail, newItem1, newItem2)

	var totalPrice int
	for _, item := range c.PurchaseOrder.OrderDetail {
		totalPrice += item.Price * item.Quantity
	}
	c.PurchaseOrder.TotalPrice = totalPrice

	result, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s", result)

}
