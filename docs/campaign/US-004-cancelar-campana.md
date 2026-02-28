# Cancelar campaña de donación

## Narrativa
**Como** usuario registrado propietario de una campaña
**Quiero** poder cancelar una campaña que he creado
**Para** detener la recepción de donaciones cuando la campaña ya no es viable o pertinente

## Criterios de Aceptación (Gherkin)

- **Escenario:** Cancelación exitosa por el propietario
  - **Dado que** el usuario está autenticado y es propietario de la campaña
  - **Cuando** envía una petición `DELETE /campaigns/{id}`
  - **Entonces** recibe una respuesta `204 No Content` y la campaña queda registrada como cancelada

- **Escenario:** Cancelación rechazada sin autenticación
  - **Dado que** el usuario no incluye un JWT o el token es inválido
  - **Cuando** envía una petición `DELETE /campaigns/{id}`
  - **Entonces** recibe una respuesta `401 Unauthorized`

- **Escenario:** Cancelación rechazada por usuario no propietario
  - **Dado que** el usuario está autenticado pero no es el propietario de la campaña
  - **Cuando** envía una petición `DELETE /campaigns/{id}`
  - **Entonces** recibe una respuesta `403 Forbidden`

- **Escenario:** Campaña no encontrada
  - **Dado que** el usuario está autenticado
  - **Cuando** envía una petición `DELETE /campaigns/{id}` con un `id` inexistente
  - **Entonces** recibe una respuesta `404 Not Found`

## Definición de Hecho (DoD)
- [ ] Endpoint `DELETE /campaigns/{id}` implementado con middleware de autenticación JWT, donde `{id}` es un UUID válido
- [ ] La autorización verifica que el `owner_id` de la campaña coincide con el claim `sub` del JWT
- [ ] La cancelación realiza un soft delete (marca la campaña como cancelada) en lugar de eliminar el registro
- [ ] Código revisado
- [ ] Cobertura de al menos 80% en el handler y la capa de servicio
- [ ] Documentación actualizada (contrato de la API en OpenAPI/Swagger o equivalente)

---
*Nota para la IA: Mantener un lenguaje técnico preciso y evitar ambigüedades.*
