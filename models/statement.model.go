package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Statement struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Title      string             `bson:"title"`
	Code       string             `bson:"code"`
	Date       time.Time          `bson:"date"`
	Signer     string             `bson:"signer"`
	FileLink   string             `bson:"file_link"`
	ReportLink string             `bson:"report_link"`
	Status     bool               `bson:"status"`
	Fwd        string             `bson:"fwd"`
	Versions   []struct {
		ID       primitive.ObjectID `bson:"_id,omitempty"`
		Date     time.Time          `bson:"date"`
		Signer   string             `bson:"signer"`
		FWD      string             `bson:"fwd"`
		FileLink string             `bson:"file_link"`
	}
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	CreatedBy primitive.ObjectID `bson:"created_by"`
	UpdatedBy primitive.ObjectID `bson:"updated_by"`
}
