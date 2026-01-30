import {
  IconHome,
  IconInfoCircle,
  IconLogin,
  IconUserCheck,
  IconCalendar,
  IconTable,
} from "@tabler/icons-react"

const Sidebar = () => {
  const links = [
    { name: "Home", path: "/", icon: IconHome },
    { name: "About", path: "/about", icon: IconInfoCircle },
    { name: "Calendar", path: "/calendar", icon: IconCalendar },
    { name: "Users", path: "/users", icon: IconTable },
    { name: "Login", path: "/login", icon: IconLogin },
    { name: "Logged In", path: "/logged", icon: IconUserCheck },
  ]

  return (
    <aside className="min-h-screen w-64 border-r border-gray-200 bg-gray-100 p-4">
      <nav>
        <ul className="space-y-2">
          {links.map((link) => (
            <li key={link.path}>
              <a
                href={link.path}
                className="flex items-center gap-2 rounded p-2 font-medium text-gray-700 hover:bg-gray-200"
              >
                <link.icon size={20} />
                <span>{link.name}</span>
              </a>
            </li>
          ))}
        </ul>
      </nav>
    </aside>
  )
}

export default Sidebar
