"use client";

import { useEffect, useMemo, useState } from "react";
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
  const [detailEvent, setDetailEvent] = useState<Event | null>(null);
  const [fetchingId, setFetchingId] = useState<string | null>(null);
  const [failedId, setFailedId] = useState<string | null>(null);

  useEffect(() => {
    if (!selectedId) {
      return;
    }

    let cancelled = false;

    (async () => {
      await Promise.resolve();
      if (cancelled) return;

      setFetchingId(selectedId);
      setFailedId(null);

      try {
        const detail = await fetchEventDetail(selectedId);
        if (!cancelled) {
          setDetailEvent(detail);
        }
      } catch (err) {
        console.error("Failed to load event detail:", err);
        if (!cancelled) {
          setFailedId(selectedId);
        }
      } finally {
        if (!cancelled) {
          setFetchingId((current) => (current === selectedId ? null : current));
        }
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
  }, [listForMerge, selectedId, onMissingFromList]);

  const event = useMemo(() => {
    if (!selectedId || !detailEvent || detailEvent.id !== selectedId) {
      return null;
    }

    const fromList = listForMerge.find((item) => item.id === selectedId);
    if (!fromList) return detailEvent;

    return {
      ...detailEvent,
      ...fromList,
      geometry: fromList.geometry ?? detailEvent.geometry,
    };
  }, [detailEvent, listForMerge, selectedId]);

  const isLoading =
    selectedId !== null &&
    failedId !== selectedId &&
    (fetchingId === selectedId || detailEvent?.id !== selectedId);

  return { event, isLoading };
}
