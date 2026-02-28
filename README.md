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
│   │   ├── application/        # Casos de uso
│   │   └── infrastructure/     # Handlers HTTP, repositorios de BD
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
|       └── db/
├── go.mod
└── go.sum
```