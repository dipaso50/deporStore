
# Project estructure

```bash
.
├── application
│   └── feeder
│       ├── feederService.go
│       ├── feederService_test.go
│       ├── productValidator.go
│       └── productValidator_test.go
├── bin 
│   ├── linux
│   │   └── feeder
│   └── windows
│       └── feeder.exe
├── domain
│   └── product.go
├── go.mod
├── infraestructure 
│   └── cmd
│       ├── socketEntry.go
│       └── socketEntry_test.go
├── main.go
├── Makefile
└── README.MD
```
* application, contiene la lógica de negocio de la aplicación, en nuestro caso, por ejemplo las validaciones de formato del código SKU y las estructuras donde los guardamos
* domain
* infraestructure

# Run test
```bash
$ make test
$ make testv //muestra más detalle de la ejecución de los test
```