import { Rocket } from "lucide-react"

export function Logo() {
  return (
    <div className="flex items-center gap-2">
      <div className="h-8 w-8 rounded-md bg-primary flex items-center justify-center shadow-sm">
        <Rocket className="h-5 w-5 text-primary-foreground" />
      </div>
      <span className="font-bold text-xl">SOAR</span>
    </div>
  )
}

