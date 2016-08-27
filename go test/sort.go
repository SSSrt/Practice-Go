package main

import (
	"fmt"
	"sort"
)

func main() {
	drinks := map[string]string{"cola": "kele", "orangejuice": "chengzhi", "tea": "cha"}
	kk := make([]string, len(drinks))
	i := 0
	for key, _ := range drinks {
		kk[i] = key
		i++
	}
	sort.Strings(kk)
	for _, k := range kk {
		fmt.Printf("the English name is %s,and the Chinese name is %s\n", k, drinks[k])
	}
}
