package handlers

import (
	"api-tareas-plus/database"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Manejador para crear un registro en el historial de tareas
func CreateHistorialTarea(w http.ResponseWriter, r *http.Request) {
	var historial database.HistorialTarea
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&historial); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := database.CreateHistorialTarea(historial.AccionRealizada, historial.FechaHoraAccion, historial.IDUsuario, historial.IDTarea)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Manejador para obtener un registro de historial de tareas por su ID
func GetHistorialTarea(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de historial de tarea no válido", http.StatusBadRequest)
		return
	}

	historial, err := database.GetHistorialTareaByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(historial)
}

// Manejador para obtener todos los registros de historial de tareas relacionados con una tarea por su ID
func GetHistorialTareasByTareaID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idTareaStr := params["idTarea"]
	idTarea, err := strconv.Atoi(idTareaStr)
	if err != nil {
		http.Error(w, "ID de tarea no válido", http.StatusBadRequest)
		return
	}

	historiales, err := database.GetHistorialTareasByTareaID(idTarea)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(historiales)
}

// Manejador para actualizar un registro de historial de tareas
func UpdateHistorialTarea(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de historial de tarea no válido", http.StatusBadRequest)
		return
	}

	var historial database.HistorialTarea
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&historial); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = database.UpdateHistorialTarea(id, historial.AccionRealizada, historial.FechaHoraAccion, historial.IDUsuario, historial.IDTarea)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Manejador para eliminar un registro de historial de tareas por su ID
func DeleteHistorialTarea(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de historial de tarea no válido", http.StatusBadRequest)
		return
	}

	err = database.DeleteHistorialTarea(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
