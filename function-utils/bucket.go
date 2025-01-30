package functionUtils

const BucketName string = "gcf-v2-uploads-1026004530618.us-west1.cloudfunctions.appspot.com"
const MaxFileSize int64 = 10 << 20 // 10MB

// Represents an object stored in a cloud storage bucket.
// It contains metadata about the object such as its name, size, and last modified timestamp.
type BucketObject struct {
	Name         string `json:"name"`
	Size         int64  `json:"size"`
	LastModified string `json:"lastModified"`
}
