package testQuestion

// abcde = ["ab","cd","e*"]
func Question1(string) []string {
	questions := "abcde"
	answer := []string{}

	if len(questions)%2 != 0 {
		questions = questions + "*"
	}

	for i := 0; i < len(questions); i = i + 2 {

		x := questions[i : i+2]
		answer = append(answer, x)

	}

	return answer
}
