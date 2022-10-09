package life

// Wait 等待结束，收到信号或者被 Cancel()
func Wait() {
	for {
		if Sleep(100) != nil {
			break
		}
	}
}
