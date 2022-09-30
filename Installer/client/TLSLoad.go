package main

import "os"
import "encoding/base64"
import "fmt"

func PEMDecoder( PemFileBase64Str string ) []uint8 {
    PEMFileObj ,_ := base64.RawStdEncoding.DecodeString( PemFileBase64Str )
    return PEMFileObj
}

func PEMEncoder( PemFileName string ) string {
    PEMFileObj , _ := os.ReadFile( PemFileName )
    PEMFileObjBase64Str := base64.RawStdEncoding.EncodeToString( PEMFileObj )
    return PEMFileObjBase64Str
}

func main() {
    CertObjBase64 := PEMEncoder( "Client-TLS-ED25519.crt" )
    KeyObjBase64 := PEMEncoder( "Client-TLS-ED25519.key" )
    CertObj := PEMDecoder( CertObjBase64 )
    KeyObj := PEMDecoder( KeyObjBase64 )
    fmt.Println( CertObj )
    fmt.Println( KeyObj )
}
