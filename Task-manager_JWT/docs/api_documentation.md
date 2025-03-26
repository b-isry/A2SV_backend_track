# Task Manager API Documentation

## Prerequisites

### MongoDB Setup
1. Install MongoDB Community Edition from [MongoDB Download Center](https://www.mongodb.com/try/download/community)
2. Start MongoDB service
3. Create a database named `taskmanager` (will be created automatically on first connection)

### Environment Variables
```bash
MONGODB_URI=mongodb://localhost:27017  
PORT=8080                             
```

## Base URL
```
http://localhost:8080
```

## Authentication

### Register a New User
```http
POST /register
Content-Type: application/json

{
    "name": "Test User",
    "email": "test@example.com",
    "password": "password123",
    "role": "user"  // or "admin"
}
```

**Response (201 Created)**
```json
{
    "id": "user_id",
    "name": "Test User",
    "email": "test@example.com",
    "role": "user",
    "created_at": "2024-03-25T12:00:00Z",
    "updated_at": "2024-03-25T12:00:00Z"
}
```

### Login
```http
POST /login
Content-Type: application/json

{
    "email": "test@example.com",
    "password": "password123"
}
```

**Response (200 OK)**
```json
{
    "token": "jwt_token",
    "user": {
        "id": "user_id",
        "name": "Test User",
        "email": "test@example.com",
        "role": "user"
    }
}
```

## Protected Endpoints

All protected endpoints require a JWT token in the Authorization header:
```
Authorization: Bearer <your_jwt_token>
```

### Users

#### Get All Users (Admin Only)
```http
GET /users
Authorization: Bearer <token>
```

**Response (200 OK)**
```json
[
    {
        "id": "user_id",
        "name": "Test User",
        "email": "test@example.com",
        "role": "user",
        "created_at": "2024-03-25T12:00:00Z",
        "updated_at": "2024-03-25T12:00:00Z"
    }
]
```

### Tasks

#### Create Task
```http
POST /tasks
Authorization: Bearer <token>
Content-Type: application/json

{
    "title": "Complete Project",
    "description": "Finish the task manager project",
    "due_date": "2024-03-30T00:00:00Z",
    "status": "pending"
}
```

**Response (201 Created)**
```json
{
    "message": "Task created successfully"
}
```

#### Get All Tasks
```http
GET /tasks
Authorization: Bearer <token>
```

**Response (200 OK)**
```json
{
    "message": "Tasks retrieved successfully",
    "data": [
        {
            "id": "task_id",
            "title": "Complete Project",
            "description": "Finish the task manager project",
            "due_date": "2024-03-30T00:00:00Z",
            "status": "pending"
        }
    ]
}
```

#### Get Task by ID
```http
GET /tasks/:id
Authorization: Bearer <token>
```

**Response (200 OK)**
```json
{
    "message": "Task retrieved successfully",
    "data": {
        "id": "task_id",
        "title": "Complete Project",
        "description": "Finish the task manager project",
        "due_date": "2024-03-30T00:00:00Z",
        "status": "pending"
    }
}
```

#### Update Task
```http
PUT /tasks/:id
Authorization: Bearer <token>
Content-Type: application/json

{
    "title": "Updated Task Title",
    "description": "Updated description",
    "status": "in_progress"
}
```

**Response (200 OK)**
```json
{
    "message": "Task updated successfully"
}
```

#### Delete Task
```http
DELETE /tasks/:id
Authorization: Bearer <token>
```

**Response (200 OK)**
```json
{
    "message": "Task deleted successfully"
}
```

## Error Responses

### 400 Bad Request
```json
{
    "error": "Validation error message"
}
```

### 401 Unauthorized
```json
{
    "error": "Authorization header is required"
}
```
or
```json
{
    "error": "Invalid token"
}
```

### 403 Forbidden
```json
{
    "error": "Insufficient permissions"
}
```

### 404 Not Found
```json
{
    "error": "Resource not found"
}
```

### 500 Internal Server Error
```json
{
    "error": "Internal server error message"
}
```

## Environment Variables

The following environment variables are required:

- `MONGODB_URI`: MongoDB connection string (default: "mongodb://localhost:27017")
- `JWT_SECRET`: Secret key for JWT token generation (required for production)

## Testing the API

1. Start the server:
```bash
go run main.go
```

2. Register a new user:
```bash
curl -X POST http://localhost:8080/register \
-H "Content-Type: application/json" \
-d '{
    "name": "Test User",
    "email": "test@example.com",
    "password": "password123",
    "role": "user"
}'
```

3. Login to get a token:
```bash
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{
    "email": "test@example.com",
    "password": "password123"
}'
```

4. Use the token for protected endpoints:
```bash
curl -X GET http://localhost:8080/tasks \
-H "Authorization: Bearer <your_token>"
```

## Data Models

### Task
```go
type Task struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Title       string            `bson:"title" json:"title"`
    Description string            `bson:"description" json:"description"`
    DueDate     time.Time         `bson:"due_date" json:"due_date"`
    Status      string            `bson:"status" json:"status"`
}
```

### Status Values
- `pending`
- `in_progress`
- `completed`

## MongoDB Integration Details

### Connection
- The application automatically connects to MongoDB on startup
- Default connection string: `mongodb://localhost:27017`
- Database name: `taskmanager`
- Collection name: `tasks`

### Error Handling
The API handles various MongoDB-related errors:
- Connection errors
- Invalid ObjectID format
- Document not found
- Database operation errors
- Validation errors

### Data Persistence
- All tasks are stored in MongoDB
- Task IDs are MongoDB ObjectIDs
- Data persists between application restarts

## Running the Application

1. Start MongoDB service
2. Set environment variables (optional)
3. Run the application:
```bash
go run main.go
```

The server will start on port 8080 by default. 