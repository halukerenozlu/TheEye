/**
 * Shared types for TheEye Dashboard
 * Matching the backend API contract
 */

export interface EventGeometry {
  longitude: number;
  latitude: number;
}

export interface Event {
  id: string;
  type: string;
  title: string;
  status: string;
  severity: number;
  started_at: string;
  updated_at: string;
  geometry?: EventGeometry;
}

export interface EventsListResponse {
  items: Event[];
  next_cursor: string;
}

export interface ErrorResponse {
  error: string;
  message: string;
}

export type SortOrder = "updated_at_desc" | "updated_at_asc";

export interface EventFilters {
  type?: string;
  started_after?: string;
  started_before?: string;
  sort?: SortOrder;
  limit?: number;
  cursor?: string;
}
