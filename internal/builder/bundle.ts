import { rolldown } from "rolldown"
import { readFileSync } from "node:fs"
import { join } from "node:path"

const isDevelopment = process.env.NODE_ENV === "development"

let bundle,
  failed = false
try {
  const input = readFileSync(0, "utf-8")

  bundle = await rolldown({
    input: "stdin.tsx",
    plugins: [
      {
        name: "stdin",
        resolveId(id) {
          if (id === "stdin.tsx") return id
        },
        load(id) {
          if (id === "stdin.tsx") return input
        },
      },
    ],
    tsconfig: join(process.cwd(), "tsconfig.json"),
  })
  const output = await bundle.generate({
    format: "esm",
    minify: !isDevelopment,
    sourcemap: isDevelopment ? "inline" : false,
  })
  console.log(output.output[0].code)
} catch (e) {
  console.error(e)
  failed = true
}
if (bundle) {
  await bundle.close()
}

process.exitCode = failed ? 1 : 0
