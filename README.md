# E-Commerce API with Fiber and MongoDB

## 📌 Overview
This is a clean and scalable E-Commerce API built with Golang using Fiber as the web framework and MongoDB as the database.

## 🚀 Features
- User authentication with JWT (login, register, profile update)
- Role-based access control (admin, user)
- CRUD operations for products
- Middleware for request validation
- Secure password handling with SRP (Salted Secure Remote Password)
- DTO pattern for request handling
- Graceful database connection handling

## 📂 Project Structure
```
📦 go-rest
├── 📁 config          # Configuration files (DB connection, env variables)
├── 📁 controllers     # API controllers
├── 📁 dtos            # Data Transfer Objects (DTOs)
├── 📁 middlewares     # Request validation and authentication
├── 📁 models          # Database models
├── 📁 repository      # Data access layer
├── 📁 routes          # API route handlers
├── 📁 services        # Business logic
├── 📁 utils           # Utility functions (hashing, JWT, etc.)
├── .env              # Environment variables
├── main.go           # Entry point of the application
├── go.mod            # Go modules file
├── go.sum            # Go modules dependencies
└── README.md         # Documentation
```

## 🔧 Installation & Setup
### 1️⃣ Clone the Repository
```bash
git clone https://github.com/sarafraz7697/go-rest.git
cd go-rest
```

### 2️⃣ Install Dependencies
```bash
go mod tidy
```

### 3️⃣ Configure Environment Variables
Create a `.env` file in the root directory and add the following:
```env
PORT=8080
SECRET_KEY=mysecretkey

DB_HOST=localhost
DB_USER=admin
DB_PASS=12345
DB_PORT=27017
DB_NAME=ecommerce
```

### 4️⃣ Run the Server
```bash
go run main.go
```
Server will start at `http://localhost:8080`

## 📌 API Endpoints
### 🔑 Authentication Routes
| Method | Endpoint     | Description         | Protected |
|--------|-------------|---------------------|-----------|
| POST   | `/auth/register` | Register a new user | ❌ No |
| POST   | `/auth/login`    | Login user         | ❌ No |
| GET    | `/auth/profile`  | Get user profile   | ✅ Yes |

### 🛍️ Product Routes
| Method | Endpoint       | Description         | Protected |
|--------|---------------|---------------------|-----------|
| POST   | `/product`     | Create a product   | ✅ Admin |
| GET    | `/product`     | Get all products   | ❌ No |
| GET    | `/product/:id` | Get product by ID  | ❌ No |
| PUT    | `/product/:id` | Update product     | ✅ Admin |
| DELETE | `/product/:id` | Delete product     | ✅ Admin |

## 🛠️ Built With
- **Golang** – Backend language
- **Fiber** – Web framework
- **MongoDB** – NoSQL database
- **JWT** – Authentication
- **Validator** – Input validation

## 📜 License
This project is licensed under the MIT License.

## 📌 Author
👨‍💻 Developed by Mahdi(https://github.com/sarafraz7697)

