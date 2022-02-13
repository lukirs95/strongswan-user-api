package configparser

import (
	"os"
	"testing"
)

func TestConfigparser_Deserialize(t *testing.T) {
	file, _ := os.OpenFile("../../ipsec.secrets", os.O_RDONLY, os.ModePerm)
	defer file.Close()
	_, err := Deserialize(file)
	if err != nil {
		t.Log("Deserialize failed.", err)
		t.Fail()
	}
}
