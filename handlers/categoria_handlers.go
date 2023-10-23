// handlers/categoria_handlers.go

package handlers

import (
	"api-tareas-plus/database"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Manejador para obtener todas las categorías
func GetCategorias(w http.ResponseWriter, r *http.Request) {
	categorias, err := database.GetCategorias()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categorias)
}

// Manejador para obtener una categoría por ID
func GetCategoria(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de categoría no válido", http.StatusBadRequest)
		return
	}

	categoria, err := database.GetCategoriaByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categoria)
}

// Manejador para crear una nueva categoría
func CreateCategoria(w http.ResponseWriter, r *http.Request) {
	var categoria database.Categoria
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&categoria); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := database.CreateCategoria(categoria.Nombre, categoria.Color)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Manejador para actualizar una categoría
func UpdateCategoria(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de categoría no válido", http.StatusBadRequest)
		return
	}

	var categoria database.Categoria
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&categoria); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = database.UpdateCategoria(id, categoria.Nombre, categoria.Color)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Manejador para eliminar una categoría
func DeleteCategoria(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de categoría no válido", http.StatusBadRequest)
		return
	}

	err = database.DeleteCategoria(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
