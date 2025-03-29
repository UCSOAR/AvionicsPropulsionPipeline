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

    const sizes = ["Bytes", "KB", "MB", "GB"];
    const i = Math.floor(Math.log(bytes) / Math.log(1024));

    // Ensure we don't get out of bounds
    if (i >= sizes.length) return `${bytes} Bytes`;

    const value = bytes / Math.pow(1024, i);

    // For values less than 1 KB, don't display decimals
    if (i === 0 || value < 1) {
        return `${Math.floor(value)} ${sizes[i]}`; // No decimals
    }

    // For values above 1 KB, use the specified decimals
    return `${value.toFixed(decimals)} ${sizes[i]}`;
}

export { formatBytes };



// setting up a frame work jest set up tests for the function formatbytes