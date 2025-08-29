# Go URL Shortener

![Go Version](https://img.shields.io/badge/Go-1.24.1-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)

## Overview

This project is a lightweight and fully functional URL shortener built with Go (Golang) backend using the modernc SQLite driver. It includes a minimal frontend developed with HTML, CSS, and JavaScript, allowing users to shorten URLs and redirect via a simple web interface.

## Features

* **Shorten URLs:** Generate concise 6-character short codes for any URL.
* **Redirect Functionality:** Access the short URL to automatically redirect to the original link.
* **Lightweight Frontend:** Minimal UI for inputting URLs and viewing results.
* **Persistent Storage:** SQLite database stores URL mappings.
* **No CGO Dependency:** Works out-of-the-box on Windows, Linux, and macOS.

## Technologies Used

* **Backend:** Go (Golang)
* **Database:** SQLite (via `modernc.org/sqlite` driver)
* **Frontend:** HTML, CSS, JavaScript

## Getting Started

### Prerequisites

* Go version 1.24.1 or higher ([Download Go](https://golang.org/dl/))

### Installation & Running

1. **Clone the repository:**

```bash
git clone https://github.com/firatmio/go-url-shortener
cd url-shortener/backend
```

2. **Install dependencies:**

```bash
go get modernc.org/sqlite
```

3. **Run the server:**

```bash
go run .
```

Server will start at `http://localhost:8080`.

4. **Access the application:**
   Open your browser and navigate to `http://localhost:8080`.

## API Endpoint

| Method | Endpoint       | Description                                    |
| ------ | -------------- | ---------------------------------------------- |
| `POST` | `/api/shorten` | Generates a short URL for a given original URL |

### Request Example

```json
{
    "url": "https://example.com"
}
```

### Response Example

```json
{
    "short_url": "http://localhost:8080/s/abc123"
}
```

## Project Structure

```
.
├── backend/         # Go backend source code
│   ├── main.go      # Main application entry point, server setup
│   ├── handlers.go  # HTTP request handlers for API endpoints
│   ├── db.go        # Database initialization and connection
│   ├── go.mod       # Go module definitions
│   └── go.sum       # Go module checksums
├── frontend/        # Frontend files
│   ├── index.html   # Main HTML page
│   ├── style.css    # CSS for styling
│   └── script.js    # JavaScript for frontend logic
├── LICENSE          # Project License file
└── README.md        # This file
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.