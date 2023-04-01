package stacks

const (
	openParenthesis    = '('
	closeParenthesis   = ')'
	openSquareBracket  = '['
	closeSquareBracket = ']'
	openCurlyBracket   = '{'
	closeCurlyBracket  = '}'

	space = ' '
)

// isValid keeps track of the closing brackets that are expected to match the open brackets encountered in the string,
//and checks whether each closing bracket matches the corresponding open bracket.
// Return:
// 	False:
//		If any brackets are not balanced
//	True:
//		otherwise
func isValid(s string) bool {
	var stack Stack
	for _, c := range s {
		switch c {
		case openCurlyBracket:
			stack.Push(closeCurlyBracket)
		case openSquareBracket:
			stack.Push(closeSquareBracket)
		case openParenthesis:
			stack.Push(closeParenthesis)
		case space:
			continue
		default:
			if stack.IsEmpty() || stack.Pop() != byte(c) {
				return false
			}
		}
	}
	return stack.IsEmpty()
}
