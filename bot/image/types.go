package image

const (
	mimeJPEG = "image/jpeg"
	mimePNG  = "image/png"
)

var supportedMIMEs = map[string]bool{
	mimeJPEG: true,
	mimePNG:  true,
}

// IsSupportedType returns whether the MIME is a supported one for images.
func IsSupportedType(mime string) bool {
	return supportedMIMEs[mime]
}
