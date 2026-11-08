package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{
			name: "positive size",
			size: 100,
		},
		{
			name: "zero size",
			size: 0,
		},
		{
			name: "negative size",
			size: -5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)
			expectedLength := tt.size
			if expectedLength < 0 {
				expectedLength = 0
			}

			if len(data) != expectedLength {
				t.Errorf(
					"Expected length %d, got %d",
					expectedLength,
					len(data),
				)
			}
			for i, value := range data {
				if value <= 0 {
					t.Errorf(
						"Expected positive number at index %d, got %d",
						i,
						value,
					)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "empty slice",
			data: []int{},
			want: 0,
		},
		{
			name: "one element",
			data: []int{7},
			want: 7,
		},
		{
			name: "positive numbers",
			data: []int{2, 9, 4, 7},
			want: 9,
		},
		{
			name: "negative numbers",
			data: []int{-8, -2, -10},
			want: -2,
		},
		{
			name: "all equal",
			data: []int{5, 5, 5},
			want: 5,
		},
		{
			name: "maximum at the beginning",
			data: []int{100, 2, 3, 4},
			want: 100,
		},
		{
			name: "maximum at the end",
			data: []int{1, 2, 3, 100},
			want: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)

			if got != tt.want {
				t.Errorf(
					"maximum() = %d, want %d",
					got,
					tt.want,
				)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "empty slice",
			data: []int{},
			want: 0,
		},
		{
			name: "one element",
			data: []int{42},
			want: 42,
		},
		{
			name: "fewer elements than chunks",
			data: []int{3, 8, 1},
			want: 8,
		},
		{
			name: "exactly eight elements",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8},
			want: 8,
		},
		{
			name: "not divisible by eight",
			data: []int{1, 2, 3, 4, 5, 6, 7, 99, 9, 10},
			want: 99,
		},
		{
			name: "maximum at the beginning",
			data: []int{100, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 100,
		},
		{
			name: "maximum at the end",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 100},
			want: 100,
		},
		{
			name: "negative numbers",
			data: []int{-8, -2, -10, -4},
			want: -2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)

			if got != tt.want {
				t.Errorf(
					"maxChunks() = %d, want %d",
					got,
					tt.want,
				)
			}
		})
	}
}

func TestMaximumAndMaxChunksReturnSameResult(t *testing.T) {
	data := generateRandomElements(10003)

	oneThreadResult := maximum(data)
	multiThreadResult := maxChunks(data)

	if oneThreadResult != multiThreadResult {
		t.Errorf(
			"Results differ: maximum() = %d, maxChunks() = %d",
			oneThreadResult,
			multiThreadResult,
		)
	}
}
