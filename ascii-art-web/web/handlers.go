// Handles all HTTP routes: form, index, export, etc.

package web

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// indexHandler serves the homepage (index.html)
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Return 404 if the path is anything other than "/"
	if r.URL.Path != "/" {
		renderErrorPage(w, r, http.StatusNotFound)
		return
	}

	// Build path to the index template
	path := filepath.Join("templates", "index.html")

	// Check if the template file exists before serving it
	if _, err := os.Stat(path); err != nil {
		renderErrorWithMessage(w, http.StatusInternalServerError, "Homepage is temporarily unavailable.")
		return
	}

	// Serve the index page
	http.ServeFile(w, r, path)
}

// asciiArtHandler handles the form submission for ASCII art generation.
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// Redirect GET requests to the homepage to prevent resubmission on refresh
	if r.Method == http.MethodGet {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Only allow POST requests; reject others with a 404 error
	if r.Method != http.MethodPost {
		renderErrorPage(w, r, http.StatusNotFound)
		return
	}

	// Extract and validate form parameters from the POST request
	params, err := extractAsciiParams(r)
	if err != nil {
		renderErrorWithMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	// Generate ASCII art based on the extracted parameters
	ascii, err := generateAsciiArt(params)
	if err != nil {
		// Handle internal vs. client-caused errors differently
		if strings.HasPrefix(err.Error(), "internal:") {
			renderErrorWithMessage(w, http.StatusInternalServerError, strings.TrimPrefix(err.Error(), "internal: "))
		} else {
			renderErrorWithMessage(w, http.StatusBadRequest, err.Error())
		}
		return
	}

	// Send successful response with the ASCII art in HTML format
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(ascii))
}

// asciiTableHandler serves the ASCII table reference page
func asciiTableHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("templates", "ascii_table.html")

	// Return error if file doesn't exist
	if _, err := os.Stat(path); err != nil {
		renderErrorWithMessage(w, http.StatusInternalServerError, "ASCII Table is temporarily unavailable.")
		return
	}

	// Serve the ASCII table HTML page
	http.ServeFile(w, r, path)
}

// extractAsciiParams parses and validates user input from the form
func extractAsciiParams(r *http.Request) (*asciiRequest, error) {
	// Normalize text input and validate length
	text := strings.ReplaceAll(r.FormValue("inputText"), "\r", "")
	if len(text) > 1_000_000 {
		return nil, fmt.Errorf("input too long - max is 1,000,000")
	}

	// Ensure only ASCII characters (except newline) are present
	for _, ch := range text {
		if ch != '\n' && (ch < 32 || ch > 126) {
			return nil, fmt.Errorf(`only ASCII characters are allowed - <a href="/ascii-table"><u>see the full ASCII chart HERE</u></a>`)
		}
	}

	// Get banner style and ensure both text and banner are present
	banner := r.FormValue("banner")
	if text == "" || banner == "" {
		return nil, fmt.Errorf("missing text or banner")
	}

	// Support for optional color highlighting for specific words
	colorTargets := r.Form["colorTarget"]
	targetColors := r.Form["targetColor"]

	// Return the parsed request object
	return &asciiRequest{
		Text:         text,
		Banner:       banner,
		Align:        r.FormValue("align"),
		GlobalColor:  r.FormValue("color"),
		ColorTargets: colorTargets,
		TargetColors: targetColors,
	}, nil
}

// withRecover wraps HTTP handlers with panic recovery to prevent crashes
func withRecover(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Recover from any unexpected panic during handler execution
		defer func() {
			if rec := recover(); rec != nil {
				renderErrorWithMessage(w, http.StatusInternalServerError, "Unexpected server error.")
			}
		}()
		h(w, r)
	}
}

// handleExport handles exporting the generated ASCII art in various formats
func handleExport(w http.ResponseWriter, r *http.Request) {
	// Only allow POST method for exporting
	if r.Method != http.MethodPost {
		renderErrorPage(w, r, http.StatusNotFound)
		return
	}

	// Get ASCII art from form data
	text := r.FormValue("asciiText")
	if text == "" {
		renderErrorWithMessage(w, http.StatusBadRequest, "Empty output, nothing to export.")
		return
	}

	// Get export format and filename
	format := r.FormValue("format")
	filename := r.FormValue("filename")
	if filename == "" {
		filename = "ascii-art-web-export"
	}
	if format == "" {
		format = "txt"
	}

	// Determine MIME type based on format
	contentType := map[string]string{
		"txt":  "text/plain; charset=utf-8",
		"html": "text/html; charset=utf-8",
		"json": "application/json",
		"svg":  "image/svg+xml",
	}[format]

	// Fallback to default format if unknown
	if contentType == "" {
		contentType = "application/octet-stream"
		format = "txt"
	}

	var output string

	// Format output based on selected format
	switch format {
	case "json":
		output = fmt.Sprintf(`{"ascii":"%s"}`, template.JSEscapeString(text))
	case "html":
		output = "<pre>" + template.HTMLEscapeString(text) + "</pre>"
	case "svg":
		// Embed ASCII text in SVG format for vector export
		output = fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="100%%" height="100%%">
  <text x="0" y="15" font-family="monospace" font-size="14">%s</text>
</svg>`, template.HTMLEscapeString(text))
	default:
		output = text
	}

	// Set headers and deliver file as a downloadable attachment
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.%s\"", filename, format))
	w.Header().Set("Content-Length", strconv.Itoa(len(output)))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(output))
}
