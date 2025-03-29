# Todo Application with Microservices
This is a simple Todo application built with React and TypeScript for the frontend and Go and MongoDB for the backend with microservices architecture. The application allows users to perform basic CRUD (Create, Read, Update, Delete) operations on todo items. The todo items are stored in a MongoDB database and each action performed on the todo items is logged by a logger service in the MongoDB database.

# How to run the application

## Prerequisites

- Docker
- Docker Compose
- Go
- Node
- Bun or npm

## Running the application
- Clone the repository
- Go to the project directory
- Run `make up_build` to build and run all the services
- Go to `http://localhost:8080` to access the application