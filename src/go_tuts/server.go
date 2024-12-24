package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

// http://localhost:4242/create-payment-intent

func main() {
	stripe.Key = "sk_test_51QWJOm15vWxwhRMGBD8pTKPQqKknpkzJU3SuGPpqfaJ1Nnx1wmmC7YwdW3oDM6ksUVSbrlu6Q98CokO1N6DsXfAj00sdrcooDc"

	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)

	http.HandleFunc("/health", handleHealth)

	log.Println("Listening on localhost:4242...")

	var err error = http.ListenAndServe("localhost:4242", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("END POINT CALLED !!!")
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)

		return
	}

	// fmt.Println("Request method was correct!")

	var req struct {
		ProductId string `json:"product_id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address1  string `json:"address1"`
		Address2  string `json:"address2"`
		City      string `json:"city"`
		State     string `json:"state"`
		Zip       string `json:"zip"`
		Country   string `json:"country"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	// code for creating payment intent
	// params for stripe to use to create payment intent
	// params to be sent to stripe API ... so that at we get the Payment Intent object
	params := &stripe.PaymentIntentParams{ // from stripe library
		Amount:   stripe.Int64(calculateOrderAmount(req.ProductId)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		// AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
		// 	Enabled: stripe.Bool(true),
		// }
		PaymentMethodOptions: &stripe.PaymentIntentPaymentMethodOptionsParams{
			Card: &stripe.PaymentIntentPaymentMethodOptionsCardParams{
				RequestThreeDSecure: stripe.String("automatic"),
			},
		},
	}
	paymentIntent, err := paymentintent.New(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// printing of payment intent for testing purposes
	// fmt.Println(paymentIntent.ClientSecret)

	// response from client
	var response struct {
		ClientSecret string `json:"clientSecret"`
	}

	response.ClientSecret = paymentIntent.ClientSecret

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//sending (json-object) response to caller
	w.Header().Set("Content-Type", "application/json")

	_, err = io.Copy(w, &buf) // caller of API will read from here
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println("\n")
	// to be printed on console/ backend
	fmt.Println(req.ProductId)
	fmt.Println(req.FirstName)
	fmt.Println(req.LastName)
	fmt.Println(req.Address1)
	fmt.Println(req.Address2)
	fmt.Println(req.City)
	fmt.Println(req.State)
	fmt.Println(req.Zip)
	fmt.Println(req.Country)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("HEALTH MESSAGES !!!")

	var response []byte = []byte("Server is up and running!")

	_, err := w.Write(response)
	if err != nil {
		fmt.Println(err)
	}
}

func calculateOrderAmount(productId string) int64 {
	switch productId {
	case "Forever Pants":
		return 260.00
	case "Forever Shirt":
		return 155.00
	case "Forever Shorts":
		return 300.00
	}

	return 0
}
