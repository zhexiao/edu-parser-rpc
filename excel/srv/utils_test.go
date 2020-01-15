package srv

import (
	"io/ioutil"
	"log"
	"testing"
)

func initExcel() *CT_Excel {
	return NewCT_Excel()
}

func TestExcelPaper(t *testing.T) {
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


func TestExcelOutline(t *testing.T) {
	fileBytes, err := ioutil.ReadFile("../_testdoc/outline-20190129.xlsx")
	if err != nil {
		t.Errorf("文件读取失败 err=%s", err)
	}

	jsonData, err := initExcel().outline(fileBytes)
	if err != nil {
		t.Errorf("文件解析失败 err=%s", err)
	}

	log.Println(jsonData)
}

func TestExcelBook(t *testing.T) {
	fileBytes, err := ioutil.ReadFile("../_testdoc/book.xlsx")
	if err != nil {
		t.Errorf("文件读取失败 err=%s", err)
	}

	jsonData, err := initExcel().book(fileBytes)
	if err != nil {
		t.Errorf("文件解析失败 err=%s", err)
	}

	log.Println(jsonData)
}

func TestExcelCognitionMap(t *testing.T) {
	fileBytes, err := ioutil.ReadFile("../_testdoc/cognition-map-20190129.xlsx")
	if err != nil {
		t.Errorf("文件读取失败 err=%s", err)
	}

	jsonData, err := initExcel().cognition_map(fileBytes)
	if err != nil {
		t.Errorf("文件解析失败 err=%s", err)
	}

	log.Println(jsonData)
}

func TestExcelCognitionSp(t *testing.T) {
	fileBytes, err := ioutil.ReadFile("../_testdoc/cognition-sp-20190301.xlsx")
	if err != nil {
		t.Errorf("文件读取失败 err=%s", err)
	}

	jsonData, err := initExcel().cognition_sp(fileBytes)
	if err != nil {
		t.Errorf("文件解析失败 err=%s", err)
	}

	log.Println(jsonData)
}
