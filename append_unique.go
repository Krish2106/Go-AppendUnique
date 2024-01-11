package main 

func AppendUnique(slice []string, element string) []string {
	elementMap := make(map[string]struct{})
	for _, s := range slice {
		elementMap[s] = struct{}{}
	}

	if _, exists := elementMap[element]; !exists {
		slice = append(slice, element)
	}

	return slice
}

