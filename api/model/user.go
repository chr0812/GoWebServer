package model

type User struct {
	EmailId     string `json:"emailId"`
	Password    string `json:"password"`
	Nickname    string `json:"nickname"`
	Birth       int    `json:"birth"`
	PhoneNumber string `json:"phonenumber"`
}
