RP/0/RP0/CPU0:bdr03-ipt-454stpau-bne.au#show route ipv6 unicast 2a01:7e00::f03c:92ff:fe60:b2b0

Mon Apr 14 04:54:43.097 UTC

Routing entry for 2a01:7e00::/32
  Known via "bgp 38195", distance 200, metric 0
  Tag 63949
  Number of pic paths 1 , type internal
  Installed Apr 11 05:20:12.435 for 2d23h
  Routing Descriptor Blocks
    2401:d000:8100::1, from fdff:dd:ddaa:1::1
      Route metric is 0
    2401:d001:1100::2, from fdff:dd:ddaa:1::1, BGP backup path
      Route metric is 0
  No advertising protos.
 

Routing entry for 2401:d000:8100::1/128
  Known via "isis core", distance 115, metric 26110, type level-2
  Installed Apr 14 02:55:14.937 for 02:06:47
  Routing Descriptor Blocks
    fe80::9e09:8bff:fe03:bd0e, from 2401:d000:8100::1, via Bundle-Ether42.100
      Route metric is 26110
    fe80::9e09:8bff:fe03:bd0d, from 2401:d000:8100::1, via Bundle-Ether41.111
      Route metric is 26110
    fe80::9e09:8bff:fe03:df0f, from 2401:d000:8100::1, via Bundle-Ether42.421
      Route metric is 26110
  No advertising protos.
RP/0/RP0/CPU0:bdr01-ipt-1william-per.au#show bgp ipv6 unicast 2a01:7e01::f03c:92ff:fe60:4204

Mon Apr 14 05:25:17.851 UTC
BGP routing table entry for 2a01:7e01::/32
Versions:
  Process           bRIB/RIB  SendTblVer
  Speaker         22045211914  22045211914
Last Modified: Apr 14 05:03:51.614 for 00:21:26
Paths: (6 available, best #1)
  Advertised IPv6 Unicast paths to update-groups (with more than one peer):
    0.14 0.19 
  Advertised IPv6 Unicast paths to peers (in unique update groups):
    2620:129:1:2::1                         
    2401:d000:b201::12                      
    2401:d000:b201::42                      
  Path #1: Received by speaker 0
  Advertised IPv6 Unicast paths to update-groups (with more than one peer):
    0.14 0.19 
  Advertised IPv6 Unicast paths to peers (in unique update groups):
    2620:129:1:2::1                         
    2401:d000:b201::12                      
    2401:d000:b201::42                      
  63949
    2401:d000:8100::1 (metric 20310) from fdff:dd:ddaa:1::1 (202.130.207.1)
      Origin IGP, localpref 130, valid, internal, best, group-best, import-candidate
      Received Path ID 1, Local Path ID 1, version 22045211907
      Community: 38195:6676 38195:7130 38195:7324 38195:7675 38195:8013 38195:9705 63949:1004 65010:63949
      Originator: 202.130.207.1, Cluster list: 192.168.255.11
Routing entry for 2a01:7e01::/32
  Known via "bgp 38195", distance 200, metric 0
  Tag 63949
  Number of pic paths 1 , type internal
  Installed Apr 14 05:03:51.838 for 00:25:44
  Routing Descriptor Blocks
    2401:d000:8100::1, from fdff:dd:ddaa:1::1
      Route metric is 0
    2401:d001:1100::2, from fdff:dd:ddaa:1::1, BGP backup path
      Route metric is 0
  No advertising protos.
RP/0/RP0/CPU0:bdr01-ipt-15pionee-sin.sg#show route ipv6 unicast 2a01:7e01::f03c:92ff:fe60:4204

Mon Apr 14 05:32:56.493 UTC

Routing entry for 2a01:7e01::/32
  Known via "bgp 38195", distance 200, metric 0
  Tag 3356
  Number of pic paths 1 , type internal
  Installed Apr 14 05:03:51.898 for 00:29:04
  Routing Descriptor Blocks
    2401:d000:8100::1, from fdff:dd:ddaa:1::1, BGP backup path
      Route metric is 0
    2401:d001:1100::2, from 2401:d001:1100::2           <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
      Route metric is 0
  No advertising protos.
RP/0/RP0/CPU0:bdr01-ipt-15pionee-sin.sg#sh bgp ipv6 unicast 2a01:7e01::f03c:92ff:fe60:4204
  3356 2914 20940 63949, (received & used)
    2401:d001:1100::2 (metric 9410) from 2401:d001:1100::2 (202.177.40.3)     <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
      Origin IGP, metric 0, localpref 142, valid, internal, best, group-best, import-candidate
      Received Path ID 0, Local Path ID 1, version 14465530920
      Community: 2914:410 2914:1212 2914:2212 2914:3200 3356:4 3356:22 3356:86 3356:601 3356:666 3356:703 3356:901 3356:2172 38195:6800 38195:7349 38195:7663 38195:8010 38195:9708
 
 (202.177.40.3) is one of the eip routers in sg
 9410 is better cost than 10000
RP/0/RP0/CPU0:bdr01-ipt-15pionee-sin.sg#sh run formal | i 34.771 

Mon Apr 14 05:52:59.551 UTC
Building configuration...
interface Bundle-Ether34.771 description NNI-IPT-L3: bdr01-eip-15pionee-sin.sg Hu0/0/0/23.771 ()[100Gbit]{1316643}
interface Bundle-Ether34.771 mtu 9212
interface Bundle-Ether34.771 ipv6 address 2401:d000:10:800::52/124
<snip>
router isis core interface Bundle-Ether34.771 address-family ipv6 unicast metric 10000      <<<<<<<<<<<<<<<<<
15:05
thats what I sent to Laureen
