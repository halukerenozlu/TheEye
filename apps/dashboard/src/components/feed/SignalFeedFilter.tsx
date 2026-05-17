"use client";

import { useState } from "react";

type FilterOption = {
  label: string;
  value: string;
};

const FILTER_OPTIONS: FilterOption[] = [
  { label: "All Categories", value: "" },
  { label: "Earthquakes", value: "earthquake" },
  { label: "Wildfires", value: "wildfire" },
  { label: "Storms", value: "storm" },
];

type SignalFeedFilterProps = {
  selectedType: string | undefined;
  onChange: (type: string) => void;
};

export function SignalFeedFilter({
  selectedType,
  onChange,
}: SignalFeedFilterProps) {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <div className="border-t border-zinc-800/50 p-3 bg-zinc-950/40 relative">
      <button
        onClick={() => setIsOpen((prev) => !prev)}
        className="flex h-7 w-full items-center justify-between rounded border border-zinc-800 bg-zinc-900/50 px-2 text-[9px] text-zinc-400 uppercase tracking-widest outline-none hover:border-zinc-700 transition-colors"
      >
        <span className="truncate">
          {selectedType ? selectedType + "s" : "All Categories"}
        </span>
        <svg
          width="10"
          height="10"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          strokeWidth="2"
          className={`shrink-0 transition-transform duration-200 ${isOpen ? "rotate-180" : ""}`}
        >
          <path d="m6 9 6 6 6-6" />
        </svg>
      </button>

      {isOpen && (
        <>
          <div
            className="fixed inset-0 z-10"
            onClick={() => setIsOpen(false)}
          />
          <div className="absolute bottom-11 left-3 right-3 z-20 overflow-hidden rounded border border-zinc-800 bg-zinc-950 shadow-2xl animate-in fade-in slide-in-from-bottom-2 duration-200">
            {FILTER_OPTIONS.map((option) => {
              const isActive =
                selectedType === option.value ||
                (!selectedType && !option.value);
              return (
                <button
                  key={option.value}
                  onClick={() => {
                    onChange(option.value);
                    setIsOpen(false);
                  }}
                  className={`flex h-8 w-full items-center px-3 text-[9px] uppercase tracking-widest transition-colors hover:bg-zinc-900 ${isActive ? "text-white font-bold bg-zinc-900" : "text-zinc-500"}`}
                >
                  {option.label}
                </button>
              );
            })}
          </div>
        </>
      )}
    </div>
  );
}
