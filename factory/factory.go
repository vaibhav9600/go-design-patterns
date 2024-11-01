package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type PaymentGateWayType int

const (
	PaytmPaymentGateWay PaymentGateWayType = iota
	PhonePePaymentGateway
)

type PaymentGateway interface {
	ProcessPayments(amount float64) error
}

type PaytmPayment struct {
	ClientId     string
	ClientSecret string
}

type PhonePePayment struct {
	ApiKey string
}

func (paytm *PaytmPayment) ProcessPayments(amount float64) error {
	fmt.Println("Paytm payment selected processing amount = ", amount)

	return nil
}

func (phonePe *PhonePePayment) ProcessPayments(amount float64) error {
	fmt.Println("PhonePe payment selected processing amount = ", amount)

	return nil
}

func getPaymentProvider(gtwType PaymentGateWayType) (PaymentGateway, error) {
	switch gtwType {
	case PaytmPaymentGateWay:
		return &PaytmPayment{}, nil
	case PhonePePaymentGateway:
		return &PhonePePayment{}, nil
	default:
		return nil, errors.New("unsupported payment gateway type")
	}
}

// here since interface is present we wont send it via *, since all the interface already points to pointer struct see above
// we are returning address of structs
type Option func(PaymentGateway) error

func withClientSecret(secret string) Option {
	return func(pg PaymentGateway) error {
		if paytm, ok := pg.(*PaytmPayment); ok {
			paytm.ClientSecret = secret
			return nil
		}
		return errors.ErrUnsupported
	}
}

func withClientId(clientId string) Option {
	return func(pg PaymentGateway) error {
		if paytm, ok := pg.(*PaytmPayment); ok {
			paytm.ClientId = clientId
			return nil
		}
		return errors.ErrUnsupported
	}
}

func withApiKey(apiKey string) Option {
	return func(pg PaymentGateway) error {
		if phonePe, ok := pg.(*PhonePePayment); ok {
			phonePe.ApiKey = apiKey
			return nil
		}
		return errors.ErrUnsupported
	}
}

func NewPaymentGateway(gwType PaymentGateWayType, opts ...Option) (PaymentGateway, error) {
	var pg PaymentGateway
	switch gwType {
	case PaytmPaymentGateWay:
		pg = &PaytmPayment{}
	case PhonePePaymentGateway:
		pg = &PhonePePayment{}
	default:
		return nil, errors.New("unsupported payment gateway type")
	}

	for _, opt := range opts {
		if err := opt(pg); err != nil {
			return nil, err
		}
	}

	return pg, nil
}

func main() {
	var gtwType PaymentGateWayType
	fmt.Scanf("%d", &gtwType)

	gtw, err := getPaymentProvider(gtwType)
	if err != nil {
		fmt.Println(err)
	} else {
		gtw.ProcessPayments(float64(rand.Int()))
	}

	paytmGateway, err := NewPaymentGateway(
		PaytmPaymentGateWay,
		withClientId("paypal-client-id"),
		withClientSecret("paypal-client-secret"),
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	paytmGateway.ProcessPayments(100.00)

	phonePeGtw, err := NewPaymentGateway(
		PhonePePaymentGateway,
		withApiKey("stripe-api-key"),
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	phonePeGtw.ProcessPayments(150.50)
}
