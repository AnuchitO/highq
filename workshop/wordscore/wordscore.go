package wordscore

import "strings"

/*
# Word Score

Given a word, computes the word score for that word.

## Letter Values

You'll need these:

```plain
Letter                           Value
A, E, I, O, U, L, N, R, S, T       1
D, G                               2
B, C, M, P                         3
F, H, V, W, Y                      4
K                                  5
J, X                               8
Q, Z                               10
```

## Examples
"anuchitoO" should be scored as worth 14 points:

- 1 points for A
- 1 points for N
- 1 points for U
- 3 points for C
- 4 points for H
- 1 points for I
- 1 points for T
- 2 points for O, twice

And to total:

- `1 + 1 + 1 + 3 + 4 + 1 + 1 + 2*2`
- = 14

## Extensions
- You can play a `:double` or a `:triple` letter.
- You can play a `:double` or a `:triple` word.
*/

// Score the word.
func Score(word string) int {
	score := 0
	for _, l := range strings.ToUpper(word) {
		score += letterScores[l]
	}
	return score
}

var letterScores = map[rune]int{
	'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1,
	'L': 1, 'N': 1, 'R': 1, 'S': 1, 'T': 1,
	'D': 2, 'G': 2,
	'B': 3, 'C': 3, 'M': 3, 'P': 3,
	'F': 4, 'H': 4, 'V': 4, 'W': 4, 'Y': 4,
	'K': 5,
	'J': 8, 'X': 8,
	'Q': 10, 'Z': 10,
}
