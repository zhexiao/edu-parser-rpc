package word

import (
	"github.com/zhexiao/office-parser/utils"
	"io/ioutil"
	"log"
	"testing"
)

func initWord() *CT_Word {
	return NewCT_Word()
}

func initDemoQiniu() {
	//初始化七牛的配置
	utils.OfficeParserQiniuCfg = &utils.Qiniu{
		//七牛key
		AccessKey: "test",
		//七牛secret
		SecretKey: "test",
		//七牛存储的bucket
		Bucket: "test",
		//所属区域
		Zone: "ZoneHuanan",
		//访问的域名地址
		Domain: "test",
		//路径
		UploadPrefix: "test",
	}
}

func TestWordPaper(t *testing.T) {
	fileBytes, err := ioutil.ReadFile("../_testdoc/paper.docx")
	if err != nil {
		t.Errorf("文件读取失败 err=%s", err)
	}

	jsonData, err := initWord().paper(fileBytes)
	if err != nil {
		t.Errorf("文件解析失败 err=%s", err)
	}

	log.Println(jsonData)
}

func TestWordQuestion(t *testing.T) {
	initDemoQiniu()

	fileBytes, err := ioutil.ReadFile("../_testdoc/question-fill.docx")
	if err != nil {
		t.Errorf("文件读取失败 err=%s", err)
	}

	jsonData, err := initWord().question(fileBytes)
	if err != nil {
		t.Errorf("文件解析失败 err=%s", err)
	}

	log.Println(jsonData)
}
