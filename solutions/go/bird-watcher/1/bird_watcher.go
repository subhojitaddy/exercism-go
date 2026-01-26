package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    birdDataLength := len(birdsPerDay)
    birdCount := 0
	for i := 0; i < birdDataLength; i++ {
        birdCount += birdsPerDay[i]
    }
    return birdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    birdInWeek := 0
    startIndex := (week - 1)*7
	for i := startIndex; i <= startIndex+6; i++ {
        birdInWeek += birdsPerDay[i]
    }
    return birdInWeek
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	birdDataLength := len(birdsPerDay)
    for i := 0; i < birdDataLength; i = i+2 {
        birdsPerDay[i]++
    }
    return birdsPerDay
}
