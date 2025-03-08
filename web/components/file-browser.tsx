"use client"

import { useState } from "react"
import { useRouter } from "next/navigation"
import { ChevronRight, File, Folder, FolderOpen } from "lucide-react"

import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"

// Mock data structure for files and folders
const mockFileSystem = [
  {
    id: "folder-1",
    name: "Financial Reports",
    type: "folder",
    children: [
      { id: "file-1", name: "Q1_2023.csv", type: "file" },
      { id: "file-2", name: "Q2_2023.csv", type: "file" },
      { id: "file-3", name: "Q3_2023.csv", type: "file" },
    ],
  },
  {
    id: "folder-2",
    name: "User Analytics",
    type: "folder",
    children: [
      { id: "file-4", name: "user_engagement.csv", type: "file" },
      { id: "file-5", name: "user_demographics.csv", type: "file" },
    ],
  },
  {
    id: "folder-3",
    name: "Sales Data",
    type: "folder",
    children: [
      { id: "file-6", name: "north_region.csv", type: "file" },
      { id: "file-7", name: "south_region.csv", type: "file" },
      { id: "file-8", name: "east_region.csv", type: "file" },
      { id: "file-9", name: "west_region.csv", type: "file" },
    ],
  },
  { id: "file-10", name: "annual_report_2023.csv", type: "file" },
]

export function FileBrowser() {
  const router = useRouter()
  const [expandedFolders, setExpandedFolders] = useState<string[]>([])

  const toggleFolder = (folderId: string) => {
    setExpandedFolders((prev) => (prev.includes(folderId) ? prev.filter((id) => id !== folderId) : [...prev, folderId]))
  }

  const handleFileSelect = (fileName: string) => {
    router.push(`/dashboard/${encodeURIComponent(fileName)}`)
  }

  const renderFileSystem = (items: any[], level = 0) => {
    return items.map((item) => (
      <div key={item.id} style={{ marginLeft: `${level * 16}px` }}>
        {item.type === "folder" ? (
          <div>
            <Button variant="ghost" className="w-full justify-start p-2 mb-1" onClick={() => toggleFolder(item.id)}>
              <ChevronRight
                className={`h-4 w-4 mr-2 transition-transform ${expandedFolders.includes(item.id) ? "rotate-90" : ""}`}
              />
              {expandedFolders.includes(item.id) ? (
                <FolderOpen className="h-4 w-4 mr-2" />
              ) : (
                <Folder className="h-4 w-4 mr-2" />
              )}
              {item.name}
            </Button>
            {expandedFolders.includes(item.id) && item.children && (
              <div className="ml-4">{renderFileSystem(item.children, level + 1)}</div>
            )}
          </div>
        ) : (
          <Button variant="ghost" className="w-full justify-start p-2 mb-1" onClick={() => handleFileSelect(item.name)}>
            <File className="h-4 w-4 mr-2 ml-6" />
            {item.name}
          </Button>
        )}
      </div>
    ))
  }

  return (
    <Card className="w-full">
      <CardHeader>
        <CardTitle>File Browser</CardTitle>
        <CardDescription>Select a file to view its dashboard</CardDescription>
      </CardHeader>
      <CardContent>
        <div className="border rounded-md">{renderFileSystem(mockFileSystem)}</div>
      </CardContent>
    </Card>
  )
}

