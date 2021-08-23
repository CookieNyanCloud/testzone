//+build unit

package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2+2 = 4")
	}
}
func TestTabCalculate(t *testing.T)  {
	var tests = []struct{
		input int
		exp int
	}{
		{2,4},
		{-1,1},
		{0,2},
		{99999,100001},
	}

	for _, test:= range tests {
		if out:= Calculate(test.input); out!=test.exp{
			t.Error("Failed: {} inp, {} exp, {} res", test.input,test.exp, out)
		}
	}
}
type AddResult struct {
	x int
	y int
	exp int
}
var addResult  = []AddResult{
	{1,1,2},
}
func TestAdd(t *testing.T) {
	for _, test:= range addResult{
		res:= Add(test.x,test.y)
		if res!=test.exp{
			t.Fatal("Expected not given")
		}
	}
}
func TestReadFile(t *testing.T)  {
	data, err:= ioutil.ReadFile("testdata/test.data")
	if err!=nil{
		t.Fatal("Could not open file")
	}
	if string(data) != "hello world"{
		t.Fatal("Strings do not match")
	}
}

func TestHttpRequest (t *testing.T){
	handler := func(w http.ResponseWriter, r *http.Request) {
		_,_=io.WriteString(w,"{\"status\":\"good\"")
	}

	req, _ := http.NewRequest("GET","http://tutorialedge.net", nil)
	w:= httptest.NewRecorder()
	handler(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	println(body)
	if resp.StatusCode != 200{
		t.Fatal("Status not ok")
	}
}