package leetcode83

import (
	"log"
	"reflect"
	"strconv"
	"testing"
)

//  3297. 统计重新排列后包含另一个字符串的子字符串数目 I

/*
给你两个字符串 word1 和 word2 。

如果一个字符串 x 重新排列后，word2 是重排字符串的
前缀
 ，那么我们称字符串 x 是 合法的 。

请你返回 word1 中 合法
子字符串
 的数目。



示例 1：

输入：word1 = "bcca", word2 = "abc"

输出：1

解释：

唯一合法的子字符串是 "bcca" ，可以重新排列得到 "abcc" ，"abc" 是它的前缀。

示例 2：

输入：word1 = "abcabc", word2 = "abc"

输出：10

解释：

除了长度为 1 和 2 的所有子字符串都是合法的。

示例 3：

输入：word1 = "abcabc", word2 = "aaabc"

输出：0



解释：

1 <= word1.length <= 105
1 <= word2.length <= 104
word1 和 word2 都只包含小写英文字母。
*/

/*
chatGPT 辅助理解题目

### 具体化题目要求

题目要求判断字符串 `word1` 中的子字符串是否是“合法”的。以下是对题目要求的详细解析：

1. **什么是合法子字符串？**
   - 一个字符串 `x` 是“合法”的，当且仅当：
     1. 它的所有字符经过重新排列后，得到的新字符串可以以 `word2` 为前缀。
     2. `x` 需要包含至少和 `word2` 的前缀匹配的字符。

2. **如何判断？**
   - 取出 `word1` 的所有子字符串。
   - 对每个子字符串 `x`，统计字符的频次（也就是每个字符出现的次数）。
   - 检查这些频次是否能满足 `word2` 的前缀要求。

---

### 示例 1 的分析

#### 输入：
- `word1 = "bcca"`
- `word2 = "abc"`

#### 解释：
- `word2` 的前缀是 `abc`，所以我们需要满足 `a=1, b=1, c=1` 的频次要求。

#### 遍历 `word1` 的所有子字符串：
1. `"b"`：频次 `b=1`，不满足 `abc` 的前缀。
2. `"bc"`：频次 `b=1, c=1`，不满足 `abc` 的前缀。
3. `"bcc"`：频次 `b=1, c=2`，不满足 `abc` 的前缀。
4. `"bcca"`：频次 `b=1, c=2, a=1`，可以重新排列为 `abcc`，满足前缀 `abc`。**合法**。
5. `"c"`：频次 `c=1`，不满足 `abc` 的前缀。
6. `"cc"`：频次 `c=2`，不满足 `abc` 的前缀。
7. `"cca"`：频次 `c=2, a=1`，不满足 `abc` 的前缀。
8. `"a"`：频次 `a=1`，不满足 `abc` 的前缀。

#### 总结：
- 唯一合法的子字符串是 `"bcca"`，所以答案是 `1`。

---

### 示例 2 的分析

#### 输入：
- `word1 = "abcabc"`
- `word2 = "abc"`

#### 解释：
- `word2` 的前缀是 `abc`，需要满足 `a=1, b=1, c=1` 的频次要求。

#### 遍历 `word1` 的所有子字符串：
1. `"abc"`：频次 `a=1, b=1, c=1`，**合法**。
2. `"abca"`：频次 `a=2, b=1, c=1`，可以重新排列为 `abc...`，**合法**。
3. `"abcab"`：频次 `a=2, b=2, c=1`，可以重新排列为 `abc...`，**合法**。
4. `"abcabc"`：频次 `a=2, b=2, c=2`，可以重新排列为 `abc...`，**合法**。

其他合法子字符串分别为长度为 1 或 2 的子字符串（可以直接满足要求）：
- 单字符：`"a"`, `"b"`, `"c"`。
- 长度为 2 的子字符串：`"ab"`, `"bc"`, `"ca"`。

#### 总计：
- 共有 10 个合法子字符串：
  - 长度为 1 的 3 个：`"a"`, `"b"`, `"c"`;
  - 长度为 2 的 3 个：`"ab"`, `"bc"`, `"ca"`;
  - 长度为 3 的 1 个：`"abc"`;
  - 长度为 4 的 1 个：`"abca"`;
  - 长度为 5 的 1 个：`"abcab"`;
  - 长度为 6 的 1 个：`"abcabc"`。

答案为 `10`。

---

### 示例 3 的分析

#### 输入：
- `word1 = "abcabc"`
- `word2 = "aaabc"`

#### 解释：
- `word2` 的前缀要求是 `a=3, b=1, c=1`。
- `word1` 的所有子字符串均无法满足这个频次条件，因此没有合法子字符串。

#### 答案：
- `0`。
*/

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f       func(string, string) int64
		word1   string
		word2   string
		want    int64
		isDebug bool
	}{
		{
			f:       validSubstringCount,
			word1:   "bcca",
			word2:   "abc",
			want:    int64(1),
			isDebug: false,
		},
		{
			f:       validSubstringCount,
			word1:   "abcabc",
			word2:   "abc",
			want:    int64(10),
			isDebug: false,
		},
		{
			f:       validSubstringCount,
			word1:   "bbbb",
			word2:   "b",
			want:    int64(10),
			isDebug: false,
		},
		{
			f:       validSubstringCount,
			word1:   "dcbdcdccb",
			word2:   "cdd",
			want:    int64(18),
			isDebug: false,
		},
		{
			f:       validSubstringCount,
			word1:   "aeghijedddbabcaefgdhfefjbgjfbjeihdbjgfacaafdjjdadciaighdigabdjaacbcfffjdebbghdehbafdffefcjcdagdeahicaiiacdcafidjhfceaibcdacjidgccdjegcicgeefdjjfgjcebjdcbjibcbacjgfbahffaeejfgcejbcjagbadibfhdaahedhedgehehheedcgibbjgjjfcecdiebiajbebfiediibfejdhehbehejggcfejhhchfbddiafgceedadabcicajfcaccghhbahabfifigegbgaghbabhieiejbgbgchcigcfbdeebehbeihcjddhjhbjbhbjaejgcjeaeggfjebeehhcfahceaihcbdhegebceiigdfbhjibeifchacgjjbfcbcfigeffgacbdbfiehciajfjbdhebhbdeaihbghdaiajggdagjgjafcahbddeibhhjbibbjfjfigaiedgdagficdcfgcegiejgaigdhiiagfgdehfajjihfegccgcaehhbhiadcagiijdfaejfefebeajbibdgfgdjbeaedhedjaddcjjefbbichjbihidiaccadcfjcbjbacbjgbebijichjgjaeffcefgiejcchhffcgjbchjhfeigjgehbbgfbbabiejeijgcbgedhgfaahjaedfceadggfhaiihbajjajdhggfhibachfbfahaieddbdaahcaiedhfebgeggibcggdhdijbjfhajiijbaiaahhfhjidcbjejecdjeehbbbcjheiaghgdcfegddiehgbjadiedcdbffhegbgaggcfebfhagfaajbcdhhdbejggfheihfefddddjjeeaiedbahjhfddhbhhfjjjgdahbhfbjfhgehdbfbfcchjgbbcaggcigbjijbeiaajiadhfegijaahcaehegjghcjjbjjahdaafaihigagficgbhfahcchadgcdaiggihfbjgbgbhiccbffjahhfhaifdagcajfeabeejdhjejjggcehibahgccehchaijibgbgdebdbehgfjggcidbcbcjfffidhhhjajegjaejeheajbiabbdafhehbggfjhbjggcgecagbfecbchhiajaibidfjhgighicaghhhdffhhcbfjijiagbceafgcgjbhcidcicghhgfdhidjadeecfbccdjiifijbhichhafcgfcbbigefafecegaicaacgjcaaddcdiheehiebbidjifaebddfcdahdgaahdehaajeciccbjibibbbidebjjgbabdhheeiaaiiacehagfejjcahcfgbgefibafbacijjddheeegdediciacabaifhafejabeeaeihbhfedfdhcfedfcafahagcgidfcbfcgjeeeijejhigigdghijbgjihhbeigjjbabfgjaefdafbihjhghjgdhgccdcgahhhaafciddddfaeffhbjhjiifihjbffaagegidbggbaaeeiageiidehdfbddbbcbcchffeabhheggbfgjggfedjfiaagbgcjbabdfidgfceabgedbjhabeiecchefchefbgdhfcjcejcdcdbcaggcjhfgeighebchgchieaecgjehiicfjbggefacaiaehjjaafihaagfgggjdbccbcbhbdjegggjgdfddaaeijfgchfdjafjgiaefbjciagcjadejigidbcgfhaicdiaacichiiaahiffhjbijgefdhbcjecaddghdiachadaeceafhbbdidhecjgdbhfjehbfbdbdjegabcjgfgjcgeacaccgddiidagggfbjifgddgaiidahegebdiegahgighbacdjjecgbjaajjdffdecbbgagcejdfbbfgchjgaijcggiiejgbigjjefhdfibdggadahjcehechfefihfgedeidigdbigjjeaejffhafcbdifbhfjheefgjghbfbjgjfcddagffgeidhdjffdgibgdaagegficfigghhdjafacicjcjjhfbdcffdcbcfjhhdfcjjjjeadhejjciabdgggjdbfjcadegebdbdgejgfcgfffgaadfififbaaefigejefcacjbfeijicffjgiadfcjedggjefafhbjjhaiibajhefejbbchccgdjjcbdidiccfdhggcgddghgbbcfgfjgdghadcdgiihjhchhadjgjffdeffbbgafcifibdccedfeecebhfccdedhgcjgfchadajiefcebdibaehdihbieaagchfgggfjbgfjcbhghebddaaajefjafafheejcdhajijihacgfbbiibjggagfecibhebjbcjbdjbafbjffaegeidadhffcjdbfbfddihhjehiehcgjhhgagacdagfeidfdggbgibgcigcfgheggbifajjccdcfdebhgjbhbcfacdbicgfjfhgaegbceajidhbaagcjgfieachigddfdcjeiajibehcihedahdieijgafddhgidfigaehjigdbedecifffeddehfiajijagajbhcafhjebdabjjhfahfbccbhecaaiibgeibcbacdjfjaababidghcaecfajdhbcaebgagdahcdhcefecjchgaccbfjagbbhgjahifijeachcahgdagecbdicedihadbbcghfgaebgibaggcajhjiagicfdbfaieacdhbdfjcagcdjeeejecfdicfahijcafhecffjcjjciddidiejccfdddjaejjhaefagdiifbgdibdffdgccgjhgbidjdffdgggfbheeagageejfidcjahdbjbefgbabfhggjicbbhfbhgidbceggifeddafcjfjeajdfadhacadfacifdcccebfeijedjjaiadhdjgdaecgajhhabgfbggajbhihbjjgbifaefgchjedgehecejdafajjbefahjcbgcheheageiefddhaehecfhfdehgcegjeggafagejecegadfggggfjacjiadbagaddeecjchifececcfhagjhjfjggjfdhgifdebdiidjgdgdahbbeigecdfdgjdhjjfjgdiidcjbcgjjdcgggeefehgiedejafhedhibaibcceghbdijjcgdcbeifificchbghbddhcggjccafcdgajidhiaeacgigfcgddhjhdcaaajbedbidhaibiefdighjdhebbddhgdiiacbghegbeejchiefcieghjehfheejgacgcafhcbcgfbaejhfcefihjfediejbaadigdjgcdbdjjjijieceeijhhicegfgecjdgaabacahgbjccidbcjigfejbjabhaiibfhfcjagiijhaddbfhcfbajigaijbdcjcchjbifijegbcfagedacihiehchbihedbieihejfjafghiabeibecejaifgigaieeaidgicbehagbchdgijffiehjbgacaiaeaeabfdbjbebddaefbacfjbfcbdcgjhddiebijcaafiejjgeebagchhagihhaifbjgicgahgdbjijffdccbajagcjhgegjbcgieggjhfhecjiejihhhdecgahbbbbaigafcijgdceghjebiaceaehjjigicaebaecjiidchhdiaiaeedhggegadidechaebcifhbbgeeaghbeehiagcfbcchhigcfaicjgaaiiabdjebbjbiacbdegcdeggahjdhefchhebefgjaaffdjbcahjeaceiigbgagjidbfbefbcbejfifdchfebiagcebcjfcdhaaibabcjhaghachgbbcjgdjcbdahiiccfajdjgfehdhfcfabfhheihbdajeidjagchjhfccbaeiccehjaddidicecjccaaeddcbhiiiiiieiaaecbfiigifdjbafiicbicahdcajjcicfcgbadfeffdidjfiehjhghacjjjeaafhfgebfhfghiacabffeabcfdacjbgggjdcddcgccbjhcjefbcgjfjbejhegjgadihajfeijjbbbbecciijabafaafeebgdibhjhgjgbbbibcajghaeadjfbjefcgbhfchdfeicjghgdcdfjffjjfidegeihdjfefbiadehefbcbjbaeeaiigjbfcdffghcaiffbidadfbdcjdgeebjhfahfaeidciabbiedggfeefbgjejigbbffgedfbiddcijeaieedaeadbcajdjghjggccdhbcchabhaghiibicdhgagbjdfacejgafebhcchdahgjfaebbgdjhbfbhgjibgigcdbaajffihjbdgfeajgcgcfbbhfhadgabgijihgcaiihgejbedjiadjeecfcfdchbagbihhdghjbbccjjcahbfhehfdecddbggcficdhbjedidddajedfibaebghccgfgecgdhiihbehadhcihjicbgjhjihdiachghacgfgibcejdbdicafjffhcbeibjdagddfihdefheaieghggfjihiedjacageehaaicbjjfdgeededfgbbaaeecgfbibihaicchgaddedbijjgahabdbhjffadjegdgdjhfdiedeejiicjdhcbdejgijhfigjhaaacaegdbfgdccbgcbbehbdccjghbeebgbidjfbfcjhidjbciddgjchgegdgjgbegdagifhhfcjdddffjgdgbehcdiadgfccbfeiagbfccbedijfhdfgddccdgbjbciaijaiebfcicgeibfhgbfaeheaccgfbidaehacabgjchaeeefcdifjicdabdadgcejedgcgbegifgfeagjbbhgihiefjbhejihjcbgebjhdcdciagghihfbjhhjaegfcgihdcgideaiccajcichibjgabhiahgebegfcgebfcicjebfjajdchagcaehccicahbhcaidhfbiajfbgaiceigbddddajifcfdhijgahfajcghjjgiicdjejbjebchhacddadfhijcjafbjdcjjieiicghgadiddfihdgdghhccjddhedjfbdceffgddddigcehghidibfgacgjjjiidhgbebididbagjgbijfbefefbceedeghidhchbbcfdihfcjhggebejfffeaejigbgjhfbhfeiidcjehjhcebbdghidajijcidfjcgiagieicabgdbdbafagcebbjbdfcebfeejchhbaaaabihgececibiaddfgeahidfggbgggghechidjfcadccgegabaeecifadchbajcbjbfcjjdajgagbiehihafjdghchhdbfbgbhahcigacjbgbbdgigehfgahfjafcfabbeaagbebfdfjjagdgddbfgbgeihfjaibhjhdbjhdbhjabbedhcghffbejcbiaegdhcijbbecghiebaiigjccfghagagaicbabdgbfgbihchgjagjceaaicidcgcdbdcigdagdcchgdbcbfhdbbegjebceafjggajfdjcbjeeidbagicfihfgjhcgaafacjhedcjbfejhjcjfgjjcahbfhhbbheabjgghccjigcfdhbhdbbachbhebjecacfacagfehchjdajehajefejjfafjdhcfbcbacfibjbfjbdbieiagbbdeecdjfgigeabhjdiggiadjdbdjacbdhcdccbbgcffifijafdchhiiaiadcjcjifjdeifdeefaghbjjajhjabadffjjchjbjccdihfeiieeahefbggdfcagcjbccbcchbdfejfdbcifhfgdjihjagcgdfeiaicfcfjbheheeabifbeeffjgaafbbibgaehggehdifaeefafaehgjbfadgaebbaadigijiabjgjehbabicebbijabggdceefgjchaeaccdebfhbjbgaeeggidigdbggigacchgceihjdiihbcjcgacejbigcajbidjciaeifjfchhchdcehijdcecfbgdbeidhbbabjfcfaahhahihdebjahbfghejbfdhihdaiijbdfbadfeehafgfddjacdfcgafadiiaacdcadfecedbfeajfdjhjcidbeidhedgjjejiiibcfbacddebidgbcdijbadbagjaccdcfejcchebheebhihajbchigdcajjgdihfjgbfbdjjfihhgejchajdebfiebdihhchfjcdecdgejahfggdjgdghffcfhhbcjcbcbjbaicdbfdejddidbafdfdiajgcjcedfijbgiagjhgdjjhjdfjjacdejccdebgbjchfajdbafjedjjbbijfjgghdfffbccgbbaiejfcgfgaaajjjjhhaighigdhecbfhgigfcdgigchebddiabecfbffhifdhjgfdjhdfhafadcbgaefacebcafbdbajgbcgjjjcbgbhiidhbaaegbjcfbgjecjjffadjjchidjgghbgfddbacgggdiifbeccgdjhgigicjejhcbbeefchhdhddaefjcjijfjcgdejjjadfbhjbbffdbfifghijfiaiicbebiebgdgbijhgibabiegcacgigecgcachihdaehdjhgcgaidgfcijecjjbddjafheabbjaahegfefedhhicccahhhcegeebgajfeiebcgegebdccjcehccfdceciececbgdbajcaaeajacchcdcgcffdeacjeacgdaebaabhfdhcceaaaccdhhcijcfajihhfefjcgibccicdcffibiiecacfcbcegeiidjbbcjbggidhaahdhcdhjaghbjgigjbcbfjhgffdgidjihjjieghjcdjciefaahbigiaibdeegfcabgcbhhafebdebdeeejgcgibfcgjgejgfbgdfgdffechedjjafibafadcaadfefbhfjceghfdbaieajaeeejadiffhgegddfgiafhjbjjgihchghffbijcfjiiffebeijadcbegdjabjjgacdigchacahigacfffjhdjiacdfffbdfjejjiaibdedhehiifbhebgfcaaebfifhiiddfcjbjjgigejjjdidcggihdejghbeedhegijhhbdheabgdffagedfffjefafjhafffbadgcbgacgiibijajdcddghbgiejcijaicechgfeaicbhcejahdgdjhfabbfhfhbddajejbjhfbaddbigbjajifdcbfdjfjifdgjcjhbjciiceagfjfeaijgbifahcjjbhfgfagdjjbdjbghgceahafjajijgjdafbfdjdaibbiajfeifcdgefhiedeeiddicejfaffeehbagbgccgefijgfefcjadijjijhdacjdfdhjfiaicicghjfbeaiaifgaicajebhijbhihjfgigabdcghhbgddfhhgiididcaaihgdfjhcdfjfdegibeeeaaddcdcghadhadaifdeahjbadejbhejcadeddhdfebbjdeegegfcccfcfdejehfcgfidhdejiacbfabgbiifdgfcbigdcbghigbbeicdhbehfaijhaiaccfbeafdacgbafjgbiebgjchagbcdjcajdjifbdedcjhgdbhagecafabdgjibhjeebceejjdejdcgdajjfaibdehhgeaijihdjiegafeehhbjhhjccijefcchhihcfhcdahdhfbabijhehbhdbebbjcdeedcjcgacjhdahiabdhjgfgfddiedbdddbgfjcddjfiiebadfddiacjgjdedigifcbaadafafjccdabfcjjhfjjdbdgcehadbjjcjdihbjejcaihgjghibjabfccbdbjcahdihebcddjhgjcajedfhgbgccajiddbhbeiajdhehgejcifdbccidbbdcibajdehfjcihjchhafieafjciafjhajeghdhdfghfihbjijhjjeghhebacbdeafjffaieedaeifdbhaagacijidchdbhbdgcjihebagfehccjchdhdhhfgjcjciibedfhigbhdebcehahbeddhhdbfbhcbcdacgdbibhhdchggbefaajfcfjbegidafijfdigffebhgajfadeifhffeedcfecidadgahbdbbghejdehbcbgdacegdjjjccjgaefejcgiijbcbidhbeicgdfifeidheiijicdacfjhbfafejbjgajdfgejiihjafgejgdeafbcggihddihjjideadhabcaeggicgdjafefabaifchabjabagbjfccgiajgihigaffhaicbhfhaibiejhccacfdjchhgdgedgjbhccbaefhfgjgdcijjdbbdjhbjgiihdiaabjdefjaffjhabgiffiaaecgdeabihgfjfdgiecejfcbcbhhgecdaaeieicefhbejejijgegfchbiihgbidhcdceiddbbbhccgdaefieicdagafcbecjefgagijdbaeifjgifcdbhachiedccggjgiigdiaghbechedbhhfdbjddeegjfhejdfdccicabgajegbfdefjgjhehbbgdefcafihidcigdcggjehfeidaafaahfjjifaccajgijjaegbcebhchgidijjdgahbfbcjaihajcegfdcagfdiefijhhjbciiiihhcgefhgfgddfifdfjaebcigaidaidfgidigghfdafjfeaahfjhcjdcciddjhigigheecjaagciefccbgeehgfihiajjjfgfebfgibdjjjfagfgfbfbggfdjhajbdhaffhiddcacegdcghdbcjhehbjjbgdhbhahaahjbgiiaeihafdhbhbadhigdeahagidibabgjbhfifehebbbfigihbfgihfehcjeabigcgeaabifedhhhaeeejhhjdfhedicjhigfaecccgcdacbchgjcfgdfhiaijchdejjgjbeaigbjdihjabadfjhfcfhigfeaabegijefchiedcifadghefggaajcbifiaaeccaejhcggeeieeaieagccagfdgajeaechajacaahjebgiebjdgbeejccafbbjabeihbefbjhegaccbabebejabeaadafcbhbddaedaajjchechebbcdegbebfedbbfdiddeieggjbfejdjhfbcefhfhfhcgecajghefbaeghaebabjhfajaidfidcegaibhjfadedfgdjjadfhgdhbbgiffjghghjfgcjejjfchgbaaihcaagedifibjhcgghfcdaggabhjdafgjhcicjjaabcbadjjbbiahbjcggafhciechbgdhgdebdacacebcjdifighechacabhdjbafhdefiajdbidhabfdcfgdcdacgjfhdggbecebgejgdajcafhjgcfdafijgbfdgbichedgffiabhdebaejfbbiiaecgdhbbeiddhdcccagdbdefffccdibbejgabhdfeibaijdabaabaaiebfhajcjbgbbfbgigajieacfijgejjigccfbiheiicaibhehjibbbccebiceiachcjhdeahcjghffbeigdgeeababbfidahdcajgjiijbhaiaaahedechccdgddabbaiagcdihdjeicggbibdhjhiggbdddchigjfgghccbbfjhgeeghjajggdadjjbegdeafebfcaibegccgcfjdfjfgebfihjfijgihejgficbieaeaadgcigiehbbcgjfebidcachgbgjhgfbfchjeghaajcadeedebjadjiijggbdhfjgdcaibgigjdfdjdijacbegihjaahhaibfafceggiabeebefebcjccidaeegfcgbjdidbgbbeigdhcadficcaaigdahccfagcdchaaeafeaehghcjacebdcgjfadjcagdhheechhjbiedcddfdhfbeicjcheiejdhhajfbghbehghcjiiadjcfcdbihifghbadibbjeafcibfdeggaeegedifagjgdcfaecighfhejfgbgggeijjcfbicebbjbchhciigchfjcidbdcjaedhijfidechihddddiddeaecahgiicedchbgdbhjfdcejcgigdbfdeaifgihgdeaeicacfabdcibbhfebiaeaehgcgfjcffdcffaceccdcfaaajjegdiddidacfedgfhcjhjgaghddhfcgacffbcgegibahjeigfiajcdaggigbegficdhgaibfcchihadjdbgcbhabeeeccadaheiegfccbgdehjchcfacdhgigheaiggbjejacghaiiibfcghaigbfbebjjjhdjjgihedhhacfafehicfgbcjhagcfgjdeaehjbjhhfbajdfdaahieifbjajegbabhdfaddjaafbfiifggfjabifhhbjcgeegcgcdhjcjchggiajfbcjajafcedgccbighcigcfjfhgadbiijegfhjbgacjdfecibdjheeaidiahjjfeiehejjdbajcdaddcgdhdbchdhbecdfcfjdbhejiejggbdiibfhigaeiiidjidajfacidedfjdjhbagcaifchedeebajeajjjffcagceghhhicidgecdhgheedfdagcifjfecbddhbbgdcbeigaaeihhidhjfhcjeefigjcaegbeefjgfgaafgbdbihahfgddhgegehagaedjdccdccfaiabedjhjdcbdhebjfhcddaciajjbbcffedechghedhaeifedicjhcbddbgabafiijjdichddgjheejgfcaaagejjgfdibjjbgbffjgcbjihcbdjbbjcjcebeagdhgcjdheeaficeccjfgedjdfhjdfdhhegbjeijaijgbbaadfchcjdgijfbjbfghhiighebacahghecefibagjgdiejbcheagcdfifiahachifffabadfddidjdfchifghbjijbacgicfbiecejjciifhdegdjaajaajcgajfhfgddcdhdjiibdbifhhcfjjgadeajbibdihfeieaeabaeachcjiiiaeedgcfjihecabdiggcicdabidciciehfjeicdfjgjgiehfebegjjegchccddhhgjjidjahdbehjfjfihacgajecbfbjfjjdfjhejgdfdjcdaeidihfcccbfidefagcafdijhbfgdhdjccahhejcfjdjbbhgadajdbjijjijhhgigaejgjigichhdahfhaiiagibfjhijdfjeceabfebeegadidcjhgedcadffbfdjhhcegfjibfbfjaibfjfigijfidijjjaajfigdeiehhfahgggiaibiedjajcdcgjdaachiefciddjdfdaiagbdeejhfebicgdibjhbdiicbbcifjeebgihjgigicdhgijdchjgicfbdcccdihffhchgccggjgeegfagfaiabfdhedhgecgdehajfaijbgbdcfehegjedbaehjfagcbgjjjedjcgjidjdcdgfjiiefhijgcdhjdfhebifafdedbifieiigjabcijbgibghfibefedibffijccedabjeaiddffddghihaaiideeaegcijegdgbbahajaiebafbfcgggajecdjbfbbahacabjbchbbhahahdbbjeeediegdgjbeiehbchfaeieajfbgciiaeedefffgcfgcdebgeidgjhaadgddcjjgabafddfbdfedgiedbechihfbieccbcjdebdeafccfcjegchbefgcgjjeggfjfahfdhddheccbbifcbbacaajacjjfajfbhhgaacafafbfdbiibhhfghegdigjchjdejefcbgjebaacddehijiaahifdjhajiebhcbhaifdhbhbbeidgfhcajchhidaabjfegbbfdeajhbcjehebjibgeidccfeidhfgjbijgcdedbjbhbfggacdiaijagfdjcjgjgcbfcdegidfijbfaejjbjdijjfbffgdhdhbfbeijjjefggebafhjddfiibbhiejegieaddijfbchadggacgjdafccidhgiachjebejgaiidgecjeicgbadjgahbicieeeghdeaaggifbcdaegcieagibefhgbbbajeebdjhbdeacbhgdaajdbbadgcgaaegbebiidbdbaccjdajffabbajfccidceadadgbajhcehhdaabechdegjcaebhhefbdhajbcjefjbfcehgbiiabhidcfceddefjheefdaajgjhggcgebfdifcejeajheegbddehefjdfagbggeejcaejjifffbhjibdeacbefdfibiihfhjgeaidabdchhgjeebbihacihfgidgbcaiaihecffjafachebicgghebgijdbiiabgcacejcgfhjiaihijjabfgfcaiagjdfacaghedbajechjfbieedhaagaciiagddcgbbbhfjbbhhechiccjacfdeahehffihfdbfefheeabcdejbhgibgcaccfeajdfchjejdbagdidagaijcbdbaiabbahjhehaieaffffefbgdjideciegihjhhihfgibfjjeehjejafbgggejcffcaebahhcfiaffagghecjhhajceiejiaegahgbghbdhfdcbgjjhhjagidjhgjdibjbdbbgcjeadaijcjcgghfbedaghfaaejijbhjchebbjieccjcbdedjdfjeccbecfijhehhdbigacbhbaecffiaigjeggfifdhjdacfeahgbichcbjaaifbeafahgfccecdahifcahfhacaifhgaccjgjcfehjcheaaffggcjaabccaicbejijghfaigehhcegfjibdahjdfffcjbebafahgibhjdahfgcebcefigddcaabfdibefegfaffgbhjgbcdjabhbhdjdjfcjchgajaegdjhaibchabehffjigdheabcdgcibhcahfedhbbjfbbebchhjjafedcidbaegdahhdhebgdbadcehbbhbjgagagagjgcffhiifgcghiihgjjighgaagbgghedbaediheiafjbccfghdbfejfaafbiahbhjacdbjeceiaajeahijcijbeaaghgeeaehajdiibhdicdahgjiibjbbhigaifaddidhbceedhccefjdabhaeehjjcgebfbbgfaiggejeedbjjbjchfiaaggfbdfjjeggaffjbcbffgdehdjgbefgffhbdgjbbigecjiiffjfdihcjadigbeecjabaaeaafjfbafbbdaafdbehjcjhadeeicefcabdadjeajhajdjceegedeiigjbbadbfiifajeaahadhifihejbgggigaidjchdcdcdcadjhicfccbdeebhajaiijfcgdhjdieacaiehigbafedifjjhifijhcjgjgighjabbchcfagjiabbjehdijbgaifhihhfgdadbcfafgifcdbaghceegdigdbcahcbcebbbchijgbjdaaeghdgfgihidchjcbijafdcdegahjgceahfbibdfbecfiejjbidacabcjcjhhfbbgfciejebjgigbgjhibfbjchgbbiijcbgcaacajhehhejfdhjafejfgeicacggjhefhhfhgjchfgjhheehieiibfhjfadifbahgfjdjfgcbcabihajbihebijhidfbbdbjhegfjbbcicihffdbbecedhabcbcggdgddaecghggjgbaaehdfjcejheigfhhfdcjccjcacfjfhighfifbefbfiiajfbaijhiigfabahejbcdbfdggaehgjdedjaaggeafffbcfdfghejiacddfhdebdedgdiifjigaijaddhhihjcjjghjcfgfafhhfciahchbeihcjjbiddgcdhifacjefehehecbeffgfbegfahggfggdcigdadgjcbahiehcchghgejigggjbfchbjjhfjbhhfhhgcacfgcbbaabddejjjaahdjjichbdebbebbeebjhbfdebdacjdebhidedebechaiagddedihdeacidgbceecacaicbibjeghifhgbabgggedhgjcficggfdfjfbidbbcahifcaaiadacidfbeeciebfdaabfchbdhdabaifgidbaachhdjdgeeeiidifiachhefgijhfhdaefeeihbfbejeiifgabdcdiigjbjafdedgiigbciffdfdddeeefabhbbibhjfhfgfdhicedchggbebihefjggjcfffjjgeadhehcbghfehcffhjdabiciaffaieciaajjieeddhgiiafghfjgacgagfbjbiigajieahgjidefcgijdbbchhbgbcdcbdcjhehjeddhghedgdeegahcdhbedegechaheijbcijfaghfehhjhaahhedabdacbeegcdideehhdhehheebfccbibebfeaifhaddbbdfgcgafbdhiccabbbadcegbaecjfajagjbehbbcbfjfdibabidjdhbhbaiagjejfbjgajhfdiihbfdgcbedidfiejaiejjdffggdfjdfehicdjgehbdajgfgjeaghjifjfihbhcgcjfgcfdiecjebhhajibhcgcdbgdfdecdfjbjicccdiadecefdadeebccebhhigbbgbidbdiieddeajegcighagcacicdbcehbieeaaadibgbjcccihbgdjhdfeafheibbjccbdjjjgegbjjfgdjedaehbdjejaefecjjcebcbbhiggfiadbcddgibagjighagcacidbegbegfbfbiichghfgajedabcjacecaffjgfdcfhhjifjdcggfgeijfcfbfjggjifhabahdbaefadaafbhiefajehgihfcdadhghgcddcehfhhcceiighfgijahhgdcegicddfgihjcehjffbaihgifbfbgjcdggghbciggcajghgdhaijgcbahbdidhhaejiecidibgdjjaedhgiaafbfebcfdebbbihahacabbiaibafgcbafdbdhacaadecgagbejfjdahbacghjjcbfeefhjdfcaijajjhgbfgfbifggjhbbbjaijahddicfecgfjjhifacagfcdebbfhfffbjecgibdchbjbbfiaheaihcddfibbiadgbjgccjcaaidbgjeejabbddgejbgbgdbhagdhedabjhcdjjbebedgfdbcidhebjjeifaieggaghjadhgcjjagjidajhbjjjajejgbhhhbjeeadgcbiaggbcgdfcdghcedbdfbfdbbegbfejbigaffbegjgjcbcadiacffhhgeiafcideffidafifgihifhddfagbahibifejbdheebbgbadcaebdcjahgbefcahedijfgjbafabejaahgfbhdbddedjhcgjhdefdbhghhgjhiaafhhgaafcccdiegjdaehfbhbieacjdacfjehhcbaagcfdcbhgegjcififibiegaegefjfgdaabeehgihhdejeigbcibcjhcbjhgchdaibjcjediiaggbfejaijfjdagfhcgecgfgdbfgbgafbefaicbiaaefdgehdebffgeiagjdhhehjgdaidfbibaefbgjabdfifcjbagefeccfggheidhjjbjeecahbebibiceceahdabhjgfcfbcabahbaiiihbjbahcjcigegdicejfihfijadifbfjdcdhgffejgbehheiafefgfadccieajjebcejaeafiiidfaiehbgfcaihbaggabghghifgeabafdiieajejhccdejecjhcabgbaijfafbefccebhibgchiddijdfgbgedjhcejihabbcfeegbifedejgehjbgejajeghiifjaciceijajfdgidadibhcgijdhbicgibfjagefhidchihfbhbgjadfajgjdjgghdiihddbehdjfhgedeeebdjcgccggcdggaidibgdaccaafbdcibbcgfadicccgjibahcbaijcbgfcfedajeaiieaffhcigebigffheefeiifjjffbcbdhfghjagaaabieaebdjhabhajdjejhhdgffddbhbeecdehciiicdbacccgchbejbehidcbijiajgjjbfacifjgegeededffhfdjgaicdbgebiieeeafiggejjbfadadcbajecjcjhdcjdifdeadadbcgiicghjdiajjccbdaffhbgcffachcbfhhifbgcgbhjejbeabbhijfbjbbfhchicaagfjcadfaegahgbidcfefifeaahgidahiafbcjicfjejjdgedjgfhebbddahijhihdcdadcagifjgbdcahfgaadcciegfiedgbdfchaeafifhfejjjehgjhibejifighhcajebjgiegfcfggdjfjbbhcdcgcdjbjegebgadfegeceefeebdaagbjahbdddfddbgbhdbcfffighcjfdifgbaeeecfcadbgjidcfefdbeiagdgbdjfeeagdjheafddidadecejbgfgcfcjehgjjaccfaadchbbeidjfdgccicfdeihhdgajgbfdfgijdeeiajdafahijiebaeifgadfheidbcjihcccdgbaifaaibgjhaeeccjfhjcjccidcghbijcbfagfffchgcjdijdaabcfafjeceidbgfhdccijgbihigifccehhjabaegjchcdjjcgjdegiigcgfjgdfebcjbhagcjahdgbcebifaibccfedhffdgfceieiaiiiadagbiegiifaegjijbahhbcediibjchbjafeidbihcdfadfhefajhfibdiafdaaigbiefjfejjabgaffjcghibbjejddaabgfjceadchjddgabdgajdidaaigbeahibdfbcjaacdegahhbdadfbaagheebjbeibhaegbgigefghjaafdhiibgdcjecdbhiebfidifaejajehbcdcgfijaghbccegccbigddadfhjibcieghgghhjjgebdeabachifcffbjjjbjibaghgghjghjhebfbadfbeaafjciagceibdfbfaffhjgjgfahggabcjaicgbcadfghahccgceigajijadfeeijeeebdhcdjcgacgfagdefbdaijjajeccjgajbdeidbhgjfgcijfagedbfdfcdhcjiecfdfaigeiacbegjbaajhchdhfgjgigfieabedicjgbcgihddadaggajjcccdbijgbbgijicebibfefjfiegabegjcbhfjgeeiiebdeeibjgaggacgicacgceaejggfghgdedbacdhfhgfcjbaagggijiejhchbbdbbcgaagbifhehgbcijbdiehjgfbihfhbbhagdjgiicbifbjiagaaehgacebggjfajefjaeghbbbbiajiaifhehgegajccehejehcbieeicibfjcffabbiahijahgjddggihjbbeacfaebjhfcgichiiahhhbehjacejjfibaacdebjfejecgbeejabjegichgjbbafacbdicfgjfdagjggffieiffccbhidbgiefghabfhehddciadgaefghhbjacdifiejjbheecgjebacjhjgfedibdajaajbccdebcfcieaeiabaigbiefeiijhcdheiahcibggjijefgidghfejccgfijdgcjidgabahdagagbjaefbfechhcaahfbgdhjiadhbahdeaaafjaegjfjhhhjbjgeigdajcgddbfdefiedhiiacdhfafjefgjifhafbgaffhhfeiaebfbbbaedihigcddibegcghcjhihidjfacaihjheiabfiddbhcjacfieggiifgafbfddgfgahghdbacifgfdbfhjfgffbhhhecgcjbiijdddbdiadceiibgjdefchiadgggefcfjecdcieghhecibcfbidchgbihgbddcahdjjahdgcjgfeieabgbcgjaiiaecabbcfjibeghijcebjcadiidjeadhgfdfbfehdbfgfdcagcfbehffafifidbebjjehbgcjhbjdeadcecffhieiibgdceffjjcbdijidfabdhjbafhghaaeifdcjfaaidiaihdefbihbigfcbciegdefecgfegbcgaidcbacadfdabcigbbeebfaeeighgaiheaghjfjaabfjfidhadihfhdgfejhcedgdadfhhjaichabefcbgaihcijjdbijaieiiccdcfcadjhaffhhedafbbbfjdfjadjcbcbehifihhiggdijbbjjafbhhggfjjaijgajecheahjefcbifcheebfdfejeejcgbcagcbebjgijcciicgbcifdiaheiidbggeabjegegahfciecjjibajigfefgceajeffdfbehdbdhbjibjfjacccahebhbhebajfeibeajcehhbjjbbijjdcejbdefeaeaehifbhjhbaacgcfibdaedcjhbjjcbdibaigebcgjieicgdcfgjcdchgbggjidbcbbebcidbbiffjhbaigbdbfheajdhjaachigcabiaejfficfhieiiigeagacgfahchiaegbdbedhcfhihhjbgcdfcihdfajgdbjjhfbjahajedcaccjfdidecgddfeiegfecdfjejgfefdiefhfbfaggedafiibiecfdgcfcjdaiidigfaiaedgehajediagceihjfdcjedhabaiadjjcfcbbcgcddccjiaicdcehjacjjdhjdfabejebcaejgcdjdhcghbecbgdabeabfdeadheigddecaedgeeigdeacdjbjcacacggcjjhceggjcbhacfjhdibijccbibgggcbcfbijbdahcgiihgjgeeahhcbcagifhegidfhebiehdfbejefgejcccgadgabhggibfbfedjhbficchaggcggfcdccedjcejhbfdicejcagdedfjfegdahagaifebdeccjdfdigjcbbcgifdigjbjdifjghfggfadgbiecdfhbibbfdbgifijeabgihbegbefciicifgcajgibbbfiahbbgjjhigjjcahhcgdijaechjhjiejhjfdaefjiciahhjjhgbfjafghgjbebahiiefceaejchbjfbcfcfigfihgcjhabieafbdbgcfbidbccigcbeeijfcfjjbbdjgdfbidccgfcfcibjbdcfhgdcaacgebajefjfeifdbiebebiiahajjhdcfafghbajcbbeacaifg",
			word2:   "fcbegdgcfeaeecbbgaddbcgihahgieeeggghijjgihhjdigjccbhjdfagbjdjdcfbaihagbibcfegdgdjeabfdjbjibdghhdifdadbiijdeiaefjeefgggeejehehaejgaebcdjedgbeehgidjdacjhgdcfgbcjbiaihdbgbjjjchdabbjdiejcddegaidjjhacagdafiabeebcedgdchiacgfffjbbjgcjbhibedgbaheeiegiehjfeicgeibdhcbibfeacafeieaiejigeabjjhjacfcgfiafhbcjfeggehedhbfeadddddifbagcicfajheafchjbbbhadjbjejdfddhacabhffdaadfgbcdiiddafiiceahajfihceebjfgifbcgaigheedicfihgfhaedcgbcjbcgjcbibceeechefcbcfhijegfgfgabhcbidadeabhhjhbcagfaagfjcbffiehbcheafdfjgcfghaefhjaiieidgbbahacfiedefgaagehchfgiebcbgdadjbghdbdfgffjciichiejeifaeebeigbcafijhcbjafiadjafhcijfeidibdgdcadhfedafdebhdibejjbhaiidiedaddgbddbceadddfecdbjffccbagabgjbajhhidiiajgbhfifjhfjdbighhdbgcjjfgaiahbjecceeedfgdgecagieijcfgjdbegaefbaieedafigigiaedhgbghidhbebafjhiibhfdbhgjabfaeidiffdeedcbdjiadgbcaiehgggjegfbdeefehdffbebiadagdbbbadegiafjgjfgegjadjdfheejahdbibiffabidddahcecbefcfggdcbijefidgaefhjbffihgfgaaefcagghjifaafbicjbahiheiagceiichhgjfbiiehjhfcbaihbbadaiihhcbaiaaiagfidhegijhiagadggfaeediddcdaabaaddiihhfhheccecehhhahjdhbbdfhdifcgjbijhaihddcefijhhdebibjdabebhjcjfiaiddcgfdhjhhaacciieficegefeebiggfdabfajhcjdhddeifcaaecjbhjfbabdcfbifceahfaefibghdabifafibaceagddhahbjjbdcehjafgehejcdjcgadcgigjbiafadbgifehdffdbceehjifddcehfchbgebjfdidbidhjgffbifjbcabjhhaajjhjbaddfgaiddehiggfigajghgdafcejjgajacjjggdjgbbbddchehhjjdjeahaiceefbhghiecgbjchhjbcgbbghjedadadcgfbbhcaddgbhhcghgaajeahcccgadbeahihgccgffecgcbbfdifhjjfjccdajbjdidhdejihgchjfefhbgbghifddfigcbbebbcfjbfccjjcheabhhhajghdfbdcbgcafhggjjcdbagbjgdejadhcabjhbbgheicfcdcadgcehcgfdjbcccejfjijccbhfibfbehcecbiiacijheiddjedgghfciedjbfefddhcdggchbghgfjefieijagdfaejhbafdeiccbibieegabbbdjdhcbidhiabjdgfbdebdjheibcgffbhijiagdfagdghaihcibiahhjfgacdgdgajahiibbibcbgcdccaahghdfbcfeeaeaidieijdiajddjgbejcbbaacdefjaifihadfafhfgfhfigjeghfjbjefgehgijgfdbhagjiiddjahgebdbfafghhdhibfefgcehhaaeicgiaegiaicahcdhgjaadccadhjhhbbbdehhecbjechjeacchhdbdfjffihaeaccffcfjfeiifgjifcbejjjffghedabaeaicdigffjiddiicdihiehjifiiiahaiheibfaibfidhihecdiggajadhehgjdgebbggdbaidbgecdhgfhedbibhhidhcddijficggidhghciigfadbefifhihhcfjecgdcehjecfeefhgdebeiibdeehiicbhacjjfjaeciciejjheijcbjcdgajjbcgahcheehggffieihcbbagdejjcjjbebccbjgcfgcagfhieifhbgbahihjcciibiedagdfjbeijbbicabjfadggjgjehjfhbabbdhfhaaihfbgciaafdjgcabhegbgacidjihdgaedfiighjhbabhdfddaafeadhighdbihfcdgadifgfcdgafjdiagebbjhebhcgbifgegeeggjiechgcdcbdafceigaccjajcdhaajdgdfdhjahaacgciedahdifaagedeggcffcedigbaiidcdeifjiihejbegaggabdbjiaeicchhhgcdhbjajgjihcgbjfibbacebdeddfjjcfejbhfefifihjhcfddfihcjbecdjajdhgiggbaeibgaiajbgbhegfdiecacgjgbdfigajbefdjjdgdfgdiefiijjiggegaehgjgjbhdjgfcfigfehdfbaibggeffebcdicihbggfggehjgahhbaffhjjjeeggjjabfgfjjgdfhagbgcjdddcgghhdbgcebifiafhcjjhbiaigjaheffhgbhfhdgjhhhbafdghjgecfihbighddjdiaeficbbiiiebjjjbhfchagecafhcifhicafjgcbgccahghfjfdccgeffgbacgcbfccheedgjbfbedjagcjcbaiefdfbhahdcdbdjdgbghbecciiicjgehdhjihijihhidhjagbcbeaadbidfdbcihagihdjcihfejfdccbbiecdbjhggghaahdchbbbggehjddfcjijhjibbaffjeicfeccdeciaeibbaagciabjafcbcjggcadbgbejdgahiheieabhcggfgehbiiadiifeehegchgbdfaedejihbaibcbidbdedgeajcahihabdajaciigigjafgfjcdgidhbjffbceecjcgbehhejcfdbaicgbiiicbcidaddcbijgefchichbjbjjibfbbbeigbdidahgidfchgbdbiebhdbabgegafajggibbgefbeffeccfaiiceidcjfhbcggabfejjhajaijjajgdcgdibegbifehbjghehjaegddcibchajghfcgfaccdfaidfgbagejfejigcafdggcbjbeffaagigfbcchahaeadcchffbffacefceegefdjdbbjdfjgafiadjieigibfbcccdeedcdfceighifcaeigccaejbfceeiaehdcddgiffcbidghggjacicebaehjhaidehdgeeahbfahcjbjchfghdbjjfaheicfbdcaabfcibhdehdjjjeedfdabafegahdiegcgjaicjigegiaeacdeagdceiiibjbjibeadijfdjibdbegafddjghadhhafcdbigjghddhcaiibdjfjfhegaihdhajcjdjifggjibijjhijbcjcjccbdgbjgehefhgidhffgheagdcgbbdedgegjjjfhchfjhhbfieahcchaagafhcdfjjgicdfhcaihbibfafaahjachjdhigbdfdabdeedfhgjjcjedbfdichjhhjjbcaafaicajffhbbfgiaecfhhidbiaabdehfaechhbjdgihbeabgjghfahcfafjjjciahgieejicbfjgaijgbaigcbhiabfbjibceejfhgbgbdgibgaaedbbfehhdjdeichcdgjjdfechjicdifjjadjbbjcdbbajdbihcihdegbdgehjefhaahgbeiichiaahfbdgdejbagjdaigbadeggijbifeejhcdfbfijhdfjhggabibjdehdaedhjcjcdbadbbffajbhehhfidbgbhfcjceggbebbdfeeaafaibbafhbbefahjjcgiihicffeijeheeiaccbiibeaajhcefbgfeahdjdhaeiahadiacegifgfcccdgejffbfidfjbdfeijbbgcgeaihfcgfgbfdbbafdbhieebhaefieieghicjdfbihdgdfbjahhhcbiiddefhhbafcfgcafdegghgjeghbacjffcchicgebfecajbcgeehgadjijceeifhibiighaeedjccgjcbchaafidcjgjjcgedegdjaeijhijdheihfibgjhhebecghgdiehjbjijjeaahcgddbdbjhfbecaaeihcghhdbbidjcjabhhjfbefjfhaiijcfehihaegabefjjcijgejiifgciieghjicfjhacbcddejgfahebdgghibcedcaeddfcafafddahfgcehhiajedjaidgdccbebjfbijgechdahcdjjgcehjgahgaifhegegccaagebaiigigcijijhggbhadgbcacdbiegdbhhibagghchddfhidhgcehfhieidcfjbfijcjdchcafhddccdhediddgaacgfffdheafhjefbgbcicahedeadgheaaeagfheicfbhgbddfefaifecegebgibeffhffjhfegjeicgajhdheecfgahfabcfabgjajceccigadfbbdhjghcageddcaghcdghjdejifhfbiifgeicidfjfjjaafgaahihcgaahbigfaebffbjcegifdedabbbjhdfdbdbgjjbejhbbhgcdfigeigbeafhgaghfbagdjejcghaiihagfcfiejahdggdihcbceaebfahebdgabcdhhcejaghjhbddhjgddfcbcaiaahcehddcecbbgabjaibbggehfhjghcbgjhdadafifajjfhfbejhieejhecddcahfhafijiagfcfdcegbhefbbbhaeeejdecejbcfhhjeaabbcjfdgchaefgdfeeebaccbgebfehiefdcicjabcgjijcjcgeccfjhgfbchaijjfcij",
			want:    int64(97971163),
			isDebug: false,
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ans := v.f(v.word1, v.word2)
			if !reflect.DeepEqual(ans, v.want) {
				t.Errorf("got %v want %v", ans, v.want)
			}
		})
	}
}

func validSubstringCount(word1 string, word2 string) int64 {
	diff := make([]int, 26)
	for _, c := range word2 {
		diff[c-'a']--
	}
	cnt := 0
	for _, c := range diff {
		if c < 0 {
			cnt++
		}
	}
	var res int64
	l, r := 0, 0
	for l < len(word1) {
		for r < len(word1) && cnt > 0 {
			update(diff, int(word1[r]-'a'), 1, &cnt)
			r++
		}
		if cnt == 0 {
			res += int64(len(word1) - r + 1)
		}
		update(diff, int(word1[l]-'a'), -1, &cnt)
		l++
	}

	return res
}

func update(diff []int, c, add int, cnt *int) {
	diff[c] += add
	if add == 1 && diff[c] == 0 {
		// 表明 diff[c] 由 -1 变为 0
		*cnt--
	} else if add == -1 && diff[c] == -1 {
		// 表明 diff[c] 由 0 变为 -1
		*cnt++
	}
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/count-substrings-that-can-be-rearranged-to-contain-a-string-i/solutions/3037271/tong-ji-zhong-xin-pai-lie-hou-bao-han-li-2kiv/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

/*
方法一：哈希表 + 二分
思路与算法

我们的目标是求解 word1 中有多少子串经过重新排列后存在一个前缀是 word2，也就是说要求解有多少子串包含 word2 中的全部字符。

对于每个 l(1≤l≤n)，找到最小的 r(l≤r≤n)，使得 word1 区间 [l,r] 内包含 word2 的全部字符，可以发现子串 [l,r+1],[l,r+2],⋯,[l,n] 都是满足要求的，计数 n−r+1 。将所有的计数都加起来就是答案。

而找到每个 l 对应的最小的 r 可以使用二分算法，我们提前预处理出 word2 中所有字符的出现次数，再预处理 word1 每个前缀中每种字符的出现次数。因此在二分查找 r 的过程中，可以 O(C) 时间判断是否满足要求（C 是字符数量，此处等于 26），而那个最小的那个满足要求的下标就是我们要找的 r。

由于本方法时间复杂度较高，有些语言可能会超时，建议学习方法二

*/

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/count-substrings-that-can-be-rearranged-to-contain-a-string-i/solutions/3037271/tong-ji-zhong-xin-pai-lie-hou-bao-han-li-2kiv/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func TestXxx2(t *testing.T) {

	for i, v := range []struct {
		f       func(string, string) int64
		word1   string
		word2   string
		want    int64
		isDebug bool
	}{
		{
			f:       validSubstringCountLeetCode2,
			word1:   "bcca",
			word2:   "abc",
			want:    int64(1),
			isDebug: false,
		},
		// {
		// 	f:       validSubstringCount,
		// 	word1:   "abcabc",
		// 	word2:   "abc",
		// 	want:    int64(10),
		// 	isDebug: false,
		// },
		// {
		// 	f:       validSubstringCount,
		// 	word1:   "bbbb",
		// 	word2:   "b",
		// 	want:    int64(10),
		// 	isDebug: false,
		// },
		// {
		// 	f:       validSubstringCount,
		// 	word1:   "dcbdcdccb",
		// 	word2:   "cdd",
		// 	want:    int64(18),
		// 	isDebug: false,
		// },
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ans := v.f(v.word1, v.word2)
			if !reflect.DeepEqual(ans, v.want) {
				t.Errorf("got %v want %v", ans, v.want)
			}
		})
	}
}

func validSubstringCountLeetCode2(word1 string, word2 string) int64 {
	count := make([]int, 26)
	for _, c := range word2 {
		count[c-'a']++
	}
	n := len(word1)
	preCount := make([][]int, n+1)
	for i := range preCount {
		preCount[i] = make([]int, 26)
	}
	for i := 1; i <= n; i++ {
		copy(preCount[i], preCount[i-1])
		preCount[i][word1[i-1]-'a']++
	}
	log.Println("count")
	log.Println(count)
	log.Println("preCount")
	log.Println(preCount)
	var res int64
	for l := 1; l <= n; l++ {
		r := get(l, n+1, preCount, count)
		res += int64(n - r + 1)
	}
	return res
}

func get(l, r int, preCount [][]int, count []int) int {
	border := l
	for l < r {
		m := (l + r) >> 1
		f := true
		for i := 0; i < 26; i++ {
			if preCount[m][i]-preCount[border-1][i] < count[i] {
				f = false
				break
			}
		}
		if f {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
