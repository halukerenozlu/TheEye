import { afterEach, describe, expect, it, vi } from "vitest";

import { fetchEventDetail, fetchEvents } from "./api";

const BASE = "http://localhost:8080";

function mockFetch(response: Partial<Response> & { json?: () => unknown }) {
  const fetchMock = vi.fn().mockResolvedValue({
    ok: true,
    statusText: "OK",
    json: () => Promise.resolve({}),
    ...response,
  });
  vi.stubGlobal("fetch", fetchMock);
  return fetchMock;
}

afterEach(() => {
  vi.unstubAllGlobals();
});

describe("fetchEvents", () => {
  it("omits filters that were not provided", async () => {
    const fetchMock = mockFetch({});

    await fetchEvents();

    expect(fetchMock).toHaveBeenCalledWith(`${BASE}/v1/events?`, {
      cache: "no-store",
    });
  });

  it("serialises every supported filter", async () => {
    const fetchMock = mockFetch({});

    await fetchEvents({
      type: "earthquake",
      started_after: "2026-01-01T00:00:00Z",
      started_before: "2026-02-01T00:00:00Z",
      sort: "updated_at_asc",
      limit: 25,
      cursor: "42",
    });

    const url = new URL(fetchMock.mock.calls[0][0]);
    expect(Object.fromEntries(url.searchParams)).toEqual({
      type: "earthquake",
      started_after: "2026-01-01T00:00:00Z",
      started_before: "2026-02-01T00:00:00Z",
      sort: "updated_at_asc",
      limit: "25",
      cursor: "42",
    });
  });

  it("returns the decoded body", async () => {
    const body = { items: [], next_cursor: "" };
    mockFetch({ json: () => Promise.resolve(body) });

    await expect(fetchEvents()).resolves.toEqual(body);
  });

  it("throws when the response is not ok", async () => {
    mockFetch({ ok: false, statusText: "Bad Gateway" });

    await expect(fetchEvents()).rejects.toThrow("failed to fetch events: Bad Gateway");
  });
});

describe("fetchEventDetail", () => {
  it("requests the event by id", async () => {
    const fetchMock = mockFetch({});

    await fetchEventDetail("usgs:abc123");

    expect(fetchMock).toHaveBeenCalledWith(`${BASE}/v1/events/usgs:abc123`, {
      cache: "no-store",
    });
  });

  it("names the id in the error message", async () => {
    mockFetch({ ok: false, statusText: "Not Found" });

    await expect(fetchEventDetail("missing")).rejects.toThrow(
      "failed to fetch event detail for missing: Not Found",
    );
  });
});
