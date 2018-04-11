package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	i := 0
	for scanner.Scan() {
		i++
		if i == 1 {
			continue
		}

		saveTheWorld(scanner.Text(), i-1)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func saveTheWorld(testCase string, caseNum int) {

	args := strings.Split(testCase, " ")
	program := args[1]

	if shield, err := strconv.Atoi(args[0]); err == nil {

		swaps := numOfSwaps(shield, program)
		if swaps != -1 {
			fmt.Fprintf(os.Stdout, "Case #%d: %d\n", caseNum, swaps)
		} else {
			fmt.Fprintf(os.Stdout, "Case #%d: IMPOSSIBLE\n", caseNum)
		}

	}
}

func numOfSwaps(d int, p string) int {
	if calcDamage(p) <= d {
		return 0
	}

	ideal := idealProgram(p)
	if calcDamage(ideal) > d {
		return -1
	}

	swaps := 0
	for p != ideal {
		for i := len(p) - 1; i >= 0; i-- {
			if p[i] == 'S' {
				if i != 0 && p[i-1] == 'C' {
					p = swap(p, i)
					swaps++
					if calcDamage(p) <= d {
						return swaps
					}
				}
			}
		}

	}
	return 0

}

//calcDamage calculates the damage done with the inputted program
func calcDamage(p string) int {
	currentDamage := 1
	doneDamage := 0
	for _, char := range p {
		if char == 'C' {
			currentDamage *= 2
		}
		if char == 'S' {
			doneDamage += currentDamage
		}
	}

	return doneDamage
}

func swap(p string, pos int) string {
	prefix := p[:pos-1]
	mid := "SC"
	suffix := p[pos+1:]

	return (prefix + mid + suffix)
}

func idealProgram(p string) string {
	ideal := ""
	for _, char := range p {
		if char == 'S' {
			ideal += "S"
		}
	}

	for len(ideal) != len(p) {
		ideal += "C"
	}

	return ideal
}
