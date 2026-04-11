import type { NotFoundProps } from "~/types.gen"
import Layout from "~/components/Layout"
import { Container, Title, Text, Anchor, Code, Box } from "@mantine/core"

const NotFoundPage = ({ path }: NotFoundProps) => {
  return (
    <Layout>
      <Container size="md" ta="center">
        <Title order={1} mb="md">
          404 - Not Found
        </Title>
        <Text size="xl">
          The page <Code>{path}</Code> could not be found.
        </Text>
        <Box mt="xl">
          <Anchor href="/" size="lg">
            Return to Home
          </Anchor>
        </Box>
      </Container>
    </Layout>
  )
}

export default NotFoundPage
