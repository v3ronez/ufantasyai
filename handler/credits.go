package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/v3ronez/ufantasyai/db"
	"github.com/v3ronez/ufantasyai/view/credits"
)

func HandleCreditsIndex(w http.ResponseWriter, r *http.Request) error {
	return RenderComponent(w, r, credits.Index())
}

func HandleStripeCheckout(w http.ResponseWriter, r *http.Request) error {
	productID := chi.URLParam(r, "productID")
	stripe.Key = os.Getenv("STRIPE_KEY")
	checkoutParams := &stripe.CheckoutSessionParams{
		SuccessURL: stripe.String("http://localhost:3000/checkout/success/{CHECKOUT_SESSION_ID}"),
		CancelURL:  stripe.String("http://localhost:3000/checkout/cancel"),
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(productID),
				Quantity: stripe.Int64(1),
			},
		},
	}
	sessin, err := session.New(checkoutParams)
	if err != nil {
		return err
	}

	return HtmxRedirect(w, r, sessin.URL)
}

func HandleStripeCheckoutSuccess(w http.ResponseWriter, r *http.Request) error {
	//session not contais priceID
	sessionID := chi.URLParam(r, "sessionID")
	user := GetAuthenticatedUser(r)
	stripe.Key = os.Getenv("STRIPE_KEY")
	sess, err := session.Get(sessionID, nil)
	if err != nil {
		return err
	}
	lineItemParams := stripe.CheckoutSessionListLineItemsParams{}
	lineItemParams.Session = stripe.String(sess.ID)
	iter := session.ListLineItems(&lineItemParams)
	iter.Next()
	item := iter.LineItem()
	priceID := item.Price.ID

	switch priceID {
	case os.Getenv("100_CREDITS"):
		user.Account.Credits += 100
	case os.Getenv("500_CREDITS"):
		user.Account.Credits += 500
	case os.Getenv("1000_CREDITS"):
		user.Account.Credits += 1000
	default:
		return fmt.Errorf("invalid price id")

	}
	if err := db.UpdateAccount(&user.Account); err != nil {
		return err
	}

	http.Redirect(w, r, "/generate", http.StatusSeeOther)
	return nil
}

func HandleStripeCheckoutCancel(w http.ResponseWriter, r *http.Request) error {
	return nil
}
