## 📚 Golang REST API Demo

This is a **simple REST API** built in **pure Golang** (without frameworks) as a learning project. It demonstrates basic CRUD operations on a collection of **books**.

### 🚀 Features

✅ Pure Golang HTTP server (using `net/http`)
✅ JSON-based REST API
✅ CRUD operations:

* Create a book
* Get all books
* Get a single book by ID
  ✅ Learning-focused: no database, in-memory data store

---

### 📦 API Endpoints

| Method | Endpoint      | Description       |
| ------ | ------------- | ----------------- |
| GET    | `/books`      | Get all books     |
| GET    | `/books/{id}` | Get a book by ID  |
| POST   | `/books`      | Create a new book |

---

### 🛠️ How to Run

1️⃣ **Clone the repo**

```bash
git clone https://github.com/your-username/your-repo.git
cd your-repo
```

2️⃣ **Run the server**

```bash
go run main.go
```

Server starts at: [http://localhost:8000](http://localhost:8000)

---

### 📄 Example Data (Book)

```json
{
  "id": "1",
  "title": "Learn Go",
  "author": "John Doe"
}
```

---

### 🏗️ Tech Stack

* [Golang](https://go.dev/)
* `net/http` standard package
* JSON encoding/decoding with `encoding/json`

---

### 🚧 Future Improvements (Optional)

* Add a real database (e.g., PostgreSQL, MongoDB)
* Implement more endpoints (PUT, DELETE)
* Add error handling and validation
* Use a web framework like Gin, Echo, or Fiber

---

### 📬 Feedback

Feel free to fork, improve, or ask questions! This is just a demo for learning. 🚀