package main

import (
	"fmt"
	"strings"

	"github.com/wangshizebin/jiebago"
)

func main() {
	sigleTest(`高级会计师职称评审申报材料未通过省高评委办事机构审核,可以多次进行修改吗? 
不可以。省高评委办事机构在审核过程中发现申报人员不符合申报条件或申报材料弄虚作假的,将直接退回并说明不得再次报送。对不规范、不完整的申报材料直接退回申报人并一次性告知修改意见,提供一次修改机会。申报人员应一次修改完善,如修改提报后,仍不符合申报要求的,不再提供修改机会`)
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

func allTest() {
	jieBaGo := jiebago.NewJieBaGo("data/mydict")
	// 可以指定字典库的位置
	// jieBaGo := jiebago.NewJieBaGo("/data/mydict")

	sentence := "我喜欢上班"
	fmt.Println("原始语句：", sentence)
	fmt.Println()

	// 默认模式分词
	words := jieBaGo.Cut(sentence)
	fmt.Println("默认模式分词：", strings.Join(words, "/"))

	// 精确模式分词
	words = jieBaGo.CutAccurate(sentence)
	fmt.Println("精确模式分词：", strings.Join(words, "/"))

	// 全模式分词
	words = jieBaGo.CutFull(sentence)
	fmt.Println("全模式分词：", strings.Join(words, "/"))

	// NoHMM模式分词
	words = jieBaGo.CutNoHMM(sentence)
	fmt.Println("NoHMM模式分词：", strings.Join(words, "/"))

	// 搜索引擎模式分词
	words = jieBaGo.CutForSearch(sentence)
	fmt.Println("搜索引擎模式分词：", strings.Join(words, "/"))
	fmt.Println()

	// 提取关键词，即Tag标签
	keywords := jieBaGo.ExtractKeywords(sentence, 20)
	fmt.Println("提取关键词：", strings.Join(keywords, "/"))

	// 提取带权重的关键词，即Tag标签
	keywordsWeight := jieBaGo.ExtractKeywordsWeight(sentence, 20)
	fmt.Println("提取带权重的关键词：", keywordsWeight)
	fmt.Println()

	// 向字典加入单词
	exist, err := jieBaGo.AddDictWord("编程宝库", 3, "n")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("向字典加入单词：编程宝库")
		if exist {
			fmt.Println("单词已经存在")
		}
	}

	// 向字典加入停止词
	exist, err = jieBaGo.AddStopWord("the")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("向字典加入停止词：the")
		if exist {
			fmt.Println("单词已经存在")
		}
	}
}
