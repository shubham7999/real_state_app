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

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/your_username/real-estate-app.git
   cd real-estate-app

## Install Dependencies
```bash
go mod download

## Docker
```bash
docker build -t real-estate-app .
docker run -p 8080:8080 real-estate-app

