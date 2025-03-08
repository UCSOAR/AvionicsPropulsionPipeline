import type React from "react"
import type { Metadata } from "next"
import { Inter } from "next/font/google"
import "./globals.css"
import { ThemeProvider } from "@/components/theme-provider"
import { Navbar } from "@/components/navbar"
import { Sidebar } from "@/components/sidebar"

const inter = Inter({ subsets: ["latin"] })

export const metadata: Metadata = {
  title: "Data Dashboard",
  description: "File-based data visualization dashboard",
    generator: 'v0.dev'
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en" suppressHydrationWarning>
      <body className={inter.className}>
        <ThemeProvider
  attribute="class"
  defaultTheme="system"
  enableSystem
  disableTransitionOnChange
>
          <div className="min-h-screen flex flex-col">
            <Navbar />
            <div className="flex-1 flex">
              <Sidebar />
              <main className="flex-1 p-6">{children}</main>
            </div>
          </div>
        </ThemeProvider>
      </body>
    </html>
  )
}



import './globals.css'