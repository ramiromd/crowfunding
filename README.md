# Crowfunding

## Estructura del proyecto

El proyecto sigue un enfoque de **Clean Architecture** organizado por dominio. Cada dominio contiene sus propias capas internas y las dependencias fluyen hacia adentro: `infrastructure → application → domain`.

```
crowfunding/
├── cmd/
│   └── api/                    # Entrypoint de la aplicación
├── internal/
│   ├── campaign/               # Dominio: Campañas
│   │   ├── domain/             # Entidades, interfaces de repositorio, errores de dominio
│   │   │   └── *_test.go       # Tests unitarios co-locados
│   │   ├── application/        # Casos de uso
│   │   │   └── *_test.go       # Tests unitarios co-locados
│   │   └── infrastructure/     # Handlers HTTP, repositorios de BD
│   │       └── *_test.go       # Tests unitarios co-locados
│   ├── donation/               # Dominio: Donaciones
│   │   ├── domain/
│   │   ├── application/
│   │   └── infrastructure/
│   ├── identity/               # Dominio: Identidad y autenticación
│   │   ├── domain/
│   │   ├── application/
│   │   └── infrastructure/
│   └── shared/                 # Código compartido entre dominios
│       ├── values/
│       └── db/
├── test/
│   ├── integration/            # Tests de integración que cruzan dominios
│   └── contract/               # Tests de contrato (Gherkin / godog)
│       ├── features/           # Archivos .feature
│       └── steps/              # Implementación de los steps
├── iac/                        # Infrastructure as Code
│   ├── webapp/ 
│   ├── redis/   
│   ├── rabbit/                   
│   └── postgre/                     
├── go.mod
└── go.sum
```

## Desarrollo

Los comandos Go se ejecutan dentro del contenedor `crowfunding-devtool`, que monta el proyecto en `/app` y provee el toolchain de Go.

### Levantar el contenedor

```bash
docker compose up -d devtool
```

### Comandos disponibles

```bash
# Ejecutar la aplicación
docker compose exec devtool go run cmd/api/main.go

# Compilar
docker compose exec devtool go build -o bin/api cmd/api/main.go

# Correr tests
docker compose exec devtool go test ./...

# Agregar una dependencia
docker compose exec devtool go get <paquete>

# Abrir una shell interactiva
docker compose exec devtool bash
```

### Reporte de cobertura HTML

Se utiliza `gcov2lcov` + `genhtml` (lcov) para generar reportes de cobertura con navegación por paquete y resaltado por línea, similar a JaCoCo.

```bash
# 1. Generar el perfil de cobertura
docker compose exec devtool go test -coverprofile=reports/coverage/coverage.out ./...

# 2. Convertir a formato LCOV
docker compose exec devtool gcov2lcov -infile=reports/coverage/coverage.out -outfile=reports/coverage/coverage.lcov

# 3. Generar el HTML
docker compose exec devtool genhtml reports/coverage/coverage.lcov --output-directory reports/coverage/html/
```

El reporte queda disponible en `reports/coverage/html/index.html`.

### Convenciones de tests
**TODO : Definir ubicación de tests de integración y de contratos.**
| Tipo | Ubicación |
|---|---|
| Unitarios | Co-locados en cada capa del dominio (`*_test.go`) |