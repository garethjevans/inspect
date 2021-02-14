package kube

// ImageLister an interface to wrap getting an list of images from a namespace.
type ImageLister interface {
	// GetImagesForNamespace returns a list of images for a namespace.
	GetImagesForNamespace(namespace string) ([]string, error)
}
