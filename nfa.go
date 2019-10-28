package main

import (
	"bufio"
	"io"
	"strings"
	"unicode/utf8"
)

type nfa struct {
	state  []string
	alph   []rune
	statei map[string]int
	alphi  map[rune]int
	start  int
	delta  [][]*bitset
	final  *bitset
}

func newNfa(in io.Reader) *nfa {
	reader := bufio.NewReader(in)
	var nfa nfa
	nfa.statei = make(map[string]int)

	line := nextLine(reader)
	states := strings.Split(line, ",")
	for i := range states {
		states[i] = strings.Trim(states[i], " ")
		nfa.statei[states[i]] = i
	}
	nfa.state = states
	nfa.delta = make([][]*bitset, len(nfa.state))

	line = nextLine(reader)
	nfa.alph = make([]rune, 0)
	nfa.alphi = make(map[rune]int)
	for i, str := range strings.Split(line, ",") {
		r, _ := utf8.DecodeRuneInString(strings.Trim(str, " "))
		nfa.alph = append(nfa.alph, r)
		nfa.alphi[r] = i
	}

	line = nextLine(reader)
	nfa.start = nfa.statei[strings.Trim(line, " ")]
	nfa.final = newBitset(len(nfa.state))

	line = nextLine(reader)
	for _, str := range strings.Split(line, ",") {
		str = strings.Trim(str, " ")
		if i, ok := nfa.statei[str]; ok {
			nfa.final.set(i)
		}
	}

	for line = nextLine(reader); len(line) > 0; line = nextLine(reader) {
		d := strings.Split(line, ",")
		st := nfa.statei[strings.Trim(d[0], " ")]
		nfa.delta[st] = make([]*bitset, len(nfa.alph))
		for r, sts := range d[1:] {
			sts = strings.Trim(sts, " ")
			if len(sts) > 0 && sts != "_" {
				nfa.delta[st][r] = newBitset(len(nfa.state))
				for _, srt := range strings.Split(sts, " ") {
					std := nfa.statei[strings.Trim(srt, " ")]
					nfa.delta[st][r].set(std)
				}
			}
		}
	}
	return &nfa
}

func (a *nfa) String() string {
	var sb strings.Builder
	sb.WriteRune('{')
	for i, s := range a.state {
		sb.WriteString(s)
		if i < len(a.state)-1 {
			sb.WriteRune(' ')
		}
	}
	sb.WriteString("}\n{")
	for i, r := range a.alph {
		sb.WriteRune(r)
		if i < len(a.alph)-1 {
			sb.WriteRune(' ')
		}
	}
	sb.WriteString("}\n")
	sb.WriteString(a.state[a.start] + "\n{")
	f := a.final.array()
	for i, k := range f {
		sb.WriteString(a.state[k])
		if i < len(f)-1 {
			sb.WriteRune(' ')
		}
	}
	sb.WriteString("}\n")
	for st, rs := range a.delta {
		if rs != nil {
			sb.WriteString(a.state[st])
			for _, bs := range rs {
				sb.WriteRune(',')
				if bs != nil {
					sb.WriteString(a.tuple(bs))
				} else {
					sb.WriteRune('_')
				}
			}
			sb.WriteRune('\n')
		}
	}
	return sb.String()
}

func (a *nfa) tuple(b *bitset) string {
	var sb strings.Builder
	sb.WriteRune('[')
	sts := b.array()
	for i, si := range sts {
		sb.WriteString(a.state[si])
		if i < len(sts)-1 {
			sb.WriteRune(' ')
		}
	}
	sb.WriteRune(']')
	return sb.String()
}

func nextLine(reader *bufio.Reader) string {
	str := ""
	for len(str) == 0 || str[0] == '#' {
		s, err := reader.ReadString('\n')
		s = strings.Trim(s, "\n")
		if err != nil {
			return s
		}
		str = s
	}
	return str
}
