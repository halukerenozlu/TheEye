"use client";

import { useEffect, useState } from "react";
import { Event } from "../types";
import { fetchEventDetail } from "../lib/api";

export type UseEventDetailResult = {
  event: Event | null;
  isLoading: boolean;
};

export function useEventDetail(
  selectedId: string | null,
  options: { listForMerge: Event[]; onMissingFromList?: () => void },
): UseEventDetailResult {
  const { listForMerge, onMissingFromList } = options;
  const [event, setEvent] = useState<Event | null>(null);
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    if (!selectedId) {
      setEvent(null);
      return;
    }

    let cancelled = false;
    setIsLoading(true);

    (async () => {
      try {
        const detail = await fetchEventDetail(selectedId);
        if (!cancelled) setEvent(detail);
      } catch (err) {
        console.error("Failed to load event detail:", err);
      } finally {
        if (!cancelled) setIsLoading(false);
      }
    })();

    return () => {
      cancelled = true;
    };
  }, [selectedId]);

  useEffect(() => {
    if (!selectedId) return;

    const fromList = listForMerge.find((item) => item.id === selectedId);

    if (!fromList) {
      onMissingFromList?.();
      return;
    }

    setEvent((current) => {
      if (!current || current.id !== selectedId) return current;
      return {
        ...current,
        ...fromList,
        geometry: fromList.geometry ?? current.geometry,
      };
    });
  }, [listForMerge, selectedId, onMissingFromList]);

  return { event, isLoading };
}
