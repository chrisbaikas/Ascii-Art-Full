# ASCII Art Suite — Terminal & Web

A complete ASCII art experience in Go:  
- 🎨 Terminal tool for CLI users  
- 🌐 Web-based interface for interactive use

---

## 🧰 Project 1: ascii-art-terminal

A command-line tool written in Go that converts text to stylized ASCII art with support for alignment, color, output to file, and reverse decoding.

### 👤 Authors

- chbaikas  
- askordal  
- nektariospan  

### ✨ Features

- ✅ Text to ASCII Art conversion
- 🎨 Color highlighting (full or partial text)
- 📐 Alignment options: left, center, right, justify
- 📤 Output to file or terminal
- 🔁 Reverse: ASCII Art → Text
- 🧱 Custom banner fonts
- 📦 No external dependencies (pure Go)

### 🚀 Usage

```bash
go run . [--reverse=<input.txt>] [--color=<color>[:substring]] [--align=left|center|right|justify] [--output=file.txt] "text" [banner]
```

#### Examples:

- Basic:
  ```bash
  go run . "Hello"
  go run . "Hi\nThere" shadow
  ```

- Color:
  ```bash
  go run . --color=red "Alert"
  go run . --color=blue:Go --color=green:Lang "Go Lang"
  ```

- Align:
  ```bash
  go run . --align=justify "Text here"
  ```

- Reverse:
  ```bash
  go run . --reverse=example.txt
  ```

### 🎨 Color Options

- Named: `red`, `blue`, `green`, etc.
- RGB: `'rgb(255,0,0)'`
- HEX: `'#ff0000'`
- HSL: `'hsl(0,100%,50%)'`

### 🧪 Tests

```bash
go test ./utils
```

Covers alignment, coloring, ASCII parsing, and reverse logic.

---

## 🌍 Project 2: ascii-art-web

A modern web-based GUI that converts text to ASCII art. Built using a Go web server and a responsive HTML/CSS/JS frontend.

### 👤 Authors

- chbaikas  
- askordal  
- npanouts  

### ✨ Features

- ✅ 3 banner styles: `standard`, `shadow`, `thinkertoy`
- 🎨 Color highlighting (targeted + global)
- 📐 Left/right alignment
- 🧱 Responsive mobile/tablet layout
- 🧑‍🎨 Background color support
- 🔄 Live preview (debounce)
- 💾 Export to `.txt`, `.html`, `.json`, `.svg`
- 🔧 Built using only Go standard library

### 🚀 Getting Started

```bash
git clone <repository-url>
cd ascii-art-web
go run .
```

Then open [http://localhost:8080](http://localhost:8080)

### 📁 Structure

```
ascii-art-web/
├── banners/             # Banner fonts
├── templates/           # HTML templates
│   └── index.html
├── static/              # CSS and JS
│   ├── style.css
│   ├── main.js
│   ├── color.js
│   ├── dropdown.js
│   ├── generate.js
│   └── export.js
├── web/                 # Go handlers
│   └── handlers.go
├── go.mod
└── go.sum
```

### 🌐 Endpoints

- `GET /` → Main UI
- `POST /ascii-art` → Generate ASCII
- `POST /export` → Download file
- `GET /ascii-table` → View ASCII table

### 🧪 Export Test Flow

1. Enter text
2. Choose colors, alignment, etc.
3. View preview
4. Click export
5. Download output in desired format

---

## 📝 License

This project is licensed under the [MIT License](LICENSE).

---

## 🧠 Learning Outcomes

- Go backends (standard `net/http`)
- HTML form processing
- Frontend feedback (JS debounce, color pickers)
- File streaming + MIME headers
- Clean, responsive UI/UX

---

## 🗣 Feedback

Pull requests, suggestions, and contributions are welcome!

Happy ASCII-ing ✨
