
# rsync-web-ui

![Badge](https://img.shields.io/badge/Go-1.20%2B-blue)
![Badge](https://img.shields.io/badge/License-MIT-green)
![Badge](https://img.shields.io/badge/Status-Prototype-yellow)

Interfaz web simple en Go para configurar y generar comandos `rsync` entre máquinas dentro de la red local. La aplicación permite seleccionar directorios locales, un usuario remoto y una ruta de destino, generando comandos listos para sincronizar.

## ✨ Funcionalidades

* 📂 Añade múltiples rutas locales a sincronizar.
* 🧠 Recuerda rutas previamente guardadas.
* 🧪 Verifica existencia del directorio y sugiere comandos de comprobación.
* 🖥️ Detecta IP local automáticamente.
* 👥 Muestra usuarios en `/home` como sugerencias para uso remoto.
* ⚙️ Genera comandos `rsync` basados en entrada y configuración.
* 🌐 Interfaz web mínima corriendo en `localhost:3001`.

## ⚙️ Uso

1. Ejecuta la aplicación:

```bash
go run main.go
```

2. Abre tu navegador en:

```
http://localhost:3001
```

3. En la web podrás:

   * Ingresar directorios locales que deseas sincronizar.
   * Especificar el usuario remoto.
   * Establecer el directorio raíz remoto.
   * Obtener comandos `rsync` listos para copiar/pegar.

## 📁 Archivos

* `config`: contiene el path remoto de destino.
* `log`: lista los directorios locales agregados anteriormente.
* `index.html`: archivo de plantilla HTML que renderiza la interfaz.
* `main.go`: lógica del servidor y manejo de datos.

## 💡 Ejemplo de comando generado

```bash
rsync -t /home/pi/Documents user@192.168.1.100:/home/user/syncdir
```

> Asegúrate de que los directorios existen y que el usuario remoto tenga permisos para acceder al destino.

## 🚧 Estado

Este proyecto está en una etapa prototipo. No realiza la sincronización directamente por seguridad. Solo genera los comandos `rsync`, que puedes ejecutar manualmente.
