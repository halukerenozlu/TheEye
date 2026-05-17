"use client";

export function EventDetailEmpty() {
  return (
    <div className="flex h-full flex-col items-center justify-center text-center space-y-6">
      <div className="h-px w-6 bg-zinc-800" />
      <p className="max-w-50 text-[10px] leading-relaxed text-zinc-600 uppercase tracking-widest font-medium">
        Select a signal from the feed or map to inspect metadata and
        intelligence
      </p>
      <div className="h-px w-6 bg-zinc-800" />
    </div>
  );
}

export function EventDetailLoading() {
  return (
    <div className="flex h-full flex-col items-center justify-center space-y-4">
      <div className="relative h-0.5 w-16 overflow-hidden bg-zinc-900">
        <div className="h-full w-1/3 animate-[loading_1s_infinite_ease-in-out] bg-zinc-600" />
      </div>
      <span className="text-[9px] font-mono uppercase tracking-[0.3em] text-zinc-600">
        Analyzing Entity...
      </span>
    </div>
  );
}
