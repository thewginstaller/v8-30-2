package main

import "fmt"
import "encoding/base64"
import "net/url"
import "crypto/rand"
import "crypto/sha256"
import "net/http"
import "io"
import "encoding/json"

func Base64URL( Data []byte ) string {
    return base64.RawURLEncoding.EncodeToString( Data )
}
func Sha256( Data []byte ) []uint8 {
    Sha256 := sha256.New()
    Sha256.Write( Data )
    return Sha256.Sum( nil )
}
func UserName( username string ) string {
    return Base64URL( []byte( username ) )
}
func Password( password string ) string {
    return Base64URL( Sha256( []byte( password  ) ) )
}
func SessionID( ) string {
    RandomBytes := make( []byte , 32 )
    _ , _ = rand.Read( RandomBytes )
    return Base64URL( RandomBytes )
}
func URLValues( QueryMap map[string]string ) string {
    URLValuesObj := url.Values{}
    for QueryMapKey , QueryMapValue := range QueryMap {
        URLValuesObj.Add( QueryMapKey , QueryMapValue )
    }
    return URLValuesObj.Encode()
}
func URLCreds( URLUserName string , URLPassword string ) *url.Userinfo {
    return url.UserPassword( UserName( URLUserName ) , Password( URLPassword ) )
}
func ApiIP( ApiConfigURL string ) string {
    ConfigServerResponse , _ := http.Get( ApiConfigURL )
    defer ConfigServerResponse.Body.Close()
    ApiConfigObj , _ := io.ReadAll( ConfigServerResponse.Body )
    type ApiConfig struct {
        IPv4 string
        IPv6 string
    }
    var ApiConfigJson ApiConfig
    _ = json.Unmarshal( ApiConfigObj , &ApiConfigJson )
    return ApiConfigJson.IPv4
}
func URLQuery( URLQueryUserName string , URLQueryPassword string , FQDNorIP string , QueryMap map[string]string ) string {
    URLQueryObj := &url.URL {
        Scheme : "https" ,
        User : URLCreds( URLQueryUserName , URLQueryPassword ) ,
        Host : FQDNorIP ,
        Path : "/" ,
        RawQuery : URLValues( QueryMap ) ,
    }
    return URLQueryObj.String()
}

func main() {
    ApiConfigURL := "https://raw.githubusercontent.com/thewginstaller/config/main/conf.json"
    InputUserName := "Sajjad"
    InputPassword := "drama"
    SecondRequestQueryMap := map[ string ]string { "Profile" : "US" , "ID" : SessionID() }
    fmt.Println( URLQuery( InputUserName , InputPassword , ApiIP( ApiConfigURL ) , SecondRequestQueryMap ) )
}
