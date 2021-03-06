package models

import "time"

type ExclusionService struct {
	Id                   int        `json:"id" validate:"required"`
	CreatedBy            string     `json:"created_by":"required"`
	CreatedDate          time.Time  `json:"created_date" validate:"required"`
	ModifiedBy           *string    `json:"modified_by"`
	ModifiedDate         *time.Time `json:"modified_date"`
	DeletedBy            *string    `json:"deleted_by"`
	DeletedDate          *time.Time `json:"deleted_date"`
	IsDeleted            int        `json:"is_deleted" validate:"required"`
	IsActive             int        `json:"is_active" validate:"required"`
	ExclusionServiceName string     `json:"exclusion_service_name"`
	ExclusionServiceType int        `json:"exclusion_service_type"`
}
type ExclusionServiceDto struct {
	Id                   int        `json:"id" validate:"required"`
	ExclusionServiceName string     `json:"exclusion_service_name"`
	ExclusionServiceType int        `json:"exclusion_service_type"`
}

type ExclusionServiceWithPagination struct {
	Data []*ExclusionServiceDto `json:"data"`
	Meta *MetaPagination    `json:"meta"`
}