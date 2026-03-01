# Project Standards

## Testing

### Convención de nombres

Los tests siguen un estilo **BDD (Behavior-Driven Development)**. El nombre de cada test debe leerse como una especificación del comportamiento esperado.

**Formato:**
```
TestNombreDeLaFuncion/should_[resultado]_when_[condición]
```

**Ejemplos:**
```
TestNewEntityId/should_create_entity_id_when_uuid_is_valid
TestNewEntityId/should_return_error_when_uuid_is_invalid
```

### Estructura

Se usa `t.Run()` para agrupar casos relacionados bajo una misma función de test.

```go
func TestAlgo(t *testing.T) {
    t.Run("should ... when ...", func(t *testing.T) {
        // assertions
    })

    t.Run("should ... when ...", func(t *testing.T) {
        // assertions
    })
}
```

### Package

Los test files usan el sufijo `_test` en el nombre del package (`package foo_test`) para realizar **black-box testing**, testeando únicamente la API pública del package.

### Assertions

Se usa `github.com/stretchr/testify` con la siguiente convención:

| Herramienta | Cuándo usarla |
|---|---|
| `require` | Cuando un fallo hace que las assertions siguientes no tengan sentido (ej: verificar que `err == nil` antes de usar el valor retornado) |
| `assert` | Para el resto de assertions, permitiendo ver todos los fallos en una sola ejecución |
| `assert.ErrorIs` | Para comparar errores (respeta error wrapping) |

### Ejecución

```bash
# Correr todos los tests
docker compose exec devtool go test ./...

# Con output detallado
docker compose exec devtool go test -v ./...

# Un package específico
docker compose exec devtool go test -v ./internal/shared/domain/value_object/...

# Sin cache
docker compose exec devtool go test -count=1 ./...
```
