package life

// Cancel ...
func Cancel() {
	Stop = true
	cancel()
}
