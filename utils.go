package aircallgo

func isValidAvailabilityStatus(v string) bool {
	for _, s := range []string{"available", "custom", "unavailable"} {
		if v == s {
			return true
		}
	}
	return false
}
