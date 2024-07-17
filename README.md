
# Book Store Management




## Table of Contents

- Installation
- Usage
- Routes

## Installation

- Clone the repository
  git clone https://github.com/your-username/your-repository.git
- cd your-repository
- go mod tidy


## Usage

- go run main.go
  Server will start on http://localhost:9001
  
## Routes

- /book/ Get aall the books stored in DB
- /book/ (POST) Create a new Book
- /book/bookId (GET) Get existing book by bookId
- /book/bookId (PUT) Update a Book
- /book/bookId (DELETE) Delete a book
