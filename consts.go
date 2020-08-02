package main

var MidLandList = [10]string{"11B00000", "11D10000", "11D20000", "11C20000", "11C10000", "11F20000", "11F10000", "11H10000", "11H20000", "11G00000"}

//x,y를 concat한 값으로 사용했다.
var ShortTempList = [1628]string{
	"82134",
	"5591",
	"5472",
	"83102",
	"6075",
	"58123",
	"64129",
	"67143",
	"87103",
	"98101",
	"77123",
	"85120",
	"6291",
	"5460",
	"62125",
	"38129",
	"65134",
	"66132",
	"7878",
	"8689",
	"8582",
	"8785",
	"8373",
	"98125",
	"8095",
	"8893",
	"4575",
	"8974",
	"7874",
	"9379",
	"74137",
	"6298",
	"55110",
	"5592",
	"9086",
	"61125",
	"59122",
	"6491",
	"5971",
	"5995",
	"6076",
	"4265",
	"85107",
	"9976",
	"9091",
	"58114",
	"77126",
	"61104",
	"5695",
	"6682",
	"6278",
	"51128",
	"65113",
	"58119",
	"75136",
	"8486",
	"5364",
	"59114",
	"53129",
	"81126",
	"63102",
	"6387",
	"5067",
	"10294",
	"80108",
	"63124",
	"65116",
	"83116",
	"66109",
	"8180",
	"8672",
	"7678",
	"56126",
	"53106",
	"10489",
	"7993",
	"9692",
	"8368",
	"9775",
	"8891",
	"79120",
	"5595",
	"5975",
	"7599",
	"57104",
	"6466",
	"9197",
	"8376",
	"8183",
	"8992",
	"57134",
	"66135",
	"5784",
	"5337",
	"56106",
	"4859",
	"8375",
	"9284",
	"58128",
	"70134",
	"79136",
	"5798",
	"8593",
	"8792",
	"8881",
	"9875",
	"7367",
	"6276",
	"4772",
	"8096",
	"8384",
	"68123",
	"55111",
	"47108",
	"6580",
	"9890",
	"102103",
	"94131",
	"93132",
	"5664",
	"9996",
	"9483",
	"8181",
	"57129",
	"88118",
	"7169",
	"96103",
	"91107",
	"71117",
	"72132",
	"65111",
	"5380",
	"8571",
	"9576",
	"62121",
	"50109",
	"90104",
	"5491",
	"7470",
	"51130",
	"51120",
	"66100",
	"59107",
	"5484",
	"6576",
	"6070",
	"5962",
	"8889",
	"10084",
	"7797",
	"6093",
	"96100",
	"8980",
	"8184",
	"6088",
	"87101",
	"8588",
	"81139",
	"5892",
	"5994",
	"6774",
	"6787",
	"6381",
	"6375",
	"4960",
	"8687",
	"63108",
	"94118",
	"71104",
	"8075",
	"9781",
	"5563",
	"8990",
	"64114",
	"67142",
	"6585",
	"9169",
	"7581",
	"72107",
	"6469",
	"9492",
	"80100",
	"6684",
	"7363",
	"7370",
	"8573",
	"9874",
	"58133",
	"75117",
	"6184",
	"66137",
	"80118",
	"63103",
	"5783",
	"4573",
	"9476",
	"56107",
	"5782",
	"7181",
	"5179",
	"52105",
	"6960",
	"96105",
	"7987",
	"91110",
	"89110",
	"8291",
	"8679",
	"74102",
	"5285",
	"5260",
	"4858",
	"9090",
	"8071",
	"8968",
	"85119",
	"90105",
	"9389",
	"95100",
	"73105",
	"6578",
	"5771",
	"5865",
	"9190",
	"69100",
	"60119",
	"89123",
	"6294",
	"7189",
	"8699",
	"103105",
	"62127",
	"68121",
	"70117",
	"81118",
	"8480",
	"7482",
	"6186",
	"6189",
	"5567",
	"8796",
	"8989",
	"6798",
	"66105",
	"89130",
	"88111",
	"103107",
	"7587",
	"56127",
	"68110",
	"6783",
	"94100",
	"8490",
	"8577",
	"5233",
	"9474",
	"66121",
	"92134",
	"5792",
	"56133",
	"7593",
	"97108",
	"7979",
	"8481",
	"7677",
	"7789",
	"92131",
	"6997",
	"9990",
	"7996",
	"82106",
	"65115",
	"71137",
	"5986",
	"7471",
	"6896",
	"5381",
	"5561",
	"8198",
	"84100",
	"97101",
	"9879",
	"69119",
	"89137",
	"5497",
	"10393",
	"9180",
	"9779",
	"65123",
	"91123",
	"68105",
	"5276",
	"6593",
	"6669",
	"5960",
	"5163",
	"58122",
	"6397",
	"6497",
	"6599",
	"5985",
	"6968",
	"8979",
	"9076",
	"69115",
	"53112",
	"5889",
	"83100",
	"8787",
	"6999",
	"68141",
	"77111",
	"59125",
	"54126",
	"5177",
	"8789",
	"65104",
	"5980",
	"85110",
	"61132",
	"4968",
	"8292",
	"7198",
	"8688",
	"10285",
	"60139",
	"73124",
	"6996",
	"4565",
	"4658",
	"91106",
	"7689",
	"55132",
	"55135",
	"92130",
	"6379",
	"8492",
	"91111",
	"8677",
	"59126",
	"60104",
	"56109",
	"6662",
	"8676",
	"62126",
	"73111",
	"5870",
	"90113",
	"6090",
	"5834",
	"53124",
	"10086",
	"61115",
	"56129",
	"73101",
	"70100",
	"5471",
	"7775",
	"8682",
	"66114",
	"73103",
	"61110",
	"5680",
	"76118",
	"84128",
	"73114",
	"73107",
	"57125",
	"61114",
	"57122",
	"63129",
	"79102",
	"8072",
	"8780",
	"6479",
	"5167",
	"4962",
	"5779",
	"8494",
	"61135",
	"64128",
	"62108",
	"6068",
	"10299",
	"7871",
	"61118",
	"5799",
	"60100",
	"7560",
	"79107",
	"87142",
	"7099",
	"64112",
	"7070",
	"5765",
	"5667",
	"9089",
	"60123",
	"63130",
	"96120",
	"5866",
	"78124",
	"77125",
	"78110",
	"54104",
	"6882",
	"6166",
	"9085",
	"8470",
	"51131",
	"72133",
	"92133",
	"82118",
	"4937",
	"9877",
	"69106",
	"83114",
	"5372",
	"81102",
	"96104",
	"90115",
	"8172",
	"59124",
	"10484",
	"62114",
	"6177",
	"9580",
	"8496",
	"8975",
	"8174",
	"7785",
	"4848",
	"50108",
	"10594",
	"10097",
	"87109",
	"6175",
	"9374",
	"60120",
	"54129",
	"5599",
	"5794",
	"9377",
	"5032",
	"63117",
	"6194",
	"6980",
	"95114",
	"6871",
	"7965",
	"8790",
	"51129",
	"62101",
	"5691",
	"7091",
	"6581",
	"86106",
	"63122",
	"59121",
	"69122",
	"70105",
	"90109",
	"9269",
	"5332",
	"54125",
	"63100",
	"10394",
	"8695",
	"10283",
	"53108",
	"58104",
	"102114",
	"88110",
	"9593",
	"102106",
	"9776",
	"6768",
	"10395",
	"8595",
	"60110",
	"6894",
	"5989",
	"58127",
	"64127",
	"80138",
	"90136",
	"57132",
	"74131",
	"76141",
	"67106",
	"55101",
	"6892",
	"6257",
	"8178",
	"57116",
	"59137",
	"74110",
	"72110",
	"8669",
	"7766",
	"7283",
	"7784",
	"6390",
	"7463",
	"8295",
	"8277",
	"5075",
	"8175",
	"61100",
	"49104",
	"5785",
	"5570",
	"81116",
	"68113",
	"6979",
	"5275",
	"21135",
	"21132",
	"76129",
	"68109",
	"10095",
	"9577",
	"9170",
	"7686",
	"66120",
	"88138",
	"5277",
	"7265",
	"9491",
	"92126",
	"55112",
	"7092",
	"7186",
	"75110",
	"6886",
	"6170",
	"5458",
	"68103",
	"67104",
	"62118",
	"71132",
	"88100",
	"97106",
	"9082",
	"5169",
	"90108",
	"8774",
	"61129",
	"58129",
	"61109",
	"5791",
	"6369",
	"6258",
	"94106",
	"8099",
	"77113",
	"6096",
	"55108",
	"6389",
	"80101",
	"77100",
	"8391",
	"5788",
	"4264",
	"9788",
	"8094",
	"9773",
	"56122",
	"59118",
	"6183",
	"89115",
	"82101",
	"8488",
	"7588",
	"65129",
	"96126",
	"6188",
	"8374",
	"5890",
	"5568",
	"68102",
	"10286",
	"61119",
	"70119",
	"49110",
	"5587",
	"6059",
	"88106",
	"5774",
	"66103",
	"80127",
	"52103",
	"7793",
	"9189",
	"8892",
	"64140",
	"71143",
	"5872",
	"8188",
	"67101",
	"52102",
	"48110",
	"83106",
	"52107",
	"8383",
	"8991",
	"66101",
	"66125",
	"69123",
	"68114",
	"5299",
	"57110",
	"92105",
	"10076",
	"8690",
	"62115",
	"60138",
	"89111",
	"46119",
	"7898",
	"6890",
	"6172",
	"90101",
	"8779",
	"10082",
	"99124",
	"79127",
	"5887",
	"7968",
	"8484",
	"90102",
	"93113",
	"8273",
	"60127",
	"60125",
	"72111",
	"5596",
	"8275",
	"8197",
	"84103",
	"98109",
	"95111",
	"50102",
	"5267",
	"99114",
	"8381",
	"59116",
	"69132",
	"81141",
	"63106",
	"8272",
	"7672",
	"7981",
	"69133",
	"71106",
	"6197",
	"7692",
	"5775",
	"5574",
	"6799",
	"61137",
	"6483",
	"6962",
	"9375",
	"6982",
	"5660",
	"5877",
	"80106",
	"64115",
	"76116",
	"76111",
	"61111",
	"6567",
	"5174",
	"10088",
	"8196",
	"61128",
	"79143",
	"6489",
	"7085",
	"87105",
	"86107",
	"7980",
	"68116",
	"6995",
	"4970",
	"8195",
	"5669",
	"3364",
	"8170",
	"9477",
	"9290",
	"9177",
	"58124",
	"53120",
	"6297",
	"5582",
	"94103",
	"54130",
	"70123",
	"10191",
	"91105",
	"73133",
	"80123",
	"64141",
	"74114",
	"9878",
	"6073",
	"10281",
	"55118",
	"9981",
	"8382",
	"7888",
	"61103",
	"5999",
	"53104",
	"6778",
	"10079",
	"70124",
	"79132",
	"69114",
	"5760",
	"10193",
	"62119",
	"67131",
	"70113",
	"10487",
	"5279",
	"6766",
	"92100",
	"9286",
	"85105",
	"73143",
	"68107",
	"63113",
	"54105",
	"57103",
	"60102",
	"7876",
	"7275",
	"63128",
	"70130",
	"70111",
	"60105",
	"6796",
	"8397",
	"82103",
	"7573",
	"9873",
	"68126",
	"93131",
	"75125",
	"96118",
	"76114",
	"56108",
	"5687",
	"10284",
	"9983",
	"61116",
	"67121",
	"7170",
	"5566",
	"102102",
	"8587",
	"7768",
	"7385",
	"55109",
	"5789",
	"7076",
	"9099",
	"5764",
	"78106",
	"9098",
	"8591",
	"62132",
	"5496",
	"5396",
	"5569",
	"7787",
	"59100",
	"5681",
	"6574",
	"7891",
	"56123",
	"66126",
	"71115",
	"5793",
	"5232",
	"10091",
	"64123",
	"84115",
	"6378",
	"5557",
	"5885",
	"5790",
	"7290",
	"7266",
	"66130",
	"69116",
	"5593",
	"4395",
	"7271",
	"9794",
	"9274",
	"8668",
	"55105",
	"6765",
	"5559",
	"5071",
	"9095",
	"8271",
	"63114",
	"64136",
	"83113",
	"7069",
	"58109",
	"4462",
	"90106",
	"81106",
	"65125",
	"69129",
	"70125",
	"76121",
	"92113",
	"56119",
	"55107",
	"77101",
	"8877",
	"84106",
	"9975",
	"74111",
	"5685",
	"5269",
	"7192",
	"7171",
	"6178",
	"127127",
	"50132",
	"61113",
	"58111",
	"51113",
	"9181",
	"53114",
	"5997",
	"7475",
	"71110",
	"65110",
	"7582",
	"9486",
	"8387",
	"6899",
	"55133",
	"73120",
	"6762",
	"74100",
	"6976",
	"8179",
	"7380",
	"5777",
	"8097",
	"103110",
	"9278",
	"73135",
	"63112",
	"7280",
	"6884",
	"6196",
	"6295",
	"6094",
	"57135",
	"7497",
	"7493",
	"7295",
	"7366",
	"6179",
	"8279",
	"48120",
	"95117",
	"7398",
	"63110",
	"50128",
	"68101",
	"72112",
	"6468",
	"7574",
	"9978",
	"63120",
	"85145",
	"5689",
	"8898",
	"86101",
	"9066",
	"65107",
	"62117",
	"55113",
	"6789",
	"9791",
	"7991",
	"89125",
	"78111",
	"78118",
	"6880",
	"6959",
	"6168",
	"5756",
	"8065",
	"62128",
	"63127",
	"57106",
	"6677",
	"7670",
	"6686",
	"5770",
	"5161",
	"5271",
	"50129",
	"62123",
	"65138",
	"84113",
	"4635",
	"9380",
	"67118",
	"76117",
	"5066",
	"7371",
	"72129",
	"73116",
	"8393",
	"8479",
	"61126",
	"61117",
	"61133",
	"74121",
	"6038",
	"6699",
	"9986",
	"81115",
	"6399",
	"57102",
	"6695",
	"5359",
	"9394",
	"5875",
	"59123",
	"59115",
	"85146",
	"7983",
	"6596",
	"58102",
	"7168",
	"10090",
	"59128",
	"9977",
	"83120",
	"68106",
	"8086",
	"5333",
	"57127",
	"85102",
	"56112",
	"9176",
	"9281",
	"9069",
	"9884",
	"72119",
	"75114",
	"56110",
	"7796",
	"62120",
	"92132",
	"90125",
	"81113",
	"9885",
	"64116",
	"75143",
	"8976",
	"61138",
	"72134",
	"7288",
	"8882",
	"63132",
	"62140",
	"76124",
	"93108",
	"64113",
	"58107",
	"6978",
	"4666",
	"70109",
	"53100",
	"4964",
	"4468",
	"65106",
	"61134",
	"58136",
	"96127",
	"9282",
	"8584",
	"5338",
	"144123",
	"9882",
	"6182",
	"5786",
	"9695",
	"9487",
	"6598",
	"48109",
	"6376",
	"5475",
	"64126",
	"58131",
	"94123",
	"68111",
	"8288",
	"5166",
	"92102",
	"8884",
	"57123",
	"60137",
	"87140",
	"77115",
	"90107",
	"9376",
	"7680",
	"89118",
	"55104",
	"5365",
	"5858",
	"4932",
	"89120",
	"75102",
	"6975",
	"10592",
	"9495",
	"69107",
	"51110",
	"5257",
	"10190",
	"8878",
	"63123",
	"60132",
	"5996",
	"7872",
	"6769",
	"7774",
	"10485",
	"75112",
	"5797",
	"6292",
	"10696",
	"9183",
	"55129",
	"101119",
	"5465",
	"4663",
	"86144",
	"52110",
	"57105",
	"83110",
	"62110",
	"6679",
	"7293",
	"5478",
	"55126",
	"64119",
	"65124",
	"7199",
	"6750",
	"5863",
	"88113",
	"84124",
	"8868",
	"7798",
	"7395",
	"52112",
	"8674",
	"60128",
	"9774",
	"9674",
	"95119",
	"7483",
	"6080",
	"5958",
	"52125",
	"96123",
	"60111",
	"5991",
	"6380",
	"7468",
	"102117",
	"7682",
	"5773",
	"9883",
	"77131",
	"57101",
	"5362",
	"9578",
	"9678",
	"63126",
	"5597",
	"7071",
	"5964",
	"6467",
	"8289",
	"101109",
	"7975",
	"73119",
	"66110",
	"5899",
	"5998",
	"6888",
	"5778",
	"103109",
	"48131",
	"60124",
	"64122",
	"64124",
	"7781",
	"65103",
	"6989",
	"8899",
	"8866",
	"7466",
	"6977",
	"6863",
	"96116",
	"8686",
	"66128",
	"67107",
	"58106",
	"5938",
	"6359",
	"4455",
	"63125",
	"57119",
	"75130",
	"6482",
	"6780",
	"7187",
	"6385",
	"6666",
	"58125",
	"57128",
	"70122",
	"61102",
	"5539",
	"76103",
	"9188",
	"288",
	"55124",
	"49130",
	"63111",
	"59110",
	"8495",
	"9294",
	"86102",
	"4832",
	"61124",
	"58121",
	"64120",
	"69121",
	"9496",
	"9484",
	"7585",
	"6074",
	"65105",
	"76122",
	"6493",
	"48107",
	"7386",
	"54123",
	"66106",
	"85126",
	"73117",
	"66122",
	"84132",
	"7292",
	"8575",
	"8895",
	"9676",
	"92120",
	"5666",
	"10298",
	"60115",
	"6091",
	"8671",
	"7769",
	"6461",
	"57124",
	"56131",
	"59135",
	"86122",
	"80124",
	"75104",
	"6691",
	"9283",
	"6296",
	"9574",
	"97124",
	"71102",
	"84117",
	"68119",
	"69143",
	"53113",
	"5795",
	"55127",
	"5974",
	"10280",
	"60122",
	"6163",
	"101102",
	"85108",
	"7786",
	"5692",
	"6973",
	"6670",
	"6078",
	"10178",
	"82121",
	"7095",
	"91113",
	"50130",
	"67125",
	"5677",
	"67140",
	"88114",
	"7473",
	"5633",
	"5976",
	"6371",
	"9179",
	"9075",
	"8284",
	"64104",
	"64108",
	"73128",
	"5893",
	"10184",
	"55106",
	"5376",
	"88101",
	"57109",
	"7282",
	"6777",
	"5274",
	"53125",
	"51125",
	"56125",
	"5499",
	"8986",
	"9077",
	"7867",
	"10383",
	"5990",
	"5586",
	"6476",
	"8569",
	"7667",
	"50131",
	"5671",
	"5772",
	"10189",
	"5984",
	"6077",
	"6176",
	"6273",
	"57126",
	"79125",
	"51109",
	"57108",
	"6156",
	"69104",
	"7081",
	"5470",
	"75122",
	"8977",
	"9373",
	"5988",
	"9797",
	"9289",
	"98106",
	"9677",
	"60121",
	"95129",
	"6271",
	"8596",
	"8696",
	"64134",
	"84123",
	"54112",
	"5883",
	"10093",
	"77105",
	"86103",
	"8985",
	"8691",
	"69125",
	"72105",
	"5353",
	"5676",
	"10385",
	"7496",
	"63133",
	"6674",
	"5371",
	"10288",
	"60126",
	"9777",
	"64130",
	"53130",
	"87108",
	"6037",
	"82119",
	"72101",
	"6290",
	"7383",
	"5894",
	"54124",
	"10384",
	"56128",
	"61136",
	"6289",
	"6782",
	"5579",
	"5461",
	"55123",
	"62124",
	"60118",
	"73134",
	"5576",
	"4567",
	"92110",
	"87112",
	"5377",
	"8497",
	"8176",
	"67102",
	"71121",
	"56102",
	"5977",
	"9083",
	"7772",
	"7985",
	"5675",
	"54100",
	"7175",
	"9198",
	"67116",
	"5696",
	"5369",
	"6764",
	"5473",
	"83104",
	"8578",
	"8187",
	"50124",
	"55125",
	"5874",
	"69108",
	"76123",
	"5138",
	"86109",
	"8475",
	"5433",
	"6590",
	"6165",
	"8299",
	"82107",
	"8697",
	"57130",
	"87141",
	"74103",
	"6873",
	"5280",
	"9587",
	"8870",
	"5683",
	"5272",
	"80109",
	"8191",
	"10075",
	"9673",
	"98123",
	"5884",
	"7570",
	"6970",
	"6071",
	"4562",
	"99110",
	"9092",
	"77122",
	"69113",
	"74113",
	"9268",
	"7994",
	"8472",
	"10077",
	"65126",
	"79112",
	"5856",
	"10590",
	"8398",
	"9186",
	"8594",
	"10181",
	"85128",
	"6396",
	"6573",
	"6663",
	"10293",
	"9489",
	"9171",
	"59120",
	"59130",
	"5498",
	"58110",
	"53128",
	"64135",
	"4884",
	"4671",
	"5973",
	"67100",
	"64117",
	"56134",
	"8597",
	"69117",
	"7465",
	"80103",
	"8768",
	"5553",
	"81105",
	"94114",
	"47128",
	"57121",
	"71140",
	"7096",
	"8879",
	"8283",
	"9675",
	"6597",
	"5661",
	"9296",
	"6785",
	"5761",
	"8070",
	"5438",
	"5763",
	"103111",
	"8076",
	"61120",
	"81130",
	"84147",
	"86117",
	"59117",
	"6190",
	"6393",
	"6095",
	"6192",
	"6562",
	"8378",
	"7094",
	"6092",
	"5780",
	"8981",
	"68100",
	"56130",
	"60134",
	"76115",
	"46109",
	"6486",
	"100103",
	"78134",
	"90121",
	"78141",
	"5668",
	"6061",
	"96108",
	"9175",
	"9277",
	"10483",
	"63115",
	"69138",
	"6492",
	"8784",
	"8585",
	"5238",
	"65109",
	"5165",
	"5176",
	"81107",
	"59127",
	"65114",
	"92118",
	"77116",
	"93112",
	"58126",
	"4873",
	"9980",
	"78139",
	"53101",
	"51111",
	"6366",
	"55128",
	"61131",
	"66118",
	"78121",
	"8074",
	"7681",
	"4836",
	"6364",
	"102115",
	"8874",
	"8667",
	"5382",
	"71125",
	"75115",
	"67112",
	"6779",
	"9393",
	"9594",
	"61130",
	"64131",
	"86119",
	"7087",
	"6377",
	"8580",
	"62122",
	"69118",
	"72113",
	"8073",
	"8897",
	"8386",
	"63104",
	"5969",
	"5859",
	"10296",
	"5672",
	"5953",
	"9392",
	"102100",
	"65139",
	"87138",
	"65101",
	"5584",
	"5132",
	"5594",
	"5480",
	"7374",
	"5965",
	"9876",
	"8890",
	"91134",
	"98119",
	"8781",
	"61127",
	"9974",
	"87129",
	"82146",
	"62130",
	"53102",
	"5468",
	"10295",
	"5682",
	"9391",
	"65120",
	"97127",
	"77139",
	"60109",
	"72139",
	"6085",
	"6087",
	"6488",
	"70121",
	"72122",
	"72127",
	"97125",
	"5580",
	"91103",
	"85101",
	"9479",
	"5694",
	"7083",
	"77103",
	"8270",
	"10179",
	"54127",
	"59119",
	"71138",
	"56124",
	"6760",
	"6266",
	"7976",
	"127129",
	"8463",
	"9383",
	"50126",
	"61121",
	"7075",
	"9096",
	"9299",
	"62129",
	"5873",
	"54128",
	"6372",
	"72126",
	"97126",
	"53110",
	"101118",
	"81104",
	"83136",
	"77114",
	"6198",
	"4860",
	"7967",
	"10494",
	"79105",
	"126127",
	"9068",
	"9178",
	"8783",
	"5876",
	"71114",
	"6089",
	"6477",
	"63136",
	"5697",
	"48111",
	"6388",
}

var MidTempList = [173]string{
	"11B10101",
	"11B20201",
	"11B20601",
	"11B20605",
	"11B20602",
	"11B10103",
	"11B10102",
	"11B20606",
	"11B20603",
	"11B20609",
	"11B20612",
	"11B20610",
	"11B20611",
	"11B20604",
	"11B20503",
	"11B20501",
	"11B20502",
	"11B20504",
	"11B20701",
	"11B20703",
	"11B20702",
	"11B20301",
	"11B20302",
	"11B20305",
	"11B20304",
	"11B20401",
	"11B20402",
	"11B20403",
	"11B20404",
	"11B20101",
	"11B20102",
	"11B20202",
	"11B20204",
	"11B20203",
	"11A00101",
	"11H20201",
	"11H20101",
	"11H20304",
	"11H20102",
	"11H20301",
	"11H20601",
	"11H20603",
	"11H20604",
	"11H20602",
	"11H20701",
	"11H20704",
	"11H20402",
	"11H20502",
	"11H20503",
	"11H20703",
	"11H20501",
	"11H20401",
	"11H20403",
	"11H20404",
	"11H20405",
	"11H10701",
	"11H10702",
	"11H10703",
	"11H10704",
	"11H10705",
	"11H10601",
	"11H10602",
	"11H10603",
	"11H10604",
	"11H10605",
	"11H10501",
	"11H10502",
	"11H10503",
	"11H10302",
	"11H10301",
	"11H10303",
	"11H10401",
	"11H10402",
	"11H10403",
	"11H10101",
	"11H10102",
	"11H10201",
	"11H10202",
	"11E00101",
	"11E00102",
	"11F20501",
	"11F20503",
	"11F20502",
	"11F20504",
	"11F20505",
	"21F20102",
	"21F20101",
	"21F20801",
	"21F20804",
	"21F20802",
	"21F20201",
	"21F20803",
	"11F20701",
	"11F20603",
	"11F20402",
	"11F20601",
	"11F20602",
	"11F20301",
	"11F20303",
	"11F20304",
	"11F20302",
	"11F20401",
	"11F20403",
	"11F20404",
	"11F10201",
	"11F10202",
	"21F10501",
	"11F10203",
	"21F10502",
	"11F10401",
	"21F10601",
	"11F10302",
	"21F10602",
	"11F10403",
	"11F10204",
	"11F10402",
	"11F10301",
	"11F10303",
	"11C20401",
	"11C20404",
	"11C20402",
	"11C20602",
	"11C20403",
	"11C20601",
	"11C20301",
	"11C20302",
	"11C20303",
	"11C20101",
	"11C20102",
	"11C20103",
	"11C20104",
	"11C20201",
	"11C20202",
	"11C20502",
	"11C20501",
	"11C10301",
	"11C10304",
	"11C10303",
	"11C10102",
	"11C10101",
	"11C10103",
	"11C10201",
	"11C10202",
	"11C10302",
	"11C10403",
	"11C10402",
	"11C10401",
	"11D10101",
	"11D10102",
	"11D10201",
	"11D10202",
	"11D10301",
	"11D10302",
	"11D10401",
	"11D10402",
	"11D10501",
	"11D10502",
	"11D10503",
	"11D20201",
	"11D20401",
	"11D20402",
	"11D20403",
	"11D20501",
	"11D20601",
	"11D20602",
	"11D20301",
	"11G00201",
	"11G00401",
	"11G00101",
	"11G00501",
	"11G00302",
	"11G00601",
	"11G00800",
}
