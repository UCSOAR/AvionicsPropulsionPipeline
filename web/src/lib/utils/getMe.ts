import { endpointMapping } from "./constants";
import type { User } from "$lib/models/user";

export const fetchMe = async (): Promise<User | null> => {
  try {
    const response = await fetch(endpointMapping.getMeUrl, {
      method: "GET",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      return null;
    }

    return await response.json();
  } catch {
    return null;
  }
};
