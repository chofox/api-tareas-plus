package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

// Inicializar la conexión a la base de datos
func InitDB() {
	var err error
	db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/tareas_plus")
	if err != nil {
		panic(err)
	}
}

// Cerrar la conexión a la base de datos
func CloseDB() error {
	return db.Close()
}

type Categoria struct {
	ID     int
	Nombre string
	Color  string
}

type Prioridad struct {
	ID          int
	Nombre      string
	Descripcion string
}

type Tarea struct {
	ID               int
	Titulo           string
	Descripcion      string
	FechaVencimiento string
	Estado           string
	IDCategoria      int
	IDPrioridad      int
	IDUsuario        int
}

// Estructura para usuarios
type Usuario struct {
	ID                int
	Nombre            string
	CorreoElectronico string
	Contrasena        string
}

// Notificacion representa una notificación en la base de datos.
type Notificacion struct {
	ID                    int
	Descripcion           string
	FechaHoraNotificacion time.Time
	IDTarea               int
}

// HistorialTarea representa un registro en la tabla HistorialTareas
type HistorialTarea struct {
	ID              int
	AccionRealizada string
	FechaHoraAccion time.Time
	IDUsuario       int
	IDTarea         int
}

// GetCategorias obtiene todas las categorías
func GetCategorias() ([]Categoria, error) {
	rows, err := db.Query("SELECT id, nombre, color FROM categorias")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categorias []Categoria
	for rows.Next() {
		var categoria Categoria
		if err := rows.Scan(&categoria.ID, &categoria.Nombre, &categoria.Color); err != nil {
			return nil, err
		}
		categorias = append(categorias, categoria)
	}

	return categorias, nil
}

// GetCategoriaByID obtiene una categoría por su ID
func GetCategoriaByID(id int) (Categoria, error) {
	categoria := Categoria{}
	err := db.QueryRow("SELECT id, nombre, color FROM categorias WHERE id = ?", id).Scan(&categoria.ID, &categoria.Nombre, &categoria.Color)
	if err != nil {
		return Categoria{}, err
	}
	return categoria, nil
}

// CreateCategoria crea una nueva categoría
func CreateCategoria(nombre, color string) error {
	_, err := db.Exec("INSERT INTO categorias (nombre, color) VALUES (?, ?)", nombre, color)
	if err != nil {
		return err
	}
	return nil
}

// UpdateCategoria actualiza una categoría por su ID
func UpdateCategoria(id int, nombre, color string) error {
	_, err := db.Exec("UPDATE categorias SET nombre = ?, color = ? WHERE id = ?", nombre, color, id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCategoria elimina una categoría por su ID
func DeleteCategoria(id int) error {
	_, err := db.Exec("DELETE FROM categorias WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

// GetPrioridades obtiene todas las prioridades
func GetPrioridades() ([]Prioridad, error) {
	rows, err := db.Query("SELECT id, nombre, descripcion FROM prioridades")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prioridades []Prioridad
	for rows.Next() {
		var prioridad Prioridad
		if err := rows.Scan(&prioridad.ID, &prioridad.Nombre, &prioridad.Descripcion); err != nil {
			return nil, err
		}
		prioridades = append(prioridades, prioridad)
	}

	return prioridades, nil
}

// GetPrioridadByID obtiene una prioridad por su ID
func GetPrioridadByID(id int) (Prioridad, error) {
	prioridad := Prioridad{}
	err := db.QueryRow("SELECT id, nombre, descripcion FROM prioridades WHERE id = ?", id).Scan(&prioridad.ID, &prioridad.Nombre, &prioridad.Descripcion)
	if err != nil {
		return Prioridad{}, err
	}
	return prioridad, nil
}

// CreatePrioridad crea una nueva prioridad
func CreatePrioridad(nombre, descripcion string) error {
	_, err := db.Exec("INSERT INTO prioridades (nombre, descripcion) VALUES (?, ?)", nombre, descripcion)
	if err != nil {
		return err
	}
	return nil
}

// UpdatePrioridad actualiza una prioridad por su ID
func UpdatePrioridad(id int, nombre, descripcion string) error {
	_, err := db.Exec("UPDATE prioridades SET nombre = ?, descripcion = ? WHERE id = ?", nombre, descripcion, id)
	if err != nil {
		return err
	}
	return nil
}

// DeletePrioridad elimina una prioridad por su ID
func DeletePrioridad(id int) error {
	_, err := db.Exec("DELETE FROM prioridades WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

// GetTareas obtiene todas las tareas
func GetTareas() ([]Tarea, error) {
	rows, err := db.Query("SELECT id, titulo, descripcion, fecha_vencimiento, estado, id_categoria, id_prioridad, id_usuario FROM tareas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tareas []Tarea
	for rows.Next() {
		var tarea Tarea
		if err := rows.Scan(&tarea.ID, &tarea.Titulo, &tarea.Descripcion, &tarea.FechaVencimiento, &tarea.Estado, &tarea.IDCategoria, &tarea.IDPrioridad, &tarea.IDUsuario); err != nil {
			return nil, err
		}
		tareas = append(tareas, tarea)
	}

	return tareas, nil
}

// GetTareaByID obtiene una tarea por su ID
func GetTareaByID(id int) (Tarea, error) {
	tarea := Tarea{}
	err := db.QueryRow("SELECT id, titulo, descripcion, fecha_vencimiento, estado, id_categoria, id_prioridad, id_usuario FROM tareas WHERE id = ?", id).Scan(&tarea.ID, &tarea.Titulo, &tarea.Descripcion, &tarea.FechaVencimiento, &tarea.Estado, &tarea.IDCategoria, &tarea.IDPrioridad, &tarea.IDUsuario)
	if err != nil {
		return Tarea{}, err
	}
	return tarea, nil
}

// CreateTarea crea una nueva tarea
func CreateTarea(t Tarea) (int, error) {
	result, err := db.Exec("INSERT INTO tareas (titulo, descripcion, fecha_vencimiento, estado, id_categoria, id_prioridad, id_usuario) VALUES (?, ?, ?, ?, ?, ?, ?)",
		t.Titulo, t.Descripcion, t.FechaVencimiento, t.Estado, t.IDCategoria, t.IDPrioridad, t.IDUsuario)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	return int(id), nil
}

// UpdateTarea actualiza una tarea por su ID
func UpdateTarea(id int, t Tarea) error {
	_, err := db.Exec("UPDATE tareas SET titulo = ?, descripcion = ?, fecha_vencimiento = ?, estado = ?, id_categoria = ?, id_prioridad = ? WHERE id = ?", t.Titulo, t.Descripcion, t.FechaVencimiento, t.Estado, t.IDCategoria, t.IDPrioridad, id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTarea elimina una tarea por su ID
func DeleteTarea(id int) error {
	_, err := db.Exec("DELETE FROM tareas WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

// CreateUsuario crea un nuevo usuario
func CreateUsuario(nombre, correo, contrasena string) error {
	_, err := db.Exec("INSERT INTO Usuarios (Nombre, CorreoElectronico, Contrasena) VALUES (?, ?, ?)", nombre, correo, contrasena)
	if err != nil {
		return err
	}
	return nil
}

// GetUsuario obtiene un usuario por ID
func GetUsuario(id int) (Usuario, error) {
	u := Usuario{}
	row := db.QueryRow("SELECT ID, Nombre, CorreoElectronico, Contrasena FROM Usuarios WHERE ID = ?", id)
	err := row.Scan(&u.ID, &u.Nombre, &u.CorreoElectronico, &u.Contrasena)
	if err != nil {
		return Usuario{}, err
	}
	return u, nil
}

// UpdateUsuario actualiza un usuario
func UpdateUsuario(id int, nombre, correo, contrasena string) error {
	_, err := db.Exec("UPDATE Usuarios SET Nombre = ?, CorreoElectronico = ?, Contrasena = ? WHERE ID = ?", nombre, correo, contrasena, id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUsuario elimina un usuario por ID
func DeleteUsuario(id int) error {
	_, err := db.Exec("DELETE FROM Usuarios WHERE ID = ?", id)
	if err != nil {
		return err
	}
	return nil
}

// CreateNotificacion crea una nueva notificación.
func CreateNotificacion(descripcion string, fechaHoraNotificacion time.Time, idTarea int) error {
	_, err := db.Exec("INSERT INTO notificaciones (descripcion, fecha_hora_notificacion, id_tarea) VALUES (?, ?, ?)", descripcion, fechaHoraNotificacion, idTarea)
	if err != nil {
		return err
	}
	return nil
}

// GetNotificacionByID obtiene una notificación por su ID.
func GetNotificacionByID(id int) (Notificacion, error) {
	notificacion := Notificacion{}
	err := db.QueryRow("SELECT id, descripcion, fecha_hora_notificacion, id_tarea FROM notificaciones WHERE id = ?", id).
		Scan(&notificacion.ID, &notificacion.Descripcion, &notificacion.FechaHoraNotificacion, &notificacion.IDTarea)
	if err != nil {
		return Notificacion{}, err
	}
	return notificacion, nil
}

// GetNotificacionesPorUsuario obtiene todas las notificaciones de un usuario a través de sus tareas.
func GetNotificacionesPorUsuario(idUsuario int) ([]Notificacion, error) {
	var notificaciones []Notificacion
	rows, err := db.Query("SELECT n.id, n.descripcion, n.fecha_hora_notificacion, n.id_tarea FROM notificaciones n JOIN tareas t ON n.id_tarea = t.id WHERE t.id_usuario = ?", idUsuario)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var notificacion Notificacion
		err := rows.Scan(&notificacion.ID, &notificacion.Descripcion, &notificacion.FechaHoraNotificacion, &notificacion.IDTarea)
		if err != nil {
			return nil, err
		}
		notificaciones = append(notificaciones, notificacion)
	}
	return notificaciones, nil
}

// GetNotificacion obtiene una notificación por ID
func GetNotificacion(id int) (Notificacion, error) {
	notificacion := Notificacion{}
	err := db.QueryRow("SELECT id, descripcion, fecha_hora_notificacion, id_tarea FROM notificaciones WHERE id = ?", id).
		Scan(&notificacion.ID, &notificacion.Descripcion, &notificacion.FechaHoraNotificacion, &notificacion.IDTarea)
	if err != nil {
		return Notificacion{}, err
	}
	return notificacion, nil
}

// UpdateNotificacion actualiza una notificación existente
func UpdateNotificacion(id int, descripcion string, fechaHoraNotificacion time.Time, idTarea int) error {
	_, err := db.Exec("UPDATE notificaciones SET descripcion = ?, fecha_hora_notificacion = ?, id_tarea = ? WHERE id = ?", descripcion, fechaHoraNotificacion, idTarea, id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteNotificacion elimina una notificación por su ID
func DeleteNotificacion(id int) error {
	_, err := db.Exec("DELETE FROM notificaciones WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

// CreateHistorialTarea crea un nuevo registro en el historial de tareas
func CreateHistorialTarea(accionRealizada string, fechaHoraAccion time.Time, idUsuario, idTarea int) error {
	_, err := db.Exec("INSERT INTO HistorialTareas (AccionRealizada, FechaHoraAccion, ID_Usuario, ID_Tarea) VALUES (?, ?, ?, ?)", accionRealizada, fechaHoraAccion, idUsuario, idTarea)
	if err != nil {
		return err
	}
	return nil
}

// GetHistorialTareaByID obtiene un registro de historial de tareas por su ID
func GetHistorialTareaByID(id int) (HistorialTarea, error) {
	historial := HistorialTarea{}
	err := db.QueryRow("SELECT ID, AccionRealizada, FechaHoraAccion, ID_Usuario, ID_Tarea FROM HistorialTareas WHERE ID = ?", id).
		Scan(&historial.ID, &historial.AccionRealizada, &historial.FechaHoraAccion, &historial.IDUsuario, &historial.IDTarea)
	if err != nil {
		return HistorialTarea{}, err
	}
	return historial, nil
}

// GetHistorialTareasByTareaID obtiene todos los registros de historial de tareas relacionados con una tarea por su ID
func GetHistorialTareasByTareaID(idTarea int) ([]HistorialTarea, error) {
	rows, err := db.Query("SELECT ID, AccionRealizada, FechaHoraAccion, ID_Usuario, ID_Tarea FROM HistorialTareas WHERE ID_Tarea = ?", idTarea)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var historiales []HistorialTarea
	for rows.Next() {
		var historial HistorialTarea
		if err := rows.Scan(&historial.ID, &historial.AccionRealizada, &historial.FechaHoraAccion, &historial.IDUsuario, &historial.IDTarea); err != nil {
			return nil, err
		}
		historiales = append(historiales, historial)
	}

	return historiales, nil
}

// UpdateHistorialTarea actualiza un registro de historial de tareas por su ID
func UpdateHistorialTarea(id int, accionRealizada string, fechaHoraAccion time.Time, idUsuario, idTarea int) error {
	_, err := db.Exec("UPDATE HistorialTareas SET AccionRealizada = ?, FechaHoraAccion = ?, ID_Usuario = ?, ID_Tarea = ? WHERE ID = ?", accionRealizada, fechaHoraAccion, idUsuario, idTarea, id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteHistorialTarea elimina un registro de historial de tareas por su ID
func DeleteHistorialTarea(id int) error {
	_, err := db.Exec("DELETE FROM HistorialTareas WHERE ID = ?", id)
	if err != nil {
		return err
	}
	return nil
}

// Estructura para la solicitud de inicio de sesión
type LoginRequest struct {
	CorreoElectronico string `json:"correoElectronico"`
	Contrasena        string `json:"contrasena"`
}

// Obtener un usuario por su correo electrónico
func GetUserByCorreoElectronico(correoElectronico string) (Usuario, error) {
	user := Usuario{}
	row := db.QueryRow("SELECT ID, Nombre, CorreoElectronico, Contrasena FROM Usuarios WHERE CorreoElectronico = ?", correoElectronico)
	err := row.Scan(&user.ID, &user.Nombre, &user.CorreoElectronico, &user.Contrasena)
	if err != nil {
		return Usuario{}, err
	}
	return user, nil
}

// Verificar las credenciales del usuario
func VerifyCredentials(correoElectronico, contrasena string) (bool, int, error) {
	user := Usuario{}
	row := db.QueryRow("SELECT ID, Nombre, CorreoElectronico, Contrasena FROM Usuarios WHERE CorreoElectronico = ?", correoElectronico)
	err := row.Scan(&user.ID, &user.Nombre, &user.CorreoElectronico, &user.Contrasena)
	if err != nil {
		return false, 0, err
	}

	// Compara la contraseña ingresada con el hash almacenado
	err = bcrypt.CompareHashAndPassword([]byte(user.Contrasena), []byte(contrasena))
	if err != nil {
		return false, 0, nil
	}

	return true, user.ID, nil
}
