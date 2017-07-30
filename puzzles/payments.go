package puzzles

func calc(incoming []int, target int) []int {
	return find(0, incoming, make([]int, 0), target)
}

func find(index int, incoming []int, comprised []int, target int) []int {
	if index < 0 || index == len(incoming) {
		if target == 0 {
			return comprised
		} else {
			return comprised[0:0]
		}
	}

	c, t := NewTarget(comprised, target, incoming[index]);

	lf := find(index+1, incoming, c, t)
	rf := find(index-1, incoming, c, t)

	if len(lf) > 0 {
		return lf
	}
	return rf
}

func NewTarget(comprised []int, target int, amount int) ([]int, int){
	deduct := target - amount

	if deduct >= 0 {
		return append(comprised, amount), deduct
	}

	return comprised, target;
}
