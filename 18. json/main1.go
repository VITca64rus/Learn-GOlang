package main

import (
	"io/ioutil"
	"os"
	"encoding/json"
	"fmt"
	// "strconv"
)

func main() {
	type GlobalIds []struct {
		Id int `json:"global_id"`
	}
	var globalIds GlobalIds
	var res int
	buf, _ := os.Open("data-20190514T0100.json")
	data, _ := ioutil.ReadAll(buf)
	json.Unmarshal(data, &globalIds)
	for i:=0; i < len(globalIds); i++ {
		res += globalIds[i].Id
	}
	fmt.Println(res)
}