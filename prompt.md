# Go File Server Exercise

## Prompt:-
Write a Go program with the following features:
A function that reads the entire contents of a file and returns it as a string.
A basic HTTP server that responds with a message on the root route (/).
Now, imagine this is a live coding interview. The interviewer asks:
How will your file reading function handle errors like missing files, permission denied, or partial read failures? Update your function to handle these errors properly and explain your approach.
How does your HTTP server handle routing? What happens if a user requests an undefined route or sends a malformed request? Improve your server code to manage these cases and explain your design.
Consider edge cases such as reading very large files without exhausting memory, handling concurrent HTTP requests efficiently, and managing network timeouts or failures. Refine your code accordingly and describe your solutions.
Respond step-by-step, improving your code and explaining each change as if answering an interviewer’s questions.

## Requirements

### File Reading Function
- Should read entire file contents and return as string
- Must handle various error cases:
  - File not found
  - Permission denied
  - Empty files
  - Files too large
  - Invalid paths
- Should implement size limits to prevent memory exhaustion

### HTTP Server
- Should serve files from a designated directory
- Must handle routing properly
- Should implement proper error handling
- Must include security measures against directory traversal
- Should set appropriate Content-Type headers

## Implementation Details

### File Structure
```
.
├── main.go
└── files/
    └── sample.txt
```

### Key Features Implemented

1. **File Reading Function**
   - Size limit of 10MB
   - Path validation and cleaning
   - Comprehensive error handling
   - Resource cleanup with defer

2. **HTTP Server**
   - Routes:
     - `/` - Welcome message
     - `/health` - Health check endpoint
     - `/files/` - File serving endpoint
   - Security features:
     - Directory traversal prevention
     - Path cleaning
     - Access restriction to files directory
   - Error handling:
     - 404 for non-existent files
     - 403 for permission denied
     - 413 for files too large
     - 400 for invalid requests
     - 405 for wrong HTTP methods

3. **Content Type Handling**
   - Automatic Content-Type based on file extension
   - Support for:
     - .txt files (text/plain)
     - .json files (application/json)
     - .html files (text/html)
     - Other files (application/octet-stream)

## Testing

1. Start the server:
```bash
go run main.go
```

2. Access endpoints:
- Welcome page: `http://localhost:8080/`
- Health check: `http://localhost:8080/health`
- Sample file: `http://localhost:8080/files/sample.txt`

## Security Considerations

1. **Path Traversal Prevention**
   - Cleaning file paths
   - Checking for ".." in paths
   - Validating against absolute paths

2. **Resource Management**
   - File size limits
   - Proper file handle cleanup
   - Memory usage control

3. **Error Handling**
   - Comprehensive error messages
   - Proper HTTP status codes
   - Secure error responses

## Learning Outcomes

1. **Go Concepts**
   - HTTP server implementation
   - File operations
   - Error handling
   - Path manipulation
   - Content-Type headers

2. **Security Best Practices**
   - Path validation
   - Resource limits
   - Error handling
   - Directory traversal prevention

3. **Web Server Concepts**
   - Routing
   - HTTP methods
   - Status codes
   - Content types
   - Error responses

## Next Steps

1. Add authentication
2. Implement file upload functionality
3. Add caching mechanisms
4. Implement rate limiting
5. Add logging and monitoring
6. Implement file compression
7. Add support for range requests 