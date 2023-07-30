package rest_err

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

//Prar satisfazer a interface de error
func (*RestErr) Error() string {
	return r.Message
}

func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

// Metodos para erros
func NewBadRequestError(massage string) *RestErr {

	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValitationError(massage string, causes []Causes) *RestErr {

	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes: causes,
	}
}

func NewInternalSeverError(message string) *RestErr {
		return &RestErr{
		Message: message,
		Err:     "intenal_sever_error",
		Code:    http.StatusBadRequest,
		}

}

func NewNotFoundError(message string) *RestErr {

		return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusBadRequest,
		}
}

func NewForbiddenError(message string) *RestErr {

		return &RestErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusBadRequest,
		}
}