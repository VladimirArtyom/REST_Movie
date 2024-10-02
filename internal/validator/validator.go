package validator

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) IsValid() bool {
	return len(v.Errors) == 0
}
func (v *Validator) Check(ok bool, key string, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

func (v *Validator) AddError(key string, message string) {
	v.Errors[key] = message
}

func (v *Validator) IsUnique(values []string) bool {

	uniqueValues := make(map[string]bool)
	for _, value := range values {
		uniqueValues[value] = true
	}

	return len(uniqueValues) == len(values)

}
