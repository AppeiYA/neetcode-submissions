func isOperation(char string) bool {
	switch char {
		case "+", "-", "*", "/":
			return true
		default:
			return false
	}
}

func evalRPN(tokens []string) int {
	cStack := []int{}
	if len(tokens) == 1 {
		num, _ := strconv.Atoi(tokens[0])
		return num
	}

	for _, v := range tokens {
		is_operation := isOperation(v)
		if !is_operation {
			num, _ := strconv.Atoi(v)
			cStack = append(cStack, num)
			continue
		}
		if is_operation && len(cStack) < 2 {
			return 0
		}else {
			var top int = 0
			var op1, op2 = cStack[len(cStack)-2], cStack[len(cStack)-1]
			cStack = cStack[:len(cStack)-2]
			switch v {
				case "+":
					top = op1 + op2
				case "*": 
					top = op1 * op2
				case "-": 
					top = op1 - op2
				case "/":
					if op2 == 0 {
						return 0
					}
					top = op1 / op2
				default: 
					top = 0
			}

			cStack = append(cStack, top)
		}
	}

	return cStack[len(cStack)-1]
}
