"use client";

import { useEffect, useState, useCallback, useRef } from "react";
import { Event, EventFilters } from "../types";
import { fetchEvents } from "../lib/api";
import maplibregl from "maplibre-gl";

export default function DashboardPage() {
  // --- State ---
  const [events, setEvents] = useState<Event[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const [filters, setFilters] = useState<EventFilters>({
    sort: "updated_at_desc",
  });

  const [selectedEventId, setSelectedEventId] = useState<string | null>(null);
  const [isFilterOpen, setIsFilterOpen] = useState(false);
  const [isMapReady, setIsMapReady] = useState(false);

  // --- Refs ---
  const mapContainer = useRef<HTMLDivElement>(null);
  const map = useRef<maplibregl.Map | null>(null);
  const markers = useRef<maplibregl.Marker[]>([]);

  // --- Actions ---
  const loadEvents = useCallback(async () => {
    setIsLoading(true);
    setError(null);
    try {
      const data = await fetchEvents(filters);
      setEvents(data.items);
    } catch (err) {
      console.error(err);
      setError("Failed to sync with signal intelligence");
    } finally {
      setIsLoading(false);
    }
  }, [filters]);

  useEffect(() => {
    loadEvents();
  }, [loadEvents]);

  // --- Map Initialization ---
  useEffect(() => {
    if (map.current || !mapContainer.current) return;

    console.log("Initializing map...");
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
      console.log("Map loaded successfully");
      setIsMapReady(true);
      mapInstance.resize();
    });

    mapInstance.on("error", (e) => {
      console.error("Map error:", e);
    });

    return () => {
      console.log("Cleaning up map...");
      mapInstance.remove();
      map.current = null;
      setIsMapReady(false);
    };
  }, []);

  // --- Marker Management ---
  useEffect(() => {
    const currentMap = map.current;
    if (!currentMap || !isMapReady) {
      console.log("Map not ready for markers", { hasMap: !!currentMap, isMapReady });
      return;
    }

    console.log("Syncing markers...", events.length);
    // Clear existing markers
    markers.current.forEach((m) => m.remove());
    markers.current = [];

    const bounds = new maplibregl.LngLatBounds();
    let hasGeometry = false;

    // Add new markers for events with geometry
    events.forEach((event) => {
      if (!event.geometry) return;

      hasGeometry = true;
      const coords: [number, number] = [event.geometry.longitude, event.geometry.latitude];
      bounds.extend(coords);

      const el = document.createElement("div");
      el.className = "group relative cursor-pointer";

      const inner = document.createElement("div");
      inner.className = `h-2.5 w-2.5 rounded-full border border-white/20 transition-transform duration-200 group-hover:scale-125 ${getSeverityColor(event.severity)}`;
      el.appendChild(inner);

      const marker = new maplibregl.Marker({ element: el })
        .setLngLat(coords)
        .addTo(currentMap);

      el.addEventListener("click", () => {
        setSelectedEventId(event.id);
      });

      markers.current.push(marker);
    });

    if (hasGeometry) {
      currentMap.fitBounds(bounds, {
        padding: 64,
        maxZoom: 10,
        duration: 2000,
      });
    }
  }, [events, isMapReady]);

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
              className={`h-1.5 w-1.5 rounded-full shadow-[0_0_8px_rgba(16,185,129,0.5)] ${isLoading ? "bg-amber-500 animate-pulse" : "bg-emerald-500"}`}
            />
            <span className="text-[9px] font-semibold text-zinc-400 uppercase tracking-tighter">
              {isLoading ? "Syncing Signals..." : "System Ready"}
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
                onClick={loadEvents}
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
              <div className="flex h-full items-center justify-center p-8 text-center">
                <span className="text-[9px] font-mono uppercase tracking-[0.3em] text-zinc-700 animate-pulse">
                  Acquiring Stream...
                </span>
              </div>
            ) : error ? (
              <div className="flex h-full flex-col items-center justify-center p-8 text-center">
                <div className="mb-2 text-rose-900">
                  <svg
                    width="24"
                    height="24"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    strokeWidth="1"
                  >
                    <circle cx="12" cy="12" r="10" />
                    <line x1="12" y1="8" x2="12" y2="12" />
                    <line x1="12" y1="16" x2="12.01" y2="16" />
                  </svg>
                </div>
                <p className="text-[9px] uppercase tracking-widest text-rose-900/80">
                  {error}
                </p>
              </div>
            ) : events.length === 0 ? (
              <div className="flex h-full flex-col items-center justify-center p-8 text-center text-zinc-700">
                <p className="text-[9px] uppercase tracking-[0.2em]">
                  No active signals detected
                </p>
              </div>
            ) : (
              <div className="divide-y divide-zinc-800/30">
                {events.map((event) => (
                  <button
                    key={event.id}
                    onClick={() => setSelectedEventId(event.id)}
                    className={`flex w-full flex-col gap-1 px-4 py-3 text-left transition-colors hover:bg-zinc-800/30 ${selectedEventId === event.id ? "bg-zinc-800/50" : ""}`}
                  >
                    <div className="flex items-center justify-between">
                      <span className="text-[8px] font-mono font-bold text-zinc-500 uppercase tracking-tighter">
                        {event.type}
                      </span>
                      <span className="text-[8px] font-mono text-zinc-600">
                        {new Date(event.started_at).toLocaleTimeString([], {
                          hour: "2-digit",
                          minute: "2-digit",
                          second: "2-digit",
                          hour12: false,
                        })}
                      </span>
                    </div>
                    <h3 className="line-clamp-2 text-[11px] font-medium leading-tight text-zinc-300">
                      {event.title}
                    </h3>
                    <div className="mt-1 flex items-center gap-1.5">
                      <div
                        className={`h-1 w-1 rounded-full ${getSeverityColor(event.severity)}`}
                      />
                      <span className="text-[8px] font-medium text-zinc-600 uppercase tracking-widest">
                        Lvl {event.severity}{" "}
                        <span className="mx-1 opacity-30">|</span>{" "}
                        {event.status}
                      </span>
                    </div>
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
          <div ref={mapContainer} className="theeye-map absolute inset-0 h-full w-full" />

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
          <div className="border-b border-zinc-800/50 px-4 py-3">
            <h2 className="text-[10px] font-bold uppercase tracking-[0.2em] text-zinc-500">
              Event Intelligence
            </h2>
          </div>

          <div className="flex-1 overflow-y-auto p-6">
            {selectedEventId ? (
              <div className="flex flex-col gap-6 animate-in fade-in slide-in-from-right-2 duration-300">
                <div>
                  <div className="mb-1 text-[8px] font-mono font-bold text-zinc-600 uppercase tracking-widest">
                    ID: {selectedEventId}
                  </div>
                  <h2 className="text-lg font-medium leading-tight text-white">
                    {events.find((e) => e.id === selectedEventId)?.title}
                  </h2>
                </div>

                <div className="grid grid-cols-2 gap-4">
                  <div className="flex flex-col gap-1">
                    <span className="text-[8px] font-bold text-zinc-600 uppercase tracking-widest">
                      Status
                    </span>
                    <span className="text-xs text-zinc-400 capitalize">
                      {events.find((e) => e.id === selectedEventId)?.status}
                    </span>
                  </div>
                  <div className="flex flex-col gap-1">
                    <span className="text-[8px] font-bold text-zinc-600 uppercase tracking-widest">
                      Severity
                    </span>
                    <span className="text-xs text-zinc-400">
                      Level{" "}
                      {events.find((e) => e.id === selectedEventId)?.severity}
                    </span>
                  </div>
                </div>

                <div className="h-px w-full bg-zinc-800/50" />

                <div className="flex flex-col gap-4">
                  <p className="text-[10px] leading-relaxed text-zinc-500 uppercase tracking-wider">
                    Further intelligence awaiting integration. Full metadata
                    wiring scheduled for Step 4.
                  </p>
                </div>
              </div>
            ) : (
              <div className="flex h-full flex-col items-center justify-center text-center">
                <div className="mb-4 h-px w-8 bg-zinc-800" />
                <p className="max-w-45 text-[10px] leading-relaxed text-zinc-600 uppercase tracking-wider">
                  Select a signal from the feed or map to inspect metadata and
                  intelligence.
                </p>
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

function getSeverityColor(severity: number): string {
  if (severity >= 4) return "bg-rose-500 shadow-[0_0_4px_rgba(244,63,94,0.5)]";
  if (severity >= 2)
    return "bg-amber-500 shadow-[0_0_4px_rgba(245,158,11,0.5)]";
  return "bg-emerald-500 shadow-[0_0_4px_rgba(16,185,129,0.5)]";
}
