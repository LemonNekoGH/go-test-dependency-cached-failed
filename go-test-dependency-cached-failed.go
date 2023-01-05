package gotestdependencycachedfailed

var (
	num     = []int{0}
	aCached = []int{}
)

func A() []int {
	if len(aCached) != 0 {
		return aCached
	}

	aCached = []int{num[0]}
	return aCached
}

func B() [][]int {
	return [][]int{A()}
}
