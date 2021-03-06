package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	guuid "github.com/google/uuid"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/misc/currency"
	"github.com/service/exclude"
	"github.com/service/exp_Include"
	"github.com/service/exp_exclude"
	"github.com/service/exp_facilities"
	"github.com/service/facilities"
	"github.com/service/include"

	"github.com/booking/booking_exp"
	"github.com/service/temp_user_preferences"

	"github.com/service/filter_activity_type"

	"github.com/product/experience_add_ons"

	"github.com/auth/merchant"
	"github.com/product/reviews"
	"github.com/service/cpc"
	"github.com/service/exp_availability"
	"github.com/service/exp_photos"
	"github.com/service/harbors"

	"github.com/models"
	inspiration "github.com/service/exp_inspiration"
	payment "github.com/service/exp_payment"
	types "github.com/service/exp_types"
	"github.com/service/experience"
)

type experienceUsecase struct {
	bookingRepo       booking_exp.Repository
	tempUserPreRepo   temp_user_preferences.Repository
	filterATRepo      filter_activity_type.Repository
	adOnsRepo         experience_add_ons.Repository
	experienceRepo    experience.Repository
	harborsRepo       harbors.Repository
	cpcRepo           cpc.Repository
	paymentRepo       payment.Repository
	reviewsRepo       reviews.Repository
	typesRepo         types.Repository
	inspirationRepo   inspiration.Repository
	expPhotos         exp_photos.Repository
	mUsecase          merchant.Usecase
	contextTimeout    time.Duration
	exp_availablitiy  exp_availability.Repository
	expFacilitiesRepo exp_facilities.Repository
	expIncludeRepo    exp_Include.Repository
	expExcludeRepo    exp_exclude.Repository
	facilitiesRepo    facilities.Repository
	includeRepo       include.Repository
	excludeRepo       exclude.Repository
	currencyUsecase   currency.Usecase
}

// NewexperienceUsecase will create new an experienceUsecase object representation of experience.Usecase interface
func NewexperienceUsecase(
	currencyUsecase currency.Usecase,
	expFacilitiesRepo exp_facilities.Repository,
	expIncludeRepo exp_Include.Repository,
	expExcludeRepo exp_exclude.Repository,
	facilitiesRepo facilities.Repository,
	includeRepo include.Repository,
	excludeRepo exclude.Repository,
	b booking_exp.Repository,
	tup temp_user_preferences.Repository,
	fac filter_activity_type.Repository,
	adOns experience_add_ons.Repository,
	ea exp_availability.Repository,
	ps exp_photos.Repository,
	a experience.Repository,
	h harbors.Repository,
	c cpc.Repository,
	p payment.Repository,
	r reviews.Repository,
	t types.Repository,
	i inspiration.Repository,
	m merchant.Usecase,
	timeout time.Duration,
) experience.Usecase {
	return &experienceUsecase{
		currencyUsecase:   currencyUsecase,
		expFacilitiesRepo: expFacilitiesRepo,
		expIncludeRepo:    expIncludeRepo,
		expExcludeRepo:    expExcludeRepo,
		facilitiesRepo:    facilitiesRepo,
		includeRepo:       includeRepo,
		excludeRepo:       excludeRepo,
		bookingRepo:       b,
		tempUserPreRepo:   tup,
		filterATRepo:      fac,
		adOnsRepo:         adOns,
		exp_availablitiy:  ea,
		experienceRepo:    a,
		harborsRepo:       h,
		cpcRepo:           c,
		paymentRepo:       p,
		reviewsRepo:       r,
		typesRepo:         t,
		inspirationRepo:   i,
		mUsecase:          m,
		contextTimeout:    timeout,
		expPhotos:         ps,
	}
}

func (m experienceUsecase) UpdateStatus(ctx context.Context, status int, id string, token string) (*models.NewCommandChangeStatus, error) {
	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()

	currentMerchant, err := m.mUsecase.ValidateTokenMerchant(ctx, token)
	if err != nil {
		return nil, err
	}

	errorUpdate := m.experienceRepo.UpdateStatus(ctx, status, id, currentMerchant.MerchantEmail)
	if errorUpdate != nil {
		return nil, errorUpdate
	}
	result := models.NewCommandChangeStatus{
		ExpId:   id,
		TransId: "",
		Status:  status,
	}
	return &result, nil
}


func (m experienceUsecase) GetExpPendingTransactionCount(ctx context.Context, token string) (*models.Count, error) {
	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()

	currentMerchant, err := m.mUsecase.ValidateTokenMerchant(ctx, token)
	if err != nil {
		return nil, err
	}

	count, err := m.experienceRepo.GetExpPendingTransactionCount(ctx, currentMerchant.Id)
	if err != nil {
		return nil, err
	}

	return &models.Count{Count: count}, nil
}

func (m experienceUsecase) GetExperienceMasterData(ctx context.Context) ([]*models.MasterDataExperience, error) {
	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()

	count, err := m.experienceRepo.GetAllExperience(ctx)
	if err != nil {
		return nil, err
	}
	result := []*models.MasterDataExperience{}
	for _, element := range count {
		res := models.MasterDataExperience{
			Id: element.Id,
			Title: element.ExpTitle,
		}
		result = append(result, &res)
	}


	return result, nil
}

func (m experienceUsecase) GetExpFailedTransactionCount(ctx context.Context, token string) (*models.Count, error) {
	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()

	currentMerchant, err := m.mUsecase.ValidateTokenMerchant(ctx, token)
	if err != nil {
		return nil, err
	}

	count, err := m.experienceRepo.GetExpFailedTransactionCount(ctx, currentMerchant.Id)
	if err != nil {
		return nil, err
	}

	return &models.Count{Count: count}, nil
}

func (m experienceUsecase) GetPublishedExpCount(ctx context.Context, token string) (*models.Count, error) {
	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()

	currentMerchant, err := m.mUsecase.ValidateTokenMerchant(ctx, token)
	if err != nil {
		return nil, err
	}

	count, err := m.experienceRepo.GetPublishedExpCount(ctx, currentMerchant.Id)
	if err != nil {
		return nil, err
	}

	return &models.Count{Count: count}, nil
}

func (m experienceUsecase) GetSuccessBookCount(ctx context.Context, token string) (*models.Count, error) {
	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()

	currentMerchant, err := m.mUsecase.ValidateTokenMerchant(ctx, token)
	if err != nil {
		return nil, err
	}

	count, err := m.experienceRepo.GetSuccessBookCount(ctx, currentMerchant.Id)
	if err != nil {
		return nil, err
	}

	return &models.Count{Count: count}, nil
}

func (m experienceUsecase) GetByCategoryID(ctx context.Context, categoryId int) ([]*models.ExpSearchObject, error) {
	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()

	expList, err := m.experienceRepo.GetByCategoryID(ctx, categoryId)
	if err != nil {
		return nil, err
	}

	results := make([]*models.ExpSearchObject, len(expList))
	for i, exp := range expList {
		var expType []string
		if errUnmarshal := json.Unmarshal([]byte(exp.ExpType), &expType); errUnmarshal != nil {
			return nil, models.ErrInternalServerError
		}

		expPayment, err := m.paymentRepo.GetByExpID(ctx, exp.Id)
		if err != nil {
			return nil, err
		}

		var currency string
		if expPayment[0].Currency == 1 {
			currency = "USD"
		} else {
			currency = "IDR"
		}

		var priceItemType string
		if expPayment[0].PriceItemType == 1 {
			priceItemType = "Per Pax"
		} else {
			priceItemType = "Per Trip"
		}

		countRating, err := m.reviewsRepo.CountRating(ctx, 0, exp.Id)
		if err != nil {
			return nil, err
		}

		results[i] = &models.ExpSearchObject{
			Id:          exp.Id,
			ExpTitle:    exp.ExpTitle,
			ExpType:     expType,
			Rating:      exp.Rating,
			CountRating: countRating,
			Currency:    currency,
			Price:       expPayment[0].Price,
			PaymentType: priceItemType,
		}
	}

	return results, nil
}

func (m experienceUsecase) GetUserDiscoverPreference(ctx context.Context, page *int, size *int, currencyPrice string) ([]*models.ExpUserDiscoverPreferenceDto, error) {
	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()

	tempUserPreference, err := m.tempUserPreRepo.GetAll(ctx, page, size)
	if err != nil {
		return nil, err
	}
	var expListDto []*models.ExpUserDiscoverPreferenceDto
	for _, tup := range tempUserPreference {
		expList, err := m.experienceRepo.GetUserDiscoverPreferenceByHarborsIdOrProvince(ctx, tup.HarborsId, tup.ProvinceId)
		if err != nil {
			return nil, err
		}
		for _, element := range expList {

			var expType []string
			if errUnmarshal := json.Unmarshal([]byte(element.ExpType), &expType); errUnmarshal != nil {
				return nil, models.ErrInternalServerError
			}
			countRating, err := m.reviewsRepo.CountRating(ctx, 0, element.Id)
			if err != nil {
				return nil, err
			}
			//expPhotos, err := m.expPhotos.GetByExperienceID(ctx,element.Id)
			//if err != nil {
			//	return nil, models.ErrInternalServerError
			//}

			var coverPhotos models.CoverPhotosObj
			cityPhotos := make([]models.CoverPhotosObj, 0)
			if element.ExpCoverPhoto != nil {
				covertPhoto := models.CoverPhotosObj{
					Original:  *element.ExpCoverPhoto,
					Thumbnail: "",
				}
				coverPhotos = covertPhoto
				//if errUnmarshal := json.Unmarshal([]byte(expPhotos[0].ExpPhotoImage), &coverPhotos); errUnmarshal != nil {
				//	return nil,models.ErrInternalServerError
				//}
			}
			if element.CityPhotos != nil {
				if errUnmarshal := json.Unmarshal([]byte(*element.CityPhotos), &cityPhotos); errUnmarshal != nil {
					return nil, models.ErrInternalServerError
				}
			}

			expPayment, err := m.paymentRepo.GetByExpID(ctx, element.Id)
			if err != nil {
				return nil, err
			}

			var priceItemType string
			var currency string
			if expPayment != nil {
				if expPayment[0].Currency == 1 {
					currency = "USD"
				} else {
					currency = "IDR"
				}

				if expPayment[0].PriceItemType == 1 {
					priceItemType = "Per Pax"
				} else {
					priceItemType = "Per Trip"
				}

				if currencyPrice == "USD" {
					if currency == "IDR" {
						convertCurrency, _ := m.currencyUsecase.ExchangeRatesApi(ctx, "IDR", "USD")
						calculatePrice := convertCurrency.Rates.USD * expPayment[0].Price
						expPayment[0].Price = calculatePrice
						currency = "USD"
					}
				} else if currencyPrice == "IDR" {
					if currency == "USD" {
						convertCurrency, _ := m.currencyUsecase.ExchangeRatesApi(ctx, "USD", "IDR")
						calculatePrice := convertCurrency.Rates.IDR * expPayment[0].Price
						expPayment[0].Price = calculatePrice
						currency = "IDR"
					}
				}
			} else {
				priceItemType = ""
				currency = ""
			}

			//expDto := models.ExperienceUserDiscoverPreferenceDto{}

			if len(expListDto) == 0 {
				cityDto := models.ExpUserDiscoverPreferenceDto{
					ProvinceId:   element.ProvinceId,
					ProvinceName: element.ProvinceName,
					CityId:       element.CityId,
					City:         element.CityName,
					CityDesc:     element.CityDesc,
					Item:         nil,
					CityPhotos:   cityPhotos,
					HarborsId:    element.IdHarbors,
					HarborsName:  element.HarborsName,
				}
				expDto := models.ExperienceUserDiscoverPreferenceDto{
					Id:          element.Id,
					ExpTitle:    element.ExpTitle,
					ExpType:     expType,
					Rating:      element.Rating,
					CountRating: countRating,
					Currency:    currency,
					//Price:        expPayment[0].Price,
					Payment_type: priceItemType,
					Cover_Photo:  coverPhotos,
				}
				if expPayment != nil {
					expDto.Price = expPayment[0].Price
				}
				cityDto.Item = append(cityDto.Item, expDto)
				expListDto = append(expListDto, &cityDto)
			} else if len(expListDto) != 0 {
				var searchDto *models.ExpUserDiscoverPreferenceDto
				for _, dto := range expListDto {
					if dto.CityId == element.CityId {
						searchDto = dto
					}
				}
				if searchDto == nil {
					cityDto := models.ExpUserDiscoverPreferenceDto{
						ProvinceId:   element.ProvinceId,
						ProvinceName: element.ProvinceName,
						CityId:       element.CityId,
						City:         element.CityName,
						CityDesc:     element.CityDesc,
						Item:         nil,
						CityPhotos:   cityPhotos,
						HarborsId:    element.IdHarbors,
						HarborsName:  element.HarborsName,
					}
					expDto := models.ExperienceUserDiscoverPreferenceDto{
						Id:          element.Id,
						ExpTitle:    element.ExpTitle,
						ExpType:     expType,
						Rating:      element.Rating,
						CountRating: countRating,
						Currency:    currency,
						//Price:        expPayment[0].Price,
						Payment_type: priceItemType,
						Cover_Photo:  coverPhotos,
					}
					if expPayment != nil {
						expDto.Price = expPayment[0].Price
					}
					cityDto.Item = append(cityDto.Item, expDto)
					expListDto = append(expListDto, &cityDto)
				} else {
					for _, dto := range expListDto {
						if dto.CityId == element.CityId {
							expDto := models.ExperienceUserDiscoverPreferenceDto{
								Id:          element.Id,
								ExpTitle:    element.ExpTitle,
								ExpType:     expType,
								Rating:      element.Rating,
								CountRating: countRating,
								Currency:    currency,
								//Price:        expPayment[0].Price,
								Payment_type: priceItemType,
								Cover_Photo:  coverPhotos,
							}
							if expPayment != nil {
								expDto.Price = expPayment[0].Price
							}
							dto.Item = append(dto.Item, expDto)
						}
					}
				}
			}
		}
	}

	return expListDto, nil
}

func (m experienceUsecase) FilterSearchExp(
	ctx context.Context,
	isMerchant bool,
	search,
	token,
	qStatus,
	cityID string,
	harborsId string,
	activityType string,
	startDate string,
	endDate string,
	guest string,
	trip string,
	bottomPrice string,
	upPrice string,
	sortBy string,
	page int,
	limit int,
	offset int,
	provinceId string,
	currencyPrice string,
	bookingType string,
	paymentType string,
) (*models.FilterSearchWithPagination, error) {
	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()

	var activityTypeArray []int
	if activityType != "" && activityType != "[]" {
		if errUnmarshal := json.Unmarshal([]byte(activityType), &activityTypeArray); errUnmarshal != nil {
			return nil, models.ErrInternalServerError
		}
	}
	query := `
	select distinct
		e.id,
		e.exp_title,
		e.exp_type,
		e.status as exp_status,
		e.rating,
		e.exp_location_latitude as latitude,
		e.exp_location_longitude as longitude,
		e.exp_cover_photo as cover_photo,
		province_name AS province,
		e.exp_location_map_name,
		e.exp_latitude_map,
		e.exp_longitude_map
	from 
		experiences e
	JOIN harbors ha ON e.harbors_id = ha.id
	JOIN cities ci ON ha.city_id = ci.id
	JOIN provinces p ON ci.province_id = p.id
	JOIN experience_payments ep on ep.exp_id = e.id`

	qCount := `
	select COUNT(*) from experiences e
	JOIN harbors ha ON e.harbors_id = ha.id
	JOIN cities ci ON ha.city_id = ci.id
	JOIN provinces p ON ci.province_id = p.id
	JOIN experience_payments ep on ep.exp_id = e.id`

	if bottomPrice != "" && upPrice != "" && qStatus != "draft" {
		//query = ", ep.price" + query[:(strings.Index(query,"from"))]
		//qCount = ", ep.price" + query[:(strings.Index(query,"from"))]
		//query = query + ` join experience_payments ep on ep.exp_id = e.id`

		//qCount = qCount + ` join experience_payments ep on ep.exp_id = e.id`
	}
	if startDate != "" && endDate != "" && qStatus != "draft" {
		query = query + ` join exp_availabilities ead on ead.exp_id = e.id`
		qCount = qCount + ` join exp_availabilities ead on ead.exp_id = e.id`
	}
	//if len(activityTypeArray) != 0 {
	//	query = query + ` join filter_activity_types fat on fat.exp_id = e.id`
	//	qCount = qCount + ` join filter_activity_types fat on fat.exp_id = e.id`
	//}
	if cityID != "" && qStatus != "draft" {
		query = query + ` join harbors h on h.id = e.harbors_id`
		qCount = qCount + ` join harbors h on h.id = e.harbors_id`
	}

	query = query + ` WHERE ep.is_deleted = 0 AND ep.is_active = 1 AND e.is_deleted = 0 AND e.is_active = 1 `
	qCount = qCount + ` WHERE ep.is_deleted = 0 AND ep.is_active =1  AND e.is_deleted = 0 AND e.is_active = 1 `

	if isMerchant {
		if token == "" {
			return nil, models.ErrUnAuthorize
		}

		currentMerchant, err := m.mUsecase.ValidateTokenMerchant(ctx, token)
		if err != nil {
			return nil, err
		}

		query = query + ` AND e.merchant_id = '` + currentMerchant.Id + `'`
		qCount = qCount + ` AND e.merchant_id = '` + currentMerchant.Id + `'`
	}

	if bottomPrice != "" && upPrice != "" {
		query = query + `AND ep.price between ` + bottomPrice + ` and ` + upPrice
		qCount = qCount + `AND ep.price between ` + bottomPrice + ` and ` + upPrice
	}

	if search != "" {
		keyword := `'%` + search + `%'`
		query = query + ` AND LOWER(e.exp_title) LIKE LOWER(` + keyword + `)`
		qCount = qCount + ` AND LOWER(e.exp_title) LIKE LOWER(` + keyword + `)`
	}
	if bookingType != ""{
		if bookingType == "instant"{
			query = query + ` AND e.exp_booking_type = 'Instant Booking' `
			qCount = qCount + ` AND e.exp_booking_type = 'Instant Booking' `
		}else if bookingType == "noinstant"{
			query = query + ` AND e.exp_booking_type = 'No Instant Booking' `
			qCount = qCount + ` AND e.exp_booking_type = 'No Instant Booking' `
		}
	}
	if paymentType != ""{
		if paymentType == "full"{
			query = query + ` AND ep.exp_payment_type_id = '8a5e3eef-a6db-4584-a280-af5ab18a979b' `
			qCount = qCount + ` AND ep.exp_payment_type_id = '8a5e3eef-a6db-4584-a280-af5ab18a979b' `
		}else if paymentType == "down"{
			query = query + ` AND ep.exp_payment_type_id = '86e71b8d-acc3-4ade-80c0-de67b9100633' `
			qCount = qCount + ` AND ep.exp_payment_type_id = '86e71b8d-acc3-4ade-80c0-de67b9100633' `
		}
	}
	if provinceId != "" {
		query = query + ` AND ci.province_id = '` + provinceId + `'`
		qCount = qCount + ` AND ci.province_id = '` + provinceId + `'`
	}
	if cityID != "" {
		city_id, _ := strconv.Atoi(cityID)
		query = query + ` AND h.city_id = ` + strconv.Itoa(city_id)
		qCount = qCount + ` AND h.city_id = ` + strconv.Itoa(city_id)
	} else if harborsId != "" {
		query = query + ` AND e.harbors_id = '` + harborsId + `'`
		qCount = qCount + ` AND e.harbors_id = '` + harborsId + `'`
	}
	if qStatus != "" {
		var status int
		if qStatus == "preview" {
			status = 0
		} else if qStatus == "draft" {
			status = 1
		} else if qStatus == "published" {
			status = 2
		} else if qStatus == "unpublished" {
			status = 3
		} else if qStatus == "archived" {
			status = 4
		}
		if qStatus == "inService" {
			query = query + ` AND e.status IN (2,3)`
			qCount = qCount + ` AND e.status IN (2,3)`
		} else {
			query = query + ` AND e.status =` + strconv.Itoa(status)
			qCount = qCount + ` AND e.status =` + strconv.Itoa(status)
		}

	}
	if guest != "" {
		guests, _ := strconv.Atoi(guest)
		query = query + ` AND e.exp_max_guest >=` + strconv.Itoa(guests)
		qCount = qCount + ` AND e.exp_max_guest >=` + strconv.Itoa(guests)
	}
	if trip != "" {
		trips, _ := strconv.Atoi(trip)
		var tripType string
		if trips == 0 {
			tripType = "Private Trip"
		} else if trips == 1 {
			tripType = "Share Trip"
		} else {
			return nil, models.ErrInternalServerError
		}
		query = query + ` AND e.exp_trip_type = '` + tripType + `'`
		qCount = qCount + ` AND e.exp_trip_type = '` + tripType + `'`
	}

	if len(activityTypeArray) != 0 {
		for index, id := range activityTypeArray {
			if index == 0 && index != (len(activityTypeArray)-1) {
				query = query + ` AND (e.id = (SELECT distinct exp_id FROM filter_activity_types where exp_id = e.id and exp_type_id = ` + strconv.Itoa(id) + ` )`
				qCount = qCount + ` AND (e.id = (SELECT distinct exp_id FROM filter_activity_types where exp_id = e.id and exp_type_id = ` + strconv.Itoa(id) + ` )`
			} else if index == 0 && index == (len(activityTypeArray)-1) {
				query = query + ` AND (e.id = (SELECT distinct exp_id FROM filter_activity_types where exp_id = e.id and exp_type_id = ` + strconv.Itoa(id) + ` )` + ` ) `
				qCount = qCount + ` AND (e.id = (SELECT distinct exp_id FROM filter_activity_types where exp_id = e.id and exp_type_id = ` + strconv.Itoa(id) + ` )` + ` ) `
			} else if index == (len(activityTypeArray) - 1) {
				query = query + ` OR e.id = (SELECT distinct exp_id FROM filter_activity_types where exp_id = e.id and exp_type_id = ` + strconv.Itoa(id) + ` )` + ` )`
				qCount = qCount + ` OR e.id = (SELECT distinct exp_id FROM filter_activity_types where exp_id = e.id and exp_type_id = ` + strconv.Itoa(id) + ` )` + ` )`
			} else {
				query = query + ` OR e.id = (SELECT distinct exp_id FROM filter_activity_types where exp_id = e.id and exp_type_id = ` + strconv.Itoa(id) + ` )`
				qCount = qCount + ` OR e.id = (SELECT distinct exp_id FROM filter_activity_types where exp_id = e.id and exp_type_id = ` + strconv.Itoa(id) + ` )`
			}
		}

	}
	if bottomPrice != "" && upPrice != "" && qStatus != "draft" {
		bottomprices, _ := strconv.ParseFloat(bottomPrice, 64)
		upprices, _ := strconv.ParseFloat(upPrice, 64)

		query = query + ` AND (ep.price between ` + fmt.Sprint(bottomprices) + ` AND ` + fmt.Sprint(upprices) + `)`
		qCount = qCount + ` AND (ep.price between ` + fmt.Sprint(bottomprices) + ` AND ` + fmt.Sprint(upprices) + `)`
		if sortBy != "" {
			if sortBy == "priceup" {
				query = query + ` ORDER BY ep.price DESC`
			} else if sortBy == "pricedown" {
				query = query + ` ORDER BY ep.price ASC`
			}
		}
	}
	if sortBy != "" {
		isOrderBy := strings.Index(query, "ORDER BY")
		if sortBy == "ratingup" {
			//query = query[:(strings.Index(query,"from"))] + ", ep.price " + query[(strings.Index(query,"from")):]
			//query = query[:strings.Index(query, "join")] + ` join experience_payments ep on ep.exp_id = e.id ` +
			query = query[:strings.Index(query, "JOIN")] + query[strings.Index(query, "JOIN"):] + " ORDER BY e.rating desc"
		} else if sortBy == "ratingdown" {
			//query = query[:(strings.Index(query,"from"))] + ", ep.price " + query[(strings.Index(query,"from")):]
			//query = query[:strings.Index(query, "join")] + ` join experience_payments ep on ep.exp_id = e.id ` +
			query = query[:strings.Index(query, "JOIN")] + query[strings.Index(query, "JOIN"):] + " ORDER BY e.rating asc"
		} else if sortBy == "newest" {
			query = query + ` ORDER BY e.created_date DESC`
		} else if sortBy == "latest" {
			query = query + ` ORDER BY e.created_date ASC`
		} else if sortBy == "priceup" && isOrderBy == -1 {
			//query = query[:(strings.Index(query,"from"))] + ", ep.price " + query[(strings.Index(query,"from")):]
			//qCount = qCount[:(strings.Index(qCount,"from"))] + ", ep.price " + qCount[(strings.Index(qCount,"from")):]
			//query = query[:strings.Index(query, "join")] + ` join experience_payments ep on ep.exp_id = e.id ` +
			query = query[:strings.Index(query, "JOIN")] + query[strings.Index(query, "JOIN"):] + " ORDER BY ep.price desc"
			//qCount = qCount[:strings.Index(qCount, "join")]  + ` join experience_payments ep on ep.exp_id = e.id` + qCount[strings.Index(qCount, "join"):]
		} else if sortBy == "pricedown" && isOrderBy == -1 {
			//query = query[:(strings.Index(query,"from"))] + ", ep.price " + query[(strings.Index(query,"from")):]
			//qCount = qCount[:(strings.Index(qCount,"from"))] + ", ep.price " + qCount[(strings.Index(qCount,"from")):]
			//query = query[:strings.Index(query, "join")] + ` join experience_payments ep on ep.exp_id = e.id ` +
			query = query[:strings.Index(query, "JOIN")] + query[strings.Index(query, "JOIN"):] + " ORDER BY ep.price asc"
			//qCount = qCount[:strings.Index(qCount, "join")]  + ` join experience_payments ep on ep.exp_id = e.id` + qCount[strings.Index(qCount, "join"):]
		}
	}

	if startDate != "" && endDate != "" {
		var startDates []string

		layoutFormat := "2006-01-02"
		start, errDateDob := time.Parse(layoutFormat, startDate)
		if errDateDob != nil {
			return nil, errDateDob
		}
		end, errDateDob := time.Parse(layoutFormat, endDate)
		if errDateDob != nil {
			return nil, errDateDob
		}
		startDates = append(startDates, start.Format("2006-01-02"))
	datess:

		start = start.AddDate(0, 0, 1)
		startDates = append(startDates, start.Format("2006-01-02"))
		if start == end {

		} else {
			startDates = append(startDates, start.String())
			goto datess
		}
		tmpQFilterDate := ""
		tmpQCount := ""
		for index, id := range startDates {

			if index == 0 && index != (len(startDates)-1) {
				//query = query + ` AND (ead.exp_availability_date like '%` + id + `%' `
				//qCount = qCount + ` AND (ead.exp_availability_date like '%` + id + `%' `
				tmpQFilterDate = tmpQFilterDate + ` AND (ead.exp_availability_date like '%` + id + `%' `
				tmpQCount = tmpQCount + ` AND (ead.exp_availability_date like '%` + id + `%' `
			} else if index == 0 && index == (len(startDates)-1) {
				//query = query + ` AND (ead.exp_availability_date like '%` + id + `%' ) `
				//qCount = qCount + ` AND (ead.exp_availability_date like '%` + id + `%' ) `
				tmpQFilterDate = tmpQFilterDate + ` AND (ead.exp_availability_date like '%` + id + `%' ) `
				tmpQCount = tmpQCount + ` AND (ead.exp_availability_date like '%` + id + `%' ) `
			} else if index == (len(startDates) - 1) {
				//query = query + ` OR ead.exp_availability_date like '%` + id + `%' ) `
				//qCount = qCount + ` OR ead.exp_availability_date like '%` + id + `%' ) `
				tmpQFilterDate = tmpQFilterDate + ` OR ead.exp_availability_date like '%` + id + `%' ) `
				tmpQCount = tmpQCount + ` OR ead.exp_availability_date like '%` + id + `%' ) `
			} else {
				//query = query + ` OR ead.exp_availability_date like '%` + id + `%' `
				//qCount = qCount + ` OR ead.exp_availability_date like '%` + id + `%' `
				tmpQFilterDate = tmpQFilterDate + ` OR ead.exp_availability_date like '%` + id + `%' `
				tmpQCount = tmpQCount + ` OR ead.exp_availability_date like '%` + id + `%' `
			}
		}
		query = query[:strings.Index(query, "AND")] + tmpQFilterDate + query[strings.Index(query, "AND"):]
		qCount = qCount[:strings.Index(qCount, "AND")] + tmpQCount + qCount[strings.Index(qCount, "AND"):]
	}
	expList, err := m.experienceRepo.QueryFilterSearch(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	results := make([]*models.ExpSearchObject, len(expList))
	for i, exp := range expList {
		var expType []string
		if errUnmarshal := json.Unmarshal([]byte(exp.ExpType), &expType); errUnmarshal != nil {
			return nil, models.ErrInternalServerError
		}

		expPayment, err := m.paymentRepo.GetByExpID(ctx, exp.Id)
		if err != nil {
			return nil, err
		}

		var currency string
		var price float64
		var priceItemType string

		if expPayment != nil {

			price = expPayment[0].Price
			if expPayment[0].Currency == 1 {
				currency = "USD"
			} else {
				currency = "IDR"
			}

			if expPayment[0].PriceItemType == 1 {
				priceItemType = "Per Pax"
			} else {
				priceItemType = "Per Trip"
			}

			if currencyPrice == "USD" {
				if currency == "IDR" {
					convertCurrency, _ := m.currencyUsecase.ExchangeRatesApi(ctx, "IDR", "USD")
					calculatePrice := convertCurrency.Rates.USD * price
					price = calculatePrice
					currency = "USD"
				}
			} else if currencyPrice == "IDR" {
				if currency == "USD" {
					convertCurrency, _ := m.currencyUsecase.ExchangeRatesApi(ctx, "USD", "IDR")
					calculatePrice := convertCurrency.Rates.IDR * price
					price = calculatePrice
					currency = "IDR"
				}
			}

		}
		countRating, err := m.reviewsRepo.CountRating(ctx, 0, exp.Id)
		if err != nil {
			return nil, err
		}
		coverPhoto := models.CoverPhotosObj{
			Original:  exp.CoverPhoto,
			Thumbnail: "",
		}
		var listPhotos []models.ExpPhotosObj
		expPhotoQuery, errorQuery := m.expPhotos.GetByExperienceID(ctx, exp.Id)
		if errorQuery != nil {
			return nil, errorQuery
		}
		if expPhotoQuery != nil {
			for _, element := range expPhotoQuery {
				expPhoto := models.ExpPhotosObj{
					Folder:        element.ExpPhotoFolder,
					ExpPhotoImage: nil,
				}
				var expPhotoImage []models.CoverPhotosObj
				errObject := json.Unmarshal([]byte(element.ExpPhotoImage), &expPhotoImage)
				if errObject != nil {
					return nil, models.ErrInternalServerError
				}
				expPhoto.ExpPhotoImage = expPhotoImage
				listPhotos = append(listPhotos, expPhoto)
			}
		}

		var transStatus string
		if exp.ExpStatus == 0 {
			transStatus = "Preview"
		} else if exp.ExpStatus == 1 {
			transStatus = "Draft"
		} else if exp.ExpStatus == 2 {
			transStatus = "Published"
		} else if exp.ExpStatus == 3 {
			transStatus = "Unpublished"
		} else if exp.ExpStatus == 4 {
			transStatus = "Archived"
		}

		results[i] = &models.ExpSearchObject{
			Id:                 exp.Id,
			ExpTitle:           exp.ExpTitle,
			ExpType:            expType,
			ExpStatus:          transStatus,
			Rating:             exp.Rating,
			CountRating:        countRating,
			Currency:           currency,
			Price:              price,
			PaymentType:        priceItemType,
			Longitude:          exp.Longitude,
			Latitude:           exp.Latitude,
			Province:           exp.Province,
			CoverPhoto:         coverPhoto,
			ListPhoto:          listPhotos,
			ExpLocationMapName: exp.ExpLocationMapName,
			ExpLatitudeMap:     exp.ExpLatitudeMap,
			ExpLongitudeMap:    exp.ExpLongitudeMap,
		}
	}
	qCount = "select count(*) from (" + query + ") q"
	totalRecords, _ := m.experienceRepo.CountFilterSearch(ctx, qCount)
	totalPage := int(math.Ceil(float64(totalRecords) / float64(limit)))
	prev := page
	next := page
	if page != 1 {
		prev = page - 1
	}

	if page != totalPage {
		next = page + 1
	}

	meta := &models.MetaPagination{
		Page:          page,
		Total:         totalPage,
		TotalRecords:  totalRecords,
		Prev:          prev,
		Next:          next,
		RecordPerPage: len(results),
	}

	response := &models.FilterSearchWithPagination{
		Data: results,
		Meta: meta,
	}

	return response, nil

}
func (m experienceUsecase) GetExpInspirations(ctx context.Context) ([]*models.ExpInspirationDto, error) {
	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()

	query, err := m.inspirationRepo.GetExpInspirations(ctx)
	var results []*models.ExpInspirationDto
	for _, element := range query {
		getCountReview, err := m.reviewsRepo.CountRating(ctx, 0, element.ExpId)
		if err != nil {
			return nil, err
		}
		var expType []string
		if element.ExpType != "" {
			if errUnmarshal := json.Unmarshal([]byte(element.ExpType), &expType); errUnmarshal != nil {
				return nil, models.ErrInternalServerError
			}
		}
		dto := models.ExpInspirationDto{
			ExpInspirationID: element.ExpInspirationID,
			ExpId:            element.ExpId,
			ExpTitle:         element.ExpTitle,
			ExpDesc:          element.ExpDesc,
			ExpCoverPhoto:    element.ExpCoverPhoto,
			ExpType:          expType,
			Rating:           element.Rating,
			CountRating:      getCountReview,
		}
		results = append(results, &dto)
	}
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (m experienceUsecase) GetExpTypes(ctx context.Context) ([]*models.ExpTypeObject, error) {
	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()

	results, err := m.typesRepo.GetExpTypes(ctx)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (m experienceUsecase) SearchExp(ctx context.Context, harborID, cityID string) ([]*models.ExpSearchObject, error) {
	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()

	expList, err := m.experienceRepo.SearchExp(ctx, harborID, cityID)
	if err != nil {
		return nil, err
	}

	results := make([]*models.ExpSearchObject, len(expList))
	for i, exp := range expList {
		var expType []string
		if errUnmarshal := json.Unmarshal([]byte(exp.ExpType), &expType); errUnmarshal != nil {
			return nil, models.ErrInternalServerError
		}
		expPayment, err := m.paymentRepo.GetByExpID(ctx, exp.Id)
		if err != nil {
			return nil, err
		}

		var currency string
		if expPayment[0].Currency == 1 {
			currency = "USD"
		} else {
			currency = "IDR"
		}

		var priceItemType string
		if expPayment[0].PriceItemType == 1 {
			priceItemType = "Per Pax"
		} else {
			priceItemType = "Per Trip"
		}

		countRating, err := m.reviewsRepo.CountRating(ctx, 0, exp.Id)
		if err != nil {
			return nil, err
		}

		var listPhotos []models.ExpPhotosObj
		var coverPhotos models.CoverPhotosObj
		if exp.CoverPhoto != "" {
			coverPhotos.Original = exp.CoverPhoto
			coverPhotos.Thumbnail = ""
		}
		expPhotoQuery, errorQuery := m.expPhotos.GetByExperienceID(ctx, exp.Id)
		if errorQuery != nil {
			return nil, errorQuery
		}
		if expPhotoQuery != nil {
			for _, element := range expPhotoQuery {
				expPhoto := models.ExpPhotosObj{
					Folder:        element.ExpPhotoFolder,
					ExpPhotoImage: nil,
				}
				var expPhotoImage []models.CoverPhotosObj
				errObject := json.Unmarshal([]byte(element.ExpPhotoImage), &expPhotoImage)
				if errObject != nil {
					//fmt.Println("Error : ",err.Error())
					return nil, models.ErrInternalServerError
				}
				expPhoto.ExpPhotoImage = expPhotoImage
				listPhotos = append(listPhotos, expPhoto)
			}
		}
		results[i] = &models.ExpSearchObject{
			Id:          exp.Id,
			ExpTitle:    exp.ExpTitle,
			ExpType:     expType,
			Rating:      exp.Rating,
			CountRating: countRating,
			Currency:    currency,
			Price:       expPayment[0].Price,
			PaymentType: priceItemType,
			CoverPhoto:  coverPhotos,
			ListPhoto:   listPhotos,
		}
	}

	return results, nil
}

func (m experienceUsecase) CreateExperience(c context.Context, commandExperience models.NewCommandExperience, token string) (*models.ResponseCreateExperience, error) {
	ctx, cancel := context.WithTimeout(c, m.contextTimeout)
	defer cancel()
	currentUserMerchant, err := m.mUsecase.ValidateTokenMerchant(ctx, token)
	if err != nil {
		return nil, err
	}

	//if commandExperience.ExpType != ""
	expItinerary, _ := json.Marshal(commandExperience.ExpInternary)
	expFacilities, _ := json.Marshal(commandExperience.ExpFacilities)
	expInclusion, _ := json.Marshal(commandExperience.ExpInclusion)
	expRules, _ := json.Marshal(commandExperience.ExpRules)
	expTypes, _ := json.Marshal(commandExperience.ExpType)
	experiences := models.Experience{
		Id:                       "",
		CreatedBy:                currentUserMerchant.MerchantEmail,
		CreatedDate:              time.Time{},
		ModifiedBy:               nil,
		ModifiedDate:             nil,
		DeletedBy:                nil,
		DeletedDate:              nil,
		IsDeleted:                0,
		IsActive:                 0,
		ExpTitle:                 commandExperience.ExpTitle,
		ExpType:                  string(expTypes),
		ExpTripType:              commandExperience.ExpTripType,
		ExpBookingType:           commandExperience.ExpBookingType,
		ExpDesc:                  commandExperience.ExpDesc,
		ExpMaxGuest:              commandExperience.ExpMaxGuest,
		ExpPickupPlace:           commandExperience.ExpPickupPlace,
		ExpPickupTime:            commandExperience.ExpPickupTime,
		ExpPickupPlaceLongitude:  commandExperience.ExpPickupPlaceLongitude,
		ExpPickupPlaceLatitude:   commandExperience.ExpPickupPlaceLatitude,
		ExpPickupPlaceMapsName:   commandExperience.ExpPickupPlaceMapsName,
		ExpInternary:             string(expItinerary),
		ExpFacilities:            string(expFacilities),
		ExpInclusion:             string(expInclusion),
		ExpRules:                 string(expRules),
		Status:                   commandExperience.Status,
		Rating:                   0,
		ExpLocationLatitude:      commandExperience.ExpLocationLatitude,
		ExpLocationLongitude:     commandExperience.ExpLocationLongitude,
		ExpLocationName:          commandExperience.ExpLocationName,
		ExpCoverPhoto:            commandExperience.ExpCoverPhoto,
		ExpDuration:              commandExperience.ExpDuration,
		MinimumBookingId:         &commandExperience.MinimumBookingId,
		MerchantId:               currentUserMerchant.Id,
		HarborsId:                &commandExperience.HarborsId,
		ExpPaymentDeadlineAmount: &commandExperience.ExpPaymentDeadlineAmount,
		ExpPaymentDeadlineType:   &commandExperience.ExpPaymentDeadlineType,
		IsCustomisedByUser:       &commandExperience.IsCustomisedByUser,
		ExpLocationMapName:       commandExperience.ExpLocationMapName,
		ExpLatitudeMap:           commandExperience.ExpLatitudeMap,
		ExpLongitudeMap:          commandExperience.ExpLongitudeMap,
		ExpMaximumBookingAmount:commandExperience.ExpMaximumBookingAmount,
		ExpMaximumBookingType:commandExperience.ExpMaximumBookingType,
	}
	if *experiences.HarborsId == "" && experiences.Status == 1 {
		experiences.HarborsId = nil
	}
	if *experiences.MinimumBookingId == "" && experiences.Status == 1 {
		experiences.MinimumBookingId = nil
	}
	insertToExperience, err := m.experienceRepo.Insert(ctx, &experiences)

	for _, element := range commandExperience.ExpType {
		getExpType, err := m.typesRepo.GetByName(ctx, element)
		if err != nil {
			return nil, err
		}
		filterActivityT := models.FilterActivityType{
			Id:           0,
			CreatedBy:    currentUserMerchant.MerchantEmail,
			CreatedDate:  time.Now(),
			ModifiedBy:   nil,
			ModifiedDate: nil,
			DeletedBy:    nil,
			DeletedDate:  nil,
			IsDeleted:    0,
			IsActive:     1,
			ExpTypeId:    getExpType.ExpTypeID,
			ExpId:        insertToExperience,
		}
		insertToFilterAT := m.filterATRepo.Insert(ctx, &filterActivityT)
		if insertToFilterAT != nil {
			return nil, insertToFilterAT
		}
	}

	for _, element := range commandExperience.ExpPhotos {
		images, _ := json.Marshal(element.ExpPhotoImage)
		expPhoto := models.ExpPhotos{
			Id:             guuid.New().String(),
			CreatedBy:      currentUserMerchant.MerchantEmail,
			CreatedDate:    time.Now(),
			ModifiedBy:     nil,
			ModifiedDate:   nil,
			DeletedBy:      nil,
			DeletedDate:    nil,
			IsDeleted:      0,
			IsActive:       0,
			ExpPhotoFolder: element.Folder,
			ExpPhotoImage:  string(images),
			ExpId:          *insertToExperience,
		}

		id, err := m.expPhotos.Insert(ctx, &expPhoto)
		if err != nil {
			return nil, err
		}
		element.Id = *id
	}

	for _, element := range commandExperience.ExpPayment {
		var priceItemType int
		if element.PriceItemType == "Per Pax" {
			priceItemType = 1
		} else {
			priceItemType = 0
		}
		var currency int
		if element.Currency == "USD" {
			currency = 1
		} else {
			currency = 0
		}
		var customPriceJson string
		if len(element.CustomPrice) != 0 {
			customPrice, _ := json.Marshal(element.CustomPrice)
			customPriceJson = string(customPrice)
		}

		payments := models.ExperiencePayment{
			Id:               guuid.New().String(),
			CreatedBy:        currentUserMerchant.MerchantEmail,
			CreatedDate:      time.Now(),
			ModifiedBy:       nil,
			ModifiedDate:     nil,
			DeletedBy:        nil,
			DeletedDate:      nil,
			IsDeleted:        0,
			IsActive:         0,
			ExpPaymentTypeId: element.PaymentTypeId,
			ExpId:            *insertToExperience,
			PriceItemType:    priceItemType,
			Currency:         currency,
			Price:            element.Price,
			CustomPrice:      &customPriceJson,
		}

		id, err := m.paymentRepo.Insert(ctx, payments)
		if err != nil {
			return nil, err
		}
		element.Id = id
	}

	for _, element := range commandExperience.ExpAvailability {
		date, _ := json.Marshal(element.Date)
		expAvailability := models.ExpAvailability{
			Id:                   guuid.New().String(),
			CreatedBy:            currentUserMerchant.MerchantEmail,
			CreatedDate:          time.Now(),
			ModifiedBy:           nil,
			ModifiedDate:         nil,
			DeletedBy:            nil,
			DeletedDate:          nil,
			IsDeleted:            0,
			IsActive:             0,
			ExpAvailabilityMonth: element.Month,
			ExpAvailabilityDate:  string(date),
			ExpAvailabilityYear:  element.Year,
			ExpId:                *insertToExperience,
		}

		id, err := m.exp_availablitiy.Insert(ctx, expAvailability)
		if err != nil {
			return nil, err
		}
		element.Id = id
	}

	for _, element := range commandExperience.ExperienceAddOn {
		var currency int
		if element.Currency == "USD" {
			currency = 1
		} else {
			currency = 0
		}
		addOns := models.ExperienceAddOn{
			Id:           guuid.New().String(),
			CreatedBy:    currentUserMerchant.MerchantEmail,
			CreatedDate:  time.Now(),
			ModifiedBy:   nil,
			ModifiedDate: nil,
			DeletedBy:    nil,
			DeletedDate:  nil,
			IsDeleted:    0,
			IsActive:     0,
			Name:         element.Name,
			Desc:         element.Desc,
			Currency:     currency,
			Amount:       element.Amount,
			ExpId:        *insertToExperience,
		}
		id, err := m.adOnsRepo.Insert(ctx, addOns)
		if err != nil {
			return nil, err
		}
		element.Id = id
	}

	for _, element := range commandExperience.ExpFacilities {
		getFacilitiesByName, err := m.facilitiesRepo.GetByName(ctx, element.Name)
		if err != nil {
			return nil, err
		}
		facilities := models.ExperienceFacilities{
			Id:           0,
			ExpId:        insertToExperience,
			TransId:      nil,
			FacilitiesId: getFacilitiesByName.Id,
			Amount:       element.Amount,
		}
		err = m.expFacilitiesRepo.Insert(ctx, &facilities)
		if err != nil {
			return nil, err
		}
	}

	for _, element := range commandExperience.ExpInclusion {
		if element.Type == 0 {
			getIncludeById, err := m.includeRepo.GetByName(ctx, element.Name)
			if err != nil {
				return nil, err
			}
			include := models.ExperienceInclude{
				Id:        0,
				ExpId:     *insertToExperience,
				IncludeId: getIncludeById.Id,
			}
			err = m.expIncludeRepo.Insert(ctx, &include)
			if err != nil {
				return nil, err
			}
		} else if element.Type == 1 {
			getExcludeById, err := m.excludeRepo.GetByName(ctx, element.Name)
			if err != nil {
				return nil, err
			}
			exclude := models.ExperienceExclude{
				Id:        0,
				ExpId:     *insertToExperience,
				ExcludeId: getExcludeById.Id,
			}
			err = m.expExcludeRepo.Insert(ctx, &exclude)
			if err != nil {
				return nil, err
			}
		}
	}

	var status string
	if commandExperience.Status == 1 {
		status = "Draft"
	} else if commandExperience.Status == 2 {
		status = "Publish"
	}
	response := models.ResponseCreateExperience{
		Id:      *insertToExperience,
		Message: "Success " + status,
	}
	return &response, nil

}
func (m experienceUsecase) UpdateExperience(c context.Context, commandExperience models.NewCommandExperience, token string) (*models.ResponseCreateExperience, error) {
	ctx, cancel := context.WithTimeout(c, m.contextTimeout)
	defer cancel()
	currentUserMerchant, err := m.mUsecase.ValidateTokenMerchant(ctx, token)
	if err != nil {
		return nil, err
	}

	//if commandExperience.ExpType != ""
	expItinerary, _ := json.Marshal(commandExperience.ExpInternary)
	expFacilities, _ := json.Marshal(commandExperience.ExpFacilities)
	expInclusion, _ := json.Marshal(commandExperience.ExpInclusion)
	expRules, _ := json.Marshal(commandExperience.ExpRules)
	expTypes, _ := json.Marshal(commandExperience.ExpType)
	experiences := models.Experience{
		Id:                       commandExperience.Id,
		CreatedBy:                currentUserMerchant.MerchantEmail,
		CreatedDate:              time.Time{},
		ModifiedBy:               &currentUserMerchant.MerchantEmail,
		ModifiedDate:             &time.Time{},
		DeletedBy:                nil,
		DeletedDate:              nil,
		IsDeleted:                0,
		IsActive:                 0,
		ExpTitle:                 commandExperience.ExpTitle,
		ExpType:                  string(expTypes),
		ExpTripType:              commandExperience.ExpTripType,
		ExpBookingType:           commandExperience.ExpBookingType,
		ExpDesc:                  commandExperience.ExpDesc,
		ExpMaxGuest:              commandExperience.ExpMaxGuest,
		ExpPickupPlace:           commandExperience.ExpPickupPlace,
		ExpPickupTime:            commandExperience.ExpPickupTime,
		ExpPickupPlaceLongitude:  commandExperience.ExpPickupPlaceLongitude,
		ExpPickupPlaceLatitude:   commandExperience.ExpPickupPlaceLatitude,
		ExpPickupPlaceMapsName:   commandExperience.ExpPickupPlaceMapsName,
		ExpInternary:             string(expItinerary),
		ExpFacilities:            string(expFacilities),
		ExpInclusion:             string(expInclusion),
		ExpRules:                 string(expRules),
		Status:                   commandExperience.Status,
		Rating:                   0,
		ExpLocationLatitude:      commandExperience.ExpLocationLatitude,
		ExpLocationLongitude:     commandExperience.ExpLocationLongitude,
		ExpLocationName:          commandExperience.ExpLocationName,
		ExpCoverPhoto:            commandExperience.ExpCoverPhoto,
		ExpDuration:              commandExperience.ExpDuration,
		MinimumBookingId:         &commandExperience.MinimumBookingId,
		MerchantId:               currentUserMerchant.Id,
		HarborsId:                &commandExperience.HarborsId,
		ExpPaymentDeadlineAmount: &commandExperience.ExpPaymentDeadlineAmount,
		ExpPaymentDeadlineType:   &commandExperience.ExpPaymentDeadlineType,
		IsCustomisedByUser:       &commandExperience.IsCustomisedByUser,
		ExpLocationMapName:       commandExperience.ExpLocationMapName,
		ExpLatitudeMap:           commandExperience.ExpLatitudeMap,
		ExpLongitudeMap:          commandExperience.ExpLongitudeMap,
		ExpMaximumBookingAmount:commandExperience.ExpMaximumBookingAmount,
		ExpMaximumBookingType:commandExperience.ExpMaximumBookingType,
	}
	if *experiences.HarborsId == "" && experiences.Status == 1 {
		experiences.HarborsId = nil
	}
	if *experiences.MinimumBookingId == "" && experiences.Status == 1 {
		experiences.MinimumBookingId = nil
	}
	insertToExperience, err := m.experienceRepo.Update(ctx, &experiences)

	if err != nil {
		return nil, err
	}
	err = m.filterATRepo.DeleteByExpId(ctx, experiences.Id)
	if err != nil {
		return nil, err
	}
	for _, element := range commandExperience.ExpType {
		getExpType, err := m.typesRepo.GetByName(ctx, element)
		if err != nil {
			return nil, err
		}
		filterActivityT := models.FilterActivityType{
			Id:           0,
			CreatedBy:    currentUserMerchant.MerchantEmail,
			CreatedDate:  time.Now(),
			ModifiedBy:   nil,
			ModifiedDate: nil,
			DeletedBy:    nil,
			DeletedDate:  nil,
			IsDeleted:    0,
			IsActive:     1,
			ExpTypeId:    getExpType.ExpTypeID,
			ExpId:        insertToExperience,
		}
		insertToFilterAT := m.filterATRepo.Insert(ctx, &filterActivityT)
		if insertToFilterAT != nil {
			return nil, insertToFilterAT
		}
	}

	var photoIds []string
	err = m.expPhotos.DeleteByExpId(ctx, experiences.Id, currentUserMerchant.MerchantEmail)
	if err != nil {
		return nil, err
	}
	for _, element := range commandExperience.ExpPhotos {
		if element.Id == "" {
			images, _ := json.Marshal(element.ExpPhotoImage)
			expPhoto := models.ExpPhotos{
				Id:             guuid.New().String(),
				CreatedBy:      currentUserMerchant.MerchantEmail,
				CreatedDate:    time.Now(),
				ModifiedBy:     nil,
				ModifiedDate:   nil,
				DeletedBy:      nil,
				DeletedDate:    nil,
				IsDeleted:      0,
				IsActive:       0,
				ExpPhotoFolder: element.Folder,
				ExpPhotoImage:  string(images),
				ExpId:          *insertToExperience,
			}

			id, err := m.expPhotos.Insert(ctx, &expPhoto)
			if err != nil {
				return nil, err
			}
			photoIds = append(photoIds, *id)
			element.Id = *id
		}
	}

	var expPaymentIds []string
	err = m.paymentRepo.DeleteByExpId(ctx, experiences.Id, currentUserMerchant.MerchantEmail)
	if err != nil {
		return nil, err
	}
	for _, element := range commandExperience.ExpPayment {
		var priceItemType int
		if element.PriceItemType == "Per Pax" {
			priceItemType = 1
		} else {
			priceItemType = 0
		}
		var currency int
		if element.Currency == "USD" {
			currency = 1
		} else {
			currency = 0
		}
		var customPriceJson string
		if len(element.CustomPrice) != 0 {
			customPrice, _ := json.Marshal(element.CustomPrice)
			customPriceJson = string(customPrice)
		}
		payments := models.ExperiencePayment{
			Id:               guuid.New().String(),
			CreatedBy:        currentUserMerchant.MerchantEmail,
			CreatedDate:      time.Now(),
			ModifiedBy:       nil,
			ModifiedDate:     nil,
			DeletedBy:        nil,
			DeletedDate:      nil,
			IsDeleted:        0,
			IsActive:         0,
			ExpPaymentTypeId: element.PaymentTypeId,
			ExpId:            *insertToExperience,
			PriceItemType:    priceItemType,
			Currency:         currency,
			Price:            element.Price,
			CustomPrice:      &customPriceJson,
		}

		id, err := m.paymentRepo.Insert(ctx, payments)
		if err != nil {
			return nil, err
		}
		expPaymentIds = append(expPaymentIds, id)
		element.Id = id
	}

	var expAvailabilityIds []string
	err = m.exp_availablitiy.DeleteByExpId(ctx, experiences.Id, currentUserMerchant.MerchantEmail)
	if err != nil {
		return nil, err
	}
	for _, element := range commandExperience.ExpAvailability {
		date, _ := json.Marshal(element.Date)
		expAvailability := models.ExpAvailability{
			Id:                   guuid.New().String(),
			CreatedBy:            currentUserMerchant.MerchantEmail,
			CreatedDate:          time.Now(),
			ModifiedBy:           nil,
			ModifiedDate:         nil,
			DeletedBy:            nil,
			DeletedDate:          nil,
			IsDeleted:            0,
			IsActive:             0,
			ExpAvailabilityMonth: element.Month,
			ExpAvailabilityDate:  string(date),
			ExpAvailabilityYear:  element.Year,
			ExpId:                *insertToExperience,
		}

		id, err := m.exp_availablitiy.Insert(ctx, expAvailability)
		if err != nil {
			return nil, err
		}
		expAvailabilityIds = append(expAvailabilityIds, id)
		element.Id = id
	}

	var addOnIds []string
	err = m.adOnsRepo.DeleteByExpId(ctx, experiences.Id, currentUserMerchant.MerchantEmail)
	if err != nil {
		return nil, err
	}
	for _, element := range commandExperience.ExperienceAddOn {
		var currency int
		if element.Currency == "USD" {
			currency = 1
		} else {
			currency = 0
		}
		addOns := models.ExperienceAddOn{
			Id:           guuid.New().String(),
			CreatedBy:    currentUserMerchant.MerchantEmail,
			CreatedDate:  time.Now(),
			ModifiedBy:   nil,
			ModifiedDate: nil,
			DeletedBy:    nil,
			DeletedDate:  nil,
			IsDeleted:    0,
			IsActive:     0,
			Name:         element.Name,
			Desc:         element.Desc,
			Currency:     currency,
			Amount:       element.Amount,
			ExpId:        *insertToExperience,
		}
		id, err := m.adOnsRepo.Insert(ctx, addOns)
		if err != nil {
			return nil, err
		}
		addOnIds = append(addOnIds, id)
		element.Id = id
	}

	err = m.expFacilitiesRepo.Delete(ctx, *insertToExperience, "")
	if err != nil {
		return nil, err
	}
	for _, element := range commandExperience.ExpFacilities {
		getFacilitiesByName, err := m.facilitiesRepo.GetByName(ctx, element.Name)
		if err != nil {
			return nil, err
		}
		facilities := models.ExperienceFacilities{
			Id:           0,
			ExpId:        insertToExperience,
			TransId:      nil,
			FacilitiesId: getFacilitiesByName.Id,
			Amount:       element.Amount,
		}
		err = m.expFacilitiesRepo.Insert(ctx, &facilities)
		if err != nil {
			return nil, err
		}
	}
	err = m.expIncludeRepo.Delete(ctx, *insertToExperience)
	if err != nil {
		return nil, err
	}
	err = m.expExcludeRepo.Delete(ctx, *insertToExperience)
	if err != nil {
		return nil, err
	}
	for _, element := range commandExperience.ExpInclusion {
		if element.Type == 0 {
			getIncludeById, err := m.includeRepo.GetByName(ctx, element.Name)
			if err != nil {
				return nil, err
			}
			include := models.ExperienceInclude{
				Id:        0,
				ExpId:     *insertToExperience,
				IncludeId: getIncludeById.Id,
			}
			err = m.expIncludeRepo.Insert(ctx, &include)
			if err != nil {
				return nil, err
			}
		} else if element.Type == 1 {
			getExcludeById, err := m.excludeRepo.GetByName(ctx, element.Name)
			if err != nil {
				return nil, err
			}
			exclude := models.ExperienceExclude{
				Id:        0,
				ExpId:     *insertToExperience,
				ExcludeId: getExcludeById.Id,
			}
			err = m.expExcludeRepo.Insert(ctx, &exclude)
			if err != nil {
				return nil, err
			}
		}
	}

	var status string
	if commandExperience.Status == 1 {
		status = "Draft"
	} else if commandExperience.Status == 2 {
		status = "Publish"
	}
	response := models.ResponseCreateExperience{
		Id:      *insertToExperience,
		Message: "Success " + status,
	}
	return &response, nil

}
func (m experienceUsecase) PublishExperience(c context.Context, commandExperience models.NewCommandExperience, token string) (*models.ResponseCreateExperience, error) {
	ctx, cancel := context.WithTimeout(c, m.contextTimeout)
	defer cancel()
	var response *models.ResponseCreateExperience
	if commandExperience.Id == "" {
		create, err := m.CreateExperience(ctx, commandExperience, token)
		if err != nil {
			return nil, err
		}
		response = create
	} else {
		update, err := m.UpdateExperience(ctx, commandExperience, token)
		if err != nil {
			return nil, err
		}
		response = update
	}
	return response, nil
}
func (m experienceUsecase) GetByID(c context.Context, id string, currencyPrice string,isMerchant string) (*models.ExperienceDto, error) {
	ctx, cancel := context.WithTimeout(c, m.contextTimeout)
	defer cancel()

	res, err := m.experienceRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	var expPhotos []models.ExpPhotosObj
	expPhotoQuery, errorQuery := m.expPhotos.GetByExperienceID(ctx, res.Id)
	if expPhotoQuery != nil {
		for _, element := range expPhotoQuery {
			expPhoto := models.ExpPhotosObj{
				Id:            element.Id,
				Folder:        element.ExpPhotoFolder,
				ExpPhotoImage: nil,
			}
			var expPhotoImage []models.CoverPhotosObj
			errObject := json.Unmarshal([]byte(element.ExpPhotoImage), &expPhotoImage)
			if errObject != nil {
				//fmt.Println("Error : ",err.Error())
				return nil, models.ErrInternalServerError
			}
			expPhoto.ExpPhotoImage = expPhotoImage
			expPhotos = append(expPhotos, expPhoto)
		}
	}
	if err != nil {
		return nil, err
	}
	var expAddOns []models.ExperienceAddOnObj
	expAddOnsQuery, errorQuery := m.adOnsRepo.GetByExpId(ctx, res.Id)
	if expAddOnsQuery != nil {
		for _, element := range expAddOnsQuery {
			var currency string
			if element.Currency == 1 {
				currency = "USD"
			} else {
				currency = "IDR"
			}
			if currencyPrice == "USD" {
				if currency == "IDR" {
					convertCurrency, _ := m.currencyUsecase.ExchangeRatesApi(ctx, "IDR", "USD")
					calculatePrice := convertCurrency.Rates.USD * element.Amount
					element.Amount = calculatePrice
					currency = "USD"
				}
			} else if currencyPrice == "IDR" {
				if currency == "USD" {
					convertCurrency, _ := m.currencyUsecase.ExchangeRatesApi(ctx, "USD", "IDR")
					calculatePrice := convertCurrency.Rates.IDR * element.Amount
					element.Amount = calculatePrice
					currency = "IDR"
				}
			}
			addOns := models.ExperienceAddOnObj{
				Id:       element.Id,
				Name:     element.Name,
				Desc:     element.Desc,
				Currency: currency,
				Amount:   element.Amount,
			}
			expAddOns = append(expAddOns, addOns)
		}
	}
	var expPayment []models.ExpPaymentObj
	expPaymentQuery, errorQuery := m.paymentRepo.GetByExpID(ctx, res.Id)
	for _, elementPayment := range expPaymentQuery {
		var currency string
		if elementPayment.Currency == 1 {
			currency = "USD"
		} else {
			currency = "IDR"
		}

		var priceItemType string
		if elementPayment.PriceItemType == 1 {
			priceItemType = "Per Pax"
		} else {
			priceItemType = "Per Trip"
		}
		customPrice := make([]models.CustomPrice, 0)
		if elementPayment.CustomPrice != nil {
			if *elementPayment.CustomPrice != "" {
				errObject := json.Unmarshal([]byte(*elementPayment.CustomPrice), &customPrice)
				if errObject != nil {
					return nil, models.ErrInternalServerError
				}
			}
		}

		if currencyPrice == "USD" {
			if currency == "IDR" {
				convertCurrency, _ := m.currencyUsecase.ExchangeRatesApi(ctx, "IDR", "USD")
				calculatePrice := convertCurrency.Rates.USD * elementPayment.Price
				elementPayment.Price = calculatePrice
				currency = "USD"
			}
		} else if currencyPrice == "IDR" {
			if currency == "USD" {
				convertCurrency, _ := m.currencyUsecase.ExchangeRatesApi(ctx, "USD", "IDR")
				calculatePrice := convertCurrency.Rates.IDR * elementPayment.Price
				elementPayment.Price = calculatePrice
				currency = "IDR"
			}
		}

		for index, elementCustomPrice := range customPrice {
			if currencyPrice == "USD" {
				if elementCustomPrice.Currency == "IDR" {
					convertCurrency, _ := m.currencyUsecase.ExchangeRatesApi(ctx, "IDR", "USD")
					calculatePrice := convertCurrency.Rates.USD * elementCustomPrice.Price
					customPrice[index].Price = calculatePrice
					customPrice[index].Currency = "USD"
				}
			} else if currencyPrice == "IDR" {
				if elementCustomPrice.Currency == "USD" {
					convertCurrency, _ := m.currencyUsecase.ExchangeRatesApi(ctx, "USD", "IDR")
					calculatePrice := convertCurrency.Rates.IDR * elementCustomPrice.Price
					customPrice[index].Price = calculatePrice
					customPrice[index].Currency = "IDR"
				}
			}
		}
		expPayobj := models.ExpPaymentObj{
			Id:              elementPayment.Id,
			Currency:        currency,
			Price:           elementPayment.Price,
			PriceItemType:   priceItemType,
			PaymentTypeId:   elementPayment.ExpPaymentTypeId,
			PaymentTypeName: elementPayment.ExpPaymentTypeName,
			PaymentTypeDesc: elementPayment.ExpPaymentTypeDesc,
			CustomPrice:     customPrice,
		}
		expPayment = append(expPayment, expPayobj)
	}

	var expAvailability []models.ExpAvailablitityObj
	expAvailabilityQuery, errorQuery := m.exp_availablitiy.GetByExpId(ctx, res.Id)
	if errorQuery != nil {
		return nil, errorQuery
	}
	if expAvailabilityQuery != nil {
		for _, element := range expAvailabilityQuery {
			expA := models.ExpAvailablitityObj{
				Id:    element.Id,
				Year:  element.ExpAvailabilityYear,
				Month: element.ExpAvailabilityMonth,
				Date:  nil,
			}
			var dates []string
			errObject := json.Unmarshal([]byte(element.ExpAvailabilityDate), &dates)
			if errObject != nil {
				return nil, models.ErrInternalServerError
			}
			if isMerchant == "true"{
				expA.Date = dates
			}else {
				if res.ExpTripType == "Private Trip" {
					for _, date := range dates {
						checkBookingCount, err := m.bookingRepo.GetCountByBookingDateExp(ctx, date, element.ExpId)
						if err != nil {
							return nil, err
						}
						if res.ExpMaximumBookingType != nil && res.ExpMaximumBookingAmount != nil{
							if *res.ExpMaximumBookingType == "Days"{
								convertDate ,_ := time.Parse("2006-01-02",date)
								now := time.Now().AddDate(0,0,*res.ExpMaximumBookingAmount)
								if (convertDate.After(now) || convertDate.Equal(now)) && checkBookingCount == 0{
									expA.Date = append(expA.Date, date)
								}

							}else if *res.ExpMaximumBookingType == "Week"{
								convertDate ,_ := time.Parse("2006-01-02",date)
								now := time.Now().AddDate(0,0,*res.ExpMaximumBookingAmount * 7)
								if (convertDate.After(now) || convertDate.Equal(now)) && checkBookingCount == 0{
									expA.Date = append(expA.Date, date)
								}
							}else if *res.ExpMaximumBookingType == "Month"{
								convertDate ,_ := time.Parse("2006-01-02",date)
								now := time.Now().AddDate(0,*res.ExpMaximumBookingAmount,0 )
								if (convertDate.After(now) || convertDate.Equal(now)) && checkBookingCount == 0 {
									expA.Date = append(expA.Date, date)
								}
							}
						}else {
							if checkBookingCount == 0 {
								expA.Date = append(expA.Date, date)
							}
						}
					}
				} else {
					for _, date := range dates {
						checkBookingCount, err := m.bookingRepo.GetCountByBookingDateExp(ctx, date, element.ExpId)
						if err != nil {
							return nil, err
						}
						if res.ExpMaximumBookingType != nil && res.ExpMaximumBookingAmount != nil{
							if *res.ExpMaximumBookingType == "Days"{
								convertDate ,_ := time.Parse("2006-01-02",date)
								now := time.Now().AddDate(0,0,*res.ExpMaximumBookingAmount)
								if (convertDate.After(now) || convertDate.Equal(now)) && checkBookingCount < res.ExpMaxGuest{
									expA.Date = append(expA.Date, date)
								}

							}else if *res.ExpMaximumBookingType == "Week"{
								convertDate ,_ := time.Parse("2006-01-02",date)
								now := time.Now().AddDate(0,0,*res.ExpMaximumBookingAmount * 7)
								if (convertDate.After(now) || convertDate.Equal(now)) && checkBookingCount < res.ExpMaxGuest{
									expA.Date = append(expA.Date, date)
								}
							}else if *res.ExpMaximumBookingType == "Month"{
								convertDate ,_ := time.Parse("2006-01-02",date)
								now := time.Now().AddDate(0,*res.ExpMaximumBookingAmount,0 )
								if (convertDate.After(now) || convertDate.Equal(now)) && checkBookingCount < res.ExpMaxGuest{
									expA.Date = append(expA.Date, date)
								}
							}
						}else {
							if checkBookingCount < res.ExpMaxGuest {
								expA.Date = append(expA.Date, date)
							}
						}
					}
				}
			}
			expAvailability = append(expAvailability, expA)
		}
	}


	expItinerary := models.ExpItineraryObject{}
	errObject := json.Unmarshal([]byte(res.ExpInternary), &expItinerary)
	if errObject != nil {
		//fmt.Println("Error : ",err.Error())
		return nil, models.ErrInternalServerError
	}
	expFacilities := make([]models.ExpFacilitiesObject, 0)
	//errObject = json.Unmarshal([]byte(res.ExpFacilities), &expFacilities)
	//if errObject != nil {
	//	//fmt.Println("Error : ",err.Error())
	//	return nil, models.ErrInternalServerError
	//}
	getFacilities, err := m.expFacilitiesRepo.GetJoin(ctx, res.Id, "")
	for _, element := range getFacilities {
		facility := models.ExpFacilitiesObject{
			Name:   element.FacilityName,
			Icon:   *element.FacilityIcon,
			Amount: element.Amount,
		}
		expFacilities = append(expFacilities, facility)
	}
	expInclusion := make([]models.ExpInclusionObject, 0)
	//if res.ExpInclusion != ""{
	//	errObject = json.Unmarshal([]byte(res.ExpInclusion), &expInclusion)
	//	if errObject != nil {
	//		//fmt.Println("Error : ",err.Error())
	//		return nil, models.ErrInternalServerError
	//	}
	//}
	getInclude, err := m.expIncludeRepo.GetByExpIdJoin(ctx, res.Id)
	for _, element := range getInclude {
		inclusion := models.ExpInclusionObject{
			Name: element.IncludeName,
			Type: 0,
		}
		expInclusion = append(expInclusion, inclusion)
	}
	getExclude, err := m.expExcludeRepo.GetByExpIdJoin(ctx, res.Id)
	for _, element := range getExclude {
		inclusion := models.ExpInclusionObject{
			Name: element.ExcludeName,
			Type: 1,
		}
		expInclusion = append(expInclusion, inclusion)
	}
	expRules := make([]models.ExpRulesObject, 0)
	if res.ExpRules != "" {
		errObject = json.Unmarshal([]byte(res.ExpRules), &expRules)
		if errObject != nil {
			//fmt.Println("Error : ",err.Error())
			return nil, models.ErrInternalServerError
		}
	}

	harbors, err := m.harborsRepo.GetByID(ctx, res.HarborsId)
	city, err := m.cpcRepo.GetCityByID(ctx, harbors.CityId)
	province, err := m.cpcRepo.GetProvinceByID(ctx, city.ProvinceId)
	minimumBooking := models.MinimumBookingObj{
		MinimumBookingDesc:   res.MinimumBookingDesc,
		MinimumBookingAmount: res.MinimumBookingAmount,
	}


	var expType []string
	expTypes,_ := m.filterATRepo.GetJoinExpType(ctx,res.Id)
	for _,elementType := range expTypes{
		expType = append(expType,elementType.ExpTypeName)
	}
	countRating, err := m.reviewsRepo.CountRating(ctx, 0, res.Id)
	experiences := models.ExperienceDto{
		Id:                       res.Id,
		ExpTitle:                 res.ExpTitle,
		ExpType:                  expType,
		ExpTripType:              res.ExpTripType,
		ExpBookingType:           res.ExpBookingType,
		ExpDesc:                  res.ExpDesc,
		ExpMaxGuest:              res.ExpMaxGuest,
		ExpPickupPlace:           res.ExpPickupPlace,
		ExpPickupTime:            res.ExpPickupTime,
		ExpPickupPlaceLongitude:  res.ExpPickupPlaceLongitude,
		ExpPickupPlaceLatitude:   res.ExpPickupPlaceLatitude,
		ExpPickupPlaceMapsName:   res.ExpPickupPlaceMapsName,
		ExpInternary:             expItinerary,
		ExpFacilities:            expFacilities,
		ExpInclusion:             expInclusion,
		ExpRules:                 expRules,
		ExpAvailability:          expAvailability,
		ExpPayment:               expPayment,
		ExpPhotos:                expPhotos,
		ExperienceAddOn:          expAddOns,
		Status:                   res.Status,
		Rating:                   res.Rating,
		CountRating:              countRating,
		ExpLocationLatitude:      res.ExpLocationLatitude,
		ExpLocationLongitude:     res.ExpLocationLongitude,
		ExpLocationName:          res.ExpLocationName,
		ExpCoverPhoto:            res.ExpCoverPhoto,
		ExpDuration:              res.ExpDuration,
		MinimumBooking:           minimumBooking,
		MerchantId:               res.MerchantId,
		HarborsName:              harbors.HarborsName,
		City:                     city.CityName,
		Province:                 province.ProvinceName,
		GuideReview:              res.GuideReview,
		ActivitiesReview:         res.ActivitiesReview,
		ServiceReview:            res.ServiceReview,
		CleanlinessReview:        res.CleanlinessReview,
		ValueReview:              res.ValueReview,
		ExpPaymentDeadlineAmount: res.ExpPaymentDeadlineAmount,
		ExpPaymentDeadlineType:   res.ExpPaymentDeadlineType,
		IsCustomisedByUser:       res.IsCustomisedByUser,
		ExpLocationMapName:       res.ExpLocationMapName,
		ExpLatitudeMap:           res.ExpLatitudeMap,
		ExpLongitudeMap:          res.ExpLongitudeMap,
		ExpMaximumBookingAmount:res.ExpMaximumBookingAmount,
		ExpMaximumBookingType:res.ExpMaximumBookingType,
	}
	return &experiences, nil
}

/*
* In this function below, I'm using errgroup with the pipeline pattern
* Look how this works in this package explanation
* in godoc: https://godoc.org/golang.org/x/sync/errgroup#ex-Group--Pipeline
 */
