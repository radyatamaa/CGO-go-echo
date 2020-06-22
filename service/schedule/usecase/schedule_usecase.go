package usecase

import (
	"encoding/json"
	"time"

	"github.com/auth/merchant"
	"github.com/models"
	"github.com/service/exp_availability"
	"github.com/service/experience"
	"github.com/service/schedule"
	"github.com/service/time_options"
	"github.com/service/transportation"
	"golang.org/x/net/context"
)

type scheduleUsecase struct {
	expAvailability    exp_availability.Repository
	experieceRepo      experience.Repository
	transportationRepo transportation.Repository
	merchantUsecase    merchant.Usecase
	scheduleRepo       schedule.Repository
	timeOptionsRepo    time_options.Repository
	contextTimeout     time.Duration
}

func NewScheduleUsecase(tr transportation.Repository, mr merchant.Usecase, s schedule.Repository, tmo time_options.Repository, exp experience.Repository, expA exp_availability.Repository, timeout time.Duration) schedule.Usecase {
	return &scheduleUsecase{
		transportationRepo: tr,
		merchantUsecase:    mr,
		scheduleRepo:       s,
		timeOptionsRepo:    tmo,
		experieceRepo:      exp,
		expAvailability:    expA,
		contextTimeout:     timeout,
	}
}

func (s scheduleUsecase) InsertSchedule(c context.Context, command *models.NewCommandSchedule) (*models.NewCommandSchedule, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	defer cancel()
	//for _, month := range year.Month {
	for _, day := range command.DayPrice {
		for _, times := range command.TimeObj {
			var currency int
			if day.Currency == "USD" {
				currency = 1
			} else {
				currency = 0
			}
			priceObj := models.PriceObj{
				AdultPrice:    day.AdultPrice,
				ChildrenPrice: day.ChildrenPrice,
				Currency:      currency,
			}
			departureTimeOption, err := s.timeOptionsRepo.GetByTime(ctx, times.DepartureTime)
			if err != nil {
				//return nil, err
			}
			arrivalTimeOption, err := s.timeOptionsRepo.GetByTime(ctx, times.ArrivalTime)
			if err != nil {
				//return nil, err
			}
			price, _ := json.Marshal(priceObj)
			schedule := models.Schedule{
				Id:                    "",
				CreatedBy:             command.CreatedBy,
				CreatedDate:           time.Time{},
				ModifiedBy:            nil,
				ModifiedDate:          nil,
				DeletedBy:             nil,
				DeletedDate:           nil,
				IsDeleted:             0,
				IsActive:              0,
				TransId:               command.TransId,
				DepartureTime:         times.DepartureTime,
				ArrivalTime:           times.ArrivalTime,
				Day:                   day.Day,
				Month:                 command.Month,
				Year:                  command.Year,
				DepartureDate:         day.DepartureDate,
				Price:                 string(price),
				DepartureTimeoptionId: &departureTimeOption.Id,
				ArrivalTimeoptionId:   &arrivalTimeOption.Id,
			}
			_, err = s.scheduleRepo.Insert(ctx, schedule)
			if err != nil {
				//return nil, err
			}
		}
	}
	return command, nil
}
func (s scheduleUsecase) GetScheduleByMerchantId(c context.Context, merchantId string, date string) (*models.ScheduleDtoObj, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	dateParse := date + "-" + "01"
	layoutFormat := "2006-01-02"
	dt, _ := time.Parse(layoutFormat, dateParse)
	//year := dt.Year()
	month := dt.Month().String()

	var dates []string

	start, errDateDob := time.Parse(layoutFormat, dateParse)
	if errDateDob != nil {
		return nil, errDateDob
	}
	dates = append(dates, start.Format("2006-01-02"))

datess:
	start = start.AddDate(0, 0, 1)
	if start.Month().String() != month {

	} else {
		dates = append(dates, start.Format("2006-01-02"))
		goto datess
	}
	var result models.ScheduleDtoObj
	result.MerchantId = merchantId

	for _, element := range dates {
		obj := models.ScheduleObjDate{
			Date:                element,
			TransportationCount: 0,
			ExperienceCount:     0,
		}

		transportationCount, err := s.scheduleRepo.GetCountSchedule(ctx, merchantId, element)
		if err != nil {
			return nil, err
		}
		obj.TransportationCount = transportationCount

		experienceCount, err := s.expAvailability.GetCountDate(ctx, element, merchantId)
		if err != nil {
			return nil, err
		}

		obj.ExperienceCount = experienceCount

		if obj.TransportationCount != 0 || obj.ExperienceCount != 0 {
			result.ScheduleDate = append(result.ScheduleDate, obj)
		}

	}
	return &result, nil

}
