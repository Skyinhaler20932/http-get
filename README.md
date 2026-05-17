# 🌐 HTTP Client CLI (Go)

A simple HTTP client built in Go for learning how HTTP requests, error handling, and retries work in real-world network environments.

---

## 🚀 Project Overview

This project is a lightweight CLI HTTP client that:

- Accepts a URL from command-line arguments
- Validates the URL structure
- Sends HTTP GET requests using Go's `net/http`
- Implements retry logic for unstable networks
- Reads and prints the HTTP response body
- Displays HTTP status code and final resolved URL

The main goal is to understand how Go handles networking, errors, and HTTP response lifecycle in real-world conditions.

---

## 🧠 Key Features

- CLI argument handling using `os.Args`
- URL validation using `net/url`
- HTTP GET requests using `net/http`
- Retry mechanism for failed requests
- Safe handling of HTTP responses
- Reading response body using `io.ReadAll`
- Proper resource cleanup with `defer resp.Body.Close()`

---

## ⚙️ How It Works

1. The program reads a URL from the command line
2. It validates the URL format
3. It attempts an HTTP GET request
4. If the request fails, it retries up to N times
5. On success:
   - Reads the response body
   - Prints status code, URL, and body

---

## 📦 Example Usage

```bash
go run main.go https://google.com


Welcome to the HTTP client app

Retrying... 1
Retrying... 2

Request completed!
Status Code: 200
Status: 200 OK
URL: https://www.google.com
Body: <html>...</html>
