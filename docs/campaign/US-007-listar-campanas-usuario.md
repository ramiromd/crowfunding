# Listar campañas de un usuario

## Narrativa
**Como** usuario de la plataforma
**Quiero** obtener el listado de campañas creadas por un usuario específico
**Para** explorar la actividad de un creador y decidir si donar a alguna de sus campañas

## Criterios de Aceptación (Gherkin)

- **Escenario:** Listado exitoso de campañas de un usuario existente
  - **Dado que** el usuario con el `id` indicado existe y tiene campañas activas
  - **Cuando** envía una petición `GET /users/{id}/campaigns`
  - **Entonces** recibe una respuesta `200 OK` con las campañas del usuario ordenadas por `(created_at DESC, id DESC)`, incluyendo `id`, `title`, `description`, `goal_amount`, `raised_amount`, `end_date` y un campo `next_cursor`

- **Escenario:** Página siguiente usando cursor válido
  - **Dado que** el cliente dispone de un `next_cursor` obtenido en una respuesta previa
  - **Cuando** envía una petición `GET /users/{id}/campaigns?after=<cursor>`
  - **Entonces** recibe una respuesta `200 OK` con los registros inmediatamente posteriores al cursor, pertenecientes al mismo usuario

- **Escenario:** Usuario sin campañas activas
  - **Dado que** el usuario con el `id` indicado existe pero no tiene campañas activas
  - **Cuando** envía una petición `GET /users/{id}/campaigns`
  - **Entonces** recibe una respuesta `200 OK` con una colección vacía y `next_cursor` nulo

- **Escenario:** Usuario no encontrado
  - **Dado que** el `id` indicado no corresponde a ningún usuario registrado
  - **Cuando** envía una petición `GET /users/{id}/campaigns`
  - **Entonces** recibe una respuesta `404 Not Found`

- **Escenario:** Cursor inválido o malformado
  - **Dado que** el cliente envía un valor de cursor que no corresponde a un cursor válido
  - **Cuando** envía una petición `GET /users/{id}/campaigns?after=<cursor_invalido>`
  - **Entonces** recibe una respuesta `400 Bad Request` indicando que el cursor es inválido

- **Escenario:** Las campañas canceladas y cerradas no aparecen en el listado
  - **Dado que** el usuario tiene campañas con `cancelled_at` o `closed_at` no nulos
  - **Cuando** envía una petición `GET /users/{id}/campaigns`
  - **Entonces** dichas campañas no están incluidas en la respuesta

## Definición de Hecho (DoD)
- [ ] Endpoint `GET /users/{id}/campaigns` implementado, accesible sin autenticación, donde `{id}` es un UUID válido
- [ ] El cursor se construye codificando en base64 el par `(created_at, id)` del último registro devuelto
- [ ] La consulta filtra por `owner_id = $id`, `cancelled_at IS NULL` y `closed_at IS NULL`
- [ ] La consulta usa `WHERE (created_at, id) < ($cursor_created_at, $cursor_id)` con índice compuesto sobre `(owner_id, created_at DESC, id DESC)`
- [ ] La respuesta incluye `next_cursor` (string opaco) o `null` si no hay más registros
- [ ] El tamaño de página por defecto es configurable via variable de entorno (`CAMPAIGN_PAGE_SIZE`, default 20)
- [ ] Código revisado
- [ ] Cobertura de al menos 80% en el handler y la capa de repositorio
- [ ] Documentación actualizada (contrato de la API en OpenAPI/Swagger o equivalente)

---
*Nota para la IA: Mantener un lenguaje técnico preciso y evitar ambigüedades.*
*Nota para la IA: Cuando se requiera paginar listados utilizar la estrategia de cursor por sobre paginado por offset, cuando sea necesario*
