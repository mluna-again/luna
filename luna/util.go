package luna

func memberOf[K comparable](enum []K, needle K) bool {
	for _, item := range enum {
		if item == needle {
			return true
		}
	}

	return false
}

func availableVariants(name LunaPet) []LunaVariant {
	switch name {
	case CAT:
		return []LunaVariant{CAT_RAGDOLL, CAT_BLACK}

	default:
		return []LunaVariant{}
	}
}

func getDefaultVariant(name LunaPet) LunaVariant {
	switch name {
	case CAT:
		return CAT_RAGDOLL

	default:
		return DEFAULT_VARIANT
	}
}
