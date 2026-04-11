import { MantineProvider, AppShell, Box } from "@mantine/core"
import Header from "./Header"
import Sidebar from "./Sidebar"

interface LayoutProps {
  children: React.ReactNode
}

const Layout = ({ children }: LayoutProps) => {
  return (
    <MantineProvider>
      <AppShell
        header={{ height: 60 }}
        navbar={{
          width: 300,
          breakpoint: "sm",
        }}
        padding="md"
      >
        <AppShell.Header>
          <Header />
        </AppShell.Header>

        <AppShell.Navbar p="md">
          <Sidebar />
        </AppShell.Navbar>

        <AppShell.Main>
          <Box>{children}</Box>
        </AppShell.Main>
      </AppShell>
    </MantineProvider>
  )
}

export default Layout
