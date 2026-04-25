"use client";

import { useEffect, useState, useCallback, useRef } from "react";
import { Event, EventFilters } from "../types";
import { fetchEvents, fetchEventDetail } from "../lib/api";
import maplibregl from "maplibre-gl";

const EVENTS_POLL_INTERVAL_MS = 30000;

export default function DashboardPage() {
  // --- State ---
  const [events, setEvents] = useState<Event[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [isRefreshing, setIsRefreshing] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const [filters, setFilters] = useState<EventFilters>({
    sort: "updated_at_desc",
  });

  const [selectedEventId, setSelectedEventId] = useState<string | null>(null);
  const [selectedEvent, setSelectedEvent] = useState<Event | null>(null);
  const [isDetailLoading, setIsDetailLoading] = useState(false);
  const [isFilterOpen, setIsFilterOpen] = useState(false);
  const [isMapReady, setIsMapReady] = useState(false);

  // --- Refs ---
  const mapContainer = useRef<HTMLDivElement>(null);
  const map = useRef<maplibregl.Map | null>(null);
  const markers = useRef<Map<string, maplibregl.Marker>>(new Map());

  // --- Actions ---
  const loadEvents = useCallback(
    async (options?: { background?: boolean }) => {
      const isBackgroundRefresh = options?.background ?? false;

      if (isBackgroundRefresh) {
        setIsRefreshing(true);
      } else {
        setIsLoading(true);
        setError(null);
      }

      try {
        const data = await fetchEvents(filters);
        setEvents(data.items);
      } catch (err) {
        console.error(err);

        if (!isBackgroundRefresh || events.length === 0) {
          setError("Failed to sync with signal intelligence");
        }
      } finally {
        if (isBackgroundRefresh) {
          setIsRefreshing(false);
        } else {
          setIsLoading(false);
        }
      }
    },
    [events.length, filters],
  );

  useEffect(() => {
    loadEvents();
  }, [loadEvents]);

  useEffect(() => {
    const poll = () => {
      if (document.hidden) {
        return;
      }

      void loadEvents({ background: true });
    };

    const intervalId = window.setInterval(poll, EVENTS_POLL_INTERVAL_MS);

    return () => {
      window.clearInterval(intervalId);
    };
  }, [loadEvents]);

  useEffect(() => {
    if (!selectedEventId) {
      return;
    }

    const selectedEventFromList = events.find(
      (event) => event.id === selectedEventId,
    );

    if (!selectedEventFromList) {
      setSelectedEventId(null);
      setSelectedEvent(null);
      return;
    }

    setSelectedEvent((current) => {
      if (!current || current.id !== selectedEventId) {
        return current;
      }

      return {
        ...current,
        ...selectedEventFromList,
        geometry: selectedEventFromList.geometry ?? current.geometry,
      };
    });
  }, [events, selectedEventId]);

  // Load detail when selection changes
  useEffect(() => {
    if (!selectedEventId) {
      setSelectedEvent(null);
      return;
    }

    const loadDetail = async () => {
      setIsDetailLoading(true);
      try {
        const detail = await fetchEventDetail(selectedEventId);
        setSelectedEvent(detail);

        // Focus map if geometry is available
        if (map.current && detail.geometry) {
          map.current.easeTo({
            center: [detail.geometry.longitude, detail.geometry.latitude],
            zoom: 6,
            duration: 1500,
          });
        }
      } catch (err) {
        console.error("Failed to load event detail:", err);
      } finally {
        setIsDetailLoading(false);
      }
    };

    loadDetail();
  }, [selectedEventId]);

  // --- Map Initialization ---
  useEffect(() => {
    if (map.current || !mapContainer.current) return;

    const mapInstance = new maplibregl.Map({
      container: mapContainer.current,
      style: "https://basemaps.cartocdn.com/gl/dark-matter-gl-style/style.json",
      center: [0, 20],
      zoom: 1.5,
      attributionControl: false,
    });

    map.current = mapInstance;

    mapInstance.addControl(
      new maplibregl.NavigationControl({
        showCompass: false,
      }),
      "bottom-right",
    );

    mapInstance.on("load", () => {
      setIsMapReady(true);
      mapInstance.resize();
    });

    mapInstance.on("error", (e) => {
      console.error("Map error:", e);
    });

    return () => {
      mapInstance.remove();
      map.current = null;
      setIsMapReady(false);
    };
  }, []);

  // --- Marker Management ---
  useEffect(() => {
    const currentMap = map.current;
    if (!currentMap || !isMapReady) return;
    // Clear existing markers
    markers.current.forEach((m) => m.remove());
    markers.current.clear();

    const bounds = new maplibregl.LngLatBounds();
    let hasGeometry = false;

    // Add new markers for events with geometry
    events.forEach((event) => {
      if (!event.geometry) return;

      hasGeometry = true;
      const coords: [number, number] = [
        event.geometry.longitude,
        event.geometry.latitude,
      ];
      bounds.extend(coords);

      const isSelected = selectedEventId === event.id;
      const severityTone = getSeverityTone(event.severity);

      const el = document.createElement("div");
      el.className = "group relative cursor-pointer";

      const inner = document.createElement("div");
      inner.className = `rounded-full border transition-all duration-300 group-hover:scale-125 ${severityTone.markerClass} ${isSelected ? "h-4 w-4 shadow-[0_0_12px_rgba(255,255,255,0.4)] ring-2 ring-white/30" : severityTone.markerSizeClass}`;
      el.appendChild(inner);

      const marker = new maplibregl.Marker({ element: el })
        .setLngLat(coords)
        .addTo(currentMap);

      el.addEventListener("click", (e) => {
        e.stopPropagation();
        setSelectedEventId(event.id);
      });

      markers.current.set(event.id, marker);
    });

    // Only fit bounds on initial multi-event load if no selection is active
    if (hasGeometry && !selectedEventId && events.length > 0) {
      currentMap.fitBounds(bounds, {
        padding: 64,
        maxZoom: 10,
        duration: 2000,
      });
    }
  }, [events, isMapReady, selectedEventId]);

  const toggleSort = () => {
    setFilters((prev) => ({
      ...prev,
      sort:
        prev.sort === "updated_at_desc" ? "updated_at_asc" : "updated_at_desc",
    }));
  };

  const updateTypeFilter = (type: string) => {
    setFilters((prev) => ({
      ...prev,
      type: type || undefined,
    }));
  };

  return (
    <div className="flex h-screen w-full flex-col overflow-hidden bg-black font-sans text-zinc-300 antialiased selection:bg-zinc-800 selection:text-white">
      {/* 
        TOP BAR
      */}
      <header className="flex h-12 shrink-0 items-center justify-between border-b border-zinc-800 bg-zinc-950 px-4">
        <div className="flex items-center gap-4">
          <div className="text-sm font-black tracking-[0.2em] text-white">
            THE EYE
          </div>
          <div className="h-4 w-px bg-zinc-800" />
          <div className="text-[10px] font-medium text-zinc-500 uppercase tracking-widest">
            World Monitoring <span className="mx-1 opacity-30">|</span> v0.3.0
          </div>
        </div>

        <div className="flex items-center gap-4">
          <div className="flex items-center gap-2 rounded-full border border-zinc-800 bg-zinc-900/50 px-2.5 py-1">
            <div
              className={`h-1.5 w-1.5 rounded-full shadow-[0_0_8px_rgba(16,185,129,0.5)] ${isLoading || isRefreshing ? "bg-amber-500 animate-pulse" : "bg-emerald-500"}`}
            />
            <span className="text-[9px] font-semibold text-zinc-400 uppercase tracking-tighter">
              {isLoading
                ? "Syncing Signals..."
                : isRefreshing
                  ? "Refreshing Feed..."
                  : "System Ready"}
            </span>
          </div>
        </div>
      </header>

      {/* 
        MAIN CONTENT AREA
      */}
      <div className="flex flex-1 overflow-hidden">
        {/* 
          LEFT PANEL: SIGNAL FEED
        */}
        <aside className="flex w-72 shrink-0 flex-col border-r border-zinc-800 bg-zinc-900/20 backdrop-blur-md">
          <div className="flex items-center justify-between border-b border-zinc-800/50 px-4 py-3">
            <h2 className="text-[10px] font-bold uppercase tracking-[0.2em] text-zinc-500">
              Signal Feed
            </h2>
            <div className="flex gap-1.5">
              <button
                onClick={() => void loadEvents()}
                title="Refresh feed"
                className="h-1.5 w-1.5 rounded-full bg-zinc-700 hover:bg-zinc-500 transition-colors"
              />
              <button
                onClick={toggleSort}
                title={`Sort: ${filters.sort === "updated_at_desc" ? "Newest" : "Oldest"}`}
                className={`h-1.5 w-1.5 rounded-full ${filters.sort === "updated_at_asc" ? "bg-emerald-500" : "bg-zinc-700"} hover:opacity-80 transition-colors`}
              />
            </div>
          </div>

          <div className="flex-1 overflow-y-auto">
            {isLoading && events.length === 0 ? (
              <div className="flex h-full flex-col items-center justify-center p-8 text-center space-y-4">
                <div className="relative h-px w-24 overflow-hidden bg-zinc-900">
                  <div className="h-full w-1/3 animate-[loading_1.5s_infinite_ease-in-out] bg-zinc-700" />
                </div>
                <span className="text-[9px] font-mono uppercase tracking-[0.3em] text-zinc-600">
                  Acquiring Stream...
                </span>
              </div>
            ) : error ? (
              <div className="flex h-full flex-col items-center justify-center p-8 text-center">
                <div className="mb-3 flex h-10 w-10 items-center justify-center rounded-full border border-rose-900/30 bg-rose-950/10 text-rose-900">
                  <svg
                    width="18"
                    height="18"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    strokeWidth="1.5"
                  >
                    <circle cx="12" cy="12" r="10" />
                    <line x1="12" y1="8" x2="12" y2="12" />
                    <line x1="12" y1="16" x2="12.01" y2="16" />
                  </svg>
                </div>
                <p className="text-[10px] font-medium uppercase tracking-widest text-rose-900/80">
                  Signal Sync Failed
                </p>
                <button
                  onClick={() => void loadEvents()}
                  className="mt-4 text-[9px] font-bold uppercase tracking-tighter text-zinc-500 hover:text-zinc-300 transition-colors underline underline-offset-4"
                >
                  Retry Connection
                </button>
              </div>
            ) : events.length === 0 ? (
              <div className="flex h-full flex-col items-center justify-center p-8 text-center text-zinc-700">
                <div className="mb-4 h-12 w-12 opacity-20 grayscale filter">
                  <svg
                    width="48"
                    height="48"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    strokeWidth="0.5"
                  >
                    <path d="M2 12h3l2 8 4-16 4 16 2-8h3" />
                  </svg>
                </div>
                <p className="text-[9px] uppercase tracking-[0.2em] font-medium text-zinc-600">
                  No active signals detected
                </p>
              </div>
            ) : (
              <div className="divide-y divide-zinc-800/20">
                {events.map((event) => (
                  <button
                    key={event.id}
                    onClick={() => setSelectedEventId(event.id)}
                    className={`flex w-full flex-col gap-1.5 px-4 py-4 text-left transition-all duration-200 border-l-2 ${selectedEventId === event.id ? "bg-zinc-800/30 border-emerald-500/50" : "hover:bg-zinc-800/10 border-transparent"}`}
                  >
                    {(() => {
                      const severityTone = getSeverityTone(event.severity);
                      return (
                        <>
                          <div className="flex items-center justify-between">
                            <span
                              className={`text-[8px] font-mono font-bold uppercase tracking-tighter ${selectedEventId === event.id ? "text-emerald-500" : "text-zinc-500"}`}
                            >
                              {event.type}
                            </span>
                            <span className="text-[8px] font-mono text-zinc-600">
                              {new Date(event.started_at).toLocaleTimeString(
                                [],
                                {
                                  hour: "2-digit",
                                  minute: "2-digit",
                                  second: "2-digit",
                                  hour12: false,
                                },
                              )}
                            </span>
                          </div>
                          <h3
                            className={`line-clamp-2 text-[11px] font-medium leading-snug transition-colors ${selectedEventId === event.id ? "text-white" : "text-zinc-400 group-hover:text-zinc-300"}`}
                          >
                            {event.title}
                          </h3>
                          <div className="mt-0.5 flex items-center gap-2">
                            <div
                              className={`h-1.5 w-1.5 rounded-full ${severityTone.dotClass}`}
                            />
                            <span
                              className={`rounded border px-1.5 py-px text-[8px] font-bold uppercase tracking-widest ${severityTone.badgeClass}`}
                            >
                              Lvl {clampSeverityLevel(event.severity)}
                            </span>
                            <span className="text-[8px] font-bold text-zinc-600 uppercase tracking-widest">
                              {event.status}
                            </span>
                          </div>
                        </>
                      );
                    })()}
                  </button>
                ))}
              </div>
            )}
          </div>

          <div className="border-t border-zinc-800/50 p-3 bg-zinc-950/40 relative">
            <button
              onClick={() => setIsFilterOpen(!isFilterOpen)}
              className="flex h-7 w-full items-center justify-between rounded border border-zinc-800 bg-zinc-900/50 px-2 text-[9px] text-zinc-400 uppercase tracking-widest outline-none hover:border-zinc-700 transition-colors"
            >
              <span className="truncate">
                {filters.type ? filters.type + "s" : "All Categories"}
              </span>
              <svg
                width="10"
                height="10"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                strokeWidth="2"
                className={`shrink-0 transition-transform duration-200 ${isFilterOpen ? "rotate-180" : ""}`}
              >
                <path d="m6 9 6 6 6-6" />
              </svg>
            </button>

            {isFilterOpen && (
              <>
                <div
                  className="fixed inset-0 z-10"
                  onClick={() => setIsFilterOpen(false)}
                />
                <div className="absolute bottom-11 left-3 right-3 z-20 overflow-hidden rounded border border-zinc-800 bg-zinc-950 shadow-2xl animate-in fade-in slide-in-from-bottom-2 duration-200">
                  {[
                    { label: "All Categories", value: "" },
                    { label: "Earthquakes", value: "earthquake" },
                    { label: "Wildfires", value: "wildfire" },
                    { label: "Storms", value: "storm" },
                  ].map((option) => (
                    <button
                      key={option.value}
                      onClick={() => {
                        updateTypeFilter(option.value);
                        setIsFilterOpen(false);
                      }}
                      className={`flex h-8 w-full items-center px-3 text-[9px] uppercase tracking-widest transition-colors hover:bg-zinc-900 ${filters.type === option.value || (!filters.type && !option.value) ? "text-white font-bold bg-zinc-900" : "text-zinc-500"}`}
                    >
                      {option.label}
                    </button>
                  ))}
                </div>
              </>
            )}
          </div>
        </aside>

        {/* 
          CENTER PANEL: MAP VIEWPORT
        */}
        <main className="relative min-h-0 flex-1 bg-zinc-950 overflow-hidden">
          <div
            ref={mapContainer}
            className="theeye-map absolute inset-0 h-full w-full"
          />

          {/* Map Attribution/Status Overlay */}
          <div className="absolute bottom-2 left-4 pointer-events-none">
            <span className="text-[8px] font-mono text-zinc-700 uppercase tracking-widest">
              Engine: MapLibre GL <span className="mx-1 opacity-30">|</span>{" "}
              Style: Dark Matter
            </span>
          </div>
        </main>

        {/* 
          RIGHT PANEL: EVENT INTELLIGENCE
        */}
        <aside className="flex w-96 shrink-0 flex-col border-l border-zinc-800 bg-zinc-900/20 backdrop-blur-md">
          <div className="flex items-center justify-between border-b border-zinc-800/50 px-4 py-3">
            <h2 className="text-[10px] font-bold uppercase tracking-[0.2em] text-zinc-500">
              Event Intelligence
            </h2>
            {selectedEventId && (
              <button
                onClick={() => setSelectedEventId(null)}
                className="text-[9px] font-bold uppercase tracking-widest text-zinc-600 hover:text-zinc-400 transition-colors"
              >
                Clear
              </button>
            )}
          </div>

          <div className="flex-1 overflow-y-auto p-6">
            {isDetailLoading ? (
              <div className="flex h-full flex-col items-center justify-center space-y-4">
                <div className="relative h-0.5 w-16 overflow-hidden bg-zinc-900">
                  <div className="h-full w-1/3 animate-[loading_1s_infinite_ease-in-out] bg-zinc-600" />
                </div>
                <span className="text-[9px] font-mono uppercase tracking-[0.3em] text-zinc-600">
                  Analyzing Entity...
                </span>
              </div>
            ) : selectedEvent ? (
              (() => {
                const severityTone = getSeverityTone(selectedEvent.severity);
                return (
                  <div className="flex flex-col gap-8 animate-in fade-in slide-in-from-right-4 duration-500">
                    <div className="space-y-2">
                      <div className="flex items-center gap-2">
                        <div
                          className={`h-1.5 w-1.5 rounded-full ${severityTone.dotClass}`}
                        />
                        <span className="text-[8px] font-mono font-bold text-zinc-600 uppercase tracking-widest">
                          Signal ID: {selectedEvent.id}
                        </span>
                      </div>
                      <h2 className="text-xl font-semibold leading-tight text-white tracking-tight">
                        {selectedEvent.title}
                      </h2>
                    </div>

                    <div className="grid grid-cols-2 gap-y-6 gap-x-4">
                      <div className="flex flex-col gap-1.5">
                        <span className="text-[8px] font-bold text-zinc-600 uppercase tracking-[0.15em]">
                          Status
                        </span>
                        <span className="text-[11px] text-zinc-300 capitalize flex items-center gap-2">
                          <span className="h-1 w-1 rounded-full bg-zinc-500" />
                          {selectedEvent.status}
                        </span>
                      </div>
                      <div className="flex flex-col gap-1.5">
                        <span className="text-[8px] font-bold text-zinc-600 uppercase tracking-[0.15em]">
                          Severity
                        </span>
                        <span
                          className={`text-[11px] font-semibold ${severityTone.textClass}`}
                        >
                          Level {clampSeverityLevel(selectedEvent.severity)}
                        </span>
                      </div>
                      <div className="flex flex-col gap-1.5">
                        <span className="text-[8px] font-bold text-zinc-600 uppercase tracking-[0.15em]">
                          Category
                        </span>
                        <span className="text-[11px] text-zinc-300 capitalize">
                          {selectedEvent.type}
                        </span>
                      </div>
                      <div className="flex flex-col gap-1.5">
                        <span className="text-[8px] font-bold text-zinc-600 uppercase tracking-[0.15em]">
                          Temporal Mark
                        </span>
                        <span className="text-[11px] text-zinc-300">
                          {new Date(
                            selectedEvent.started_at,
                          ).toLocaleTimeString([], {
                            hour: "2-digit",
                            minute: "2-digit",
                            second: "2-digit",
                          })}
                        </span>
                      </div>
                    </div>

                    <div className="h-px w-full bg-zinc-800/40" />

                    <div className="space-y-4">
                      <div className="flex flex-col gap-2">
                        <span className="text-[8px] font-bold text-zinc-600 uppercase tracking-[0.15em]">
                          Geospatial Position
                        </span>
                        <div className="rounded border border-zinc-800/50 bg-zinc-950/50 p-3">
                          <span className="text-[10px] font-mono text-emerald-500/80">
                            {selectedEvent.geometry
                              ? `${selectedEvent.geometry.latitude.toFixed(6)}°N, ${selectedEvent.geometry.longitude.toFixed(6)}°E`
                              : "Position data inaccessible"}
                          </span>
                        </div>
                      </div>
                    </div>
                  </div>
                );
              })()
            ) : (
              <div className="flex h-full flex-col items-center justify-center text-center space-y-6">
                <div className="h-px w-6 bg-zinc-800" />
                <p className="max-w-50 text-[10px] leading-relaxed text-zinc-600 uppercase tracking-widest font-medium">
                  Select a signal from the feed or map to inspect metadata and
                  intelligence
                </p>
                <div className="h-px w-6 bg-zinc-800" />
              </div>
            )}
          </div>

          <div className="border-t border-zinc-800/50 p-4 bg-zinc-950/20">
            <div className="flex gap-2">
              <div className="h-1.5 w-1.5 rounded-full bg-zinc-800" />
              <div className="h-1.5 w-1.5 rounded-full bg-zinc-800" />
              <div className="h-1.5 w-1.5 rounded-full bg-zinc-800" />
            </div>
          </div>
        </aside>
      </div>
    </div>
  );
}

type SeverityTone = {
  dotClass: string;
  markerClass: string;
  markerSizeClass: string;
  badgeClass: string;
  textClass: string;
};

function clampSeverityLevel(severity: number): 1 | 2 | 3 {
  if (severity >= 3) return 3;
  if (severity === 2) return 2;
  return 1;
}

function getSeverityTone(severity: number): SeverityTone {
  const level = clampSeverityLevel(severity);

  if (level === 3) {
    return {
      dotClass: "bg-rose-500 shadow-[0_0_6px_rgba(244,63,94,0.6)]",
      markerClass:
        "bg-rose-500 border-rose-300/60 shadow-[0_0_10px_rgba(244,63,94,0.75)]",
      markerSizeClass: "h-3.5 w-3.5",
      badgeClass: "border-rose-400/50 bg-rose-500/15 text-rose-300",
      textClass: "text-rose-300",
    };
  }

  if (level === 2) {
    return {
      dotClass: "bg-amber-400 shadow-[0_0_5px_rgba(251,191,36,0.55)]",
      markerClass:
        "bg-amber-400 border-amber-200/60 shadow-[0_0_8px_rgba(251,191,36,0.6)]",
      markerSizeClass: "h-3 w-3",
      badgeClass: "border-amber-300/40 bg-amber-400/10 text-amber-200",
      textClass: "text-amber-200",
    };
  }

  return {
    dotClass: "bg-emerald-500 shadow-[0_0_4px_rgba(16,185,129,0.45)]",
    markerClass:
      "bg-emerald-500 border-emerald-300/50 shadow-[0_0_6px_rgba(16,185,129,0.5)]",
    markerSizeClass: "h-2.5 w-2.5",
    badgeClass: "border-emerald-400/35 bg-emerald-500/10 text-emerald-200",
    textClass: "text-emerald-200",
  };
}
