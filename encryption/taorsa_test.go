package encryption

import (
	"encoding/json"
	"fmt"
	"testing"
)

type KeyData struct {
	SendKey    byte
	ReceiveKey byte
}

// Test RsaEncrypt
func TestRsaEncrypt(t *testing.T) {

	//key, err := LoadPublicKeyFromFile("../conf/rsa/public.key")
	key, err := RsaLoadPublicKeyBase64("MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC5Jk+Qi85rNZAwkEh2FzbwUbt9Ifumm6Ob+erbI/fWC1nq+fVW5mvSSN+uktmKpV5l5Qgh8V2Y7ttjqqX+qCAHFm7TmM+t0kdGVnAVpHu3gMOnk752BsALJadbDd9C4+bk4sWeKFUgr4VisWz7FmhBiQBB+DpX/MynemnyTKb+QQIDAQAB")
	if err != nil {
		fmt.Println("LoadPublicKeyBase64 error", err)
		t.Fail()
		return
	}
	str, err := RsaEncrypt("123", key)
	if err != nil {
		fmt.Println("RsaEncrypt error", err)
		t.Fail()
		return
	}
	fmt.Println("加密", str)

	//初始化KeyData
	keyData := KeyData{
		SendKey:    15,
		ReceiveKey: 16,
	}

	// 序列化keyData成json
	jsonByte, err := json.Marshal(keyData)

	// 将byte[]转换成string
	jsonStr := string(jsonByte)
	str2, err := RsaEncrypt(jsonStr, key)
	if err != nil {
		fmt.Println("RsaEncrypt error", err)
		t.Fail()
		return
	}
	fmt.Println("加密2", str2)
}

func TestRsaDecrypt(t *testing.T) {
	//key, err := LoadPublicKeyFromFile("../conf/rsa/public.key")
	key, err := RsaLoadPrivateKeyBase64("MIICXQIBAAKBgQC5Jk+Qi85rNZAwkEh2FzbwUbt9Ifumm6Ob+erbI/fWC1nq+fVW5mvSSN+uktmKpV5l5Qgh8V2Y7ttjqqX+qCAHFm7TmM+t0kdGVnAVpHu3gMOnk752BsALJadbDd9C4+bk4sWeKFUgr4VisWz7FmhBiQBB+DpX/MynemnyTKb+QQIDAQABAoGABAjz9HI7jCYVb52BWsN8QMYDOEGolpOvP9u1NNPml41VRxPt9xgb8vAAIfGU2JrsgPt6DyVPtSDFerlSF0fqNLvWrG3p7wjXQA8D1HaAk2orf7bRIi3OBsJ9P8XQVYbnY4n3240VNW2Nb2fHlbgJGeekloYBMNZmIwLACJEuUHECQQDQlVL7IdH3sx/6fh4PBHfdB5/xZDKSefPzZoR6LoHH/SeorPwPwAbnv1ISX54oWS2ZnwHlDsitdus4sWKGau+pAkEA4z0/FoL8HNi6L1WVetd4gRydx+5Fnbsq9GntSMstDwXYaFBfm8jauwcAlemmTqgo6EGhHe5VLNX0BU8ch3oY2QJARu3gcAHKMt84yqfEdPrh/8mt/BpEkEbkTCradeoGvAk8SUG53WlBb+FBeXoGgFYDCbBmpovmdgZwarD3fhozMQJBAIVvCmOaqJX7wWBYHf1TFyShfBjRVjmnlTKOHNJ4082VfhzKzUl56M6X7wUYfRqE7fhryUL9FzDLtY2EhdkbngkCQQDN4A3I5NaqB0I78CDRRZJ5Ps0qlgM8PEuukQcwPLdrOISgtKvEeqtfjORXt/lzcxAALC+b8r/fh4c0hCv5KfVN")
	if err != nil {
		fmt.Println("LoadPrivateKeyBase64 error", err)
		t.Fail()
		return
	}
	str, err := RsaDecrypt("KTaWBe/MdMUe50pUdrFJ3bXTaENJQHbKfjmEwkejpe9LezVeMJAmiv8aqG5tcOkU7Sd3KUhZasZ+2RMJHfPq2483JVrRZFQhzijyvVY6v5+bCbUetK3m1Eli8Ud3GQtKKJIqKNcZYZfylnKWGW1IcNAa7x30pWdBveM6eKHkesk=", key)
	if err != nil {
		fmt.Println("RsaEncrypt error", err)
		t.Fail()
		return
	}
	fmt.Println("解密", str)
}
