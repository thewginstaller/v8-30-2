const Shell = require("child_process");

exports.Run = function( Command ) {
    let CommandArgs = Command.split(" ");
    let Result = Shell.spawnSync( CommandArgs[ 0 ] , CommandArgs.slice( 1 ) , { encoding : "utf8" } ).stdout ;
    Result = Result.slice( 0 , Result.length - 1 )
    return Result
}

exports.Exec = function( Command ) {
    Shell.execSync( Command );
}
