package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {

	s1 := "kyzjmmhhtgqltvseraagwufcqjihkjzrkrvxgryhgspmnjnaanujadpshqomilawhdwdnxfdkfpbelnzeozhaezueyduvrblommahfucyprpxritjexcfctrkfjgxtffqbwrvteqaoyawnjdopkpqgjkwwffxnmuxuuqisbvkmddngmbcmkyaueqkfalixhxlvdxjozjvapjihjbnmskgyknokxovpifsrqoqdktqvzkhhdjunqgjlcitybdemfssddhxmblerupdmunrpwlkwpuegiesrthomkgklrcrqceulwsfnwvwqpbvstgiofbdszwjpzvfltomdyuttnmtxpfsiaechkncoibymkwueunkdacwazyvyybivfxskqmkigqeuainobzbowyittpzdtpholpirwmiumixjcodjivsbxhgqxwmpmyiexrvfttrrfxpgcgnnnqhaqihaabxdpzeeybugimwlctfohsovxwsdpjznqvzocqkzqwsonlkjswzncqwbgfnekuncpifucwzmfdswdoymvthzswwtpjgisefwcdhdbxfahpnqlwkogkmyybgxfyoffifesvxlrnqrvdtfyibfvvooodkcrxffkxbgtxbjwjiuvjrpttwsstbnworvrfvgndignkmpxgmntlhivrcxwxhudfyuxlausblbugldknvqyqbcgpydwakhmoizbttcgjlefevrtyevkzlpbahfccritmtlnsoadkmxgikoeokibgvmfjwzgdeszaxtzygdsmwokozacbbvnpkhwbnvrsndwrqcfsogshewhmqzkusjyvcradbvnfnziuqjcsubvanhquxcrlnwjjrbddmkqdeoycenmteifdyhuufgvwjgzwvjizqxeaucojskrcrqwefnoopvqsixmkswugjcocysrwotiehginrszyflfxjmfltmrtxslwawugxeunqxdgquqfnuwlnnacpqxkowgzwvcuuesotbduhjttxtmatfjnmtjwcgcsncwjzlbpcjkxgbkykmiawcwcazslrqfesklgbxxoapettuzdgamvusnpmvcgpugkmafhykrzkafzkcmsifbaogkkwiykjqgapaovguhhdydsbtxrahaxmcgqepgrhhngysqfsphnhwpclhwddkfddxshnhtzlyybwibkcqdwxrarbctvxgvglpnmwxdhghqmfqmvdwjvnjfrujlitdtruldmpsfddipcshtjcykfdgiydiadsnfdupcchrihwxhalogqrryyfowbmsmlfqzhjiarzknyabuoquagtnikokxoadevrpbxfmfweiuwesaymdcibzgjrogxulrfbtvlgtenlczcfbqbkdvkzrnjnwkrcdttaicuwdqghnsyxancpxynhtalcynuxpolenbqmbuwnzowqthqkmkyujwwktrzvblxadrgyeqapoyexrvnexjyzzxzypzdrcvplmfkpwzckwvpnnrapolsicytccjfrqqbijgmukiwlrcniyxhpjdetdywvwfbhajecmpkvamcziziwtbhicbwldeuqczgwngduyjtvbqfhnabfefmfjvmjkspglwcstxxxbyfttxvqdhirbwklvaezklzcumojlmmyuxltesuotcbsayhyrvehqhdkfnjqfjuqwpszrvowjyzgscvliwrbcvzkqmwciwwtoxxiwtqczvmsmbkoczkpnvakodoznnruvmppaookjivkxdxhzriirmzersxexhyilujgufpyjnqqnrukzvpejdtdbcoxdxvfggsqtcchmkvvqgtyfmwgfjvqpizvfvssoogxjqsrghbmxhrpmwtvtrxfpeyftjjvnhwxqjvfibidgvveecsurrkqfaamiasvbgutooothbfsxiqlaxfctmdrckqvizxtnaxfekiwufltnsjietugjzdsvsubxbmsqpvpgfqgukdngcwcmfkbjkpvzwkjaxtdugmibblhsldsgqjnueqzjuwjtuypsduymlydxyteluogogpcqiduxzwgtrohdalwqipbbvfhrkgukxknlztfbqqbsvdkcqeqofbzgjahkubhzhsflacljjrruewbquvwdluwsfcmtasbjoprfobearkcwhsmffwawztkgqlnhinvzlmljhaezmqqvadpdnsoafdhxyamzfebhufmunzwtxlewgywivjzakhraiqtrmjipxotxcrnupwncrlcutoaobtxjvkxrdmgefqjzbwdgpsxamlogduipgewxorszsnogeitzsvpybxfzyulfntrlzjstodqgorzloqftkytgiessvcwgqistzmtmtrsyupysgbvnfrzoauijjfozshdsbwenyvjdzemhvxljhprirljvpkloynukyuhursicdjwpyvkxsxadrhghftuhodljndycickjqazvunsjvvtnwsbvusoirgvcabpdsakzlonohhkjttamqhyddbulbfzbauixziqcixrwuebciogkdnrtdwipcacgxrniwexbbyvakepkamoymrvpfbadphwbmzvvrlxwcgbpwlugndiksqwoqkmhsyihqgeisxhwaotcdognndfcggxawvdzdhagurdkaudjptnmghnkvfpytpsqkwetudvkebmvdrbthhhngludmefkjfndvvlcttqsmswroyppfzsczkdtgotvmueixkqzxrvkjkafxupwmvxmkixrcfewamlafqndrgrcuyojvpduuejptjhhdzqgvgrkomfeygbznstyetvkpglfgkdwcmwkxnzyuvajaxdykowfgzxqagvgtvsexrogdbcuqutzzkkzyltdzdigavpwvnjryhbmojolvtwxqmywzxeultwgtxctoqhnumltbvnybpkuuwfucndoukgccoudmxerwtcocfupeftrwgtqswnluhqkcpkqslutlewqppxuasvqnytxanoorvldxnccpcczdezfvejqlumedsjrgisiarfcfwaclixseujwsrzycgypmigpjlotogjlibavbjkwjvmsxvqignwzpaooqeponsydoxfvkmlbruoaudmagpbylzwrlgrlkvenxpvptvorxfsyqctksumxsoeofpgrhyvnkggbroohorgrfumugeicqjzfztmxafboezhykkcwibpycudbykrsflztnxxldfnoidcxjeuypsfsxgnmnukkjgeakcbpsdwpsliwmljcpbjwlumsgfunxxyybjfqlwrnatbpobflwdhebqzswgmpbrxntjzeehdifvswjgllflptvjyuieqvaqeliyxtasajeyjkyxygnneqlsixivrerefbfgpngulxewdjpixpckniuhwwobobciequhgkljrnfrlhlncmxggffsugmyepwlyjgncariytuduzinpdlhzxtcwuawehosgtylhtoeyecsdojnotdtcoaztkwoywknqkmvxjvlsthnhnoaktwlkgmnfbmmkwomwkmggfqnztutluogbsbdcpdxzwxksrtwfcwzxqhocagthobcphiguveebxqaoxreqmjuzzazjdsvyoblxiuhnyxwxbwkrvyenwfmrervwbpthplpqsmmxfldwgzncgmnapvglxqnwpkoogzfybgfywrppzsmmqgxefrspxvwfhnsqrazsewfmbhxdfhuxjgztdqbjdakunokkjqksunzztgokvdyyamfuyexlwwljwhyuhbivnqgysadmohalgnsfanoyhtzhjxueqoigrqgapmafbvohsrfymbvwmtnfrhesgxamjltvmzftfkwxzchgcmhkibpayzibssxefgjnwlrnxdjjjxtgezxqprrbdzwisidxuastrthwkvdetvdzaltejxiejnzcadlzxkrxbljgsmrgkkvmdtfqxcgfzrvmewywkwxqlvkiwmhylzccquguyfwquiivthjeavbiwltndavtmrfhojuexdeasvdgbfmtexzcqalhjhvohkbebrwnrdcfdycsibbxsqzqcmebtkjuormxsqxbjamdmeltztjbnfdcwfhthypqyrugogqnthbanrygkjveqrlofqmpbzqlnysfvfcamlngawnedspiuiutpgfxzownidpexndbdekgddkdulcjqleageksjvsvmfrgnlyawrhvhbytiodeiksujyexhzncwpajzjuvgdldeghwiaregonicdgzzgmmgwuzwbexeryflddiavshekhyinuncmkfstkuxckitwpugkdseqpbwjbuonbupahukzdslrbodxagpppdsfwdgoognjdnwpiliziwfpdorntogvjfxndpmrktnuiftunkdqxzoyujpsxzumewgwjheqodakmhmlbwcpjzgwilwznifvdpaptttpyyjjolptcpmgdsjrlmyuxuoncjovxcosnarxuxvsnvwimjqewcgtcsauvpicxldseeanylxdrkbnbkldcnvekvjqeetjdzngjjklzpvzarjhcayzierngohenskwnrkobasrxeyckohlvyndfpbihrzsxmqoamfvzaolbrpxxniwoadxinbuzklgordikvxxtruvkbcuadjrlcoecwvhvbvykfbnhpzhhuvyyptzygaogknuaepmoaspauanknwsiziajhgbiqgroyxlgipzywdkduvpolmcshgkajeekwwdbryhphqlswoefzrgsqkxvzqrrssmbeybjtildsifqoggtxtlzhdglilqqakyijadfpkesnjvfexeekbmczmcyovrijhfyhlhccpgxjsxhncpndsianvumpyxwjgudjcrfbqblyxjdfdprmbwnzwqoblzdjaipuhpwhybwktbytyordtzuccvajqmlcsxrgotesmcyfuwjfruibxsgrdedfqrprpsmfmlfscheayorzpkhwfcwlhkjyvflqvhawelwqvvjtsrcxanykbgotignppkwbsappkwfavlhtjhemjdfwongpbytncqoursmnkygmbtvdjonmrclpqnjmshfvkdkrdjmoaplmcrhvmhrymwgsdiyfxwzvrlngcglmtuobtbziaqecxtigrwbklfsfaulajgyuxrmzuslqclcnvtdtjiuwevmtylgntjfhdediptteodudramfooedumbsvauwaqufjtbacshjmfqqvqqukzhwvljwfliznwisvyiltfwdibgunqcazdjreqcdupzkrajmhdglpeomvnsepumsynfkhpvqeafdeusgvebmxrwcecnfgiayhsjecawienfnuvdrvenvrkxejlyuhoxurooteazruzraowotxfkwcgqshqgmcjudqcxmxqivavpeieycvmscjvrnxpueaevynbqxbjgwvjazrmdqftoqvcsdxwugfneacsilquhhypqbgznjjuqmrdbhxjrohxcxeaoagabusfjimvlsogxzgiocbeikbgbutwwcyivuotsingmveimfcbdsdqfougktbinmygxtribfddodzehjgvowarndrahtqtcapijhbmrrzzqhqpnkiiofgcmnnvnwzprvxobikshuvmlsxnydiwsdigxfkjbwoyjpldivsxumdolrzwoybecwxuimkntnifwguethhvcrdrcroynehscosaqrxcoeknyngwvaghmethvuetqsezaugloybmqmwnmtpsodymeojtvqlmcafogabenjjvjzpcdynbwfxpcqcwqohliptxcnzssavvztbqorxbxtekfztoyfqksejggdgoibllhbgnlkvfigulqwxlbbdsbwyhkvqhyccfwgcplvzowqulqvanqnzejcsgkpjufaprtvkhacenehfsvmlxkzxocokueadgjqiysxwxrpkrtivwhsrqtpgjkgrwhaiqicxrorfbxxuuqcqmqdzckzepkbwdurrwfwbhunumptualwctrqkykfeblyoixgwcfgjavjxiewvgxdkdabhwhohsworlvmxqxtethzzelpyxosnpqdymcmjifpeudmcmoxcnhbatucgleugcehljngkqcsgeaxgyyqeddasdnemliwkdgflplkiivjzpeoxlgtttuevxkdmurzguesiqkgxphgzznnhjqeyczqzrxllfinhsejxlkjfrnjvwnwysotthphpdphiuvhkgeivqtirkhhwmhfqbtjzpvkhicqoefxykmjcipnduvacwasfbtihunawcalclifewnyfgklxhcirfkztmuanhsygaafzavchhlodsdstgjqwlnuqcnixfkuacomcyfyftnhieuefpwbbvikedueeoojxokbjjiefozslugruczbwznsooxuriwvqurztxjgeebdlrmxfqtpfpnxpebywbdmcylzjeflwvwyqvqhxuyzibfbuqqzyukpbzhtfnncpjmjwwzxxlswcfcghujdravthqabzbacpxalubyblmnssctyzzulzeqnqdwggdlgfdcrspdanufjoarkthfzlfyrvfbrmebzmktmgzagofhgumbtbanwvjpvjcmeljynynpxlqymecvkfglepotxttvblxzhfswncrpsuwtskwnrqtqmdjjhsycfvwiipavrefgmjlaskstnvqbaqcfdaydlzlrbegxczbrwqygbqrktqhrjutwdvxvmbmqturfpsmebvbmyllboxrnwqjwsdfufxlzxojhfxqbltlwssodgthlycjpyfhltfaoooujspfgiatzsggznbjrfwyrsokgxmutqpayxjxpxsqvypojurslbdmgvbkgujmrzytbvgfzjehodyldqajntdgmkojsupvgaxyaresyeaqucaxpbsthvxciiqgdsqftyeumpvyornllzolnwvsiuylhedsadwlzqgjtchfusmolrmfelqypudgvahwigpbuzpmsqeockskkibgpnkmlylzuugxhfrxbubwnmmggxhzbhsfzznsssulqgaxfboxcycjmfipvhuurkhyoevphmxbifrfatlbzrhsfrwccvwqtkpusyavjmtcmewxrxgbpbiafjoqgsszxuncznyluggsvpyypofyrvgqpvwjiekhclxberlgxufjfzaaideostfenggbeeuqehjkkxyvwfqhcasmjtsditojadywozdgumrwsyqyhbfpxfcifviikehkycfeojfshryahoobahetobvwztachkhfvldtubbcmousojgeejrceokrhoghknuedkmqicqtlauhgtpuwtzqmgndvzryoqfitfigvrmatcttlbqdbidgwxduhpsertgcowtefsckdjyayxyivqzpziffvozevttdgudmdrtkzldheefevgbvbeqohwzxovvodhlxpsyxirnpltqemnbvptqgylpzossszmgcwmzctairrvxeskkiqnetrhvcpqvuowblrixpiodlbrjrpszdqymxktsyvnbbbbkpixntzaogpzrzahknizpsqmrhxvbotchcbpaupmgosssotfkgnbfeyfjxpfgkdldsxafudhezlwyvjxmmqwczkxcsnhtugkeskidioutovrhofawhgxfugefuuttfzoeewvyokzkbumxihfoqvzqhibqeocsrbhwrgokyjhoglejhqxzoezrwdgftevraxavppemrdytyvmoxhrxzdhnofvajjbwgwiufyczuzpmfttwptadfiwsuxxxwjqbkuxnipbhckpmrmfbatpbqpqhvpjuzyvycdraaljjzzukxxxqbowddhvcbtqazdmciiezcsrwxvjsqmpasvlnalevchnftnknjoqsuomulaporzyupujjfxtooikzkpvmyagzknchifvfwxcgyfvlpxccfrueyxsjwtiqievgtpeemhpykjexxuvatqrudbaocjydbszancpdgkpiwcpgltloanuhvqckuheecaowimpuawhcvwtetyhszafckivvedhineljrqzkgcbrrpjupcaggggkqxgveooukveqhlofcghgqxvznffhqueletsyjzgrzexiobscbwkycerflerdxyjhejbvwqyzanbofuhvqlxwhtbbjdhiktybddppvkpbefcmuvmkpuoonzjrxtielipixefcddokuxplzvtmcykvlondudgwiiaqublpdryjehezingznydddkmnesdfpamatejvvvsgadszspmvzljohvipwnquwpcgdfhxyroyzzttqhcqglqrztzzqkjgynpmalkwczsmtvjndlxunseyuvgwjxurwnhzbgzuhsngmhzuchjivgfctosuxqbchamxvuhjxznzjqkbogrhhxiyugqwzvpmghsmwewbrrbbadzdwlrenpxfcmfoitgidjpikotnwpvlsvazjmdvfatohoazuzktqvwpgbgawkwmveaktdxfuwyupejoqrwbaqgzkcozpdzrtyeuqhgflhxliauiuaaaqgolqqlgadpmvftggxyydzcyadnppztgdmswszzcuxldeoulmanoduevkrcnptggfyoeunnnpxjqhjafzmczgciqeztznmnmbkkwalyliuyiikqchtfffbrqqvuankviydvvrcwwphjvwbzhwiocbgslzjbwdpzptgfxqohuimukiytmgutxdivmcymnvsoodjbrfaqpfkyfbhcpljitkluukosdpvgzioitbmpogitkbaeidwvrrecwxyatscjoviwsrxrbevwafdigpdhebsykeerwetkqyoznxtpntoltvawhwvhxfjqznazgvqfttisaishecexdxxagvixssfqjxritjrawnyyaucckixajkuqtgdohwykgdbvvmicxwajebnbukksigqmdhnmaphakmzsdtnyidocevvmjtggteoayidadxgtztspupahcypzywflirbcvbxgqlgocfqplealpukyelnulozyehrtorveuqgzlflumiijgstlohqhkfwokocaqatqblgqtlrmpgmowxinibpotreegnxubbnjomlnpuuzxlxlcplhlrgudfhmrmwylogkrgtilgxkikoenxoxswbllxqismvxtndkxtshdgstkpeoztcjugtyczgwtztvvhatozonkqyshwdysgltuqvhkxfsyuyorsfvmeyutzaulbgqnjkofnulyubzoobhpytpuxugtmfntjtehcqxevujuqcpfzrogczoaljiywlibhihcocubmsaqfxtxbdmmcpwekhymhcrmnfzxkdfkkuccjndedfjourwazjnlicfamigzycwxglixsegxsnuppxsxqlwuheboygltaurhnptcivnhvxcblgqbeelscpvtweookxmnhpcajgfqvvlmqdqlvqigaojnzbhvfurjliigazhhyglasxfzzubjbcplhckunccmumcmmmiqvveawjfrtvgzfvaiqsptzcfrfjzvdohpgiulxtujvenljaplkmfatjpdsktcwhzswmzdicmtpkavqvzphynzaeqddtbzhutcpxggogygrvgybdcslycwjomlhyxjwkuyptmrzxyrauksxlfzkaijzsblpnimjkdljsmsbkwhvcygpxoyczzcoyphkrapttseejlllrbjxqdsymmwwzcjgkroosoaysdnxiidkvxntdgakdgfbtjjeljudaavdeqryxiouvsiojhgcdabmubztfbbsnfowedmjmuujrahecmdpbhxvyxulepteftcpaotaruwhdagnejklhansmkowtcrdkiwoibdihggdcdwcoghsuhumjapvborjmdxsyvrieqjgcddeqzjirtjhgyznkexhuuykghwccctjsdehyepbnakgsvlkznqqizdoczkzdghrwyhdfabjobkeugrbmwthkkmukmbuyeznnpfposlkrgsdpzbohoezalrtulbgeacfdnbucvohqhbyzcnhdxrgrttdylzvwbkvvglqgmodqittkuvmphtjxjgyqxfoqibhfojexqbrifrkptshdrbmtgdiwkjbamgozaptktpgtwbnskrvuvhcppkcnstkhcgoqydxrcuxorfumchgcelmhsjyjexorylvxhdjkwrrqqpffkztjlcopnqwezihmdmzmhqtvmunfsasxvzfmfohvbcmlxdzbgncwefcdumddvbzibqsqmdreesxqngyvrqtfaucbbywqhnecgwdluihsngakmjwwoepffhijtplmmxqxgeedepwnvjvlqsqybxwaxdqrdsvhapoxcezhesnhtyuqjvbbctacktjsifghbokuzmblzxyiwdeyshwlpliehxzqhpgcbrvywdqbwajcdpieonmqjtgqacxblrkgxgkpybbbabvrmddtgkzjebignkrjxbejbcvwgimgdpegmvnmnaugctjhzapoyqwueucbypwdxvqtzcrbzasvnnraijcvjfkuwaxravnbuhngiaguznaxzzmbjmonllfjrvnzzrivxigxqbakveovhzcfhqmyxzcrglhrasaxinomuhnfcborjfvpcexmltnmedzlremjehxqntwbzjvyvkjodwqijozrxeqebnclfmzepwlndxjnpjlxclufqogwqpuqsbewxiewcueugbmddmmkmjamduhqkxuqipuxdyvlkgnsfpcxeabcgpamrtvbhifbvgzchejlianewbntglynyxvogrrcpgamtudjarkzpuideyzanqqchkocdmmivrkvljpzoidqoztcmihlhmorhrehxtgfhuvzleeqzmursfgdjvjplaatouienxwtlaxhbvjiygzehjdqjaksipibxtcnfwaspsaxaskketoknbtqwqpxlmuhwyfrwytqdoybffdlejamlwmjqmhqgceoisefwmumiunmxqcphsagsdlyflfyucrdszcgreqemkxfbnrkxpfurkvfteboydlmuzrclnguhirvitbivtyaonkfwzytupjmdjoxtfxicerwgjkcfwjhbmgscmmntdktxcxufqfipgccwwhicodznatsemvfhlekyfvrurxrjygnihpogpzqmynafvkjomignevkihjdlhtynggfjiaculaizqmvoizjemlepwpnyhplqexiwbqucywthnohcdbdgniegdbowuljoicfxyeidhjxspvpzyvgkegggakhnmbaagmnqlzpunsxylrabbrdnpsckajrtcgjceddwovodmgexznpgcqdqvtnteepshymcexwmjsejkdwgbcdvzirjxazeapnrngwrehxxkfrplzwucbhfqztskpmmyithtktpwlkecjxjjiieenjvppmgsuvaqnhfnhzcnymmhcbdupycydzomlrfqbjnmqkldcbooqjoqaebxmxrzfhqrqgmpcajcwhwyhhppszmiqcwevjyripepegzjxzshaaunxuvyqlkluamszccumwepxsqyseerzqnaumevszepewkccsmwysnypissjfbikpjdpjkbpjktdxyghambduitlulargzacszmhkosfarjkzefuqpjdiegpveaxngbiknuwyhuepsrtxotoclzgvyzufwvnfyufqzekvsujcadfaixvoqahrbgmedafuvxgymhemxdyttgcyydtvxttrcibbkupuvurqgbgjwnlmrtgkkkwibanldpjoxjpuhladqmtkumpvagbotzwajstsgdoxriqwenqjolzcqzwmerjjyozjoniqqcpeiinovypsuebddnpndqtxsabokfpyuofmxbojqroaewuhdyfofstqkijxcfldqlxibalhzybbqmileponzqfferzigfpwlazbojlhrcuotvqayxornhygvsznvkwtpwdzmxotnnhttzndrhrvdudccqjzjpdzmhakzlddjwmanooznkjdgkisnoialxiabwhsokmmokzpwicajkdxxnczkyhfspbywxhztlncjtlwikoymsjdovaroekirvuezehxybmgoxoeamqmgwzxkitujbswgmpugokcnzmegkenwwhowumraqmpgjnfiyqvdttdygtkznepxmlvpdykjmbxvqryjdqcluzfitltqcyyrkuuuofrmqojdtcyhxwqtvpjbtntnoyahetomtbxsimvuyjsijssvizxzhvfppxkmvjrlczmbljmnnyvicqlqreedztmenqmuerfgbenyylxxrwwkycdcbzwekjqzijjfgqgiockkikjnvxhfoaderdjlnaztgtysxkxxpmihvqavuypbqdzqfykmdyaprjuofkuaxspawxisrlafbyveenihxtpvpuvrqycubtukolhbguhgkparbcxbmcdsiskyhmtzghjinnceqytgjltcvjzdqhvbovyknmstfyshculaplqzjyzlfzhybsensljowuzqgwgmwllkbthenxeuybkbzhokaqyylbmtdhswergffwirhbhhxtwfrzpnqzlmliyfrffzxatssufahqmoxtbqdiwgmkwllhuqhfcqbtpypiferevchbjmmnyyfaryyczqagwialancrecefoxbadocbievglwypnsgockeekmrldywuhkcjjukzfejwgrnlvfbytxlbnbnutmoxebofddppyuaonaiwalbmxonhcgddaollcbouwmbkhdlbcppoxefnavjvxsavzvpdummjeyurddcvaxkwdqsvftuaahsofuuzjqpmbuorwdwdbtdnaopwkjvvphrxyriamyxmxeaorzkaozcnrdgwkipqkaapcdidvizvlvlbedmpptvykhnnjjadsdpqqfnvpvqtsxhddbwyvrxkaymnzkqwvnqmvjzpveqigolrhnfyviqinoekrjxwghaofwqcvhgcbhjnzfvioevrvcmknswspptogopnktenihvpetkycltiuqsqbhjlnzdfqvsnqwjzegczxogparfeyifyfueserjqmqvcycpjgpsxolfwdfaeelsueqtcbdrhclfucgeklexbnfiirhbuvmycgrkjrhmhfivswpbcqaksbjuxufjnxlgzjflipikxvdybyqruksuygokccrlslviuhtvdgybdopxdmbapkkzyopnapyuuyompqueytkfetbluygfkttbesxytsxlfjiqwnoluvrtxgzcfhihrmjnecvoeezvpvvpahkhtecpsuukssodwvetxjmgqwpcbwlyqqoskskioaztbnmwhyepmofdhlbxiranfzjbsaopresuqoszqnjevajkmqdbszbomafoehmwcovdzkfjbyoaptxwovpwknqtrtdtmsjncembohodmfqsjonvesaurapnysatxvckohlqwtxjqrskwsvbidnlrusuklzehbgsboqasoiywynqgifeizwjbewpgvwzhsgyslogdguzrdgnouqveenksoebexqokbbmvqjgfrnbvjlopzvbswowtmdeafufjzitydywpkjqoefastcysgdzcxixdpxjaaalbgxpnctqlyaugtcjvjdvcncxygplvnzzodzzvitdvukrekehqgtapjhvdwcacblfkkrojeljueajlmtqlhknezpjcbpfvbtbapzcoggdquenytdthnnxkqwxxkmhothnomsxujdnnnyuocodzqcqletpuyxnffbcvafmsymddvphduwdpdyxuhuqbgnuelyjksfzxlwdiakimvkabazslqjjsvenpbyvfnhmgwtgwhxkerbplljpojkbijqvsjzkospzwtnutvkgwcdtuaasmfsucanddotdvsfspaknokvewnwvyjbycdfjiqoupobwxmscbgvrchpyudogakrvsiuceroiwbmubpqfcneunwnkkbpedynmzyddvoxrkbedclenbgvofwdcshnhuhqhomdxrynnqgighytvyywkxytraqgilwawcndtybtpshxhietbabwtqtejgwhitcuqheeotmvxwcbojfadbhoqmtvydmdgzjovlgnuuiyycgzzcnjmahsokdfelixvblrhnwgwocypagiytakjljphzvimebgrfsakfpkfcmlikglqsmpjetlndgynvyuhrftjeinnhenktsudsxwwsxnyudlkpmbsdcvlarakedfhwrsvovbvcjntkfltjstirpafrcbmihbptlzfqwedkqimjchvcefsoraeioltondcvvymzbjlwyqzjpxzovksgfhplxlymajmvvljazjsaxghboelqbhscnmcysaajrfokcrulhqqqtxbfgycchfwozkxipvdavnyphbhprlxgxbsgqzfriogusgwbzbajcsmbqknyfwhfbuogguidpdwrefmaisfjhofnoktwshpytiabfokkffwfgxxjowqjolnofkysxewuauvmvxwnyfymeyshyebcuefhwegphcyjrbhktqdkvttzocjrcasvnwvckxtwfjdqfyjozzwxgcrjywdsabeayywseibemezhlhkmynlkifzahmpprdweepuuwejrvjusisgnizwxszxsmfnamuczpocuyfjwzlvvvxjwvivavcccyzuvjfwyqmitdmguhjhfuakoxmduogecnnwg"

	// run Cost 338.839ms
	TestFunc(longestContinuousSubstring, s1)

	// i := longestContinuousSubstring(s1)
	// log.Printf("%d \n", i)
	log.Println("----------------")
	// run Cost 999.3µs
	// run Cost 700.6µs
	TestFunc(longestContinuousSubstringOptimize, s1)
}

func longestContinuousSubstring(s string) int {
	maxLen := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s)+1; j++ {
			if strings.Contains(strContinuousLetter, s[i:j]) {
				if maxLen < len(s[i:j]) {
					log.Printf("max str %s \n", s[i:j])
					maxLen = len(s[i:j])
				}
			}
		}
	}

	return maxLen
}

// longestContinuousSubstring returns the length of the longest continuous substring
// in 's' that appears in 'strContinuousLetter'.
func longestContinuousSubstringOptimize(s string) int {
	if len(s) == 0 {
		return 0
	}

	maxLen := 1
	currentLen := 1
	startIndex := 0 // To keep track of the starting index of the current substring

	for i := 1; i < len(s); i++ {
		// Check if current character is consecutive to the previous character
		// if s[i]-s[i-1] == 1 && strings.Contains(strContinuousLetter, string(s[i])) {
		if s[i]-s[i-1] == 1 {
			currentLen++
			// Update maximum length and log the substring if a new maximum is found
			if currentLen > maxLen {
				maxLen = currentLen
				startIndex = i - currentLen + 1
				log.Printf("New max substring: %s\n", s[startIndex:i+1])
			}
		} else {
			currentLen = 1 // Reset the current length
		}
	}

	return maxLen
}

const strContinuousLetter = "abcdefghijklmnopqrstuvwxyz"

func TestFunc(f func(string) int, s string) {
	defer timeCost(time.Now())
	log.Println(f(s))
	// defer PrintMemUsage()

}

func timeCost(sT time.Time) {
	tc := time.Since(sT)
	fmt.Printf("run Cost %v \n", tc)
}