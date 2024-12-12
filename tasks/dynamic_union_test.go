package tasks

import (
	"bufio"
	"fmt"
	"os"
	"sandbox/utils"
	"strconv"
	"strings"
	"testing"
)

type pq struct {
	p, q int
}

func dynamicSumFile(filename string) (int, []pq) {
	defer fmt.Println("finish read file:" + filename)
	file := utils.Must(os.Open(filename))
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	n := utils.Must(strconv.Atoi(scanner.Text()))
	result := make([]pq, 0, n)
	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), " ")
		p, q := utils.Must(strconv.Atoi(pair[0])), utils.Must(strconv.Atoi(pair[1]))
		result = append(result, pq{p, q})
	}
	return n, result
}

func TestDynamicUnion1(t *testing.T) {
	cases := []struct {
		name     string
		filename string
		want     int
	}{
		{
			name:     "tiny",
			filename: "./testdata/tinyUF.txt",
			want:     2,
		},
		{
			name:     "medium",
			filename: "./testdata/mediumUF.txt",
			want:     3,
		},
		// {
		// 	name:     "large",
		// 	filename: "./testdata/largeUF.txt", // > 30sec
		// 	want:     6,
		// },
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			n, result := dynamicSumFile(c.filename)
			var uf DynamicUnion = NewDynamicUnion1(n)
			for _, pq := range result {
				if uf.Connected(pq.p, pq.q) {
					continue
				}
				uf.Union(pq.p, pq.q)
			}
			if got := uf.Count(); got != c.want {
				t.Errorf("Count() = %v, want %v", got, c.want)
			}
		})
	}
}

func TestDynamicUnion2(t *testing.T) {
	cases := []struct {
		name     string
		filename string
		want     int
	}{
		{
			name:     "tiny",
			filename: "./testdata/tinyUF.txt",
			want:     2,
		},
		{
			name:     "medium",
			filename: "./testdata/mediumUF.txt",
			want:     3,
		},
		// {
		// 	name:     "large",
		// 	filename: "./testdata/largeUF.txt", // > 30
		// 	want:     6,
		// },
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			n, result := dynamicSumFile(c.filename)
			var uf DynamicUnion = NewDynamicUnion2(n)
			for _, pq := range result {
				if uf.Connected(pq.p, pq.q) {
					continue
				}
				uf.Union(pq.p, pq.q)
			}

			if got := uf.Count(); got != c.want {
				t.Errorf("Count() = %v, want %v", got, c.want)
			}
		})
	}
}

func TestDynamicUnion3(t *testing.T) {
	cases := []struct {
		name     string
		filename string
		want     int
	}{
		{
			name:     "tiny",
			filename: "./testdata/tinyUF.txt",
			want:     2,
		},
		{
			name:     "medium",
			filename: "./testdata/mediumUF.txt",
			want:     3,
		},
		{
			name:     "large",
			filename: "./testdata/largeUF.txt",
			want:     6,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			n, result := dynamicSumFile(c.filename)
			var uf DynamicUnion = NewDynamicUnion3(n)
			for _, pq := range result {
				if uf.Connected(pq.p, pq.q) {
					continue
				}
				uf.Union(pq.p, pq.q)
			}

			if got := uf.Count(); got != c.want {
				t.Errorf("Count() = %v, want %v", got, c.want)
			}
		})
	}
}

func TestDynamicUnion4(t *testing.T) {
	cases := []struct {
		name     string
		filename string
		want     int
	}{
		{
			name:     "tiny",
			filename: "./testdata/tinyUF.txt",
			want:     2,
		},
		{
			name:     "medium",
			filename: "./testdata/mediumUF.txt",
			want:     3,
		},
		{
			name:     "large",
			filename: "./testdata/largeUF.txt",
			want:     6,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			n, result := dynamicSumFile(c.filename)
			var uf DynamicUnion = NewDynamicUnion4(n)
			for _, pq := range result {
				if uf.Connected(pq.p, pq.q) {
					continue
				}
				uf.Union(pq.p, pq.q)
			}

			if got := uf.Count(); got != c.want {
				t.Errorf("Count() = %v, want %v", got, c.want)
			}
		})
	}
}
