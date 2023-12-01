package main

import (
	"errors"
	"fmt"
	"github.com/philfreshman/advent-of-code-2022/day07"
	"io"
	"sort"
	"time"

	"github.com/philfreshman/advent-of-code-2022/day01"
	"github.com/philfreshman/advent-of-code-2022/day02"
	"github.com/philfreshman/advent-of-code-2022/day03"
	"github.com/philfreshman/advent-of-code-2022/day04"
	"github.com/philfreshman/advent-of-code-2022/day05"
	"github.com/philfreshman/advent-of-code-2022/day06"

	"github.com/spf13/cobra"
)

const (
	benchNum = 10000
)

var (
	puzzles = [][]Runner{
		{day01.PuzzleA{}, day01.PuzzleB{}},
		{day02.PuzzleA{}, day02.PuzzleB{}},
		{day03.PuzzleA{}, day03.PuzzleB{}},
		{day04.PuzzleA{}, day04.PuzzleB{}},
		{day05.PuzzleA{}, day05.PuzzleB{}},
		{day06.PuzzleA{}, day06.PuzzleB{}},
		{day07.PuzzleA{}, day07.PuzzleB{}},
	}
)

type Runner interface {
	fmt.Stringer
	Run() any
}

type Options struct {
	Bench  bool
	Puzzle int
	Times  int
}

func newRootCmd(args []string, out io.Writer) (*cobra.Command, error) {
	opts := Options{}

	cmd := &cobra.Command{
		Use:          "aoc-2022",
		Short:        "Advent of Code 2022 Solutions",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			toRun := make([]Runner, 0, len(puzzles)*2)
			if opts.Puzzle > 0 {
				if opts.Puzzle > len(puzzles) {
					return errors.New("no puzzle found")
				}

				toRun = append(toRun, puzzles[opts.Puzzle-1]...)

			} else {
				for _, pz := range puzzles {
					toRun = append(toRun, pz...)
				}
			}

			if opts.Bench {
				benchmarkPuzzles(out, toRun, opts)
			} else {
				runPuzzles(out, toRun)
			}

			return nil
		},
	}
	f := cmd.Flags()
	f.BoolVarP(&opts.Bench, "benchmark", "b", false, "run in benchmarking mode")
	f.IntVarP(&opts.Times, "times", "t", benchNum, "number of puzzle executions during benchmark")
	f.IntVarP(&opts.Puzzle, "day", "p", 0, "run a specific puzzle only")

	return cmd, nil
}

func runPuzzles(out io.Writer, pzs []Runner) {
	fmt.Fprintln(out, "🎄 Advent of Code 2022")
	fmt.Fprintln(out)

	// Capture all durations for displaying total elapsed time
	elapsed := make([]time.Duration, len(pzs))

	for _, pz := range pzs {
		now := time.Now()
		res := pz.Run()
		dur := time.Since(now)
		fmt.Fprintf(out, "🧩 Puzzle %s: %-14d [time taken: %s]\n", pz, res, dur)

		elapsed = append(elapsed, dur)
	}

	fmt.Fprintln(out)
	fmt.Fprintf(out, "Total elapsed time: [%s]\n", sumElapsedTime(elapsed))
}

func sumElapsedTime(dur []time.Duration) time.Duration {
	var t time.Duration
	for _, d := range dur {
		t += d
	}

	return t
}

func benchmarkPuzzles(out io.Writer, pzs []Runner, opts Options) {
	fmt.Fprintf(out, "🎄 Advent of Code 2022 - Benchmark [executions: %d]\n", opts.Times)
	fmt.Fprintln(out)

	// Capture all durations for displaying total elapsed time
	elapsed := make([]time.Duration, len(pzs))

	exec := make([]time.Duration, opts.Times)
	for _, pz := range pzs {
		//var res any

		// Run each puzzle the required benchmark sample rate, collecting each duration
		for b := 0; b < opts.Times; b++ {
			now := time.Now()
			pz.Run()
			exec[b] = time.Since(now)
		}
		sort.Slice(exec, func(i, j int) bool {
			return exec[i] < exec[j]
		})

		fmt.Fprintf(out, "🧩 Puzzle %s:\t[time taken: %s]\n", pz, exec[0])
		elapsed = append(elapsed, exec[0])
	}

	fmt.Fprintln(out)
	fmt.Fprintf(out, "Total elapsed time: [%s]\n", sumElapsedTime(elapsed))
}
