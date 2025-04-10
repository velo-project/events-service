package io

type CreateEventInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (cei *CreateEventInput) Validate() []string {
	errors := []string{}

	if cei.Name == "" {
		errors = append(errors, "Name is required")
	}
	if cei.Description == "" {
		errors = append(errors, "Description is required")
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}
