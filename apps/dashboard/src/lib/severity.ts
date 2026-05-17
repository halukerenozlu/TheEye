export type SeverityLevel = 1 | 2 | 3;

export type SeverityTone = {
  dotClass: string;
  markerClass: string;
  markerSizeClass: string;
  badgeClass: string;
  textClass: string;
};

export function clampSeverityLevel(severity: number): SeverityLevel {
  if (severity >= 3) return 3;
  if (severity === 2) return 2;
  return 1;
}

export function getSeverityTone(severity: number): SeverityTone {
  const level = clampSeverityLevel(severity);

  if (level === 3) {
    return {
      dotClass: "bg-rose-500 shadow-[0_0_6px_rgba(244,63,94,0.6)]",
      markerClass:
        "bg-rose-500 border-rose-300/60 shadow-[0_0_10px_rgba(244,63,94,0.75)]",
      markerSizeClass: "h-3.5 w-3.5",
      badgeClass: "border-rose-400/50 bg-rose-500/15 text-rose-300",
      textClass: "text-rose-300",
    };
  }

  if (level === 2) {
    return {
      dotClass: "bg-amber-400 shadow-[0_0_5px_rgba(251,191,36,0.55)]",
      markerClass:
        "bg-amber-400 border-amber-200/60 shadow-[0_0_8px_rgba(251,191,36,0.6)]",
      markerSizeClass: "h-3 w-3",
      badgeClass: "border-amber-300/40 bg-amber-400/10 text-amber-200",
      textClass: "text-amber-200",
    };
  }

  return {
    dotClass: "bg-emerald-500 shadow-[0_0_4px_rgba(16,185,129,0.45)]",
    markerClass:
      "bg-emerald-500 border-emerald-300/50 shadow-[0_0_6px_rgba(16,185,129,0.5)]",
    markerSizeClass: "h-2.5 w-2.5",
    badgeClass: "border-emerald-400/35 bg-emerald-500/10 text-emerald-200",
    textClass: "text-emerald-200",
  };
}
