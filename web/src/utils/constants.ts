export const devPort = 8080;

const devEndpointMapping = {
  uploadStaticFireUrl: new URL(`http://localhost:${devPort}/UploadStaticFire`),
  getStaticFireMetadataUrl: new URL(
    `http://localhost:${devPort}/GetStaticFireMetadata`
  ),
  getStaticFireColumnsUrl: new URL(
    `http://localhost:${devPort}/GetStaticFireColumns`
  ),
};

const prodEndpointMapping: typeof devEndpointMapping = {
  uploadStaticFireUrl: new URL(`http://localhost:${devPort}/UploadStaticFire`),
  getStaticFireMetadataUrl: new URL(
    `http://localhost:${devPort}/GetStaticFireMetadata`
  ),
  getStaticFireColumnsUrl: new URL(
    `http://localhost:${devPort}/GetStaticFireColumns`
  ),
};

// Freeze the objects to prevent accidental modification
Object.freeze(devEndpointMapping);
Object.freeze(prodEndpointMapping);

// Set to true when testing deployed version of the functions
export const inProd = false;

// Switch between development and production endpoints
export const endpointMapping = inProd
  ? prodEndpointMapping
  : devEndpointMapping;
