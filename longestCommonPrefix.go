package leetcodes

/**

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".



Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.


Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters if it is non-empty.
*/

func longestCommonPrefix(strs []string) string {
	var prefix = []byte(strs[0])
	if len(strs) > 1 {
		for i := 1; i < len(strs); i++ {
			if len(strs[i]) == 0 {
				return ""
			}

			if len(strs[i]) < len(prefix) {
				prefix = prefix[:len(strs[i])]
			}
			for j := 0; j < len(strs[i]) && j < len(prefix); j++ {

				if strs[i][j] != prefix[j] {
					prefix = prefix[:j]
				}
			}
		}

		return string(prefix)
	}
	return string(prefix)
}
