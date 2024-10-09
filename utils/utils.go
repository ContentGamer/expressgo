package utils

func HandleError(e error) {
	if e == nil {
		return
	}

	panic(e)
}
