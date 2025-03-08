"use client";

import * as React from "react";
import { useRouter } from "next/navigation";
import { useMetadataStore } from "../stores/metadataStore";
import { endpointMapping } from "@/utils/constants";
import { File, PanelLeftClose, PanelLeftOpen, Upload } from "lucide-react";
import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import FileUploader from "@/components/FileUploader";

export function Sidebar() {
  const router = useRouter();
  const [isCollapsed, setIsCollapsed] = React.useState(false);
  const [files, setFiles] = React.useState<Record<string, any>>({});
  const [error, setError] = React.useState<string | null>(null);
  const setMetadata = useMetadataStore((state) => state.setMetadata);

  React.useEffect(() => {
    const fetchFiles = async () => {
      try {
        const response = await fetch(endpointMapping.getStaticFireMetadataUrl);
        if (!response.ok) throw new Error("Failed to fetch files");
        const data = await response.json();
        setFiles(data);
        setError(null);
      } catch (err) {
        setError((err as Error).message || "An error occurred.");
      }
    };
    fetchFiles();
  }, []);

  // Handle file selection and navigate to Dashboard
  const handleFileClick = (name: string, metadata: any) => {
    // Set metadata in store
    setMetadata(metadata, metadata.xColumnNames[0], metadata.yColumnNames[0], name);
    
    // Navigate to Dashboard with fileId
    router.push(`/dashboard/${encodeURIComponent(name)}`);
  };

  return (
    <div className={cn("h-[calc(100vh-3.5rem)] border-r bg-background transition-all duration-300", isCollapsed ? "w-[60px]" : "w-[280px]")}>
      
      {/* Upload Section */}
      {isCollapsed ? (
        <div className="p-4 flex items-center justify-center">
          <Button variant="ghost" size="icon">
            <Upload className="h-5 w-5" />
          </Button>
        </div>
      ) : (
        <div className="p-4">
          <FileUploader />
        </div>
      )}

      {/* Sidebar Header */}
      <div className="p-4 flex justify-between items-center">
        <h2 className={cn("font-semibold", isCollapsed && "hidden")}>Files</h2>
        <Button variant="ghost" size="icon" onClick={() => setIsCollapsed(!isCollapsed)}>
          {isCollapsed ? <PanelLeftOpen className="h-4 w-4" /> : <PanelLeftClose className="h-4 w-4" />}
        </Button>
      </div>

      {/* File List */}
      <div className="overflow-auto h-[calc(100vh-7.5rem)]">
        {Object.keys(files).length > 0 ? (
          Object.entries(files).map(([name, metadata]) => (
            <Button
              key={name}
              variant="ghost"
              className="w-full justify-start p-2 mb-1 hover:bg-accent hover:text-accent-foreground"
              onClick={() => handleFileClick(name, metadata)}
            >
              <File className="h-4 w-4 mr-2" />
              {!isCollapsed && <span>{name.replace(/\.[^.]+$/, "")}</span>}
            </Button>
          ))
        ) : (
          <p className="p-4 text-sm text-muted-foreground">{error || "No uploaded files yet."}</p>
        )}
      </div>
    </div>
  );
}
