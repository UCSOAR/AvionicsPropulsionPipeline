export const devPort = 8080;

const devEndpointMapping = {
  uploadStaticFireUrl: new URL(
    `http://localhost:${devPort}/api/staticfire/upload`
  ),
  getStaticFireMetadataUrl: new URL(
    `http://localhost:${devPort}/api/staticfire/metadata`
  ),
  getStaticFireColumnsUrl: new URL(
    `http://localhost:${devPort}/api/staticfire/columns`
  ),
};

const prodEndpointMapping: typeof devEndpointMapping = devEndpointMapping; // For now...

// Freeze the objects to prevent accidental modification
Object.freeze(devEndpointMapping);
Object.freeze(prodEndpointMapping);

// Set to true when testing deployed version of the functions
export const inProd = false;

// Switch between development and production endpoints
export const endpointMapping = inProd
  ? prodEndpointMapping
  : devEndpointMapping;
