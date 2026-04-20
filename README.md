# URL Shortener (Go)

Proyecto de practica para aprender y reforzar **arquitectura hexagonal (Ports & Adapters)** en Go.

Es mi primer proyecto aplicando este estilo de arquitectura, asi que el foco principal esta en la separacion de responsabilidades entre dominio, casos de uso y adaptadores.

## Objetivo de aprendizaje

Este repositorio prioriza:

- Diseñar capas con dependencias hacia adentro.
- Mantener el dominio desacoplado de detalles de infraestructura.
- Practicar composicion explicita de dependencias en el `main`.
- Evolucionar de forma incremental (primero memoria, luego base de datos, etc.).

## Que hace hoy

- Crea shortlinks por HTTP.
- Resuelve shortlinks por codigo y redirige a la URL original.
- Guarda datos en memoria (sin persistencia real aun).

## Stack actual

- Go
- `net/http`
- `github.com/go-chi/chi/v5`
- `github.com/google/uuid`

## Estructura del proyecto

```text
cmd/
  api/            # punto de entrada

internal/
  domain/         # entidades, reglas de negocio y errores de dominio
  application/    # casos de uso
  adapters/
    httpa/        # handlers HTTP
    persistence/
      memory/     # repositorio en memoria
      postgres/   # placeholder para repositorio postgres (en progreso)
    system/       # clock, code generator, id generator
```

## Ejecutar localmente

```bash
go run ./cmd/api
```

Servidor en:

- `http://localhost:8080`

## Uso de la API

### 1) Crear shortlink

Metodo y endpoint:

- `POST http://localhost:8080/shortlinks`

Body:

```json
{
  "url": "https://www.facebook.com/"
}
```

Ejemplo con `curl`:

```bash
curl -X POST http://localhost:8080/shortlinks \
  -H "Content-Type: application/json" \
  -d '{"url":"https://www.facebook.com/"}'
```

Respuesta esperada (ejemplo):

- Status: `201 Created`

```json
{
  "id": "355e4e68-7ed9-4bdd-b296-f082eaf0f984",
  "originalUrl": "https://www.instagram.com/",
  "code": "DGwJNs",
  "createdAt": "2026-04-19T20:34:44.3672022-05:00",
  "expiresAt": null,
  "visitCount": 0
}
```

### 2) Resolver shortlink

Con el valor de `Code`, usa:

- `GET http://localhost:8080/{code}`

Ejemplo:

- `GET http://localhost:8080/DGwJNs`

Eso redirige a la URL original (por ejemplo, Instagram).

## Nota importante sobre expiracion

Aunque el modelo ya contempla `expiresAt`, en este estado del proyecto la creacion se hace sin expiracion (`nil`), por eso en la respuesta aparece `"expiresAt": null`.

## Estado actual

Este proyecto esta bien para aprendizaje y practica de arquitectura hexagonal. Aun asi, los siguientes pasos naturales serian:

- Implementar persistencia real (Postgres).
- Formalizar/mejorar puertos de entrada y salida para mantener consistencia arquitectonica.
- Agregar validaciones mas completas y mejor manejo de errores HTTP.
- Agregar tests de casos de uso y adaptadores.

## Licencia

Ver `LICENSE`.