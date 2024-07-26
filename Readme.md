# CloudStorage Project

## Overview

CloudStorage is a file-sharing application developed using Go version 1.22.5. It allows users to upload, download, share, and manage files and directories. The application provides a RESTful API for these operations.

## Features

- User registration and authentication
- Upload and download files
- Create, rename, and delete directories
- Share files with other users
- Manage user permissions for shared files

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/MirsaidovUmed/CloudStorage.git
    ```
2. Navigate to the project directory:
    ```sh
    cd CloudStorage
    ```
3. Install the required dependencies:
    ```sh
    go get ./...
    ```
4. Set up the PostgreSQL database and create the necessary tables:
    ```sql
    CREATE DATABASE cloud_storage;

    CREATE TABLE users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(255) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL
    );

    CREATE TABLE files (
        id SERIAL PRIMARY KEY,
        file_name VARCHAR(255) NOT NULL,
        user_id INT REFERENCES users(id),
        directory_id INT REFERENCES directories(id)
    );

    CREATE TABLE directories (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        user_id INT REFERENCES users(id)
    );

    CREATE TABLE file_access (
        id SERIAL PRIMARY KEY,
        file_id INT REFERENCES files(id),
        grantor_id INT REFERENCES users(id),
        grantee_id INT REFERENCES users(id)
    );
    ```
5. Run the application:
    ```sh
    go run cmd/app/main.go
    ```

## Libraries Used

The project uses the following libraries:

- [github.com/go-playground/validator/v10 v10.22.0](https://github.com/go-playground/validator) - For validation
- [github.com/gorilla/mux v1.8.1](https://github.com/gorilla/mux) - For routing
- [github.com/jackc/pgx/v5 v5.6.0](https://github.com/jackc/pgx) - For PostgreSQL database interaction
- [github.com/mattn/go-colorable v0.1.13](https://github.com/mattn/go-colorable) - For colored console output
- [github.com/sirupsen/logrus v1.9.3](https://github.com/sirupsen/logrus) - For logging
- [golang.org/x/crypto v0.19.0/](https://pkg.go.dev/golang.org/x/crypto) - For cryptographic functions

## API Endpoints

### User Endpoints

#### Registration
- **URL:** `/api/registration`
- **Method:** `POST`
- **Description:** Register a new user.

#### Login
- **URL:** `/api/login`
- **Method:** `POST`
- **Description:** Authenticate a user and return a token.

### User Management Endpoints (Admin)

#### Get User List
- **URL:** `/admin/users/list`
- **Method:** `GET`
- **Description:** Get a list of all users.

#### Get User By ID
- **URL:** `/admin/users/get/{id}`
- **Method:** `GET`
- **Description:** Get details of a specific user.

#### Update User By ID
- **URL:** `/admin/users/update/{id}`
- **Method:** `PUT`
- **Description:** Update a specific user.

#### Delete User
- **URL:** `/admin/users/delete/{id}`
- **Method:** `DELETE`
- **Description:** Delete a specific user.

### File Endpoints

#### Upload File
- **URL:** `/files/upload`
- **Method:** `POST`
- **Description:** Upload a new file.

#### Get File List
- **URL:** `/files/list`
- **Method:** `GET`
- **Description:** Get a list of all files for the authenticated user.

#### Get File By ID
- **URL:** `/files/get/{id}`
- **Method:** `GET`
- **Description:** Get details of a specific file.

#### Delete File
- **URL:** `/files/remove/{id}`
- **Method:** `DELETE`
- **Description:** Delete a specific file.

#### Rename File
- **URL:** `/files/rename/{id}`
- **Method:** `PUT`
- **Description:** Rename a specific file.

### Directory Endpoints

#### Create Directory
- **URL:** `/directories/create`
- **Method:** `POST`
- **Description:** Create a new directory.

#### Rename Directory
- **URL:** `/directories/rename/{id}`
- **Method:** `PUT`
- **Description:** Rename a specific directory.

#### Get Directory By ID
- **URL:** `/directories/get/{id}`
- **Method:** `GET`
- **Description:** Get details of a specific directory.

#### Delete Directory
- **URL:** `/directories/delete/{id}`
- **Method:** `DELETE`
- **Description:** Delete a specific directory.

### File Sharing Endpoints

#### Get File Access Users
- **URL:** `/files/share/{id}`
- **Method:** `GET`
- **Description:** Get the list of users who have access to a specific file.

#### Share File
- **URL:** `/files/share/{id}/{user_id}`
- **Method:** `PUT`
- **Description:** Share a file with another user.

#### Delete File Access
- **URL:** `/files/share/{id}/{user_id}`
- **Method:** `DELETE`
- **Description:** Remove a user's access to a specific file.
