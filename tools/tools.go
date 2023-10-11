package tools

func StringInSlice(element string, slice []string) (int, bool) {
	for index, s := range slice {
		if s == element {
			return index, true
		}
	}
	return -1, false
}

func StringInMapKeys(element string, m map[string]string) bool {
	_, exist := m[element]
	return exist
}
