package problems

//有效的字母异位词
//判断 s和t是不是字母异位词
//如果s和t每个字母出现的次数相同，那么就是字母异位词

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	countS := map[rune]int{}
	countT := map[rune]int{}

	for _, c := range s {
		_, ok := countS[c]
		if !ok {
			countS[c] = 0
		}
		countS[c] += 1
	}

	for _, c := range t {
		_, ok := countT[c]
		if !ok {
			countT[c] = 0
		}
		countT[c] += 1
	}

	for key := range countS {
		_, ok := countT[key]

		if !ok || countT[key] != countS[key] {
			return false
		}

	}

	return true
}

func init() {
	_ = isAnagram("anagram", "nagaram")
}
