package internal

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
)

// 读取CSV文件并返回记录
func readCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

// 按domain列对记录进行排序
func sortRecords(records [][]string) [][]string {
	if len(records) <= 1 {
		return records
	}

	header := records[0]
	rows := records[1:]

	sort.Slice(rows, func(i, j int) bool {
		return rows[i][0] < rows[j][0] // 假设domain列是第一列
	})

	sortedRecords := append([][]string{header}, rows...)
	return sortedRecords
}

// 将记录写入新的CSV文件
func writeCSV(filePath string, records [][]string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range records {
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return writer.Error()
}

// 表格排序，以domain为列
func Sorted(inputFilePath string, outputFilePath string) {
	records, err := readCSV(inputFilePath)
	if err != nil {
		log.Fatalf("[-] 表格读取失败 : %v", err)
	}
	sortedRecords := sortRecords(records)
	if err := writeCSV(outputFilePath, sortedRecords); err != nil {
		log.Fatalf("[-] 表格排序失败 : %v", err)
	} else {
		err := os.Remove(inputFilePath)
		if err != nil {
			return
		}
	}
}
