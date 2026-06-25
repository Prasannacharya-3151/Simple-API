# 🔐 Simple Authentication API (Go + Gin + GORM + PostgreSQL + Docker + JWT)

A simple REST API built using **Go**, **Gin**, **GORM**, **PostgreSQL**, **Docker**, and **JWT Authentication**.

This project demonstrates how to build a secure authentication system with:

- User Signup
- User Login
- Password Hashing using bcrypt
- JWT Token Authentication
- Protected Routes
- PostgreSQL Database
- Docker Container

---

# 📚 Technologies Used

- Go
- Gin Framework
- GORM ORM
- PostgreSQL
- Docker
- JWT
- bcrypt
- godotenv

---

# 📁 Project Structure

```
simple-api/

├── config/
│   └── db.go
│
├── handlers/
│   └── auth.go
│
├── middleware/
│   └── auth.go
│
├── models/
│   └── user.go
│
├── routes/
│   └── routes.go
│
├── utils/
│   └── jwt.go
│
├── .env
├── go.mod
└── main.go
```

---

# 📖 Folder Explanation

## main.go

Application entry point.

Responsibilities:

- Load environment variables
- Connect PostgreSQL database
- Auto create database tables
- Create Gin server
- Register all API routes
- Start HTTP server

---

## config/

Responsible for database connection.

Responsibilities:

- Read database credentials from `.env`
- Connect PostgreSQL using GORM
- Store database connection globally

---

## models/

Contains database models.

### User Model

Stores:

- ID
- Name
- Email
- Password
- Created Time
- Updated Time

Also contains request models:

- SignupInput
- LoginInput

These models are used to validate incoming JSON requests.

---

## handlers/

Contains the business logic.

### Signup

Steps:

1. Read JSON request
2. Validate request
3. Check if email already exists
4. Hash password using bcrypt
5. Save user into PostgreSQL
6. Generate JWT Token
7. Return success response

---

### Login

Steps:

1. Read login request
2. Find user by email
3. Compare entered password with hashed password
4. Generate JWT Token
5. Return JWT Token

---

### Profile

Protected API.

Returns logged-in user information.

---

## middleware/

Contains JWT Authentication Middleware.

Responsibilities:

- Read Authorization Header
- Extract JWT Token
- Validate JWT
- Get User ID and Email
- Store user information in Gin Context
- Allow access to protected routes

---

## routes/

Defines all API routes.

```
POST /api/signup

POST /api/login

GET  /api/profile
```

Protected routes use:

```
AuthMiddleware()
```

---

## utils/

Contains JWT helper functions.

### GenerateJWT()

Creates JWT token after successful signup or login.

### ValidateJWT()

Checks:

- Token Signature
- Expiration Time
- User Claims

Returns:

- User ID
- Email

---

# 🔄 Authentication Flow

```
User

↓

Signup

↓

Validate Request

↓

Hash Password

↓

Save User

↓

Generate JWT

↓

Return Token

↓

Frontend Stores Token

===========================

User Login

↓

Verify Email

↓

Compare Password

↓

Generate JWT

↓

Return Token

===========================

Protected Request

↓

Authorization Header

↓

JWT Middleware

↓

Validate Token

↓

Extract User Information

↓

Allow Request

↓

Return Response
```

---

# 🔐 JWT Authentication

JWT contains:

- User ID
- Email
- Expiration Time

JWT is **NOT stored in the database**.

JWT is stored on the client (Browser or Mobile App).

Every protected request sends:

```
Authorization: Bearer <JWT_TOKEN>
```

The server validates the token using the secret key stored in `.env`.

---

# 🔒 Password Security

Passwords are never stored as plain text.

Process:

```
User Password

↓

bcrypt Hash

↓

Store Hashed Password
```

During Login:

```
Entered Password

↓

Compare With Stored Hash

↓

Login Success
```

---

# 🐳 Running PostgreSQL with Docker

## Step 1 - Run PostgreSQL Container

```bash
docker run --name auth-postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_DB=authdb -p 5432:5432 -d postgres:16
```

### Explanation

- `docker run` → Creates and starts a new Docker container.
- `--name auth-postgres` → Gives the container the name **auth-postgres**.
- `POSTGRES_USER` → Creates the PostgreSQL username.
- `POSTGRES_PASSWORD` → Creates the PostgreSQL password.
- `POSTGRES_DB` → Creates the database named **authdb**.
- `-p 5432:5432` → Maps Docker's PostgreSQL port to your local machine.
- `-d` → Runs the container in the background.
- `postgres:16` → Downloads and runs PostgreSQL version 16.

---

## Step 2 - Check Running Containers

```bash
docker ps
```

Shows all running Docker containers.

---

## Step 3 - Enter PostgreSQL

```bash
docker exec -it auth-postgres psql -U postgres -d authdb
```

### Explanation

- `docker exec` → Run a command inside an existing container.
- `-it` → Opens an interactive terminal.
- `auth-postgres` → Container name.
- `psql` → PostgreSQL command-line tool.
- `-U postgres` → Login using the postgres user.
- `-d authdb` → Connect to the authdb database.

---

## Step 4 - Show All Databases

```sql
\l
```

Lists all databases.

---

## Step 5 - Show Tables

```sql
\dt
```

Lists all tables inside the current database.

---

## Step 6 - Show Table Structure

```sql
\d users
```

Displays the columns of the `users` table.

---

## Step 7 - Show User Data

```sql
SELECT * FROM users;
```

Displays all rows stored in the `users` table.

---

## Step 8 - Exit PostgreSQL

```sql
\q
```

Closes the PostgreSQL terminal.

---

## Stop PostgreSQL Container

```bash
docker stop auth-postgres
```

Stops the PostgreSQL container.

---

## Start PostgreSQL Container Again

```bash
docker start auth-postgres
```

Starts the existing PostgreSQL container.

---

# ⚙️ Environment Variables

Create a `.env` file.

```env
PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=mysecretpassword
DB_NAME=authdb

JWT_SECRET=your_secret_key
```

---

# ▶️ Install Dependencies

```bash
go mod tidy
```

Downloads all required Go packages.

---

# ▶️ Run the Application

```bash
go run main.go
```

Starts the Go server.

During startup:

1. Loads `.env`
2. Connects to PostgreSQL
3. Creates the `users` table (AutoMigrate)
4. Registers API routes
5. Starts the HTTP server

---

# 📌 API Endpoints

## Signup

```
POST /api/signup
```

Request

```json
{
  "name": "Prasanna",
  "email": "abc@gmail.com",
  "password": "123456"
}
```

---

## Login

```
POST /api/login
```

Request

```json
{
  "email": "abc@gmail.com",
  "password": "123456"
}
```

---

## Profile

```
GET /api/profile
```

Header

```
Authorization: Bearer <JWT_TOKEN>
```

---

# 🎯 Learning Outcomes

This project demonstrates:

- REST API Development
- Gin Framework
- GORM ORM
- PostgreSQL
- Docker
- JWT Authentication
- Password Hashing (bcrypt)
- Middleware
- Protected Routes
- Environment Variables
- Clean Project Structure

---

# 📄 License

This project was created for learning and educational purposes.
