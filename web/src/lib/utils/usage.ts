import { endpointMapping } from "./constants"; // Adjust the path if needed

// Define a Type for the Expected Response
interface StorageUsageResponse {
    storageSizeBytes: number;
}

export async function getStorageUsage(): Promise<StorageUsageResponse | null> {
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

        const data: StorageUsageResponse = await response.json();
        return data;

    } catch (error) {
        return null;
    }
}

// Function to Format Bytes
function formatBytes(bytes: number, decimals: number = 2): string {
    if (bytes === 0) return "0 Bytes";

    const sizes = ["Bytes", "KB", "MB", "GB"];
    const i = Math.floor(Math.log(bytes) / Math.log(1024));

    if (i >= sizes.length) return `${bytes} Bytes`;

    const value = bytes / Math.pow(1024, i);
    return i === 0 || value < 1 ? `${Math.floor(value)} ${sizes[i]}` : `${value.toFixed(decimals)} ${sizes[i]}`;
}

export { formatBytes };
