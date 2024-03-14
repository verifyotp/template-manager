package payment_processors

import (
	"context"
	"errors"

	"github.com/stripe/stripe-go/v76"
	billingPortalSession "github.com/stripe/stripe-go/v76/billingportal/session"
	checkoutSession "github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/stripe/stripe-go/v76/subscription"
)

const (
	StripeSubscriptionStatusActive   = "active"
	StripeSubscriptionStatusTrialing = "trialing"
	StripeSubscriptionStatusPastDue  = "past_due"
)

type (
	StripeService interface {
		NewCheckoutSession(
			ctx context.Context,
			input StripeNewCheckoutSessionInput,
		) (*stripe.CheckoutSession, error)

		GetCheckoutSession(
			ctx context.Context,
			sessionId string,
		) (*stripe.CheckoutSession, error)

		GetSubscription(
			ctx context.Context,
			subscriptionId string,
		) (*stripe.Subscription, error)

		GenerateSubscriptionManagementLink(
			ctx context.Context,
			customerId,
			returnUrl string,
		) (*stripe.BillingPortalSession, error)

		CancelSubscription(
			ctx context.Context,
			subscriptionId string,
		) error
	}

	StripeCheckoutSessionApi interface {
		New(params *stripe.CheckoutSessionParams) (*stripe.CheckoutSession, error)
		Get(sessionId string, params *stripe.CheckoutSessionParams) (*stripe.CheckoutSession, error)
	}

	StripeSubscriptionApi interface {
		Get(subscriptionId string, params *stripe.SubscriptionParams) (*stripe.Subscription, error)
		Update(subscriptionId string, params *stripe.SubscriptionParams) (*stripe.Subscription, error)
	}

	StripeBillingPortalSessionApi interface {
		New(params *stripe.BillingPortalSessionParams) (*stripe.BillingPortalSession, error)
	}

	StripeClient struct {
		secretKey                     string
		checkoutSessionApiClient      StripeCheckoutSessionApi
		subscriptionApiClient         StripeSubscriptionApi
		billingPortalSessionApiClient StripeBillingPortalSessionApi
	}

	StripeOptions func(s *StripeClient) error

	StripeNewCheckoutSessionInput struct {
		SuccessURL *string
		CancelURL  *string
		PriceId    *string
		UIMode     string
	}

	StripeNewCheckoutSessionOutput struct {
		URL string
	}
)

func NewStripeClient(secretKey string, opts ...StripeOptions) (*StripeClient, error) {
	if secretKey == "" {
		return nil, errors.New("secret Key is required")
	}

	s := &StripeClient{secretKey: secretKey}

	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, err
		}
	}

	if s.checkoutSessionApiClient == nil {
		s.checkoutSessionApiClient = &checkoutSession.Client{B: stripe.GetBackend(stripe.APIBackend), Key: s.secretKey}
	}

	if s.subscriptionApiClient == nil {
		s.subscriptionApiClient = &subscription.Client{B: stripe.GetBackend(stripe.APIBackend), Key: s.secretKey}
	}

	if s.billingPortalSessionApiClient == nil {
		s.billingPortalSessionApiClient = &billingPortalSession.Client{B: stripe.GetBackend(stripe.APIBackend), Key: s.secretKey}
	}

	return s, nil
}

func (s *StripeClient) NewCheckoutSession(
	ctx context.Context,
	input StripeNewCheckoutSessionInput,
) (*stripe.CheckoutSession, error) {

	uiMode := stripe.CheckoutSessionUIModeHosted
	var redirectOnCompletion *string

	if input.UIMode == "embedded" {
		uiMode = stripe.CheckoutSessionUIModeEmbedded
		input.SuccessURL = nil
		input.CancelURL = nil
		redirectOnCompletion = stripe.String("never")
	}

	if uiMode == stripe.CheckoutSessionUIModeHosted {
		if input.SuccessURL == nil {
			return nil, errors.New("SuccessURL is required")
		}
		if input.CancelURL == nil {
			return nil, errors.New("CancelURL is required")
		}
	}

	params := &stripe.CheckoutSessionParams{
		SuccessURL:           input.SuccessURL,
		CancelURL:            input.CancelURL,
		Mode:                 stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		UIMode:               stripe.String(string(uiMode)),
		RedirectOnCompletion: redirectOnCompletion,
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				Price: input.PriceId,
				// For metered billing, do not pass quantity
				Quantity: stripe.Int64(1),
			},
		},
	}
	return s.checkoutSessionApiClient.New(params)
}

func (s *StripeClient) GetCheckoutSession(
	ctx context.Context,
	sessionId string,
) (*stripe.CheckoutSession, error) {
	return s.checkoutSessionApiClient.Get(sessionId, nil)
}

func (s *StripeClient) GetSubscription(
	ctx context.Context,
	subscriptionId string,
) (*stripe.Subscription, error) {
	return s.subscriptionApiClient.Get(subscriptionId, nil)
}

func (s *StripeClient) GenerateSubscriptionManagementLink(
	ctx context.Context,
	customerId,
	returnUrl string,
) (*stripe.BillingPortalSession, error) {
	params := &stripe.BillingPortalSessionParams{
		Customer:  stripe.String(customerId),
		ReturnURL: stripe.String(returnUrl),
	}

	return s.billingPortalSessionApiClient.New(params)
}

// CancelSubscription cancels a subscription at the end of the current billing period.
func (s *StripeClient) CancelSubscription(
	ctx context.Context,
	subscriptionId string,
) error {
	params := &stripe.SubscriptionParams{CancelAtPeriodEnd: stripe.Bool(true)}
	_, err := s.subscriptionApiClient.Update(subscriptionId, params)
	return err
}
