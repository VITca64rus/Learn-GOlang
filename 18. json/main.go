package main

import (
	"io/ioutil"
	"os"
	"encoding/json"
	"fmt"
	// "strconv"
)

func main() {
	type myStruct struct {
        Students []struct {
            Rating []float64 `json:"Rating"`
        } `json:"Students"`
    }
	data, _ := ioutil.ReadAll(os.Stdin)
	var group myStruct
	json.Unmarshal(data, &group)
	var res int
	for i:=0; i < len(group.Students); i++ {
		res += len(group.Students[i].Rating)
	}
	// fmt.Println(float64(res) / float64(len(group.Students)))
	r:=float64(res) / float64(len(group.Students))
	type res_s struct{
		Average float64
	}
	var result *res_s = new(res_s)
	result.Average = r
	data, _ = json.MarshalIndent(result, "", "    ")
	fmt.Printf("%s", data)
}