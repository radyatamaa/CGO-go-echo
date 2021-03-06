package xendit

import (
	"context"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/card"
	"github.com/xendit/xendit-go/client"
	"github.com/xendit/xendit-go/virtualaccount"
)

var XenClient *client.API

const (
	// dev
	defXenditSecretKey = "xnd_development_Dms6iAkgd6b4p5f9jpLdP41uaCVBdCLPNqJ00XDiFQL0oIpsTZYVLlERGFnxi"

	// Production
	// defXenditSecretKey = "xnd_production_wUqt0xBrasJpktiTTgOgOIojpewhY455AGFik0AxizdVAL1pIUYBic8EGeStyDs"
)

func XenditSetup() {
	XenClient = client.New(defXenditSecretKey)
}

type VirtualAccount struct {
	*virtualaccount.Client
	ExternalID string
	BankCode   string
	Name       string
	ExpireDate *time.Time
}

type VACallbackRequest struct {
	Amount                   float64 `json:"amount"`
	CallbackVirtualAccountID string  `json:"callback_virtual_account_id"`
	PaymentID                string  `json:"payment_id"`
	ExternalID               string  `json:"external_id"`
	AccountNumber            string  `json:"account_number"`
	MerchantCode             string  `json:"merchant_code"`
	BankCode                 string  `json:"bank_code"`
	TransactionTimestamp     string  `json:"transaction_timestamp"`
	Currency                 string  `json:"currency"`
	Created                  string  `json:"created"`
	Updated                  string  `json:"updated"`
	ID                       string  `json:"id"`
	OwnerID                  string  `json:"owner_id"`
}

func (va *VirtualAccount) CreateFixedVA(ctx context.Context) (*xendit.VirtualAccount, error) {
	data := virtualaccount.CreateFixedVAParams{
		ExternalID:     va.ExternalID,
		BankCode:       va.BankCode,
		Name:           va.Name,
		ExpirationDate: va.ExpireDate,
	}
	resVa, err := va.CreateFixedVAWithContext(ctx, &data)
	if err != nil {
		return nil, err
	}

	return resVa, nil
}

type CreditCard struct {
	*card.Client
	TokenID    string
	AuthID     string
	ExternalID string
	Amount     float64
	IsCapture  bool
}

func (cc *CreditCard) CreateCharge(ctx context.Context) (*xendit.CardCharge, error) {
	data := &card.CreateChargeParams{
		TokenID:          cc.TokenID,
		AuthenticationID: cc.AuthID,
		ExternalID:       cc.ExternalID,
		Amount:           cc.Amount,
		Capture:          &cc.IsCapture,
	}

	resCc, err := cc.CreateChargeWithContext(ctx, data)
	if err != nil {
		return nil, err
	}

	return resCc, nil
}
