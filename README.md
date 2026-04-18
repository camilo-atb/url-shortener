# URL Shortener (Go) — Proyecto de aprendizaje (Arquitectura Hexagonal)

Proyecto de aprendizaje en **Go** para construir un **servicio de acortamiento de URLs** aplicando **arquitectura hexagonal (Ports & Adapters)**, con foco en separar el **dominio** y los **casos de uso** de los detalles de infraestructura.

> Objetivo: practicar diseño limpio, dependencias hacia adentro y evolución incremental del proyecto (sin “magia” de frameworks).

---

## Qué hace hoy

- Expone un API HTTP para:
  - **Crear** un shortlink.
  - **Resolver** un shortlink y redirigir a la URL original.
- Usa **persistencia en memoria** (todavía **sin base de datos**).
- Usa adaptadores de sistema para dependencias como:
  - reloj (clock),
  - generador de códigos,
  - generador de IDs.

---

## Stack / Dependencias

- **Go**
- HTTP con `net/http`
- Router: `chi` (ligero)
- Sin frameworks “pesados” y sin ORM (por ahora).

---

## Arquitectura (Hexagonal)

La estructura principal sigue la idea de separar:

- **Domain**: reglas/entidades del negocio.
- **Application**: casos de uso (use cases).
- **Ports**: interfaces que definen lo que la app necesita del mundo exterior (entrada/salida).
- **Adapters**: implementaciones concretas (HTTP, memoria, sistema, etc.).

Estructura (alto nivel):

```
cmd/
  api/            # punto de entrada (main)
internal/
  application/    # casos de uso
  domain/         # modelo y lógica de dominio
  ports/
    input/        # puertos de entrada (contratos hacia la app)
    output/       # puertos de salida (repositorio, reloj, generator, etc.)
  adapters/
    httpa/        # adaptador HTTP (handlers)
    persistence/
      memory/     # repositorio en memoria
    system/       # clock / code generator / id generator
```

> Nota honesta (aprendizaje): me di cuenta tarde de que me faltó **formalizar mejor los puertos de entrada**. Como el objetivo del repo es aprender, queda documentado como parte del proceso y próximos pasos.

---

## Cómo ejecutar

1. Clona el repositorio
2. Ejecuta el API:

```bash
go run ./cmd/api
```

El servidor inicia en:

- `http://localhost:8080`

---

## Endpoints

- `POST /shortlinks` → crea un shortlink
- `GET /{code}` → resuelve el shortlink

(La forma exacta del body/response puede variar; la intención principal es practicar la separación por capas y el flujo de dependencias.)

---

## Estado actual y próximos pasos (idea)

- [ ] Agregar persistencia real (por ejemplo Postgres/SQLite) mediante un adaptador de `ports/output`
- [ ] Completar/ajustar el modelado de **puertos de entrada**
- [ ] Mejorar validaciones y manejo de errores
- [ ] Tests de casos de uso y adaptadores

---

## Motivación

Este repositorio existe para practicar:
- arquitectura hexagonal en un proyecto pequeño pero real,
- composición explícita de dependencias en `main`,
- adaptadores intercambiables (memoria hoy, DB mañana),
- claridad de diseño antes que complejidad.

---

## Licencia

Ver archivo `LICENSE`.