import Layout from "~/components/Layout"
import { useEffect, useMemo, useState } from "react"
import {
  Container,
  Title,
  TextInput,
  Button,
  Group,
  Checkbox,
  ActionIcon,
  Text,
  Paper,
  Stack,
} from "@mantine/core"
import { IconTrash } from "@tabler/icons-react"

type TodoItem = {
  id: number
  text: string
  completed: boolean
}

type TodoProps = {
  initialTodos: TodoItem[]
}

const TodoPage = (props: TodoProps) => {
  const [todos, setTodos] = useState<TodoItem[]>(props.initialTodos ?? [])
  const [text, setText] = useState("")

  const stats = useMemo(() => {
    const total = todos.length
    const done = todos.filter((t) => t.completed).length
    const remaining = total - done
    return { total, done, remaining }
  }, [todos])

  const addTodo = () => {
    const trimmed = text.trim()
    if (!trimmed) return
    const nextId = (todos.at(-1)?.id ?? 0) + 1
    setTodos((prev) => [
      ...prev,
      { id: nextId, text: trimmed, completed: false },
    ])
    setText("")
  }

  const toggle = (id: number) => {
    setTodos((prev) =>
      prev.map((t) => (t.id === id ? { ...t, completed: !t.completed } : t))
    )
  }

  const remove = (id: number) => {
    setTodos((prev) => prev.filter((t) => t.id !== id))
  }

  useEffect(() => {
    const onKey = (e: KeyboardEvent) => {
      if (e.key === "Enter") {
        addTodo()
      }
    }
    window.addEventListener("keydown", onKey)
    return () => window.removeEventListener("keydown", onKey)
  }, [text, todos])

  return (
    <Layout>
      <Container size="sm">
        <Title order={1} mb="md">
          Todo List
        </Title>

        <Group mb="md">
          <TextInput
            value={text}
            onChange={(e) => setText(e.target.value)}
            placeholder="Add a task..."
            style={{ flex: 1 }}
          />
          <Button onClick={addTodo}>Add</Button>
        </Group>

        <Stack gap="sm">
          {todos.map((t) => (
            <Paper key={t.id} withBorder p="xs" shadow="xs">
              <Group justify="space-between">
                <Checkbox
                  checked={t.completed}
                  onChange={() => toggle(t.id)}
                  label={
                    <Text
                      style={{
                        textDecoration: t.completed ? "line-through" : "none",
                        color: t.completed
                          ? "var(--mantine-color-dimmed)"
                          : "inherit",
                      }}
                    >
                      {t.text}
                    </Text>
                  }
                />
                <ActionIcon
                  variant="subtle"
                  color="red"
                  onClick={() => remove(t.id)}
                  aria-label={`remove ${t.text}`}
                >
                  <IconTrash size={16} />
                </ActionIcon>
              </Group>
            </Paper>
          ))}
        </Stack>

        <Text mt="md" size="sm" c="dimmed">
          Total: {stats.total} / Done: {stats.done} / Remaining:{" "}
          {stats.remaining}
        </Text>
      </Container>
    </Layout>
  )
}

export default TodoPage
