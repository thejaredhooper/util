package query

import (
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
)

func And(conditions []bson.M) bson.M {
	return bson.M{"$and": conditions}
}

func Or(conditions []bson.M) bson.M {
	return bson.M{"$or": conditions}
}

func CaseInsensitiveRegex(value string) bson.M {
	return bson.M{
		"$regex":   fmt.Sprintf("%v", value),
		"$options": "i",
	}
}
