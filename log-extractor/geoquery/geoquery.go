package geoquery

import (
	"errors"
	"os"
	"strings"

	ipdb "github.com/ipipdotnet/ipdb-go"
)

var (
	//
	ipipDB *ipdb.City
	// CountryCode
	cc map[string]string
	// province code
	pc map[string]string
)

// GEOInfo ...
type GEOInfo struct {
	country string
	region  string
	city    string
}

// GetGEOInfo ...
func GetGEOInfo(ip string) (GEOInfo, error) {
	ipipRes, err := ipipDB.FindInfo(ip, "CN")
	if err != nil {
		return GEOInfo{}, err
	}
	return GEOInfo{
		// country_name 国家
		// region_name 省、直辖市、港澳台
		country: ipipRes.CountryName,
		region:  ipipRes.RegionName,
		city:    ipipRes.CityName,
	}, nil
}

// GetCCAndPC CountryCode and ProvinceCode
func GetCCAndPC(ip string) (string, string, error) {
	CC := ""
	PC := ""
	areaInfo, err := GetGEOInfo(ip)
	if err != nil {
		return "", "", err
	}
	//fmt.Printf("country:%v region:%v\n", areaInfo.country, areaInfo.region)
	if strings.Contains(areaInfo.country, "中国") {
		if provinceCode, ok := pc[areaInfo.region]; ok {
			CC = "CN"
			PC = provinceCode
			return CC, PC, nil
		}
		//
		return "CN", "", nil
	}
	if countryCode, ok := cc[areaInfo.country]; ok {
		CC = countryCode
		return CC, PC, nil
	}
	return CC, PC, errors.New("no GEO info in IP-GEO database, ip: " + ip)
}

// IsChinaMainlandIP ...
func IsChinaMainlandIP(ip string) bool {
	areaInfo, err := GetGEOInfo(ip)
	if err != nil {

	}
	if strings.Contains(areaInfo.country, "中国") {
		if strings.Contains(areaInfo.region, "香港") ||
			strings.Contains(areaInfo.region, "澳门") ||
			strings.Contains(areaInfo.region, "台湾") {
			return false
		}
		return true
	}
	return false
}

// GetProvinceCode ...
func GetProvinceCode(provinceName string) string {
	return pc[provinceName]
}

func fileExists(path string) bool {
	//os.Stat获取文件信息
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// Init ...
func Init(geoPath string) error {
	if !fileExists(geoPath) {
		return errors.New("IP-GEO file doesn't exits")
	}

	err := *new(error)
	ipipDB, err = ipdb.NewCity(geoPath)
	if err != nil {
		return err
	}

	pc = map[string]string{
		// 34个
		"北京":  "BJ110",
		"天津":  "TJ120",
		"河北":  "HE130",
		"山西":  "SX140",
		"内蒙古": "NM150",
		"辽宁":  "LN210",
		"吉林":  "JL220",
		"黑龙江": "HL230",
		"上海":  "SH310",
		"江苏":  "JS320",
		"浙江":  "ZJ330",
		"安徽":  "AH340",
		"福建":  "FJ350",
		"江西":  "JX360",
		"山东":  "SD370",
		"河南":  "HA410",
		"湖北":  "HB420",
		"湖南":  "HN430",
		"广东":  "GD440",
		"广西":  "GX450",
		"海南":  "HI460",
		"重庆":  "CQ500",
		"四川":  "SC510",
		"贵州":  "GZ520",
		"云南":  "YN530",
		"西藏":  "XZ540",
		"陕西":  "SN610",
		"甘肃":  "GS620",
		"青海":  "QH630",
		"宁夏":  "NX640",
		"新疆":  "XJ650",
		"台湾":  "TW710",
		"香港":  "HK810",
		"澳门":  "MO820",
	}

	cc = make(map[string]string)
	cc["安道尔"] = "AD"
	cc["阿联酋"] = "AE"
	cc["阿富汗"] = "AF"
	cc["安提瓜和巴布达"] = "AG"
	cc["安圭拉"] = "AI"
	cc["阿尔巴尼亚"] = "AL"
	cc["亚美尼亚"] = "AM"
	cc["安哥拉"] = "AO"
	cc["南极洲"] = "AQ"
	cc["阿根廷"] = "AR"
	cc["美属萨摩亚"] = "AS"
	cc["奥地利"] = "AT"
	cc["澳大利亚"] = "AU"
	cc["阿鲁巴"] = "AW"
	cc["奥兰群岛"] = "AX"
	cc["阿塞拜疆"] = "AZ"
	cc["波黑"] = "BA"
	cc["巴巴多斯"] = "BB"
	cc["孟加拉"] = "BD"
	cc["比利时"] = "BE"
	cc["布基纳法索"] = "BF"
	cc["保加利亚"] = "BG"
	cc["巴林"] = "BH"
	cc["布隆迪"] = "BI"
	cc["贝宁"] = "BJ"
	cc["圣巴泰勒米岛"] = "BL"
	cc["百慕大"] = "BM"
	cc["文莱"] = "BN"
	cc["玻利维亚"] = "BO"
	cc["荷兰加勒比区"] = "BQ"
	cc["巴西"] = "BR"
	cc["巴哈马"] = "BS"
	cc["不丹"] = "BT"
	cc["布韦岛"] = "BV"
	cc["博茨瓦纳"] = "BW"
	cc["白俄罗斯"] = "BY"
	cc["伯利兹"] = "BZ"
	cc["加拿大"] = "CA"
	cc["科科斯群岛"] = "cc"
	cc["中非"] = "CF"
	cc["瑞士"] = "CH"
	cc["智利"] = "CL"
	cc["喀麦隆"] = "CM"
	cc["哥伦比亚"] = "CO"
	cc["哥斯达黎加"] = "CR"
	cc["古巴"] = "CU"
	cc["佛得角"] = "CV"
	cc["圣诞岛"] = "CX"
	cc["塞浦路斯"] = "CY"
	cc["捷克"] = "CZ"
	cc["德国"] = "DE"
	cc["吉布提"] = "DJ"
	cc["丹麦"] = "DK"
	cc["多米尼克"] = "DM"
	cc["多米尼加"] = "DO"
	cc["阿尔及利亚"] = "DZ"
	cc["厄瓜多尔"] = "EC"
	cc["爱沙尼亚"] = "EE"
	cc["埃及"] = "EG"
	cc["西撒哈拉"] = "EH"
	cc["厄立特里亚"] = "ER"
	cc["西班牙"] = "ES"
	cc["芬兰"] = "FI"
	cc["斐济群岛"] = "FJ"
	cc["马尔维纳斯群岛（福克兰）"] = "FK"
	cc["密克罗尼西亚联邦"] = "FM"
	cc["法罗群岛"] = "FO"
	cc["法国"] = "FR"
	cc["加蓬"] = "GA"
	cc["格林纳达"] = "GD"
	cc["格鲁吉亚"] = "GE"
	cc["法属圭亚那"] = "GF"
	cc["加纳"] = "GH"
	cc["直布罗陀"] = "GI"
	cc["格陵兰"] = "GL"
	cc["几内亚"] = "GN"
	cc["瓜德罗普"] = "GP"
	cc["赤道几内亚"] = "GQ"
	cc["希腊"] = "GR"
	cc["南乔治亚岛和南桑威奇群岛"] = "GS"
	cc["危地马拉"] = "GT"
	cc["关岛"] = "GU"
	cc["几内亚比绍"] = "GW"
	cc["圭亚那"] = "GY"
	cc["香港"] = "CNHK810"
	cc["亚太地区"] = "CNHK810"
	cc["赫德岛和麦克唐纳群岛"] = "HM"
	cc["洪都拉斯"] = "HN"
	cc["克罗地亚"] = "HR"
	cc["海地"] = "HT"
	cc["匈牙利"] = "HU"
	cc["印尼"] = "ID"
	cc["爱尔兰"] = "IE"
	cc["以色列"] = "IL"
	cc["马恩岛"] = "IM"
	cc["印度"] = "IN"
	cc["英属印度洋领地"] = "IO"
	cc["伊拉克"] = "IQ"
	cc["伊朗"] = "IR"
	cc["冰岛"] = "IS"
	cc["意大利"] = "IT"
	cc["泽西岛"] = "JE"
	cc["牙买加"] = "JM"
	cc["约旦"] = "JO"
	cc["日本"] = "JP"
	cc["柬埔寨"] = "KH"
	cc["基里巴斯"] = "KI"
	cc["科摩罗"] = "KM"
	cc["科威特"] = "KW"
	cc["开曼群岛"] = "KY"
	cc["黎巴嫩"] = "LB"
	cc["列支敦士登"] = "LI"
	cc["斯里兰卡"] = "LK"
	cc["利比里亚"] = "LR"
	cc["莱索托"] = "LS"
	cc["立陶宛"] = "LT"
	cc["卢森堡"] = "LU"
	cc["拉脱维亚"] = "LV"
	cc["利比亚"] = "LY"
	cc["摩洛哥"] = "MA"
	cc["摩纳哥"] = "MC"
	cc["摩尔多瓦"] = "MD"
	cc["黑山"] = "ME"
	cc["法属圣马丁"] = "MF"
	cc["马达加斯加"] = "MG"
	cc["马绍尔群岛"] = "MH"
	cc["马其顿"] = "MK"
	cc["马里"] = "ML"
	cc["缅甸"] = "MM"
	cc["澳门"] = "CNMO820"
	cc["马提尼克"] = "MQ"
	cc["毛里塔尼亚"] = "MR"
	cc["蒙塞拉特岛"] = "MS"
	cc["马耳他"] = "MT"
	cc["马尔代夫"] = "MV"
	cc["马拉维"] = "MW"
	cc["墨西哥"] = "MX"
	cc["马来西亚"] = "MY"
	cc["纳米比亚"] = "NA"
	cc["尼日尔"] = "NE"
	cc["诺福克岛"] = "NF"
	cc["尼日利亚"] = "NG"
	cc["尼加拉瓜"] = "NI"
	cc["荷兰"] = "NL"
	cc["挪威"] = "NO"
	cc["尼泊尔"] = "NP"
	cc["瑙鲁"] = "NR"
	cc["阿曼"] = "OM"
	cc["巴拿马"] = "PA"
	cc["秘鲁"] = "PE"
	cc["法属波利尼西亚"] = "PF"
	cc["巴布亚新几内亚"] = "PG"
	cc["菲律宾"] = "PH"
	cc["巴基斯坦"] = "PK"
	cc["波兰"] = "PL"
	cc["皮特凯恩群岛"] = "PN"
	cc["波多黎各"] = "PR"
	cc["巴勒斯坦"] = "PS"
	cc["帕劳"] = "PW"
	cc["巴拉圭"] = "PY"
	cc["卡塔尔"] = "QA"
	cc["留尼汪"] = "RE"
	cc["罗马尼亚"] = "RO"
	cc["塞尔维亚"] = "RS"
	cc["俄罗斯"] = "RU"
	cc["卢旺达"] = "RW"
	cc["所罗门群岛"] = "SB"
	cc["塞舌尔"] = "SC"
	cc["苏丹"] = "SD"
	cc["瑞典"] = "SE"
	cc["新加坡"] = "SG"
	cc["斯洛文尼亚"] = "SI"
	cc["斯瓦尔巴群岛和"] = "SJ"
	cc["斯洛伐克"] = "SK"
	cc["塞拉利昂"] = "SL"
	cc["圣马力诺"] = "SM"
	cc["塞内加尔"] = "SN"
	cc["索马里"] = "SO"
	cc["苏里南"] = "SR"
	cc["南苏丹"] = "SS"
	cc["圣多美和普林西比"] = "ST"
	cc["萨尔瓦多"] = "SV"
	cc["叙利亚"] = "SY"
	cc["斯威士兰"] = "SZ"
	cc["特克斯和凯科斯群岛"] = "TC"
	cc["乍得"] = "TD"
	cc["多哥"] = "TG"
	cc["泰国"] = "TH"
	cc["托克劳"] = "TK"
	cc["东帝汶"] = "TL"
	cc["突尼斯"] = "TN"
	cc["汤加"] = "TO"
	cc["土耳其"] = "TR"
	cc["图瓦卢"] = "TV"
	cc["坦桑尼亚"] = "TZ"
	cc["乌克兰"] = "UA"
	cc["乌干达"] = "UG"
	cc["美国"] = "US"
	cc["乌拉圭"] = "UY"
	cc["梵蒂冈"] = "VA"
	cc["委内瑞拉"] = "VE"
	cc["英属维尔京群岛"] = "VG"
	cc["美属维尔京群岛"] = "VI"
	cc["越南"] = "VN"
	cc["瓦利斯和富图纳"] = "WF"
	cc["萨摩亚"] = "WS"
	cc["也门"] = "YE"
	cc["马约特"] = "YT"
	cc["南非"] = "ZA"
	cc["赞比亚"] = "ZM"
	cc["津巴布韦"] = "ZW"
	cc["中国"] = "CN"
	cc["刚果（布）"] = "CG"
	cc["刚果（金）"] = "CD"
	cc["莫桑比克"] = "MZ"
	cc["根西岛"] = "GG"
	cc["冈比亚"] = "GM"
	cc["北马里亚纳群岛"] = "MP"
	cc["埃塞俄比亚"] = "ET"
	cc["新喀里多尼亚"] = "NC"
	cc["瓦努阿图"] = "VU"
	cc["法属南部领地"] = "TF"
	cc["纽埃"] = "NU"
	cc["美国本土外小岛屿"] = "UM"
	cc["库克群岛"] = "CK"
	cc["英国"] = "GB"
	cc["特立尼达和多巴哥"] = "TT"
	cc["圣文森特和格林纳丁斯"] = "VC"
	cc["中华民国"] = "CNTW710"
	cc["台湾"] = "CNTW710"
	cc["新西兰"] = "NZ"
	cc["沙特阿拉伯"] = "SA"
	cc["老挝"] = "LA"
	cc["朝鲜"] = "KP"
	cc["韩国"] = "KR"
	cc["葡萄牙"] = "PT"
	cc["吉尔吉斯斯坦"] = "KG"
	cc["哈萨克斯坦"] = "KZ"
	cc["塔吉克斯坦"] = "TJ"
	cc["土库曼斯坦"] = "TM"
	cc["乌兹别克斯坦"] = "UZ"
	cc["圣基茨和尼维斯"] = "KN"
	cc["圣皮埃尔和密克隆"] = "PM"
	cc["圣赫勒拿"] = "SH"
	cc["圣卢西亚"] = "LC"
	cc["毛里求斯"] = "MU"
	cc["科特迪瓦"] = "CI"
	cc["肯尼亚"] = "KE"
	cc["蒙古国"] = "MN"

	//for k, v := range cc {
	//	fmt.Printf("k:%v\tv:%v\n", k, v)
	//}
	return nil
}
