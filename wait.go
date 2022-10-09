package life

// Wait ...
func Wait() {
	for {
		if Sleep(100) != nil {
			break
		}
	}
}
