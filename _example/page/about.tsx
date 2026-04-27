import type { AboutProps } from "~/types.gen.ts"
import Layout from "~/components/Layout"
import { Title, Text, Container } from "@mantine/core"

const AboutPage = ({ count }: AboutProps) => {
  return (
    <Layout>
      <Container size="md">
        <Title order={1} mb="md">
          About Page
        </Title>
        <Text size="lg">Count: {count}</Text>
      </Container>
    </Layout>
  )
}

export default AboutPage
