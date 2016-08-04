package httptool

import (
	"testing"
)

type TestJsonrResp struct {
	Data struct {
		Order int `json:"targetOrder"`
	} `json:"data"`
}

func TestRespJsonr(t *testing.T) {
	respbody := []byte(`**YGKJ{"jsonr":{"targetOrder":39}}YGKJ##`)
	r := &TestJsonrResp{}

	err := respjsonr(respbody, r)
	if err != nil {
		t.Fatal(err)
	} else if r.Order != 39 {
		t.Fatalf("expect the targetorder 39,but get %d", r.Order)
	}

}
