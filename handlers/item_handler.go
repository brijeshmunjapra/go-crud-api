package handlers

import (
	"crud-api/config"
	"crud-api/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item

	// Decode the incoming JSON request body into the item struct
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Insert the item into the database
	query := `INSERT INTO items (name, price) VALUES ($1, $2) RETURNING id`
	err = config.DB.QueryRow(query, item.Name, item.Price).Scan(&item.ID)
	if err != nil {

		http.Error(w, "Error saving item", http.StatusInternalServerError)
		return
	}

	// Respond with the created item
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id, name, price FROM items")
	if err != nil {
		http.Error(w, "Unable to fetch items", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.Name, &item.Price)
		if err != nil {
			http.Error(w, "Error scanning items", http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	var item models.Item
	err = config.DB.QueryRow("SELECT id, name, price FROM items WHERE id = $1", id).Scan(&item.ID, &item.Name, &item.Price)
	if err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	var updatedItem models.Item
	_ = json.NewDecoder(r.Body).Decode(&updatedItem)

	_, err = config.DB.Exec("UPDATE items SET name = $1, price = $2 WHERE id = $3",
		updatedItem.Name, updatedItem.Price, id)
	if err != nil {
		http.Error(w, "Unable to update item", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedItem)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	_, err = config.DB.Exec("DELETE FROM items WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Unable to delete item", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Item deleted"})
}
