package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// PingHandler godoc
// @Summary Health check
// @Tags health
// @Success 200 {string} string "pong"
// @Router /ping [get]
func PingHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var result string
		err := db.QueryRow("SELECT 'pong' AS result").Scan(&result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "DB error: %v", err)
			return
		}
		fmt.Fprintf(w, result)
	}
}

// CreateItemHandler godoc
// @Summary Create a new item
// @Tags items
// @Accept json
// @Produce json
// @Param item body Item true "Item to create"
// @Success 200 {object} Item
// @Router /item/create [post]
func CreateItemHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var item Item
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid request body")
			return
		}
		res, err := db.Exec("INSERT INTO items (name, description) VALUES (?, ?)", item.Name, item.Description)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "DB error: %v", err)
			return
		}
		id, _ := res.LastInsertId()
		item.ID = int(id)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(item)
	}
}

// GetItemsHandler godoc
// @Summary List all items
// @Tags items
// @Produce json
// @Success 200 {array} Item
// @Router /items [get]
func GetItemsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, description FROM items")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "DB error: %v", err)
			return
		}
		defer rows.Close()
		items := make([]Item, 0) // Always return an array, even if empty
		for rows.Next() {
			var item Item
			if err := rows.Scan(&item.ID, &item.Name, &item.Description); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "DB error: %v", err)
				return
			}
			items = append(items, item)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
	}
}

// GetItemHandler godoc
// @Summary Get an item by ID
// @Tags items
// @Produce json
// @Param id query int true "Item ID"
// @Success 200 {object} Item
// @Failure 404 {string} string "Item not found"
// @Router /item [get]
func GetItemHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid id")
			return
		}
		var item Item
		err = db.QueryRow("SELECT id, name, description FROM items WHERE id = ?", id).Scan(&item.ID, &item.Name, &item.Description)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Item not found!")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(item)
	}
}

// UpdateItemHandler godoc
// @Summary Update an item
// @Tags items
// @Accept json
// @Produce json
// @Param id query int true "Item ID"
// @Param item body Item true "Updated item"
// @Success 200 {object} Item
// @Router /item/update [put]
func UpdateItemHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid id")
			return
		}
		var item Item
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid request body")
			return
		}
		_, err = db.Exec("UPDATE items SET name = ?, description = ? WHERE id = ?", item.Name, item.Description, id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "DB error: %v", err)
			return
		}
		item.ID = id
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(item)
	}
}

// DeleteItemHandler godoc
// @Summary Delete an item
// @Tags items
// @Param id query int true "Item ID"
// @Success 204 {string} string "No Content"
// @Router /item/delete [delete]
func DeleteItemHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid id")
			return
		}
		_, err = db.Exec("DELETE FROM items WHERE id = ?", id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "DB error: %v", err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
