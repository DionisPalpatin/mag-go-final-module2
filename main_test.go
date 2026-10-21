package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FindMaximumSuite struct {
	suite.Suite

	reallyBigArray []int
}

func TestFindMaximumSuite(t *testing.T) {
	suite.Run(t, new(FindMaximumSuite))
}

func (s *FindMaximumSuite) SetupSuite() {
	s.reallyBigArray = generateRandomElements(SIZE)
}

func (s *FindMaximumSuite) TestFindMaximumOneThread() {
	tests := []struct {
		name    string
		input   []int
		wantMax int
	}{
		{
			"Пустой массив",
			[]int{},
			-1,
		},
		{
			"Массив с одним элементом",
			[]int{1},
			1,
		},
		{
			"Значения идут по возрастанию",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			10,
		},
		{
			"Значения идут по убыванию",
			[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			10,
		},
		{
			"Значения идут вразнобой",
			[]int{5, 2, 1, 10, 6, 8, 9, 7, 3, 4},
			10,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			max := maximum(tt.input)

			assert.Equal(s.T(), tt.wantMax, max)
		})
	}
}

func (s *FindMaximumSuite) TestFindMaximumSomeThreads() {
	tests := []struct {
		name    string
		input   []int
		wantMax int
	}{
		{
			"Пустой массив",
			[]int{},
			-1,
		},
		{
			"Массив с количеством элементов меньше количества горутин",
			[]int{1},
			1,
		},
		{
			"Массив с количеством элементов, кратным количеству горутин",
			[]int{12, 45, 7, 89, 23, 56, 34, 18, 67, 91, 42, 15, 73, 28, 54, 36},
			91,
		},
		{
			"Массив с количеством элементов, не кратным количеству горутин",
			[]int{12, 45, 7, 89, 23, 56, 34, 18, 67, 91, 42, 15, 73, 28, 54, 36, 99, 21, 63},
			99,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			max := maximum(tt.input)

			assert.Equal(s.T(), tt.wantMax, max)
		})
	}
}
