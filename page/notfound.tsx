import type { NotFoundProps } from "~/types.gen"
import Layout from "~/components/Layout"

const NotFoundPage = ({ path }: NotFoundProps) => {
  return (
    <Layout>
      <div className="container mx-auto text-center">
        <h1 className="mb-4 text-4xl font-bold">404 - Not Found</h1>
        <p className="text-xl">
          The page <code className="rounded bg-gray-200 px-1">{path}</code>{" "}
          could not be found.
        </p>
        <div className="mt-8">
          <a href="/" className="text-blue-500 hover:underline">
            Return to Home
          </a>
        </div>
      </div>
    </Layout>
  )
}

export default NotFoundPage
