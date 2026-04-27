import type { LoggedProps } from "~/types.gen.ts"
import Layout from "~/components/Layout"
import { Container, Title, Paper, Text, Button, Group } from "@mantine/core"

const LoggedPage = ({ message }: LoggedProps) => {
  const handleLogout = async () => {
    try {
      const response = await fetch("/api/logout", {
        method: "POST",
      })
      if (response.ok) {
        window.location.href = "/login"
      } else {
        alert("Logout failed")
      }
    } catch (error) {
      console.error("Logout error:", error)
      alert("An error occurred during logout")
    }
  }

  return (
    <Layout>
      <Container size="md">
        <Title order={1} mb="md">
          Logged In
        </Title>
        <Paper shadow="xs" p="md" withBorder bg="green.0" style={{ borderColor: "var(--mantine-color-green-2)" }}>
          <Text c="green.9" fw={500}>{message}</Text>
          <Text mt="xs" c="dimmed">
            If you can see this, you are authenticated!
          </Text>
          <Group mt="md">
            <Button color="red" onClick={handleLogout}>
              Logout
            </Button>
          </Group>
        </Paper>
      </Container>
    </Layout>
  )
}

export default LoggedPage
