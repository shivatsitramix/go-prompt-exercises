## 🧠 Prompting Practice and Fullstack Integration Guide (Flutter + Go)

### Overview

This repository documents a **3-day intensive fullstack development and prompt engineering training** where I practiced structured task breakdown, API design, and prompt improvement for Flutter frontend and Go backend.

---

## 🔁 Project Summary

* **Frontend:** Flutter app for offline-first expense tracking.
* **Backend:** Go server using JSON-based storage per user (no DB).
* **Integration:** Full CORS-enabled API with sync, fetch, and delete routes.
* **Focus:** Improve technical prompting, backend modularity, and real-world dev workflow simulation.

---

## 🛠️ Backend API Routes (Go)

| Route       | Method | Description                    | Auth Required |
| ----------- | ------ | ------------------------------ | ------------- |
| `/sync`     | POST   | Sync expense list              | Yes (Bearer)  |
| `/expenses` | GET    | Fetch all expenses             | Yes (Bearer)  |
| `/expenses` | DELETE | Delete expense by `?id=` param | Yes (Bearer)  |

📁 Stored JSON files in `data/data_<token>.json`
🔒 File locking implemented to avoid race conditions

---

## 📲 Flutter Features

* Add, edit, and view expenses
* Filter by category and date
* Sync to backend using token
* Offline-first functionality
* Bottom navigation and drawer-based navigation

---

## 📌 Prompting Best Practices

### ✅ Flutter-Specific Prompting

* **Be UI/UX precise**: Mention widget structure (`Column`, `Padding`, `TextField`) with reasons.
* **Break into steps**: "Step 1: Create widget → Step 2: Reuse widget → Step 3: Display result"
* **Mention file names**: e.g., “Create `custom_button.dart` for reusable widget”
* **State the output clearly**: “Expect button to print to console”
* **Use scenario-based phrasing**: “As a user enters an expense…”

### ✅ Go-Specific Prompting

* **Always specify route, method, header, and body format**
* **Modular thinking**: Split prompts into `routes/`, `models/`, `utils/`, `middleware/`
* **Ask like an engineer**: “Refactor CORS handling middleware-wise”
* **Mention actual log errors**: Instead of “backend broken,” give logs
* **Request precise behavior**: “When file not found, return 401 with message”

---

## 📘 Prompting Evolution (3-Day Learning)

### 🔴 **Day 1–2: What I Did Wrong**

* Wrote **vague prompts** with no input/output or API structure
* Combined multiple tasks into **confusing one-liners**
* Missed key context like headers, tokens, payload formats
* Didn’t anticipate edge cases like CORS, file not found, or bad date format
* Just asked for solutions — not how to think or plan like a dev

### 🟢 **Day 3: What I Learned**

* Used **Chain of Thought, Tree of Thought, multi-shot, role-based** prompting
* Wrote **structured instructions** with clear context, logs, inputs, and outputs
* Asked prompts like a **team lead**, not a stuck developer
* **Linked problems and solutions** across multiple prompts
* Used **real-world scenarios** to guide prompt clarity

---

## ⚙️ Challenges Faced

| Challenge                        | Solution Through Prompting                                      |
| -------------------------------- | --------------------------------------------------------------- |
| CORS errors                      | Modularized middleware and asked for dev-friendly CORS handling |
| Sync failing due to date parsing | Clarified expected RFC format in prompt                         |
| Structuring Go backend           | Wrote refactor prompts with folders and responsibilities        |
| Flutter input not saving         | Prompted using widget logic breakdown + controllers             |
| Backend port conflict            | Prompted about proper process kill / port reuse                 |

---

## 🏁 Final Takeaways

* **Prompting is engineering communication** — not just asking questions.
* Good prompts include: **goal, input format, desired output, edge cases, context.**
* Advanced prompting techniques (CoT, ToT, Interview Pattern) make a huge difference when used properly.
* By Day 3, I wasn’t just writing prompts to fix issues — I was designing systems via prompts.

---

## 📎 Folder Structure Summary

### Flutter

```
lib/
├── screens/
├── widgets/
├── models/
├── providers/
└── main.dart
```

### Go

```
backend/
├── routes/
├── middleware/
├── models/
├── utils/
└── main.go
```

---

## 🚀 How to Run

### Flutter

```bash
flutter run
```

### Go

```bash
go run main.go
# Runs at http://localhost:8080
```
