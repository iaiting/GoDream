package main

import (
	"GoDream/cms/recipient"
	"encoding/hex"
	"fmt"
)

func GetSerialNumberFromDerBytesTest(certDerHex string) {
	fmt.Println("== Start Testing: GetSerialNumberFromDerBytes")
	bs, _ := hex.DecodeString(certDerHex)
	rv, _ := recipient.GetSerialNumberFromDerBytes(bs)
	fmt.Print(rv,"\012\012")
}

func GetSerialNumberFromDerHexTest(certDerHex string) {
	fmt.Println("== Start Testing: GetSerialNumberFromDerHex")
	rv , _ := recipient.GetSerialNumberFromDerHex(certDerHex)
	fmt.Print(rv,"\012\012")
}


func GetIssuerNameFromDerBytesTest(certDerHex string) {
	fmt.Println("== Start Testing: GetIssuerNameFromDerBytes")
	bs, _ := hex.DecodeString(certDerHex)
	rv, _ := recipient.GetIssuerNameFromDerBytes(bs)
	fmt.Print(rv,"\012\012")
}

func GetIssuerNameFromDerHexTest(certDerHex string) {
	fmt.Println("== Start Testing: GetIssuerNameFromDerHex")
	rv, _ := recipient.GetIssuerNameFromDerHex(certDerHex)
	fmt.Print(rv,"\012\012")
}

func GetSubjectNameFromDerBytesTest(certDerHex string) {
	fmt.Print("\012\012")
}

func GetSubjectNameFromDerHexTest(certDerHex string) {
	fmt.Println("== Start Testing: GetSubjectNameFromDerHex")
	rv, _ := recipient.GetSubjectNameFromDerHex(certDerHex)
	fmt.Print(rv,"\012\012")
}


func TEST_RI_GettypeFromDerhex() {
	rsa_riDerhex := "308203643082024CA003020102020857FAC3C581FA9C07300D06092A864886F70D01010B0500302D310F300D06035504030C0658594A414341310D300B060355040A0C0458594A41310B300906035504061302434E301E170D3136313131353032353932345A170D3138313131353032353932345A307C3120301E06092A864886F70D0109011611313333343732363036324071712E636F6D3112301006035504030C09E99988E6989FE5AE8731153013060355040B0C0CE4BFA1E6BA90E4B985E5AE89310F300D06035504070C06E58D97E4BAAC310F300D06035504080C06E6B19FE88B8F310B300906035504061302434E30819F300D06092A864886F70D010101050003818D0030818902818100879D1A10899A82A9A0C99AFCF0F24788143A0F91325E478757F34060192B6EB5190DDE783CCBADC83B11E7285A2734F871C17D25E95EF7F913F39DB8061ABCEF462D4823BF5BE7E9360C20999ED49F7B1B36CA5E2BEB6F70D0BF786F3009C09CCCBF393D83ABB0020FDD79CC45A745B05E381E1CBA96F7066DD1FEAD875C068D0203010001A381BC3081B9301D0603551D0E04160414D96BE03E2D4FF86395F2C008416F4F7E515DCD97300C0603551D130101FF04023000301F0603551D23041830168014A002E381E6CDDEC42874917BB2F87D8C03E47D30300E0603551D0F0101FF0404030204F0303B0603551D250434303206082B0601050507030106082B0601050507030206082B0601050507030306082B0601050507030406082B06010505070308301C0603551D11041530138111313333343732363036324071712E636F6D300D06092A864886F70D01010B05000382010100AE1421AAE7B39C8448FFA26E756D1B83AE7444C87D3838D05E245C6F716E1603F0F0C258002F431D2FC600D6BF0D80D70BA833AEDBDB7880EDDEBA8EC7B866427E0F448A56BC6B2C4DF915B33680319FBC6AB52F057CCB60B453DC774701EC6FB6B03BE05F846E78C2BF95CDB12C3ED47EC10CD389EFB20AB72AFC3E2FDFD8FF8DB64C3AA0648ED43B72CE3D9EA91997AB2CCC54E8CD63AFC532921DEC4D70719B16DE2CE301CE75CE1249E4646EF1418F449F309CBFEB05EB297E751258E539A674DA60028DAF294ABBAE3E98E53E90CA62F39F5233B18CC4F169F645C8261586F04A16BF60B4CF3406BC262BFC1D6972E30B38C95A6FD968DB1410FF34EC47";
	recipient.RI_GettypeFromDerhex(rsa_riDerhex)

	sm2_riDerhex := "308202aa30820255a0030201020203123123300c06082a811ccf5501837505003081a7312a302806035504030c214c4d20534d3220526f6f7420436572746966696361746520417574686f726974793120301e06092a864886f70d0109010c116c6d406c6f6e676d61692e636f6d2e636e31143012060355040b0c0b646576656c6f706d656e743110300e060355040a0c076c6f6e676d61693110300e06035504070c074265696a696e673110300e06035504080c074265696a696e67310b300906035504060c02434e301e170d3134303932333034303030305a170d3135303932333034303030305a308192310d300b06035504030c04746573743126302406092a864886f70d0109010c176d6963726f646f6e65406d6963726f646f6e652e636f6d31143012060355040b0c0b646576656c6f706d656e7431123010060355040a0c096d6963726f646f6e653110300e06035504070c074265696a696e673110300e06035504080c074265696a696e67310b300906035504060c02434e3053300d06092a811ccf5501822d030500034200046549b2e18db0903970ac2aeb477e193ffe293225fc6397b465dbe1252fb5c8f78e687495c6c3525d3887121dbdf315cf7a31ad97f59aae729883f24e59659967a3818630818330250603551d11041e301c811a6d6963726f646f6e65406d6963726f646f6e652e636f6d2e636e302c0603551d1f042530233021a01fa01d861b687474703a2f2f31302e332e34332e37332f7465737463616d732f300b0603551d0f040403020338301f0603551d250418301606082b06010505070302060a2b060104018237140202300c06082a811ccf550183750500034100b58f1382820ed8d5faa74a4ba9743e8cd47b7d2e76408823ea7b49bfb19f2f05b66dccfab7c76e8e3f3365c8ff0f2ca58ef7c4847a8839c605d63557dc1ffb03";
	recipient.RI_GettypeFromDerhex(sm2_riDerhex)
}




////////////////////////////////////////////////////////////////////////////////////////////////////
func main()  {
	certDerHex := "308203643082024CA003020102020857FAC3C581FA9C07300D06092A864886F70D01010B0500302D310F300D06035504030C0658594A414341310D300B060355040A0C0458594A41310B300906035504061302434E301E170D3136313131353032353932345A170D3138313131353032353932345A307C3120301E06092A864886F70D0109011611313333343732363036324071712E636F6D3112301006035504030C09E99988E6989FE5AE8731153013060355040B0C0CE4BFA1E6BA90E4B985E5AE89310F300D06035504070C06E58D97E4BAAC310F300D06035504080C06E6B19FE88B8F310B300906035504061302434E30819F300D06092A864886F70D010101050003818D0030818902818100879D1A10899A82A9A0C99AFCF0F24788143A0F91325E478757F34060192B6EB5190DDE783CCBADC83B11E7285A2734F871C17D25E95EF7F913F39DB8061ABCEF462D4823BF5BE7E9360C20999ED49F7B1B36CA5E2BEB6F70D0BF786F3009C09CCCBF393D83ABB0020FDD79CC45A745B05E381E1CBA96F7066DD1FEAD875C068D0203010001A381BC3081B9301D0603551D0E04160414D96BE03E2D4FF86395F2C008416F4F7E515DCD97300C0603551D130101FF04023000301F0603551D23041830168014A002E381E6CDDEC42874917BB2F87D8C03E47D30300E0603551D0F0101FF0404030204F0303B0603551D250434303206082B0601050507030106082B0601050507030206082B0601050507030306082B0601050507030406082B06010505070308301C0603551D11041530138111313333343732363036324071712E636F6D300D06092A864886F70D01010B05000382010100AE1421AAE7B39C8448FFA26E756D1B83AE7444C87D3838D05E245C6F716E1603F0F0C258002F431D2FC600D6BF0D80D70BA833AEDBDB7880EDDEBA8EC7B866427E0F448A56BC6B2C4DF915B33680319FBC6AB52F057CCB60B453DC774701EC6FB6B03BE05F846E78C2BF95CDB12C3ED47EC10CD389EFB20AB72AFC3E2FDFD8FF8DB64C3AA0648ED43B72CE3D9EA91997AB2CCC54E8CD63AFC532921DEC4D70719B16DE2CE301CE75CE1249E4646EF1418F449F309CBFEB05EB297E751258E539A674DA60028DAF294ABBAE3E98E53E90CA62F39F5233B18CC4F169F645C8261586F04A16BF60B4CF3406BC262BFC1D6972E30B38C95A6FD968DB1410FF34EC47";
	////////////////////////////////////////////////////////////////////////////////////////////////////
	GetSerialNumberFromDerBytesTest(certDerHex)

	////////////////////////////////////////////////////////////////////////////////////////////////////
	GetSerialNumberFromDerHexTest(certDerHex)

	////////////////////////////////////////////////////////////////////////////////////////////////////
	GetIssuerNameFromDerBytesTest(certDerHex)
	
	////////////////////////////////////////////////////////////////////////////////////////////////////
	GetIssuerNameFromDerHexTest(certDerHex)

	GetSubjectNameFromDerBytesTest(certDerHex)
	GetSubjectNameFromDerHexTest(certDerHex)


	TEST_RI_GettypeFromDerhex()
}