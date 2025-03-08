"use client"
import { useState } from "react"
import Link from "next/link"
import { usePathname } from "next/navigation"
import { Menu, X } from "lucide-react"
import { cn } from "@/lib/utils"
import { Logo } from "@/components/logo"
import { ModeToggle } from "@/components/mode-toggle"
import { Button } from "@/components/ui/button"

const navItems = [
  { name: "pyTac", href: "/pytac" },
  { name: "pyHac", href: "/pyhac" },
  { name: "mapleleaf", href: "/mapleleaf" },
]

export function Navbar() {
  const pathname = usePathname()
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false)

  return (
    <header className="sticky top-0 z-50 w-full border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
      <div className="container mx-auto px-4 flex h-14 items-center justify-between w-full">
        
        {/* Left: Logo (Fully Left-Aligned) */}
        <div className="flex items-center justify-start w-1/3">
        <img src="/soar-logo.svg" alt="Soar Logo" draggable="false" className="h-10 w-auto" />

        </div>

        {/* Center: Desktop Navigation (Fully Centered) */}
        <nav className="hidden md:flex justify-center w-1/3">
          <ul className="flex space-x-8">
            {navItems.map((item) => (
              <li key={item.href}>
                <Link
                  href={item.href}
                  className={cn(
                    "text-sm font-medium transition-colors hover:text-primary",
                    pathname === item.href ? "text-primary font-semibold" : "text-muted-foreground"
                  )}
                >
                  {item.name}
                </Link>
              </li>
            ))}
          </ul>
        </nav>

        {/* Right: Mode Toggle (Fully Right-Aligned) */}
        <div className="flex items-center justify-end w-1/3">
          <ModeToggle />

          {/* Mobile Menu Button */}
          <Button
            variant="ghost"
            size="icon"
            className="ml-4 md:hidden"
            onClick={() => setMobileMenuOpen(!mobileMenuOpen)}
          >
            {mobileMenuOpen ? <X className="h-5 w-5" /> : <Menu className="h-5 w-5" />}
            <span className="sr-only">Toggle menu</span>
          </Button>
        </div>
      </div>

      {/* Mobile Navigation */}
      {mobileMenuOpen && (
        <div className="md:hidden bg-background text-muted-foreground">
          <div className="container mx-auto py-4 border-t">
            <nav className="flex flex-col space-y-4 items-center">
              {navItems.map((item) => (
                <Link
                  key={item.href}
                  href={item.href}
                  className={cn(
                    "text-sm font-medium transition-colors hover:text-primary py-2",
                    pathname === item.href ? "text-primary font-semibold" : "text-muted-foreground"
                  )}
                  onClick={() => setMobileMenuOpen(false)}
                >
                  {item.name}
                </Link>
              ))}
            </nav>
          </div>
        </div>
      )}
    </header>
  )
}

