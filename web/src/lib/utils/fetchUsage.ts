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
