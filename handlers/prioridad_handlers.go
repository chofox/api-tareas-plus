// handlers/prioridad_handlers.go

package handlers

import (
	"api-tareas-plus/database"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Manejador para obtener todas las prioridades
func GetPrioridades(w http.ResponseWriter, r *http.Request) {
	prioridades, err := database.GetPrioridades()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prioridades)
}

// Manejador para obtener una prioridad por ID
func GetPrioridad(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de prioridad no válido", http.StatusBadRequest)
		return
	}

	prioridad, err := database.GetPrioridadByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prioridad)
}

// Manejador para crear una nueva prioridad
func CreatePrioridad(w http.ResponseWriter, r *http.Request) {
	var prioridad database.Prioridad
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&prioridad); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := database.CreatePrioridad(prioridad.Nombre, prioridad.Descripcion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Manejador para actualizar una prioridad
func UpdatePrioridad(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de prioridad no válido", http.StatusBadRequest)
		return
	}

	var prioridad database.Prioridad
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&prioridad); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = database.UpdatePrioridad(id, prioridad.Nombre, prioridad.Descripcion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Manejador para eliminar una prioridad
func DeletePrioridad(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de prioridad no válido", http.StatusBadRequest)
		return
	}

	err = database.DeletePrioridad(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
