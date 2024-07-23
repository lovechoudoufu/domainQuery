package internal

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"
)

// 创建writer
func CreateCSV(filename string) (*csv.Writer, *os.File) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()
	// 写入CSV文件标题行
	header := []string{"domain", "cname", "ip", "region", "owner"}
	err = writer.Write(header)
	if err != nil {
		panic(err)
	}
	return writer, file
}

// 结果写入csv文件
func WriteCSV(writer *csv.Writer, domain string, cname string, ip string, region string, owner string, mu *sync.Mutex) {
	row := []string{domain, cname, ip, region, owner}
	mu.Lock()
	defer mu.Unlock()
	if err := writer.Write(row); err != nil {
		fmt.Println("[-] 写入数据失败: %v", err)
	}
	writer.Flush() // 确保所有数据都写入文件
	if err := writer.Error(); err != nil {
		fmt.Println("[-] 保存数据失败: %v", err)
	}
	fmt.Println("[+] " + domain)
}
