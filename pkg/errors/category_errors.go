package errors

type CategoryAlreadyExists struct {
}

func (e *CategoryAlreadyExists) Error() string {
	return "Category with given name already exists"
}

type CategoryNotFound struct {
}

func (e *CategoryNotFound) Error() string {
	return "Category not found"
}
