package error

func ErrMapping(err error) bool {
	allErrors := make([]error, 0)
	allErrors = append(allErrors, GeneralErrors[:]...)
	allErrors = append(allErrors, UserErrors[:]...)

	for _, item := range allErrors {
		if item.Error() == err.Error() {
			return true
		}
	}
	return false
}
