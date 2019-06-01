package problem0991

func brokenCalc(X int, Y int) int {
	if X >= Y {
		return X - Y
	}

	if Y%2 == 0 {
		return brokenCalc(X, Y/2) + 1
	}

	if 2*X >= Y {
		return brokenCalc(X*2, Y) + 1
	}

	return brokenCalc(X, (Y+1)/2) + 2

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
