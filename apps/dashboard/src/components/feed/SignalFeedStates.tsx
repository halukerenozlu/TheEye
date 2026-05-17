"use client";

export function SignalFeedLoadingState() {
  return (
    <div className="flex h-full flex-col items-center justify-center p-8 text-center space-y-4">
      <div className="relative h-px w-24 overflow-hidden bg-zinc-900">
        <div className="h-full w-1/3 animate-[loading_1.5s_infinite_ease-in-out] bg-zinc-700" />
      </div>
      <span className="text-[9px] font-mono uppercase tracking-[0.3em] text-zinc-600">
        Acquiring Stream...
      </span>
    </div>
  );
}

type SignalFeedErrorStateProps = {
  onRetry: () => void;
};

export function SignalFeedErrorState({ onRetry }: SignalFeedErrorStateProps) {
  return (
    <div className="flex h-full flex-col items-center justify-center p-8 text-center">
      <div className="mb-3 flex h-10 w-10 items-center justify-center rounded-full border border-rose-900/30 bg-rose-950/10 text-rose-900">
        <svg
          width="18"
          height="18"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          strokeWidth="1.5"
        >
          <circle cx="12" cy="12" r="10" />
          <line x1="12" y1="8" x2="12" y2="12" />
          <line x1="12" y1="16" x2="12.01" y2="16" />
        </svg>
      </div>
      <p className="text-[10px] font-medium uppercase tracking-widest text-rose-900/80">
        Signal Sync Failed
      </p>
      <button
        onClick={onRetry}
        className="mt-4 text-[9px] font-bold uppercase tracking-tighter text-zinc-500 hover:text-zinc-300 transition-colors underline underline-offset-4"
      >
        Retry Connection
      </button>
    </div>
  );
}

export function SignalFeedEmptyState() {
  return (
    <div className="flex h-full flex-col items-center justify-center p-8 text-center text-zinc-700">
      <div className="mb-4 h-12 w-12 opacity-20 grayscale filter">
        <svg
          width="48"
          height="48"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          strokeWidth="0.5"
        >
          <path d="M2 12h3l2 8 4-16 4 16 2-8h3" />
        </svg>
      </div>
      <p className="text-[9px] uppercase tracking-[0.2em] font-medium text-zinc-600">
        No active signals detected
      </p>
    </div>
  );
}
