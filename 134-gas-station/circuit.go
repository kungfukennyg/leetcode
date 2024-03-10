package circuit

func canCompleteCircuit(gas, costs []int) int {
	required, cost := 0, 0

	for i := 0; i < len(gas); i++ {
		required += gas[i]
		cost += costs[i]
	}

	if required < cost {
		return -1
	}

	tank := 0
	start := 0
	for i := 0; i < len(gas); i++ {
		tank += gas[i] - costs[i]
		if tank < 0 {
			tank = 0
			start = i + 1
		}
	}

	if start == len(gas) {
		return -1
	}

	return start
}
