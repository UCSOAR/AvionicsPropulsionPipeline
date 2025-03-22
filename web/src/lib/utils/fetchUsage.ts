import { endpointMapping } from "./constants"; // Adjust the path if needed

export async function getStorageUsage() {
    try {
        const response = await fetch(endpointMapping.getUsageURL.toString(), {
            method: "GET",
            headers: {
                "Accept": "application/json",
            },
        });

        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        }

        const data = await response.json();
        console.log("Storage Size (Bytes):", data.storageSizeBytes);
    } catch (error) {
        console.error("Error fetching storage usage:", error);
    }
}


function formatBytes(bytes: number, decimals = 2): string {
    if (bytes === 0) return "0 Bytes";

    const sizes = ["Bytes", "KB", "MB", "GB", "TB", "PB"];
    const i = Math.floor(Math.log(bytes) / Math.log(1024));

    const value = bytes / Math.pow(1024, i);
    if (value >= 1) {
        return `${value.toFixed(decimals)} ${sizes[i]}`;
    }

    return `${bytes} Bytes`;
}


