package amazon

func findSecondSmallestAndLargest(arr []int) (int, int) {
	secondMin := MinInt
	secondMax := MaxInt

	smallest := make([]int, 0)
	largest := make([]int, 0)

	for _, x := range arr {
		if len(smallest) == 0 {
			smallest = append(smallest, x)
		} else if len(smallest) == 1 {
			if smallest[0] < x {
				smallest = append(smallest, x)
			} else if smallest[0] > x {
				smallest = append(smallest, smallest[0])
				smallest[0] = x
			}
		} else {
			if x < smallest[0] {
				smallest[1] = smallest[0]
				smallest[0] = x
			} else if x < smallest[1] {
				smallest[1] = x
			}
		}
		if  len(largest) == 0 {
			largest = append(largest, x)
		} else if len(largest) == 1 {
			if largest[0] > x {
				largest = append(largest, x)
			} else {
				largest = append(largest, largest[0])
				largest[0] = x
			}
		} else {
			if x > largest[0] {
				largest[1] = largest[0]
				largest[0] = x
			} else if x > largest[1] {
				largest[1] = x
			}
		}
	}
	if len(smallest) > 0 {
		secondMin = smallest[len(smallest) - 1]
	}
	if len(largest) > 0 {
		secondMax = largest[len(largest) - 1]
	}
	return secondMin, secondMax
}

const UintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64

const (
	MaxInt  = 1<<(UintSize-1) - 1 // 1<<31 - 1 or 1<<63 - 1
	MinInt  = -MaxInt - 1         // -1 << 31 or -1 << 63
	MaxUint = 1<<UintSize - 1     // 1<<32 - 1 or 1<<64 - 1
)
