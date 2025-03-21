export const backendDevPort = 8080;

export const endpointMapping = Object.freeze({
  uploadStaticFireUrl: new URL(
    `http://localhost:${backendDevPort}/api/staticfire/upload`
  ),
  getStaticFireMetadataUrl: new URL(
    `http://localhost:${backendDevPort}/api/staticfire/metadata`
  ),
  getStaticFireColumnsUrl: new URL(
    `http://localhost:${backendDevPort}/api/staticfire/columns`
  ),
});
