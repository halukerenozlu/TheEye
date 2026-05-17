"use client";

import { Event } from "../../types";
import { clampSeverityLevel, getSeverityTone } from "../../lib/severity";

type EventDetailContentProps = {
  event: Event;
};

export function EventDetailContent({ event }: EventDetailContentProps) {
  const severityTone = getSeverityTone(event.severity);

  return (
    <div className="flex flex-col gap-8 animate-in fade-in slide-in-from-right-4 duration-500">
      <div className="space-y-2">
        <div className="flex items-center gap-2">
          <div className={`h-1.5 w-1.5 rounded-full ${severityTone.dotClass}`} />
          <span className="text-[8px] font-mono font-bold text-zinc-600 uppercase tracking-widest">
            Signal ID: {event.id}
          </span>
        </div>
        <h2 className="text-xl font-semibold leading-tight text-white tracking-tight">
          {event.title}
        </h2>
      </div>

      <div className="grid grid-cols-2 gap-y-6 gap-x-4">
        <div className="flex flex-col gap-1.5">
          <span className="text-[8px] font-bold text-zinc-600 uppercase tracking-[0.15em]">
            Status
          </span>
          <span className="text-[11px] text-zinc-300 capitalize flex items-center gap-2">
            <span className="h-1 w-1 rounded-full bg-zinc-500" />
            {event.status}
          </span>
        </div>
        <div className="flex flex-col gap-1.5">
          <span className="text-[8px] font-bold text-zinc-600 uppercase tracking-[0.15em]">
            Severity
          </span>
          <span className={`text-[11px] font-semibold ${severityTone.textClass}`}>
            Level {clampSeverityLevel(event.severity)}
          </span>
        </div>
        <div className="flex flex-col gap-1.5">
          <span className="text-[8px] font-bold text-zinc-600 uppercase tracking-[0.15em]">
            Category
          </span>
          <span className="text-[11px] text-zinc-300 capitalize">
            {event.type}
          </span>
        </div>
        <div className="flex flex-col gap-1.5">
          <span className="text-[8px] font-bold text-zinc-600 uppercase tracking-[0.15em]">
            Temporal Mark
          </span>
          <span className="text-[11px] text-zinc-300">
            {new Date(event.started_at).toLocaleTimeString([], {
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
              {event.geometry
                ? `${event.geometry.latitude.toFixed(6)}°N, ${event.geometry.longitude.toFixed(6)}°E`
                : "Position data inaccessible"}
            </span>
          </div>
        </div>
      </div>
    </div>
  );
}
