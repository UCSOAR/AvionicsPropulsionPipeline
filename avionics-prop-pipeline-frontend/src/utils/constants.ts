interface EndpointMapping {
  bucketUploadUrl: URL;
}

export const devPort = 8080;

const devEndpointMapping: EndpointMapping = {
  bucketUploadUrl: new URL(`http://localhost:${devPort}/BucketUpload`),
};

const prodEndpointMapping: EndpointMapping = {
  bucketUploadUrl: new URL(
    "https://us-west1-avionic-propulsion-pipeline.cloudfunctions.net/BucketUpload"
  ),
};

// Set to true when testing deployed version of the functions
export const inProd = false;

// Switch between development and production endpoints
export const endpointMapping = inProd
  ? prodEndpointMapping
  : devEndpointMapping;
