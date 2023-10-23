package main

import (
	"api-tareas-plus/database"
	"api-tareas-plus/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Inicializar la conexión a la base de datos
	database.InitDB()
	defer database.CloseDB() // Asegura que la conexión a la base de datos se cierre al final

	r := mux.NewRouter()

	// Manejadores para Categorías
	r.HandleFunc("/categorias", handlers.GetCategorias).Methods("GET")
	r.HandleFunc("/categorias/{id:[0-9]+}", handlers.GetCategoria).Methods("GET")
	r.HandleFunc("/categorias", handlers.CreateCategoria).Methods("POST")
	r.HandleFunc("/categorias/{id:[0-9]+}", handlers.UpdateCategoria).Methods("PUT")
	r.HandleFunc("/categorias/{id:[0-9]+}", handlers.DeleteCategoria).Methods("DELETE")

	// Manejadores para Prioridades
	r.HandleFunc("/prioridades", handlers.GetPrioridades).Methods("GET")
	r.HandleFunc("/prioridades/{id:[0-9]+}", handlers.GetPrioridad).Methods("GET")
	r.HandleFunc("/prioridades", handlers.CreatePrioridad).Methods("POST")
	r.HandleFunc("/prioridades/{id:[0-9]+}", handlers.UpdatePrioridad).Methods("PUT")
	r.HandleFunc("/prioridades/{id:[0-9]+}", handlers.DeletePrioridad).Methods("DELETE")

	// Manejadores para Tareas
	r.HandleFunc("/tareas", handlers.GetTareas).Methods("GET")
	r.HandleFunc("/tareas/{id:[0-9]+}", handlers.GetTarea).Methods("GET")
	r.HandleFunc("/tareas", handlers.CreateTarea).Methods("POST")
	r.HandleFunc("/tareas/{id:[0-9]+}", handlers.UpdateTarea).Methods("PUT")
	r.HandleFunc("/tareas/{id:[0-9]+}", handlers.DeleteTarea).Methods("DELETE")

	// Rutas para usuarios
	r.HandleFunc("/usuarios", handlers.CreateUsuario).Methods("POST")
	r.HandleFunc("/usuarios/{id:[0-9]+}", handlers.GetUsuario).Methods("GET")
	r.HandleFunc("/usuarios/{id:[0-9]+}", handlers.UpdateUsuario).Methods("PUT")
	r.HandleFunc("/usuarios/{id:[0-9]+}", handlers.DeleteUsuario).Methods("DELETE")

	// Rutas para las notificaciones
	r.HandleFunc("/notificaciones/usuario/{idUsuario:[0-9]+}", handlers.GetNotificacionesPorUsuario).Methods("GET")
	r.HandleFunc("/notificaciones", handlers.CreateNotificacion).Methods("POST")
	r.HandleFunc("/notificaciones/{id:[0-9]+}", handlers.GetNotificacion).Methods("GET")
	r.HandleFunc("/notificaciones/{id:[0-9]+}", handlers.UpdateNotificacion).Methods("PUT")
	r.HandleFunc("/notificaciones/{id:[0-9]+}", handlers.DeleteNotificacion).Methods("DELETE")

	// Rutas para historial de tareas
	r.HandleFunc("/historialtareas", handlers.CreateHistorialTarea).Methods("POST")
	r.HandleFunc("/historialtareas/{id:[0-9]+}", handlers.GetHistorialTarea).Methods("GET")
	r.HandleFunc("/historialtareas/tarea/{idTarea:[0-9]+}", handlers.GetHistorialTareasByTareaID).Methods("GET")
	r.HandleFunc("/historialtareas/{id:[0-9]+}", handlers.UpdateHistorialTarea).Methods("PUT")
	r.HandleFunc("/historialtareas/{id:[0-9]+}", handlers.DeleteHistorialTarea).Methods("DELETE")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
