package main

import (
	"fmt"
	"strings"

	"github.com/wangshizebin/jiebago"
)

func main() {
	sigleTest(`孙女士来电反映：高新区舜华路街道玉兰花园四期13号楼1单元901室业主，801室业主将排烟管道私自改装至自己客厅窗户下，请求有关部门介入督促801室业主恢复排烟管道原貌。环保督察希望相关单位落实处理，请处理并回复来电人。`)
	// allTest()

}

func sigleTest(str string) {
	jieBaGo := jiebago.NewJieBaGo("data/mydict")

	// 提取关键词，即Tag标签
	keywords := jieBaGo.ExtractKeywords(str, 20)
	fmt.Println("提取关键词：", strings.Join(keywords, "/"))

	// 提取带权重的关键词，即Tag标签
	keywordsWeight := jieBaGo.ExtractKeywordsWeight(str, 20)
	fmt.Println("提取带权重的关键词：", keywordsWeight)
	fmt.Println()

}
