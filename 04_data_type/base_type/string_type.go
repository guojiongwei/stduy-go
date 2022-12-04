package basetype

import (
	"fmt"
	"strings"
)

func StrigType() string {
	const name string = "小明"

	/** 字符串分割 */
	var line = strings.Join([]string{name, "哈哈", name}, "")

	fmt.Println("line", line)

	fmt.Println(`strings.Contains(name, "小")`, strings.Contains(name, "小"))
	fmt.Println(`strings.Index(name, "小")`, strings.Index(name, "小"))
	return name
}
