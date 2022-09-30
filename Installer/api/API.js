// const http = require( "http" );
const FS = require( "./FS.js" );
const Https = require( "https" );

const ApiTlsCACrt = FS.ReadFile( "TLS" , "API-TLS-CA-ED25519.crt" );
const ApiTlsCrt = FS.ReadFile( "TLS" , "API-TLS-ED25519.crt" );
const ApiTlsKey = FS.ReadFile( "TLS" , "API-TLS-ED25519.key" );

let ConfFilesNameList = FS.ListDir( "CONFS" );
let ConfFilesNumber = ConfFilesNameList.length ;
let ConfFilesObjList = [];
for( let i = 0 ; i < ConfFilesNumber ; i++ ) {
    ConfFilesObjList[ i ] = FS.ReadFile( "CONFS" , ConfFilesNameList[i] );
}

const RequestListener = function( Req , Res ) {
    let Response ;
    if( Req.method == "GET" ) {
        let ReqURL = Req.url.slice( 1 ) ;
        Response = ConfFilesObjList[ ConfFilesNameList.indexOf( ReqURL ) ];
    }
    Res.end( Response );
};

const TlsOptions = {
    ca : ApiTlsCACrt ,
    cert : ApiTlsCrt ,
    key : ApiTlsKey ,
    minVersion : "TLSv1.3"
};


// requestCert : True ?? for client auth

// const server = http.createServer( RequestListener );
const Server = Https.createServer( TlsOptions , RequestListener );
// server.listen( 80 );
Server.listen( 443 , "192.168.1.106" );

// https://nodejs.org/api/tls.html#tlscreatesecurecontextoptions
