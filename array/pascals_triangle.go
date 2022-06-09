package array
//给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
//
//在杨辉三角中，每个数是它左上方和右上方的数的和。

func Generate(numRows int) [][]int {
	var res [][]int
	for i := 0; i < numRows; i++ {
		v := make([]int, i+1)
		for j := range v {
			v[j] = 1
		}
		res = append(res, v)
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res
}
