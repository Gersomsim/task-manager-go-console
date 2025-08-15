# 📋 Gestor de Tareas (Task Manager)

![coverage](https://raw.githubusercontent.com/Gersomsim/task-manager-go-console/refs/heads/badges/.badges/main/coverage.svg)

Un gestor de tareas simple y eficiente desarrollado en Go que te permite administrar tus tareas diarias desde la línea de comandos.

## ✨ Características

- **💾 Agregar tareas**: Crea nuevas tareas con título y descripción
- **📝 Listar tareas**: Visualiza todas tus tareas en una tabla organizada
- **✅ Completar tareas**: Marca tareas como completadas
- **💾 Persistencia**: Las tareas se guardan automáticamente en formato JSON
- **🔄 Interfaz intuitiva**: Menú claro y fácil de usar
- **⚡ Rendimiento**: Aplicación rápida y eficiente en Go

## 🚀 Instalación

### Prerrequisitos

- Go 1.25.0 o superior
- Terminal o línea de comandos

### Pasos de instalación

1. **Clona el repositorio**
   ```bash
   git clone https://github.com/Gersomsim/task-manager-go-console
   cd task-manager
   ```

2. **Ejecuta la aplicación**
   ```bash
   go run main.go
   ```

## 🎯 Uso

Una vez ejecutada la aplicación, verás un menú con las siguientes opciones:

```
--------------------------------
 
Seleccione una opción:
1. 💾 Agregar tarea
2. 📝 Listar tareas
3. ✅ Marcar tarea como completada
4. 🚪 Salir
 
--------------------------------
```

### Agregar una tarea
- Selecciona la opción **1**
- Ingresa el título de la tarea
- Ingresa la descripción
- La tarea se guardará automáticamente

### Listar tareas
- Selecciona la opción **2**
- Verás todas tus tareas en una tabla organizada
- Presiona Enter para volver al menú principal

### Completar una tarea
- Selecciona la opción **3**
- Ingresa el ID de la tarea a completar
- Confirma la acción
- La tarea se marcará como completada

## 🏗️ Estructura del Proyecto

```
task-manager/
├── main.go              # Punto de entrada de la aplicación
├── go.mod               # Dependencias del proyecto
├── menu/
│   └── menu.go         # Lógica del menú principal
├── task/
│   ├── task.go         # Definición de la estructura Task
│   ├── handler.go      # Manejo de operaciones de tareas
│   └── storage.go      # Persistencia de datos en JSON
├── storage/
│   └── tasks.json      # Archivo de almacenamiento de tareas
└── README.md           # Este archivo
```

## 🔧 Tecnologías Utilizadas

- **Go 1.25.0**: Lenguaje de programación principal
- **encoding/json**: Para serialización/deserialización de datos
- **os**: Para operaciones del sistema de archivos
- **bufio**: Para lectura de entrada del usuario
- **text/tabwriter**: Para formateo de tablas en consola

## 📊 Estructura de Datos

Cada tarea se representa con la siguiente estructura:

```go
type Task struct {
    Id          int       // Identificador único
    Title       string    // Título de la tarea
    Description string    // Descripción detallada
    Completed   bool      // Estado de completado
    CreatedAt   time.Time // Fecha de creación
    UpdatedAt   time.Time // Fecha de última modificación
}
```

## 💾 Almacenamiento

- Las tareas se guardan automáticamente en `storage/tasks.json`
- El archivo se crea automáticamente si no existe
- Los datos persisten entre ejecuciones del programa
- Formato JSON legible y fácil de editar manualmente

## 🎨 Características de la Interfaz

- **Limpieza de pantalla**: Interfaz clara y organizada
- **Emojis**: Iconos visuales para mejor experiencia de usuario
- **Tablas formateadas**: Presentación ordenada de las tareas
- **Mensajes de confirmación**: Feedback claro para cada acción
- **Validación de entrada**: Verificación de datos ingresados

## 🚀 Desarrollo

### Ejecutar en modo desarrollo
```bash
go run main.go
```

### Compilar la aplicación
```bash
go build -o task-manager
```

### Ejecutar el binario compilado
```bash
./task-manager
```

## 🤝 Contribuir

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 📝 Licencia

Este proyecto está bajo la Licencia MIT. Ver el archivo `LICENSE` para más detalles.

## 🙏 Agradecimientos

- Desarrollado como proyecto de aprendizaje de Go
- Inspirado en la necesidad de gestionar tareas de manera eficiente
- Utiliza las mejores prácticas de Go para aplicaciones de línea de comandos

---

**¡Disfruta organizando tus tareas con este gestor simple y eficiente! 🎉**
