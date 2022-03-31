package model

import "time"

type Order struct {
	ID string
	CreatedAt time.Time
	TotalCount int64
	TotalAmount float64
	State int //订单的状态 0未发货 1已发货 2交易完成
	UserID int
	Username string
}

func (order *Order) CreatedTime() string {
	return order.CreatedAt.Format("2006-01-02 15:04:05")
}

func (order *Order) StateNotSend() bool {
	return order.State == 0
}

func (order *Order) StateNotConfirm() bool {
	return order.State != 2
}