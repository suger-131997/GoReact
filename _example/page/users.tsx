import type { UsersProps } from "~/types.gen.ts"
import Layout from "~/components/Layout"
import {
  Container,
  Title,
  Table,
  Group,
  Text,
  Pagination,
  Paper,
} from "@mantine/core"

const UsersPage = (props: UsersProps) => {
  const params = new URLSearchParams(window.location.search)
  const page = Number(params.get("page") || 0)
  const pageSize = Number(params.get("pageSize") || 10)

  const setPagination = (pageIndex: number) => {
    window.location.href = `?page=${pageIndex}&pageSize=${pageSize}`
  }

  const totalPages = Math.ceil(props.totalCount / pageSize)

  return (
    <Layout>
      <Container size="xl">
        <Title order={1} mb="md">
          Users
        </Title>
        <Paper withBorder shadow="xs">
          <Table striped highlightOnHover>
            <Table.Thead>
              <Table.Tr>
                <Table.Th>ID</Table.Th>
                <Table.Th>Name</Table.Th>
                <Table.Th>Email</Table.Th>
                <Table.Th>Role</Table.Th>
              </Table.Tr>
            </Table.Thead>
            <Table.Tbody>
              {props.users.map((user) => (
                <Table.Tr key={user.id}>
                  <Table.Td>{user.id}</Table.Td>
                  <Table.Td>{user.name}</Table.Td>
                  <Table.Td>{user.email}</Table.Td>
                  <Table.Td>{user.role}</Table.Td>
                </Table.Tr>
              ))}
            </Table.Tbody>
          </Table>
        </Paper>

        <Group justify="space-between" mt="md">
          <Text size="sm" c="dimmed">
            Total: {props.totalCount} users
          </Text>
          <Pagination
            total={totalPages || 1}
            value={page + 1}
            onChange={(p) => setPagination(p - 1)}
          />
        </Group>
      </Container>
    </Layout>
  )
}

export default UsersPage
