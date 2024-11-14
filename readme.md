# Rest API Server on Fiber

This project is a Rest API server built using the Fiber framework, with a React front-end and a Python bot for link sharing. The server is structured to maintain clear separation of concerns, making the codebase easy to understand and extend.

## Project Structure

### 1. `cmd`
This is the entry point of the application, containing the `main.go` file that initializes and starts the server.

### 2. `internal`
This directory contains the core logic of the project.

#### 2.1. `app`
This folder contains the application configuration and the `Run` method, which sets up and starts the server.

#### 2.2. `endpoints`
This folder contains methods that handle incoming requests and send appropriate responses.

#### 2.3. `services`
This folder stores the business logic of the application.

#### 2.4. `databases`
This folder contains the code for interacting with the database using Gorm.

#### 2.5. `models`
This folder defines the models for working with the database, including creating tables and associated methods.

#### 2.6. `tests`
This folder contains tests for the application to ensure everything works as expected.

### `bot.py`
A Python bot that sends a link to the application, facilitating easy sharing and access.

### `frontend`
The front-end of the application is built with React, providing a user-friendly interface for interacting with the server.

## Getting Started

### Prerequisites
- Go
- Python
- React


