package mock

import "strings"

func ContainWord(word string, sentence string) bool {
	return strings.Contains(sentence, word)
   }
   