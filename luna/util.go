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

func getSelectedVariant(name LunaPet, variant LunaVariant) LunaVariant {
	switch name {
	case CAT:
		if variant == "black" {
			return DEFAULT_VARIANT
		}
		return variant

	default:
		return DEFAULT_VARIANT
	}
}

func translateVariant(name LunaVariant) LunaVariant {
	if name == "black" {
		return CAT_BLACK
	}

	return name
}

func nextVariant(variant LunaVariant, pet LunaPet) LunaVariant {
	if !hasNextVariant(pet) {
		return variant
	}

	switch pet {
	case CAT:
		switch variant {
		case CAT_BLACK:
			return CAT_RAGDOLL

		case CAT_RAGDOLL:
			return CAT_BLACK
		}
	}

	return DEFAULT_VARIANT
}

func hasNextVariant(pet LunaPet) bool {
	return pet == CAT
}
