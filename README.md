# Go-based API Rate Limiter using Token Bucket Algorithm

## Overview

This project is a simple API rate limiter implemented in Go. The rate limiting mechanism is based on the Token Bucket Algorithm. Users need to obtain a client key by accessing the `/token` endpoint and then include this key in the header (`X-Client-Key`) of subsequent requests to the `/ping` endpoint.

## Token Bucket Algorithm

### Introduction

The Token Bucket Algorithm is a widely used method for controlling the rate of traffic or events in a system. In the context of this project, it helps in limiting the number of API requests a client can make over a specified time period.

### How it Works

The Token Bucket Algorithm uses a metaphorical "bucket" that holds tokens. Tokens are added to the bucket at a fixed rate, and clients can only make a request if there is a token available in the bucket. If there are no tokens, the request is denied. This allows for bursts of requests within the limits of the bucket's capacity, but it also enforces an overall rate limit.

### Implementation in this Project

1. When a client accesses the `/token` endpoint, a certain number of tokens are added to their bucket. The number of tokens added per request and the rate at which tokens are replenished can be configured in the project settings.

2. Clients need to include their client key in the header (`X-Client-Key`) of the `/ping` endpoint request.

3. When a request hits the `/ping` endpoint, the middleware checks if the client has a valid key and if there are available tokens in their bucket.

4. If the conditions are met, the request is allowed to proceed, and a token is consumed. Otherwise, the request is rejected.

## Getting Started

### Prerequisites

- Golang 1.17+
- Git installed on your machine

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/sahildotexe/go-rate-limit.git
   cd go-rate-limit
   ```
2. Run the application
   ```bash
   go run main.go
   ```
3. The API server should now be running locally.

You can test the API endpoints by making requests as described in the following section.

## API Endpoints

The following endpoints are available for testing:

### `/token`

- **Method:** `GET`
- **Description:** Generates client key and adds tokens to the client's bucket.
- **Headers:**
  - None
- **Response:** Returns the client key

### `/ping`

- **Method:** `GET`
- **Description:** A sample API endpoint that requires a valid client key and available tokens for access.
- **Headers:**
  - `X-Client-Key`: Client key obtained from the `/token` endpoint.

You can use tools like `curl` or Postman to interact with these endpoints and test the functionality.
