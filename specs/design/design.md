---
skillsApplied:
  - high-level-architecture
  - go
  - openapi-conventions
---

# Design: Hello World API

## 1. Overview

A minimal, stateless HTTP API that returns a friendly (optionally
personalized) greeting and exposes a health check, deployed as a
single public backend service.

## 2. Components

- **hello-api** (`service`) — exposes `GET /hello` and `GET /health`;
  the sole component in this system.

## 3. Capabilities

### hello-api

- **Greeting** — returns a JSON greeting message; accepts an optional
  `name` query parameter to personalize the message, defaulting to
  "Hello, World!" when absent.
- **Health check** — returns a 200 response indicating the service is
  live, for liveness probes.

## 4. Data model

No persistent entities. The only payload shape is the greeting
response:

- **Greeting**: `message` (string) — the greeting text returned to
  the caller.

## 5. Roles & access

- **Anonymous caller** — any client, unauthenticated. No roles or
  per-user data exist in this system.

## 6. Interactions

None — `hello-api` has no dependencies on other components, org
services, external systems, or platform resources. It is a single,
self-contained, stateless service.

## 7. Data flow

1. Client issues `GET /hello` (optionally with `?name=<name>`).
2. `hello-api` builds the greeting message (personalized if `name`
   was supplied, generic otherwise) and returns it as JSON with a 200
   status.
3. Client (or platform liveness probe) issues `GET /health` and
   receives a 200 response confirming the service is running.
