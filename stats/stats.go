package stats

func Min(nums []float64) float64 {
	var min float64 = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}

func Max(nums []float64) float64 {
	var max float64 = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func Avg(nums []float64) float64 {
	var avg float64
	for _, v := range nums {
		avg += v
	}
	return avg / float64(len(nums))
}

type Report struct {
	Max     float64
	Min     float64
	Average float64
}

func Reporting(nums []float64) Report {
	report := Report{
		Max:     Max(nums),
		Min:     Min(nums),
		Average: Avg(nums),
	}

	return report
}
