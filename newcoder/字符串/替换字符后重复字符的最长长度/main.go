package main

/*
思路：
统计出现频次最多的字符，窗口长度减去窗口中出现频次最大的长度
如果差值比 操作步骤 k 大，则缩小左窗口直到等于 k
取出窗口长度的最大值即可
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func characterReplacement(s string, k int) int {
	res, left, counter, freq := 0, 0, 0, make([]int, 26)
	// 结果值，窗口做边界，统计计数器，字符映射表
	for right := 0; right < len(s); right++ {
		// 统计每一个字符的出现次数
		freq[s[right]-'A']++
		// 记录当前窗口中出现频率最多的字母
		counter = max(counter, freq[s[right]-'A'])
		// 窗口中除了出现频率最多的字符其他的字符都要进行替换操作
		for right-left+1-counter > k {
			// 需要替换的字符大于指定的替换次数
			// 窗口左边界移动
			freq[s[left]-'A']--
			left++
		}
		// 记录过程中出现的最大窗口值即为进行 k 次替换操作后的最长重复字符串
		res = max(res, right-left+1)
	}
	return res
}
