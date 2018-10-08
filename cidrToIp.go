package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	minIp, maxIp := getCidrIpRange("1.64.0.0/11")
	fmt.Println("CIDR最小IP：", minIp, " CIDR最大IP：", maxIp)
}

func getCidrIpRange(cidr string) (string,string){
	cidrArr := strings.Split(cidr,"/")
    ip := cidrArr[0]
    ipSegs := strings.Split(ip,".")
    maskLen,_ := strconv.Atoi(cidrArr[1])
	seg1MinIp,seg1MaxIp := getIpSeg1Range(ipSegs,maskLen)
	seg2MinIp,seg2MaxIp := getIpSeg2Range(ipSegs,maskLen)
	seg3MinIp,seg3MaxIp := getIpSeg3Range(ipSegs,maskLen)
	seg4MinIp,seg4MaxIp := getIpSeg4Range(ipSegs,maskLen)

	return fmt.Sprint(seg1MinIp,".",seg2MinIp,".",seg3MinIp,".",seg4MinIp),fmt.Sprint(seg1MaxIp,".",seg2MaxIp,".",seg3MaxIp,".",seg4MaxIp)
}

//得到第一段IP的区间
func getIpSeg1Range(ipSegs []string, maskLen int)(int,int){
	segIp,_ := strconv.Atoi(ipSegs[2])
	if maskLen>8{
		return segIp,segIp
	}

	return getIpSegRange(uint8(segIp),uint8(8-maskLen))
}

//得到第二段IP的区间
func getIpSeg2Range(ipSegs []string, maskLen int)(int,int){
	segIp,_ := strconv.Atoi(ipSegs[1])
	if maskLen>16{
		return segIp,segIp
	}

	return getIpSegRange(uint8(segIp),uint8(16-maskLen))
}

//得到第三段IP的区间
func getIpSeg3Range(ipSegs []string, maskLen int)(int,int){
	segIp,_ := strconv.Atoi(ipSegs[2])
	if maskLen>24{
		return segIp,segIp
	}

return getIpSegRange(uint8(segIp),uint8(24-maskLen))
}

//得到第四段IP的区间（第一片段.第二片段.第三片段.第四片段）
func getIpSeg4Range(ipSegs []string, maskLen int) (int,int) {
	if maskLen<=24{
		return 0,254
	}
	ipSeg,_ := strconv.Atoi(ipSegs[3])
	segMinIp,segMaxIp := getIpSegRange(uint8(ipSeg), uint8(32-maskLen))
	if maskLen>24{
		segMaxIp--
	}
	return segMinIp+1,segMaxIp
}


//根据用户输入的基础IP地址和CIDR掩码计算一个IP片段的区间
func getIpSegRange(segIp,offset uint8)(int,int){
	var ipSegMax uint8 = 255
	netSegIp := ipSegMax << offset
	segMinIp := netSegIp & segIp
	segMaxIp := segIp &(255<<offset) | ^(255<<offset)

	return int(segMinIp), int(segMaxIp)
}
