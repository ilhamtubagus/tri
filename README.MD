## 📝 Simple To-Do Application (CLI) with Cobra and Golang

A simple command-line To-Do application to help you manage your tasks with priority levels. This tool allows you to quickly add tasks, list them, and mark them as done—all from the terminal. This project demonstrate how we can create CLI app in golang at ease with the help of [cobra](https://cobra.dev/).

### 🚀 Features
- ✅ **Add To-Do with Priority:** Create tasks and set their priority (`low`, `medium`, `high`).
- 📋 **List To-Dos:** View all tasks, sorted by their status and priority.
- ☑️ **Mark To-Do as Done:** Quickly complete a task by marking it as done.

---

### 📂 Project Structure
```
todo-app/
├── cmd/
│   ├── add.go
│   ├── done.go
│   ├── list.go
│   └── root.go
├── main.go
└── go.mod
```

---

### 🛠️ Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/ilhamtubagus/tri.git
   cd tri
   ```

2. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

3. **Build the Application**:
   ```bash
   go build -o tri
   ```
---

### 💻 Usage

#### Add a To-Do
```bash
./todo add "Buy groceries" -p3
```
> Adds a new to-do with a **high priority**.

#### List All To-Dos
```bash
./todo list
```
> Displays all to-dos with their status and priority.

#### List Completed To-Dos
```bash
./todo list -d
```
> Displays to-dos marked as done.


#### Mark a To-Do as Done
```bash
./todo done 1
```
> Marks the to-do with **ID 1** as completed.

---

### ⚙️ Commands Overview

| Command             | Description                  | Example Usage                 |
|---------------------|-----------------------------|-------------------------------|
| `add`              | Add a new to-do             | `./todo add "Read book" -p3,` |
| `list`             | List all to-dos             | `./todo list`                 |
| `done`             | Mark a to-do as completed   | `./todo done 2`               |

---

### 📝 Example Output
![CleanShot 2025-02-17 at 09 50 57@2x](https://github.com/user-attachments/assets/0f50f420-8010-4bc9-b34d-962ea708f086)

---

### 🚧 Dependencies
- [Golang](https://go.dev/) (v1.19 or above)
- [Cobra](https://github.com/spf13/cobra) (for CLI framework)

---

### 📜 License
This project is licensed under the [MIT License](LICENSE).

---

### 📞 Contact
- **GitHub:** [ilhamtubagus](https://github.com/ilhamtubagus)
- **Email:** ilhamta27@gmail.com
