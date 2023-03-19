package main

type WordDictionary struct {
	words []string
}

func Constructor() WordDictionary {
	return WordDictionary{
		words: []string{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	this.words = append(this.words, word)
}

func (this *WordDictionary) Search(word string) bool {
	lenWord := len(word)
	for _, wInDic := range this.words {
		if lenWord != len(wInDic) {
			continue
		}

		isFound := true

		for i := 0; i < lenWord; i++ {
			if (word[i] != '.') && (word[i] != wInDic[i]) {
				isFound = false
				break
			}
		}

		if isFound {
			return true
		}

	}

	return false
}
