import { Group, Text, Box } from "@mantine/core"

const Header = () => {
  return (
    <Box p="md" bg="dark.7" style={{ height: "100%" }}>
      <Group justify="space-between" align="center" style={{ height: "100%" }}>
        <Text size="xl" fw={700} c="white">
          GoReact App
        </Text>
      </Group>
    </Box>
  )
}

export default Header
