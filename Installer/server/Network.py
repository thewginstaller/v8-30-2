import ipaddress
import json
import Bash

def ListNICs() :
    NoneLoopBackNICs = []
    NICs = json.loads( Bash.Run( "ip -j address" ) )
    NICObj = {
        "Name" : "" ,
        "Type" : "" ,
        "IPv4" : { "Addr" : "" , "Type" : "" } ,
        "IPv6" : { "Addr" : "" , "Type" : "" }
    }
    for NIC in NICs :
        if NIC[ "link_type" ] != "loopback" :
            NoneLoopBackNIC = NICObj
            NoneLoopBackNIC[ "Name" ] = NIC[ "ifname" ]
            NoneLoopBackNIC[ "Type" ] = NIC[ "link_type" ]
            for Addr in NIC[ "addr_info" ] :
                if Addr[ "family" ] == "inet" :
                    NoneLoopBackNIC[ "IPv4" ][ "Addr" ] = Addr[ "local" ]
                if Addr[ "family" ] == "inet6" :
                    NoneLoopBackNIC[ "IPv6" ][ "Addr" ] = Addr[ "local" ]
            NoneLoopBackNICs.append( NoneLoopBackNIC )
    return NoneLoopBackNICs

def IsLinkLocal( IP ) :
    return ipaddress.ip_address( IP ).is_link_local
def IsPrivate( IP ) :
    return ipaddress.ip_address( IP ).is_private
def IsGlobal( IP ) :
    return ipaddress.ip_address( IP ).is_global
def HostsList( IPNet ) :
    Hosts = []
    for Host in ipaddress.ip_network( IPNet ).hosts() :
        Hosts.append( str( Host ) )
    return Hosts

# def DoHResolve( FQDN , Type )

# def PortScanApi()
