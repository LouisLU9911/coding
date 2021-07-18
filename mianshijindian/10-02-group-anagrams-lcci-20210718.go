package mianshijindian

func groupAnagrams(strs []string) [][]string {
    words := make(map[string][]string)
    for _, str := range(strs) {
        strSorted := sortWord(str)
        if _, ok := words[strSorted]; ok {
            words[strSorted] = append(words[strSorted], str)
        } else {
            words[strSorted] = []string{str}
        }
    }
    ret := [][]string{}
    for _, v := range(words) {
        ret = append(ret, v)
    }
    return ret
}

func sortWord(word string) string {
    wordBytes := []byte(word)
    sort.Slice(wordBytes, func (i, j int) bool {
        return wordBytes[i] < wordBytes[j]
    })
    return string(wordBytes)
}
