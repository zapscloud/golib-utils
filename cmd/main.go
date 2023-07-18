package main

import (
	"fmt"

	"github.com/zapscloud/golib-utils/utils"
)

func main() {
	fmt.Println("Utils Module")

	fmt.Println("Guid ", utils.GenerateChecksumId("test", "Test sting for hash generation"))
	fmt.Println("Guid ", utils.GenerateUniqueId("test"))

	fmt.Println("ToTitle ", utils.ToTitle("this is test message"))

	str := " This  is  ,a - 	/	\n test  message "
	//str := "new001"
	keyId := utils.GenerateSeoKeyId(str)
	fmt.Println("prodId:", keyId)

	for i := 0; i < 10; i++ {
		// NextKeyId
		keyId = utils.GenerateNextKeyId(keyId)
		fmt.Println("NextKeyId:", keyId)
	}

}
