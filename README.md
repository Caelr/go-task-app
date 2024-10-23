
# Go Task App

A simple task management application built with a Go backend using Fiber and a React frontend using Vite.js.

## Table of Contents

- [Features](#features)
- [Technologies](#technologies)
- [Setup](#setup)
  - [Backend Setup](#backend-setup)
  - [Frontend Setup](#frontend-setup)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

## Features

- Create, read, update, and delete tasks
- User-friendly interface with React
- Fast performance with GoFiber
- Persistent storage using MongoDB

## Technologies

- **Backend:** Go, GoFiber
- **Frontend:** React, TypeScript, Vite.js
- **Database:** MongoDB
- **Development Tools:** Air (for backend live reload), Bun (for frontend dev server)

## Setup

### Backend Setup

1. Navigate to the root directory and install necessary Go dependencies:

   ```bash
   go mod tidy
   ```

2. Set up your MongoDB database. Make sure to have a running instance of MongoDB. Update the database connection string in the Go application configuration.

3. Start the backend server using Air:

   ```bash
   air
   ```

### Frontend Setup

1. Navigate to the `client` directory:

   ```bash
   cd client
   ```

2. Install the frontend dependencies using Bun:

   ```bash
   bun install
   ```

3. Start the frontend development server:

   ```bash
   bun dev
   ```

## Running the Application

1. Make sure both the backend and frontend servers are running.
2. Open your browser and navigate to `http://localhost:5173` (or the port specified by Vite) to access the application.

## API Endpoints

| Method | Endpoint             | Description                |
|--------|----------------------|----------------------------|
| GET    | `/api/v1/todos`      | Retrieve all tasks         |
| POST   | `/api/v1/todos`      | Create a new task          |
| PATCH  | `/api/v1/todos/:id`  | Update an existing task    |
| DELETE | `/api/v1/todos/:id`  | Delete a task              |

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
