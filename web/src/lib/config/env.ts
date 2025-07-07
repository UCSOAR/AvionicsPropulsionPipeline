const getBackendHost = (): string => {
  if (
    typeof import.meta !== "undefined" &&
    import.meta.env?.VITE_BACKEND_HOST
  ) {
    return import.meta.env.VITE_BACKEND_HOST;
  }

  // Fallback for Node.js / test environment
  return process.env.VITE_BACKEND_HOST ?? "http://localhost:8080";
};

export const backendHost = getBackendHost();
