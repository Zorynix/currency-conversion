package models

type Currency struct {
	Id            int
	Code          string `json:"code"`
	Name          string `json:"name"`
	Symbol        string `json:"symbol"`
	DecimalNumber int    `json:"decimal_digits"`
	Active        bool
	MainAreaId    int
}
