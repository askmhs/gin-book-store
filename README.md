# 📚 Gin Book Store API

A simple RESTful API for managing a bookstore, built with [Gin](https://github.com/gin-gonic/gin) and [GORM](https://gorm.io/) using SQLite.

## 🚀 Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) installed (version 1.16 or higher)

### Installation

**1. Clone the repository:**

   ```bash
   git clone https://github.com/askmhs/gin-book-store.git
   cd gin-book-store
   ```

**2. Install dependencies:**

   ```bash
   go get .
   ```

**3. Set up ENV config:**

   ```bash
   cp .env.example .env
   ```
Update the ENV variables value

**4. Run the application:**

   ```bash
   go run main.go
   ```
The server will start on http://localhost:8080.

## 📖 API Endpoints

- `POST /users/register` - Register a new user
- `POST /useres/login` - Login user
- `GET /books` - Retrieve all books
- `GET /books/:id` - Retrieve a book by ID
- `POST /books` - Create a new book
- `PUT /books/:id` - Update an existing book
- `DELETE /books/:id` - Delete a book

## 🛠 Technologies Used

- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [GORM](https://gorm.io/) - ORM library
- SQLite - Lightweight relational database

## 📄 License

This project is licensed under the MIT License.
