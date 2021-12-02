package day1

// CountMeasurementsIncreases takes a list of int and return the amount of increases between two following numbers
func CountMeasurementsIncreases(l []int) int {
	var count int
	for k, _ := range l {
		if k == 0 {
			continue
		}
		if l[k] > l[k-1] {
			count++
		}
	}

	return count
}

// CountMeasurementsIncreasesSlidingWindow takes a list of int and return the amount of increases between two following numbers
func CountMeasurementsIncreasesSlidingWindow(l []int, windowSize int) int {
	var listOfSummedByWindowMeasurements []int
	for k, _ := range l {

		if k < windowSize {
			continue
		}

		elementsWindow := make([]int, windowSize)
		elementsWindow = l[k-windowSize:k]

		listOfSummedByWindowMeasurements = append(listOfSummedByWindowMeasurements, windowSum(elementsWindow))
	}

	lastWindow := l[:len(l)-windowSize]
	listOfSummedByWindowMeasurements = append(listOfSummedByWindowMeasurements, windowSum(lastWindow))
	return CountMeasurementsIncreases(listOfSummedByWindowMeasurements)
}

func windowSum(l []int) int {
	var s int
	for _, v := range l {
		s += v
	}

	return s
}