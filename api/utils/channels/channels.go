package channels

// OK returns if a operation was successful
func OK(done chan bool) bool {
	select {
	case ok := <-done:
		if ok {
			return ok
		}

	}
	return false
}
