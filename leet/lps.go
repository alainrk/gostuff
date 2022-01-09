package main

func longestCommonPrefix(strs []string) string {
	for i := 0; i < len(strs[0]); i++ {
		curr := strs[0][i]
		for s := 1; s < len(strs); s++ {
			if len(strs[s]) <= i || strs[s][i] != curr {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
