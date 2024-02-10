# Real Estate Application

This is a simple real estate application that provides APIs for managing houses.

## Table of Contents
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Testing](#testing)
- [Docker](#docker)


## Features
- Upload house details
- Get house details
- List all houses
- Update house details (e.g., mark as sold)
- Delete a house
- Logging middleware for request logging

## Prerequisites
- Go installed on your machine
- Docker (for containerization)

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/shubham7999/real_state_assignment.git
   cd real_state_assignment

## Usage

1. Run the application:

```bash
go run main.go

```
2. Access the APIs at http://localhost:8080/api/v1/houses

   

## API Endpoints

### Upload House Details

- **Method:** POST
- **Endpoint:** `/api/v1/houses`
- **Payload:** JSON data with house details

### Get House Details

- **Method:** GET
- **Endpoint:** `/api/v1/houses/{id}`
- **Parameters:** `id` (House ID)

### List All Houses

- **Method:** GET
- **Endpoint:** `/api/v1/houses`

### Update House Details

- **Method:** PUT
- **Endpoint:** `/api/v1/houses/{id}`
- **Parameters:** `id` (House ID)
- **Payload:** JSON data with updated details

### Delete House

- **Method:** DELETE
- **Endpoint:** `/api/v1/houses/{id}`
- **Parameters:** `id` (House ID)




## Docker
To run the application using Docker:

```bash
docker build -t real-estate-app .
docker run -p 8080:8080 real-estate-app
```

## Testing

To ensure the correctness of the application, run the following unit tests:

```bash
1. go test ./...
2. go test ./... -coverprofile=coverage.out
3. go tool cover -func=coverage.out  
```







