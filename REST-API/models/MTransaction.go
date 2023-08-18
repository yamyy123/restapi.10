package models

import "time"

type Transaction struct {
	ID               int       `json:"id" bson:"_id"`
	Transaction_Date time.Time `json:"date" bson:"date"`
	Amount           int       `json:"amount" bson:"amount"`
	Transaction_Type string    `json:"transaction_type" bson:"transaction_type"`
	Customer_id      int       `json:"customer_id" bson:"customer_id"`
}

type DBResponse struct {
	
	Name               string             `json:"name" bson:"name"`
	Email              string             `json:"email" bson:"email"`
	Password           string             `json:"password" bson:"password"`
	PasswordConfirm    string             `json:"passwordConfirm,omitempty" bson:"passwordConfirm,omitempty"`
	Role               string             `json:"role" bson:"role"`
	VerificationCode   string             `json:"verificationCode,omitempty" bson:"verificationCode"`
	ResetPasswordToken string             `json:"resetPasswordToken,omitempty" bson:"resetPasswordToken,omitempty"`
	ResetPasswordAt    time.Time          `json:"resetPasswordAt,omitempty" bson:"resetPasswordAt,omitempty"`
	Verified           bool               `json:"verified" bson:"verified"`
	CreatedAt          time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at" bson:"updated_at"`
}