"use client";

import React, { useState, useRef } from "react";
import { UploadCloud, Trash2 } from "lucide-react";
import { endpointMapping } from "@/utils/constants";
import { Button } from "@/components/ui/button";
import { Progress } from "@/components/ui/progress";

export default function FileUploader() {
  const fileInputRef = useRef<HTMLInputElement>(null);
  const [selectedFiles, setSelectedFiles] = useState<File[]>([]);
  const [uploading, setUploading] = useState(false);
  const [uploadProgress, setUploadProgress] = useState(0);
  const [showConfirmPopup, setShowConfirmPopup] = useState(false);

  // Handle file selection
  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const files = event.target.files ? Array.from(event.target.files) : [];
    if (files.length > 0) {
      setSelectedFiles(files);
      setShowConfirmPopup(true);
    }
  };

  // Handle file selection via button click
  const handleClick = () => {
    fileInputRef.current?.click();
  };

  // Handles file upload confirmation
  const confirmUpload = async () => {
    setShowConfirmPopup(false);
    if (selectedFiles.length === 0) return;

    setUploading(true);
    setUploadProgress(0);

    for (const file of selectedFiles) {
      const formData = new FormData();
      formData.append("file", file);

      try {
        const response = await fetch(endpointMapping.uploadStaticFireUrl, {
          method: "POST",
          body: formData,
        });

        if (response.ok) {
          console.log(`File ${file.name} uploaded successfully`);
        } else {
          console.error(`Failed to upload file ${file.name}: ${await response.text()}`);
        }
      } catch (error) {
        console.error(`Error uploading file ${file.name}:`, error);
      }
    }

    // Simulate progress bar
    const interval = setInterval(() => {
      setUploadProgress((prev) => {
        if (prev >= 100) {
          clearInterval(interval);
          setUploading(false);
          setSelectedFiles([]);
          return 100;
        }
        return prev + 10;
      });
    }, 500);
  };

  return (
    <div className="px-4 space-y-2">
      {/* Hidden File Input */}
      <input type="file" ref={fileInputRef} onChange={handleFileChange} className="hidden" />

      {/* Upload Button */}
      <Button variant="outline" className="w-full justify-start" onClick={handleClick} disabled={uploading}>
        <UploadCloud className="mr-2 h-4 w-4" />
        {uploading ? "Uploading..." : "Upload File"}
      </Button>

      {/* Upload Progress Bar */}
      {uploading && (
        <div className="space-y-1">
          <Progress value={uploadProgress} className="h-1" />
          <p className="text-xs text-muted-foreground">{uploadProgress}% complete</p>
        </div>
      )}

      {/* Confirmation Popup */}
      {showConfirmPopup && (
        <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
          <div className="bg-white p-6 rounded-lg shadow-lg">
            <h2 className="text-lg font-semibold mb-4">Confirm Upload</h2>
            <p className="text-gray-600">Are you sure you want to upload {selectedFiles.length} files?</p>
            <div className="mt-4 flex justify-end space-x-4">
              <Button variant="ghost" onClick={() => setShowConfirmPopup(false)}>No</Button>
              <Button variant="outline" onClick={confirmUpload}>Yes</Button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
