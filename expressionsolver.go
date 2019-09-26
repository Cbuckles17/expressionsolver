package main

import (
	"fmt"
	"math"
)

// Letter struct stores the base and the power of a given letter in
// an expression.
type Letter struct {
	base int
	power int	
}

// primeFactorize takes in a value and prime factorizes it.
func primeFactorize(value int) (map[int]int, error) {
	if value == 1 {
	// cannot prime factorize a value of 1
		return nil, fmt.Errorf("primeFactorize: Cannot Prime Factorize a Value of 1")
	}

	// hardcoded prime numbers...these could by done on the fly but doing this for simplicity
	primes := [...]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113}
	index := 0
	pFactorMap := make(map[int]int)

	// while index is not OB
	// AND value is gt 0
	for index < len(primes) && value > 0 {
		if value % primes[index] == 0 {
		// if value is divisible by the current prime
			value/=primes[index]
			pFactorMap[primes[index]]++
			//fmt.Println(strconv.Itoa(primes[index]) + ": " + strconv.Itoa(value))
		} else {
			index++
		}
	}

	if value != 1 {
	// final value was not divisible by a prime number
		return nil, fmt.Errorf("primeFactorize: Could not Prime Factorize Given Value")
	} else {
		printPFM(pFactorMap)
		return pFactorMap, nil
	}
}

// matchCombined goes through the prime factor map and tries to match itself to the
// combined expression. This is O(mn) or really just O(n^2). There is no way to do 
// a 0(1) lookup on the expressions to grab power/exponent. This is because there 
// could be an expression that had two or more letters with the same exponent/power.
func matchCombined(pFactorMap map[int]int, combinedExp map[rune]*Letter) error {
	fmt.Println("\nAttempting to Match Combined Expression with PFM")
	printExp(combinedExp, false, false)
	
	// loop through every prime in the prime factor map
	for prime, primePower := range pFactorMap {
		foundMatch := false

		// loop through every letter of the combined expression
        for letterKey, letterValue := range combinedExp {
        	if primePower == letterValue.power && letterValue.base == -1 {
        	// if the powers/exponents match
        	// AND the letter has not been assigned a value already
        		combinedExp[letterKey].base = prime
        		foundMatch = true
        		//fmt.Println(string(letterKey) + ": " + strconv.Itoa(combinedExp[letterKey].base) + " exp: " + strconv.Itoa(letterValue.power))
        		break
        	}
        }

        if foundMatch == false {
        // there wasn't a match
        	return fmt.Errorf("matchCombined: Failed to Find a Match for %d^%d", prime, primePower)
        }
	}

	printExp(combinedExp, true, false)

	return nil
}

// checkExps goes through all the expressions and checks to make sure
// that with the solved letters the expression sums match. This is O(mn).
// These is no way to do this check faster as you need to loop through all of the
// functions and check each one for correctness.
func checkExps(exps []map[rune]*Letter) error {
	// store the combined expression because it holds the letter values at the moment
	// and we will be calling it quite a bit
	combinedExp := exps[0]

	// loop through the expressions
	for expKey, expValue := range exps {
		// start the sum at 1 rather than 0 because the expressions are multiplication 
		var sum float64 = 1

		fmt.Println("\nChecking Expression for Correct Sum")
		printExp(expValue, false, true)

		// loop through every letter of the expression
		for letterKey, letterValue := range expValue {
			if combinedExp[letterKey].base == -1 {
			// a letter was unmatched in the combined expression
				return fmt.Errorf("checkExps: Unmatched letter %s", letterKey)
			} else if letterKey != '-' {
			// - means it is the sum so ignore it
				sum*=math.Pow(float64(combinedExp[letterKey].base), float64(letterValue.power))
				// set the letter value
				exps[expKey][letterKey].base = combinedExp[letterKey].base
			}			
		}

		printExp(exps[expKey], true, true)

		if sum != float64(expValue['-'].base) {
		// original sum and new sum do not match
			return fmt.Errorf("checkExps: Sums Do Not Match %f != %f", sum, float64(expValue['-'].base))
		} 
	}

	return nil
}

// printExp pretty prints an expression. It has 2 bool flags to select 
// whether or not to show the solved expression and whether or not to
// show the sum of the expression. 
func printExp(exp map[rune]*Letter, solved bool, showSum bool) {
	firstKey := true

	fmt.Print("EXP: ")

	// loop through every letter of the expression
	for letterKey, letterValue := range exp {
		if letterKey == '-' {
		// - means it is the sum so ignore it
			continue
		} else if firstKey {
		// this is the first key in the loop
			firstKey = false

			if solved {
			// show solved letter
				fmt.Printf("%d^%d", letterValue.base, letterValue.power)
			} else {
				fmt.Printf("%c^%d", letterKey, letterValue.power)
			}
		} else {
			if solved {
			// show solved letter
				fmt.Printf(" + %d^%d", letterValue.base, letterValue.power)
			} else {
				fmt.Printf(" + %c^%d", letterKey, letterValue.power)
			}
		}	
	}

	if showSum {
	// show sum
		fmt.Printf(" = %d\n", exp['-'].base)
	} else {
		fmt.Printf("\n")
	}
}

// printPFM pretty prints a prime factor map.
func printPFM(pFactorMap map[int]int) {
	firstKey := true

	fmt.Print("PFM: ")

	// loop through every prime in the prime factor map
	for prime, primePower := range pFactorMap {
		if firstKey {
		// this is the first key in the loop
			firstKey = true
			fmt.Printf("%d^%d", prime, primePower)
		} else {
			fmt.Printf(" + %d^%d", prime, primePower)
		}
	}

	fmt.Printf("\n")
}

// printSolvedLetters prints out the solved letters in a list.
func printSolvedLetters(exp map[rune]*Letter) {
	fmt.Println("\nSolved Letters:")

	// loop through every letter of the expression
	for letterKey, letterValue := range exp {
		if letterKey == '-' {
		// - means it is the sum so ignore it
			continue
		}

		fmt.Printf("%c: %d\n", letterKey, letterValue.base)
	}
}

// main pulls in the choosen test case and then attempts to solve
// the expressions.
func main() {
	
	exps := test_case1()

	pFactorMap, err := primeFactorize(exps[0]['-'].base)
	if err != nil {
    // checking if primeFactorize returned an error
    	fmt.Println(err)
    	return
    }

    if err := matchCombined(pFactorMap, exps[0]); err != nil {
    // checking if matchCombined returned an error
    	fmt.Println(err)
    	return
    }

	if err := checkExps(exps); err != nil {
    // checking if checkExps returned an error
    	fmt.Println(err)
    	return
    }

    printSolvedLetters(exps[0])
}

// test cases

// good scenario
func test_case1() []map[rune]*Letter {
	// a^2 × b × c^2 × g = 5100
	exp1 := make(map[rune]*Letter)
	exp1['a'] = &Letter{-1, 2}
	exp1['b'] = &Letter{-1, 1}
	exp1['c'] = &Letter{-1, 2}
	exp1['g'] = &Letter{-1, 1}
	exp1['-'] = &Letter{5100, 1}
	// a × b^2 × e × f^2 = 33462
	exp2 := make(map[rune]*Letter)
	exp2['a'] = &Letter{-1, 1}
	exp2['b'] = &Letter{-1, 2}
	exp2['e'] = &Letter{-1, 1}
	exp2['f'] = &Letter{-1, 2}
	exp2['-'] = &Letter{33462, 1}
	// a × c^2 × d^3 = 17150
	exp3 := make(map[rune]*Letter)
	exp3['a'] = &Letter{-1, 1}
	exp3['c'] = &Letter{-1, 2}
	exp3['d'] = &Letter{-1, 3}
	exp3['-'] = &Letter{17150, 1}
	// a^3 × b^3 × c × d × e^2 = 914760
	exp4 := make(map[rune]*Letter)
	exp4['a'] = &Letter{-1, 3}
	exp4['b'] = &Letter{-1, 3}
	exp4['c'] = &Letter{-1, 1}
	exp4['d'] = &Letter{-1, 1}
	exp4['e'] = &Letter{-1, 2}
	exp4['-'] = &Letter{914760, 1}
	// a^7 × b^6 × c^5 × d^4 × e^3 × f^2 × g = 2677277333530800000
	combinedExp := make(map[rune]*Letter)
	combinedExp['a'] = &Letter{-1, 7}
	combinedExp['b'] = &Letter{-1, 6}
	combinedExp['c'] = &Letter{-1, 5}
	combinedExp['d'] = &Letter{-1, 4}
	combinedExp['e'] = &Letter{-1, 3}
	combinedExp['f'] = &Letter{-1, 2}
	combinedExp['g'] = &Letter{-1, 1}
	combinedExp['-'] = &Letter{2677277333530800000, 1}

	exps := make([]map[rune]*Letter, 0)
	exps = append(exps, combinedExp)
	exps = append(exps, exp1)
	exps = append(exps, exp2)
	exps = append(exps, exp3)
	exps = append(exps, exp4)

	return exps
}

// checks for failure in matching
func test_case2() []map[rune]*Letter {
	// a = 10
	exp1 := make(map[rune]*Letter)
	exp1['a'] = &Letter{-1, 1}
	exp1['-'] = &Letter{10, 1}
	// b^2 = 25
	exp2 := make(map[rune]*Letter)
	exp2['b'] = &Letter{-1, 2}
	exp2['-'] = &Letter{25, 1}
	// c^3 = 27
	exp3 := make(map[rune]*Letter)
	exp3['c'] = &Letter{-1, 3}
	exp3['-'] = &Letter{27, 1}
	// a x b^2 x c^3 = 6750
	combinedExp := make(map[rune]*Letter)
	combinedExp['a'] = &Letter{-1, 1}
	combinedExp['b'] = &Letter{-1, 2}
	combinedExp['c'] = &Letter{-1, 3}
	combinedExp['-'] = &Letter{6750, 0}

	exps := make([]map[rune]*Letter, 0)
	exps = append(exps, combinedExp)
	exps = append(exps, exp1)
	exps = append(exps, exp2)
	exps = append(exps, exp3)

	return exps
}

// matches the combined expression but not the sub expressions
func test_case3() []map[rune]*Letter {
	// a = 10
	exp1 := make(map[rune]*Letter)
	exp1['a'] = &Letter{-1, 1}
	exp1['-'] = &Letter{10, 1}
	// b^2 = 9
	exp2 := make(map[rune]*Letter)
	exp2['b'] = &Letter{-1, 2}
	exp2['-'] = &Letter{9, 1}
	// c^3 = 25
	exp3 := make(map[rune]*Letter)
	exp3['c'] = &Letter{-1, 3}
	exp3['-'] = &Letter{25, 1}
	// a x b^2 x c^3 = 2250
	combinedExp := make(map[rune]*Letter)
	combinedExp['a'] = &Letter{-1, 1}
	combinedExp['b'] = &Letter{-1, 2}
	combinedExp['c'] = &Letter{-1, 3}
	combinedExp['-'] = &Letter{2250, 1}

	exps := make([]map[rune]*Letter, 0)
	exps = append(exps, combinedExp)
	exps = append(exps, exp1)
	exps = append(exps, exp2)
	exps = append(exps, exp3)

	return exps
}

// checks for unable to factorize
func test_case4() []map[rune]*Letter {
	// a = 2253
	exp1 := make(map[rune]*Letter)
	exp1['a'] = &Letter{-1, 1}
	exp1['-'] = &Letter{2253, 1}
	// a = 2253
	combinedExp := make(map[rune]*Letter)
	combinedExp['a'] = &Letter{-1, 1}
	combinedExp['-'] = &Letter{2253, 1}

	exps := make([]map[rune]*Letter, 0)
	exps = append(exps, combinedExp)
	exps = append(exps, exp1)

	return exps
}

// checks for value of 1
func test_case5() []map[rune]*Letter {
	// a = 1
	exp1 := make(map[rune]*Letter)
	exp1['a'] = &Letter{-1, 1}
	exp1['-'] = &Letter{1, 1}
	// a = 1
	combinedExp := make(map[rune]*Letter)
	combinedExp['a'] = &Letter{-1, 1}
	combinedExp['-'] = &Letter{1, 1}

	exps := make([]map[rune]*Letter, 0)
	exps = append(exps, combinedExp)
	exps = append(exps, exp1)

	return exps
}
