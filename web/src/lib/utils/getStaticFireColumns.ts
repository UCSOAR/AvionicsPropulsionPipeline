import { endpointMapping } from "./constants";
import type {
  PostStaticFireColumnsRequest,
  PostStaticFireColumnsResponse,
} from "$lib/models/dashboardModels";

export const fetchStaticFireColumns = async (
  req: PostStaticFireColumnsRequest
): Promise<PostStaticFireColumnsResponse | null> => {
  try {
    const response = await fetch(endpointMapping.postStaticFireColumnsUrl, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(req),
    });

    if (!response.ok) {
      return null;
    }

    return await response.json();
  } catch {
    return null;
  }
};
