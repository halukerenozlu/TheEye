"use client";

import { APP_VERSION } from "../../generated/version";

type TopBarProps = {
  isLoading: boolean;
  isRefreshing: boolean;
};

export function TopBar({ isLoading, isRefreshing }: TopBarProps) {
  return (
    <header className="flex h-12 shrink-0 items-center justify-between border-b border-zinc-800 bg-zinc-950 px-4">
      <div className="flex items-center gap-4">
        <div className="text-sm font-black tracking-[0.2em] text-white">
          THE EYE
        </div>
        <div className="h-4 w-px bg-zinc-800" />
        <div className="text-[10px] font-medium text-zinc-500 uppercase tracking-widest">
          World Monitoring <span className="mx-1 opacity-30">|</span>{" "}
          <span className="normal-case">{APP_VERSION}</span>
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
  );
}
