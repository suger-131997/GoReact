import type { UsersPageProps } from "~/types.gen"
import Layout from "~/components/Layout"

const UsersPage = (props: UsersPageProps) => {
  const params = new URLSearchParams(window.location.search)
  console.log(props)
  const page = Number(params.get("page") || 0)
  const pageSize = Number(params.get("pageSize") || 10)

  console.log(page)

  const setPagination = (pagination: {
    pageIndex: number
    pageSize: number
  }) => {
    console.log(pagination)
    window.location.href = `?page=${pagination.pageIndex}&pageSize=${pagination.pageSize}`
  }

  const totalPages = Math.ceil(props.totalCount / pageSize)

  return (
    <Layout>
      <div className="container mx-auto">
        <h1 className="mb-4 text-2xl font-bold">Users</h1>
        <div className="overflow-x-auto">
          <table className="min-w-full border border-gray-200 bg-white">
            <thead>
              <tr className="bg-gray-100">
                <th className="border-b px-4 py-2 text-left">ID</th>
                <th className="border-b px-4 py-2 text-left">Name</th>
                <th className="border-b px-4 py-2 text-left">Email</th>
                <th className="border-b px-4 py-2 text-left">Role</th>
              </tr>
            </thead>
            <tbody>
              {props.users.map((user) => (
                <tr key={user.id} className="hover:bg-gray-50">
                  <td className="border-b px-4 py-2">{user.id}</td>
                  <td className="border-b px-4 py-2">{user.name}</td>
                  <td className="border-b px-4 py-2">{user.email}</td>
                  <td className="border-b px-4 py-2">{user.role}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
        <div className="mt-4 flex items-center justify-between">
          <div className="text-sm text-gray-700">
            Total: {props.totalCount} users
          </div>
          <div className="flex space-x-2">
            <button
              onClick={() => setPagination({ pageIndex: page - 1, pageSize })}
              disabled={page <= 0}
              className="rounded border px-3 py-1 disabled:opacity-50"
            >
              Previous
            </button>
            <span className="px-3 py-1">
              Page {page + 1} of {totalPages || 1}
            </span>
            <button
              onClick={() => setPagination({ pageIndex: page + 1, pageSize })}
              disabled={page >= totalPages - 1}
              className="rounded border px-3 py-1 disabled:opacity-50"
            >
              Next
            </button>
          </div>
        </div>
      </div>
    </Layout>
  )
}

export default UsersPage
