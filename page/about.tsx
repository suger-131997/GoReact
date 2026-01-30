import type { AboutProps } from "~/types.gen"
import Layout from "~/components/Layout"

const AboutPage = ({ count }: AboutProps) => {
  return (
    <Layout>
      <div className="container mx-auto">
        <h1 className="mb-4 text-2xl font-bold">About Page</h1>
        <p>Count: {count}</p>
      </div>
    </Layout>
  )
}

export default AboutPage
