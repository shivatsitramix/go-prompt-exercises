Prompt 1:

Project Prompt: Multi-Page Flutter Expense Tracker with Go Backend (JSON Dataset Storage)
Background:
Create an expense tracker app using Flutter that works offline but supports syncing to a Go backend. The Go backend should store user expenses in individual JSON files based on a token.

Objective:
Build a Flutter app with multiple pages:
Add expenses
View history
Filter by category/date
See total summary
Sync to backend

Build a Go backend that:
Accepts expense data
Stores each user's data in a separate JSON file

Input:
Flutter Side:
Title (string)
Amount (float)
Category (string): Food, Travel, Rent, Misc
Date (date)
Token (string for backend sync)

Go API:
JSON payload:
{
  "id": 1,
  "title": "Groceries",
  "amount": 150.5,
  "category": "Food",
  "date": "2025-06-05T10:00:00Z"
}
Header:
Authorization: Bearer <your-token>

Data Structures:

Dart (Flutter):
class Expense {
  final int id;
  final String title;
  final double amount;
  final String category;
  final DateTime date;
}
Go:
type Expense struct {
  ID       int       `json:"id"`
  Title    string    `json:"title"`
  Amount   float64   `json:"amount"`
  Category string    `json:"category"`
  Date     time.Time `json:"date"`
}
Processing:
Flutter:
Add screen: validate inputs, save locally
History screen: list all entries, allow filtering
Summary screen: calculate totals
Sync: send list of expenses to Go backend via /sync


Go Backend:
On sync, load existing JSON file for the token
Merge or replace entries
Save to file: data/data_<token>.json
On fetch, return file contents as JSON

Output:
Flutter:
List of expenses
Filtered views
Summary with total amount and category-wise breakdown
Sync status (Success/Fail)

Go:
JSON responses
Data saved per token in files like data/data_abc123.json
HTTP status codes: 200 (OK), 401 (Unauthorized), 400 (Bad Request)

Constraints:

Flutter:
Must work offline
Use provider, shared_preferences, http
Clean, basic UI with bottom navigation or drawer
No backend required to run core features

Go Backend:
Store each user's expenses in a separate JSON file
File path: data/data_<token>.json
Use Authorization header for simple token-based auth
Do not use SQL or external DBs
Ensure basic file locking/concurrency safety


Prompt 2: 
For the code present, we get this error:
```
Undefined name 'SharedPreferences'.
Try correcting the name to one that is defined, or defining the name.
```
How can I resolve this? If you propose a fix, please make it concise.


Prompt 3:
Background:
You have a working multi-page Flutter frontend and a functional Go backend. Currently, everything works, but the Go backend needs to be modularized, organized into folders, and structured to better support future scalability and integration with the frontend. You also want to enable full CORS (Access-Control-Allow-Origin: *) during development.

Objective:
Restructure the Go backend by separating routes, middleware, models, and utils. Update the API design based on the frontend needs, and ensure clear CORS support for cross-origin requests. The Flutter frontend will remain unchanged but will seamlessly integrate with this structured backend.

Input:
Flutter Side (already functional):
Add expenses
View history
Filter by category/date
Sync expenses via token-based auth

Expected API Usage from Flutter:
POST /sync — Sync expenses list
GET /expenses — Get all expenses for token
DELETE /expenses?id=<id> — Delete expense
API Request Headers:
Authorization: Bearer <token>

Request Body for /sync (JSON):
[
  {
    "id": 1,
    "title": "Groceries",
    "amount": 150.5,
    "category": "Food",
    "date": "2025-06-05T10:00:00Z"
  }
]
Data Structures:
Dart (Flutter):

class Expense {
  final int id;
  final String title;
  final double amount;
  final String category;
  final DateTime date;
}
Go (models/expense.go):
type Expense struct {
  ID       int       `json:"id"`
  Title    string    `json:"title"`
  Amount   float64   `json:"amount"`
  Category string    `json:"category"`
  Date     time.Time `json:"date"`
}

Processing:
Flutter Side:
Collect expense details from UI
Save locally and sync to backend when online
Display synced expenses and allow deletion

Go Backend (Refactored):
Layer	Responsibility
routes/	Handle /sync, /expenses, and /expenses/delete
middleware/	CORS middleware with Access-Control-Allow-*
models/	Define the Expense struct
utils/	JSON file operations and mutex locking for tokens

Output:
Flutter:
Synced & filtered expenses view
Summary screen using backend data

Go:
Store JSON files like data/data_<token>.json

Return HTTP codes:
200 for OK
401 for Unauthorized
400 for Bad Request

Constraints:
Must support multiple users via token
No SQL or external DBs — JSON file storage only
Must handle CORS for all origins during development
Each handler must run in isolated route files
File locking to prevent race conditions when writing

Command to Run:
go run main.go
Runs server at: http://localhost:8080


Prompt 4:
sync function in frontend is not working correctly please checknvalid request body: parsing time "2025-06-05T15:02:06.339" as "2006-01-02T15:04:05Z07:00": cannot parse "" as "Z07:00