package solution

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f     func([]string) string
		votes []string
		want  string
	}{
		{
			f:     rankTeams,
			votes: []string{"ABC", "ACB", "ABC", "ACB", "ACB"},
			want:  "ACB",
		},
		{
			f:     rankTeams,
			votes: []string{"WXYZ", "XYZW"},
			want:  "XWYZ",
		},
		{
			f:     rankTeams,
			votes: []string{"BCA", "CAB", "CBA", "ABC", "ACB", "BAC"},
			want:  "ABC",
		},
		{
			f:     rankTeams,
			votes: []string{"FVSHJIEMNGYPTQOURLWCZKAX", "AITFQORCEHPVJMXGKSLNZWUY", "OTERVXFZUMHNIYSCQAWGPKJL", "VMSERIJYLZNWCPQTOKFUHAXG", "VNHOZWKQCEFYPSGLAMXJIUTR", "ANPHQIJMXCWOSKTYGULFVERZ", "RFYUXJEWCKQOMGATHZVILNSP", "SCPYUMQJTVEXKRNLIOWGHAFZ", "VIKTSJCEYQGLOMPZWAHFXURN", "SVJICLXKHQZTFWNPYRGMEUAO", "JRCTHYKIGSXPOZLUQAVNEWFM", "NGMSWJITREHFZVQCUKXYAPOL", "WUXJOQKGNSYLHEZAFIPMRCVT", "PKYQIOLXFCRGHZNAMJVUTWES", "FERSGNMJVZXWAYLIKCPUQHTO", "HPLRIUQMTSGYJVAXWNOCZEKF", "JUVWPTEGCOFYSKXNRMHQALIZ", "MWPIAZCNSLEYRTHFKQXUOVGJ", "EZXLUNFVCMORSIWKTYHJAQPG", "HRQNLTKJFIEGMCSXAZPYOVUW", "LOHXVYGWRIJMCPSQENUAKTZF", "XKUTWPRGHOAQFLVYMJSNEIZC", "WTCRQMVKPHOSLGAXZUEFYNJI"},
			want:  "VWFHSJARNPEMOXLTUKICZGYQ",
		},
		{
			f:     rankTeams,
			votes: []string{"XDELQOKZUGMCBTWAJNPFRVIS", "IQKGUDCWMNLSABZEPVXFJRTO", "SFPUJWZABNRLQEMXKIGDVCTO", "FISQWOZAGXJDNLTKPRECVUBM", "BMXIRQOWJSLPDNGAEFKZTUVC", "SCUFTIAQMWNXBGZEOJVDLKRP", "LMNXWKDBTAGPFUCJSROIQZEV", "NVDICPQFBZLRGWTMXJOAUESK", "PGZRUWNFMKTVJSLCDEXOABIQ", "BFKNDXRIGWTVSUJOCMPZAEQL", "MJLDKRNWPQXUFSTEGCOBVAIZ", "VCIMLFJSPZTDRKBAGXWUOEQN", "TERGMVDPILNQSFJCAKZWXBOU", "KOPFCIDQSVAMRBLNTWUZXJGE", "NKWPJGUAFOTZIERXMSLBDCQV", "QUKLSIFCDVZEAMXGJNTWROPB", "DSAGZRFWTMJEILKOPUNQVXCB", "ARQIVZNCMBJGWUPKEOSLFTDX", "ESMATLNJKPRFGXWVUCOQDIZB", "QNGTVDFICPSKMWUJBOZELAXR", "JAOMTSGFQPDWKRNLUXEVBCIZ", "FORAQBEDSTNPUZICVXGWMKLJ", "UWAFSXITRMPEGVQDLZBCNJOK", "IKFOVMNWJLCZBUAGERQPXTSD", "SLEOPWNDBUVIGQJZMRACXTFK", "GTDKMOJSILNAVXFPZQCUWRBE", "AZGRVFJDEUSCNLTBXWQIPMOK", "VQPNUMCJKSBXWIAFTGDOEZRL", "PVNODUXIMJSGACQTWRFBKLZE", "QGKDUAIMEVTXSZFBJLOWNPRC", "MDUOXBVQJPSAGNIKLRFZTCWE", "FMRTEDKWPVLGBICXSNZAQJOU", "WUEFMGXSIRNAOPCBDLQTKJZV", "UDGSPEQVZOBMNICJKXTWALRF", "SFNWLTARVKPJQZMGUDBOICEX", "FSPCMAJZEKNRLDWBGOXQIUVT", "VRUOAMJQGFSIZNXLEPWTDCBK", "VNCMBRLKUPDISJEWOAZGTXFQ", "VGJQZXMPKEIRLTNODCABWFSU", "IMALJVBCXSQPOUGEKNTRFDWZ", "UPJBVKQOAIDGWEZRLTNSXFMC", "IUAPCZFWOQTXBSNLKGVRDMEJ", "DOIRNGMKJLEUCPTWZQSVXFAB", "VJMWTPZFOGBNIDCQAULERKSX", "OAISKWCDLVRQFGNZPUTJEMXB", "NBTGWMEUROJDCKAZLFQXISPV", "DSOBAIJXFWNTUVERGQMCLKPZ", "SBLFZOGCNDWPXAEKTVRJUMIQ", "EKWMAQNURJBZPLFDSGOVIXCT", "BKXRSNVMWGAPCIFLZJTEUQOD", "VQNWJSZTCPEURAKBLIXOFGMD", "TQKFDXCWIUARVBELNOPJZGMS", "MSCPWBTKQUFVDGEXROANJLIZ", "XFQVITPRZMONWCLKBUAJGSED", "MVKLRFQWICOPDJBUSTXNEZGA", "DMKBRGLONJEIFTXZPUACQWSV", "NKSJBQZORVFTXWAUCDLGEMIP", "JQUIXNBRDMOLGSTFVAZCPEWK", "MEBXTJUNVWCQAZGRLIKSPDFO", "EKNUQILMTDPOSFJZWXGVCARB", "JOZNFECPVSXRLGDUKQWBITAM", "UXNZISEGTCVABJMKFLWDQPOR", "IJCFLPGZRXQUMANODKBVSETW", "KCJQRFEWNGAMBIXZVLPDSUTO", "TMSEDRAUWZGQXJPIFVLCOBKN", "ZPRXJNOEKCGFUSDQAILWMVTB", "WTBXANSJLRKFEPOIVDMQUZCG", "IRLVZNQPECOJATFWUSDMGKBX", "BGMPAZOIKWCSUDJNXLTVREQF", "BZXSTPRMLEKWCFIVONGADUQJ", "UFGZCWTARKSBDJXPMVQIENOL", "ZBROLDUFMNQAPJEGITXWCSKV", "FWLJXQDTRCGNVUZMAPBKISEO", "SJXLKAVZCPBIEWFQTNODGRMU", "VAIBDLOEXGUZQWMPNSJTCRKF", "XFLMPECKOQINGWSZJRTVBADU", "ZAJCPMSKUBFWEXQDTVGNLORI", "FUVECNMKZWJBGXTDLPASIOQR", "COFKSVMTNLBWXAREZDQIJUGP", "TKRVFXUIONEALPCDSJGBZWQM", "VJNQLTARZSGCPOMFBUWIXKDE", "GMUSFBPDZATKWEQONXCJILVR", "RFNCLTQZXMSVWPIUBGDEAKOJ", "NARSWLXMEIGQCFBOVDUJPZTK", "OBZVJGSKDUCXAQWPIFNEMTRL", "RIEQDWFSLCGTPKNBAJUVMZOX", "GBAXDSVUEOJNTLPIRKZCWQMF", "ZOETNALWVPGXRIMQKDBFUSCJ", "NAMBJKFZQTOSGWCDPIRXLEUV", "NWQJUZMKIXCGDSBREFPATLOV", "KWAECTUZODSPIJMNLVQRXFBG", "JXPWZNOCGRDKLQITUSEVAFBM", "FNAXSLPRWUBMQJICZTVKDOEG", "GMAWSUXVPZINTODEBRJFKQLC", "FACUKTLWNJSXRQMEDZVGBOIP", "FXEGPVLUJDRNQMTOWSIKAZCB", "ZNEKOJXDVQSUFARMLWPGICTB", "IRSZQGUKVJXCANFWLOBDEMPT", "IUEDPRJFVGQSAXNOWCBMZTLK", "XGKLOPJVQCWTBMSNUZARIEDF", "VSOJUPRFDCNQXTLMBKIGZWEA", "UDGIKPAFNTZWCJRMSLVOXEQB", "GQOIULRNMCJVPFEAZBSDWKTX", "ZJUTMKIFXRSCOPENVGLBDAQW", "ZUQOBCPWNJKVARLTDEGIXFSM", "ADBOKTGCWFMNVUEZLRPIQJSX", "ALDVEQZKCTOJRWBMISPNXFGU", "CEGUFDRXLAVTQNWOKMPZIJBS", "EMBXCVQDJLNWSGIZPRAFTUKO", "KJSQEPNOVGRWZLUIADFBCTXM", "UWVFQEDASJBGXCNPTZIORMLK", "XFLCZSMJAVEPIQDRTUBNKGWO", "URFBQMVATJWEPDCNGSILOZXK", "ZLDPUVIWOJEXGNTQFRBKSACM", "IEQBVDUPNLCGJASFRXTWZOMK", "TVWMROKSUPCXNFEALDQZIGBJ", "GFKEZOICNTXPASMBJVRUWLDQ", "NVWKLBMQJFPIRCSDOZGTXEAU", "QANRCGLPMDZKFOTXBVJEIWUS", "UPCLXNVQZTRKMEBGAOSFDJWI", "PANMJZIFCWGKQDXSRUETOLVB", "JGWVNUXLKAEODTBZIQRCMPFS", "TEOMUIZFSXPGWLADRBQKJVCN", "XPFDJLGZAWEKTRVCMBNUSOQI", "OCQSTRKAZPIMUBNDGEXJVWLF", "MILWTBRAVGZODFJUCSPXQENK", "FREBOSZDPMCNUJXAVKIGLWQT", "TQWPUCERNGIBMXDZKJVASFOL", "APLDRMWCQNFUTSKXVEGJIZBO", "SUWEMPJNFITABXVOCRZLDQKG", "FGRZBWMUTIACNEJKVQPXDOSL", "PBUWQEIDOLXCRNFVSZJTKAMG", "BRFZUGVLWEJKNIASPMTXDQOC", "XLIPMTZQSUFVANKOGBEJDRCW", "GSCRPUAKMEIDBVOLQXWZNFTJ", "KCQMJXRWSFGNBDLAUOTEIVZP", "GKDPJSRMALVEXUNWOICTFQZB"},
			want:  "FVUIGNZSXTMBAKJQEPDORWCL",
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := v.f(v.votes)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/rank-teams-by-votes/solutions/123922/tong-guo-tou-piao-dui-tuan-dui-pai-ming-by-leetcod/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func rankTeams(votes []string) string {
	// 初始化哈希映射
	ranking := make(map[byte][]int)
	for i := 0; i < len(votes[0]); i++ {
		ranking[votes[0][i]] = make([]int, len(votes[0]))
	}
	// 遍历统计
	for _, vote := range votes {
		for i := 0; i < len(vote); i++ {
			ranking[vote[i]][i]++
		}
	}

	// 取出所有的键值对
	result := make([]struct {
		vid  byte
		rank []int
	}, 0, len(ranking))
	for k, v := range ranking {
		result = append(result, struct {
			vid  byte
			rank []int
		}{k, v})
	}
	// 排序
	sort.Slice(result, func(i, j int) bool {
		for k := 0; k < len(result[i].rank); k++ {
			if result[i].rank[k] != result[j].rank[k] {
				return result[i].rank[k] > result[j].rank[k]
			}
		}
		return result[i].vid < result[j].vid
	})
	ans := make([]byte, 0, len(result))
	for _, r := range result {
		ans = append(ans, r.vid)
	}
	return string(ans)
}
