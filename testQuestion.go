package testQuestion

// abcde = ["ab","cd","e*"]
func Question1(text string) []string {
	answer := []string{}

	if len(text)%2 != 0 {
		text = text + "*"
	}

	for i := 0; i < len(text); i = i + 2 {

		x := text[i : i+2]
		answer = append(answer, x)

	}

	return answer
}
