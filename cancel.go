package life

// Cancel 手动中止（而不等信号）
func Cancel() {
	Stop = true
	cancel()
}
