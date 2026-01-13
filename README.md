# MuchToDo

MuchToDo is a task management backend service built in **Go**.  
It provides a REST API to manage users and todos, with data stored in **MongoDB**. The application is containerized using **Docker** for easy deployment.

---

## Table of Contents

- [Features](#features)  
- [Requirements](#requirements)  
- [Setup & Installation](#setup--installation)  
- [Running with Docker](#running-with-docker)  
- [Environment Variables](#environment-variables)  
- [API Endpoints](#api-endpoints)  
- [License](#license)  

---

## Features

- User registration and authentication (JWT-based)  
- CRUD operations for todos  
- MongoDB as the backend database  
- Fully containerized with Docker  
- Swagger documentation available  

---

## Requirements

- [Go](https://golang.org/) >= 1.25  
- [Docker](https://www.docker.com/)  
- [MongoDB](https://www.mongodb.com/) (optional if using Docker)  

---

## Setup & Installation

1. Clone the repository:

```bash
git clone https://github.com/Innocent9712/much-to-do.git
cd much-to-do/Server/MuchToDo
