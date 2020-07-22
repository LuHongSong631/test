//这是我新添加的内容哦
package main
import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

func main(){
	info :=[]Website{{"Golang","http://c.biancheng.net/golang/",[]string{"http://c.biancheng.net/cplus/","http://c.biancheng.net/linux_tutorial/"}},
	{"Java","http://c.biancheng.net/java",[]string{"http://c.biancheng.net/socket/","http://c.biancheng.net/python/"}}}

	//创建文件
	filePtr,err := os.Create("info.json")
	if err != nil{
		fmt.Println("文件创建失败",err.Error())
		return
	}
	 filePtr.Close()//我也对二十行做了修改，将defer去掉了
	//创建Json编码器
	encoder := json.NewEncoder(filePtr)

	err = encoder.Encode(info)
	if err != nil{
		fmt.Println("编码错误", err.Error())
	}else{
		fmt.Println("编码成功")
	}

	//读json文件
	filePtr1,err1 := os.Open("./info.json")
	if err1 != nil{
		fmt.Println("文件打开失败[Err:%s]",err1.Error())
		return
	}
	defer filePtr1.Close()
	var info1 []Website

	//创建json解码器
	decoder1 := json.NewDecoder(filePtr1)
	err1 = decoder1.Decode(&info1)
	if err1 != nil{
		fmt.Println("解码失败", err1.Error())
	}else{
		fmt.Println("解码成功")
		fmt.Println(info1)
	}

	//写xml文件
	//实例化对象
	info2 := Website{"C语言中文网","http://c.biancheng.net/golang/",[]string{"go语言入门教程","golang入门教程"}}
	f,err2 := os.Create("./info2.xml")
	if err2 != nil {
		fmt.Println("文件创建失败", err2.Error())
		return
	}
	defer f.Close()
	//序列化到文件中
	encoder3 := xml.NewEncoder(f)
	err2 = encoder3.Encode(info2)
	if err2 != nil {
		fmt.Println("编码错误：", err2.Error())
		return
	} else {
		fmt.Println("编码成功")
	}

	//读取xml文件
	file4,err4 := os.Open("./info2.xml")
	if err4 != nil {
		fmt.Printf("文件打开失败：%v", err4)
		return
	}
	defer file4.Close()

	//info4 := Website{}
	//创建xml解析器
	decoder4 := xml.NewDecoder(file4)
	err4 = decoder4.Decode(&info2)
	if err4 != nil {
		fmt.Printf("解码失败：%v", err4)
		return
	} else {
		fmt.Println("解码成功")
		fmt.Println(info2)

	}
    fmt.Println("==============================")
	fmt.Sprintf("%s","bedg")
	fmt.Print("WWWW")
	fmt.Println(fmt.Sprint("abc","dec"))
	fmt.Printf("aaaaa%s","nnnnn")
	fmt.Println(fmt.Sprintln(12445))

	
	n,err:=fmt.Scan("是否能看到看对方")

	fmt.Println(n,err)

	name :=&A{
		2,
		"2",
		AB{
			O: 2,
		},
		&AB{
			O:3,
		},
	}
	fmt.Println(name)






}
type Website struct{
	Name string `xml:"name,attr"`
	Url string
	Course []string
}
type A struct {
	X int
	Y string
	Z AB
	K *AB
}

type AB struct {
	O int
}
