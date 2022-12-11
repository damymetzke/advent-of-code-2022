package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type OperationType = byte

const (
	Multiply OperationType = iota
	Add
	Squared
)

type Monkey struct {
	items             []int
	operationType     OperationType
	operationAmount   int
	testDivisibleBy   int
	resultTrueMonkey  int
	resultFalseMonkey int
	inspectedNum      int
}

func GetInput() string {
	data, err := os.ReadFile("input/11")
	if err != nil {
		log.Fatal("Could not read file 'input/11':\n  * ", err)
	}

	return string(data)
}

func ParseMonkey(lines []string, i int) Monkey {
	// Parse items
	itemStrings := strings.Split(lines[i+1][18:], ", ")
	items := make([]int, len(itemStrings))

	for i, item := range itemStrings {
		value, err := strconv.ParseInt(item, 10, 64)

		if err != nil {
			log.Fatalf("Could not parse items '%v'\n  line: '%v'", item, lines[i+1])
		}

		items[i] = int(value)
	}

	// Parse operation type
	operationTypeString := lines[i+2][23]
	var operationType OperationType

	switch operationTypeString {
	case '*':
		operationType = Multiply
	case '+':
		operationType = Add
	default:
		log.Fatalf("Could not parse operation type '%v'\n  line: '%v'", operationTypeString, lines[i+2])
	}

	// Parse operation amount
	operationAmountString := lines[i+2][25:]
	var operationAmount int

	if operationAmountString == "old" {
		operationType = Squared
		operationAmount = 2
	} else {
		result, err := strconv.ParseInt(operationAmountString, 10, 64)

		if err != nil {
			log.Fatalf("Could not parse operation amount '%v'\n  line: '%v'", operationTypeString, lines[i+2])
		}

		operationAmount = int(result)
	}

	// Parse test divisible by
	testDivisibleBy, err := strconv.ParseInt(lines[i+3][21:], 10, 64)

	if err != nil {
		log.Fatalf("Could not parse test divisible by\n  line: '%v'", lines[i+3])
	}

	// Parse result true monkey
	resultTrueMonkey, err := strconv.ParseInt(lines[i+4][29:], 10, 64)

	if err != nil {
		log.Fatalf("Could not parse result true monkey\n  line: '%v'", lines[i+4])
	}

	// Parse result false monkey
	resultFalseMonkey, err := strconv.ParseInt(lines[i+5][30:], 10, 64)

	if err != nil {
		log.Fatalf("Could not parse result false monkey\n  line: '%v'", lines[i+5])
	}

	return Monkey{
		items:             items,
		operationType:     operationType,
		operationAmount:   operationAmount,
		testDivisibleBy:   int(testDivisibleBy),
		resultTrueMonkey:  int(resultTrueMonkey),
		resultFalseMonkey: int(resultFalseMonkey),
		inspectedNum:      0,
	}
}

func main() {
	lines := strings.Split(GetInput(), "\n")
	numMonkeys := (len(lines) + 1) / 7
	monkeys := make([]Monkey, numMonkeys)

	var wait sync.WaitGroup

	for i := 0; i < numMonkeys; i++ {
		monkeyI := i
		textOffset := i * 7
		wait.Add(1)

		go func() {
			monkeys[monkeyI] = ParseMonkey(lines, textOffset)
			wait.Done()
		}()
	}

	wait.Wait()

	// Go for 20 rounds
	for i := 0; i < 20; i++ {
		// Each monkey gets a round
		for j, monkey := range monkeys {
			// Consider each item
			for _, item := range monkey.items {
				// Transform worry
				switch monkey.operationType {
				case Multiply:
					item *= monkey.operationAmount
				case Add:
					item += monkey.operationAmount
				case Squared:
					item = item * item
				default:
					log.Fatalf("Invalid state:\n%v", monkey)
				}

				// Monkey is bored
				item /= 3

				if item%monkey.testDivisibleBy == 0 {
					monkeys[monkey.resultTrueMonkey].items = append(monkeys[monkey.resultTrueMonkey].items, item)
				} else {
					monkeys[monkey.resultFalseMonkey].items = append(monkeys[monkey.resultFalseMonkey].items, item)
				}

				// Increment inspectednum
				monkey.inspectedNum++
			}

			monkey.items = []int{}
			// I initially forgot to do this.
			// This duplicated all items.
			// And exponental growth is very real.
			// This caused my memory to overflow and my desktop to crash.
			// :)
			monkeys[j] = monkey
		}
	}

	maxInspected := [2]int{0, 0}

	for _, monkey := range monkeys {
		max := monkey.inspectedNum
    if max > maxInspected[0] {
      maxInspected[0], max = max, maxInspected[0]
    }

    if max > maxInspected[1] {
      maxInspected[1] = max
    }
	}

  monkeyBusiness := maxInspected[0] * maxInspected[1]

	fmt.Println("=-= PART 1 =-=")
	fmt.Println(monkeyBusiness)
	fmt.Println("=-= PART 2 =-=")
}
