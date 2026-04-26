# TheEye API

## Current API Baseline

Base version prefix:

```text
/v1
```

Current active endpoints:

- `GET /v1/healthz`
- `GET /v1/readyz`
- `GET /v1/meta`
- `GET /v1/events`
- `GET /v1/events/{id}`

Planned later, not active in the current baseline:

- SSE event stream
- event changes endpoint
- alert endpoints
- saved view endpoints

---

## Health and Metadata

### `GET /v1/healthz`

Basic process health endpoint.

### `GET /v1/readyz`

Readiness endpoint for local/service checks.

### `GET /v1/meta`

Service metadata endpoint.

---

## Events List

### `GET /v1/events`

Returns stored normalized events.

Response shape:

```json
{
  "items": [],
  "next_cursor": ""
}
```

Fields:

- `items`: array of Event objects
- `next_cursor`: cursor for the next page, or an empty string when no next page exists

### Supported Filters

Current filters:

- `type`
- `started_after`
- `started_before`

Filter behavior:

- filters are optional
- invalid query values return consistent JSON `400` errors
- default behavior without filters returns stored events according to default list behavior

### Supported Sorting

Current sort values:

- `updated_at_desc`
- `updated_at_asc`

Invalid sort values return consistent JSON `400` errors.

### Supported Pagination

Current pagination parameters:

- `limit`
- `cursor`

Pagination behavior:

- `next_cursor` is an empty string when there is no next page
- invalid cursor usage returns consistent JSON `400` errors
- the response shape remains stable across paginated and unpaginated requests

---

## Event Detail

### `GET /v1/events/{id}`

Returns one stored normalized event when found.

Unknown ids return a consistent JSON `404` error shape.

---

## Event Object

Current Event fields:

- `id`
- `type`
- `category`
- `title`
- `status`
- `severity`
- `severity_level`
- `started_at`
- `updated_at`
- `geometry`
- `source`

Field notes:

- `type` is the specific event type, such as `earthquake`.
- `category` is the broader normalized grouping, currently `natural_disaster` for USGS earthquakes.
- `severity` is retained for compatibility and source-readable severity behavior.
- `severity_level` is the backend-owned normalized display level for the dashboard.
- `geometry` is used by the dashboard for marker rendering.
- `source` identifies the source context exposed by the API.

---

## Error Shape

The current baseline uses consistent JSON errors for route, method, query, and not-found cases.

Representative shape:

```json
{
  "error": "event_not_found",
  "message": "event not found"
}
```

Exact error codes may vary by route and validation failure, but errors should remain structured JSON.

---

## Planned API Work

The following capabilities are planned but not active in the current baseline:

- SSE endpoint for realtime event stream
- event changes endpoint
- source/category filters beyond the current MVP filters
- bbox/geospatial filters
- alert and saved-view endpoints

Planned endpoints should not be documented as active until implemented.
