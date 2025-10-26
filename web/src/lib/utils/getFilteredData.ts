import { endpointMapping } from "./constants";
import type {
    PostFilterDataRequest,
  PostFilterDataResponse,
} from "$lib/models/dashboardModels";

export const fetchStaticFireFilteredData = async (
  req: PostFilterDataRequest,
): Promise<PostFilterDataResponse | null> => {
  try {
    const response = await fetch(endpointMapping.getFilterData, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
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
