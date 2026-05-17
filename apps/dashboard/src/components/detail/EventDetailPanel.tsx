"use client";

import { Event } from "../../types";
import { EventDetailContent } from "./EventDetailContent";
import { EventDetailEmpty, EventDetailLoading } from "./EventDetailEmpty";

type EventDetailPanelProps = {
  event: Event | null;
  isLoading: boolean;
  hasSelection: boolean;
  onClear: () => void;
};

export function EventDetailPanel({
  event,
  isLoading,
  hasSelection,
  onClear,
}: EventDetailPanelProps) {
  return (
    <aside className="flex w-96 shrink-0 flex-col border-l border-zinc-800 bg-zinc-900/20 backdrop-blur-md">
      <div className="flex items-center justify-between border-b border-zinc-800/50 px-4 py-3">
        <h2 className="text-[10px] font-bold uppercase tracking-[0.2em] text-zinc-500">
          Event Intelligence
        </h2>
        {hasSelection && (
          <button
            onClick={onClear}
            className="text-[9px] font-bold uppercase tracking-widest text-zinc-600 hover:text-zinc-400 transition-colors"
          >
            Clear
          </button>
        )}
      </div>

      <div className="flex-1 overflow-y-auto p-6">
        {isLoading ? (
          <EventDetailLoading />
        ) : event ? (
          <EventDetailContent event={event} />
        ) : (
          <EventDetailEmpty />
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
  );
}
