---
milestone_id: 1
---

# Generar JWT de acceso mediante credenciales

## Narrativa
**Como** usuario registrado en la plataforma
**Quiero** autenticarme con mi email y contraseña para obtener un token JWT
**Para** poder acceder a los recursos protegidos de la API sin reenviar mis credenciales en cada petición

## Criterios de Aceptación (Gherkin)

- **Escenario:** Autenticación exitosa con credenciales válidas
  - **Dado que** el usuario existe en el sistema y su contraseña es correcta
  - **Cuando** envía una petición `POST /auth/tokens` con `email` y `password` válidos
  - **Entonces** recibe una respuesta `201 OK` con un JWT firmado (HS256 o RS256) en el cuerpo, junto con el tiempo de expiración

- **Escenario:** Credenciales incorrectas
  - **Dado que** el usuario envía un email o contraseña incorrectos
  - **Cuando** realiza la petición `POST /auth/tokens`
  - **Entonces** recibe una respuesta `401 Unauthorized` con un mensaje genérico que no revela cuál campo es incorrecto

- **Escenario:** Campos requeridos ausentes
  - **Dado que** el cuerpo de la petición omite `email` o `password`
  - **Cuando** realiza la petición `POST /auth/tokens`
  - **Entonces** recibe una respuesta `400 Bad Request` con el detalle del campo faltante

- **Escenario:** Usuario inactivo o deshabilitado
  - **Dado que** el usuario existe pero su cuenta está deshabilitada
  - **Cuando** realiza la petición `POST /auth/tokens` con credenciales correctas
  - **Entonces** recibe una respuesta `401 Forbidden` indicando que la cuenta no está activa

## Definición de Hecho (DoD)
- [ ] Endpoint `POST /auth/tokens` implementado en Go
- [ ] Contraseña comparada usando bcrypt (sin comparación en texto plano)
- [ ] JWT contiene claims: `sub` (UUID del usuario), `email`, `exp`, `iat`
- [ ] Tiempo de expiración configurable via variable de entorno (`JWT_EXPIRATION`)
- [ ] Clave secreta de firma cargada desde variable de entorno (`JWT_SECRET`), nunca hardcodeada
- [ ] Código revisado
- [ ] Cobertura de al menos 80% en el handler y la capa de servicio
- [ ] Documentación actualizada (contrato de la API en OpenAPI/Swagger o equivalente)

---
*Nota para la IA: Mantener un lenguaje técnico preciso y evitar ambigüedades.*