package srv

import (
	"io/ioutil"
	"log"
	"testing"
)

func initExcel() *CT_Excel {
	return &CT_Excel{}
}

func TestWordPaper(t *testing.T) {
	fileBytes, err := ioutil.ReadFile("../_testdoc/paper_20190702.xlsx")
	if err != nil {
		t.Errorf("文件读取失败 err=%s", err)
	}

	jsonData, err := initExcel().paper(fileBytes)
	if err != nil {
		t.Errorf("文件解析失败 err=%s", err)
	}

	log.Println(jsonData)
}
