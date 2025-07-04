# ascii-art-web

**ascii-art-web** is a full-featured web-based ASCII Art generator written in Go. It features a clean user interface, responsive design, live preview, customizable colors, and multi-format export capabilities — all backed by a minimal but powerful Go HTTP server.

---

## ✨ Features

- ✅ Convert text to ASCII art using 3 banner styles (`standard`, `shadow`, `thinkertoy`)
- 🎨 Highlight substrings with color (targeted or global)
- 📐 Left or right alignment support
- 🧱 Responsive layout (mobile/tablet friendly)
- 🧑‍🎨 Background color customization
- 🎛️ Live updates via JavaScript debounce
- 💾 Export result as `.txt`, `.html`, `.json`, or `.svg`
- 🖌️ Beautiful UI using flexbox + iro.js color pickers
- 🚀 Fast and safe — built with only Go standard libraries

---

## 👤 Authors

- chbaikas  
- askordal  
- npanouts  

---

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone <repository-url>
cd ascii-art-web
```

### 2. Prepare assets

- Ensure `banners/` directory contains:
  - `standard.txt`
  - `shadow.txt`
  - `thinkertoy.txt`
- Make sure `templates/index.html` and `static/` (CSS/JS) folders are present.

### 3. Run the server

```bash
go run .
```

Then visit:

```
http://localhost:8080
```

---

## 🛠️ Implementation Details

### ASCII Engine

#### `LoadBanner`:
- Reads a banner `.txt` file line-by-line
- Splits lines into blocks of 8 (each ASCII character)
- Maps from rune to `[]string`

#### `AsciiArt`:
- Processes user text line-by-line
- Builds row-aligned ASCII output per line
- Applies ANSI coloring to targets and converts to `<span>`
- Handles left/right alignment using CSS

---

### 🌐 Web Server (Go)

#### Endpoints

- `GET /` → serves `index.html`
- `POST /ascii-art` → processes input, returns formatted HTML
- `POST /export` → returns downloadable file in chosen format
- `GET /ascii-table` → optional ASCII table reference

### HTTP Response Handling

- Color, banner, alignment, and highlighting parsed from form
- Responses formatted in:
  - Plain text
  - HTML (`<pre>`)
  - JSON
  - SVG (`<text>` inside `<svg>`)

---

## 🎨 Front-end

- Pure HTML/CSS/JS (no frameworks or modules)
- JavaScript:
  - Debounced input with `fetch()` request
  - Color sync between pickers and radio buttons
  - Dropdown banner selection
  - Live character count and validation
- CSS:
  - Uses `flexbox` for layout
  - Fully responsive for mobile
  - Custom dropdowns and buttons
  - Preserves formatting with `white-space: pre`

---

## 📁 File Structure

```
ascii-art-web/
│
├── banners/              # Text banner font files
├── templates/
│   └── index.html        # Main page template
├── static/
│   ├── style.css         # All CSS
│   ├── main.js           # init only
│   ├── color.js          # iro.js integration
│   ├── dropdown.js       # Banner dropdown logic
│   ├── generate.js       # ASCII art fetch logic
│   └── export.js         # Export to file logic
├── web/                  # Go handlers
│   └── handlers.go
├── go.mod
└── go.sum
```

---

## 📦 Exporting ASCII Art

### Features

- Export formats: `.txt`, `.html`, `.json`, `.svg`
- Custom filename
- Modal toggle to show export form
- Exports with proper headers (`Content-Disposition`, `Content-Type`, etc)

### API Endpoint

```
POST /export
```

**Body Parameters:**
- `asciiText`: the ASCII content
- `format`: one of `txt`, `html`, `json`, `svg`
- `filename`: name of exported file

**Server responds with downloadable file.**

---

## 🧪 How to Test

1. Enter text and choose options
2. Wait for live preview
3. Click **Export**
4. Select format and filename
5. Download and verify file

---

## 📌 Design Goals

- Immediate preview (debounced)
- Visibility under any color background
- Highlight individual substrings
- Clean, accessible UI (ARIA friendly)
- Fully responsive layout
- Minimal external dependencies (only `iro.js`)

---

## 📚 Learning Outcomes

- Go HTTP server development
- Form processing and error handling
- HTML + JavaScript + Go integration
- CSS architecture for responsiveness
- UX principles for live feedback
- File export and MIME headers in Go

---

## License

This project is licensed under the [MIT License](LICENSE).

---

## 🗣️ Feedback

Feel free to report issues or contribute improvements via GitHub Pull Requests or Issues.

Happy ASCII-ing 🎨