export const devPort = 8080;

const devEndpointMapping = {
  bucketUploadUrl: new URL(`http://localhost:${devPort}/BucketUpload`),
  getBucketUploadsUrl: new URL(`http://localhost:${devPort}/GetBucketUploads`),
  downloadBucketObjectUrl: new URL(
    `http://localhost:${devPort}/DownloadBucketObject`
  ),
};

const prodEndpointMapping: typeof devEndpointMapping = {
  bucketUploadUrl: new URL(
    "https://us-west1-avionic-propulsion-pipeline.cloudfunctions.net/BucketUpload"
  ),
  getBucketUploadsUrl: new URL(
    "https://us-west1-avionic-propulsion-pipeline.cloudfunctions.net/GetBucketUploads"
  ),
  downloadBucketObjectUrl: new URL(
    "https://us-west1-avionic-propulsion-pipeline.cloudfunctions.net/DownloadBucketObject"
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
