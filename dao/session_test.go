package dao

import (
	"bookstore/model"
	"fmt"
	"testing"
)

func TestSession(t *testing.T) {
	t.Run("", TestAddSession)
	t.Run("", TestDeleteSession)
	t.Run("", TestGetSession)
}

func TestAddSession(t *testing.T) {
	sess := &model.Session{
		SessionID: "13124235254524",
		Username: "root",
		UserID: 3,
	}
	err := AddSession(sess)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteSession(t *testing.T) {
	err := DeleteSession("13124235254524")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetSession(t *testing.T) {
	sess, err := GetSession("6f5153b7-c018-4aa5-633c-0ef6c52a8d5e")
	fmt.Printf("%+v\n", sess)
	if err != nil {
		t.Fatal(err)
	}
}