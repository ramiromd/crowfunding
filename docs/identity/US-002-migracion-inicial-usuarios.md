# Migración inicial y seed de la tabla de usuarios

## Narrativa
**Como** desarrollador del sistema
**Quiero** ejecutar la migración inicial que crea la tabla `users` y poblarla con usuarios de arranque
**Para** disponer de la estructura persistente necesaria y de datos base que permitan operar y probar la plataforma desde el primer despliegue

## Criterios de Aceptación (Gherkin)

- **Escenario:** Migración aplicada en una base de datos limpia
  - **Dado que** la base de datos no contiene la tabla `users`
  - **Cuando** se ejecuta la migración inicial (`migrate up`)
  - **Entonces** la tabla `users` es creada con las columnas: `id`, `email`, `password_hash`, `nickname`, `created_at`, `updated_at`

- **Escenario:** Restricción de unicidad sobre email
  - **Dado que** la tabla `users` ha sido creada
  - **Cuando** se intenta insertar dos registros con el mismo `email`
  - **Entonces** la base de datos rechaza la segunda inserción con un error de constraint de unicidad

- **Escenario:** Restricción de unicidad sobre nickname
  - **Dado que** la tabla `users` ha sido creada
  - **Cuando** se intenta insertar dos registros con el mismo `nickname`
  - **Entonces** la base de datos rechaza la segunda inserción con un error de constraint de unicidad

- **Escenario:** Reversión de la migración
  - **Dado que** la migración inicial ha sido aplicada
  - **Cuando** se ejecuta la operación de reversión (`migrate down`)
  - **Entonces** la tabla `users` es eliminada sin dejar objetos huérfanos (índices, secuencias)

- **Escenario:** Idempotencia al reaplicar la migración
  - **Dado que** la migración ya fue aplicada previamente
  - **Cuando** se intenta ejecutar `migrate up` nuevamente
  - **Entonces** la operación no produce errores ni duplica estructuras (uso de `CREATE TABLE IF NOT EXISTS` o control por versión de migración)

- **Escenario:** Seed ejecutado tras la migración
  - **Dado que** la tabla `users` ha sido creada
  - **Cuando** se ejecuta el script de seed (`seed.sql` o equivalente)
  - **Entonces** se insertan los usuarios iniciales con sus contraseñas almacenadas como hash bcrypt y sus roles asignados correctamente

- **Escenario:** Seed no duplica registros al re-ejecutarse
  - **Dado que** el seed ya fue aplicado previamente
  - **Cuando** se ejecuta el script de seed nuevamente
  - **Entonces** la operación no produce errores ni genera registros duplicados (idempotencia)

- **Escenario:** Contraseñas del seed no se almacenan en texto plano
  - **Dado que** el script de seed define usuarios iniciales
  - **Cuando** se inspeccionan los valores en la columna `password_hash`
  - **Entonces** todos los valores corresponden a hashes bcrypt válidos, sin ninguna contraseña en texto plano

## Definición de Hecho (DoD)
- [ ] Archivo de migración creado bajo `db/migrations/` con nombre versionado (ej. `000001_create_users_table.up.sql` / `.down.sql`)
- [ ] La tabla `users` incluye los campos: `id` (UUID PRIMARY KEY), `email` (VARCHAR UNIQUE NOT NULL), `password_hash` (TEXT NOT NULL), `nickname` (UNIQUE VARCHAR NOT NULL), `created_at` (TIMESTAMPTZ), `updated_at` (TIMESTAMPTZ)
- [ ] Índice único creado sobre la columna `email` y `nickname`
- [ ] Migración de reversión (`down`) implementada y verificada
- [ ] Migración ejecutada exitosamente en entorno local
- [ ] Script de seed creado bajo `db/seeds/` (ej. `users.sql` o `seed.go`)
- [ ] Las contraseñas del seed se generan como hashes bcrypt antes de la inserción (nunca texto plano en el script)
- [ ] El seed debe garantizar idempotencia
- [ ] Código revisado
- [ ] Cobertura de al menos 80% en tests de integración que validen la estructura de la tabla y la correcta inserción del seed
- [ ] Documentación actualizada (esquema de base de datos en `docs/`)

---
*Nota para la IA: Mantener un lenguaje técnico preciso y evitar ambigüedades.*