const getBackendHost = (): string => {
    if (typeof import.meta !== 'undefined' && import.meta.env?.VITE_BACKEND_HOST) {
        return import.meta.env.VITE_BACKEND_HOST;
    }

    // Fallback for Node.js / test environment
<<<<<<< HEAD
    return import.meta.env.VITE_BACKEND_HOST ?? 'http://localhost:8080';
=======
    return process.env.VITE_BACKEND_HOST ?? 'http://localhost:8080';
>>>>>>> 1e8dd8df07ff932ac7b4307293e3137af452d9d9
};

export const backendHost = getBackendHost();
