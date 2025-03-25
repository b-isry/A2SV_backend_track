# Library Management System

This document outlines a command-line Library Management System (LMS) developed in Go. The LMS facilitates basic library operations, enabling librarians to manage books and member accounts.

## Key Features

The system offers the following core functionalities:

* **Add Book:** Registers new books into the library catalog.
* **Add Member:** Creates new member accounts within the system.
* **Remove Book:** Deletes books from the library catalog.
* **Borrow Book:** Facilitates the asynchronous borrowing of books by members.
* **Return Book:** Manages the return of borrowed books.
* **Reserve Book:** Allows members to reserve books, handling concurrent reservations.
* **List Available Books:** Displays all books currently available in the library.
* **List Borrowed Books:** Displays all books currently checked out by members.

## Getting Started

### Prerequisites

* **Go:** Version 1.x or higher is required. Download and install from [https://go.dev/dl/](https://go.dev/dl/).
* **Git:** (Optional) Recommended for cloning the repository. Download and install from [https://git-scm.com/downloads](https://git-scm.com/downloads).

### Installation

1. **Clone the Repository:**

    ```bash
    git clone https://github.com/yourusername/Library-Management.git
    cd Library-Management
    ```

    Alternatively, you can download the source code as a ZIP file and extract it.

2. **Build the Application:**

    ```bash
    go build
    ```

    This command compiles the Go source code and generates an executable file.

3. **Run the Application:**

    ```bash
    ./Library-Management
    ```

    This will start the Library Management System.

## Usage

Upon execution, the system presents a menu-driven interface with the following options:

1. **Add Book:** Prompts for the following information:
    * **Book ID:** A unique integer identifier for the book.
    * **Book Title:** The title of the book.
    * **Book Author:** The author of the book.
    * The book is automatically marked as "available."

2. **Add Member:** Prompts for the following information:
    * **Member ID:** A unique integer identifier for the member.
    * **Member Name:** The name of the member.
    * The member's borrowed books list is initialized as empty.

3. **Remove Book:** Prompts for the following information:
    * **Book ID:** The ID of the book to remove.

4. **Borrow Book:** Prompts for the following information:
    * **Member ID:** The ID of the member borrowing the book.
    * **Book ID:** The ID of the book being borrowed.
    * The system validates:
        * The member exists.
        * The book exists.
        * The book is currently available.
    * The system handles asynchronous borrowing.

5. **Return Book:** Prompts for the following information:
    * **Member ID:** The ID of the member returning the book.
    * **Book ID:** The ID of the book being returned.
    * The system validates:
        * The member exists.
        * The book exists.
        * The book is currently borrowed by the specified member.

6. **Reserve Book:** Prompts for the following information:
    * **Member ID:** The ID of the member reserving the book.
    * **Book ID:** The ID of the book being reserved.
    * The system handles concurrent reservations.

7. **List Available Books:** Displays a list of all books with the status "available."

8. **List Borrowed Books:** Displays a list of all books with the status "borrowed."

9. **Exit:** Terminates the program.

After each operation, the system prompts the user to continue (y/n) to perform another action.

## Project Structure

```
Library-Management/
├── main.go                 # Entry point
├── workers/ 
│   └── reservation_worker.go # Handles multiple reservation requests
├── controllers/
│   └── library_controller.go  # User interface handling
├── models/
│   └── models.go          # Data structures
└── services/
│   └── library_service.go # Business logic
├── Tests/
    └── library_test.go    #Tests functionalities

```

## Data Structures

### Book

* **Id:** `int` - Unique identifier for the book.
* **Title:** `string` - Title of the book.
* **Author:** `string` - Author of the book.
* **Status:** `string` - Current status of the book ("available" or "borrowed").

### Member

* **Id:** `int` - Unique identifier for the member.
* **Name:** `string` - Name of the member.
* **BorrowedBooks:** `[]Book` - A list of books currently borrowed by the member.

## Error Handling

The system includes robust error handling to manage common issues:

* **Member Not Found:** Indicates that the entered member ID does not exist in the system.
* **Book Not Found:** Indicates that the entered book ID does not exist in the system.
* **Book Not Available:** Indicates that the requested book is currently borrowed and unavailable.
* **Book Not Borrowed:** Indicates that the book is not currently borrowed by the specified member (when attempting to return).
* **Invalid Menu Choice:** Indicates that the user entered an invalid option in the main menu.

## Contributing

Contributions are welcome! Please submit issues and enhancement requests through the project's issue tracker. Fork the repository, create a branch for your changes, and submit a pull request.

## License

This project is licensed under the MIT License - see the `LICENSE` file for details.