package day1_test

import (
	"bufio"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xorilog/aoc2021/day1"
)

var example = []int{
	199,
	200,
	208,
	210,
	200,
	207,
	240,
	269,
	260,
	263,
}

func TestCountMeasurementsIncreases(t *testing.T) {
	increases := day1.CountMeasurementsIncreases(example)
	assert.Equal(t, 7, increases)
}

func TestCountMeasurementsIncreases_Answer_Part1(t *testing.T) {
	input, err := readInput(t)
	require.NoError(t, err)

	t.Log("[part1] Total of increases: ", day1.CountMeasurementsIncreases(input))
}

func TestCountMeasurementsIncreasesSlidingWindow(t *testing.T) {
	increases := day1.CountMeasurementsIncreasesSlidingWindow(example, 3)
	assert.Equal(t, 5, increases)
}

func TestCountMeasurementsIncreasesSlidingWindow_Answer(t *testing.T) {
	input, err := readInput(t)
	require.NoError(t, err)

	t.Log("[part2] Total of increases: ", day1.CountMeasurementsIncreasesSlidingWindow(input, 3))
}

func readInput(t *testing.T) ([]int, error){
	t.Helper()

	file, err := os.Open("./data/input.txt")
	require.NoError(t, err)
	defer file.Close()

	return parseInput(file)
}

func parseInput(f io.Reader) ([]int, error) {
	var (
		scanner = bufio.NewScanner(f)
		result []int
	)

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		result = append(result, value)
	}
	return result, scanner.Err()
}
