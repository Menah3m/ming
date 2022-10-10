package errors

/*
   @Auth: menah3m
   @Desc:
*/

type Errors struct {
	errors map[string][]string
}

func New() *Errors {
	return &Errors{
		errors: map[string][]string{},
	}
}

//AddError 新增错误
func (e *Errors) AddError(key, err string) {
	if _, ok := e.errors[key]; !ok {
		e.errors[key] = make([]string, 0, 5)
	}
	e.errors[key] = append(e.errors[key], err)
}

//Errors 返回错误
func (e *Errors) Errors() map[string][]string {
	return e.errors
}

//ErrorsByKey 根据key获取错误
func (e *Errors) ErrorsByKey(key string) []string {
	return e.errors[key]
}

func (e *Errors) HasErrors() bool {
	return len(e.errors) != 0
}
