# go-mongodb-graphql-crud
A schema-first GraphQL API for job listings built with Go and MongoDB. Features full CRUD functionality and custom BSON mapping.

## 🚀 Features
- **GraphQL API**: Flexible data querying and mutations using `gqlgen`.
- **MongoDB Integration**: Persistent storage using the official Go MongoDB driver.
- **Full CRUD Support**: 
  - **Create**: Add new job listings with title, company, and description.
  - **Read**: Fetch all listings or a specific job by its unique ID.
  - **Update**: Partially update job details (Title, URL, etc.).
  - **Delete**: Remove listings from the database by ID.
- **Type Safety**: Custom Go structs mapped to GraphQL types with specific BSON tags for MongoDB `_id` compatibility.

## 🛠 Tech Stack
- **Language**: Go (Golang) 1.22+
- **GraphQL Library**: [gqlgen](https://github.com/99designs/gqlgen)
- **Database**: MongoDB (v7.0)
- **OS**: Linux (Ubuntu 24.04 / Noble Numbat) on Dell Inspiron

## 📦 Getting Started

### 1. Prerequisites
Ensure you have Go installed and MongoDB running on your local machine.

**To start MongoDB on Ubuntu:**
```bash
sudo systemctl start mongod
# Verify it is running (Look for "active (running)")
sudo systemctl status mongod

```

### 2. Installations
Clone the repository and install dependencies:

```bash
git clone [https://github.com/doha-ms/go-mongodb-graphql-crud.git](https://github.com/doha-ms/go-mongodb-graphql-crud.git)
cd go-mongodb-graphql-crud
go mod tidy
```

### 3. Running the server
Start the backend service:

```bash
go run server.go
```


