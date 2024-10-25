package utils

func IsAllowedMimeType(mimeType string) bool {
	allowedMimeTypes := []string{"image/png", "image/jpeg"}
	for _, allowed := range allowedMimeTypes {
		if mimeType == allowed {
			return true
		}
	}
	return false
}
