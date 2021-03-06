package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDecodeOp(t *testing.T) {
	tt := []struct {
		input    int
		expected Instruction
	}{
		{2, Instruction{
			OpCode:         2,
			ParameterMode1: 0,
			ParameterMode2: 0,
			ParameterMode3: 0,
			ParameterMode4: 0,
		}},
		{1, Instruction{
			OpCode:         1,
			ParameterMode1: 0,
			ParameterMode2: 0,
			ParameterMode3: 0,
			ParameterMode4: 0,
		}},
		{1002, Instruction{
			OpCode:         2,
			ParameterMode1: 0,
			ParameterMode2: 1,
			ParameterMode3: 0,
			ParameterMode4: 0,
		}},
		{111102, Instruction{
			OpCode:         2,
			ParameterMode1: 1,
			ParameterMode2: 1,
			ParameterMode3: 1,
			ParameterMode4: 1,
		}},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("Decoding %v into %v", tc.input, tc.expected), func(t *testing.T) {
			result := decodeInstruction(tc.input)
			if tc.expected != result {
				t.Fatalf("expected %v; got %v", tc.expected, result)
			}
		})
	}
}

func TestIntcode(t *testing.T) {
	tt := []struct {
		inputs   []int
		outputs  []int
		code     []int
		expected []int
		err      bool
	}{
		{[]int{}, []int{}, []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, false},
		{[]int{}, []int{}, []int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}, false},
		{[]int{}, []int{}, []int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}, false},
		{[]int{}, []int{}, []int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}, false},
		{[]int{}, []int{}, []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, false},
		{[]int{}, []int{}, []int{1002, 4, 3, 4, 33}, []int{1002, 4, 3, 4, 99}, false},      //Immediate mode
		{[]int{}, []int{}, []int{1101, 100, -1, 4, 0}, []int{1101, 100, -1, 4, 99}, false}, //Immediate mode
		{[]int{99}, []int{}, []int{3, 2, 0}, []int{3, 2, 99}, false},                       //Input mode
		{[]int{99}, []int{99}, []int{3, 4, 4, 4, 0}, []int{3, 4, 4, 4, 99}, false},         //Input & Output mode
		{[]int{}, []int{}, []int{1, 0, 0, 0}, []int{}, true},                               //error - No end instruction
		{[]int{}, []int{}, []int{7, 0, 0, 0, 99}, []int{}, true},                           //error - invalid opcode
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v vs %v", tc.code, tc.expected), func(t *testing.T) {
			result, outputs, err := runIntCode(tc.inputs, tc.code)

			// if err != nil {
			// 	fmt.Printf("%v -> %v (%v)\n", tc.input, tc.expected, err)
			// }

			if tc.err && err == nil {
				t.Fatalf("expected an error but got none %v", err)
			} else if !tc.err {
				if err != nil {
					t.Fatalf("expected no error but got %v", err)
				}
				if !reflect.DeepEqual(tc.expected, result) {
					t.Fatalf("expected %v; got %v", tc.expected, result)
				}
				if !reflect.DeepEqual(tc.outputs, outputs) {
					t.Fatalf("expected output %v; got %v", tc.outputs, outputs)
				}
			}

		})
	}
}
