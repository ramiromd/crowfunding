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

### Convenciones de tests
**TODO : Definir ubicación de tests de integración y de contratos.**
| Tipo | Ubicación |
|---|---|
| Unitarios | Co-locados en cada capa del dominio (`*_test.go`) |