package util

// Contains -> SliceIncludes
// SliceIncludes checks if a string slice contains a specified string
func SliceIncludes(collection []string, searchTerm string) bool {  // slice -> collection, text -> searchTerm
	for _, element := range collection {  // item -> element
		if element == searchTerm {
			return true
		}
	}

	return false
}

// Difference -> SliceDifference
// SliceDifference returns the difference of two string slices
func SliceDifference(collection1 []string, collection2 []string) (difference []string) {  // slice -> collection1, slice2 -> collection2
	// Loop two times, first to find collection1 strings not in collection2,
	// second loop to find collection2 strings not in collection1
	for i := 0; i < 2; i++ {
		for _, element1 := range collection1 {  // s1 -> element1
			found := false
			for _, element2 := range collection2 {  // s2 -> element2
				if element1 == element2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				difference = append(difference, element1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			collection1, collection2 = collection2, collection1
		}
	}

	return difference
}

// Index -> SliceIndex
// SliceIndex returns the index of a string in a string slice
func SliceIndex(collection []string, searchTerm string) int {  // slice -> collection, text -> searchTerm
	for i, element := range collection {  // item -> element
		if element == searchTerm {
			return i
		}
	}

	return 0
}
