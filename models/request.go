package models

type Users struct {
	NAME     string `json:"name"`
	EMAIL    string `json:"email"`
	PASSWORD string `json:"password"`
}
