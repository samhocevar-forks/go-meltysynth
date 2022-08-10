package meltysynth

import (
	"math"
)

const halfPi float32 = math.Pi

const nonAudible float32 = 1.0e-3

var logNonAudible float32 = float32(math.Log(1.0e-3))

func calcTimecentsToSeconds(x float32) float32 {
	return float32(math.Pow(float64(2), (float64(1)/float64(1200))*float64(x)))
}

func calcCentsToHertz(x float32) float32 {
	return float32(float64(8.176) * math.Pow(float64(2), (float64(1)/float64(1200))*float64(x)))
}

func calcCentsToMultiplyingFactor(x float32) float32 {
	return float32(math.Pow(float64(2), (float64(1)/float64(1200))*float64(x)))
}

func calcDecibelsToLinear(x float32) float32 {
	return float32(math.Pow(float64(10), float64(0.05)*float64(x)))
}

func calcLinearToDecibels(x float32) float32 {
	return float32(float64(20) * math.Log10(float64(x)))
}

func calcKeyNumberToMultiplyingFactor(cents int, key int) float32 {
	return calcTimecentsToSeconds(float32(cents * (60 - key)))
}

func calcExpCutoff(x float64) float64 {
	if x < float64(logNonAudible) {
		return 0
	} else {
		return math.Exp(x)
	}
}

func calcClamp(value float32, min float32, max float32) float32 {
	if value < min {
		return min
	} else if value > max {
		return max
	} else {
		return value
	}
}