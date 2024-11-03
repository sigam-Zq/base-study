package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/go-ego/gse"
	"github.com/go-ego/gse/hmm/pos"
	"github.com/yanyiwu/gojieba"
)

func main() {

	log.Println("-------------------------------------")
	testGse()

	log.Println("-------------------------------------")

	testJieba()
}

var (
	seg    gse.Segmenter
	posSeg pos.Segmenter

	new, _ = gse.New("zh,testdata/test_en_dict3.txt", "alpha")

	text = "你好世界, Hello world, Helloworld."
)

func testGse() {
	// 加载默认词典
	seg.LoadDict()
	// 加载默认 embed 词典
	// seg.LoadDictEmbed()
	//
	// 加载简体中文词典
	seg.LoadDict("zh_s")
	seg.LoadDictEmbed("zh_s")
	//
	// 加载繁体中文词典
	// seg.LoadDict("zh_t")
	//
	// 加载日文词典
	// seg.LoadDict("jp")
	//
	// 载入词典
	// seg.LoadDict("your gopath"+"/src/github.com/go-ego/gse/data/dict/dictionary.txt")

	cut()

	segCut()
}

func cut() {
	hmm := new.Cut(text, true)
	fmt.Println("cut use hmm: ", hmm)

	hmm = new.CutSearch(text, true)
	fmt.Println("cut search use hmm: ", hmm)
	fmt.Println("analyze: ", new.Analyze(hmm, text))

	hmm = new.CutAll(text)
	fmt.Println("cut all: ", hmm)

	reg := regexp.MustCompile(`(\d+年|\d+月|\d+日|[\p{Latin}]+|[\p{Hangul}]+|\d+\.\d+|[a-zA-Z0-9]+)`)
	text1 := `헬로월드 헬로 서울, 2021年09月10日, 3.14`
	hmm = seg.CutDAG(text1, reg)
	fmt.Println("Cut with hmm and regexp: ", hmm, hmm[0], hmm[6])
}

func analyzeAndTrim(cut []string) {
	a := seg.Analyze(cut, "")
	fmt.Println("analyze the segment: ", a)

	cut = seg.Trim(cut)
	fmt.Println("cut all: ", cut)

	fmt.Println(seg.String(text, true))
	fmt.Println(seg.Slice(text, true))
}

func cutPos() {
	po := seg.Pos(text, true)
	fmt.Println("pos: ", po)
	po = seg.TrimPos(po)
	fmt.Println("trim pos: ", po)

	posSeg.WithGse(seg)
	po = posSeg.Cut(text, true)
	fmt.Println("pos: ", po)

	po = posSeg.TrimWithPos(po, "zg")
	fmt.Println("trim pos: ", po)
}

func segCut() {
	// 分词文本
	tb := "山达尔星联邦共和国联邦政府"

	// 处理分词结果
	fmt.Println("输出分词结果, 类型为字符串, 使用搜索模式: ", seg.String(tb, true))
	fmt.Println("输出分词结果, 类型为 slice: ", seg.Slice(tb))

	tb2 := []byte("山达尔星联邦共和国联邦政府")
	segments := seg.Segment(tb2)
	// 处理分词结果, 普通模式
	fmt.Println(gse.ToString(segments))

	segments1 := seg.Segment([]byte(text))
	// 搜索模式
	fmt.Println(gse.ToString(segments1, true))
}

func testJieba() {

	var s string
	var words []string
	use_hmm := true
	x := gojieba.NewJieba()
	defer x.Free()

	s = `不好了,我们两个兄弟被老汉奸杀了`
	words = x.CutAll(s)
	fmt.Println(s)
	fmt.Println("全模式:", strings.Join(words, "/"))

	s = "我来到北京清华大学"
	words = x.CutAll(s)
	fmt.Println(s)
	fmt.Println("全模式:", strings.Join(words, "/"))

	words = x.Cut(s, use_hmm)
	fmt.Println(s)
	fmt.Println("精确模式:", strings.Join(words, "/"))
	s = "比特币"
	words = x.Cut(s, use_hmm)
	fmt.Println(s)
	fmt.Println("精确模式:", strings.Join(words, "/"))

	x.AddWord("比特币")
	// `AddWordEx` 支持指定词语的权重，作为 `AddWord` 权重太低加词失败的补充。
	// `tag` 参数可以为空字符串，也可以指定词性。
	// x.AddWordEx("比特币", 100000, "")
	s = "比特币"
	words = x.Cut(s, use_hmm)
	fmt.Println(s)
	fmt.Println("添加词典后,精确模式:", strings.Join(words, "/"))

	s = "他来到了网易杭研大厦"
	words = x.Cut(s, use_hmm)
	fmt.Println(s)
	fmt.Println("新词识别:", strings.Join(words, "/"))

	s = "小明硕士毕业于中国科学院计算所，后在日本京都大学深造"
	words = x.CutForSearch(s, use_hmm)
	fmt.Println(s)
	fmt.Println("搜索引擎模式:", strings.Join(words, "/"))

	s = "长春市长春药店"
	words = x.Tag(s)
	fmt.Println(s)
	fmt.Println("词性标注:", strings.Join(words, ","))

	s = "区块链"
	words = x.Tag(s)
	fmt.Println(s)
	fmt.Println("词性标注:", strings.Join(words, ","))

	s = "长江大桥"
	words = x.CutForSearch(s, !use_hmm)
	fmt.Println(s)
	fmt.Println("搜索引擎模式:", strings.Join(words, "/"))

	wordinfos := x.Tokenize(s, gojieba.SearchMode, !use_hmm)
	fmt.Println(s)
	fmt.Println("Tokenize:(搜索引擎模式)", wordinfos)

	wordinfos = x.Tokenize(s, gojieba.DefaultMode, !use_hmm)
	fmt.Println(s)
	fmt.Println("Tokenize:(默认模式)", wordinfos)

	keywords := x.ExtractWithWeight(s, 5)
	fmt.Println("Extract:", keywords)
}
