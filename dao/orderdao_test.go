package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
	"testing"
)

func TestAddOrder(t *testing.T) {
	err := AddOrder(&model.Order{
		ID: utils.CreateUUID(),
		TotalCount: 2,
		TotalAmount: 1.2,
		State: 1,
		UserID: 1,
	})
	fmt.Println(err)
}

func TestGetOrders(t *testing.T) {
	orders, err := GetOrders()
	for _, order := range orders {
		fmt.Println(order)
	}
	fmt.Println(err)
}