# Go-React-App

The goal of the project is to build a simple business loan application system.

## Features

- Business loan application form
- Balance sheet data input
- Decision engine to evaluate loan applications
- Integration with accounting software (MYOB, Xero) and Decision Engine

## Technologies

- Backend: Go, Gin Web Framework
- Frontend: React, Vite.js, TailwindCSS
- Deployment: Docker, Docker Compose

## Backend

The backend is a Go API server built with Gin Web Framework.

### Setup

1. Install Go: [https://golang.org/doc/install](https://golang.org/doc/install)
2. Install dependencies: `go mod download`

### Running the server

```bash
go run main.go
```

### Running tests

```bash
go test ./tests
```

## Frontend

The frontend is a React application built with `Vite.js`.
Create a `.env` file in the `frontend` folder and set the following variables: `VITE_API_BASE_URL=http://localhost:8080/api/v1`
Replace `http://localhost:8080` with the actual backend API URL if it's different.

### Setup

1. Install Node.js: https://nodejs.org/en/download/
2. Install dependencies: `npm install`

### Running the development server

```bash
npm start
```

### Building for production

```bash
npm run build
``` 

### Usage

1. Open the web application in your browser at http://localhost:5173.
2. Enter your Business Details such as:
   - Name
   - Year established
   - Loan Amount
   - Acconting Provider 
3. Review and click "Submit" to apply loan.
4. Final Outcome will show the result of your loan application.

## Docker and Docker Compose

### Building and running with Docker Compose

```bash
docker-compose up --build
```

### Stopping and removing containers

```bash
docker-compose down
```

## Authors

Alex Aung Myo OO

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
