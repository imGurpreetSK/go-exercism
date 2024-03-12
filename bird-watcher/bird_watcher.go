package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var count int = 0
	for _, element := range birdsPerDay {
		count += element
	}
	return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	count := 0
	startDay := (week - 1) * 7
	endDay := startDay + 7
	for _, element := range birdsPerDay[startDay:endDay] {
		count += element
	}
	return count
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for index, element := range birdsPerDay {
		if index%2 == 0 {
			birdsPerDay[index] = element + 1
		}
	}

	return birdsPerDay
}
