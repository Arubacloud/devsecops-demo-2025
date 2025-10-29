# CloudServer REST API

This service provides a RESTful API for CRUD operations on a MySQL database, written in Go.

## Features
- CRUD endpoints for `items` table
- MySQL connection via environment variables
- OpenAPI/Swagger documentation auto-generated and served at `/swagger/index.html`
- Docker and Docker Compose support

## Usage

1. **Build and run with Docker Compose**  
   See `docker-compose.yml` for configuration.  
   Supports both local MySQL container and external MySQL instance.

2. **Access the API**  
   - Default: `http://localhost:8080`  
   - With Nginx: `https://api.devsecops2025-arubacloud.com` (if configured)

3. **View Swagger UI**  
   - `http://localhost:8080/swagger/index.html`

## Endpoints

- `GET /items` — List all items
- `GET /item?id=...` — Get one item
- `POST /item/create` — Create item
- `PUT /item/update?id=...` — Update item
- `DELETE /item/delete?id=...` — Delete item
- `GET /ping` — Health check

## Environment Variables

- `DB_HOST` — MySQL host (e.g., `mysql` for container, or external IP)
- `DB_USERNAME` — MySQL username
- `DB_PASSWORD` — MySQL password
- `DB_NAME` — MySQL database name

## Swagger/OpenAPI

- Documentation is auto-generated at build time.
- Served at `/swagger/index.html`.
- To update docs locally, install swag CLI and run `swag init`.

## Example Usage (with curl)

### Health Check
```sh
curl http://localhost:8080/ping
```

### List All Items
```sh
curl http://localhost:8080/items
```

### Get One Item
```sh
curl "http://localhost:8080/item?id=1"
```

### Create Item
```sh
curl -X POST http://localhost:8080/item/create \
  -H "Content-Type: application/json" \
  -d '{"name":"Sample Item","description":"A test item"}'
```

### Update Item
```sh
curl -X PUT "http://localhost:8080/item/update?id=1" \
  -H "Content-Type: application/json" \
  -d '{"name":"Updated Item","description":"Updated description"}'
```

### Delete Item
```sh
curl -X DELETE "http://localhost:8080/item/delete?id=1"
```

---

