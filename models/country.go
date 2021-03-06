package models

import "time"

type Country struct {
	Id           int        `json:"id" validate:"required"`
	CreatedBy    string     `json:"created_by":"required"`
	CreatedDate  time.Time  `json:"created_date" validate:"required"`
	ModifiedBy   *string    `json:"modified_by"`
	ModifiedDate *time.Time `json:"modified_date"`
	DeletedBy    *string    `json:"deleted_by"`
	DeletedDate  *time.Time `json:"deleted_date"`
	IsDeleted    int        `json:"is_deleted" validate:"required"`
	IsActive     int        `json:"is_active" validate:"required"`
	CountryName  string     `json:"country_name"`
	Iso 		*string 		`json:"iso"`
	Name 		*string		`json:"name"`
	NiceName 	*string 		`json:"nicename"`
	Iso3 		*string 		`json:"iso3"`
	NumCode 	*int    	`json:"numcode"`
	PhoneCode 	*int 		`json:"phonecode"`
}
type CountryDto struct {
	Id           int        `json:"id"`
	CountryName  string     `json:"country_name"`
	PhoneCode	 *int		`json:"phone_code"`
}
type NewCommandCountry struct {
	Id           int        `json:"id"`
	CountryName  string     `json:"country_name"`
	Iso 		*string 		`json:"iso"`
	Name 		*string		`json:"name"`
	NiceName 	*string 		`json:"nicename"`
	Iso3 		*string 		`json:"iso3"`
	NumCode 	*int    	`json:"numcode"`
	PhoneCode 	*int 		`json:"phonecode"`
}

type CountryDtoWithPagination struct {
	Data []*CountryDto `json:"data"`
	Meta *MetaPagination    `json:"meta"`
}