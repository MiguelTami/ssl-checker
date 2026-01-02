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