import subprocess
import sys
import os
import math

def Run( Command ) :
    CommandArgs = Command.split(" ")
    Result = subprocess.run( CommandArgs , capture_output=True , text=True ).stdout.strip()
    return Result

def Exec( Command ) :
    CommandArgs = Command.split(" ")
    Result = subprocess.run( CommandArgs )
    return Result

def InputLoop( Text , Default , Options ) :
    Options.append( Default )
    Text += " [ " + Default + " ] "
    Input = input( Text )
    if Input == "" :
        Input = Default
    while Input not in Options :
        print( "Input Not Found in Available Options , Try Again !" )
        Input = input( Text )
    return Input

def ErrorMsg( Text ) :
    print()
    sys.exit( Text )

def TerminalRow( DataList , ColoumnWidthList ) :
    TerminalWidth = os.get_terminal_size().columns
    TotalColoumnsWidth = 0
    for ColoumnWidth in ColoumnWidthList :
        TotalColoumnsWidth += ColoumnWidth
    TerminalColoumnWidthList = []
    for ColoumnWidth in ColoumnWidthList :
        TerminalColoumnWidthList.append( math.trunc( ColoumnWidth * TerminalWidth / TotalColoumnsWidth ) )
    i = 0
    for Data in DataList :
        spaces = ( TerminalColoumnWidthList[ i ] - len( Data ) ) * " "
        if i == len( DataList ) - 1 :
            spaces += "\n"
        print( Data , end = spaces )
        i = i + 1

