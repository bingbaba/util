package httptool

import (
	"testing"
)

type TestJsonrResp struct {
	Order int `json:"targetOrder"`
}

func TestRespJsonr(t *testing.T) {
	respbody := []byte(`**YGKJ{"jsonr":{"data":{"targetOrder":39}}}YGKJ##`)
	r := &TestJsonrResp{}

	err := respjsonr(respbody, r)
	if err != nil {
		t.Fatal(err)
	}

}
