
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
* **application**, contiene la lógica de negocio de la aplicación,  por ejemplo las validaciones de formato del código SKU.
* **domain**, definiciones de objetos propias de nuestra lógica de negocio
* **infraestructure**, implementaciones concretas de las interfaces definidas por la capa de aplicación, así como también el código que usa el aplicativo para comunicarse con el exterior, por ejemplo el código que expone el socket para que los clientes se conecten
* **bin**, ficheros ejecutables de nuestra aplicación una vez construida, para generarlos lanzar el comando "make" sin parámetros, generará un ejecutable para linus y uno para windows

# Build
```bash
$ make
```

# Run test
```bash
$ make test
$ make testv //muestra más detalle de la ejecución de los test
```