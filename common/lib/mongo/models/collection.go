package models

type Collection struct {
	Name string `bson:"name" json:"name"`
}

func CollectionFilter(collections []Collection, filter func(Collection) bool) (filtered []Collection) {
	for _, collection := range collections {
		if filter(collection) {
			filtered = append(filtered, collection)
		}
	}
	return
}
