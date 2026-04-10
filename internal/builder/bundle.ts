import { rolldown } from "rolldown"
import { readFileSync, writeFileSync, rmSync } from "node:fs"
import { join } from "node:path"

const isDevelopment = process.env.NODE_ENV === "development"

let bundle,
  failed = false
try {
  const input = readFileSync(0, "utf-8")
  const tmpFile = join(join(process.cwd(), ".tmp"), `bundle-${Date.now()}.tsx`)
  writeFileSync(tmpFile, input)

  bundle = await rolldown({
    input: tmpFile,
    tsconfig: join(process.cwd(), "tsconfig.json"),
  })
  const output = await bundle.generate({
    format: "esm",
    minify: !isDevelopment,
    sourcemap: isDevelopment ? "inline" : false,
  })
  console.log(output.output[0].code)
  rmSync(tmpFile)
} catch (e) {
  console.error(e)
  failed = true
}
if (bundle) {
  await bundle.close()
}

process.exitCode = failed ? 1 : 0
