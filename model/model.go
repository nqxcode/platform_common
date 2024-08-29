package model

// HasID model interface
type HasID interface {
	GetID() int64
}

// ExtractIDs extract ids from models
func ExtractIDs[T HasID](models []T) []int64 {
	modelIDs := make([]int64, 0, len(models))
	for _, model := range models {
		modelIDs = append(modelIDs, model.GetID())
	}

	return modelIDs
}
