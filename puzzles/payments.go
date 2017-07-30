package puzzles

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

Solution:
===

The algorithm is trivial:

1) Scan the list of incoming payments having the target amount.
2) If the element is greater than the amount just skip it.
2) If the element is equals to the amount it means the chain of incoming payments has been found.
3) If element is less than the amount remember it, deduct it from the amount then recursively apply same
   steps to the list excluding the element and having new amount. Merge result upon return.

Issues:

- Complexity: O(n2). It can be improved using BTS: no need to continue with scan if (v > target) => O(n x n x logN)
- The algorithm does not return all possible chains. Introduce special structure (Go) or class (Java) at next steps:

struct Payment {
	id int
	amount int
}

and return all possible combinations i.e. ([1:100],[6:80]),([2:100],[6:80])


*/
func find(a []int, target int) []int {

	comprise := make([]int, 0)

	if len(a) == 0 || target == 0 {
		return comprise
	}

	for i, v := range a {

		if v > target {
			continue
		}

		if v == target {
			return append(comprise, v)
		}

		c := find(a[i+1:], target-v)

		if len(c) > 0 {
			comprise = append(comprise, v)
			for _, cv := range c {
				comprise = append(comprise, cv)
			}
			break
		}
	}

	return comprise
}
