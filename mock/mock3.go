package mock

func CloseBracket(input string) bool {
	mapAns := map[string]string{
		"[": "]",
		"{": "}",
		"(": ")",
	}

	store := []string{}
	for _, v := range input {
		strValue := string(v)
		value, ok := mapAns[strValue]
		if ok {
			store = append([]string{value}, store...)
			continue
		}
 
		if len(store) == 0 {
			return false
		}

		if store[0] == strValue {
			store = store[1:]
		} else {
			return false
		}
	}
	return true
}
