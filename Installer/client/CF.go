package main

import "fmt"
import "net/http"
import "io"
// import "encoding/json"

func main() {
    FQDN := "google.com"
    Type := "A"
    DoHQueryURL := "https://cloudflare-dns.com/dns-query?name=" + FQDN + "&type=" + Type + "&ct=application/dns-json"
    DoHResponse , _ := http.Get( DoHQueryURL )
    defer DoHResponse.Body.Close()
    DoHResponseBody , _ := io.ReadAll( DoHResponse.Body )
    fmt.Printf( "%s" , DoHResponseBody )
    /*
    jsonMap := make( map[string]( interface{} ) )
    _ = json.Unmarshal( DoHResponseBody , &jsonMap )
    AnswerMapList := jsonMap[ "Answer" ].( []interface{} )
    AnswerMap := AnswerMapList[0].( map[string]interface{} )
    IP := AnswerMap[ "data" ]
    fmt.Println( IP )
    */
}

