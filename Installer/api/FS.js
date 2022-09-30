const FS = require( "fs" );


exports.ReadFile = function( DirName , FileName ) {
    let FileObj = FS.readFileSync( DirName + "/" + FileName );
    return FileObj
}

exports.ListDir = function( DirName ) {
    let DirList = FS.readdirSync( DirName + "/" );
    return DirList
}

