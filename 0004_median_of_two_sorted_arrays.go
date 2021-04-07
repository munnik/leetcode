package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lengthMerged, lengthNums1, lengthNums2 := len(nums1)+len(nums2), len(nums1), len(nums2)
	numsMerged := make([]int, lengthMerged)
	index1, index2, indexMerged := 0, 0, 0

	for indexMerged < lengthMerged {
		if index1 < lengthNums1 && (index2 == lengthNums2 || nums1[index1] < nums2[index2]) {
			numsMerged[indexMerged] = nums1[index1]
			index1 += 1
		} else {
			numsMerged[indexMerged] = nums2[index2]
			index2 += 1
		}
		indexMerged += 1
	}

	if lengthMerged%2 == 0 {
		return float64(numsMerged[lengthMerged/2-1]+numsMerged[lengthMerged/2]) / 2
	}
	return float64(numsMerged[lengthMerged/2])
}
