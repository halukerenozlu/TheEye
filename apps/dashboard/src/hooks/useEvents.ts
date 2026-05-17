"use client";

import { useCallback, useEffect, useRef, useState } from "react";
import { Event, EventFilters } from "../types";
import { fetchEvents } from "../lib/api";

const EVENTS_POLL_INTERVAL_MS = 30000;

export type UseEventsResult = {
  items: Event[];
  isLoading: boolean;
  isRefreshing: boolean;
  error: string | null;
  refetch: () => void;
};

export function useEvents(filters: EventFilters): UseEventsResult {
  const [items, setItems] = useState<Event[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [isRefreshing, setIsRefreshing] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const itemsCountRef = useRef(0);

  useEffect(() => {
    itemsCountRef.current = items.length;
  }, [items.length]);

  const load = useCallback(
    async (options?: { background?: boolean }) => {
      await Promise.resolve();

      const isBackgroundRefresh = options?.background ?? false;

      if (isBackgroundRefresh) {
        setIsRefreshing(true);
      } else {
        setIsLoading(true);
        setError(null);
      }

      try {
        const data = await fetchEvents(filters);
        setItems(data.items);
      } catch (err) {
        console.error(err);

        if (!isBackgroundRefresh || itemsCountRef.current === 0) {
          setError("Failed to sync with signal intelligence");
        }
      } finally {
        if (isBackgroundRefresh) {
          setIsRefreshing(false);
        } else {
          setIsLoading(false);
        }
      }
    },
    [filters],
  );

  useEffect(() => {
    let cancelled = false;

    (async () => {
      await Promise.resolve();
      if (cancelled) return;

      setIsLoading(true);
      setError(null);

      try {
        const data = await fetchEvents(filters);
        if (!cancelled) {
          setItems(data.items);
        }
      } catch (err) {
        console.error(err);

        if (!cancelled && itemsCountRef.current === 0) {
          setError("Failed to sync with signal intelligence");
        }
      } finally {
        if (!cancelled) {
          setIsLoading(false);
        }
      }
    })();

    return () => {
      cancelled = true;
    };
  }, [filters]);

  useEffect(() => {
    const poll = () => {
      if (document.hidden) return;
      void load({ background: true });
    };

    const intervalId = window.setInterval(poll, EVENTS_POLL_INTERVAL_MS);
    return () => window.clearInterval(intervalId);
  }, [load]);

  const refetch = useCallback(() => {
    void load();
  }, [load]);

  return { items, isLoading, isRefreshing, error, refetch };
}
