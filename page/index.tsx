import type { IndexProps } from "~/types.gen"
import Layout from "~/components/Layout"
import { Title, Text, Container } from "@mantine/core"

const IndexPage = (props: IndexProps) => {
  return (
    <Layout>
      <Container size="md">
        <Title order={1}>Welcome to My Website</Title>
        <Text mt="md">This is the homepage of my website.</Text>
        <Text mt="xs">My name is {props.name}.</Text>
      </Container>
    </Layout>
  )
}

export default IndexPage
