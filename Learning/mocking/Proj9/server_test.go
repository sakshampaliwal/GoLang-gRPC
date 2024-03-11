package main

import (
	"testing"
	"github.com/stretchr/testify/mock"
)

type CalculatorMock struct {
    mock.Mock
}

func (m *CalculatorMock) Add(a, b int) int {
    args := m.Called(a, b)
    return args.Int(0)
}

func TestAdd(t *testing.T) {
    mockCalc := &CalculatorMock{}
    mockCalc.On("Add", 1, 2).Return(3)

    result := mockCalc.Add(1, 2)
    assert.Equal(t, 3, result)

    mockCalc.AssertExpectations(t)
}


// func TestAdd(t *testing.T) {
// 	type args struct {
// 		a int
// 		b int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Add(tt.args.a, tt.args.b); got != tt.want {
// 				t.Errorf("Add() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestSubtract(t *testing.T) {
// 	type args struct {
// 		a int
// 		b int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Subtract(tt.args.a, tt.args.b); got != tt.want {
// 				t.Errorf("Subtract() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
