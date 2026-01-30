import { useCalendarApp, ScheduleXCalendar } from "@schedule-x/react"
import {
  createViewDay,
  createViewMonthAgenda,
  createViewMonthGrid,
  createViewWeek,
} from "@schedule-x/calendar"
import { createEventsServicePlugin } from "@schedule-x/events-service"
import Layout from "~/components/Layout"
import "temporal-polyfill/global"

const CalendarPage = () => {
  const eventsService = createEventsServicePlugin()
  const calendar = useCalendarApp({
    views: [
      createViewDay(),
      createViewWeek(),
      createViewMonthGrid(),
      createViewMonthAgenda(),
    ],
    events: [
      {
        id: "1",
        title: "Event 1",
        start: Temporal.ZonedDateTime.from(
          "2023-12-04T10:05:00+01:00[Europe/Berlin]"
        ),
        end: Temporal.ZonedDateTime.from(
          "2023-12-04T10:35:00+01:00[Europe/Berlin]"
        ),
      },
    ],
    plugins: [eventsService],
  })

  return (
    <Layout>
      <div className="container mx-auto">
        <h1 className="mb-4 text-2xl font-bold">Calendar</h1>
        <div>
          <ScheduleXCalendar calendarApp={calendar} />
        </div>
      </div>
    </Layout>
  )
}

export default CalendarPage
