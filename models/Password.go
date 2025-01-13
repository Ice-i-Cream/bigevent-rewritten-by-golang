package models

type Password struct {
	OldPwd string `json:"old_pwd"`
	NewPwd string `json:"new_pwd"`
	RePwd  string `json:"re_pwd"`
}
