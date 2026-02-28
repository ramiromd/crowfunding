# Cierre programado de campañas expiradas

## Narrativa
**Como** sistema de la plataforma
**Quiero** detectar y cerrar automáticamente las campañas cuya `end_date` ha sido alcanzada
**Para** garantizar que ninguna campaña permanezca activa más allá de su fecha de fin sin intervención manual

## Criterios de Aceptación (Gherkin)

- **Escenario:** Cierre exitoso de campaña expirada
  - **Dado que** existe una campaña activa cuya `end_date` es anterior o igual a la fecha y hora actuales
  - **Cuando** se ejecuta el proceso de cierre programado
  - **Entonces** la campaña queda marcada como `closed` y se registra la fecha y hora exacta del cierre en `closed_at`

- **Escenario:** Campañas no expiradas no se ven afectadas
  - **Dado que** existen campañas activas cuya `end_date` es posterior a la fecha y hora actuales
  - **Cuando** se ejecuta el proceso de cierre programado
  - **Entonces** dichas campañas no son modificadas

- **Escenario:** Campañas ya cerradas o canceladas no son reprocesadas
  - **Dado que** existen campañas con `closed_at` no nulo o en estado cancelado
  - **Cuando** se ejecuta el proceso de cierre programado
  - **Entonces** dichas campañas no son modificadas ni se genera un error

- **Escenario:** Cierre de múltiples campañas en una misma ejecución
  - **Dado que** existen varias campañas activas expiradas simultáneamente
  - **Cuando** se ejecuta el proceso de cierre programado
  - **Entonces** todas las campañas expiradas quedan cerradas en la misma ejecución

- **Escenario:** Fallo de la base de datos durante el cierre
  - **Dado que** ocurre un error de base de datos durante la actualización
  - **Cuando** se ejecuta el proceso de cierre programado
  - **Entonces** la transacción es revertida, ninguna campaña queda en estado inconsistente y el error queda registrado en el log del sistema

## Definición de Hecho (DoD)
- [ ] Proceso de cierre implementado como job programado (cron) en Go con frecuencia configurable via variable de entorno (`CAMPAIGN_CLOSE_CRON`)
- [ ] La tabla `campaigns` incluye el campo `closed_at` (TIMESTAMPTZ, nullable); un valor no nulo indica que la campaña está cerrada
- [ ] La consulta de cierre selecciona únicamente campañas donde `end_date <= now()` y `closed_at IS NULL` y `cancelled_at IS NULL`
- [ ] La actualización se ejecuta dentro de una transacción para garantizar consistencia
- [ ] Los errores del job quedan registrados en el sistema de logging con nivel `ERROR`
- [ ] Código revisado
- [ ] Cobertura de al menos 80% en la lógica del job y la capa de repositorio
- [ ] Documentación actualizada (descripción del job en `docs/`)

---
*Nota para la IA: Mantener un lenguaje técnico preciso y evitar ambigüedades.*
