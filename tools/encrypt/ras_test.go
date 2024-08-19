package encrypt

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestGenRsaKey(t *testing.T) {
	bits := 2048
	path := "./"
	if err := GenRsaKey(bits, path); err != nil {
		p("秘钥生成失败")
	}
	p("秘钥生成成功")
}

func TestRsaEnc(t *testing.T) {
	// 读取公钥私钥
	var err error
	PublicKey, err = ioutil.ReadFile("public.pem")
	if err != nil {
		t.Fatal(err)
	}
	str := ""
	str = `{"level":"info","name":"lucifer"}`
	str = `{"uuid":"123"}`
	str = `{"bundleid":"111","uid":"222","uuid":"333"}`
	// str = `{"uuid":"CE6454A3-2745-457E-A28C-A2FB616E41FB","uid":"VO8YYE","bundleid":"com.wuyoulin.yw.tc","package_version":"1.0.0","sys":"IOS","sys_version":"14.4","phone_model":"iPhone10,3","sim_status":"1","mac_addr":"54:33:cb:72:42:27"}`
	str = `{"uuid":"32位uuid","uid":"6位uid","bundleid":"com.wuyoulin.yw.tc","platform_name":"巨量","ad_local":"banner"}`

	ciphetextByte, err := RsaEncrypt([]byte(str))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ciphetextByte)

	ciphetextStr, err := RsaEncryptString(str)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ciphetextStr)
}

func TestRsDec(t *testing.T) {
	// 读取公钥私钥
	var err error

	PrivateKey, err = ioutil.ReadFile("private.pem")
	if err != nil {
		os.Exit(-1)
	}
	// fmt.Println(PrivateKey)

	// ciphetextStr := `Ai4ITKJFxT27Bn/nf63mmtCs1mTHcfrpc+uhEf922dEL/b442zUPjKxfiz3o/5nw3GgHbHfUPyHJzphRvx9QCi8QyHqz4JAJ+SDmqqAGaZjzBHRXXbKTZsEeL2CLsTeihqs+joTjwNF16f7knVyELh+PARkgJNu3CFYOcjwuKtsxeLDa23vmw+IO9JDyRjBpkI/1iT+eJcjtSO++6Zu7meiXGSVcMq2jE3AvVMkOQcocPyyQpCarN4cUkOb9PRRbya8hVYe7l/aMC3dobCiRFNdlei2EumF8JOcF0WOk1mLTav6Zu5UeL0jowzcaUOlR4WZ+ZYUjsddLIonPtuB1Sg==`
	ciphetextStr := `O1bwsuOHFnZctiJoQ4FLPYMqxspNfwWCt5QJriCQw99Jv3Q7c6LsOKyGoo1mXYtuIMzW/UjEu9/yeQa06mX2PcFJLT4B48sWM/qXYEl7ZHjbhhVnvOEFexugcuZvtBs4LpW7Ru0yhVuIai7b+Twy8miJwnbRAeHjK3XXZKiuQPamo7k3ZoxOLVcMIKKwMxoU8hQw9UqWiYib1ZvOjLGhnge271hoDO050IWQQXfKcO/JjNnt+z/NplURT43exIiIqithv1BK2HM++mW3K9LIcvuPE22ljamkf1Ie7c3LVrpGQiQ7WKZQs9JTX/BEEAUI7G9oM0brBrkyohb+L8OLBg==`
	// ciphetextStr := `BjdSeEKR+osU1+DZabUBeBy/yO2AK0ME7XLNfjVjuVgUU8oRr/d1yhAcw2OG8NKWZS8zVUY0JzMayttHAi6niyA46CQtusaDzwtoUDGjpk/W6cZ31TsqXTlVj8xtwq7vGAFLCMrB0AZVmLAICPHOQ70kBZabYeCmKsiO/s9jRcwM9KKp5lPp+VN1bQgoGeyRKCon1E72NKfR0qxHk4nCC5hcHapRZujq/mymbVUnyWkzPpIybVMDddPyDT4s9+IWsWkFgaxKpgNMl16RX4OZCXlyx7hlyRRbU2z3iksyrq1K+hypErsYVxBswF85iAa7kfckI6Qlf1LXAlQjsnSBgA==`

	ciphetextByte := Base64DecodeToByte(ciphetextStr)
	plaintextByte, err := RsaDecrypt(ciphetextByte)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(plaintextByte)
	fmt.Println(string(plaintextByte))

	plaintextStr, err := RsaDecryptString(ciphetextStr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(plaintextStr)
}
