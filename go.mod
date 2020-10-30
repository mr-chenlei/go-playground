module github.com/MrVegeta/go-playground

require (
	github.com/LiamHaworth/go-tproxy v0.0.0-20190726054950-ef7efd7f24ed
	github.com/PuerkitoBio/goquery v1.5.1 // indirect
	github.com/akamensky/argparse v1.2.2
	github.com/akrennmair/gopcap v0.0.0-20150728160502-00e11033259a
	github.com/antchfx/htmlquery v1.2.3 // indirect
	github.com/antchfx/xmlquery v1.2.4 // indirect
	github.com/armon/go-socks5 v0.0.0-20160902184237-e75332964ef5
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gocolly/colly v1.2.0
	github.com/google/gopacket v1.1.18
	github.com/google/uuid v1.1.0
	github.com/hzwy23/quicksort v0.0.0-20161206141632-56727a12b0b1
	github.com/ip2location/ip2location-go v8.2.0+incompatible
	github.com/ipipdotnet/ipdb-go v1.1.0
	github.com/ipplus360/awdb-golang v0.0.0-20200707075537-c4fdcab1c251
	github.com/jbenet/go-is-domain v1.0.5
	github.com/kennygrant/sanitize v1.2.4 // indirect
	github.com/klauspost/reedsolomon v1.9.9
	github.com/mmcloughlin/avo v0.0.0-20200803215136-443f81d77104 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca // indirect
	github.com/songgao/water v0.0.0-20190725173103-fd331bda3f4b
	github.com/stretchr/testify v1.6.1 // indirect
	github.com/temoto/robotstxt v1.1.1 // indirect
	github.com/templexxx/cpufeat v0.0.0-20180724012125-cef66df7f161 // indirect
	github.com/templexxx/xor v0.0.0-20191217153810-f85b25db303b // indirect
	github.com/tidwall/gjson v1.6.0
	github.com/tjfoc/gmsm v1.3.2 // indirect
	github.com/xtaci/kcp-go v5.4.20+incompatible
	github.com/xtaci/lossyconn v0.0.0-20200209145036-adba10fffc37 // indirect
	golang.org/x/crypto v0.0.0-20191219195013-becbf705a915
	golang.org/x/net v0.0.0-20200421231249-e086a090c8fd
	taas.com/atom v0.0.8
)

replace (
	github.com/rs/zerolog v1.11.0 => gitlab.lstaas.com/pros/zerolog v1.11.1
	taas.com/atom v0.0.8 => gitlab.lstaas.com/pros/atom v0.0.8
	taas.com/orbit v0.0.4 => gitlab.lstaas.com/pros/orbit v0.0.4
)

go 1.13
