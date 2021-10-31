package algo_os

import "testing"

func TestRunScheduling(t *testing.T) {
	pcbs := InitProcess()
	RunScheduling(pcbs)
}
