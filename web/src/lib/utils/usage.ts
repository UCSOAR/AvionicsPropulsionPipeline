import { endpointMapping } from "./constants"; // Adjust the path if needed

// Define a Type for the Expected Response
interface StorageUsageResponse {
  storageSizeBytes: number;
}

export async function getStorageUsage(): Promise<StorageUsageResponse | null> {
  try {
    const response = await fetch(endpointMapping.getUsageURL.toString(), {
      method: "GET",
      credentials: "include",
      headers: {
        Accept: "application/json",
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
