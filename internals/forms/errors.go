package forms

type errors map[string][]string

// Add appends a new error message by field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get returns the first error message if any
func (e errors) Get(field string) string {
	es := e[field]

	if len(es) == 0 {
		return ""
	}

	return es[0]
}