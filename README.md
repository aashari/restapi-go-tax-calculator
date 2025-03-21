# Tax Calculator API

A Go-based REST API service for calculating taxes on various product types based on their tax codes.

## Features

- Fast tax calculation for individual products or in bulk
- Support for different tax types (Food and Beverage, Tobacco, Entertainment)
- Configurable tax rules (percentages and fixed amounts)
- Support for refundable vs. non-refundable taxes
- API endpoints for both single and bulk calculation

## Prerequisites

- Docker and Docker Compose

## How to Run

To run this project, use the following command:

```bash
docker-compose up --build
```

This will build and start:

- MySQL database container
- Golang application container

The setup process will:

1. Download all Go dependencies
2. Run database schema migration
3. Seed the database with default tax configurations

When the service is ready, you should see output similar to:

```bash
golang_service    | [GIN-debug] POST   /v1/taxes                 --> _/my_app/controllers.(*V1TaxesController).CalculateTax-fm (4 handlers)
golang_service    | [GIN-debug] POST   /v1/taxes/bulk            --> _/my_app/controllers.(*V1TaxesController).CalculateTaxBulk-fm (4 handlers)
golang_service    | 0.0.0.0:3000
golang_service    | [GIN-debug] Listening and serving HTTP on 0.0.0.0:3000
```

The API will be available at: http://localhost:3000

## API Endpoints

### Calculate Tax for Single Item

**Endpoint:** `POST /v1/taxes`

**Request Body:**

```json
{
  "name": "Product Name",
  "tax_code": 1,
  "price": 100.0
}
```

**Response:**

```json
{
  "name": "Product Name",
  "tax_code": 1,
  "type": "Food and Beverage",
  "refundable": "YES",
  "price": 100.0,
  "tax": 10.0,
  "amount": 110.0
}
```

### Calculate Tax for Multiple Items

**Endpoint:** `POST /v1/taxes/bulk`

**Request Body:**

```json
[
  {
    "name": "Product 1",
    "tax_code": 1,
    "price": 100.0
  },
  {
    "name": "Product 2",
    "tax_code": 2,
    "price": 200.0
  }
]
```

**Response:** Array of tax calculation results for each item.

## Tax Codes

| Tax Code | Description       | Refundable |
| -------- | ----------------- | ---------- |
| 1        | Food and Beverage | Yes        |
| 2        | Tobacco           | No         |
| 3        | Entertainment     | No         |

## Project Structure

```
src/
├── controllers/        # API request handlers
├── services/           # Business logic
├── repositories/       # Database interaction
├── models/             # Data models
├── objects/            # Request/response objects
├── database/           # Database connection and migrations
├── middleware/         # HTTP middleware
├── helpers/            # Utility functions
├── constants/          # Constant values
└── main.go             # Application entry point
```

## Detailed API Documentation

For comprehensive API documentation, see:
[Postman Documentation](https://documenter.getpostman.com/view/883805/RztitAQK)
