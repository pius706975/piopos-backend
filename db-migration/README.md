
# PioPOS Database 

This repository contains only for database migration with CLI. Use this repository for migrate the database.

## Table of Contents

- [Project Structure](#project-structure)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [Folder Structure](#folder-structure)
- [Contributor](#contributor)

## Project Structure

This service uses [Golang](https://golang.org/) and [GORM](https://gorm.io/) for database ORM. 

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/pius-microservices/piopos-database.git
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Configuration

1. Copy the `.env.example` file to `.env`:
   ```bash
    MODE = development

    # User DB
    USER_DB_PORT = 
    USER_DB_USERNAME = 
    USER_DB_PASSWORD = 
    USER_DB_NAME = 
    USER_DB_HOST = 

    # Product DB
    PRODUCT_DB_PORT = 
    PRODUCT_DB_USERNAME =
    PRODUCT_DB_PASSWORD =
    PRODUCT_DB_NAME = 
    PRODUCT_DB_HOST = 
   ```

2. Update the `.env` file with your environment variables.

## Running the Application

To run database migration, use:
```bash
#To see the list of the DB CLI
go run .

# migrate the database models
go run . <migration command> -u 

# drop database
go run . <migration command> -d
```

## Folder Structure

Here's a breakdown of the project folder structure:

- **cmd/**: Contains command line scripts
  - **command.line.go**: CLI execution entry point

- **config/**: Configuration files, including environment variables
  - **env.go**: Loads and parses environment variables

- **docs/**: API documentation files (Swagger)
  - **swagger.json** and **swagger.yaml**: Swagger specification files
  
- **database/**: Database configuration, models, and migrations
  - **migrations/**: Database migration logic and configuration
  - **models/**: GORM models

- **main.go**: Application entry point

- **.env.example**: Example environment configuration

## 👨‍💻 Contributor

- Pius Restiantoro - [GitHub](https://github.com/pius706975)
