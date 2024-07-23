package internal

import (
	"encoding/csv"
	"github.com/xiaoqidun/qqwry"
	"sync"
)

// 通过qqwry查询IPv4地址归属
func Search(domain string, cname string, writer *csv.Writer, IP string, mu *sync.Mutex) {
	// 从文件加载IP数据库
	if err := qqwry.LoadFile("./configs/qqwry.dat"); err != nil {
		panic(err)
	}
	// 从内存或缓存查询IP
	location, err := qqwry.QueryIP(IP)
	if err != nil {
		//fmt.Printf("错误：%v\n", err)
		WriteCSV(writer, domain, cname, IP, "", "", mu)
		return
	}
	region := location.Country + location.Province + location.City + location.District
	WriteCSV(writer, domain, cname, IP, region, location.ISP, mu)

}
