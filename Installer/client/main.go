package main

// import "fmt"
import "net/http"
import "io"
import "os"

func main() {
    Res , _ := http.Get( "http://192.168.1.106/UK1-U1.conf" )
    defer Res.Body.Close()
    ConfFileObj , _ := os.Create( "UK1-U1.conf" )
    _ , _ = io.Copy( ConfFileObj , Res.Body )
    defer ConfFileObj.Close()
}

// Body , _ := io.ReadAll( Res.Body )
// defer Res.Body.Close()
// fmt.Printf( "%s" , Body )

// https://raw.githubusercontent.com/thewginstaller/v8-30/main/Installer/client/CONFS/UK1-U1.conf
