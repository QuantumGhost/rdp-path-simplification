package rdp

// SimplifyPath accepts a list of points and epsilon as threshold, simplifies a path by dropping
// points that do not pass threshold values.
func SimplifyPath(points []Point, ep float64) []Point {
	if len(points) <= 2 {
		return points
	}

	l := Line{Start: points[0], End: points[len(points)-1]}

	idx, maxDist := seekMostDistantPoint(l, points)
	if maxDist >= ep {
		left := SimplifyPath(points[:idx+1], ep)
		right := SimplifyPath(points[idx:], ep)
		return append(left[:len(left)-1], right...)
	}

	// If the most distant point fails to pass the threshold test, then just return the two points
	return []Point{points[0], points[len(points)-1]}
}

func seekMostDistantPoint(l Line, points []Point) (idx int, maxDist float64) {
	// FIX FROM https://github.com/calvinfeng/rdp-path-simplification/pull/1/files
	// by @mholt.
	//
	// TODO(QuantumGhost): add a test case for line causing infinite recursion.
	for i := 0; i < len(points)-1; i++ {
		d := l.DistanceToPoint(points[i])
		if d > maxDist {
			maxDist = d
			idx = i
		}
	}

	return idx, maxDist
}
