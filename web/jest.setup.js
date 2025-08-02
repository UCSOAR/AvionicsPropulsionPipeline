// Polyfill import.meta for Jest environment
global.import = {};
global.import.meta = {
  env: {
    VITE_BACKEND_HOST: "http://localhost:3000",
  },
};
