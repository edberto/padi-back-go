package helper

func Wrap(data interface{}, message string) map[string]interface{} {
	return map[string]interface{}{
		"data":    data,
		"message": message,
	}
}
