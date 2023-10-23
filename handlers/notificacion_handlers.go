package handlers

import (
	"api-tareas-plus/database"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Manejador para obtener todas las notificaciones por usuario
func GetNotificacionesPorUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idUsuarioStr := params["idUsuario"]
	idUsuario, err := strconv.Atoi(idUsuarioStr)
	if err != nil {
		http.Error(w, "ID de usuario no válido", http.StatusBadRequest)
		return
	}

	notificaciones, err := database.GetNotificacionesPorUsuario(idUsuario)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notificaciones)
}

// Manejador para crear una nueva notificación
func CreateNotificacion(w http.ResponseWriter, r *http.Request) {
	var notificacion database.Notificacion
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&notificacion); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := database.CreateNotificacion(notificacion.Descripcion, notificacion.FechaHoraNotificacion, notificacion.IDTarea)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Manejador para obtener una notificación por ID
func GetNotificacion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de notificación no válido", http.StatusBadRequest)
		return
	}

	notificacion, err := database.GetNotificacion(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notificacion)
}

// Manejador para actualizar una notificación
func UpdateNotificacion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de notificación no válido", http.StatusBadRequest)
		return
	}

	var notificacion database.Notificacion
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&notificacion); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = database.UpdateNotificacion(id, notificacion.Descripcion, notificacion.FechaHoraNotificacion, notificacion.IDTarea)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Manejador para eliminar una notificación
func DeleteNotificacion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de notificación no válido", http.StatusBadRequest)
		return
	}

	err = database.DeleteNotificacion(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
