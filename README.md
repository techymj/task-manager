# Task Manager REST API – Golang

## Objective  
Build a scalable Task Management REST API in Go demonstrating clean architecture, persistence, concurrency, authentication, and background processing.


## Tech Stack  
- **Language:** Go  
- **Database:** MySQL  
- **Authentication:** JWT  
- **Concurrency:** Goroutines + Channels  
- **Documentation:** Swagger (OpenAPI)  
- **Tools:** VS Code, Postman  



## Features  

### APIs  
- **POST /register** – Register user  
- **POST /login** – Login and get JWT  
- **POST /tasks** – Create task  
- **GET /tasks** – List tasks (pagination + filtering)  
- **GET /tasks/{id}** – Get task  
- **DELETE /tasks/{id}** – Delete task (owner or admin only)  



## Task Model  

```json
{
  "id": "uuid",
  "title": "string",
  "description": "string",
  "status": "pending | in_progress | completed",
  "user_id": "uuid",
  "created_at": "timestamp",
  "updated_at": "timestamp"
}
```


## Authentication & Authorization  

- JWT-based authentication  
- Token must be sent in header:

```
Authorization: Bearer <token>
```

### Roles
- **user** → can access only own tasks  
- **admin** → can access all tasks  


## Pagination & Filtering  

Supported on **GET /tasks**

```
GET '/tasks?page=1&limit=5&status=pending'
```

### Query Parameters
- `page` → page number  
- `limit` → items per page  
- `status` → pending | in_progress | completed  


## Background Worker (Concurrency)  

- When a task is created, its ID is pushed to a buffered channel  
- A worker goroutine:
  - Waits **X minutes** (configurable via env)
  - If task status is `pending` or `in_progress`
  - Automatically marks it as `completed`
- If task is deleted or manually completed earlier:
  - Worker safely skips processing  

### Environment Configuration

```
AUTO_COMPLETE_MINUTES=5
```


## Swagger / OpenAPI  

Swagger UI is available at:

```
http://localhost:8080/swagger/index.html
```

### Generate Swagger Docs

```bash
swag init -g cmd/server/main.go
```


## Folder Structure  

```
cmd/server        → application entry point  
internal/
  handlers        → HTTP handlers  
  services        → business logic  
  repositories    → DB access layer  
  middleware      → JWT validation  
  models          → data models  
  routes          → route registration  
  database        → DB connection  
  config          → env configuration  
migrations        → SQL schema  
docs              → Swagger files  
```


## Database  

### Tables  

**users**
- id  
- email  
- password  
- role  
- created_at  

**tasks**
- id  
- title  
- description  
- status  
- user_id  
- created_at  
- updated_at  


## How to Run  

### 1. Create Database  

```sql
CREATE DATABASE taskdb;
USE taskdb;
SOURCE migrations/create.sql;
```

### 2. Create `.env`

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=yourpassword
DB_NAME=taskdb

JWT_KEY=your_secret_key
AUTO_COMPLETE_MINUTES=5
```


### 3. Run Application  

```bash
go run ./cmd/server
```

Server runs on:

```
http://localhost:8080
```


## Testing APIs  

Using Postman:

1. Register → `/register`  
2. Login → `/login`  
3. Copy JWT token  
4. Use token in header  
5. Create task  
6. Wait X minutes → verify auto-complete  
7. Test pagination & filtering  


## Design Highlights  

- Clean architecture  
- Separation of concerns  
- Repository pattern  
- Env-based configuration  
- Background worker using goroutines  
- Swagger for API documentation  


## Notes  

Unit tests, Docker, and production-grade improvements are planned for future enhancement.  
Current focus is on correctness, concurrency, and clean architecture.


## Author  

Built as part of a **Golang – Task Management REST API**
