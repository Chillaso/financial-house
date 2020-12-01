package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Entry struct {
	Date        time.Time 	`json:"date,omitempty"`
	Type        int       	`json:"type,omitempty"`
	Amount      int       	`json:"amount,omitempty"`
	Concept     string    	`json:"concept,omitempty"`
	Description string    	`json:"description,omitempty"`
}

type Month struct {
	Month int 				`json:"month,omitempty"`
	Entries []Entry 		`json:"entries,omitempty"`
}

type Book struct {
	ID primitive.ObjectID 	`bson:"_id" json:"id,omitempty"`
	CreatedAt time.Time 	`bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time 	`bson:"updated_at" json:"updated_at,omitempty"`
	Year int 				`json:"year,omitempty"`
	Months []Month 			`json:"months,omitempty"`
}