package metrics

func GetLabelsKey(labels map[string]string) []string {
	var labelsKey []string
	for key := range labels {
		labelsKey = append(labelsKey, key)
	}

	return labelsKey
}
