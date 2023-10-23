// handlers/tarea_handlers.go

package handlers

import (
	"api-tareas-plus/database"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Manejador para obtener todas las tareas
func GetTareas(w http.ResponseWriter, r *http.Request) {
	tareas, err := database.GetTareas()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tareas)
}

// Manejador para obtener una tarea por ID
func GetTarea(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de tarea no válido", http.StatusBadRequest)
		return
	}

	tarea, err := database.GetTareaByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tarea)
}

// Manejador para crear una nueva tarea
func CreateTarea(w http.ResponseWriter, r *http.Request) {
	var tarea database.Tarea
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tarea); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := database.CreateTarea(tarea)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// Manejador para actualizar una tarea
func UpdateTarea(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de tarea no válido", http.StatusBadRequest)
		return
	}

	var tarea database.Tarea
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tarea); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = database.UpdateTarea(id, tarea)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Manejador para eliminar una tarea
func DeleteTarea(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de tarea no válido", http.StatusBadRequest)
		return
	}

	err = database.DeleteTarea(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
