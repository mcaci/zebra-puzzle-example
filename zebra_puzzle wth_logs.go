// +build uselogs

package zebra

import (
	"log"
)

func recursiveBacktrackingWithLogs(assignment *[]uint8, csp *CSP) bool {
	log.Println(*assignment)
	printAssignment(csp.v)
	if complete(csp.v) {
		return true
	}
	idx := selectUnassigned(&csp.v, assignment, *csp)
	log.Println("-", "Selecting variable at idx", idx)
domainValueLoop:
	for _, d := range csp.d {
		log.Println("--", "Selecting value", catElem(idx, d))
		csp.v[idx] = d
		for _, c := range csp.c {
			log.Println("---", "Selecting constraint", c)
			switch {
			case !c.forApplicability()(csp.v...):
				log.Println("----", "Constraint not applicable")
				continue
			case !c.forSatisfaction()(csp.v...):
				log.Println("----", "Constraint not satisfied")
				log.Println("--", "Discarding value", catElem(idx, d))
				csp.v[idx] = 0
				continue domainValueLoop
			default:
				log.Println("----", "Constraint ok")
			}
		}
		log.Println("--", "Accepting value", catElem(idx, d), "at idx", idx)
		copy(*assignment, csp.v)
		outcome := recursiveBacktracking(assignment, csp)
		if !outcome {
			csp.v[idx] = 0
			log.Println("--", "Discarding value", catElem(idx, d), "at idx", idx)
			continue domainValueLoop
		}
		return true
	}
	return false
}

func printAssignment(assignment []uint8) {
	grid := make([][]domainer, 5)
	for i, v := range assignment {
		grid[i%5] = append(grid[i%5], catElem(uint8(i), v))
	}
	log.Println(grid)
}

func catElem(idx, val uint8) domainer {
	switch idx / 5 {
	case 0:
		return nationality(val)
	case 1:
		return color(val)
	case 2:
		return drink(val)
	case 3:
		return smoke(val)
	case 4:
		return pets(val)
	case 5:
		return number(val)
	default:
		return nil
	}
}
