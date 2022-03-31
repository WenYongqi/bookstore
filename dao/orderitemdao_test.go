package dao

import (
	"bookstore/model"
	"fmt"
	"testing"
)

func TestAddOrderItem(t *testing.T) {
	err := AddOrderItem(&model.OrderItem{
		Count: 2,
		Amount: 1.2,
		Title: "123",
		Author: "234",
		Price: 12.3,
		ImgPath: "111",
		OrderID: "7641d562-8f54-46eb-5ce6-d32e41c1c542",
	})
	fmt.Println(err)
}