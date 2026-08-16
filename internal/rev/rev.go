package rev

// Reverse returns a reversed copy of s (must not mutate s).
func Reverse(s []int) []int {
	out := make([]int, len(s))
	for i := range s {
		out[len(s)-1-i] = s[i]
	}
	return out
}
