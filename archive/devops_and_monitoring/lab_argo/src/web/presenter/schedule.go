package presenter

func ErrorResponse(err error) map[string]string {
	return map[string]string{"error": err.Error()}
}

func SuccessResponse(data interface{}) interface{} {
	return data
}
