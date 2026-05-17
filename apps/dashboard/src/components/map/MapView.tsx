"use client";

import { useEffect, useRef, useState } from "react";
import maplibregl from "maplibre-gl";
import { Event, EventGeometry } from "../../types";
import { getSeverityTone } from "../../lib/severity";

type MapViewProps = {
  events: Event[];
  selectedId: string | null;
  focusOn?: EventGeometry | null;
  onMarkerClick: (id: string) => void;
};

export function MapView({
  events,
  selectedId,
  focusOn,
  onMarkerClick,
}: MapViewProps) {
  const containerRef = useRef<HTMLDivElement>(null);
  const mapRef = useRef<maplibregl.Map | null>(null);
  const markersRef = useRef<Map<string, maplibregl.Marker>>(new Map());
  const didInitialFitRef = useRef(false);
  const onMarkerClickRef = useRef(onMarkerClick);
  const [isMapReady, setIsMapReady] = useState(false);

  useEffect(() => {
    onMarkerClickRef.current = onMarkerClick;
  }, [onMarkerClick]);

  useEffect(() => {
    if (mapRef.current || !containerRef.current) return;

    const mapInstance = new maplibregl.Map({
      container: containerRef.current,
      style: "https://basemaps.cartocdn.com/gl/dark-matter-gl-style/style.json",
      center: [0, 20],
      zoom: 1.5,
      attributionControl: false,
    });

    mapRef.current = mapInstance;

    mapInstance.addControl(
      new maplibregl.NavigationControl({ showCompass: false }),
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
      mapRef.current = null;
      setIsMapReady(false);
      didInitialFitRef.current = false;
    };
  }, []);

  useEffect(() => {
    const currentMap = mapRef.current;
    if (!currentMap || !isMapReady) return;

    markersRef.current.forEach((m) => m.remove());
    markersRef.current.clear();

    const bounds = new maplibregl.LngLatBounds();
    let hasGeometry = false;

    events.forEach((event) => {
      if (!event.geometry) return;

      hasGeometry = true;
      const coords: [number, number] = [
        event.geometry.longitude,
        event.geometry.latitude,
      ];
      bounds.extend(coords);

      const isSelected = selectedId === event.id;
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
        onMarkerClickRef.current(event.id);
      });

      markersRef.current.set(event.id, marker);
    });

    if (
      hasGeometry &&
      !selectedId &&
      events.length > 0 &&
      !didInitialFitRef.current
    ) {
      didInitialFitRef.current = true;
      currentMap.fitBounds(bounds, {
        padding: 64,
        maxZoom: 10,
        duration: 2000,
      });
    }
  }, [events, isMapReady, selectedId]);

  useEffect(() => {
    const currentMap = mapRef.current;
    if (!currentMap || !isMapReady || !focusOn) return;

    currentMap.easeTo({
      center: [focusOn.longitude, focusOn.latitude],
      zoom: 6,
      duration: 1500,
    });
  }, [focusOn, isMapReady]);

  return (
    <main className="relative min-h-0 flex-1 bg-zinc-950 overflow-hidden">
      <div
        ref={containerRef}
        className="theeye-map absolute inset-0 h-full w-full"
      />
      <div className="absolute bottom-2 left-4 pointer-events-none">
        <span className="text-[8px] font-mono text-zinc-700 uppercase tracking-widest">
          Engine: MapLibre GL <span className="mx-1 opacity-30">|</span> Style:
          Dark Matter
        </span>
      </div>
    </main>
  );
}
