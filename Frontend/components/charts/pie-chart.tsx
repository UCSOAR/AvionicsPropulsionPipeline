"use client"

import { PieChart as RechartsPieChart, Pie, Cell, ResponsiveContainer, Legend, Tooltip } from "recharts"
import { useTheme } from "next-themes"

const data = [
  { name: "Direct", value: 400 },
  { name: "Social", value: 300 },
  { name: "Organic", value: 300 },
  { name: "Referral", value: 200 },
]

// Use colors that complement our red theme
const COLORS = ["#dc2626", "#ea580c", "#d97706", "#ca8a04"]
const COLORS_DARK = ["#ef4444", "#f97316", "#f59e0b", "#eab308"]

export function PieChart() {
  const { resolvedTheme } = useTheme()
  const isDark = resolvedTheme === "dark"

  const textColor = isDark ? "#f8fafc" : "#0f172a"
  const pieColors = isDark ? COLORS_DARK : COLORS

  return (
    <ResponsiveContainer width="100%" height="100%">
      <RechartsPieChart>
        <Pie
          data={data}
          cx="50%"
          cy="50%"
          labelLine={false}
          outerRadius={80}
          fill="#8884d8"
          dataKey="value"
          label={({ name, percent }) => `${name} ${(percent * 100).toFixed(0)}%`}
        >
          {data.map((entry, index) => (
            <Cell key={`cell-${index}`} fill={pieColors[index % pieColors.length]} />
          ))}
        </Pie>
        <Tooltip
          contentStyle={{
            backgroundColor: isDark ? "#1e1e1e" : "#ffffff",
            color: textColor,
            borderRadius: "6px",
            boxShadow: "0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06)",
          }}
        />
        <Legend />
      </RechartsPieChart>
    </ResponsiveContainer>
  )
}

