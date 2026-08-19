import { describe, expect, it } from "vitest";

import { clampSeverityLevel, getSeverityTone } from "./severity";

describe("clampSeverityLevel", () => {
  it("defaults to 1 when severity is missing", () => {
    expect(clampSeverityLevel()).toBe(1);
    expect(clampSeverityLevel(null)).toBe(1);
    expect(clampSeverityLevel(undefined)).toBe(1);
  });

  it("maps each documented level to itself", () => {
    expect(clampSeverityLevel(1)).toBe(1);
    expect(clampSeverityLevel(2)).toBe(2);
    expect(clampSeverityLevel(3)).toBe(3);
  });

  it("clamps out-of-range values into 1..3", () => {
    expect(clampSeverityLevel(0)).toBe(1);
    expect(clampSeverityLevel(-5)).toBe(1);
    expect(clampSeverityLevel(4)).toBe(3);
    expect(clampSeverityLevel(99)).toBe(3);
  });

  it("treats a fractional value between 2 and 3 as level 1", () => {
    // clampSeverityLevel only matches 2 exactly, so 2.5 falls through to 1.
    // Pinning current behaviour so a future rewrite has to be deliberate.
    expect(clampSeverityLevel(2.5)).toBe(1);
  });
});

describe("getSeverityTone", () => {
  it("returns a distinct tone per level", () => {
    const tones = [1, 2, 3].map((level) => getSeverityTone(level).dotClass);
    expect(new Set(tones).size).toBe(3);
  });

  it("falls back to the level 1 tone for missing severity", () => {
    expect(getSeverityTone(null)).toEqual(getSeverityTone(1));
  });

  it("gives every level a full set of classes", () => {
    for (const level of [1, 2, 3]) {
      const tone = getSeverityTone(level);
      for (const value of Object.values(tone)) {
        expect(value).not.toBe("");
      }
    }
  });
});
