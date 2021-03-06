package models

import "time"

type ExperiencePaymentType struct {
	Id                 string     `json:"id" validate:"required"`
	CreatedBy          string     `json:"created_by":"required"`
	CreatedDate        time.Time  `json:"created_date" validate:"required"`
	ModifiedBy         *string    `json:"modified_by"`
	ModifiedDate       *time.Time `json:"modified_date"`
	DeletedBy          *string    `json:"deleted_by"`
	DeletedDate        *time.Time `json:"deleted_date"`
	IsDeleted          int        `json:"is_deleted" validate:"required"`
	IsActive           int        `json:"is_active" validate:"required"`
	ExpPaymentTypeName string     `json:"exp_payment_type_id"`
	ExpPaymentTypeDesc string     `json:"exp_payment_type_desc"`
}

type ExperiencePaymentTypeDto struct {
	Id   string `json:"id" validate:"required"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	RemainingPayment float64	`json:"remaining_payment"`
	OriginalPrice 	*float64	`json:"original_price"`
}
