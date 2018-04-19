// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"io/ioutil"
	"unicode/utf8"

	"github.com/shivansh/gogo/tmp/token"
)

const (
	NoState    = -1
	NumStates  = 176
	NumSymbols = 243
)

type Lexer struct {
	src    []byte
	pos    int
	line   int
	column int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (l *Lexer) Scan() (tok *token.Token) {
	tok = new(token.Token)
	if l.pos >= len(l.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = l.pos, l.line, l.column
		return
	}
	start, startLine, startColumn, end := l.pos, l.line, l.column, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
		if l.pos >= len(l.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(l.src[l.pos:])
			l.pos += size
		}

		nextState := -1
		if rune1 != -1 {
			nextState = TransTab[state](rune1)
		}
		state = nextState

		if state != -1 {

			switch rune1 {
			case '\n':
				l.line++
				l.column = 1
			case '\r':
				l.column = 1
			case '\t':
				l.column += 4
			default:
				l.column++
			}

			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				end = l.pos
			case ActTab[state].Ignore != "":
				start, startLine, startColumn = l.pos, l.line, l.column
				state = 0
				if start >= len(l.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = l.pos
			}
		}
	}
	if end > start {
		l.pos = end
		tok.Lit = l.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = start, startLine, startColumn

	return
}

func (l *Lexer) Reset() {
	l.pos = 0
}

/*
Lexer symbols:
0: ';'
1: '\n'
2: ';'
3: '\n'
4: 'b'
5: 'r'
6: 'e'
7: 'a'
8: 'k'
9: 'c'
10: 'a'
11: 's'
12: 'e'
13: 'c'
14: 'o'
15: 'n'
16: 's'
17: 't'
18: 'c'
19: 'o'
20: 'n'
21: 't'
22: 'i'
23: 'n'
24: 'u'
25: 'e'
26: 'd'
27: 'e'
28: 'f'
29: 'a'
30: 'u'
31: 'l'
32: 't'
33: 'e'
34: 'l'
35: 's'
36: 'e'
37: 'f'
38: 'u'
39: 'n'
40: 'c'
41: 'f'
42: 'o'
43: 'r'
44: 'g'
45: 'o'
46: 't'
47: 'o'
48: 'i'
49: 'f'
50: 'i'
51: 'm'
52: 'p'
53: 'o'
54: 'r'
55: 't'
56: 'p'
57: 'a'
58: 'c'
59: 'k'
60: 'a'
61: 'g'
62: 'e'
63: 'r'
64: 'a'
65: 'n'
66: 'g'
67: 'e'
68: 'r'
69: 'e'
70: 't'
71: 'u'
72: 'r'
73: 'n'
74: 's'
75: 't'
76: 'r'
77: 'u'
78: 'c'
79: 't'
80: 's'
81: 'w'
82: 'i'
83: 't'
84: 'c'
85: 'h'
86: 't'
87: 'y'
88: 'p'
89: 'e'
90: 'v'
91: 'a'
92: 'r'
93: 'b'
94: 'o'
95: 'o'
96: 'l'
97: 'i'
98: 'n'
99: 't'
100: 'f'
101: 'l'
102: 'o'
103: 'a'
104: 't'
105: '3'
106: '2'
107: 'f'
108: 'l'
109: 'o'
110: 'a'
111: 't'
112: '6'
113: '4'
114: 'b'
115: 'y'
116: 't'
117: 'e'
118: 's'
119: 't'
120: 'r'
121: 'i'
122: 'n'
123: 'g'
124: 't'
125: 'r'
126: 'u'
127: 'e'
128: 'f'
129: 'a'
130: 'l'
131: 's'
132: 'e'
133: '='
134: ':'
135: '='
136: '.'
137: '.'
138: '''
139: '\'
140: '''
141: '='
142: ','
143: '|'
144: '|'
145: '&'
146: '&'
147: '+'
148: '-'
149: '*'
150: '/'
151: '('
152: ')'
153: '='
154: '='
155: '!'
156: '='
157: '<'
158: '='
159: '<'
160: '>'
161: '='
162: '>'
163: '!'
164: '{'
165: '}'
166: '.'
167: '['
168: ']'
169: ':'
170: 'd'
171: 'e'
172: 'f'
173: 'e'
174: 'r'
175: '+'
176: '+'
177: '-'
178: '-'
179: '='
180: '='
181: '!'
182: '='
183: '<'
184: '<'
185: '='
186: '>'
187: '>'
188: '='
189: '+'
190: '-'
191: '|'
192: '^'
193: '*'
194: '/'
195: '%'
196: '<'
197: '<'
198: '>'
199: '>'
200: '&'
201: '&'
202: '^'
203: '/'
204: '/'
205: '\n'
206: '/'
207: '*'
208: '*'
209: '*'
210: '/'
211: '_'
212: '0'
213: '0'
214: 'x'
215: 'X'
216: 'e'
217: 'E'
218: '+'
219: '-'
220: '`'
221: '`'
222: '"'
223: '\'
224: '"'
225: '"'
226: '\'
227: 'n'
228: '\'
229: 'r'
230: '\'
231: 't'
232: ' '
233: '\t'
234: '\r'
235: 'a'-'z'
236: 'A'-'Z'
237: '0'-'9'
238: '0'-'7'
239: 'a'-'f'
240: 'A'-'F'
241: '1'-'9'
242: .
*/
