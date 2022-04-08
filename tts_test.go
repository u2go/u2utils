package u2utils

import "testing"

func TestTtsSpeak(t *testing.T) {
	PanicOnError(TtsSpeak("1号指标下降20%", "tmp", "zh"))
}
