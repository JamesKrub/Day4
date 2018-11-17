package stats_test

import (
	stats "KBTGCourse/day4/stats"
	"testing"
)

func TestMin(t *testing.T) {
	setOfnumber := []float64{5, 92, 45, 111, 39, 1}
	r := stats.Min(setOfnumber)
	if r != 1 {
		t.Error("Expected 1 but got ", r)
	}
}

func TestMax(t *testing.T) {
	setOfnumber := []float64{5, 92, 45, 111, 39, 1}
	r := stats.Max(setOfnumber)
	if r != 111 {
		t.Error("Expected 111 but got ", r)
	}
}

func TestAverage(t *testing.T) {
	setOfnumber := []float64{5, 92, 45, 111, 39, 1}
	r := stats.Avg(setOfnumber)
	if r != 48.833333333333336 {
		t.Error("Expected 48.8333 but got ", r)
	}
}

func TestReporting(t *testing.T) {
	setOfnumber := []float64{5, 92, 45, 111, 39, 1}

	tcs := stats.Report{111, 1, 48.833333333333336}

	r := stats.Reporting(setOfnumber)
	if r != tcs {
		t.Errorf("Expected %v but got %v", tcs, r)
	}
}
