package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Entry struct {
	Id primitive.ObjectID 	`bson:"_id,omitempty" json:"id,omitempty"`
	Year int 				`json:"year,omitempty" binding:"required"`
	Month int 				`json:"month,omitempty" binding:"required"`
	Date        *time.Time 	`json:"date,omitempty"`
	Type        int       	`json:"type,omitempty"`
	Amount      float32     `json:"amount,omitempty"`
	Category	string		`json:"category,omitempty"`
	Concept     string    	`json:"concept,omitempty"`
	Description string    	`json:"description,omitempty"`
}
