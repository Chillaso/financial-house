package model

import "time"

type Entry struct {
	Date        time.Time `json:"date"`
	Type        int       `json:"type"`
	Amount      int       `json:"amount"`
	Concept     string    `json:"concept"`
	Description string    `json:"description"`
}

type Month struct {
	Month int `json:"month"`
	Entries []Entry `json:"entry"`
}

type Book struct {
	Id int
	Year int `json:"year"`
	Months []Month `json:"month"`
}