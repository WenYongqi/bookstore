package dao

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	t.Run("", testLogin)
	t.Run("" ,testRegister)
	t.Run("", testSave)
}

func testLogin(t *testing.T) {
	user, err := CheckUsernameAndPassword("admin", "123456")
	fmt.Println("测试登录：", user, err)
}

func testRegister(t *testing.T) {
	user, err := CheckUsername("admin")
	fmt.Println("测试注册", user, err)
}

func testSave(t *testing.T) {
	err := SaveUser("admin", "123456", "admin@gmail.com")
	fmt.Println("测试保存", err)
}