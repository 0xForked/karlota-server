package utils

type FormRequest interface {
	Validate(model interface{}, err error) map[string]string
}

type FormRequestResult struct {
	Field   string
	JsonTag string
	Message string
}
