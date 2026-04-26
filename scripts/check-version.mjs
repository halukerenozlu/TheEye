import { readFile } from "node:fs/promises";
import path from "node:path";
import { fileURLToPath } from "node:url";

const tag = process.argv[2];

if (!tag) {
  console.error("Usage: pnpm version:check vMAJOR.MINOR.PATCH");
  process.exit(1);
}

if (!/^v\d+\.\d+\.\d+$/.test(tag)) {
  console.error(`Tag must use vMAJOR.MINOR.PATCH format; found "${tag}".`);
  process.exit(1);
}

const scriptDir = path.dirname(fileURLToPath(import.meta.url));
const repoRoot = path.resolve(scriptDir, "..");
const versionPath = path.join(repoRoot, "VERSION");
const rawVersion = (await readFile(versionPath, "utf8")).trim();
const expectedTag = `v${rawVersion}`;

if (!/^\d+\.\d+\.\d+$/.test(rawVersion)) {
  console.error(
    `VERSION must contain MAJOR.MINOR.PATCH without a leading "v"; found "${rawVersion}".`,
  );
  process.exit(1);
}

if (tag !== expectedTag) {
  console.error(`Version mismatch: VERSION=${rawVersion}, tag=${tag}.`);
  console.error(`Expected tag: ${expectedTag}`);
  process.exit(1);
}

console.log(`Version check passed: ${tag}`);
