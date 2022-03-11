package main

import "testing"
import "reflect"

func TestSum(t *testing.T) {

	t.Run("collections of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6
	
		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}	
	})

}

func TestSumAll(t *testing.T) {

	t.Run("sum of multiple collections", func(t *testing.T) {
		numbers_1 := []int{1, 2, 3}
		numbers_2 := []int{1, 2, 3}

		got := SumAll(numbers_1, numbers_2)
		want := []int{6, 6}
	
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}	
	})

}

func TestSumAllTails(t *testing.T) {
	
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}	
	}


	t.Run("sum of multiple collections tails", func(t *testing.T) {
		numbers_1 := []int{1, 2, 3}
		numbers_2 := []int{1, 2, 3}

		got := SumAllTails(numbers_1, numbers_2)
		want := []int{5, 5}
	
		checkSums(t, got, want)
	})

	t.Run("sum of empty collections tails", func(t *testing.T) {
		numbers_1 := []int{}
		numbers_2 := []int{1}

		got := SumAllTails(numbers_1, numbers_2)
		want := []int{0, 0}

		checkSums(t, got, want)
	})

}