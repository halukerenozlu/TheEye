"use client";

import { Event, EventFilters } from "../../types";
import { SignalFeedItem } from "./SignalFeedItem";
import { SignalFeedFilter } from "./SignalFeedFilter";
import {
  SignalFeedEmptyState,
  SignalFeedErrorState,
  SignalFeedLoadingState,
} from "./SignalFeedStates";

type SignalFeedProps = {
  events: Event[];
  isLoading: boolean;
  error: string | null;
  filters: EventFilters;
  selectedId: string | null;
  onRefresh: () => void;
  onToggleSort: () => void;
  onTypeChange: (type: string) => void;
  onSelect: (id: string) => void;
};

export function SignalFeed({
  events,
  isLoading,
  error,
  filters,
  selectedId,
  onRefresh,
  onToggleSort,
  onTypeChange,
  onSelect,
}: SignalFeedProps) {
  return (
    <aside className="flex w-72 shrink-0 flex-col border-r border-zinc-800 bg-zinc-900/20 backdrop-blur-md">
      <div className="flex items-center justify-between border-b border-zinc-800/50 px-4 py-3">
        <h2 className="text-[10px] font-bold uppercase tracking-[0.2em] text-zinc-500">
          Signal Feed
        </h2>
        <div className="flex gap-1.5">
          <button
            onClick={onRefresh}
            title="Refresh feed"
            className="h-1.5 w-1.5 rounded-full bg-zinc-700 hover:bg-zinc-500 transition-colors"
          />
          <button
            onClick={onToggleSort}
            title={`Sort: ${filters.sort === "updated_at_desc" ? "Newest" : "Oldest"}`}
            className={`h-1.5 w-1.5 rounded-full ${filters.sort === "updated_at_asc" ? "bg-emerald-500" : "bg-zinc-700"} hover:opacity-80 transition-colors`}
          />
        </div>
      </div>

      <div className="flex-1 overflow-y-auto">
        {isLoading && events.length === 0 ? (
          <SignalFeedLoadingState />
        ) : error ? (
          <SignalFeedErrorState onRetry={onRefresh} />
        ) : events.length === 0 ? (
          <SignalFeedEmptyState />
        ) : (
          <div className="divide-y divide-zinc-800/20">
            {events.map((event) => (
              <SignalFeedItem
                key={event.id}
                event={event}
                isSelected={selectedId === event.id}
                onSelect={onSelect}
              />
            ))}
          </div>
        )}
      </div>

      <SignalFeedFilter
        selectedType={filters.type}
        onChange={onTypeChange}
      />
    </aside>
  );
}
