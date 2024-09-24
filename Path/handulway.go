package graph

type Data struct {
	Col, Row int
	Realst   [][]string
	Index    int
}

func AllPathDisjoin(allPaths [][]string) map[int][][]string {
	res := make(map[int][][]string)
	indix := 0
	for _, path := range allPaths {
		passed := false
		if len(res) == 0 {
			res[indix] = append(res[indix], path)
		} else {
			for i, way := range res {
				if !HandulWay(way, path) {
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

	for _, Paths := range allPaths {
		for i, r := range res {
			if !HandulWay(r, Paths) {
				res[i] = append(res[i], Paths)
			}
		}
	}
	return res
}

func HandulWay(Paths [][]string, way []string) bool {
	for _, t := range Paths {
		if !isDisjoint(t, way) {
			return true
		}
	}
	return false
}
