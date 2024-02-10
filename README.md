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
- [Contributing](#contributing)
- [License](#license)

## Features
- Upload house details
- Get house details
- List all houses
- Update house details (e.g., mark as sold)
- Delete a house
- Logging middleware for request logging

## Prerequisites
- Go installed on your machine
- Docker (optional, for containerization)



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



## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/your_username/real-estate-app.git
   cd real-estate-app

# Usage

Run the application:

```bash
go run main.go




## Install Dependencies
Run the application:

```bash
go run main.go




## Docker
To run the application using Docker:

```bash
docker build -t real-estate-app .
docker run -p 8080:8080 real-estate-app





