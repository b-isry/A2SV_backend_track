# A2SV Backend Track

This repository contains a collection of backend-related projects and exercises, primarily written in Go, for learning and practicing backend engineering concepts as part of the A2SV program. The repository features RESTful APIs, command-line tools, and practical implementations of core backend patterns.

## Repository Structure

- **Task-manager_JWT/**: A RESTful API for task management built with Go and MongoDB, featuring JWT authentication.
- **Task Management/**: Implementation of a task management API in Go.
- **Task Management_testify/**: Comprehensive unit and integration tests for the Task Management API, using the [Testify](https://github.com/stretchr/testify) testing library.
- **Task 1/**, **Task 2/**, **Task 3/**: Individual Go-based exercises, such as a grade calculator, a word frequency counter, and a command-line library management system.

## Highlighted Projects

### Task Manager API (Task-manager_JWT)
A RESTful API for managing tasks, using Go and MongoDB.

**Features:**
- CRUD operations for tasks
- JWT-based authentication (inferred from folder name)
- MongoDB integration

**Tech Stack:**  
- Go (1.16+)
- MongoDB

**Getting Started:**
```bash
git clone https://github.com/b-isry/A2SV_backend_track.git
cd Task-manager_JWT
go mod download
go run main.go
```
The server will start at `http://localhost:8080`.

See [`Task-manager_JWT/README.md`](Task-manager_JWT/README.md) and [`Task-manager/docs/api_documentation.md`](Task-manager/docs/api_documentation.md) for detailed setup and API documentation.

### Task Management Testing (Task Management_testify)
This folder contains automated tests for the Task Management API, written with the Go [Testify](https://github.com/stretchr/testify) library. These tests cover core functionalities and ensure correctness of the API implementation.

### Library Management System (Task 3)
A command-line interface (CLI) tool to manage books and members in a library.

- Add, remove, borrow, return, and reserve books
- Manage library members
- List available and borrowed books

See [`Task 3/library_management/docs/documentation.md`](Task%203/library_management/docs/documentation.md) for details.

## Technologies Used

- **Go**: Main programming language for all projects
- **MongoDB**: Used for API data persistence in the Task Manager project
- **Testify**: Go testing framework for unit and integration tests
- **Git**: For version control

## Project Setup

Each major project or exercise is self-contained. See their respective folders and `README.md` or docs for setup and running instructions.

General prerequisites:
- [Go](https://go.dev/dl/) (version 1.16 or higher recommended)
- [MongoDB](https://www.mongodb.com/try/download/community) (for projects requiring persistence)

## Contribution

Contributions are welcome! Please open issues or pull requests for improvements or additional backend exercises.

## License

This repository is licensed under the MIT License.

---

**Note:**  
- For complete and up-to-date details, always refer to the latest files in the GitHub repository.
````
