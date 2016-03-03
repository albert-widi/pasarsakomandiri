package response

type SimpleResponse struct {
	Status string
	Message string
}

type LoginResponse struct {
	Status string
	Message string
	Token string
}

type DataResponse struct {
	Status string
	Message string
	Data map[string]interface{}
}

func NewSimpleResponse(status string, message string) SimpleResponse {
	var simpleResponse = SimpleResponse{}
	simpleResponse.Status = status
	simpleResponse.Message = message
	return simpleResponse
}

func NewLoginResponse(status string, message string, token string) LoginResponse {
	var loginResponse = LoginResponse{}
	loginResponse.Status = status
	loginResponse.Message = message
	loginResponse.Token = token
	return loginResponse
}

func NewDataResponse(status string, message string, data map[string]interface{}) DataResponse {
	var dataResponse = DataResponse{}
	dataResponse.Status = status
	dataResponse.Message = message
	dataResponse.Data = data
	return dataResponse
}
