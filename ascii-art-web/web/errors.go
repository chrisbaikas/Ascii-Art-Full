// Renders error pages using templates and helper functions

package web

import (
	"html/template"
	"net/http"
	"strconv"
)

// ErrorData holds the error code and a message that may include HTML
type ErrorData struct {
	Code    int
	Message template.HTML
}

// Preload error template at startup
var errorTemplate = template.Must(template.ParseFiles("templates/error.html"))

// errorPageHandler reads ?code= from URL and shows the appropriate error page
func errorPageHandler(w http.ResponseWriter, r *http.Request) {
	codeStr := r.URL.Query().Get("code")
	code, err := strconv.Atoi(codeStr)
	if err != nil {
		code = http.StatusInternalServerError
	}
	renderErrorPage(w, r, code)
}

// renderErrorPage shows a predefined error based on status code
func renderErrorPage(w http.ResponseWriter, _ *http.Request, status int) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)

	messages := map[int]string{
		http.StatusBadRequest:          "Bad Request: Please check your input.",
		http.StatusNotFound:            "Page Not Found.",
		http.StatusInternalServerError: "Internal Server Error: Something went wrong.",
	}

	msg, ok := messages[status]
	if !ok {
		msg = "An unexpected error occurred."
	}

	errorTemplate.Execute(w, ErrorData{
		Code:    status,
		Message: template.HTML(msg),
	})
}

// renderErrorWithMessage allows dynamic custom error messages (including HTML)
func renderErrorWithMessage(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(code)

	errorTemplate.Execute(w, ErrorData{
		Code:    code,
		Message: template.HTML(message), // Important: trust only controlled inputs
	})
}
