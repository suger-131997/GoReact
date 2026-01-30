import type { LoggedProps } from "~/types.gen"
import Layout from "~/components/Layout"

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
      <div className="container mx-auto">
        <h1 className="mb-4 text-2xl font-bold">Logged In</h1>
        <div className="rounded border border-green-200 bg-green-50 p-4">
          <p className="text-green-700">{message}</p>
          <p className="mt-2 text-gray-600">
            If you can see this, you are authenticated!
          </p>
          <button
            onClick={handleLogout}
            className="mt-4 rounded bg-red-600 px-4 py-2 font-bold text-white hover:bg-red-700 focus:outline-hidden"
          >
            Logout
          </button>
        </div>
      </div>
    </Layout>
  )
}

export default LoggedPage
