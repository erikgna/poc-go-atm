# First contact with Go

### Struct

- I learned that Go doesn't have classes, therefore it doesn't have inheritance.
- Instead of classes, Go have structs, which can hold functions or be used as objects.

### Pointers

- Pointers are variables that store the memory address of an object.
- Go has pointers, which helps not to create copies of large objects (structs).
- Pointers also allow us to modify the original state of an object.
- It also helps to share the same value between multiples layers.

### Tests

- Tests in Go are simple, you can create a file with \_test.go.
- Each method needs to begin with "Test" to be evaluated.

### Modules

- Go modules can be created with `go mod init go-atm/atm`
- To export a method from the module, it needs to start with uppercase.
- To use it, you can import and call methods with `atm.MethodName`
- To use local modules you need to pass the path in the go.mod file = `replace go-atm/atm => ./atm`
- `go mod tidy` to sync modules.
