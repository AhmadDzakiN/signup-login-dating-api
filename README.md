# Signup-Login Dating App API

- [Signup-Login Dating App API](#signup-login-dating-app-api)
  - [Tech Stack](#tech-stack)
  - [How to Start](#how-to-start)
  - [List of APIs](#list-of-apis)
  - [Project Structure](#project-structure)


## Tech Stack
- Go
- PostgreSQL

## How to Start
- Clone this repo
- Copy file `env.sample` to `.env`
- Modify `params/.env` as your needs
- Ensure your PostgreSQL database already up (You can use `database/seeds.sql` file to create the table needed for this API)
- Run the project by `go run main.go` or `go build && ./signup-login-bumble`

## List of APIs
- Signup / Register User
- Signin / Login User
- Get User by ID (Use Bearer Token Authentication)

To see the examples and descriptions of these APIs, you can import Postman collection from `documents` folder and import it to your Postman

## Project Structure
- Config: All files related to project configuration definition (Other tech stack dependencies, read env file, etc) will be here
- Database: All sql files related to database (seeds & migrations (if any)) will be here
- Documents: All files related to project documentation will be here (Postman collection, pdfs, or any images)
- Domain: I use Domain Driven Design for this project. So every use cases/domains for this project will be seperated each other including the service and route file
- Middleware: All middlewares will be here
- Model: All database model structs and payloads will be here
- Router: Parent configuration of routes in the projejct
- Utils: Utility function that can be served as common function and can used by other function in the project