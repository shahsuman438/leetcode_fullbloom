package main

import (
	"fmt"

	"github.com/shahsuman438/leetcode_fullbloom/pkg/logic"
)

func main() {
	plantTime := []int{1, 4, 3}
	growTime := []int{2, 3, 1}

	fmt.Println(logic.EarliestFullBloom(plantTime, growTime))
}
