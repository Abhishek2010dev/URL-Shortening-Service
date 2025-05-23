# URL Shortening Service

A simple, scalable URL shortening service built with Go, Fiber, PostgreSQL, and SQLx. Inspired by the [roadmap.sh URL Shortening Service project](https://roadmap.sh/projects/url-shortening-service).

## Features

- ✅ Create short URLs from long URLs
- 🔁 Redirect to the original long URL using the short code
- 📊 View statistics for each short URL
- ✏️ Update the destination of a short URL
- ❌ Delete short URLs
- 🧪 REST API with full error handling and logging

---

## Tech Stack

- **Language:** Go
- **Web Framework:** [Fiber v3](https://github.com/gofiber/fiber)
- **Database:** PostgreSQL
- **ORM:** [sqlx](https://github.com/jmoiron/sqlx)
- **Logging:** Fiber’s built-in middleware
- **Modular structure:** `handler/`, `repository/`, `database/`, `utils/`

---

## Project Structure

```bash
.
├── main.go
├── database/
│   └── postgres.go       # SQLx connection
├── handler/
│   └── shorten.go        # HTTP Handlers
├── repository/
│   └── shorten.go        # DB Queries
├── utils/
│   └── error.go          # Global error handler
```

---

## Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/Abhishek2010dev/URL-Shortening-Service.git
cd URL-Shortening-Service
```

### 2. Run the App

```bash
go run main.go
```

Server runs at: [http://localhost:3000](http://localhost:3000)

---

## API Endpoints

| Method | Endpoint                     | Description              |
| ------ | ---------------------------- | ------------------------ |
| POST   | `/shorten`                   | Create a new short URL   |
| GET    | `/shorten/:short_code`       | Get original URL         |
| GET    | `/shorten/:short_code/stats` | Get URL statistics       |
| PATCH  | `/shorten/:short_code`       | Update original URL      |
| DELETE | `/shorten/:short_code`       | Delete a short URL       |
| GET    | `/:short_code`               | Redirect to original URL |

## Contributing

PRs and feedback are welcome! If you find a bug or have a feature request, feel free to open an issue.

---

## License

MIT © [Abhishek2010dev](https://github.com/Abhishek2010dev)

---

Let me know if you’d like me to generate this as a downloadable file.
