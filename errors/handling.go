package errors

// Handles an error by logging it
func HandleError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
