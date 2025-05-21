export const backendDevPort = 8080;

export const redirectUriParam = "redirect_uri";
export const endpointMapping = Object.freeze({
  getGoogleLoginUrl: new URL(
    `http://localhost:${backendDevPort}/auth/google/login`,
  ),
  getMeUrl: new URL(`http://localhost:${backendDevPort}/auth/me`),
  postLogoutUrl: new URL(`http://localhost:${backendDevPort}/auth/logout`),
  uploadStaticFireUrl: new URL(
    `http://localhost:${backendDevPort}/api/staticfire/upload`,
  ),
  getStaticFireMetadataUrl: new URL(
    `http://localhost:${backendDevPort}/api/staticfire/metadata`,
  ),
  postStaticFireColumnsUrl: new URL(
    `http://localhost:${backendDevPort}/api/staticfire/columns`,
  ),
  getUsageURL: new URL(`http://localhost:${backendDevPort}/api/usage`),
});
