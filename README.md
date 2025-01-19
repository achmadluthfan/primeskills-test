# Comment API

A RESTful API service for managing comments built with Go. This project follows clean architecture principles and SOLID design patterns.

## Project Structure

```
.
├── main.go                    # Entry point aplikasi
├── config/
│   └── database.go           # Konfigurasi database
├── models/
│   └── comment.go            # Model Comment
├── repositories/
│   └── comment_repository.go # Mengelola query database terkait Comment
├── services/
│   └── comment_service.go    # Logika bisnis terkait Comment
├── handlers/
│   └── comment_handler.go    # Handler untuk request HTTP terkait Comment
├── routes/
│   └── routes.go             # Definisi routing aplikasi
└── data/
    └── comments.json         # Data awal untuk database
```

## Installation

1. Clone the repository

```bash
git clone <repository-url>
cd <project-directory>
```

2. Install dependencies

```bash
go mod download
```

3. Run the application

```bash
go run main.go
```

The server will start on port 8080.

## API Endpoints

### Get All Comments

```
GET /comments
```

Response Success (200):

```json
[
  {
    "id": 1,
    "userId": 1,
    "title": "Sample Title 1",
    "body": "Sample Body 1"
  }
]
```

### Create New Comment

```
POST /comments
```

Request Body:

```json
{
  "userId": 1,
  "title": "Sample Title",
  "body": "Sample Body"
}
```

Required fields:

- userId
- title
- body

### Delete Comment

```
DELETE /comments/delete?id=1
```

Parameters:

- id: Comment ID to delete

## Database

The application uses SQLite as the database. The database file will be created automatically as `comments.db` when you first run the application.

If the database is empty during initialization, the application will automatically load initial data from `data/comments.json` if the file exists.

## Project Components

1. **Models**

   - Contains the data structures
   - Defines the Comment struct with its properties

2. **Repositories**

   - Handles database operations
   - Implements CRUD operations for comments

3. **Services**

   - Contains business logic
   - Validates input data
   - Coordinates between handlers and repositories

4. **Handlers**

   - Manages HTTP requests and responses
   - Handles request validation
   - Converts between HTTP and service layer

5. **Routes**

   - Defines API endpoints
   - Sets up request routing

6. **Config**
   - Manages application configuration
   - Handles database initialization

## Error Handling Detail

### Create Comment Endpoint (POST /comments)

Status Code 400 (Bad Request):

- `"Invalid request payload"`: JSON dalam request body tidak valid
- `"UserID is required"`: Field userId kosong atau 0
- `"Title is required"`: Field title kosong
- `"Body is required"`: Field body kosong

Status Code 409 (Conflict):

- `"Comment already exists"`: Terjadi ketika mencoba membuat comment dengan ID yang sudah ada

Status Code 500 (Internal Server Error):

- `"Unexpected error: [error_message]"`: Terjadi kesalahan yang tidak terduga pada server

### Get Comments Endpoint (GET /comments)

Status Code 500 (Internal Server Error):

- `"Failed to retrieve comments"`: Terjadi kesalahan saat mengambil data dari database

### Delete Comment Endpoint (DELETE /comments/delete)

Status Code 400 (Bad Request):

- `"Missing comment ID"`: Parameter id tidak disertakan dalam URL
- `"Invalid comment ID"`: Parameter id bukan angka yang valid

Status Code 404 (Not Found):

- `"Comment not found"`: Comment dengan ID yang diminta tidak ditemukan

Status Code 500 (Internal Server Error):

- `"Failed to delete comment"`: Terjadi kesalahan saat menghapus data dari database

### Method Not Allowed (Status Code 405)

- `"Method not allowed"`: HTTP method yang digunakan tidak sesuai dengan yang diizinkan untuk endpoint tersebut
