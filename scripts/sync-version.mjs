import { mkdir, readFile, writeFile } from "node:fs/promises";
import path from "node:path";
import { fileURLToPath } from "node:url";

const scriptDir = path.dirname(fileURLToPath(import.meta.url));
const repoRoot = path.resolve(scriptDir, "..");
const versionPath = path.join(repoRoot, "VERSION");
const dashboardVersionPath = path.join(
  repoRoot,
  "apps",
  "dashboard",
  "src",
  "generated",
  "version.ts",
);

const rawVersion = (await readFile(versionPath, "utf8")).trim();

if (!/^\d+\.\d+\.\d+$/.test(rawVersion)) {
  console.error(
    `VERSION must contain MAJOR.MINOR.PATCH without a leading "v"; found "${rawVersion}".`,
  );
  process.exit(1);
}

const appVersion = `v${rawVersion}`;
const generatedContent = `export const APP_VERSION = "${appVersion}";\n`;

await mkdir(path.dirname(dashboardVersionPath), { recursive: true });
await writeFile(dashboardVersionPath, generatedContent, "utf8");

console.log(`Synced dashboard version: ${appVersion}`);
