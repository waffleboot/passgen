package main

type charType int

const (
	numberCharType charType = iota + 1
	symbolCharType
)

type replacement struct {
	ty charType
	to rune
}

type wordStats struct {
	replace map[int]replacement
	numbers []int
	symbols []int
}

func makeWordStats(word string) wordStats {
	stats := wordStats{
		replace: make(map[int]replacement),
	}

	for i, c := range word {
		switch c {
		case 'c':
			stats.replace[i] = replacement{ty: symbolCharType, to: '('}
			stats.symbols = append(stats.symbols, i)
		case 't':
			stats.replace[i] = replacement{ty: symbolCharType, to: '+'}
			stats.symbols = append(stats.symbols, i)
		case 'f':
			stats.replace[i] = replacement{ty: symbolCharType, to: '='}
			stats.symbols = append(stats.symbols, i)
		case 'k':
			stats.replace[i] = replacement{ty: symbolCharType, to: '<'}
			stats.symbols = append(stats.symbols, i)
		case 'z':
			stats.replace[i] = replacement{ty: numberCharType, to: '2'}
			stats.numbers = append(stats.numbers, i)
		case 'g':
			stats.replace[i] = replacement{ty: numberCharType, to: '9'}
			stats.numbers = append(stats.numbers, i)
		case 'a':
			stats.replace[i] = replacement{ty: symbolCharType, to: '@'}
			stats.symbols = append(stats.symbols, i)
		case 'e':
			stats.replace[i] = replacement{ty: numberCharType, to: '3'}
			stats.numbers = append(stats.numbers, i)
		case 'i':
			stats.replace[i] = replacement{ty: symbolCharType, to: '!'}
			stats.symbols = append(stats.symbols, i)
		case 'o':
			stats.replace[i] = replacement{ty: numberCharType, to: '0'}
			stats.numbers = append(stats.numbers, i)
		case 's':
			stats.replace[i] = replacement{ty: symbolCharType, to: '$'}
			stats.symbols = append(stats.symbols, i)
		case 'l':
			stats.replace[i] = replacement{ty: numberCharType, to: '1'}
			stats.numbers = append(stats.numbers, i)
		case 'x':
			stats.replace[i] = replacement{ty: numberCharType, to: '6'}
			stats.numbers = append(stats.numbers, i)
		}
	}

	return stats
}
