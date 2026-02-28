# Crear campaña de donación

## Narrativa
**Como** usuario registrado en la plataforma
**Quiero** crear una nueva campaña de donación proporcionando su información básica
**Para** poder recibir donaciones de otros usuarios hacia mi causa o proyecto

## Criterios de Aceptación (Gherkin)

- **Escenario:** Creación exitosa de campaña con datos válidos
  - **Dado que** el usuario está autenticado con un JWT válido
  - **Cuando** envía una petición `POST /campaigns` con `title`, `description`, `goal_amount` y `end_date` válidos
  - **Entonces** recibe una respuesta `201 Created` con los datos de la campaña creada, incluyendo su `id` (UUID)

- **Escenario:** Creación rechazada sin autenticación
  - **Dado que** el usuario no incluye un JWT o el token es inválido
  - **Cuando** envía una petición `POST /campaigns`
  - **Entonces** recibe una respuesta `401 Unauthorized`

- **Escenario:** Campos requeridos ausentes
  - **Dado que** el usuario está autenticado
  - **Cuando** envía una petición `POST /campaigns` omitiendo uno o más campos obligatorios (`title`, `description`, `goal_amount`, `end_date`)
  - **Entonces** recibe una respuesta `400 Bad Request` con el detalle de cada campo faltante o inválido

- **Escenario:** Monto objetivo inválido
  - **Dado que** el usuario está autenticado
  - **Cuando** envía una petición `POST /campaigns` con `goal_amount` igual a cero o negativo
  - **Entonces** recibe una respuesta `400 Bad Request` indicando que el monto debe ser mayor a cero

- **Escenario:** Fecha de fin en el pasado
  - **Dado que** el usuario está autenticado
  - **Cuando** envía una petición `POST /campaigns` con una `end_date` anterior a la fecha actual
  - **Entonces** recibe una respuesta `400 Bad Request` indicando que la fecha de fin debe ser futura

## Definición de Hecho (DoD)
- [ ] Endpoint `POST /campaigns` implementado con middleware de autenticación JWT
- [ ] El campo `owner_id` se asigna automáticamente desde el claim `sub` del JWT, sin aceptarlo del cuerpo de la petición
- [ ] Validaciones aplicadas: `title` no vacío, `description` no vacía, `goal_amount` > 0, `end_date` futura
- [ ] Archivo de migración actualizado o creado para la tabla `campaigns` con los campos: `id` (UUID PRIMARY KEY), `title`, `description`, `goal_amount`, `raised_amount` (DEFAULT 0), `owner_id` (UUID REFERENCES users(id)), `end_date`, `created_at`, `updated_at`
- [ ] Código revisado
- [ ] Cobertura de al menos 80% en el handler y la capa de servicio
- [ ] Documentación actualizada (contrato de la API en OpenAPI/Swagger o equivalente)

---
*Nota para la IA: Mantener un lenguaje técnico preciso y evitar ambigüedades.*
