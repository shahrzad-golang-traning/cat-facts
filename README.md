
## 🐱 Cat Facts CLI - Go HTTP Client

A simple and clean Go program that fetches a random **cat fact** from a public API and prints the response details like status, headers, and JSON body. Perfect as a beginner-friendly example of working with HTTP requests in Go.

---

### 🚀 What It Does

* Sends a `GET` request to the `https://catfact.ninja/fact` API
* Sets custom headers (`Accept` and `User-Agent`)
* Prints the HTTP status, status code, content type, and response body
* Handles all network and read errors gracefully
* Uses `defer` to properly close the response body

---

### 🧠 What You’ll Learn

* How to make HTTP requests in Go using `net/http`
* How to read and handle response data
* How to work with headers
* Good practices like error checking and resource cleanup

---

### 📦 How to Run

Make sure you have **Go** installed. Then:

```bash
go run main.go
```

You’ll see output like this:

```
============= API Response =============
Status: 200 OK
Status Code: 200
Headers: application/json
----------------------------------------
Body:
{"fact":"Cats sleep 70% of their lives.","length":35}
========================================
```

---

### 🔧 Code Overview

```go
req, err := http.NewRequest("GET", url, nil)
req.Header.Add("Accept", "application/json")
req.Header.Add("User-Agent", "GoLang-Client/1.0")
```

> This tells the server: "Hey, I want JSON, and I'm a Go client."

Then we handle the response:

```go
body, err := io.ReadAll(res.Body)
fmt.Println(string(body))
```

---

### 📚 API Reference

This project uses [catfact.ninja](https://catfact.ninja/fact), a free and fun API that provides random cat facts in JSON format.

---

### 💡 Want More?

This project is a great base to add more features:

* Print the fact in a pretty format
* Save it to a file
* Convert JSON to Go struct using `encoding/json`
* Add CLI arguments or flags

---

### 🐾 Author

Made with love (and curiosity for cats) by **Shahrzad Taherzadeh**
🧠 Want to learn more about HTTP clients or Go projects? Let’s connect!

---

