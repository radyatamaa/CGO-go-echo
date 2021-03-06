package models

import "time"

type City struct {
	Id           int        `json:"id" validate:"required"`
	CreatedBy    string     `json:"created_by":"required"`
	CreatedDate  time.Time  `json:"created_date" validate:"required"`
	ModifiedBy   *string    `json:"modified_by"`
	ModifiedDate *time.Time `json:"modified_date"`
	DeletedBy    *string    `json:"deleted_by"`
	DeletedDate  *time.Time `json:"deleted_date"`
	IsDeleted    int        `json:"is_deleted" validate:"required"`
	IsActive     int        `json:"is_active" validate:"required"`
	CityName     string     `json:"city_name"`
	CityDesc     string     `json:"city_desc"`
	CityPhotos   *string    `json:"city_photos"`
	ProvinceId   int        `json:"province_id"`
}

type NewCommandCity struct {
	Id           int        `json:"id"`
	CityName     string     `json:"city_name"`
	CityDesc     string     `json:"city_desc"`
	CityPhotos   []CoverPhotosObj    `json:"city_photos"`
	ProvinceId   int        `json:"province_id"`
}

type CityDto struct {
	Id           int        `json:"id"`
	CityName     string     `json:"city_name"`
	CityDesc     string     `json:"city_desc"`
	CityPhotos   []CoverPhotosObj    `json:"city_photos"`
	ProvinceId   int        `json:"province_id"`
}


type CityDtoWithPagination struct {
	Data []*CityDto `json:"data"`
	Meta *MetaPagination    `json:"meta"`
}