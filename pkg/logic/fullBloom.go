package logic

import "sort"

func EarliestFullBloom(plantTime []int, growTime []int) int {
	n := len(plantTime)
	plantOrder := make([]int, n)

	// Initialize plantOrder with indices
	for i := range plantOrder {
		plantOrder[i] = i
	}

	// Sort by growTime in descending order, since we want to plant seeds that take the longest
	// to grow first, maximizing overlap between planting and growing times.
	sort.Slice(plantOrder, func(i, j int) bool {
		return growTime[plantOrder[i]] > growTime[plantOrder[j]]
	})

	currentPlantingTime := 0
	earliestBloomTime := 0

	// Calculate the earliest full bloom time
	for _, i := range plantOrder {
		// Add the current plant's planting time
		currentPlantingTime += plantTime[i] 
		// Calculate when this plant will fully bloom
		bloomTime := currentPlantingTime + growTime[i] 
		if bloomTime > earliestBloomTime {
			// Keep track of the latest bloom time
			earliestBloomTime = bloomTime 
		}
	}

	return earliestBloomTime
}

