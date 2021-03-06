package models

import "time"

type Experience struct {
	Id                      string     `json:"id" validate:"required"`
	CreatedBy               string     `json:"created_by":"required"`
	CreatedDate             time.Time  `json:"created_date" validate:"required"`
	ModifiedBy              *string    `json:"modified_by"`
	ModifiedDate            *time.Time `json:"modified_date"`
	DeletedBy               *string    `json:"deleted_by"`
	DeletedDate             *time.Time `json:"deleted_date"`
	IsDeleted               int        `json:"is_deleted" validate:"required"`
	IsActive                int        `json:"is_active" validate:"required"`
	ExpTitle                string     `json:"exp_title"`
	ExpType                 string     `json:"exp_type"`
	ExpTripType             string     `json:"exp_trip_type"`
	ExpBookingType          string     `json:"exp_booking_type"`
	ExpDesc                 string     `json:"exp_desc"`
	ExpMaxGuest             int        `json:"exp_max_guest"`
	ExpPickupPlace          string     `json:"exp_pickup_place"`
	ExpPickupTime           string     `json:"exp_pickup_time"`
	ExpPickupPlaceLongitude float64    `json:"exp_pickup_place_longitude"`
	ExpPickupPlaceLatitude  float64    `json:"exp_pickup_place_latitude"`
	ExpPickupPlaceMapsName  string     `json:"exp_pickup_place_maps_name"`
	ExpInternary            string     `json:"exp_internary"`
	ExpFacilities           string     `json:"exp_facilities"`
	ExpInclusion            string     `json:"exp_inclusion"`
	ExpRules                string     `json:"exp_rules"`
	Status                  int        `json:"status"`
	Rating                  float64    `json:"rating"`
	ExpLocationLatitude     float64    `json:"exp_location_latitude"`
	ExpLocationLongitude    float64    `json:"exp_location_longitude"`
	ExpLocationName         string     `json:"exp_location_name"`
	ExpCoverPhoto           *string    `json:"exp_cover_photo"`
	ExpDuration             int        `json:"exp_duration"`
	MinimumBookingId        *string     `json:"minimum_booking_id"`
	MerchantId              string     `json:"merchant_id"`
	HarborsId               *string     `json:"harbors_id"`
	GuideReview *float64		`json:"guide_review"`
	ActivitiesReview *float64	`json:"activities_review"`
	ServiceReview *float64		`json:"service_review"`
	CleanlinessReview *float64	`json:"cleanliness_review"`
	ValueReview *float64		`json:"value_review"`
	ExpPaymentDeadlineAmount *int `json:"exp_payment_deadline_amount"`
	ExpPaymentDeadlineType *string	`json:"exp_payment_deadline_type"`
	IsCustomisedByUser *int 		`json:"is_customised_by_user"`
	ExpLocationMapName 	*string		`json:"exp_location_map_name"`
	ExpLatitudeMap 		*float64	`json:"exp_latitude_map"`
	ExpLongitudeMap		*float64	`json:"exp_longitude_map"`
	ExpMaximumBookingAmount *int `json:"exp_maximum_booking_amount"`
	ExpMaximumBookingType *string	`json:"exp_maximum_booking_type"`
}
type ExperienceWithExperiencePayment struct {
	Id                      string     `json:"id" validate:"required"`
	CreatedBy               string     `json:"created_by":"required"`
	CreatedDate             time.Time  `json:"created_date" validate:"required"`
	ModifiedBy              *string    `json:"modified_by"`
	ModifiedDate            *time.Time `json:"modified_date"`
	DeletedBy               *string    `json:"deleted_by"`
	DeletedDate             *time.Time `json:"deleted_date"`
	IsDeleted               int        `json:"is_deleted" validate:"required"`
	IsActive                int        `json:"is_active" validate:"required"`
	ExpTitle                string     `json:"exp_title"`
	ExpType                 string     `json:"exp_type"`
	ExpTripType             string     `json:"exp_trip_type"`
	ExpBookingType          string     `json:"exp_booking_type"`
	ExpDesc                 string     `json:"exp_desc"`
	ExpMaxGuest             int        `json:"exp_max_guest"`
	ExpPickupPlace          string     `json:"exp_pickup_place"`
	ExpPickupTime           string     `json:"exp_pickup_time"`
	ExpPickupPlaceLongitude float64    `json:"exp_pickup_place_longitude"`
	ExpPickupPlaceLatitude  float64    `json:"exp_pickup_place_latitude"`
	ExpPickupPlaceMapsName  string     `json:"exp_pickup_place_maps_name"`
	ExpInternary            string     `json:"exp_internary"`
	ExpFacilities           string     `json:"exp_facilities"`
	ExpInclusion            string     `json:"exp_inclusion"`
	ExpRules                string     `json:"exp_rules"`
	Status                  int        `json:"status"`
	Rating                  float64    `json:"rating"`
	ExpLocationLatitude     float64    `json:"exp_location_latitude"`
	ExpLocationLongitude    float64    `json:"exp_location_longitude"`
	ExpLocationName         string     `json:"exp_location_name"`
	ExpCoverPhoto           *string    `json:"exp_cover_photo"`
	ExpDuration             int        `json:"exp_duration"`
	MinimumBookingId        *string     `json:"minimum_booking_id"`
	MerchantId              string     `json:"merchant_id"`
	HarborsId               *string     `json:"harbors_id"`
	GuideReview *float64		`json:"guide_review"`
	ActivitiesReview *float64	`json:"activities_review"`
	ServiceReview *float64		`json:"service_review"`
	CleanlinessReview *float64	`json:"cleanliness_review"`
	ValueReview *float64		`json:"value_review"`
	ExpPaymentDeadlineAmount *int `json:"exp_payment_deadline_amount"`
	ExpPaymentDeadlineType *string	`json:"exp_payment_deadline_type"`
	IsCustomisedByUser *int 		`json:"is_customised_by_user"`
	ExpLocationMapName 	*string		`json:"exp_location_map_name"`
	ExpLatitudeMap 		*float64	`json:"exp_latitude_map"`
	ExpLongitudeMap		*float64	`json:"exp_longitude_map"`
	ExpMaximumBookingAmount *int `json:"exp_maximum_booking_amount"`
	ExpMaximumBookingType *string	`json:"exp_maximum_booking_type"`
	ExpPaymentTypeName	string 	`json:"exp_payment_type_name"`
	OrderId 			string `json:"order_id"`
	BookingExpId 		string `json:"booking_exp_id"`
}
type ExperienceJoinForegnKey struct {
	Id                      string     `json:"id" validate:"required"`
	CreatedBy               string     `json:"created_by":"required"`
	CreatedDate             time.Time  `json:"created_date" validate:"required"`
	ModifiedBy              *string    `json:"modified_by"`
	ModifiedDate            *time.Time `json:"modified_date"`
	DeletedBy               *string    `json:"deleted_by"`
	DeletedDate             *time.Time `json:"deleted_date"`
	IsDeleted               int        `json:"is_deleted" validate:"required"`
	IsActive                int        `json:"is_active" validate:"required"`
	ExpTitle                string     `json:"exp_title"`
	ExpType                 string     `json:"exp_type"`
	ExpTripType             string     `json:"exp_trip_type"`
	ExpBookingType          string     `json:"exp_booking_type"`
	ExpDesc                 string     `json:"exp_desc"`
	ExpMaxGuest             int        `json:"exp_max_guest"`
	ExpPickupPlace          string     `json:"exp_pickup_place"`
	ExpPickupTime           string     `json:"exp_pickup_time"`
	ExpPickupPlaceLongitude float64    `json:"exp_pickup_place_longitude"`
	ExpPickupPlaceLatitude  float64    `json:"exp_pickup_place_latitude"`
	ExpPickupPlaceMapsName  string     `json:"exp_pickup_place_maps_name"`
	ExpInternary            string     `json:"exp_internary"`
	ExpFacilities           string     `json:"exp_facilities"`
	ExpInclusion            string     `json:"exp_inclusion"`
	ExpRules                string     `json:"exp_rules"`
	Status                  int        `json:"status"`
	Rating                  float64    `json:"rating"`
	ExpLocationLatitude     float64    `json:"exp_location_latitude"`
	ExpLocationLongitude    float64    `json:"exp_location_longitude"`
	ExpLocationName         string     `json:"exp_location_name"`
	ExpCoverPhoto           *string    `json:"exp_cover_photo"`
	ExpDuration             int        `json:"exp_duration"`
	MinimumBookingId        string     `json:"minimum_booking_id"`
	MerchantId              string     `json:"merchant_id"`
	HarborsId               string     `json:"harbors_id"`
	GuideReview *float64		`json:"guide_review"`
	ActivitiesReview *float64	`json:"activities_review"`
	ServiceReview *float64		`json:"service_review"`
	CleanlinessReview *float64	`json:"cleanliness_review"`
	ValueReview *float64		`json:"value_review"`
	ExpPaymentDeadlineAmount *int `json:"exp_payment_deadline_amount"`
	ExpPaymentDeadlineType *string	`json:"exp_payment_deadline_type"`
	IsCustomisedByUser *int 		`json:"is_customised_by_user"`
	ExpLocationMapName 	*string		`json:"exp_location_map_name"`
	ExpLatitudeMap 		*float64	`json:"exp_latitude_map"`
	ExpLongitudeMap		*float64	`json:"exp_longitude_map"`
	ExpMaximumBookingAmount *int `json:"exp_maximum_booking_amount"`
	ExpMaximumBookingType *string	`json:"exp_maximum_booking_type"`
	MinimumBookingAmount    *int       `json:"minimum_booking_amount"`
	MinimumBookingDesc      string     `json:"minimum_booking_desc"`
}
type ExperienceDto struct {
	Id                      string                `json:"id" validate:"required"`
	ExpTitle                string                `json:"exp_title"`
	ExpType                 []string              `json:"exp_type"`
	ExpTripType             string                `json:"exp_trip_type"`
	ExpBookingType          string                `json:"exp_booking_type"`
	ExpDesc                 string                `json:"exp_desc"`
	ExpMaxGuest             int                   `json:"exp_max_guest"`
	ExpPickupPlace          string                `json:"exp_pickup_place"`
	ExpPickupTime           string                `json:"exp_pickup_time"`
	ExpPickupPlaceLongitude float64               `json:"exp_pickup_place_longitude"`
	ExpPickupPlaceLatitude  float64               `json:"exp_pickup_place_latitude"`
	ExpPickupPlaceMapsName  string                `json:"exp_pickup_place_maps_name"`
	ExpInternary            ExpItineraryObject    `json:"exp_internary"`
	ExpFacilities           []ExpFacilitiesObject `json:"exp_facilities"`
	ExpInclusion            []ExpInclusionObject  `json:"exp_inclusion"`
	ExpRules                []ExpRulesObject      `json:"exp_rules"`
	ExpAvailability         []ExpAvailablitityObj `json:"exp_availability"`
	ExpPayment              []ExpPaymentObj       `json:"exp_payment"`
	ExpPhotos               []ExpPhotosObj        `json:"exp_photos"`
	ExperienceAddOn         []ExperienceAddOnObj  `json:"experience_add_on"`
	Status                  int                   `json:"status"`
	Rating                  float64               `json:"rating"`
	CountRating             int                   `json:"count_rating"`
	ExpLocationLatitude     float64               `json:"exp_location_latitude"`
	ExpLocationLongitude    float64               `json:"exp_location_longitude"`
	ExpLocationName         string                `json:"exp_location_name"`
	ExpCoverPhoto           *string               `json:"exp_cover_photo"`
	ExpDuration             int                   `json:"exp_duration"`
	MinimumBooking          MinimumBookingObj     `json:"minimum_booking"`
	MerchantId              string                `json:"merchant_id"`
	HarborsName             string                `json:"harbors_name"`
	City                    string                `json:"city"`
	Province                string                `json:"province"`
	GuideReview *float64		`json:"guide_review"`
	ActivitiesReview *float64	`json:"activities_review"`
	ServiceReview *float64		`json:"service_review"`
	CleanlinessReview *float64	`json:"cleanliness_review"`
	ValueReview *float64		`json:"value_review"`
	ExpPaymentDeadlineAmount *int `json:"exp_payment_deadline_amount"`
	ExpPaymentDeadlineType *string	`json:"exp_payment_deadline_type"`
	IsCustomisedByUser *int 		`json:"is_customised_by_user"`
	ExpLocationMapName 	*string		`json:"exp_location_map_name"`
	ExpLatitudeMap 		*float64	`json:"exp_latitude_map"`
	ExpLongitudeMap		*float64	`json:"exp_longitude_map"`
	ExpMaximumBookingAmount *int `json:"exp_maximum_booking_amount"`
	ExpMaximumBookingType *string	`json:"exp_maximum_booking_type"`
}
type MasterDataExperience struct {
	Id		string `json:"id"`
	Title	string	`json:"title"`
}

type ResponseCreateExperience struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}
type NewCommandChangeStatus struct {
	ExpId		string		`json:"exp_id"`
	TransId 	string		`json:"trans_id"`
	Status 		int		`json:"status"`
}
type NewCommandExperience struct {
	Id                      string                `json:"id"`
	ExpTitle                string                `json:"exp_title"`
	ExpType                 []string              `json:"exp_type"`
	ExpTripType             string                `json:"exp_trip_type"`
	ExpBookingType          string                `json:"exp_booking_type"`
	ExpDesc                 string                `json:"exp_desc"`
	ExpMaxGuest             int                   `json:"exp_max_guest"`
	ExpPickupPlace          string                `json:"exp_pickup_place"`
	ExpPickupTime           string                `json:"exp_pickup_time"`
	ExpPickupPlaceLongitude float64               `json:"exp_pickup_place_longitude"`
	ExpPickupPlaceLatitude  float64               `json:"exp_pickup_place_latitude"`
	ExpPickupPlaceMapsName  string                `json:"exp_pickup_place_maps_name"`
	ExpInternary            ExpItineraryObject    `json:"exp_internary"`
	ExpFacilities           []ExpFacilitiesObject `json:"exp_facilities"`
	ExpInclusion            []ExpInclusionObject  `json:"exp_inclusion"`
	ExpRules                []ExpRulesObject      `json:"exp_rules"`
	ExpAvailability         []ExpAvailablitityObj `json:"exp_availability"`
	ExpPayment              []ExpPaymentObj       `json:"exp_payment"`
	ExpPhotos               []ExpPhotosObj        `json:"exp_photos"`
	ExperienceAddOn         []ExperienceAddOnObj  `json:"experience_add_on"`
	Status                  int                   `json:"status"`
	ExpLocationLatitude  float64 `json:"exp_location_latitude"`
	ExpLocationLongitude float64 `json:"exp_location_longitude"`
	ExpLocationName      string  `json:"exp_location_name"`
	ExpCoverPhoto        *string `json:"exp_cover_photo"`
	ExpDuration          int     `json:"exp_duration"`
	MinimumBookingId     string  `json:"minimum_booking_id"`
	HarborsId string `json:"harbors_id"`
	ExpPaymentDeadlineAmount int `json:"exp_payment_deadline_amount"`
	ExpPaymentDeadlineType string	`json:"exp_payment_deadline_type"`
	IsCustomisedByUser int 		`json:"is_customised_by_user"`
	ExpLocationMapName 	*string		`json:"exp_location_map_name"`
	ExpLatitudeMap 		*float64	`json:"exp_latitude_map"`
	ExpLongitudeMap		*float64	`json:"exp_longitude_map"`
	ExpMaximumBookingAmount *int `json:"exp_maximum_booking_amount"`
	ExpMaximumBookingType *string	`json:"exp_maximum_booking_type"`
}
type CustomPrice struct {
	Currency	string 		`json:"currency"`
	Price 		float64		`json:"price"`
	Guest 		int 		`json:"guest"`
}
type MinimumBookingObj struct {
	MinimumBookingDesc   string `json:"minimum_booking_desc"`
	MinimumBookingAmount *int   `json:"minimum_booking_amount"`
}
type ExperienceAddOnObj struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Desc     string  `json:"desc"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}
type ExpAvailablitityObj struct {
	Id    string   `json:"id"`
	Year  int      `json:"year"`
	Month string   `json:"month"`
	Date  []string `json:"date"`
}
type ExpPaymentObj struct {
	Id              string  `json:"id"`
	Currency        string  `json:"currency"`
	Price           float64 `json:"price"`
	PriceItemType   string  `json:"price_item_type"`
	PaymentTypeId   string  `json:"payment_type_id"`
	PaymentTypeName string  `json:"payment_type_name"`
	PaymentTypeDesc string  `json:"payment_type_desc"`
	CustomPrice    []CustomPrice	`json:"custom_price"`
}
type ExpPhotosObj struct {
	Id            string           `json:"id"`
	Folder        string           `json:"folder"`
	ExpPhotoImage []CoverPhotosObj `json:"exp_photo_image"`
}
type ExpItineraryObject struct {
	Item []ItemObject `json:"item"`
}
type ItemObject struct {
	Day       int              `json:"day"`
	Itinerary []ItineraryObjet `json:"itinerary"`
}
type ItineraryObjet struct {
	Time     string `json:"time"`
	Activity string `json:"activity"`
}
type ExpFacilitiesObject struct {
	Name   string `json:"name"`
	Icon   string `json:"icon"`
	Amount int    `json:"amount"`
}

type ExpInclusionObject struct {
	Name string `json:"name"`
	Type int    `json:"type"`
}
type ExpRulesObject struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type ExpSearch struct {
	Id         string  `json:"id" validate:"required"`
	ExpTitle   string  `json:"exp_title"`
	ExpType    string  `json:"exp_type"`
	ExpStatus  int     `json:"exp_status"`
	Rating     float64 `json:"rating"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	CoverPhoto string  `json:"cover_photo"`
	Province   string  `json:"province"`
	ExpLocationMapName 	*string		`json:"exp_location_map_name"`
	ExpLatitudeMap 		*float64	`json:"exp_latitude_map"`
	ExpLongitudeMap		*float64	`json:"exp_longitude_map"`
	//Price	   string   `json:"price"`
}

type ExpSearchObject struct {
	Id          string         `json:"id" validate:"required"`
	ExpTitle    string         `json:"exp_title"`
	ExpType     []string       `json:"exp_type"`
	ExpStatus   string         `json:"exp_status"`
	Rating      float64        `json:"rating"`
	CountRating int            `json:"count_rating"`
	Currency    string         `json:"currency"`
	Price       float64        `json:"price"`
	PaymentType string         `json:"payment_type"`
	Latitude    float64        `json:"latitude"`
	Longitude   float64        `json:"longitude"`
	Province    string         `json:"province"`
	CoverPhoto  CoverPhotosObj `json:"cover_photo"`
	ListPhoto   []ExpPhotosObj `json:"list_photo"`
	ExpLocationMapName 	*string		`json:"exp_location_map_name"`
	ExpLatitudeMap 		*float64	`json:"exp_latitude_map"`
	ExpLongitudeMap		*float64	`json:"exp_longitude_map"`

}
type ExperienceUserDiscoverPreferenceDto struct {
	Id           string         `json:"id" validate:"required"`
	ExpTitle     string         `json:"exp_title"`
	ExpType      []string       `json:"exp_type"`
	Rating       float64        `json:"rating"`
	CountRating  int            `json:"count_rating"`
	Currency     string         `json:"currency"`
	Price        float64        `json:"price"`
	Payment_type string         `json:"payment_type"`
	Cover_Photo  CoverPhotosObj `json:"cover_photo"`
}
type CoverPhotosObj struct {
	Original  string `json:"original"`
	Thumbnail string `json:"thumbnail"`
}
type ExpUserDiscoverPreference struct {
	ProvinceId				int 	`json:"province_id"`
	ProvinceName			string	`json:"province_name"`
	CityId                  int        `json:"city_id"`
	CityName                string     `json:"city_name"`
	CityDesc                string     `json:"city_desc"`
	CityPhotos              *string    `json:"city_photos"`
	IdHarbors				string		`json:"id_harbors"`
	HarborsName				string		`json:"harbors_name"`
	Id                      string     `json:"id" validate:"required"`
	CreatedBy               string     `json:"created_by":"required"`
	CreatedDate             time.Time  `json:"created_date" validate:"required"`
	ModifiedBy              *string    `json:"modified_by"`
	ModifiedDate            *time.Time `json:"modified_date"`
	DeletedBy               *string    `json:"deleted_by"`
	DeletedDate             *time.Time `json:"deleted_date"`
	IsDeleted               int        `json:"is_deleted" validate:"required"`
	IsActive                int        `json:"is_active" validate:"required"`
	ExpTitle                string     `json:"exp_title"`
	ExpType                 string     `json:"exp_type"`
	ExpTripType             string     `json:"exp_trip_type"`
	ExpBookingType          string     `json:"exp_booking_type"`
	ExpDesc                 string     `json:"exp_desc"`
	ExpMaxGuest             int        `json:"exp_max_guest"`
	ExpPickupPlace          string     `json:"exp_pickup_place"`
	ExpPickupTime           string     `json:"exp_pickup_time"`
	ExpPickupPlaceLongitude float64    `json:"exp_pickup_place_longitude"`
	ExpPickupPlaceLatitude  float64    `json:"exp_pickup_place_latitude"`
	ExpPickupPlaceMapsName  string     `json:"exp_pickup_place_maps_name"`
	ExpInternary            string     `json:"exp_internary"`
	ExpFacilities           string     `json:"exp_facilities"`
	ExpInclusion            string     `json:"exp_inclusion"`
	ExpRules                string     `json:"exp_rules"`
	Status                  int        `json:"status"`
	Rating                  float64    `json:"rating"`
	ExpLocationLatitude     float64    `json:"exp_location_latitude"`
	ExpLocationLongitude    float64    `json:"exp_location_longitude"`
	ExpLocationName         string     `json:"exp_location_name"`
	ExpCoverPhoto           *string    `json:"exp_cover_photo"`
	ExpDuration             int        `json:"exp_duration"`
	MinimumBookingId        string     `json:"minimum_booking_id"`
	MerchantId              string     `json:"merchant_id"`
	HarborsId               string     `json:"harbors_id"`
	GuideReview *float64		`json:"guide_review"`
	ActivitiesReview *float64	`json:"activities_review"`
	ServiceReview *float64		`json:"service_review"`
	CleanlinessReview *float64	`json:"cleanliness_review"`
	ValueReview *float64		`json:"value_review"`
	ExpPaymentDeadlineAmount *int `json:"exp_payment_deadline_amount"`
	ExpPaymentDeadlineType *string	`json:"exp_payment_deadline_type"`
	IsCustomisedByUser *int 		`json:"is_customised_by_user"`
	ExpLocationMapName 	*string		`json:"exp_location_map_name"`
	ExpLatitudeMap 		*float64	`json:"exp_latitude_map"`
	ExpLongitudeMap		*float64	`json:"exp_longitude_map"`
	ExpMaximumBookingAmount *int `json:"exp_maximum_booking_amount"`
	ExpMaximumBookingType *string	`json:"exp_maximum_booking_type"`
}

type ExpUserDiscoverPreferenceDto struct {
	ProvinceId int 									`json:"province_id"`
	ProvinceName string	`json:"province_name"`
	CityId     int                                   `json:"city_id"`
	City       string                                `json:"city"`
	CityDesc   string                                `json:"city_desc"`
	CityPhotos []CoverPhotosObj                      `json:"city_photos"`
	HarborsId	string							`json:"harbors_id"`
	HarborsName	string	`json:"harbors_name"`
	Item       []ExperienceUserDiscoverPreferenceDto `json:"item"`
}

type FilterSearchWithPagination struct {
	Data []*ExpSearchObject `json:"data"`
	Meta *MetaPagination    `json:"meta"`
}
