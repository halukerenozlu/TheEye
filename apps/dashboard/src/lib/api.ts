import { EventFilters, EventsListResponse, Event } from "../types";

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

export async function fetchEvents(filters: EventFilters = {}): Promise<EventsListResponse> {
  const query = new URLSearchParams();
  
  if (filters.type) query.append("type", filters.type);
  if (filters.started_after) query.append("started_after", filters.started_after);
  if (filters.started_before) query.append("started_before", filters.started_before);
  if (filters.sort) query.append("sort", filters.sort);
  if (filters.limit) query.append("limit", filters.limit.toString());
  if (filters.cursor) query.append("cursor", filters.cursor);

  const url = `${API_BASE_URL}/v1/events?${query.toString()}`;
  
  const response = await fetch(url);
  if (!response.ok) {
    throw new Error(`failed to fetch events: ${response.statusText}`);
  }

  return response.json();
}

export async function fetchEventDetail(id: string): Promise<Event> {
  const url = `${API_BASE_URL}/v1/events/${id}`;
  
  const response = await fetch(url);
  if (!response.ok) {
    throw new Error(`failed to fetch event detail for ${id}: ${response.statusText}`);
  }

  return response.json();
}
