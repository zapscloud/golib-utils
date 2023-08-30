package main

import (
	"fmt"
	"log"

	"github.com/zapscloud/golib-utils/utils"
)

func main() {
	fmt.Println("Utils Module")

	testGuid()
	testSeoKeyId()
	testGenerateNextKeyId()

	testMapUtils()

}

func testGuid() {
	fmt.Println("Guid ", utils.GenerateChecksumId("test", "Test sting for hash generation"))
	fmt.Println("Guid ", utils.GenerateUniqueId("test"))

}

func testSeoKeyId() {
	str := " This  is  ,a - 	/	\n test  message "
	//str := "new001"
	keyId := utils.GenerateSeoKeyId(str)
	fmt.Println("prodId:", keyId)
}

func testGenerateNextKeyId() {
	var keyId string

	for i := 0; i < 10; i++ {
		// NextKeyId
		keyId = utils.GenerateNextKeyId(keyId)
		fmt.Println("NextKeyId:", keyId)
	}
}

func testMapUtils() {
	test := utils.Map{
		//"test_string": "Test-String",
		//"test_string": "",
		"test_string": []string{"Test-String"},
		"test_int":    10.0,
	}
	test1 := map[string]interface{}{
		"test": "string",
	}

	val, err := utils.GetMemberDataInt(test, "test_int", true)

	log.Println("Result=>", val, err, utils.IsTypeMap(test), utils.IsTypeMapStrInterface(test1))
}
