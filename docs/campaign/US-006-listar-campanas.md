# Listar campañas

## Narrativa
**Como** usuario de la plataforma
**Quiero** obtener un listado de campañas navegable mediante cursor
**Para** explorar las campañas disponibles y decidir a cuál donar

## Criterios de Aceptación (Gherkin)

- **Escenario:** Primera página de campañas sin cursor
  - **Dado que** existen campañas registradas en el sistema
  - **Cuando** envía una petición `GET /campaigns`
  - **Entonces** recibe una respuesta `200 OK` con los primeros registros ordenados por `(created_at DESC, id DESC)`, incluyendo `id`, `title`, `description`, `goal_amount`, `raised_amount`, `end_date`, `owner_id`, y un campo `next_cursor` para obtener la siguiente página

- **Escenario:** Página siguiente usando cursor válido
  - **Dado que** el cliente dispone de un `next_cursor` obtenido en una respuesta previa
  - **Cuando** envía una petición `GET /campaigns?after=<cursor>`
  - **Entonces** recibe una respuesta `200 OK` con los registros inmediatamente posteriores al cursor, manteniendo el mismo orden

- **Escenario:** Cursor apunta al último registro
  - **Dado que** no existen campañas posteriores al cursor enviado
  - **Cuando** envía una petición `GET /campaigns?after=<cursor>`
  - **Entonces** recibe una respuesta `200 OK` con una colección vacía y `next_cursor` nulo

- **Escenario:** Listado vacío cuando no hay campañas
  - **Dado que** no existen campañas registradas en el sistema
  - **Cuando** envía una petición `GET /campaigns`
  - **Entonces** recibe una respuesta `200 OK` con una colección vacía y `next_cursor` nulo

- **Escenario:** Cursor inválido o malformado
  - **Dado que** el cliente envía un valor de cursor que no corresponde a un cursor válido
  - **Cuando** envía una petición `GET /campaigns?after=<cursor_invalido>`
  - **Entonces** recibe una respuesta `400 Bad Request` indicando que el cursor es inválido

- **Escenario:** Las campañas canceladas y cerradas no aparecen en el listado
  - **Dado que** existen campañas con `cancelled_at` o `closed_at` no nulos
  - **Cuando** envía una petición `GET /campaigns`
  - **Entonces** dichas campañas no están incluidas en la respuesta

## Definición de Hecho (DoD)
- [ ] Endpoint `GET /campaigns` implementado, accesible sin autenticación
- [ ] El cursor se construye codificando en base64 el par `(created_at, id)` del último registro devuelto
- [ ] La consulta usa `WHERE (created_at, id) < ($cursor_created_at, $cursor_id)` con índice compuesto sobre `(created_at DESC, id DESC)` para garantizar rendimiento constante
- [ ] La respuesta incluye `next_cursor` (string opaco) o `null` si no hay más registros
- [ ] El tamaño de página por defecto es configurable via variable de entorno (`CAMPAIGN_PAGE_SIZE`, default 20)
- [ ] La consulta filtra campañas donde `cancelled_at IS NULL` y `closed_at IS NULL`
- [ ] Código revisado
- [ ] Cobertura de al menos 80% en el handler y la capa de repositorio
- [ ] Documentación actualizada (contrato de la API en OpenAPI/Swagger o equivalente)

---
*Nota para la IA: Mantener un lenguaje técnico preciso y evitar ambigüedades.*
