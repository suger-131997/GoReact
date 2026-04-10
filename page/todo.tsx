import Layout from "~/components/Layout"
import { useCallback, useEffect, useMemo, useState } from "react"

// ローカル型（types.gen.ts の自動生成を待たずに利用できるように）
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

  // 未完了/完了の件数
  const stats = useMemo(() => {
    const total = todos.length
    const done = todos.filter((t) => t.completed).length
    const remaining = total - done
    return { total, done, remaining }
  }, [todos])

  const addTodo = useCallback(() => {
    const trimmed = text.trim()
    if (!trimmed) return
    const nextId = (todos.at(-1)?.id ?? 0) + 1
    setTodos((prev) => [
      ...prev,
      { id: nextId, text: trimmed, completed: false },
    ])
    setText("")
  }, [text, todos])

  const toggle = (id: number) => {
    setTodos((prev) =>
      prev.map((t) => (t.id === id ? { ...t, completed: !t.completed } : t))
    )
  }

  const remove = (id: number) => {
    setTodos((prev) => prev.filter((t) => t.id !== id))
  }

  // Enterキーで追加
  useEffect(() => {
    const onKey = (e: KeyboardEvent) => {
      if (e.key === "Enter") {
        addTodo()
      }
    }
    window.addEventListener("keydown", onKey)
    return () => window.removeEventListener("keydown", onKey)
  }, [text, todos, addTodo])

  return (
    <Layout>
      <div className="mx-auto max-w-xl">
        <h1 className="mb-4 text-2xl font-bold">Todo List</h1>

        <div className="mb-4 flex gap-2">
          <input
            value={text}
            onChange={(e) => setText(e.target.value)}
            placeholder="Add a task..."
            className="flex-1 rounded border px-3 py-2"
          />
          <button
            onClick={addTodo}
            className="rounded bg-blue-600 px-4 py-2 text-white hover:bg-blue-700"
          >
            Add
          </button>
        </div>

        <ul className="space-y-2">
          {todos.map((t) => (
            <li
              key={t.id}
              className="flex items-center justify-between rounded border bg-white px-3 py-2"
            >
              <label className="flex items-center gap-2">
                <input
                  type="checkbox"
                  checked={t.completed}
                  onChange={() => toggle(t.id)}
                />
                <span
                  className={t.completed ? "text-gray-500 line-through" : ""}
                >
                  {t.text}
                </span>
              </label>
              <button
                onClick={() => remove(t.id)}
                className="rounded px-2 py-1 text-sm text-red-600 hover:bg-red-50"
                aria-label={`remove ${t.text}`}
              >
                Remove
              </button>
            </li>
          ))}
        </ul>

        <div className="mt-4 text-sm text-gray-700">
          Total: {stats.total} / Done: {stats.done} / Remaining:{" "}
          {stats.remaining}
        </div>
      </div>
    </Layout>
  )
}

export default TodoPage
