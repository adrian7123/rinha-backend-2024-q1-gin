package helpers

func ArrayReverse(input []any) []any {
	if len(input) == 0 {
		return input
	}
	return append(ArrayReverse(input[1:]), input[0])
}
