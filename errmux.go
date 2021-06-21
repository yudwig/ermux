package ermux

// First returns the first error(!= nil) of input.
func First(errs []error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

// Last returns the last error(!= nil) of input.
func Last(errs []error) error {
	l := len(errs)
	for i := range errs {
		if errs[l-i-1] != nil {
			return errs[l-i-1]
		}
	}
	return nil
}

// Some returns true if input has error(!= nil).
func Some(errs []error) bool {
	for _, err := range errs {
		if err != nil {
			return true
		}
	}
	return false
}

// Filter removes empty(= nil) elements from input error slice.
func Filter(errs []error) []error {
	var res []error
	for _, err := range errs {
		if err != nil {
			res = append(res, err)
		}
	}
	return res
}
