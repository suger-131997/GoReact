import { NavLink, Stack } from "@mantine/core"
import {
  IconHome,
  IconInfoCircle,
  IconLogin,
  IconUserCheck,
  IconCalendar,
  IconTable,
  IconListCheck,
} from "@tabler/icons-react"

const Sidebar = () => {
  const links = [
    { name: "Home", path: "/", icon: IconHome },
    { name: "About", path: "/about", icon: IconInfoCircle },
    { name: "Calendar", path: "/calendar", icon: IconCalendar },
    { name: "Users", path: "/users", icon: IconTable },
    { name: "Todo", path: "/todo", icon: IconListCheck },
    { name: "Login", path: "/login", icon: IconLogin },
    { name: "Logged In", path: "/logged", icon: IconUserCheck },
  ]

  const currentPath = window.location.pathname

  return (
    <Stack gap="xs">
      {links.map((link) => (
        <NavLink
          key={link.path}
          href={link.path}
          label={link.name}
          leftSection={<link.icon size={20} stroke={1.5} />}
          active={currentPath === link.path}
          variant="filled"
        />
      ))}
    </Stack>
  )
}

export default Sidebar
