export const backendDevPort = 8080;

export const endpointMapping = Object.freeze({
  uploadStaticFireUrl: new URL(
    `http://localhost:${backendDevPort}/api/staticfire/upload`,
  ),
  getStaticFireMetadataUrl: new URL(
    `http://localhost:${backendDevPort}/api/staticfire/metadata`,
  ),
  postStaticFireColumnsUrl: new URL(
    `http://localhost:${backendDevPort}/api/staticfire/columns`,
  ),
  getUsageURL: new URL(
    `http://localhost:${backendDevPort}/api/usage`
  )
});
