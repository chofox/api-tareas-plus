// handlers/usuario_handlers.go
package handlers

import (
	"api-tareas-plus/database"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Manejador para crear un nuevo usuario
func CreateUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario database.Usuario
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&usuario); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := database.CreateUsuario(usuario.Nombre, usuario.CorreoElectronico, usuario.Contrasena)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Manejador para obtener un usuario por ID
func GetUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de usuario no válido", http.StatusBadRequest)
		return
	}

	usuario, err := database.GetUsuario(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuario)
}

// Manejador para actualizar un usuario
func UpdateUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de usuario no válido", http.StatusBadRequest)
		return
	}

	var usuario database.Usuario
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&usuario); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = database.UpdateUsuario(id, usuario.Nombre, usuario.CorreoElectronico, usuario.Contrasena)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Manejador para eliminar un usuario
func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de usuario no válido", http.StatusBadRequest)
		return
	}

	err = database.DeleteUsuario(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
