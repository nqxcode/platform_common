package model

type HasID interface {
	GetID() int64
}

func ExtractIDs[T HasID](models []T) []int64 {
	modelIDs := make([]int64, 0, len(models))
	for _, model := range models {
		modelIDs = append(modelIDs, model.GetID())
	}

	return modelIDs
}
