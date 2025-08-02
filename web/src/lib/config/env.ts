const getBackendHost = (): string => {
  if (
    typeof import.meta !== "undefined" &&
    import.meta.env?.VITE_BACKEND_HOST
  ) {
    return import.meta.env.VITE_BACKEND_HOST;
  } else {
    return "http://localhost:8080";
  }
};

export const backendHost = getBackendHost();
