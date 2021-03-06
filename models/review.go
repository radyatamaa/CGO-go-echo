package models

import "time"

type Review struct {
	Id           string     `json:"id" `
	CreatedBy    string     `json:"created_by"`
	CreatedDate  time.Time  `json:"created_date"`
	ModifiedBy   *string    `json:"modified_by"`
	ModifiedDate *time.Time `json:"modified_date"`
	DeletedBy    *string    `json:"deleted_by"`
	DeletedDate  *time.Time `json:"deleted_date"`
	IsDeleted    int        `json:"is_deleted" validate:"required"`
	IsActive     int        `json:"is_active" validate:"required"`
	Values       float64        `json:"values"`
	Desc         string     `json:"desc"`
	ExpId        string     `json:"exp_id"`
	UserId      *string 		`json:"user_id"`
	GuideReview *float64		`json:"guide_review"`
	ActivitiesReview *float64	`json:"activities_review"`
	ServiceReview *float64		`json:"service_review"`
	CleanlinessReview *float64	`json:"cleanliness_review"`
	ValueReview *float64		`json:"value_review"`
}
type ReviewDto struct {
	Name   string    `json:"name"`
	Image  string    `json:"image"`
	Desc   string    `json:"desc"`
	Values float64       `json:"values"`
	Date   time.Time `json:"date"`
	UserId      *string 		`json:"user_id"`
	GuideReview *float64		`json:"guide_review"`
	ActivitiesReview *float64	`json:"activities_review"`
	ServiceReview *float64		`json:"service_review"`
	CleanlinessReview *float64	`json:"cleanliness_review"`
	ValueReview *float64		`json:"value_review"`
}
type NewReviewCommand struct {
	Id 		string	`json:"id"`
	ExpId 	string	`json:"exp_id"`
	Desc   string    `json:"desc"`
	GuideReview float64		`json:"guide_review"`
	ActivitiesReview float64	`json:"activities_review"`
	ServiceReview float64		`json:"service_review"`
	CleanlinessReview float64	`json:"cleanliness_review"`
	ValueReview float64		`json:"value_review"`
}
type ReviewDtoObject struct {
	Name   string `json:"name"`
	UserId string `json:"userid"`
	Desc   string `json:"desc"`
}
type ReviewsWithPagination struct {
	Data []*ReviewDto    `json:"data"`
	Meta *MetaPagination `json:"meta"`
}
