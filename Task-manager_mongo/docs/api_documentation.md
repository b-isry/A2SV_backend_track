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

## API Endpoints

### Base URL
```
http://localhost:8080
```

### 1. Get All Tasks
```http
GET /tasks
```

**Response (200 OK)**
```json
{
    "message": "Tasks retrieved successfully",
    "data": [
        {
            "id": "507f1f77bcf86cd799439011",
            "title": "Complete Project",
            "description": "Finish the task manager API",
            "due_date": "2024-03-25T15:00:00Z",
            "status": "pending"
        }
    ]
}
```

### 2. Get Task by ID
```http
GET /tasks/:id
```

**Parameters**
- `id`: MongoDB ObjectID of the task

**Response (200 OK)**
```json
{
    "message": "Task retrieved successfully",
    "data": {
        "id": "507f1f77bcf86cd799439011",
        "title": "Complete Project",
        "description": "Finish the task manager API",
        "due_date": "2024-03-25T15:00:00Z",
        "status": "pending"
    }
}
```

### 3. Create Task
```http
POST /tasks
Content-Type: application/json
```

**Request Body**
```json
{
    "title": "Complete Project",
    "description": "Finish the task manager API",
    "due_date": "2024-03-25T15:00:00Z",
    "status": "pending"
}
```

**Response (201 Created)**
```json
{
    "message": "Task created successfully",
    "data": {
        "id": "507f1f77bcf86cd799439011",
        "title": "Complete Project",
        "description": "Finish the task manager API",
        "due_date": "2024-03-25T15:00:00Z",
        "status": "pending"
    }
}
```

### 4. Update Task
```http
PUT /tasks/:id
Content-Type: application/json
```

**Parameters**
- `id`: MongoDB ObjectID of the task

**Request Body**
```json
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

### 5. Delete Task
```http
DELETE /tasks/:id
```

**Parameters**
- `id`: MongoDB ObjectID of the task

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
    "message": "Validation failed",
    "data": [
        {
            "field": "title",
            "message": "Title is required"
        }
    ]
}
```

### 404 Not Found
```json
{
    "message": "Task not found"
}
```

### 500 Internal Server Error
```json
{
    "message": "Failed to retrieve tasks"
}
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