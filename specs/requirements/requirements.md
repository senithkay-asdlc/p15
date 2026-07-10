# Requirements: Hello World API

## 1. Overview

This project delivers a simple "Hello World" API: a minimal backend
service that demonstrates a working, deployable HTTP API on the
platform. It exists as a reference/starter service rather than to
satisfy complex business needs.

## 2. Goals

- Provide a single HTTP API that returns a friendly greeting message.
- Keep the service minimal, easy to understand, and easy to deploy.
- Demonstrate a healthy, observable service (liveness/health check).

## 3. Non-Goals

- No authentication or user management.
- No persistent storage of data.
- No UI/front-end is required for this project.

## 4. Functional Requirements

### 4.1 Hello World Endpoint

- The API MUST expose an HTTP GET endpoint (e.g. `/hello`) that
returns a JSON payload containing a greeting message, for example:
  ```json
  { "message": "Hello, World!" }
  ```
- The endpoint MUST optionally accept a `name` query parameter
(e.g. `/hello?name=Alice`) and, when provided, MUST personalize
the greeting, for example:
  ```json
  { "message": "Hello, Alice!" }
  ```
- If no `name` is provided, the API MUST default to a generic
greeting ("Hello, World!").

### 4.2 Health Check

- The API MUST expose a health check endpoint (e.g. `/health`) that
returns a success response when the service is running correctly,
suitable for liveness checks.

## 5. Non-Functional Requirements

- **Simplicity**: The service should have minimal dependencies and a
small, easy-to-read codebase.
- **Performance**: Responses should be returned in well under 1
second under normal conditions.
- **Availability**: The service should start quickly and remain
available; no external dependencies are required for it to function.
- **Statelessness**: The service must not require a database or any
persistent state.

## 6. Constraints

- The API is public (unauthenticated) — no login or API key is
required to call it.
- The API should be exposed over HTTP(S) following standard REST
conventions.

## 7. Success Criteria

- A client can issue `GET /hello` and receive a 200 response with a
greeting message.
- A client can issue `GET /hello?name=<name>` and receive a
personalized greeting.
- A client can issue `GET /health` and receive a 200 response
confirming the service is up.

