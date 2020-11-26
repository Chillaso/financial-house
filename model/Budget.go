package model

type Budget struct {
	Input   int `json:"input"`
	FixCost int `json:"fixCost"`
	Leisure int `json:"leisure"`
	Savings int `json:"savings"`
}