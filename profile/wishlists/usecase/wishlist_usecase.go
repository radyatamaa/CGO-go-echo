package usecase

import (
	"context"
	"encoding/json"
	"time"

	"github.com/auth/user"
	"github.com/models"
	"github.com/product/reviews"
	"github.com/profile/wishlists"
	"github.com/service/exp_payment"
	"github.com/service/experience"
)

type wishListUsecase struct {
	wlRepo      wishlists.Repository
	userUsecase user.Usecase
	expRepo     experience.Repository
	paymentRepo exp_payment.Repository
	reviewRepo  reviews.Repository
	ctxTimeout  time.Duration
}

func NewWishlistUsecase(
	w wishlists.Repository,
	u user.Usecase,
	e experience.Repository,
	p exp_payment.Repository,
	r reviews.Repository,
	timeout time.Duration,
) wishlists.Usecase {
	return &wishListUsecase{
		wlRepo:      w,
		userUsecase: u,
		expRepo:     e,
		paymentRepo: p,
		reviewRepo:  r,
		ctxTimeout:  timeout,
	}
}

func (w wishListUsecase) List(ctx context.Context, token string) ([]*models.WishlistOut, error) {
	ctx, cancel := context.WithTimeout(ctx, w.ctxTimeout)
	defer cancel()

	currentUser, err := w.userUsecase.ValidateTokenUser(ctx, token)
	if err != nil {
		return nil, err
	}

	wLists, err := w.wlRepo.List(ctx, currentUser.Id)
	if err != nil {
		return nil, err
	}

	results := make([]*models.WishlistOut, len(wLists))
	for i, wl := range wLists {
		exp, err := w.expRepo.GetByID(ctx, wl.ExpId.String)
		if err != nil {
			return nil, err
		}

		var expType []string
		if errUnmarshal := json.Unmarshal([]byte(exp.ExpType), &expType); errUnmarshal != nil {
			return nil, models.ErrInternalServerError
		}

		expPayment, err := w.paymentRepo.GetByExpID(ctx, exp.Id)
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

		countRating, err := w.reviewRepo.CountRating(ctx, exp.Id)
		if err != nil {
			return nil, err
		}

		wtype := "EXPERIENCE"
		if wl.TransId.String != "" {
			wtype = "TRANSPORTATION"
		}

		results[i] = &models.WishlistOut{
			WishlistID:  wl.Id,
			Type:        wtype,
			ExpID:       exp.Id,
			ExpTitle:    exp.ExpTitle,
			ExpType:     expType,
			Rating:      exp.Rating,
			CountRating: countRating,
			Currency:    currency,
			Price:       expPayment[0].Price,
			PaymentType: priceItemType,
			CoverPhoto:  *exp.ExpCoverPhoto,
		}
	}

	return results, nil
}

func (w wishListUsecase) Insert(ctx context.Context, wl *models.WishlistIn, token string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, w.ctxTimeout)
	defer cancel()


	currentUser, err := w.userUsecase.ValidateTokenUser(ctx, token)
	if err != nil {
		return "", err
	}
	checkWhislist, err := w.wlRepo.GetByUserAndExpId(ctx,currentUser.Id,wl.ExpID)
	if err != nil {
		return "", err
	}
	if len(checkWhislist) != 0 || checkWhislist != nil{
		return "",models.ErrConflict
	}
	newData := &models.Wishlist{
		Id:           "",
		CreatedBy:    currentUser.UserEmail,
		CreatedDate:  time.Now(),
		ModifiedBy:   nil,
		ModifiedDate: nil,
		DeletedBy:    nil,
		DeletedDate:  nil,
		IsDeleted:    0,
		IsActive:     1,
		UserId:       currentUser.Id,
		ExpId:        wl.ExpID,
		TransId:      wl.TransID,
	}

	res, err := w.wlRepo.Insert(ctx, newData)
	if err != nil {
		return "", err
	}

	return res.Id, nil
}