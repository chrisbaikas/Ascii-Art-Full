# ASCII Art Tool with Color, Alignment & Reverse

This project is a command-line tool written in Go that allows converting text into stylized ASCII art, supports alignment and color highlighting, and can reverse ASCII art back to readable text.

# Authors

* chbaikas
* askordal
* nektariospan


## ✨ Features

- ✅ **Text to ASCII Art Conversion**
- 🎨 **Coloring full or partial text**
- 📐 **Alignment (left, center, right, justify)**
- 📁 **Output to file or terminal**
- 🔁 **Reverse ASCII Art back to plain text**
- 🧱 **Support for custom banner fonts**
- 📦 **No external packages used – standard Go only**

---

## 🚀 Usage

```bash
go run . [--reverse=<input.txt>] [--color=<color>[:substring]] [--align=left|center|right|justify] [--output=file.txt] "text" [banner]
```

### 🔡 Text to ASCII Art

```bash
go run . "Hello"
go run . "Hello there" thinkertoy
go run . "Hi\nThere" shadow
```

### 🎨 Apply Colors

```bash
go run . --color=red "Alert"
go run . --color=blue:Go --color=green:Lang "Go Lang is awesome"
go run . '--color=rgb(255,0,0)' "RGB Colors!"
```

### 📐 Align Output

```bash
go run . --align=center "Aligned Center"
go run . --align=justify "Justified Text Here"
```

### 📤 Write to File

```bash
go run . --output=out.txt "Save me!"
```

### 🔁 Reverse ASCII Art

```bash
go run . --reverse=banner.txt
```

*Note: ASCII input must follow valid 8-line blocks per text line, optionally separated by empty lines.*

---

## 🎨 Color Options

You can use:
- **Named colors**: `red`, `green`, `blue`, `orange`, `pink`, `gray`, etc.
- **RGB**: `--color='rgb(255,0,0)'`
- **HEX**: `--color='#ff0000'`
- **HSL**: `--color='hsl(0, 100%, 50%)'`

To apply to a substring: `--color=blue:Go`

---

## 🏗️ Setup

1. Clone the project.
2. Ensure Go (1.16+) is installed.
3. Include banner files like:
   - `standard.txt`
   - `shadow.txt`
   - `thinkertoy.txt`

Each banner must support ASCII 32–126 using 8 lines per character.

---

## 🧪 Testing

```bash
go test ./utils
```

Covers:
- Alignment logic
- Color parsing & application
- ASCII parsing & generation
- Reverse decoding

---

## ❗ Error Handling

- Invalid banner: `Error reading banner file`
- Unsupported characters: `Error: unsupported character`
- Bad reverse input: `Error: invalid art height`
- Misused flags: Shows clear usage instructions

---

## License

This project is licensed under the [MIT License](LICENSE).

This project was created as part of a Zone01 educational project.  
Feel free to use, modify, and extend it.
