"use client";

import { Event } from "../../types";
import { clampSeverityLevel, getSeverityTone } from "../../lib/severity";

type SignalFeedItemProps = {
  event: Event;
  isSelected: boolean;
  onSelect: (id: string) => void;
};

export function SignalFeedItem({
  event,
  isSelected,
  onSelect,
}: SignalFeedItemProps) {
  const severityTone = getSeverityTone(event.severity);

  return (
    <button
      onClick={() => onSelect(event.id)}
      className={`flex w-full flex-col gap-1.5 px-4 py-4 text-left transition-all duration-200 border-l-2 ${isSelected ? "bg-zinc-800/30 border-emerald-500/50" : "hover:bg-zinc-800/10 border-transparent"}`}
    >
      <div className="flex items-center justify-between">
        <span
          className={`text-[8px] font-mono font-bold uppercase tracking-tighter ${isSelected ? "text-emerald-500" : "text-zinc-500"}`}
        >
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
      <h3
        className={`line-clamp-2 text-[11px] font-medium leading-snug transition-colors ${isSelected ? "text-white" : "text-zinc-400 group-hover:text-zinc-300"}`}
      >
        {event.title}
      </h3>
      <div className="mt-0.5 flex items-center gap-2">
        <div className={`h-1.5 w-1.5 rounded-full ${severityTone.dotClass}`} />
        <span
          className={`rounded border px-1.5 py-px text-[8px] font-bold uppercase tracking-widest ${severityTone.badgeClass}`}
        >
          Lvl {clampSeverityLevel(event.severity)}
        </span>
        <span className="text-[8px] font-bold text-zinc-600 uppercase tracking-widest">
          {event.status}
        </span>
      </div>
    </button>
  );
}
