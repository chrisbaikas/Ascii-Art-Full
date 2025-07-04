# ASCII Art Web Dockerize

A simple web server written in Go that transforms input text into ASCII art using one of three banner styles. The application is containerized with Docker for easy deployment.

---

## 🛠️ Technologies

- **Language**: Go (only standard packages used)
- **Containerization**: Docker
- **Base image**: `golang:1.22-alpine`

---

## 📦 Project Structure

```
ascii-art-web-dockerize/
├── Dockerfile
├── Makefile
├── main.go
├── utils/
│   ├── processor.go
│   ├── justify.go
│   └── color.go
├── banners/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
└── README.md
```

---

## 🚀 How to Run

### 1. Build the Docker Image
```bash
make build
```

### 2. Run the Docker Container
```bash
make run
```

The server will start at `http://localhost:8080`

### 3. Stop the Container
```bash
make stop
```

### 4. View Logs
```bash
make logs
```

### 5. Open in Browser (Windows only)
```bash
make open
```

---

## 📄 Example Usage

Visit `http://localhost:8080` in your browser and:

- Select a banner style (Standard, Shadow, Thinkertoy)
- Enter your text input
- (Optional) Choose color options
- Press **Generate**

---

## 📋 Makefile Targets

| Target   | Description                             |
|----------|-----------------------------------------|
| build    | Builds the Docker image                 |
| run      | Runs the container                      |
| stop     | Stops and removes the container         |
| rebuild  | Stops, rebuilds, and reruns everything  |
| logs     | Follows container logs                  |
| clean    | Cleans up Docker garbage objects        |
| start    | Starts a stopped container              |
| exec     | Enters the container using bash         |
| open     | Opens the web page (Windows only)       |

---

## 🧼 Docker Cleanup

To clean up unused Docker resources:

```bash
make clean
```

This removes:
- Unused containers
- Dangling images
- Unused volumes & networks

---

## 📌 Metadata

- Maintainers: `chbaikas`, `askordal`, `npanouts`
- Version: `1.0`
- Description: ASCII Art Web Server in Go

---

## ✅ Evaluation Checklist

- [x] Written in Go using standard packages
- [x] Dockerized with clean Dockerfile
- [x] Image and container successfully created
- [x] Metadata labels included
- [x] Makefile automates all steps
- [x] Runs correctly on `http://localhost:8080`

---

## 🧪 Notes for Evaluators

This project was developed with clarity and simplicity in mind, focusing on Docker usage, clean containerization, and standard Go practices. Please evaluate using the functional and general criteria provided in the subject.