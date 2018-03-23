package main
import (
	"fmt"
	"GoDream/xasn1"
)

func main()  {
	fmt.Println("========= parse syspid:")
	sydpidDerHex :=  "3082040D020101040453444B5304820400CECA5781B340A8A033C82873879C1D43EBA1C6A9B328CF21EEE4F24ECD138B9DAD1F7BF16F4CB8EE0E6C3F76E3DDC158889C5F89299CBA066A0DA3E1FB39B2E53C20ED1452713DB798F368E8C7483F9F9EBAF827AA5AF50A5A35745A727D7B01923714091380A92DC9D166A279824B5E6AF544E53C9020B3486968165D4F657747715FB69FF4B31653FB8F1A116173B9AC8D207389D2816C85501406339839D0259E08F55684E3C6185E1D51C3D695B45E34E1E2E9EBF8EA7D12AB435D24F034AE29909566BB67827DBC1E007128B4525750712196628209362782566B52D34D3672C34CD413E83A9D5133F68F889BB810515FAA839D74CA4DCE8247F6ED81BBFF372B1E1D4D5FFD99C7FA4974998EC68922F499461B7963142A275EAB53D896D60C5A02F820E3862E0B6966246E6033E21430A528DC6D2CA0C9D7804C67B43FF3FF978942FD651185AFCC3B36C84F14D1CD3713B4E96C90E98AEA17EB7962B01FB739BA2F572F729C5B7877D0BCE5489E7D19D228798A021EB150A5D12BCE9AE4CC7110EF7DE32B8B08DF76E8D828DC23667C8D666C97E6FCF4A81D39847D9ED247F832BB1406ECCE341FA95B91696B88C470A1FC1517B4EA0998844A891C0397484F930D307AFB0B2A2039D890C1496286B31F08D9ED46CF36EF952AC12E0E15CC5BE15E51338026E4F9842D156417E27852DB86DED426D008300AFF0BB62D6AA31923CE58EB24D63CDABF2B254BAC5F1EC6EDE125C499A7E725C74F20F09CE6295CE5E871280064E1C74CAC3DBB874922841547C7EE12691680794CDC211A5C6D74DE3095410F15A9247058F78D0C42F46B30F3E8B774E439A32709AEF61DDA8923004886FF92D21F1AF7D4EB3642E8CDFA25DA349424A4CD257097FE4CEF83659B87BAD0A04841A2B4CBC36E37E3CCFED4A9899BA8310D709AE7954C3FB33FE55156A279A2B053BCE38613A6FADB3DA7F62DC25A1EEF03A10364BC2096A3750AE6B831AC62C1280353FA2529DBD3A628DBAD18F294ABDA807A54CD59457864145833AE66B8EECA4CB5DA0AEB2B45BD5F7572BE6C6876A7F1FC12CD74BA8C0159B7AD764A8FC78A6377783CD12F71C74C38489CA64847496C35C135C1A656A5882D0DBF6AEAFFD954137E66B69D571242D0337C79F2C5E30320757FE9945C53777DC52925C873A12B58950857C2940F681A0E9C044E81DB793211CF1F25AB7409BE832FF88ABFE49065F079D7A42DAD4A2CC09A593DFEA8A7F9C7E4A0C1A9A14CEE72C17BE932B9EDA5C8D96E761C21C98EE4DFAD7A5D0DB5AA4E410B5F4E8E223CB28934D06E1DE6F59E5D3A7599A80326AE7485D16F2881CA73E7F6A88F580312346FFCACB5F9D6FBF207BF3C7A62FFAD0E3A83C5BC1071B64F4DA53CE8EF7C36188C49E5324256824DA3DF8CF1C117A825A8E213EA66AA57082B95781D"
	xasn1.ParseSysPid(sydpidDerHex)

	fmt.Println("========= parse userpid:")
	userpidDerHex :=  "3081c002010180623060020101800453444b5304077e6b6773333031044c01049f8f74d09fe35ac8c6bb317d2fcb8633a27d4997f52c635d8a586488d5329c0493a75d885cb0b141d00035085dc6db66790e01efcfca7141bcc6d650ad251053170724270724f4f780520409353436353436353436044c0004752effe64f8ae607e10d160f4c5b5986d6d50c565ca06fdb184f703259b77e74290de1b455951e5dba316dc474372843abc6aef1a88072741e70b884ef0625ea170830220830c15fe59e";
	xasn1.ParseUserPid(userpidDerHex)
}