# MuchToDo Backend

MuchToDo is a backend service built with **Go** that allows users to manage tasks efficiently.  
It provides REST APIs for creating, reading, updating, and deleting tasks, as well as managing user authentication.  
The application uses **MongoDB** for data storage and is fully containerized with **Docker** for easy deployment.

---

## Features

- User authentication with JWT tokens  
- Task management (CRUD operations)  
- MongoDB as a backend database  
- Dockerized for local development and testing  

---

## Requirements

- Go >= 1.25  
- Docker  
- MongoDB (optional if using Docker container)  

---

## Installation & Setup

1. Clone this repository:

```bash
git clone <YOUR_REPO_URL>
cd muchtodo-backend
Install Go dependencies:

go mod tidy


(Optional) Build the application locally:

GOOS=linux GOARCH=amd64 go build -o muchtodo ./cmd/api

Running with Docker

Create a Docker network:

docker network create muchtodo-net


Run MongoDB container:

docker run -d \
  --name muchtodo-mongo \
  --network muchtodo-net \
  -p 27017:27017 \
  -e MONGO_INITDB_ROOT_USERNAME=admin \
  -e MONGO_INITDB_ROOT_PASSWORD=admin123 \
  mongo:6


Run the backend container:

docker run -p 8080:8080 \
  --network muchtodo-net \
  -e MONGO_URI="mongodb://admin:admin123@muchtodo-mongo:27017/muchtodo?authSource=admin" \
  muchtodo


The backend will be available at: http://localhost:8080
