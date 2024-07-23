package main

import (
	"bufio"
	"domainQuery/internal"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"sync"
)

// 读取域名列表文件
func readDomainsFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var domains []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domain := scanner.Text()
		if domain != "" {
			domains = append(domains, domain)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return domains, nil
}

// 工作线程函数
func worker(mywrite *csv.Writer, jobs <-chan string, dnsServer string, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	for domain := range jobs {
		internal.QueryDomianRecord(mywrite, domain, dnsServer, mu)
	}
}

// 处理域名
func processDomains(mywrite *csv.Writer, filePath string, numWorkers int, dnsServer string) error {
	domains, err := readDomainsFromFile(filePath)
	if err != nil {
		return err
	}

	jobs := make(chan string, len(domains))
	var wg sync.WaitGroup
	var mu sync.Mutex

	// 启动工作线程
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(mywrite, jobs, dnsServer, &wg, &mu)
	}

	// 发送任务
	for _, domain := range domains {
		jobs <- domain
	}
	close(jobs)

	// 等待所有工作线程完成
	wg.Wait()
	return nil
}

//func processDomains(mywrite *csv.Writer, filePath string) error {
//	domains, err := readDomainsFromFile(filePath)
//	if err != nil {
//		return err
//	}
//
//	var mu sync.Mutex
//	var wg sync.WaitGroup
//	for _, domain := range domains {
//		wg.Add(1)
//		go func(domain string) {
//			defer wg.Done()
//			internal.QueryDomianRecord(mywrite, domain, &mu)
//		}(domain)
//	}
//	wg.Wait()
//	return nil
//}

func main() {
	// 定义命令行参数
	inputFile := flag.String("f", "", "输入域名文件路径")
	dnsServer := flag.String("d", "223.6.6.6", "输入dns服务器")
	numWorkers := flag.Int("t", 10, "输入线程数")
	flag.Parse()

	if *inputFile == "" {
		fmt.Println("[-] 请使用 -f 参数指定输入文件")
		return
	}
	// 创建csv文件的write
	mywrite, file := internal.CreateCSV(*inputFile + "_result.csv")
	defer file.Close()
	err := processDomains(mywrite, *inputFile, *numWorkers, *dnsServer+":53")
	if err != nil {
		return
	}
	internal.Sorted(*inputFile+"_result.csv", *inputFile+"_result2.csv")
	fmt.Println("[+] 解析完成，解析结果输出: " + *inputFile + "_result2.csv")
}
