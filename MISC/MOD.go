Existing config:

bdr01-met-15pionee-sin.sg
    set interfaces ae1 unit 101 description "Cust: MOD Mission Critical - IP-PT (None)[10000Mbit]{607088}"
    set interfaces ae1 unit 101 encapsulation vlan-ccc
    set interfaces ae1 unit 101 vlan-tags outer 0x8100.101
    set interfaces ae1 unit 101 input-vlan-map pop
    set interfaces ae1 unit 101 output-vlan-map push
    set interfaces ae1 unit 101 statistics
    set interfaces ae1 unit 101 family ccc filter input rate-limit-10000mbps


    set protocols l2circuit neighbor 172.21.55.9 interface ae1.101 virtual-circuit-id 607088
    set protocols l2circuit neighbor 172.21.55.9 interface ae1.101 description "Cust: MOD Mission Critical - IP-PT (None)[10000Mbit]{607088}"
    set protocols l2circuit neighbor 172.21.55.9 interface ae1.101 community evc-core-protected
    set protocols l2circuit neighbor 172.21.55.9 interface ae1.101 encapsulation-type ethernet
    set protocols l2circuit neighbor 172.21.55.9 interface ae1.101 ignore-mtu-mismatch
    set protocols l2circuit neighbor 172.21.55.9 interface ae1.101 pseudowire-status-tlv


bdr01-etp-15pionee-sin.sg

    interface TenGigE0/0/0/7 description Cust: MOD Mission Critical - IP-PT [10000Mbit]{607088}
    interface TenGigE0/0/0/7 mtu 9216
    interface TenGigE0/0/0/7 load-interval 30
    interface TenGigE0/0/0/7 l2transport


    l2vpn xconnect group xc_607088
    l2vpn xconnect group xc_607088 p2p p2p_607088
    l2vpn xconnect group xc_607088 p2p p2p_607088 interface TenGigE0/0/0/7
    l2vpn xconnect group xc_607088 p2p p2p_607088 neighbor ipv4 100.127.224.129 pw-id 607088
    l2vpn xconnect group xc_607088 p2p p2p_607088 neighbor ipv4 100.127.224.129 pw-id 607088 pw-class bdr01-met-15pionee-sin-cp-01

Existing setup:

bdr01-etp-15pionee-sin.sg ---------l2vc-------bdr01-met-15pionee-sin.sg

Proposed setup:


bdr02-etp-15pionee-sin.sg -----l2vc-----bdr01-eip-15pionee:



bdr02-etp-15pionee-sin.sg:


    interface TenGigE0/0/0/0 description Prov: Cust: MOD Mission Critical - IP-PT [10000Mbit]{607088}
    interface TenGigE0/0/0/0 mtu 9216
    interface TenGigE0/0/0/0 load-interval 30
    interface TenGigE0/0/0/0 l2transport


    l2vpn xconnect group xc_607088
    l2vpn xconnect group xc_607088 p2p p2p_607088
    l2vpn xconnect group xc_607088 p2p p2p_607088 interface TenGigE0/0/0/7
    l2vpn xconnect group xc_607088 p2p p2p_607088 neighbor ipv4 202.177.40.4 pw-id 607088
    l2vpn xconnect group xc_607088 p2p p2p_607088 neighbor ipv4 202.177.40.4 pw-id 607088 pw-class CW-enabled


bdr01-eip-15pionee-sin.sg:



l2vpn xconnect group xc_607088
l2vpn xconnect group xc_607088 p2p p2p_607088
l2vpn xconnect group xc_607088 p2p p2p_607088 interface HundredGigE0/0/0/22.607088
l2vpn xconnect group xc_607088 p2p p2p_607088 neighbor ipv4 172.21.55.13 pw-id 607088
l2vpn xconnect group xc_607088 p2p p2p_607088 neighbor ipv4 172.21.55.13 pw-id 607088 pw-class CW-enabled

interface HundredGigE0/0/0/22.607088 l2transport
interface HundredGigE0/0/0/22.607088 l2transport description Loop-L2: MOD Mission Critical - IP-PT (AS39855)[5000Mbps]{607088}
interface HundredGigE0/0/0/22.607088 l2transport encapsulation dot1q 103
interface HundredGigE0/0/0/22.607088 l2transport rewrite ingress tag pop 1 symmetric
interface HundredGigE0/0/0/22.607088 l2transport mtu 9216


interface HundredGigE0/0/0/23.103
interface HundredGigE0/0/0/23.103 description Prov: Cust: MOD Mission Critical - IP-PT (AS39855)[5000Mbps]{607088}
interface HundredGigE0/0/0/23.103 service-policy input POLICE-5000MBPS
interface HundredGigE0/0/0/23.103 ipv4 mtu 1500
interface HundredGigE0/0/0/23.103 ipv4 address 202.177.42.57 255.255.255.248 secondary
interface HundredGigE0/0/0/23.103 ipv4 unreachables disable
interface HundredGigE0/0/0/23.103 load-interval 30
interface HundredGigE0/0/0/23.103 flow ipv4 monitor netflow-ipv4 sampler 1in1k ingress
interface HundredGigE0/0/0/23.103 encapsulation dot1q 103


router static address-family ipv4 unicast 85.203.60.0/22 202.177.42.59 description SID:607088
router static address-family ipv4 unicast 91.142.140.0/24 202.177.42.59 description SID:607088
router static address-family ipv4 unicast 93.123.114.0/24 202.177.42.59 description SID:607088
router static address-family ipv4 unicast 185.250.42.0/23 202.177.42.59 description SID:607088

router bgp 64590 neighbor 202.177.42.58
no shutdown

router bgp 64590 neighbor 202.177.42.58 shutdown
router bgp 64590 neighbor 202.177.42.58 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 202.177.42.58 remote-as 39855
router bgp 64590 neighbor 202.177.42.58 timers 10 30
router bgp 64590 neighbor 202.177.42.58 description Cust:MOD Mission Critical-607088
router bgp 64590 neighbor 202.177.42.58 address-family ipv4 unicast
router bgp 64590 neighbor 202.177.42.58 address-family ipv4 unicast maximum-prefix 100 90 restart 15
router bgp 64590 neighbor 202.177.42.58 address-family ipv4 unicast send-community-ebgp
router bgp 64590 neighbor 202.177.42.58 address-family ipv4 unicast route-policy SLCAU-CUST-IPV4-IN(AS39855-607088-IPV4-IN) in
router bgp 64590 neighbor 202.177.42.58 address-family ipv4 unicast route-policy SLCAU-CUST-IPV4-DEFAULTONLY-OUT out
router bgp 64590 neighbor 202.177.42.58 address-family ipv4 unicast default-originate
router bgp 64590 neighbor 202.177.42.58 address-family ipv4 unicast remove-private-AS
router bgp 64590 neighbor 202.177.42.58 address-family ipv4 unicast soft-reconfiguration inbound always




prefix-set AS39855-607088-IPV4-IN
   Company: MOD Mission Critical
   Service ID: 607088
   IRT: https://apps-networks.au.superloop.systems/irt/company/616/
  136.144.16.0/24 le 32,
  136.144.18.0/24 le 32,
  136.144.32.0/24 le 32,
  136.144.34.0/24 le 32,
  185.214.11.0/24 le 32,
  185.216.0.03   Company: MOD Mission Critical    Service ID: 607088    IRT: https://apps-networks.au.superloop.systems/irt/company/616/2  136.144.16.0/24 le 32,5.0/24 le 32,
  194.99.112.0/24 le 32,
  203.159.84.0/24 le 32,
  203.159.86.0/24 le 32,
  203.159.88.0/24 le 32,
  45.136.72.0/22 le 32,
  45.144.197.0/24 le 32,
  45.144.216.0/22 le 32,
  45.146.56.0/22 le 32,
  45.86.84.0/24 le 32,
  45.91.117.0/24 le 32,
  5.10.194.0/24 le 32,
  5.10.196.0/24 le 32,
  5.10.198.0/24 le 32,
  89.36.172.0/24 le 32
end-set
!

policy-map POLICE-5000MBPS
 class class-default
  police rate 5000 mbps
  !
  set dscp default
 !
 end-policy-map
!
policy-map SHAPE-5000MBPS
 class class-default
  shape average 5000 mbps
 !
 end-policy-map
!




bdr01-ipt-15pionee-sin.sg:

        no router static address-family ipv4 unicast 91.142.140.0/24 202.177.42.59 description SID:607088
        no router static address-family ipv4 unicast 93.123.114.0/24 202.177.42.59 description SID:607088
        no router static address-family ipv4 unicast 185.250.42.0/23 202.177.42.59 description SID:607088


        no interface Bundle-Ether2.101 ipv4 address 202.177.42.57 255.255.255.248 secondary

        router bgp 38195 neighbor 202.177.42.58 address-family ipv4 unicast route-policy  DENY in
        router bgp 38195 neighbor 202.177.42.58 address-family ipv4 unicast route-policy  DENY out

        interface Bundle-Ether2.101 description Decom: Cust: MOD Mission Critical - IP-PT (AS39855)[5000Mbps]{607088}












Precheck:

RP/0/RP0/CPU0:bdr01-ipt-15pionee-sin.sgsh bgp neighbor 202.177.42.58 routes
Thu Mar 13 07:18:14.115 UTC
BGP router identifier 202.177.40.1, local AS number 38195
BGP generic scan interval 60 secs
Non-stop routing is enabled
BGP table state: Active
Table ID: 0xe0000000   RD version: 17419107430
BGP main routing table version 17419107430
BGP NSR Initial initsync version 850831 (Reached)
BGP NSR/ISSU Sync-Group versions 0/0
BGP scan interval 60 secs

Status codes: s suppressed, d damped, h history, * valid, > best
              i - internal, r RIB-failure, S stale, N Nexthop-discard
Origin codes: i - IGP, e - EGP, ? - incomplete
   Network            Next Hop            Metric LocPrf Weight Path
*> 5.10.194.0/24      202.177.42.58                 194      0 39855 39855 39855 i
*> 5.10.196.0/24      202.177.42.58                 194      0 39855 39855 39855 i
*> 5.10.198.0/24      202.177.42.58                 194      0 39855 39855 39855 i
*> 89.36.172.0/24     202.177.42.58                 194      0 39855 39855 39855 i
*> 136.144.16.0/24    202.177.42.58                 194      0 39855 39855 39855 i
*> 136.144.18.0/24    202.177.42.58                 194      0 39855 39855 39855 i
*> 136.144.32.0/24    202.177.42.58                 194      0 39855 39855 39855 i
*> 136.144.34.0/24    202.177.42.58                 194      0 39855 39855 39855 i
*> 185.67.136.0/24    202.177.42.58                 194      0 39855 39855 39855 i
*> 185.67.138.0/24    202.177.42.58                 194      0 39855 39855 39855 i
*> 185.214.11.0/24    202.177.42.58                 194      0 39855 39855 39855 i
*> 185.216.0.0/24     202.177.42.58                 194      0 39855 39855 39855 i
*> 185.216.2.0/24     202.177.42.58                 194      0 39855 39855 39855 i
*> 203.159.84.0/24    202.177.42.58                 194      0 39855 39855 39855 i
*> 203.159.86.0/24    202.177.42.58                 194      0 39855 39855 39855 i
*> 203.159.88.0/24    202.177.42.58                 194      0 39855 39855 39855 i

Processed 16 prefixes, 16 paths
RP/0/RP0/CPU0:bdr01-ipt-15pionee-sin.sgsh bgp neighbor 202.177.42.58 advertised-routes
Thu Mar 13 07:18:22.150 UTC
Network            Next Hop        From            AS Path
0.0.0.0/0          202.177.42.57   Local           38195i

Processed 1 prefixes, 1 paths


sh bgp neighbor 202.177.42.58 routes
Thu Mar 13 12:06:55.010 UTC
BGP router identifier 202.177.40.4, local AS number 64590
BGP generic scan interval 60 secs
Non-stop routing is enabled
BGP table state: Active
Table ID: 0xe0000000   RD version: 416156789
BGP table nexthop route policy:
BGP main routing table version 416156789
BGP NSR Initial initsync version 168 (Reached)
BGP NSR/ISSU Sync-Group versions 0/0
BGP scan interval 60 secs

Status codes: s suppressed, d damped, h history, * valid, > best
              i - internal, r RIB-failure, S stale, N Nexthop-discard
Origin codes: i - IGP, e - EGP, ? - incomplete
   Network            Next Hop            Metric LocPrf Weight Path
*> 5.10.194.0/24      202.177.42.58                 194      0 39855 39855 39855 i
*> 5.10.196.0/24      202.177.42.58                 194      0 39855 39855 39855 i
*> 5.10.198.0/24      202.177.42.58                 194      0 39855 39855 39855 i
*> 89.36.172.0/24     202.177.42.58                 194      0 39855 39855 39855 i
*> 136.144.16.0/24    202.177.42.58                 194      0 39855 39855 39855 i
*> 136.144.18.0/24    202.177.42.58                 194      0 39855 39855 39855 i
*> 136.144.32.0/24    202.177.42.58                 194      0 39855 39855 39855 i
*> 136.144.34.0/24    202.177.42.58                 194      0 39855 39855 39855 i
*> 185.67.136.0/24    202.177.42.58                 194      0 39855 39855 39855 i
*> 185.67.138.0/24    202.177.42.58                 194      0 39855 39855 39855 i
*> 185.214.11.0/24    202.177.42.58                 194      0 39855 39855 39855 i
*> 185.216.0.0/24     202.177.42.58                 194      0 39855 39855 39855 i
*> 185.216.2.0/24     202.177.42.58                 194      0 39855 39855 39855 i
*> 203.159.84.0/24    202.177.42.58                 194      0 39855 39855 39855 i
*> 203.159.86.0/24    202.177.42.58                 194      0 39855 39855 39855 i
*> 203.159.88.0/24    202.177.42.58                 194      0 39855 39855 39855 i
