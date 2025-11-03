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

3. **View Swagger UI**  
   - `http://localhost:8080/swagger/index.html`

## Local Development

For local development with live code reload and a local MySQL instance:

1. **Create `docker-compose.override.yml`**:
   ```yaml
   version: '3.8'
   services:
     restapi:
       build: .
       volumes:
         - .:/app
       environment:
         - DB_HOST=mysql
         - DB_USERNAME=root
         - DB_PASSWORD=rootpass
         - DB_NAME=restapi
       depends_on:
         - mysql

     mysql:
       image: mysql:8.0
       environment:
         MYSQL_ROOT_PASSWORD: rootpass
         MYSQL_DATABASE: restapi
       ports:
         - "3306:3306"
       volumes:
         - mysql_data:/var/lib/mysql

   volumes:
     mysql_data:
   ```

2. **Start local environment**:
   ```sh
   docker-compose up -d
   ```

   Docker Compose automatically merges `docker-compose.yml` and `docker-compose.override.yml`.

3. **View logs**:
   ```sh
   docker-compose logs -f restapi
   ```

4. **Stop local environment**:
   ```sh
   docker-compose down
   ```

## Endpoints

- `GET /items` — List all items
- `GET /item?id=...` — Get one item
- `POST /item/create` — Create item
- `PUT /item/update?id=...` — Update item
- `DELETE /item/delete?id=...` — Delete item
- `GET /ping` — Health check

## Environment Variables

- `DB_HOST` — MySQL host 
- `DB_USERNAME` — MySQL username
- `DB_PASSWORD` — MySQL password
- `DB_NAME` — MySQL database name

## Swagger/OpenAPI

- Documentation is auto-generated at build time.
- Served at `/swagger/index.html`.
- To update docs locally, install swag CLI and run `swag init`.

## Example Usage

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
