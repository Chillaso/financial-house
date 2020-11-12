package model

import "time"

type Book struct {
	Year []struct {
		Month struct { Entry []struct {
				Date        time.Time `json:"date"`
				Type        int       `json:"type"`
				Amount      int       `json:"amount"`
				Concept     string    `json:"concept"`
				Description string    `json:"description"`
			} `json:"entry"`
		} `json:"month"`
	} `json:"year"`
}