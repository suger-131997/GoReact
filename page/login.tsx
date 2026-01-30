import { useForm } from "@tanstack/react-form"
import { useMutation } from "@tanstack/react-query"
import Layout from "~/components/Layout"

const LoginPage = () => {
  const loginMutation = useMutation({
    mutationFn: async (value: { username: string; password: string }) => {
      const response = await fetch("/api/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          username: value.username,
          password: value.password,
        }),
      })

      if (!response.ok) {
        throw new Error("Login failed")
      }

      return response.json()
    },
    onSuccess: (data) => {
      console.log("Login success:", data)
      window.location.href = "/logged"
    },
    onError: (error) => {
      console.error("Error during login:", error)
      alert("Login failed. Please check your username and password.")
    },
  })

  const form = useForm({
    defaultValues: {
      username: "",
      password: "",
    },
    onSubmit: async ({ value }) => {
      loginMutation.mutate(value)
    },
  })

  return (
    <Layout>
      <div className="container mx-auto flex flex-col items-center">
        <h1 className="mb-4 text-2xl font-bold">Login</h1>
        <form
          onSubmit={(e) => {
            e.preventDefault()
            e.stopPropagation()
            form.handleSubmit()
          }}
          className="mb-4 w-full max-w-sm rounded border border-gray-200 bg-white px-8 pt-6 pb-8 shadow-md"
        >
          <div className="mb-4">
            <form.Field
              name="username"
              validators={{
                onChange: ({ value }) =>
                  !value ? "Username is required" : undefined,
              }}
              children={(field) => (
                <>
                  <label
                    htmlFor={field.name}
                    className="mb-2 block text-sm font-bold text-gray-700"
                  >
                    Username:
                  </label>
                  <input
                    id={field.name}
                    name={field.name}
                    value={field.state.value}
                    onBlur={field.handleBlur}
                    onChange={(e) => field.handleChange(e.target.value)}
                    className="focus:shadow-outline w-full appearance-none rounded border px-3 py-2 leading-tight text-gray-700 shadow focus:outline-none"
                  />
                  {field.state.meta.isTouched &&
                  field.state.meta.errors.length ? (
                    <em className="text-xs text-red-500 italic">
                      {field.state.meta.errors.join(", ")}
                    </em>
                  ) : null}
                </>
              )}
            ></form.Field>
          </div>
          <div className="mb-6">
            <form.Field
              name="password"
              validators={{
                onChange: ({ value }) =>
                  !value
                    ? "Password is required"
                    : value.length < 6
                      ? "Password must be at least 6 characters"
                      : undefined,
              }}
              children={(field) => (
                <>
                  <label
                    htmlFor={field.name}
                    className="mb-2 block text-sm font-bold text-gray-700"
                  >
                    Password:
                  </label>
                  <input
                    id={field.name}
                    name={field.name}
                    type="password"
                    value={field.state.value}
                    onBlur={field.handleBlur}
                    onChange={(e) => field.handleChange(e.target.value)}
                    className="focus:shadow-outline mb-3 w-full appearance-none rounded border px-3 py-2 leading-tight text-gray-700 shadow focus:outline-none"
                  />
                  {field.state.meta.isTouched &&
                  field.state.meta.errors.length ? (
                    <em className="text-xs text-red-500 italic">
                      {field.state.meta.errors.join(", ")}
                    </em>
                  ) : null}
                </>
              )}
            />
          </div>
          <div className="flex items-center justify-between">
            <form.Subscribe
              selector={(state) => [state.canSubmit]}
              children={([canSubmit]) => (
                <button
                  type="submit"
                  disabled={!canSubmit || loginMutation.isPending}
                  className="focus:shadow-outline rounded bg-blue-500 px-4 py-2 font-bold text-white hover:bg-blue-700 focus:outline-none disabled:opacity-50"
                >
                  {loginMutation.isPending ? "..." : "Login"}
                </button>
              )}
            />
          </div>
        </form>
      </div>
    </Layout>
  )
}

export default LoginPage
