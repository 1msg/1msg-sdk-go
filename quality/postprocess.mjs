import { readFile, writeFile } from "node:fs/promises";
import { resolve } from "node:path";
import { fileURLToPath } from "node:url";

/**
 * Repair OpenAPI Generator 7.22's form/JSON alternative handling.
 * Generated operation builders fill formParams, but the runtime prefers JSON and
 * otherwise sends an empty body. Select the contract's form media type only when
 * a form payload exists and no explicit JSON body was supplied.
 * This transformation belongs in the upstream producer and is idempotent.
 */
const path = process.argv[2]
  ? resolve(process.argv[2], "client.go")
  : fileURLToPath(new URL("../client.go", import.meta.url));
const original = await readFile(path, "utf8");
const marker = "\t// 1MSG: form builders must send the populated form payload.";
const anchor = "\tvar body *bytes.Buffer\n";
if (!original.includes(marker)) {
  if (!original.includes(anchor))
    throw new Error(
      "Generated runtime changed; review the form-body transform",
    );
  const fix = `${anchor}\n${marker}\n\tif postBody == nil && len(formParams) > 0 && strings.HasPrefix(headerParams["Content-Type"], "application/json") {\n\t\theaderParams["Content-Type"] = "application/x-www-form-urlencoded"\n\t}\n`;
  await writeFile(path, original.replace(anchor, fix));
}
