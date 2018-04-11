package main

import "testing"

func TestSwap(t *testing.T) {
	p := "CS"
	swapped := swap(p, 1)
	if swapped != "SC" {
		t.Error("Didnt swap SC")
	}

	p = "SSCS"
	swapped = swap(p, 3)
	if swapped != "SSSC" {
		t.Error("Didnt swap SSSC")
	}

	p = "CSCS"
	swapped = swap(p, 3)
	if swapped != "CSSC" {
		t.Error("Didnt swap CSSC")
	}

}

func TestCalcDamage(t *testing.T) {

	input := "SCS"
	output := calcDamage(input)
	if output != 3 {
		t.Error("Didnt calculate damage 3")
	}

	input = "CCCSS"
	output = calcDamage(input)

	if output != 16 {
		t.Error("Didnt calculate damage 16")
	}
}

func TestnumOfSwaps(t *testing.T) {
	shield := 5
	program := "SC"

	if numOfSwaps(shield, program) != 0 {
		t.Error()
	}

	shield = 3
	program = "CCS"
	if numOfSwaps(shield, program) != 1 {
		t.Error()
	}

	shield = 83
	program = "CCSCCSCCS" //4 + 16 + 64  CSCCCSCCS 2 + 16 + 64
	if numOfSwaps(shield, program) != 1 {
		t.Error()
	}

	shield = 75
	program = "CCSCCSCCS" //4 + 16 + 64  CSCCSCCCS 2 + 8 + 64
	if numOfSwaps(shield, program) != 2 {
		t.Error("Shield 75")
	}

	shield = 43
	program = "CCSCCSCCS" //4 + 16 + 64  CSCCSCCSC 2 + 8 + 32
	if numOfSwaps(shield, program) != 3 {
		t.Error("Shield 43")
	}

	shield = 20
	program = "CCSCCSCCS" //4 + 16 + 64  CSCCSCCSC 2 + 8 + 32
	if numOfSwaps(shield, program) != -1 {
		t.Error("Shield 20")
	}
}

func TestIdealProgram(t *testing.T) {
	program := "CSSC"
	if idealProgram(program) != "SSCC" {
		t.Error("Didnt create ideal program")
	}
}
