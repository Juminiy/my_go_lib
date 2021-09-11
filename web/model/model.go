package model

import "github.com/jinzhu/gorm"

type Arr struct {
	TArr []int `json:"arr" xml:"arr" form:"arr"`
}

type TraditionalAuth struct {
	gorm.Model
	AuthUserName string `json:"AuthUserName" form:"AuthUserName"`
	AuthPassWord string `json:"AuthPassWord" form:"AuthUserName"`
}

func SignIn() {

}
func SignUp() {

}
func DeleteUser() {

}
