package randoms

// You are given an array time where time[i] denotes the time taken by the ith bus to complete one trip.

// Each bus can make multiple trips successively; that is, 
// the next trip can start immediately after completing the current trip. 
// Also, each bus operates independently; that is, the trips of one bus do not influence the trips of any other bus.

// You are also given an integer totalTrips, 
// which denotes the number of trips all buses should make in total. 
// Return the minimum time required for all buses to complete at least totalTrips trips.

func MinimumTime(time []int, totalTrips int) int{
	low := 0
	high := time[0]
	for _, v := range time {
			if (v < high) {
				high = v
			}
	}
	high *= totalTrips

	f := func(mid int) bool {
		count := 0
		for _, v := range time {
			count += mid / v
		}
		return count >= totalTrips
	}

	for low <= high {
		mid := low + (high - low) / 2
		if (f(mid)) {
			high = mid -1
		} else {
			low = mid + 1
		}
	}
	return low
}