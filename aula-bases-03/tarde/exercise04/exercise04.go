package exercise04

import (
	"fmt"
	"math/rand"
	"time"
)

func SolveExercise04() {
	var timeMeasurements = map[string]map[string]int64{
		"Insertion": {
			"100":   0,
			"1000":  0,
			"10000": 0,
		},
		"Grouping": {
			"100":   0,
			"1000":  0,
			"10000": 0,
		},
		"Selection": {
			"100":   0,
			"1000":  0,
			"10000": 0,
		},
	}

	for i := 0; i < 3; i++ {

		switch i {
		case 0:
			var1 := rand.Perm(100)

			timeNow := time.Now()
			sortViaInsertion(var1)
			timeMeasurements["Insertion"]["100"] = time.Since(timeNow).Nanoseconds()

			timeNow = time.Now()
			sortViaGrouping(var1)
			timeMeasurements["Grouping"]["100"] = time.Since(timeNow).Nanoseconds()

			timeNow = time.Now()
			sortViaSelection(var1)
			timeMeasurements["Selection"]["100"] = time.Since(timeNow).Nanoseconds()

		case 1:
			var2 := rand.Perm(1000)

			timeNow := time.Now()
			sortViaInsertion(var2)
			timeMeasurements["Insertion"]["1000"] = time.Since(timeNow).Nanoseconds()

			timeNow = time.Now()
			sortViaGrouping(var2)
			timeMeasurements["Grouping"]["1000"] = time.Since(timeNow).Nanoseconds()

			timeNow = time.Now()
			sortViaSelection(var2)
			timeMeasurements["Selection"]["1000"] = time.Since(timeNow).Nanoseconds()

		case 2:
			var3 := rand.Perm(10000)

			timeNow := time.Now()
			sortViaInsertion(var3)
			timeMeasurements["Insertion"]["10000"] = time.Since(timeNow).Nanoseconds()

			timeNow = time.Now()
			sortViaGrouping(var3)
			timeMeasurements["Grouping"]["10000"] = time.Since(timeNow).Nanoseconds()

			timeNow = time.Now()
			sortViaSelection(var3)
			timeMeasurements["Selection"]["10000"] = time.Since(timeNow).Nanoseconds()

		}

	}

	fmt.Printf("Comparing sorting algorithms\n")
	fmt.Printf("Quantity: 100\n")
	fmt.Printf("%-8s %-8s %-8s\n", "Insertion", "Grouping", "Selection")
	fmt.Printf("%-8d %-8d %-8d\n", timeMeasurements["Insertion"]["100"], timeMeasurements["Grouping"]["100"], timeMeasurements["Selection"]["100"])

	fmt.Printf("Quantity: 1000\n")
	fmt.Printf("%-8s %-8s %-8s\n", "Insertion", "Grouping", "Selection")
	fmt.Printf("%-8d %-8d %-8d\n", timeMeasurements["Insertion"]["1000"], timeMeasurements["Grouping"]["1000"], timeMeasurements["Selection"]["1000"])

	fmt.Printf("Quantity: 10000\n")
	fmt.Printf("%-8s %-8s %-8s\n", "Insertion", "Grouping", "Selection")
	fmt.Printf("%-8d %-8d %-8d\n", timeMeasurements["Insertion"]["10000"], timeMeasurements["Grouping"]["10000"], timeMeasurements["Selection"]["10000"])
}

func sortViaInsertion(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}

	return arr
}

/* Assuming this is a Merge Sort actually */
func sortViaGrouping(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left := sortViaGrouping(arr[:len(arr)/2])
	right := sortViaGrouping(arr[len(arr)/2:])

	return merge(left, right)
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		final = append(final, a[i])
	}

	for ; j < len(b); j++ {
		final = append(final, b[j])
	}

	return final
}

/* Assuming this is a Heapsort */
func sortViaSelection(arr []int) []int {

	buildMaxHeap(arr)
	size := len(arr)

	for size > 1 {
		arr[0], arr[size-1] = arr[size-1], arr[0]
		size--
		heapify(arr[:size], 0)
	}

	return arr
}

func buildMaxHeap(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapify(arr, i)
	}
}

func heapify(arr []int, i int) {
	largest := i
	left, right := getChildren(i)
	if (left < len(arr)) && (arr[left] > arr[largest]) {
		largest = left
	}
	if (right < len(arr)) && (arr[right] > arr[largest]) {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, largest)
	}

}

func getChildren(i int) (int, int) {
	return 2*i + 1, 2*i + 2
}
