# Weather API

Weather API is a microservice that provides weather information via gRPC and REST endpoints. It retrieves the temperature of a specified country using gRPC and also offers a simple health check endpoint over REST.

## Features

- gRPC microservice endpoint (`GetWeather`) to retrieve the temperature of a specified country.
- REST endpoint (`/alive`) to check if the server is working.
- gRPC server runs on port 9003.
- REST server runs on port 8003.

## Getting Started

To run the Weather API application, you need to have Docker installed on your system.

1. Clone the repository:

   ```bash
   git clone https://github.com/ctuzelov/weather-api.git
   ```

2. Navigate to the project directory:

   ```bash
   cd weather-api
   ```

3. Run the following command to start the application using Docker Compose:

   ```bash
   docker-compose up
   ```

This command will build the necessary Docker images and start the containers for the gRPC and HTTP servers, along with a PostgreSQL database.
