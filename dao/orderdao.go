package dao

import (
	"bookstore/model"
	"log"
)

func AddOrder(order *model.Order) error {
	err := db.Create(order).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetOrders() ([]*model.Order, error) {
	orders := []*model.Order{}
	err := db.Find(&orders).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return orders, nil
}

func GetMyOrders(userID int) ([]*model.Order, error) {
	orders := []*model.Order{}
	err := db.Find(&orders, "user_id = ?", userID).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return orders, nil
}

func UpdateOrderState(orderID string, state int) error {
	err := db.Model(&model.Order{}).Where("id = ?", orderID).Update("state", state).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}