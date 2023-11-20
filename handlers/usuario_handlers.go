package handlers

import (
	"api-tareas-plus/auth"
	"api-tareas-plus/database"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

// Manejador para crear un nuevo usuario
func CreateUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario database.Usuario
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&usuario); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Verifica que el correo electrónico no esté en uso
	existingUser, err := database.GetUserByCorreoElectronico(usuario.CorreoElectronico)
	if err == nil && existingUser != (database.Usuario{}) {
		http.Error(w, "El correo electrónico ya está en uso", http.StatusBadRequest)
		return
	}

	// Genera un hash para la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usuario.Contrasena), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Almacena el usuario en la base de datos con la contraseña hash
	usuario.Contrasena = string(hashedPassword)
	err = database.CreateUsuario(usuario.Nombre, usuario.CorreoElectronico, usuario.Contrasena)
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

// Manejador para registrar un nuevo usuario
func RegisterUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario database.Usuario
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&usuario); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Verifica que el correo electrónico no esté en uso
	existingUser, err := database.GetUserByCorreoElectronico(usuario.CorreoElectronico)
	if err == nil && existingUser != (database.Usuario{}) {
		http.Error(w, "El correo electrónico ya está en uso", http.StatusBadRequest)
		return
	}

	// Genera un hash para la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usuario.Contrasena), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Almacena el usuario en la base de datos con la contraseña hash
	usuario.Contrasena = string(hashedPassword)
	err = database.CreateUsuario(usuario.Nombre, usuario.CorreoElectronico, usuario.Contrasena)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Manejador para iniciar sesión
func LoginUsuario(w http.ResponseWriter, r *http.Request) {
	var loginRequest database.LoginRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&loginRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Error al decodificar JSON: %v", err)
		return
	}

	// Verifica las credenciales y obtiene el ID del usuario si son válidas
	valid, userID, err := database.VerifyCredentials(loginRequest.CorreoElectronico, loginRequest.Contrasena)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !valid {
		http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
		return
	}

	// Genera un token de autenticación
	token, err := auth.GenerateAuthToken(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Envía el token de autenticación en la respuesta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
