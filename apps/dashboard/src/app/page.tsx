"use client";

/**
 * TheEye Dashboard Shell
 * Phase 5 / Sprint 1 / Step 1
 *
 * This shell establishes the "Command Center Lite" layout:
 * - Top Bar: Platform meta and status
 * - Left Panel: Event feed and filters scaffolding
 * - Center Panel: Map area (placeholder for MapLibre)
 * - Right Panel: Event detail and intelligence
 */

export default function DashboardPage() {
  return (
    <div className="flex h-screen w-full flex-col overflow-hidden bg-black font-sans text-zinc-300 antialiased selection:bg-zinc-800 selection:text-white">
      {/* 
        TOP BAR
        Restrained, high-contrast text for branding, low-contrast for secondary info.
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
            <div className="h-1.5 w-1.5 animate-pulse rounded-full bg-emerald-500 shadow-[0_0_8px_rgba(16,185,129,0.5)]" />
            <span className="text-[9px] font-semibold text-zinc-400 uppercase tracking-tighter">
              System Ready
            </span>
          </div>
        </div>
      </header>

      {/* 
        MAIN CONTENT AREA
        Three-column layout for high-density information management.
      */}
      <div className="flex flex-1 overflow-hidden">
        {/* 
          LEFT PANEL: SIGNAL FEED
          Primary list of incoming events.
        */}
        <aside className="flex w-72 shrink-0 flex-col border-r border-zinc-800 bg-zinc-900/20 backdrop-blur-md">
          <div className="flex items-center justify-between border-b border-zinc-800/50 px-4 py-3">
            <h2 className="text-[10px] font-bold uppercase tracking-[0.2em] text-zinc-500">
              Signal Feed
            </h2>
            <div className="flex gap-1.5">
              <div className="h-1.5 w-1.5 rounded-full bg-zinc-700" />
              <div className="h-1.5 w-1.5 rounded-full bg-zinc-700" />
            </div>
          </div>

          <div className="flex-1 overflow-y-auto">
            {/* 
              FEED PLACEHOLDER
              Establishing the visual pattern for list items.
            */}
            <div className="flex h-full flex-col items-center justify-center p-8 text-center">
              <div className="mb-4 flex h-12 w-12 items-center justify-center rounded-xl border border-dashed border-zinc-800 text-zinc-800">
                <svg
                  width="24"
                  height="24"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  strokeWidth="1"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                >
                  <path d="M2 12h3l2 8 4-16 4 16 2-8h3" />
                </svg>
              </div>
              <p className="text-[10px] uppercase tracking-[0.2em] text-zinc-600">
                Awaiting incoming signals
              </p>
            </div>
          </div>

          <div className="border-t border-zinc-800/50 p-3 bg-zinc-950/40">
            <div className="h-7 w-full rounded border border-zinc-800 bg-zinc-900/50 flex items-center px-2">
              <span className="text-[9px] text-zinc-600 uppercase tracking-widest">
                Filter signals...
              </span>
            </div>
          </div>
        </aside>

        {/* 
          CENTER PANEL: MAP VIEWPORT
          The primary exploration surface.
        */}
        <main className="relative flex-1 bg-zinc-950 overflow-hidden">
          {/* Subtle Grid Background */}
          <div
            className="absolute inset-0 opacity-[0.03]"
            style={{
              backgroundImage:
                "radial-gradient(circle, white 1px, transparent 1px)",
              backgroundSize: "32px 32px",
            }}
          />

          {/* Map Initialization State */}
          <div className="absolute inset-0 flex flex-col items-center justify-center">
            <div className="relative mb-6 h-32 w-32">
              <div className="absolute inset-0 rounded-full border border-zinc-900" />
              <div className="absolute inset-2 rounded-full border border-zinc-800/50" />
              <div className="absolute inset-0 flex items-center justify-center">
                <div className="h-[px] w-full bg-zinc-900 rotate-45" />
                <div className="h-[px] w-full bg-zinc-900 -rotate-45" />
              </div>
            </div>
            <div className="flex flex-col items-center gap-2">
              <span className="text-[10px] font-mono uppercase tracking-[0.4em] text-zinc-600">
                Initializing Geospatial Engine
              </span>
              <div className="h-0.5 w-32 overflow-hidden rounded-full bg-zinc-900">
                <div className="h-full w-1/3 animate-[loading_2s_infinite_ease-in-out] bg-zinc-700" />
              </div>
            </div>
          </div>

          {/* Map Controls (Scaffolding) */}
          <div className="absolute bottom-6 right-6 flex flex-col gap-1.5">
            <button className="flex h-8 w-8 items-center justify-center rounded border border-zinc-800 bg-zinc-900/80 text-zinc-500 hover:border-zinc-700 hover:text-zinc-300 transition-colors">
              <span className="text-xs">+</span>
            </button>
            <button className="flex h-8 w-8 items-center justify-center rounded border border-zinc-800 bg-zinc-900/80 text-zinc-500 hover:border-zinc-700 hover:text-zinc-300 transition-colors">
              <span className="text-xs">-</span>
            </button>
          </div>

          {/* Map Attribution/Status */}
          <div className="absolute bottom-2 left-4">
            <span className="text-[8px] font-mono text-zinc-700 uppercase tracking-widest">
              Lat: 0.0000 Lon: 0.0000
            </span>
          </div>
        </main>

        {/* 
          RIGHT PANEL: EVENT INTELLIGENCE
          Detailed information about the selected entity.
        */}
        <aside className="flex w-96 shrink-0 flex-col border-l border-zinc-800 bg-zinc-900/20 backdrop-blur-md">
          <div className="border-b border-zinc-800/50 px-4 py-3">
            <h2 className="text-[10px] font-bold uppercase tracking-[0.2em] text-zinc-500">
              Event Intelligence
            </h2>
          </div>

          <div className="flex-1 overflow-y-auto p-6">
            {/* 
              DETAIL PLACEHOLDER
            */}
            <div className="flex h-full flex-col items-center justify-center text-center">
              <div className="mb-4 h-px w-8 bg-zinc-800" />
              <p className="max-w-45 text-[10px] leading-relaxed text-zinc-600 uppercase tracking-wider">
                Select a signal from the feed or map to inspect metadata and
                intelligence.
              </p>
            </div>
          </div>

          {/* Secondary Panel Actions Scaffolding */}
          <div className="border-t border-zinc-800/50 p-4 bg-zinc-950/20">
            <div className="flex gap-2">
              <div className="h-1.5 w-1.5 rounded-full bg-zinc-800" />
              <div className="h-1.5 w-1.5 rounded-full bg-zinc-800" />
              <div className="h-1.5 w-1.5 rounded-full bg-zinc-800" />
            </div>
          </div>
        </aside>
      </div>

      <style jsx global>{`
        @keyframes loading {
          0% {
            transform: translateX(-100%);
          }
          100% {
            transform: translateX(300%);
          }
        }
      `}</style>
    </div>
  );
}
