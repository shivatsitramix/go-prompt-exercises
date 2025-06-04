package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// Constants for file size limits
const (
	maxFileSize = 10 * 1024 * 1024 // 10MB limit
	filesDir    = "files"          // Directory containing files
)

// Custom errors
var (
	ErrFileTooLarge = errors.New("file size exceeds maximum limit")
	ErrFileEmpty    = errors.New("file is empty")
	ErrInvalidPath  = errors.New("invalid file path")
)

// readFile reads the entire contents of a file and returns it as a string
func readFile(filename string) (string, error) {
	// Clean the file path to prevent directory traversal
	cleanPath := filepath.Clean(filename)

	// Ensure the file is within the files directory
	absFilesDir, err := filepath.Abs(filesDir)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path: %w", err)
	}

	absFilePath, err := filepath.Abs(cleanPath)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path: %w", err)
	}

	// Check if the file is within the files directory
	if !strings.HasPrefix(absFilePath, absFilesDir) {
		return "", ErrInvalidPath
	}

	// Check if file exists and get its info
	fileInfo, err := os.Stat(cleanPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("file does not exist: %w", err)
		}
		if os.IsPermission(err) {
			return "", fmt.Errorf("permission denied: %w", err)
		}
		return "", fmt.Errorf("failed to access file: %w", err)
	}

	// Check if file is empty
	if fileInfo.Size() == 0 {
		return "", ErrFileEmpty
	}

	// Check file size
	if fileInfo.Size() > maxFileSize {
		return "", ErrFileTooLarge
	}

	// Open the file
	file, err := os.Open(cleanPath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Read the entire file
	content, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return string(content), nil
}

// Server configuration
type Server struct {
	router *http.ServeMux
}

// NewServer creates a new server instance
func NewServer() *Server {
	return &Server{
		router: http.NewServeMux(),
	}
}

// setupRoutes configures all routes for the server
func (s *Server) setupRoutes() {
	s.router.HandleFunc("/", s.handleRoot)
	s.router.HandleFunc("/health", s.handleHealth)
	s.router.HandleFunc("/files/", s.handleFileRead)
}

// handleRoot handles requests to the root path
func (s *Server) handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Welcome to the Go HTTP Server!")
}

// handleHealth handles health check requests
func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status": "healthy"}`)
}

// handleFileRead handles requests to read files
func (s *Server) handleFileRead(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract filename from URL path and clean it
	filename := strings.TrimPrefix(r.URL.Path, "/files/")
	if filename == "" {
		http.Error(w, "No file specified", http.StatusBadRequest)
		return
	}

	// Clean the filename to prevent directory traversal
	filename = filepath.Clean(filename)
	if strings.Contains(filename, "..") {
		http.Error(w, "Invalid file path", http.StatusBadRequest)
		return
	}

	// Construct full file path
	filepath := filepath.Join(filesDir, filename)

	// Read file contents
	content, err := readFile(filepath)
	if err != nil {
		switch {
		case errors.Is(err, os.ErrNotExist):
			http.Error(w, "File not found", http.StatusNotFound)
		case errors.Is(err, os.ErrPermission):
			http.Error(w, "Permission denied", http.StatusForbidden)
		case errors.Is(err, ErrFileTooLarge):
			http.Error(w, "File too large", http.StatusRequestEntityTooLarge)
		case errors.Is(err, ErrFileEmpty):
			http.Error(w, "File is empty", http.StatusBadRequest)
		case errors.Is(err, ErrInvalidPath):
			http.Error(w, "Invalid file path", http.StatusBadRequest)
		default:
			log.Printf("Error reading file: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	// Set content type based on file extension
	ext := path.Ext(filename)
	switch ext {
	case ".txt":
		w.Header().Set("Content-Type", "text/plain")
	case ".json":
		w.Header().Set("Content-Type", "application/json")
	case ".html":
		w.Header().Set("Content-Type", "text/html")
	default:
		w.Header().Set("Content-Type", "application/octet-stream")
	}

	fmt.Fprint(w, content)
}

func main() {
	// Create and configure server
	server := NewServer()
	server.setupRoutes()

	// Configure HTTP server with timeouts
	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      server.router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start the server
	fmt.Println("Server starting on :8080...")
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
