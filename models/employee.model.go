package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	BasicInfo struct {
		FullName              string    `bson:"full_name"`
		DOB                   time.Time `bson:"dob"`
		Gender                int       `bson:"gender"`
		PlaceOfBirthWard      string    `bson:"place_of_birth_ward"`
		PlaceOfBirthDistrict  string    `bson:"place_of_birth_district"`
		PlaceOfBirthProvince  string    `bson:"place_of_birth_province"`
		Ethnic                string    `bson:"ethnic"`
		Religion              string    `bson:"religion"`
		Address               string    `bson:"address"`
		IdentificationNumber  string    `bson:"identification_number"`
		IssuanceDate          time.Time `bson:"issuance_date"`
		SocialInsuranceNumber string    `bson:"social_insurance_number"`
	} `bson:"basic_infor"`
	Recruitment struct {
		Code            string    `bson:"code"`
		RecruitmentDate time.Time `bson:"recruitment_date"`
		Position        string    `bson:"position"`
		SubPosition     string    `bson:"sub_position"`
	} `bson:"recruitment"`

	Education struct {
		Level                     int       `bson:"level"`
		QualificationLevel        int       `bson:"qualification_level"`
		SpecialQualificationLevel int       `bson:"special_qualification_level"`
		ForeignLanguageLevel      int       `bson:"foreign_language_level"`
		ITLevel                   int       `bson:"information_technology_level"`
		JoinPartyDate             time.Time `bson:"join_party_date"`
		JoinDate                  time.Time `bson:"join_date"`
	} `bson:"education"`

	Document struct {
		IdentificationLink string `bson:"identification_link"`
		DegreeLink         string `bson:"degree_link"`
		ContractLink       string `bson:"contract_link"`
	} `bson:"document"`

	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	CreatedBy primitive.ObjectID `bson:"created_by"`
	UpdatedBy primitive.ObjectID `bson:"updated_by"`
}
