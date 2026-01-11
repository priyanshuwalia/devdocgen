
# devdocgen

![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.21-00ADD8?logo=go)
![Built With](https://img.shields.io/badge/Built%20With-devdocgen-orange)
---
## Project Structure
```text
├── README.generated.md
├── README.md
├── cmd
│   ├── root.go
│   └── scan.go
├── go.mod
├── go.sum
├── internal
│   ├── detector
│   │   ├── detector.go
│   │   ├── docker.go
│   │   ├── env.go
│   │   ├── go.go
│   │   ├── license.go
│   │   ├── name.go
│   │   ├── node.go
│   │   └── node_scripts.go
│   ├── model
│   │   └── project.go
│   ├── render
│   │   ├── readme.go
│   │   ├── templates
│   │   │   └── readme.md.tmpl
│   │   └── templates.go
│   └── scanner
│       ├── scanner.go
│       └── tree.go
└── main.go

```
---

## Tech Stack
- Go
- Cobra (CLI)


