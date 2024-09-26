package validate

import "context"

// Condition is a function that validates a condition
type Condition func(ctx context.Context) error

// Validate validates the conditions
func Validate(ctx context.Context, conds ...Condition) error {
	ve := NewValidationErrors()

	for _, c := range conds {
		err := c(ctx)
		if err != nil {
			if IsValidationError(err) {
				ve.addError(err.Error())
				continue
			}

			return err
		}
	}

	if ve.Messages == nil {
		return nil
	}

	return ve
}
