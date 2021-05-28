package amazon

func isRobotBounded(instructions string) bool {
	var direction = North
	var x = 0
	var y = 0

	direction, y, x = runInstructions(instructions, direction, y, x)
	return (x == 0 && y == 0) || (direction != North)
}

func runInstructions(instructions string, direction Direction, y int, x int) (Direction, int, int) {
	for _, command := range instructions {
		switch string(command) {
		case "G":
			switch direction {
			case North:
				y++
			case East:
				x++
			case South:
				y--
			case West:
				x--
			}
		case "L":
			switch direction {
			case North:
				direction = West
			default:
				direction--
			}
		case "R":
			switch direction {
			case West:
				direction = North
			default:
				direction++
			}
		}
	}
	return direction, y, x
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)