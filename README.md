## 1-Go Modules: Creating, Using, and Managing Dependencies

This guide shows you how to create your own Go module, install external modules, and keep your dependencies clean using Go's module system.

---

### 🧱 Create Your Own Go Module

To start a new Go module, run the following command inside your project directory:

```bash
go mod init github.com/your_github_username/your_module_name
// example: go mod init github.com/lmarcela/test
```

This will create a go.mod file which defines your module path and tracks its dependencies.

### 📦 Install External Modules

You can add external dependencies to your project using go get:

```bash
go get github.com/username/module
//examples:
go get github.com/donvito/hellomod
go get github.com/donvito/hellomod/v2
go get github.com/lmarcela/go-test-hello --> Used in 16-mymodule dir
```

If the module uses semantic import versioning (v2 or higher), you must include the version suffix in the path (/v2, /v3, etc.).

The module and its version will be recorded in go.mod and go.sum

### 🧹 Clean Up Unused Dependencies

```bash
go mod tidy
```

This will:

Remove unused imports

Add missing ones that are referenced in your code but not yet included

## 2-Testing

```bash
go test
```

You have to create a module before run the test command (ex: go mod init github.com/your_github_username/your_module_name)

### Run tests with coverage

```bash
go test -cover
```

### Generate a coverage file

1. Generate coverage.out file

```bash
go test -coverprofile=coverage.out
```

2. Analysis with Go tool cover

```bash
go tool cover -func=coverage.out
```

3. HTML visualization

```bash
go tool cover -html=coverage.out -o coverage.html
```

Then open the coverage.html file in your browser

Or, everything in the same command:

```bash
go test -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html
```

### Profiling

Profiling is a key tool for analyzing and improving the performance of your applications.

NOTE: To be able to export to pdf, make sure you have previously installed this on linux:

```bash
sudo apt install graphviz
```

Visualize the profiling:

```bash
go test -cpuprofile cpu.out
go tool pprof cpu.out
(pprof) pdf
```
