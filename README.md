\--- README.md ---

# Go URL Shortener

A simple and fully functional URL shortener built with Go (Golang) backend, modernc SQLite driver, and a lightweight HTML/CSS/JavaScript frontend.

---

## Features

* Shorten any URL to a 6-character short code
* Redirect short URLs to the original URL
* Lightweight web interface
* Persistent storage using SQLite
* Fully functional without CGO dependency

---

## Setup Instructions

1. Clone the repository:

```bash
git clone https://github.com/yourusername/url-shortener.git
cd url-shortener/backend
```

2. Install the SQLite driver:

```bash
go get modernc.org/sqlite
```

3. Start the server:

```bash
go run .
```

4. Open your browser at:

```
http://localhost:8080
```

---

## Usage

* **Shorten a URL**: Enter the original URL in the input field and click the `Shorten` button.
* **Redirect**: Click the generated short URL or visit `http://localhost:8080/s/<shortcode>` to be redirected to the original URL.

---

## Project Structure

```
backend/      # Go backend files (API, database, handlers)
frontend/     # HTML, CSS, and JavaScript frontend files
README.md     # Project documentation
LICENSE       # MIT License
```

---

## Screenshots

*Add screenshots of the web interface here to showcase your UI.*

---

## License

This project is licensed under the [MIT License](./LICENSE).