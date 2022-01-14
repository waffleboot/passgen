package main

import (
	"bufio"
	"bytes"
	"context"
	_ "embed"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

//go:embed data/nouns.txt
var _nouns []byte

//go:embed data/verbs.txt
var _verbs []byte

//go:embed data/adjectives.txt
var _adjectives []byte

func main() {
	if err := app(); !errors.Is(err, context.Canceled) {
		log.Fatal(err)
	}
}

func app() error {
	verbs, err := parse(_verbs)
	if err != nil {
		return err
	}

	adjectives, err := parse(_adjectives)
	if err != nil {
		return err
	}

	nouns, err := parse(_nouns)
	if err != nil {
		return err
	}

	rand.Seed(time.Now().UnixNano())

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		for {
			i := rand.Intn(len(verbs))
			j := rand.Intn(len(adjectives))
			k := rand.Intn(len(nouns))

			_ = verbs[i]
			a := adjectives[j]
			n := nouns[k]

			select {
			case <-time.After(500 * time.Millisecond):
			case <-gCtx.Done():
				return gCtx.Err()
			}

			s := a + strings.Title(n)
			fmt.Printf("%s %s -> %s\n", a, n, translate(s))

			// s := v + strings.Title(a) + strings.Title(n)
			// fmt.Printf("%s %s %s -> %s\n", v, a, n, translate(s))

			// s := v + translateAlways(strings.Title(a)) + strings.Title(n)
			// fmt.Printf("%s %s %s -> %s\n", v, a, n, s)
		}
	})

	return g.Wait()
}

func parse(b []byte) ([]string, error) {
	var ans []string

	s := bufio.NewScanner(bytes.NewReader(b))
	for s.Scan() {
		ans = append(ans, s.Text())
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return ans, nil
}

func translateAlways(s string) string {
	var buf strings.Builder
	for _, c := range s {
		buf.WriteRune(change(c))
	}

	return buf.String()
}

func translate2(s string) string {
	// g -> 9
	// a -> @
	// e -> 3
	// i -> !
	// o -> 0
	// s -> $
	// l -> 1
	// x -> 6
	// z -> 2

	e := make(map[rune]*[]int)

	for i, c := range s {
		switch c {
		case 'g', 'a', 'e', 'i', 'o', 's', 'l', 'x', 'z', 'k', 'f', 't', 'c':
			x := e[c]
			if x != nil {
				*x = append(*x, i)
			} else {
				p := []int{i}
				e[c] = &p
			}
		}
	}

	m := make(map[int]struct{})

	for _, x := range e {
		q := *x
		rand.Shuffle(len(*x), func(i, j int) {
			q[i], q[j] = q[j], q[i]
		})

		m[q[0]] = struct{}{}
	}

	var buf strings.Builder

	for i, c := range s {
		d := change(c)
		if c != d {
			if _, ok := m[i]; !ok {
				d = c
			}
		}

		buf.WriteRune(d)
	}

	return buf.String()
}

func change(c rune) rune {
	switch c {
	case 'c':
		return '('
	case 't':
		return '+'
	case 'f':
		return '='
	case 'k':
		return '<'
	case 'z':
		return '2'
	case 'g':
		return '9'
	case 'a':
		return '@'
	case 'e':
		return '3'
	case 'i':
		return '!'
	case 'o':
		return '0'
	case 's':
		return '$'
	case 'l':
		return '1'
	case 'x':
		return '6'
	default:
		return c
	}
}
