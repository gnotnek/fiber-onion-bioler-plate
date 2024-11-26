# Fiber Onion Boilerplate

This is a boilerplate project for building web applications using the Fiber framework in Go with an onion architecture.

## Prerequisites

- Go 1.23 or higher
- Make
- Docker (optional)

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/fiber-onion-boilerplate.git
    cd fiber-onion-boilerplate
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage

### Running the Application

To run the application, use the following command:
```sh
make run
```

### Running in Live Mode

To run the application, use the following command:
```sh
make run/live
```

### Running Text
To run the test, use following command:
```sh
make test
```

### Building The Project
To Build the project, use the following command:
```sh
make build
```

### Environment Variable

The following environment variables are used in the project:

```
DATABASE_HOST: The database host
DATABASE_PORT: The database port 
DATABASE_USER: The database user 
DATABASE_PASSWORD: The database password 
DATABASE_NAME: The database name 
JWT_SECRET_KEY: The secret key for JWT 
```

These variables can be set in a .env file in the root directory. An example .env file is provided as .env.example.

### Docker

1. To Build the Docker Image:
    ```sh
    docker build -t fiber-onion-boilerplate .
    ```
2. Run the Docker Container:
    ```sh
    docker run ---env-file .env -p 8080:8080 fiber-onion-boiler-plate
    ```

The application will be accessible at http://localhost:8080.

### License

This project is licensed under the MIT License.
