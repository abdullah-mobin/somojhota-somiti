# Somojhota Somiti API

Backend API for the Somojhota Somiti project, built using Go (Golang) and the Fiber web framework. This application manages authentication, transactions, and member data for the association.

## Tech Stack

-   **Language:** Go
-   **Framework:** [Fiber](https://gofiber.io/)
-   **Database:** MongoDB (using official mongo-driver)
-   **Authentication:** JWT (JSON Web Tokens) with Access and Refresh tokens
-   **Documentation:** Swagger (swaggo)

## Features

-   **Authentication:** Secure login, logout, and token refresh mechanisms using JWT and bcrypt.
-   **Transactions:** Create and manage business transactions.
-   **Standardized Responses:** Unified JSON response structure for success, errors, and pagination.
-   **Database:** Optimized MongoDB connection with connection pooling.
-   **API Documentation:** Auto-generated Swagger documentation.

## Prerequisites

-   Go 1.18 or higher
-   MongoDB instance (Local or Atlas)

## Installation

1.  **Clone the repository:**
    ```bash
    git clone <repository-url>
    cd somojhota-somiti
    ```

2.  **Install dependencies:**
    ```bash
    go mod download
    ```

3.  **Configuration:**
    Create a `.env` file in the root directory. Based on the application logic, configure the necessary environment variables:

    ```env
    PORT=3000
    MONGO_URI=mongodb://localhost:27017
    DB_NAME=somojhota_db
    JWT_SECRET=your_secret_key
    # Add other specific variables used in utils/config
    ```

## Running the Application

To start the server:

```bash
go run main.go
```

## API Documentation

This project uses Swagger for API documentation. Once the application is running, you can typically access the interactive documentation at:

```
http://localhost:3000/swagger/index.html
```

*Note: Ensure you have generated the docs using `swag init` if you make changes to the API annotations.*