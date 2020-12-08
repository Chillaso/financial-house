package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Entry struct {
	Date        *time.Time 	`json:"date,omitempty"`
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
	Id primitive.ObjectID 	`bson:"_id,omitempty" json:"-"`
	Year int 				`json:"year,omitempty" binding:"required"`
	Months []Month 			`json:"months,omitempty"`
}