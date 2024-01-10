package models

//структура структур

type Currency struct {
	Id            int
	Code          string `json:"code"`
	Name          string `json:"name"`
	Symbol        string `json:"symbol_native"`
	DecimalNumber int    `json:"decimal_digits"`
	Active        bool
	MainAreaId    int
}
