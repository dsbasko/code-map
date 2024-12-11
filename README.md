# Code-Map

**Code-map** is a utility designed for recursively searching files in a directory and merging them into a single file with 
markup. It is useful for project documentation, report creation, and combining the contents of source code or 
other files into one markdown file.

## Features

- Recursively traverses all nested directories.
- Filters files based on a specified pattern (regular expression).
- Generates a `project.md` file with the contents of the found files wrapped in markdown format.

## Usage

### Run the utility:
```bash
./code-map ~/Develop/sandbox "\\.(go|json)\$"
```

After execution, a project.md file will be created with the contents of the project.

### Arguments

1. `path` — the path to the directory where the recursive search should be performed.
2. `pattern` — a regular expression for filtering files by name.

### Build

Ensure that Go (version 1.23 and above) is installed.

1. Clone the repository:
    ```bash
    git clone https://github.com/your-repository/project-tree.git
    ```

2. Navigate to the project directory:
    ```bash
    cd project-tree
    ```

3. Build the utility:
    ```bash
    go build -o project-tree
    ```

### Example

Suppose you have the following project structure:

```
project/
├── cmd/
│   └── main.go
├── internal/
│   ├── configs/
│   │   └── configs.go
│   └── repository/
│       └── postgresql.go
```

Contents of the files: 

**cmd/main.go:**
```go
package main

func main() {
    // some logic
}
```

**internal/configs/configs.go:**
```go
package configs

func someConfigFN() {
    // some logic
}
```

**internal/repository/postgresql.go:**
```go
package repository

func someRepositoryFN() {
    // some logic
}
```
