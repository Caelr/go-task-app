I’m here to help! It seems there might have been a misunderstanding. Let’s get it right this time. Here’s the complete README content in a clean Markdown format:

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

1. Clone the repository:

   ```bash
   git clone github.com/Caelr/go-task-app
   cd go-task-app

	2.	Navigate to the root directory and install necessary Go dependencies:

go mod tidy


	3.	Set up your MongoDB database. Make sure to have a running instance of MongoDB. Update the database connection string in the Go application configuration.
	4.	Start the backend server using Air:

air



Frontend Setup

	1.	Navigate to the client directory:

cd client


	2.	Install the frontend dependencies using Bun:

bun install


	3.	Start the frontend development server:

bun dev



Running the Application

	1.	Make sure both the backend and frontend servers are running.
	2.	Open your browser and navigate to http://localhost:3000 (or the port specified by Vite) to access the application.

API Endpoints

Method	Endpoint	Description
GET	/api/v1/todos	Retrieve all tasks
POST	/api/v1/todos	Create a new task
PUT	/api/v1/todos/:id	Update an existing task
DELETE	/api/v1/todos/:id	Delete a task




