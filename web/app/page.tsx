export default function Home() {
  return (
    <div className="flex flex-col gap-6">
      <div className="flex items-center justify-between">
        <h1 className="text-3xl font-bold tracking-tight">Dashboard</h1>
      </div>
      <div className="flex items-center justify-center h-[calc(100vh-10rem)]">
        <p className="text-muted-foreground">Select a file from the sidebar to view its data</p>
      </div>
    </div>
  )
}

