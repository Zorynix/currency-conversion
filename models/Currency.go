package models

type Currency struct {
	Id            int
	Code          string `json:"code"`
	Name          string `json:"name"`
	Symbol        string `json:"symbol"`
	DecimalNumber string `json:"decimal_digits"`
	Active        string `json:""`
	MainAreaId    string `json:""`
}
