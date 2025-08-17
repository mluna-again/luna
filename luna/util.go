package luna

func memberOf[K comparable](enum []K, needle K) bool {
	for _, item := range enum {
		if item == needle {
			return true
		}
	}

	return false
}
