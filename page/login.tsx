import { useForm } from "@tanstack/react-form"
import { useMutation } from "@tanstack/react-query"
import Layout from "~/components/Layout"
import {
  TextInput,
  PasswordInput,
  Button,
  Paper,
  Title,
  Container,
} from "@mantine/core"

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
      <Container size="xs">
        <Title order={1} mb="md" ta="center">
          Login
        </Title>
        <Paper withBorder shadow="md" p={30} mt={30} radius="md">
          <form
            onSubmit={(e) => {
              e.preventDefault()
              e.stopPropagation()
              form.handleSubmit()
            }}
          >
            <form.Field
              name="username"
              validators={{
                onChange: ({ value }) =>
                  !value ? "Username is required" : undefined,
              }}
              children={(field) => (
                <TextInput
                  label="Username"
                  placeholder="Your username"
                  id={field.name}
                  name={field.name}
                  value={field.state.value}
                  onBlur={field.handleBlur}
                  onChange={(e) => field.handleChange(e.target.value)}
                  error={
                    field.state.meta.isTouched &&
                    field.state.meta.errors.length > 0
                      ? field.state.meta.errors.join(", ")
                      : undefined
                  }
                  required
                />
              )}
            />
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
                <PasswordInput
                  label="Password"
                  placeholder="Your password"
                  id={field.name}
                  name={field.name}
                  value={field.state.value}
                  onBlur={field.handleBlur}
                  onChange={(e) => field.handleChange(e.target.value)}
                  error={
                    field.state.meta.isTouched &&
                    field.state.meta.errors.length > 0
                      ? field.state.meta.errors.join(", ")
                      : undefined
                  }
                  required
                  mt="md"
                />
              )}
            />
            <form.Subscribe
              selector={(state) => [state.canSubmit]}
              children={([canSubmit]) => (
                <Button
                  fullWidth
                  mt="xl"
                  type="submit"
                  disabled={!canSubmit || loginMutation.isPending}
                  loading={loginMutation.isPending}
                >
                  Login
                </Button>
              )}
            />
          </form>
        </Paper>
      </Container>
    </Layout>
  )
}

export default LoginPage
