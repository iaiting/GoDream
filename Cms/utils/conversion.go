package utils


import
 (
    "encoding/hex"
    "fmt"
    "strings"
)

func BYTE_Convert2hexlower_0(bytes []byte) string {
    return strings.ToLower(hex.EncodeToString(bytes))
}

func BYTE_Convert2hexlower_1(bytes []byte) string {
    hl := fmt.Sprintf("%x", bytes)
    return hl
}

func BYTE_Convert2hexupper_0(bytes []byte) string {
    return strings.ToUpper(hex.EncodeToString(bytes))
}

func BYTE_Convert2hexupper_1(bytes []byte) string {
    hu := fmt.Sprintf("%X", bytes)
    return hu
}

func HEX_Convert2byte(hexstr string) ([]byte, error) {
    return  hex.DecodeString(hexstr)
}
