# Receipt Processor API

A lightweight Go-based API for processing receipts and calculating reward points, containerized using Docker for easy deployment.

## Prerequisites

* Docker
* Docker Compose (optional for easier deployment)

## Getting Started

1. Clone the Repository

    ```
    git clone https://github.com/yourusername/receipt-processor.git
    cd receipt-processor
    ```

2. Build and Start the Application

    Use Docker Compose to build and run the application:
    ```
    docker compose up -d --build
    ```
    The API will be accessible at http://localhost:8080

3. Verify the Application is Running

    You can perform a status check to ensure the API is up:
    ```
    curl http://localhost:8080/status-check
    ```

    Expected Response:
    ```
    Status Check: OK
    ```

## API Endpoints

### Process Receipt

* Endpoint: POST /receipts/process

* Description: Submit a receipt to calculate and store reward points.

* Request Body:
    ```json
    {
        "retailer": "Target",
        "purchaseDate": "2022-01-01",
        "purchaseTime": "13:01",
        "items": [
            {
                "shortDescription": "Mountain Dew 12PK",
                "price": "6.49"
            },
            {
                "shortDescription": "Emils Cheese Pizza",
                "price": "12.25"
            },
            {
                "shortDescription": "Knorr Creamy Chicken",
                "price": "1.26"
            },
            {
                "shortDescription": "Doritos Nacho Cheese",
                "price": "3.35"
            },
            {
                "shortDescription": "Klarbrunn 12-PK 12 FL OZ",
                "price": "12.00"
            }
        ],
        "total": "35.35"
    }
    ```

* Response:

    ```json
    {
        "id": "7fb1377b-b223-49d9-a31a-5a02701dd310"
    }
    ```

### Get Points

* Endpoint: GET /receipts/{id}/points

* Description: Retrieve the calculated points for a specific receipt.

* Path Parameter:

    * id: The unique identifier of the receipt.

* Response:

    ```json
    {
        "points": 28
    }
    ```

## Stopping the Application

To gracefully stop the running containers, execute:

```
docker-compose down
```

This command stops and removes the containers defined in the docker-compose.yml file.

## What to Expect

* Logs: As the application runs, Docker Compose will display logs in the terminal, showing incoming requests and any errors.

* Graceful Shutdown: When stopping the application, it will finish processing ongoing requests before shutting down.

* Error Handling: The API provides clear error messages for invalid requests or internal issues.

## Running Tests

To execute the unit tests, run:
```
go test ./...
```
