"use client"

import {
  BarChart as RechartsBarChart,
  Bar,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
  Legend,
} from "recharts"
import { useTheme } from "next-themes"

const data = [
  { name: "Electronics", value: 4000 },
  { name: "Clothing", value: 3000 },
  { name: "Food", value: 2000 },
  { name: "Books", value: 2780 },
  { name: "Home", value: 1890 },
  { name: "Beauty", value: 2390 },
]

export function BarChart() {
  const { resolvedTheme } = useTheme()
  const isDark = resolvedTheme === "dark"

  const textColor = isDark ? "#f8fafc" : "#0f172a"
  const gridColor = isDark ? "#334155" : "#e2e8f0"

  // Use red shade that matches our theme
  const barColor = isDark ? "#ff5a5a" : "#dc2626"

  return (
    <ResponsiveContainer width="100%" height="100%">
      <RechartsBarChart
        data={data}
        margin={{
          top: 5,
          right: 30,
          left: 20,
          bottom: 5,
        }}
      >
        <CartesianGrid strokeDasharray="3 3" stroke={gridColor} />
        <XAxis dataKey="name" stroke={textColor} />
        <YAxis stroke={textColor} />
        <Tooltip
          contentStyle={{
            backgroundColor: isDark ? "#1e1e1e" : "#ffffff",
            color: textColor,
            borderColor: gridColor,
            borderRadius: "6px",
            boxShadow: "0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06)",
          }}
        />
        <Legend />
        <Bar dataKey="value" fill={barColor} radius={[4, 4, 0, 0]} />
      </RechartsBarChart>
    </ResponsiveContainer>
  )
}

