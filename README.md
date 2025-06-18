# Go Modules (class 15): Creating, Using, and Managing Dependencies

This guide shows you how to create your own Go module, install external modules, and keep your dependencies clean using Go's module system.

---

## 🧱 Create Your Own Go Module

To start a new Go module, run the following command inside your project directory:

```bash
go mod init github.com/your_github_username/your_module_name
// example: go mod init github.com/lmarcela/test
```

This will create a go.mod file which defines your module path and tracks its dependencies.

## 📦 Install External Modules

You can add external dependencies to your project using go get:

```bash
go get github.com/username/module
//examples:
go get github.com/donvito/hellomod
go get github.com/donvito/hellomod/v2
go get github.com/lmarcela/go-
test-hello --> Used in 16-mymodule dir
```

If the module uses semantic import versioning (v2 or higher), you must include the version suffix in the path (/v2, /v3, etc.).

The module and its version will be recorded in go.mod and go.sum

## 🧹 Clean Up Unused Dependencies

```bash
go mod tidy
```

This will:

Remove unused imports

Add missing ones that are referenced in your code but not yet included
