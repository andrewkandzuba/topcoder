package puzzles

import (
	"testing"
	//"github.com/stretchr/testify/assert"
	"fmt"
)

func TestFind(t *testing.T) {
	/**
Puzzle

At a payments processing company, optimizing incoming and outgoing payments is a type of fee optimization.

Below, we have a list of the incoming payments that we optimized as outgoing payments.

The outgoing payments are lump sums comprised of the some of the incoming payments.

Write an algorithm that will show for each outgoing amount which incoming amounts it is comprised of.

Identify any potential issues with your solution.

Incoming amounts:

100 100 225 300 473 80

Outgoing amounts:

180 773 225 100
	 */

	incoming := []int{100, 100, 225, 300, 473, 80}
	//outgoing := []int{180, 773, 225, 100}
	outgoing := []int{180}

	for _, amount := range outgoing {
		comprised := findComprise(incoming, amount)
		fmt.Println(comprised)
	}
}
