# Stripe Payment Intent API

A Go-based server application that handles Stripe payment intents and provides a basic health check endpoint. This application demonstrates how to integrate the Stripe API for processing payments in a server-side environment.

## Features

- **Create Payment Intent**: Generate a Stripe Payment Intent for processing payments.
- **Health Check Endpoint**: Verify server availability with a simple `/health` endpoint.
- **JSON Input/Output**: Accepts and returns data in JSON format.
- **Address Handling**: Includes address and user details in the payment request.

## Prerequisites

Before running the application, ensure you have the following:

- **Go** (version 1.16 or later)
- **Stripe API Key**: Obtain a test key from the [Stripe Dashboard](https://dashboard.stripe.com/apikeys).

## Setup Instructions

1. Clone the repository:
   ```bash
   git clone https://github.com/eugenius-watchman/golang_tuts.git
   cd golang_tuts
2. Install dependencies:

   go mod tidy

3. Update the Stripe API key: Replace the placeholder API key in the code (sk_test_...) with your Stripe secret key.

4. Run the application:

   go run main.go

5. Test the endpoints using a tool like Postman or curl.


Endpoints
1. Create Payment Intent

    URL: /create-payment-intent
    Method: POST
    Request Body:

{
  "product_id": "Forever Pants",
  "first_name": "John",
  "last_name": "Doe",
  "address1": "123 Main St",
  "address2": "Apt 4B",
  "city": "Springfield",
  "state": "IL",
  "zip": "62701",
  "country": "US"
}

Response:

{
  "clientSecret": "pi_12345_secret_67890___"
}

2. Health Check

    URL: /health
    Method: GET
    Response:

    Server is up and running!

Payment Intent Logic

The application calculates the order amount based on the product_id provided in the request. Current product prices:

    Forever Pants: $260.00
    Forever Shirt: $155.00
    Forever Shorts: $300.00

Technologies Used

    Go: Backend development.
    Stripe API: Payment processing.

Notes

    Security: Use environment variables to store your Stripe API key instead of hardcoding it.
    Currency: This implementation assumes all payments are in USD.

License

This project is licensed under the MIT License.


Let me know if you need any further customizations!


