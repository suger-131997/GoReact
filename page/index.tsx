import type { IndexProps } from "~/types.gen"
import Layout from "~/components/Layout"

const IndexPage = (props: IndexProps) => {
  return (
    <Layout>
      <div className="container mx-auto">
        <h1 className="text-2xl font-bold">Welcome to My Website</h1>
        <p>This is the homepage of my website.</p>
        <p>My name is {props.name}.</p>
      </div>
    </Layout>
  )
}

export default IndexPage
