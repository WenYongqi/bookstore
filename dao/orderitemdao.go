package dao

import (
	"bookstore/model"
	"log"
)

func AddOrderItem(orderItem *model.OrderItem) error {
	err := db.Create(orderItem).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetOrderItemsByOrderID(orderID string) ([]*model.OrderItem, error) {
	orderItems := []*model.OrderItem{}
	err := db.Find(&orderItems, "order_id = ?", orderID).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return orderItems, nil
}