# Task Manager API

## Overview

This API provides endpoints to manage tasks in a task management system. It supports basic CRUD operations (Create, Read, Update, Delete) for tasks.

## Base URL

```
http://localhost:8080

 ```

## Response Format

All endpoints return responses in the following JSON format:

``` json
{
"message": "Description of what happened",
"data": {
}
}

 ```

## Endpoints

### Health Check

#### GET /ping

Check if the API is running.

**Response**

``` json
{
"message": "pong"
}

 ```

### Tasks

#### GET /tasks

Retrieve all tasks.

**Response**

``` json
{
"message": "Tasks retrieved successfully",
"data": [
{
"id": "1",
"title": "Task 1",
"description": "First task",
"due_date": "2024-03-10T12:00:00Z",
"status": "pending"
}
]
}

 ```

#### GET /tasks/:id

Retrieve a specific task by ID.

**Parameters**

- `id`: Task ID (string, in path)
    

**Success Response**

``` json
{
"message": "Task retrieved successfully",
"data": {
"id": "1",
"title": "Task 1",
"description": "First task",
"due_date": "2024-03-10T12:00:00Z",
"status": "pending"
}
}

 ```

**Error Response**

``` json
{
"message": "Task not found"
}

 ```

#### POST /tasks

Create a new task.

**Request Body**

``` json
{
"id": "4",
"title": "New Task",
"description": "Task description",
"due_date": "2024-03-15T12:00:00Z",
"status": "pending"
}

 ```

**Success Response**

``` json
{
"message": "Task created successfully",
"data": {
"id": "4",
"title": "New Task",
"description": "Task description",
"due_date": "2024-03-15T12:00:00Z",
"status": "pending"
}
}

 ```

**Validation Error Response**

``` json
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

#### PUT /tasks/:id

Update an existing task.

**Parameters**

- `id`: Task ID (string, in path)
    

**Request Body**

``` json
{
"title": "Updated Task",
"description": "Updated description",
"status": "in_progress"
}

 ```

**Success Response**

``` json
{
"message": "Task updated successfully"
}

 ```

**Error Responses**

``` json
{
"message": "Task not found"
}

 ```

``` json
{
"message": "Validation failed",
"data": [
{
"field": "status",
"message": "Status must be one of: pending, in_progress, completed"
}
]
}

 ```

#### DELETE /tasks/:id

Delete a task.

**Parameters**

- `id`: Task ID (string, in path)
    

**Success Response**

``` json
{
"message": "Task deleted successfully"
}

 ```

**Error Response**

``` json
{
"message": "Task not found"
}

 ```

## Status Codes

- `200 OK`: Request successful
    
- `201 Created`: Resource created successfully
    
- `400 Bad Request`: Invalid request (validation error)
    
- `404 Not Found`: Resource not found
    
- `500 Internal Server Error`: Server error
    
