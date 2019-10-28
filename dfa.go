package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"strings"
)

type dfa struct {
	state  []string
	statei map[string]int
	alph   []rune
	alphi  map[rune]int
	start  int
	final  *bitset
	delta  [][]int
}

func newDfa() *dfa {
	var dfa dfa
	dfa.state = make([]string, 0)
	dfa.statei = make(map[string]int)
	dfa.alph = make([]rune, 0)
	dfa.alphi = make(map[rune]int)
	dfa.start = 0
	dfa.delta = make([][]int, 0)
	return &dfa
}

func (nfa *nfa) toDfa() *dfa {
	qu := list.New()
	vi := newhashset()

	dfa := newDfa()
	dfa.alph = nfa.alph
	dfa.alphi = nfa.alphi

	final := make([]int, 0)
	start := newBitset(len(nfa.state))
	start.set(nfa.start)
	qu.PushBack(start)
	for qu.Len() > 0 {
		st := qu.Remove(qu.Front()).(*bitset)
		if !vi.has(st.String()) {
			vi.insert(st.String())
			sti := dfa.insertState(nfa.tuple(st))

			if !st.clone().intersection(nfa.final).isEmpty() {
				final = append(final, sti)
			}
			for r := range nfa.alph {
				next := newBitset(len(nfa.state))
				for _, s := range st.array() {
					if nfa.delta[s] != nil {
						next.union(nfa.delta[s][r])
					}
				}
				if !next.isEmpty() {
					ni := dfa.insertState(nfa.tuple(next))
					dfa.delta[sti][r] = ni
					if !vi.has(next.String()) {
						qu.PushBack(next)
					}
				} else {
					dfa.delta[sti][r] = -1
				}
			}
			dfa.final = newBitset(len(dfa.state))
			for _, f := range final {
				dfa.final.set(f)
			}
		}
	}
	return dfa
}

//regresa el indice para el estado s en el dfa, si el estado no esta lo inserta
func (dfa *dfa) insertState(s string) int {
	if i, ok := dfa.statei[s]; ok {
		return i
	}
	i := len(dfa.state)
	dfa.statei[s] = i
	dfa.delta = append(dfa.delta, make([]int, len(dfa.alph)))
	for r := range dfa.delta[i] {
		dfa.delta[i][r] = -1
	}
	dfa.state = append(dfa.state, s)
	return i
}

func (dfa *dfa) String() string {
	var sb strings.Builder
	sm := 0
	sb.WriteRune('{')
	for i, s := range dfa.state {
		sb.WriteString(s)
		if i < len(dfa.state)-1 {
			sb.WriteRune(' ')
		}
		if len(s) > sm {
			sm = len(s)
		}
	}

	sb.WriteString("}\n{")
	for i, r := range dfa.alph {
		sb.WriteRune(r)
		if i < len(dfa.alph)-1 {
			sb.WriteRune(' ')
		}
	}

	sb.WriteString("}\n")
	sb.WriteString(dfa.state[dfa.start] + "\n{")

	f := dfa.final.array()
	for i, k := range f {
		sb.WriteString(dfa.state[k])
		if i < len(f)-1 {
			sb.WriteRune(' ')
		}
	}

	sb.WriteString("}\n")
	frmt := fmt.Sprintf("%%-%ds", sm)
	for st, rs := range dfa.delta {
		if rs != nil {
			sb.WriteString(fmt.Sprintf(frmt, dfa.state[st]))
			for _, bs := range rs {
				if bs >= 0 {
					sb.WriteRune(',')
					sb.WriteString(fmt.Sprintf(frmt, dfa.state[bs]))
				} else {
					sb.WriteString(fmt.Sprintf(frmt, "_"))
				}
			}
			sb.WriteRune('\n')
		}
	}
	return sb.String()
}

func (dfa *dfa) write(w io.Writer) {
	bw := bufio.NewWriter(w)
	bw.WriteString("#DFA file generated\n")
	bw.WriteString("#States\n")
	sm := 0
	for i, s := range dfa.state {
		bw.WriteString(s)
		if i < len(dfa.state)-1 {
			bw.WriteString(", ")
		}
		if len(s) > sm {
			sm = len(s)
		}
	}
	bw.WriteString("\n\n#Alphabet\n")
	for i, r := range dfa.alph {
		bw.WriteRune(r)
		if i < len(dfa.alph)-1 {
			bw.WriteString(", ")
		}
	}
	bw.WriteString("\n\n#Start state\n")
	bw.WriteString(dfa.state[dfa.start])
	bw.WriteString("\n\n#Final states\n")
	f := dfa.final.array()
	for i, k := range f {
		bw.WriteString(dfa.state[k])
		if i < len(f)-1 {
			bw.WriteString(", ")
		}
	}
	bw.WriteString("\n\n#Delta\n")
	frmt := fmt.Sprintf("%%-%dv", sm)
	ddfm := fmt.Sprintf("%%-%dc", sm+1)
	dd := fmt.Sprintf(frmt, "#s") + " "
	for _, r := range dfa.alph {
		dd += fmt.Sprintf(ddfm, r)
	}
	bw.WriteString(dd + "\n")
	for st, rs := range dfa.delta {
		if rs != nil {
			bw.WriteString(fmt.Sprintf(frmt, dfa.state[st]))
			for _, bs := range rs {
				if bs >= 0 {
					bw.WriteRune(',')
					bw.WriteString(fmt.Sprintf(frmt, dfa.state[bs]))
				} else {
					bw.WriteString(fmt.Sprintf(frmt, "_"))
				}
			}
			bw.WriteRune('\n')
		}
	}
	bw.Flush()
}
