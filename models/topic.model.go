package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Member struct {
	Name      string `bson:"name"`
	Code      string `bson:"code"`
	WorkPlace string `bson:"work_place"`
	Role      int    `bson:"role"`
}

type Topic struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	Title           string    `bson:"title"`
	Type            string    `bson:"type"`
	Code            string    `bson:"code"`
	Host            string    `bson:"host"`
	ManagementLevel int       `bson:"management_level"`
	MemberNumber    int       `bson:"member_number"`
	StartDate       time.Time `bson:"start_date"`
	Money           float64   `bson:"money"`
	Members         []Member  `bson:"members"`
	File            struct {
		ExplainLink string `bson:"explain_link"`
		OutlineLink string `bson:"outline_link"`
		ReportLink  string `bson:"report_link"`
	} `bson:"file"`

	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	CreatedBy primitive.ObjectID `bson:"created_by"`
	UpdatedBy primitive.ObjectID `bson:"updated_by"`
}
