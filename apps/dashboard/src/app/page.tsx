"use client";

import { useCallback, useState } from "react";
import { EventFilters } from "../types";
import { useEvents } from "../hooks/useEvents";
import { useEventDetail } from "../hooks/useEventDetail";
import { TopBar } from "../components/layout/TopBar";
import { SignalFeed } from "../components/feed/SignalFeed";
import { MapView } from "../components/map/MapView";
import { EventDetailPanel } from "../components/detail/EventDetailPanel";

export default function DashboardPage() {
  const [filters, setFilters] = useState<EventFilters>({
    sort: "updated_at_desc",
  });
  const [selectedId, setSelectedId] = useState<string | null>(null);

  const events = useEvents(filters);

  const clearSelection = useCallback(() => setSelectedId(null), []);

  const detail = useEventDetail(selectedId, {
    listForMerge: events.items,
    onMissingFromList: clearSelection,
  });

  const toggleSort = useCallback(() => {
    setFilters((prev) => ({
      ...prev,
      sort:
        prev.sort === "updated_at_desc" ? "updated_at_asc" : "updated_at_desc",
    }));
  }, []);

  const updateTypeFilter = useCallback((type: string) => {
    setFilters((prev) => ({
      ...prev,
      type: type || undefined,
    }));
  }, []);

  return (
    <div className="flex h-screen w-full flex-col overflow-hidden bg-black font-sans text-zinc-300 antialiased selection:bg-zinc-800 selection:text-white">
      <TopBar isLoading={events.isLoading} isRefreshing={events.isRefreshing} />

      <div className="flex flex-1 overflow-hidden">
        <SignalFeed
          events={events.items}
          isLoading={events.isLoading}
          error={events.error}
          filters={filters}
          selectedId={selectedId}
          onRefresh={events.refetch}
          onToggleSort={toggleSort}
          onTypeChange={updateTypeFilter}
          onSelect={setSelectedId}
        />

        <MapView
          events={events.items}
          selectedId={selectedId}
          focusOn={detail.event?.geometry ?? null}
          onMarkerClick={setSelectedId}
        />

        <EventDetailPanel
          event={detail.event}
          isLoading={detail.isLoading}
          hasSelection={selectedId !== null}
          onClear={clearSelection}
        />
      </div>
    </div>
  );
}
