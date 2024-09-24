package graph

type Data struct {
	Col, Row int
	Realst   [][]string
}

func AllPathDisjoin(allPaths [][]string) map[int][][]string {
	res := make(map[int][][]string)
	indix := 0
	for _, path := range allPaths {
		passed := false
		if len(res) == 0 {
			res[indix] = append(res[indix], path)
		} else {
			for i, roes := range res {
				if !OK(roes, path) {
					res[i] = append(res[i], path)
					passed = true
				}
			}
			if !passed {
				indix++
				res[indix] = append(res[indix], path)
			}
		}
	}

	for i, t := range res {
		for _, r := range allPaths {
			if !OK(t, r) {
				res[i] = append(res[i], r)
			}
		}
	}
	return res
}

func OK(p [][]string, f []string) bool {
	for _, t := range p {
		if !isDisjoint(t, f) {
			return true
		}
	}
	return false
}
