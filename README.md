# 🛡️ SSL Checker Pro

Una herramienta profesional para analizar el estado de seguridad y certificados SSL de dominios web. Desarrollada en **Go** implementando una arquitectura modular que comparte la lógica de negocio entre una **CLI** (Línea de comandos) y una **GUI** (Interfaz Gráfica) moderna construida con **Wails** y **Svelte**.

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![Wails](https://img.shields.io/badge/Wails-v2-red?style=flat&logo=wails)
![Svelte](https://img.shields.io/badge/Frontend-Svelte-orange?style=flat&logo=svelte)

## 📋 Prerrequisitos

Antes de ejecutar el proyecto, asegúrate de tener instaladas las siguientes herramientas en tu sistema:

1.  **Go** (v1.21 o superior): [Descargar Go](https://go.dev/dl/)
2.  **Node.js & npm** (Necesario para compilar el frontend): [Descargar Node.js](https://nodejs.org/)
3.  **Wails CLI** (Herramienta para empaquetar la GUI):
    ```bash
    go install [github.com/wailsapp/wails/v2/cmd/wails@latest](https://github.com/wailsapp/wails/v2/cmd/wails@latest)
    ```

Para verificar que tienes todo listo, puedes ejecutar:
```bash
wails doctor
```
## 🎨 Stack Frontend & UI

La interfaz gráfica (`gui/frontend`) no es una simple página web incrustada, es una aplicación reactiva moderna optimizada para escritorio.

* **Framework:** [Svelte](https://svelte.dev/) (Versión JS). Elegido por su ligereza y ausencia de "Virtual DOM", lo que lo hace extremadamente rápido.
* **Build Tool:** [Vite](https://vitejs.dev/). Proporciona tiempos de arranque casi instantáneos y **HMR (Hot Module Replacement)**, permitiendo ver cambios de diseño en tiempo real sin recargar la app.
* **Plantilla Base:** Wails Svelte Template (Vanilla JavaScript).
* **Comunicación:** Utiliza el runtime de Wails para invocar funciones de Go como si fueran promesas de JavaScript nativas.

## Configurar dependencias del Backend (Raíz)
Descarga las librerías necesarias para el módulo principal y la CLI.
```bash
go mod tidy
```
## Configurar dependencias de la GUI
La interfaz gráfica funciona como un sub-módulo que depende del módulo raíz.
```bash
cd gui
go mod tidy
cd frontend
npm install
cd ../..
```
## Cómo ejecutar?
### Opción A: Interfaz gráfica (GUI)
Ideal para uso interactivo. Incluye barra de progreso, validación visual de colores y sección educativa.

Navega a la carpeta de la GUI e inicia el modo de desarrollo (esto abrirá una ventana nativa):
```bash
cd gui
wails dev
```
### Opción B: Línea de Comandos (CLI)
Ideal para servidores, scripts automatizados o diagnósticos rápidos.

Desde la raíz del proyecto ejecuta:
```bash
go run ./cmd/cli -d google.com
```
## Arquitectura del proyecto:
Este proyecto sigue el principio DRY (Don't Repeat Yourself) mediante una estructura de paquetes organizada:

ssl-checker/
├── cmd/
│   └── cli/           # 🖥️ Entrypoint: Versión de Terminal (CLI)
│                      # Consume la lógica desde pkg/ssl
│
├── gui/               # 🎨 Entrypoint: Versión Gráfica (Wails)
│   ├── app.go         # Puente (Bridge) entre Go y JavaScript
│   ├── go.mod         # Módulo independiente. Usa 'replace' para leer la raíz
│   └── frontend/      # Código fuente de Svelte (Interfaz de Usuario)
│
├── pkg/               # 🧠 LÓGICA COMPARTIDA (Library)
│   └── ssl/           # Aquí vive el Scanner, Cliente HTTP y Modelos.
│                      # Es accesible tanto por 'cmd' como por 'gui'.
│
└── go.mod             # Definición del módulo raíz

## Compilación (Build)
Para generar los archivos ejecutables finales (.exe / .app) para distribuir la aplicación:

### Para la GUI:
```bash
cd gui
wails build
```
(El ejecutable se generará en gui/build/bin)

### Para la CLI:
```bash
go build -o ssl-checker.exe ./cmd/cli
```