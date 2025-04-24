#megaport IX


Prechecks:
RP/0/RP0/CPU0:bdr01-ipt-15pionee-sin.sg#sh bgp all all summ wide | i "103.41.12|2001:de$
Sun Apr  6 23:24:10.026 UTC
103.41.12.1                               0       63832   39558994    6083445 17861952096          0          0      1y34w       7640
103.41.12.2                               0       63832   39649068    6084694 17861952096          0          0      1y35w       5987
103.41.12.6                               0       13335    6872140    6082014 17861952096          0          0      5d18h       5421
103.41.12.9                               0        6939   38258677    5889418 17861952096          0          0      20w2d          0
103.41.12.11                              0       15133    6207169    6089081 17861952096          0          0      1y41w          1
103.41.12.13                              0       15133    6194625    6089387 17861952096          0          0      1y41w          1
103.41.12.14                              0         714    6442770    6538124 17861952096          0          0      2d23h       1158
103.41.12.16                              0       22822    6107226    5619122           0          0          0       7w4d Idle (Admin)
103.41.12.18                              0       20940    6399268    6091739 17861952096          0          0      2d23h         11
103.41.12.19                              0         714    6550391    6334040 17861952096          0          0      5d16h       1156
103.41.12.23                              0        8075    4983313    4677635 17861952096          0          0      2d23h          0
103.41.12.30                              0       10310    6242272    6086080 17861952096          0          0      1y17w        131
103.41.12.31                              0       46489    5887068    6089448 17861952096          0          0      2d23h          5
103.41.12.32                              0       46489    5886788    6088897 17861952096          0          0      2d23h          5
103.41.12.35                              0      199524    5960705   50497769 17861952096          0          0      22w5d         68
103.41.12.41                              0       55256    6916818    6087623 17861952096          0          0       3w1d          3
103.41.12.44                              0       20940    2122681    2046721 17861952096          0          0      2d23h         63
103.41.12.45                              0       49544    5693928    5543966 17861952096          0          0      2d23h         22
2001:ded::1                               0       63832   37265754    6480633 14388802566          0          0      1y34w       1566
2001:ded::2                               0       63832   37644422    6482269 14388802566          0          0      1y35w       1043
2001:ded::6                               0       13335    6524378    6490603 14388802566          0          0      5d18h        586
2001:ded::9                               0        6939   43120714    5889426 14388803300          0          0      20w2d          0
2001:ded::b                               0       15133    6207115    6486942 14388802566          0          0      1y41w          1
2001:ded::d                               0       15133    6194577    6485627 14388802566          0          0      38w2d          1
2001:ded::e                               0         714    6499481    6586769 14388802566          0          0      2d23h        349
2001:ded::10                              0       22822    6047275    5964974           0          0          0       7w4d Idle (Admin)
2001:ded::12                              0       20940    6396580    6495333 14388802566          0          0      2d23h          4
2001:ded::13                              0         714    6607772    6530434 14388802566          0          0      5d16h        349
2001:ded::17                              0        8075    4982895    5016218 14388802566          0          0      2d23h          0
2001:ded::1e                              0       10310    6468468    6509068 14388802566          0          0      1y17w        100
2001:ded::1f                              0       46489    5887053    6481382 14388802566          0          0      2d23h          4
2001:ded::20                              0       46489    5886793    6481049 14388802566          0          0      2d23h          4
2001:ded::23                              0      199524    5936258   13573078 14388802566          0          0      51w3d         10
2001:ded::29                              0       55256    6916709    6482651 14388802566          0          0       3w1d          0
2001:ded::2c                              0       20940    2122580    2152234 14388802566          0          0      2d23h         17
2001:ded::2d                              0       49544    5692656    5890624 14388802566          0          0      2d23h          7



router bgp 38195 neighbor 103.41.12.1 address-family ipv4 unicast route-policy DENY in
router bgp 38195 neighbor 103.41.12.2 address-family ipv4 unicast route-policy DENY in
router bgp 38195 neighbor 103.41.12.6 address-family ipv4 unicast route-policy DENY in
router bgp 38195 neighbor 103.41.12.9 address-family ipv4 unicast route-policy DENY in
router bgp 38195 neighbor 2001:ded::1 address-family ipv6 unicast route-policy DENY in
router bgp 38195 neighbor 2001:ded::2 address-family ipv6 unicast route-policy DENY in
router bgp 38195 neighbor 2001:ded::6 address-family ipv6 unicast route-policy DENY in
router bgp 38195 neighbor 2001:ded::9 address-family ipv6 unicast route-policy DENY in
router bgp 38195 neighbor 2001:ded::b address-family ipv6 unicast route-policy DENY in
router bgp 38195 neighbor 2001:ded::d address-family ipv6 unicast route-policy DENY in
router bgp 38195 neighbor 2001:ded::e address-family ipv6 unicast route-policy DENY in
router bgp 38195 neighbor 103.41.12.11 address-family ipv4 unicast route-policy DENY in
router bgp 38195 neighbor 103.41.12.13 address-family ipv4 unicast route-policy DENY in
router bgp 38195 neighbor 103.41.12.14 address-family ipv4 unicast route-policy DENY in
router bgp 38195 neighbor 103.41.12.16 address-family ipv4 unicast route-policy DENY in
router bgp 38195 neighbor 103.41.12.18 address-family ipv4 unicast route-policy DENY in
router bgp 38195 neighbor 103.41.12.19 address-family ipv4 unicast route-policy DENY in
router bgp 38195 neighbor 103.41.12.23 address-family ipv4 unicast route-policy DENY in
router bgp 38195 neighbor 103.41.12.30 address-family ipv4 unicast route-policy DENY in
router bgp 38195 neighbor 103.41.12.31 address-family ipv4 unicast route-policy DENY in
router bgp 38195 neighbor 103.41.12.32 address-family ipv4 unicast route-policy DENY in
router bgp 38195 neighbor 103.41.12.35 address-family ipv4 unicast route-policy DENY in
router bgp 38195 neighbor 103.41.12.41 address-family ipv4 unicast route-policy DENY in
router bgp 38195 neighbor 103.41.12.44 address-family ipv4 unicast route-policy DENY in
router bgp 38195 neighbor 103.41.12.45 address-family ipv4 unicast route-policy DENY in
router bgp 38195 neighbor 2001:ded::10 address-family ipv6 unicast route-policy DENY in
router bgp 38195 neighbor 2001:ded::12 address-family ipv6 unicast route-policy DENY in
router bgp 38195 neighbor 2001:ded::13 address-family ipv6 unicast route-policy DENY in
router bgp 38195 neighbor 2001:ded::17 address-family ipv6 unicast route-policy DENY in
router bgp 38195 neighbor 2001:ded::1e address-family ipv6 unicast route-policy DENY in
router bgp 38195 neighbor 2001:ded::1f address-family ipv6 unicast route-policy DENY in
router bgp 38195 neighbor 2001:ded::20 address-family ipv6 unicast route-policy DENY in
router bgp 38195 neighbor 2001:ded::23 address-family ipv6 unicast route-policy DENY in
router bgp 38195 neighbor 2001:ded::29 address-family ipv6 unicast route-policy DENY in
router bgp 38195 neighbor 2001:ded::2c address-family ipv6 unicast route-policy DENY in
router bgp 38195 neighbor 2001:ded::2d address-family ipv6 unicast route-policy DENY in
router bgp 38195 neighbor 103.41.12.1 address-family ipv4 unicast route-policy DENY out
router bgp 38195 neighbor 103.41.12.2 address-family ipv4 unicast route-policy DENY out
router bgp 38195 neighbor 103.41.12.6 address-family ipv4 unicast route-policy DENY out
router bgp 38195 neighbor 103.41.12.9 address-family ipv4 unicast route-policy DENY out
router bgp 38195 neighbor 2001:ded::1 address-family ipv6 unicast route-policy DENY out
router bgp 38195 neighbor 2001:ded::2 address-family ipv6 unicast route-policy DENY out
router bgp 38195 neighbor 2001:ded::6 address-family ipv6 unicast route-policy DENY out
router bgp 38195 neighbor 2001:ded::9 address-family ipv6 unicast route-policy DENY out
router bgp 38195 neighbor 2001:ded::b address-family ipv6 unicast route-policy DENY out
router bgp 38195 neighbor 2001:ded::d address-family ipv6 unicast route-policy DENY out
router bgp 38195 neighbor 2001:ded::e address-family ipv6 unicast route-policy DENY out
router bgp 38195 neighbor 103.41.12.11 address-family ipv4 unicast route-policy DENY out
router bgp 38195 neighbor 103.41.12.13 address-family ipv4 unicast route-policy DENY out
router bgp 38195 neighbor 103.41.12.14 address-family ipv4 unicast route-policy DENY out
router bgp 38195 neighbor 103.41.12.16 address-family ipv4 unicast route-policy DENY out
router bgp 38195 neighbor 103.41.12.18 address-family ipv4 unicast route-policy DENY out
router bgp 38195 neighbor 103.41.12.19 address-family ipv4 unicast route-policy DENY out
router bgp 38195 neighbor 103.41.12.23 address-family ipv4 unicast route-policy DENY out
router bgp 38195 neighbor 103.41.12.30 address-family ipv4 unicast route-policy DENY out
router bgp 38195 neighbor 103.41.12.31 address-family ipv4 unicast route-policy DENY out
router bgp 38195 neighbor 103.41.12.32 address-family ipv4 unicast route-policy DENY out
router bgp 38195 neighbor 103.41.12.35 address-family ipv4 unicast route-policy DENY out
router bgp 38195 neighbor 103.41.12.41 address-family ipv4 unicast route-policy DENY out
router bgp 38195 neighbor 103.41.12.44 address-family ipv4 unicast route-policy DENY out
router bgp 38195 neighbor 103.41.12.45 address-family ipv4 unicast route-policy DENY out
router bgp 38195 neighbor 2001:ded::10 address-family ipv6 unicast route-policy DENY out
router bgp 38195 neighbor 2001:ded::12 address-family ipv6 unicast route-policy DENY out
router bgp 38195 neighbor 2001:ded::13 address-family ipv6 unicast route-policy DENY out
router bgp 38195 neighbor 2001:ded::17 address-family ipv6 unicast route-policy DENY out
router bgp 38195 neighbor 2001:ded::1e address-family ipv6 unicast route-policy DENY out
router bgp 38195 neighbor 2001:ded::1f address-family ipv6 unicast route-policy DENY out
router bgp 38195 neighbor 2001:ded::20 address-family ipv6 unicast route-policy DENY out
router bgp 38195 neighbor 2001:ded::23 address-family ipv6 unicast route-policy DENY out
router bgp 38195 neighbor 2001:ded::29 address-family ipv6 unicast route-policy DENY out
router bgp 38195 neighbor 2001:ded::2c address-family ipv6 unicast route-policy DENY out
router bgp 38195 neighbor 2001:ded::2d address-family ipv6 unicast route-policy DENY out




#bdr02-etp-15pionee-sin.sg:

interface TenGigE0/0/0/6.732 l2transport
interface TenGigE0/0/0/6.732 l2transport description Cust: APEXn Pty Ltd - SP-MP-IX-E10G (#ba561113)[5000Mbit]{12202}
interface TenGigE0/0/0/6.732 l2transport encapsulation dot1q 732
interface TenGigE0/0/0/6.732 l2transport mtu 9216


l2vpn xconnect group xc_12202
l2vpn xconnect group xc_12202 p2p p2p_12202
l2vpn xconnect group xc_12202 p2p p2p_12202 interface TenGigE0/0/0/6.732
l2vpn xconnect group xc_12202 p2p p2p_12202 neighbor ipv4 202.177.40.3 pw-id 12202
l2vpn xconnect group xc_12202 p2p p2p_12202 neighbor ipv4 202.177.40.3 pw-id 12202 pw-class CW-enabled




#migrate from SG2 to SG3
#bdr01-eip-26aayer


interface HundredGigE0/0/0/22.12202 l2transport
interface HundredGigE0/0/0/22.12202 l2transport description Loop-L2: MegaIX SG AS63832 (ba561113)[5000Mbit]{12202}
interface HundredGigE0/0/0/22.12202 l2transport encapsulation dot1q 732
interface HundredGigE0/0/0/22.12202 l2transport mtu 9216


interface HundredGigE0/0/0/23.732
interface HundredGigE0/0/0/23.732 description Prov: Peering: MegaIX SG AS63832 (ba561113)[5000Mbit]{12202}
interface HundredGigE0/0/0/23.732 ipv4 mtu 1500
interface HundredGigE0/0/0/23.732 ipv4 address 103.41.12.26 255.255.252.0
interface HundredGigE0/0/0/23.732 arp learning local
interface HundredGigE0/0/0/23.732 ipv4 unreachables disable
interface HundredGigE0/0/0/23.732 ipv6 nd suppress-ra
interface HundredGigE0/0/0/23.732 ipv6 mtu 1500
interface HundredGigE0/0/0/23.732 ipv6 address 2001:ded::1a/64
interface HundredGigE0/0/0/23.732 ipv6 unreachables disable
interface HundredGigE0/0/0/23.732 load-interval 30
interface HundredGigE0/0/0/23.732 flow ipv4 monitor netflow-ipv4 sampler 1in1k ingress
interface HundredGigE0/0/0/23.732 flow ipv6 monitor netflow-ipv6 sampler 1in1k ingress
interface HundredGigE0/0/0/23.732  ipv4 access-group AS38195-IPV4-TRANSIT-DENY ingress
interface HundredGigE0/0/0/23.732  ipv6 access-group AS38195-IPV6-TRANSIT-DENY ingress
interface HundredGigE0/0/0/23.732 encapsulation dot1q 723

l2vpn xconnect group xc_12202
l2vpn xconnect group xc_12202 p2p p2p_12202
l2vpn xconnect group xc_12202 p2p p2p_12202 interface HundredGigE0/0/0/22.12202
l2vpn xconnect group xc_12202 p2p p2p_12202 neighbor ipv4 172.21.55.13 pw-id 12202
l2vpn xconnect group xc_12202 p2p p2p_12202 neighbor ipv4 172.21.55.13 pw-id 12202 pw-class CW-enabled




router bgp 64590 neighbor 103.41.12.1 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 103.41.12.1 remote-as 63832
router bgp 64590 neighbor 103.41.12.1 timers 10 30
router bgp 64590 neighbor 103.41.12.1 description Peering:AS63832-SID12202
router bgp 64590 neighbor 103.41.12.1 address-family ipv4 unicast
router bgp 64590 neighbor 103.41.12.1 address-family ipv4 unicast maximum-prefix 87500 90 restart 60
router bgp 64590 neighbor 103.41.12.1 address-family ipv4 unicast send-community-ebgp
router bgp 64590 neighbor 103.41.12.1 address-family ipv4 unicast route-policy AS63832-12202-IPV4-IN in
router bgp 64590 neighbor 103.41.12.1 address-family ipv4 unicast route-policy AS63832-12202-IPV4-OUT out
router bgp 64590 neighbor 103.41.12.1 address-family ipv4 unicast next-hop-self
router bgp 64590 neighbor 103.41.12.1 address-family ipv4 unicast remove-private-AS
router bgp 64590 neighbor 103.41.12.1 address-family ipv4 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 103.41.12.2 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 103.41.12.2 remote-as 63832
router bgp 64590 neighbor 103.41.12.2 timers 10 30
router bgp 64590 neighbor 103.41.12.2 description Peering:AS63832-SID12202
router bgp 64590 neighbor 103.41.12.2 address-family ipv4 unicast
router bgp 64590 neighbor 103.41.12.2 address-family ipv4 unicast maximum-prefix 87500 90 restart 60
router bgp 64590 neighbor 103.41.12.2 address-family ipv4 unicast send-community-ebgp
router bgp 64590 neighbor 103.41.12.2 address-family ipv4 unicast route-policy AS63832-12202-IPV4-IN in
router bgp 64590 neighbor 103.41.12.2 address-family ipv4 unicast route-policy AS63832-12202-IPV4-OUT out
router bgp 64590 neighbor 103.41.12.2 address-family ipv4 unicast next-hop-self
router bgp 64590 neighbor 103.41.12.2 address-family ipv4 unicast remove-private-AS
router bgp 64590 neighbor 103.41.12.2 address-family ipv4 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 103.41.12.6 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 103.41.12.6 remote-as 13335
router bgp 64590 neighbor 103.41.12.6 timers 10 30
router bgp 64590 neighbor 103.41.12.6 description Peering:AS13335-SID12202
router bgp 64590 neighbor 103.41.12.6 address-family ipv4 unicast
router bgp 64590 neighbor 103.41.12.6 address-family ipv4 unicast maximum-prefix 25000 90 restart 60
router bgp 64590 neighbor 103.41.12.6 address-family ipv4 unicast send-community-ebgp
router bgp 64590 neighbor 103.41.12.6 address-family ipv4 unicast route-policy AS13335-12202-IPV4-IN in
router bgp 64590 neighbor 103.41.12.6 address-family ipv4 unicast route-policy AS13335-12202-IPV4-OUT out
router bgp 64590 neighbor 103.41.12.6 address-family ipv4 unicast next-hop-self
router bgp 64590 neighbor 103.41.12.6 address-family ipv4 unicast remove-private-AS
router bgp 64590 neighbor 103.41.12.6 address-family ipv4 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 103.41.12.9 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 103.41.12.9 remote-as 6939
router bgp 64590 neighbor 103.41.12.9 timers 10 30
router bgp 64590 neighbor 103.41.12.9 description Peering:AS6939-SID12202
router bgp 64590 neighbor 103.41.12.9 address-family ipv4 unicast
router bgp 64590 neighbor 103.41.12.9 address-family ipv4 unicast maximum-prefix 211250 90 restart 60
router bgp 64590 neighbor 103.41.12.9 address-family ipv4 unicast send-community-ebgp
router bgp 64590 neighbor 103.41.12.9 address-family ipv4 unicast route-policy DENY in
router bgp 64590 neighbor 103.41.12.9 address-family ipv4 unicast route-policy DENY out
router bgp 64590 neighbor 103.41.12.9 address-family ipv4 unicast next-hop-self
router bgp 64590 neighbor 103.41.12.9 address-family ipv4 unicast remove-private-AS
router bgp 64590 neighbor 103.41.12.9 address-family ipv4 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 2001:ded::1 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 2001:ded::1 remote-as 63832
router bgp 64590 neighbor 2001:ded::1 timers 10 30
router bgp 64590 neighbor 2001:ded::1 description Peering:AS63832-SID12202
router bgp 64590 neighbor 2001:ded::1 address-family ipv6 unicast
router bgp 64590 neighbor 2001:ded::1 address-family ipv6 unicast maximum-prefix 50000 90 restart 60
router bgp 64590 neighbor 2001:ded::1 address-family ipv6 unicast send-community-ebgp
router bgp 64590 neighbor 2001:ded::1 address-family ipv6 unicast route-policy AS63832-12202-IPV6-IN in
router bgp 64590 neighbor 2001:ded::1 address-family ipv6 unicast route-policy AS63832-12202-IPV6-OUT out
router bgp 64590 neighbor 2001:ded::1 address-family ipv6 unicast next-hop-self
router bgp 64590 neighbor 2001:ded::1 address-family ipv6 unicast remove-private-AS
router bgp 64590 neighbor 2001:ded::1 address-family ipv6 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 2001:ded::2 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 2001:ded::2 remote-as 63832
router bgp 64590 neighbor 2001:ded::2 timers 10 30
router bgp 64590 neighbor 2001:ded::2 description Peering:AS63832-SID12202
router bgp 64590 neighbor 2001:ded::2 address-family ipv6 unicast
router bgp 64590 neighbor 2001:ded::2 address-family ipv6 unicast maximum-prefix 50000 90 restart 60
router bgp 64590 neighbor 2001:ded::2 address-family ipv6 unicast send-community-ebgp
router bgp 64590 neighbor 2001:ded::2 address-family ipv6 unicast route-policy AS63832-12202-IPV6-IN in
router bgp 64590 neighbor 2001:ded::2 address-family ipv6 unicast route-policy AS63832-12202-IPV6-OUT out
router bgp 64590 neighbor 2001:ded::2 address-family ipv6 unicast next-hop-self
router bgp 64590 neighbor 2001:ded::2 address-family ipv6 unicast remove-private-AS
router bgp 64590 neighbor 2001:ded::2 address-family ipv6 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 2001:ded::6 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 2001:ded::6 remote-as 13335
router bgp 64590 neighbor 2001:ded::6 timers 10 30
router bgp 64590 neighbor 2001:ded::6 description Peering:AS13335-SID12202
router bgp 64590 neighbor 2001:ded::6 address-family ipv6 unicast
router bgp 64590 neighbor 2001:ded::6 address-family ipv6 unicast maximum-prefix 2500 90 restart 60
router bgp 64590 neighbor 2001:ded::6 address-family ipv6 unicast send-community-ebgp
router bgp 64590 neighbor 2001:ded::6 address-family ipv6 unicast route-policy AS13335-12202-IPV6-IN in
router bgp 64590 neighbor 2001:ded::6 address-family ipv6 unicast route-policy AS13335-12202-IPV6-OUT out
router bgp 64590 neighbor 2001:ded::6 address-family ipv6 unicast next-hop-self
router bgp 64590 neighbor 2001:ded::6 address-family ipv6 unicast remove-private-AS
router bgp 64590 neighbor 2001:ded::6 address-family ipv6 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 2001:ded::9 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 2001:ded::9 remote-as 6939
router bgp 64590 neighbor 2001:ded::9 timers 10 30
router bgp 64590 neighbor 2001:ded::9 description Peering:AS6939-SID12202
router bgp 64590 neighbor 2001:ded::9 address-family ipv6 unicast
router bgp 64590 neighbor 2001:ded::9 address-family ipv6 unicast maximum-prefix 112500 90 restart 60
router bgp 64590 neighbor 2001:ded::9 address-family ipv6 unicast send-community-ebgp
router bgp 64590 neighbor 2001:ded::9 address-family ipv6 unicast route-policy DENY in
router bgp 64590 neighbor 2001:ded::9 address-family ipv6 unicast route-policy DENY out
router bgp 64590 neighbor 2001:ded::9 address-family ipv6 unicast next-hop-self
router bgp 64590 neighbor 2001:ded::9 address-family ipv6 unicast remove-private-AS
router bgp 64590 neighbor 2001:ded::9 address-family ipv6 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 2001:ded::b local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 2001:ded::b remote-as 15133
router bgp 64590 neighbor 2001:ded::b timers 10 30
router bgp 64590 neighbor 2001:ded::b description Peering:AS15133-SID12202
router bgp 64590 neighbor 2001:ded::b address-family ipv6 unicast
router bgp 64590 neighbor 2001:ded::b address-family ipv6 unicast maximum-prefix 6250 90 restart 60
router bgp 64590 neighbor 2001:ded::b address-family ipv6 unicast send-community-ebgp
router bgp 64590 neighbor 2001:ded::b address-family ipv6 unicast route-policy AS15133-12202-IPV6-IN in
router bgp 64590 neighbor 2001:ded::b address-family ipv6 unicast route-policy AS15133-12202-IPV6-OUT out
router bgp 64590 neighbor 2001:ded::b address-family ipv6 unicast next-hop-self
router bgp 64590 neighbor 2001:ded::b address-family ipv6 unicast remove-private-AS
router bgp 64590 neighbor 2001:ded::b address-family ipv6 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 2001:ded::d local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 2001:ded::d remote-as 15133
router bgp 64590 neighbor 2001:ded::d timers 10 30
router bgp 64590 neighbor 2001:ded::d description Peering:AS15133-SID12202
router bgp 64590 neighbor 2001:ded::d address-family ipv6 unicast
router bgp 64590 neighbor 2001:ded::d address-family ipv6 unicast maximum-prefix 6250 90 restart 60
router bgp 64590 neighbor 2001:ded::d address-family ipv6 unicast send-community-ebgp
router bgp 64590 neighbor 2001:ded::d address-family ipv6 unicast route-policy AS15133-12202-IPV6-IN in
router bgp 64590 neighbor 2001:ded::d address-family ipv6 unicast route-policy AS15133-12202-IPV6-OUT out
router bgp 64590 neighbor 2001:ded::d address-family ipv6 unicast next-hop-self
router bgp 64590 neighbor 2001:ded::d address-family ipv6 unicast remove-private-AS
router bgp 64590 neighbor 2001:ded::d address-family ipv6 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 2001:ded::e local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 2001:ded::e remote-as 714
router bgp 64590 neighbor 2001:ded::e timers 10 30
router bgp 64590 neighbor 2001:ded::e description Peering:AS714-SID12202
router bgp 64590 neighbor 2001:ded::e address-family ipv6 unicast
router bgp 64590 neighbor 2001:ded::e address-family ipv6 unicast maximum-prefix 1250 90 restart 60
router bgp 64590 neighbor 2001:ded::e address-family ipv6 unicast send-community-ebgp
router bgp 64590 neighbor 2001:ded::e address-family ipv6 unicast route-policy AS714-12202-IPV6-IN in
router bgp 64590 neighbor 2001:ded::e address-family ipv6 unicast route-policy AS714-12202-IPV6-OUT out
router bgp 64590 neighbor 2001:ded::e address-family ipv6 unicast next-hop-self
router bgp 64590 neighbor 2001:ded::e address-family ipv6 unicast remove-private-AS
router bgp 64590 neighbor 2001:ded::e address-family ipv6 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 103.41.12.11 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 103.41.12.11 remote-as 15133
router bgp 64590 neighbor 103.41.12.11 timers 10 30
router bgp 64590 neighbor 103.41.12.11 description Peering:AS15133-SID12202
router bgp 64590 neighbor 103.41.12.11 address-family ipv4 unicast
router bgp 64590 neighbor 103.41.12.11 address-family ipv4 unicast maximum-prefix 6250 90 restart 60
router bgp 64590 neighbor 103.41.12.11 address-family ipv4 unicast send-community-ebgp
router bgp 64590 neighbor 103.41.12.11 address-family ipv4 unicast route-policy AS15133-12202-IPV4-IN in
router bgp 64590 neighbor 103.41.12.11 address-family ipv4 unicast route-policy AS15133-12202-IPV4-OUT out
router bgp 64590 neighbor 103.41.12.11 address-family ipv4 unicast next-hop-self
router bgp 64590 neighbor 103.41.12.11 address-family ipv4 unicast remove-private-AS
router bgp 64590 neighbor 103.41.12.11 address-family ipv4 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 103.41.12.13 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 103.41.12.13 remote-as 15133
router bgp 64590 neighbor 103.41.12.13 timers 10 30
router bgp 64590 neighbor 103.41.12.13 description Peering:AS15133-SID12202
router bgp 64590 neighbor 103.41.12.13 address-family ipv4 unicast
router bgp 64590 neighbor 103.41.12.13 address-family ipv4 unicast maximum-prefix 6250 90 restart 60
router bgp 64590 neighbor 103.41.12.13 address-family ipv4 unicast send-community-ebgp
router bgp 64590 neighbor 103.41.12.13 address-family ipv4 unicast route-policy AS15133-12202-IPV4-IN in
router bgp 64590 neighbor 103.41.12.13 address-family ipv4 unicast route-policy AS15133-12202-IPV4-OUT out
router bgp 64590 neighbor 103.41.12.13 address-family ipv4 unicast next-hop-self
router bgp 64590 neighbor 103.41.12.13 address-family ipv4 unicast remove-private-AS
router bgp 64590 neighbor 103.41.12.13 address-family ipv4 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 103.41.12.14 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 103.41.12.14 remote-as 714
router bgp 64590 neighbor 103.41.12.14 timers 10 30
router bgp 64590 neighbor 103.41.12.14 description Peering:AS714-SID12202
router bgp 64590 neighbor 103.41.12.14 address-family ipv4 unicast
router bgp 64590 neighbor 103.41.12.14 address-family ipv4 unicast maximum-prefix 12500 90 restart 60
router bgp 64590 neighbor 103.41.12.14 address-family ipv4 unicast send-community-ebgp
router bgp 64590 neighbor 103.41.12.14 address-family ipv4 unicast route-policy AS714-12202-IPV4-IN in
router bgp 64590 neighbor 103.41.12.14 address-family ipv4 unicast route-policy AS714-12202-IPV4-OUT out
router bgp 64590 neighbor 103.41.12.14 address-family ipv4 unicast next-hop-self
router bgp 64590 neighbor 103.41.12.14 address-family ipv4 unicast remove-private-AS
router bgp 64590 neighbor 103.41.12.14 address-family ipv4 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 103.41.12.16 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 103.41.12.16 remote-as 22822
router bgp 64590 neighbor 103.41.12.16 timers 10 30
router bgp 64590 neighbor 103.41.12.16 shutdown
router bgp 64590 neighbor 103.41.12.16 description Peering:AS22822-SID12202
router bgp 64590 neighbor 103.41.12.16 address-family ipv4 unicast
router bgp 64590 neighbor 103.41.12.16 address-family ipv4 unicast maximum-prefix 6250 90 restart 60
router bgp 64590 neighbor 103.41.12.16 address-family ipv4 unicast send-community-ebgp
router bgp 64590 neighbor 103.41.12.16 address-family ipv4 unicast route-policy AS22822-12202-IPV4-IN in
router bgp 64590 neighbor 103.41.12.16 address-family ipv4 unicast route-policy AS22822-12202-IPV4-OUT out
router bgp 64590 neighbor 103.41.12.16 address-family ipv4 unicast next-hop-self
router bgp 64590 neighbor 103.41.12.16 address-family ipv4 unicast remove-private-AS
router bgp 64590 neighbor 103.41.12.16 address-family ipv4 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 103.41.12.18 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 103.41.12.18 remote-as 20940
router bgp 64590 neighbor 103.41.12.18 timers 10 30
router bgp 64590 neighbor 103.41.12.18 description Peering:AS20940-SID12202
router bgp 64590 neighbor 103.41.12.18 address-family ipv4 unicast
router bgp 64590 neighbor 103.41.12.18 address-family ipv4 unicast maximum-prefix 15000 90 restart 60
router bgp 64590 neighbor 103.41.12.18 address-family ipv4 unicast send-community-ebgp
router bgp 64590 neighbor 103.41.12.18 address-family ipv4 unicast route-policy AS20940-12202-IPV4-IN in
router bgp 64590 neighbor 103.41.12.18 address-family ipv4 unicast route-policy AS20940-12202-IPV4-OUT out
router bgp 64590 neighbor 103.41.12.18 address-family ipv4 unicast next-hop-self
router bgp 64590 neighbor 103.41.12.18 address-family ipv4 unicast remove-private-AS
router bgp 64590 neighbor 103.41.12.18 address-family ipv4 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 103.41.12.19 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 103.41.12.19 remote-as 714
router bgp 64590 neighbor 103.41.12.19 timers 10 30
router bgp 64590 neighbor 103.41.12.19 description Peering:AS714-SID12202
router bgp 64590 neighbor 103.41.12.19 address-family ipv4 unicast
router bgp 64590 neighbor 103.41.12.19 address-family ipv4 unicast maximum-prefix 12500 90 restart 60
router bgp 64590 neighbor 103.41.12.19 address-family ipv4 unicast send-community-ebgp
router bgp 64590 neighbor 103.41.12.19 address-family ipv4 unicast route-policy AS714-12202-IPV4-IN in
router bgp 64590 neighbor 103.41.12.19 address-family ipv4 unicast route-policy AS714-12202-IPV4-OUT out
router bgp 64590 neighbor 103.41.12.19 address-family ipv4 unicast next-hop-self
router bgp 64590 neighbor 103.41.12.19 address-family ipv4 unicast remove-private-AS
router bgp 64590 neighbor 103.41.12.19 address-family ipv4 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 103.41.12.23 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 103.41.12.23 remote-as 8075
router bgp 64590 neighbor 103.41.12.23 timers 10 30
router bgp 64590 neighbor 103.41.12.23 description Peering:AS8075-SID12202
router bgp 64590 neighbor 103.41.12.23 address-family ipv4 unicast
router bgp 64590 neighbor 103.41.12.23 address-family ipv4 unicast maximum-prefix 2000 90 restart 60
router bgp 64590 neighbor 103.41.12.23 address-family ipv4 unicast send-community-ebgp
router bgp 64590 neighbor 103.41.12.23 address-family ipv4 unicast route-policy AS8075-12202-IPV4-IN in
router bgp 64590 neighbor 103.41.12.23 address-family ipv4 unicast route-policy AS8075-12202-IPV4-OUT out
router bgp 64590 neighbor 103.41.12.23 address-family ipv4 unicast next-hop-self
router bgp 64590 neighbor 103.41.12.23 address-family ipv4 unicast remove-private-AS
router bgp 64590 neighbor 103.41.12.23 address-family ipv4 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 103.41.12.30 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 103.41.12.30 remote-as 10310
router bgp 64590 neighbor 103.41.12.30 timers 10 30
router bgp 64590 neighbor 103.41.12.30 description Peering:AS10310-SID12202
router bgp 64590 neighbor 103.41.12.30 address-family ipv4 unicast
router bgp 64590 neighbor 103.41.12.30 address-family ipv4 unicast maximum-prefix 1250 90 restart 60
router bgp 64590 neighbor 103.41.12.30 address-family ipv4 unicast send-community-ebgp
router bgp 64590 neighbor 103.41.12.30 address-family ipv4 unicast route-policy AS10310-12202-IPV4-IN in
router bgp 64590 neighbor 103.41.12.30 address-family ipv4 unicast route-policy AS10310-12202-IPV4-OUT out
router bgp 64590 neighbor 103.41.12.30 address-family ipv4 unicast next-hop-self
router bgp 64590 neighbor 103.41.12.30 address-family ipv4 unicast remove-private-AS
router bgp 64590 neighbor 103.41.12.30 address-family ipv4 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 103.41.12.31 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 103.41.12.31 remote-as 46489
router bgp 64590 neighbor 103.41.12.31 timers 10 30
router bgp 64590 neighbor 103.41.12.31 description Peering:AS46489-SID12202
router bgp 64590 neighbor 103.41.12.31 address-family ipv4 unicast
router bgp 64590 neighbor 103.41.12.31 address-family ipv4 unicast maximum-prefix 50 90 restart 60
router bgp 64590 neighbor 103.41.12.31 address-family ipv4 unicast send-community-ebgp
router bgp 64590 neighbor 103.41.12.31 address-family ipv4 unicast route-policy AS46489-12202-IPV4-IN in
router bgp 64590 neighbor 103.41.12.31 address-family ipv4 unicast route-policy AS46489-12202-IPV4-OUT out
router bgp 64590 neighbor 103.41.12.31 address-family ipv4 unicast next-hop-self
router bgp 64590 neighbor 103.41.12.31 address-family ipv4 unicast remove-private-AS
router bgp 64590 neighbor 103.41.12.31 address-family ipv4 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 103.41.12.32 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 103.41.12.32 remote-as 46489
router bgp 64590 neighbor 103.41.12.32 timers 10 30
router bgp 64590 neighbor 103.41.12.32 description Peering:AS46489-SID12202
router bgp 64590 neighbor 103.41.12.32 address-family ipv4 unicast
router bgp 64590 neighbor 103.41.12.32 address-family ipv4 unicast maximum-prefix 50 90 restart 60
router bgp 64590 neighbor 103.41.12.32 address-family ipv4 unicast send-community-ebgp
router bgp 64590 neighbor 103.41.12.32 address-family ipv4 unicast route-policy AS46489-12202-IPV4-IN in
router bgp 64590 neighbor 103.41.12.32 address-family ipv4 unicast route-policy AS46489-12202-IPV4-OUT out
router bgp 64590 neighbor 103.41.12.32 address-family ipv4 unicast next-hop-self
router bgp 64590 neighbor 103.41.12.32 address-family ipv4 unicast remove-private-AS
router bgp 64590 neighbor 103.41.12.32 address-family ipv4 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 103.41.12.35 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 103.41.12.35 remote-as 199524
router bgp 64590 neighbor 103.41.12.35 timers 10 30
router bgp 64590 neighbor 103.41.12.35 description Peering:AS199524-SID12202
router bgp 64590 neighbor 103.41.12.35 address-family ipv4 unicast
router bgp 64590 neighbor 103.41.12.35 address-family ipv4 unicast maximum-prefix 1250 90 restart 60
router bgp 64590 neighbor 103.41.12.35 address-family ipv4 unicast send-community-ebgp
router bgp 64590 neighbor 103.41.12.35 address-family ipv4 unicast route-policy AS199524-12202-IPV4-IN in
router bgp 64590 neighbor 103.41.12.35 address-family ipv4 unicast route-policy AS199524-12202-IPV4-OUT out
router bgp 64590 neighbor 103.41.12.35 address-family ipv4 unicast next-hop-self
router bgp 64590 neighbor 103.41.12.35 address-family ipv4 unicast remove-private-AS
router bgp 64590 neighbor 103.41.12.35 address-family ipv4 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 103.41.12.41 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 103.41.12.41 remote-as 55256
router bgp 64590 neighbor 103.41.12.41 timers 10 30
router bgp 64590 neighbor 103.41.12.41 description Peering:AS55256-SID12202
router bgp 64590 neighbor 103.41.12.41 address-family ipv4 unicast
router bgp 64590 neighbor 103.41.12.41 address-family ipv4 unicast maximum-prefix 50 90 restart 60
router bgp 64590 neighbor 103.41.12.41 address-family ipv4 unicast send-community-ebgp
router bgp 64590 neighbor 103.41.12.41 address-family ipv4 unicast route-policy AS55256-12202-IPV4-IN in
router bgp 64590 neighbor 103.41.12.41 address-family ipv4 unicast route-policy AS55256-12202-IPV4-OUT out
router bgp 64590 neighbor 103.41.12.41 address-family ipv4 unicast next-hop-self
router bgp 64590 neighbor 103.41.12.41 address-family ipv4 unicast remove-private-AS
router bgp 64590 neighbor 103.41.12.41 address-family ipv4 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 103.41.12.44 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 103.41.12.44 remote-as 20940
router bgp 64590 neighbor 103.41.12.44 timers 10 30
router bgp 64590 neighbor 103.41.12.44 description Peering:AS20940-SID12202
router bgp 64590 neighbor 103.41.12.44 address-family ipv4 unicast
router bgp 64590 neighbor 103.41.12.44 address-family ipv4 unicast maximum-prefix 15000 90 restart 60
router bgp 64590 neighbor 103.41.12.44 address-family ipv4 unicast send-community-ebgp
router bgp 64590 neighbor 103.41.12.44 address-family ipv4 unicast route-policy AS20940-12202-IPV4-IN in
router bgp 64590 neighbor 103.41.12.44 address-family ipv4 unicast route-policy AS20940-12202-IPV4-OUT out
router bgp 64590 neighbor 103.41.12.44 address-family ipv4 unicast next-hop-self
router bgp 64590 neighbor 103.41.12.44 address-family ipv4 unicast remove-private-AS
router bgp 64590 neighbor 103.41.12.44 address-family ipv4 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 103.41.12.45 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 103.41.12.45 remote-as 49544
router bgp 64590 neighbor 103.41.12.45 timers 10 30
router bgp 64590 neighbor 103.41.12.45 description Peering:AS49544-SID12202
router bgp 64590 neighbor 103.41.12.45 address-family ipv4 unicast
router bgp 64590 neighbor 103.41.12.45 address-family ipv4 unicast maximum-prefix 625 90 restart 60
router bgp 64590 neighbor 103.41.12.45 address-family ipv4 unicast send-community-ebgp
router bgp 64590 neighbor 103.41.12.45 address-family ipv4 unicast route-policy AS49544-12202-IPV4-IN in
router bgp 64590 neighbor 103.41.12.45 address-family ipv4 unicast route-policy AS49544-12202-IPV4-OUT out
router bgp 64590 neighbor 103.41.12.45 address-family ipv4 unicast next-hop-self
router bgp 64590 neighbor 103.41.12.45 address-family ipv4 unicast remove-private-AS
router bgp 64590 neighbor 103.41.12.45 address-family ipv4 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 2001:ded::10 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 2001:ded::10 remote-as 22822
router bgp 64590 neighbor 2001:ded::10 timers 10 30
router bgp 64590 neighbor 2001:ded::10 shutdown
router bgp 64590 neighbor 2001:ded::10 description Peering:AS22822-SID12202
router bgp 64590 neighbor 2001:ded::10 address-family ipv6 unicast
router bgp 64590 neighbor 2001:ded::10 address-family ipv6 unicast maximum-prefix 625 90 restart 60
router bgp 64590 neighbor 2001:ded::10 address-family ipv6 unicast send-community-ebgp
router bgp 64590 neighbor 2001:ded::10 address-family ipv6 unicast route-policy AS22822-12202-IPV6-IN in
router bgp 64590 neighbor 2001:ded::10 address-family ipv6 unicast route-policy AS22822-12202-IPV6-OUT out
router bgp 64590 neighbor 2001:ded::10 address-family ipv6 unicast next-hop-self
router bgp 64590 neighbor 2001:ded::10 address-family ipv6 unicast remove-private-AS
router bgp 64590 neighbor 2001:ded::10 address-family ipv6 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 2001:ded::12 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 2001:ded::12 remote-as 20940
router bgp 64590 neighbor 2001:ded::12 timers 10 30
router bgp 64590 neighbor 2001:ded::12 description Peering:AS20940-SID12202
router bgp 64590 neighbor 2001:ded::12 address-family ipv6 unicast
router bgp 64590 neighbor 2001:ded::12 address-family ipv6 unicast maximum-prefix 6250 90 restart 60
router bgp 64590 neighbor 2001:ded::12 address-family ipv6 unicast send-community-ebgp
router bgp 64590 neighbor 2001:ded::12 address-family ipv6 unicast route-policy AS20940-12202-IPV6-IN in
router bgp 64590 neighbor 2001:ded::12 address-family ipv6 unicast route-policy AS20940-12202-IPV6-OUT out
router bgp 64590 neighbor 2001:ded::12 address-family ipv6 unicast next-hop-self
router bgp 64590 neighbor 2001:ded::12 address-family ipv6 unicast remove-private-AS
router bgp 64590 neighbor 2001:ded::12 address-family ipv6 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 2001:ded::13 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 2001:ded::13 remote-as 714
router bgp 64590 neighbor 2001:ded::13 timers 10 30
router bgp 64590 neighbor 2001:ded::13 description Peering:AS714-SID12202
router bgp 64590 neighbor 2001:ded::13 address-family ipv6 unicast
router bgp 64590 neighbor 2001:ded::13 address-family ipv6 unicast maximum-prefix 1250 90 restart 60
router bgp 64590 neighbor 2001:ded::13 address-family ipv6 unicast send-community-ebgp
router bgp 64590 neighbor 2001:ded::13 address-family ipv6 unicast route-policy AS714-12202-IPV6-IN in
router bgp 64590 neighbor 2001:ded::13 address-family ipv6 unicast route-policy AS714-12202-IPV6-OUT out
router bgp 64590 neighbor 2001:ded::13 address-family ipv6 unicast next-hop-self
router bgp 64590 neighbor 2001:ded::13 address-family ipv6 unicast remove-private-AS
router bgp 64590 neighbor 2001:ded::13 address-family ipv6 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 2001:ded::17 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 2001:ded::17 remote-as 8075
router bgp 64590 neighbor 2001:ded::17 timers 10 30
router bgp 64590 neighbor 2001:ded::17 description Peering:AS8075-SID12202
router bgp 64590 neighbor 2001:ded::17 address-family ipv6 unicast
router bgp 64590 neighbor 2001:ded::17 address-family ipv6 unicast maximum-prefix 625 90 restart 60
router bgp 64590 neighbor 2001:ded::17 address-family ipv6 unicast send-community-ebgp
router bgp 64590 neighbor 2001:ded::17 address-family ipv6 unicast route-policy AS8075-12202-IPV6-IN in
router bgp 64590 neighbor 2001:ded::17 address-family ipv6 unicast route-policy AS8075-12202-IPV6-OUT out
router bgp 64590 neighbor 2001:ded::17 address-family ipv6 unicast next-hop-self
router bgp 64590 neighbor 2001:ded::17 address-family ipv6 unicast remove-private-AS
router bgp 64590 neighbor 2001:ded::17 address-family ipv6 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 2001:ded::1e local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 2001:ded::1e remote-as 10310
router bgp 64590 neighbor 2001:ded::1e timers 10 30
router bgp 64590 neighbor 2001:ded::1e description Peering:AS10310-SID12202
router bgp 64590 neighbor 2001:ded::1e address-family ipv6 unicast
router bgp 64590 neighbor 2001:ded::1e address-family ipv6 unicast maximum-prefix 625 90 restart 60
router bgp 64590 neighbor 2001:ded::1e address-family ipv6 unicast send-community-ebgp
router bgp 64590 neighbor 2001:ded::1e address-family ipv6 unicast route-policy AS10310-12202-IPV6-IN in
router bgp 64590 neighbor 2001:ded::1e address-family ipv6 unicast route-policy AS10310-12202-IPV6-OUT out
router bgp 64590 neighbor 2001:ded::1e address-family ipv6 unicast next-hop-self
router bgp 64590 neighbor 2001:ded::1e address-family ipv6 unicast remove-private-AS
router bgp 64590 neighbor 2001:ded::1e address-family ipv6 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 2001:ded::1f local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 2001:ded::1f remote-as 46489
router bgp 64590 neighbor 2001:ded::1f timers 10 30
router bgp 64590 neighbor 2001:ded::1f description Peering:AS46489-SID12202
router bgp 64590 neighbor 2001:ded::1f address-family ipv6 unicast
router bgp 64590 neighbor 2001:ded::1f address-family ipv6 unicast maximum-prefix 12 90 restart 60
router bgp 64590 neighbor 2001:ded::1f address-family ipv6 unicast send-community-ebgp
router bgp 64590 neighbor 2001:ded::1f address-family ipv6 unicast route-policy AS46489-12202-IPV6-IN in
router bgp 64590 neighbor 2001:ded::1f address-family ipv6 unicast route-policy AS46489-12202-IPV6-OUT out
router bgp 64590 neighbor 2001:ded::1f address-family ipv6 unicast next-hop-self
router bgp 64590 neighbor 2001:ded::1f address-family ipv6 unicast remove-private-AS
router bgp 64590 neighbor 2001:ded::1f address-family ipv6 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 2001:ded::20 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 2001:ded::20 remote-as 46489
router bgp 64590 neighbor 2001:ded::20 timers 10 30
router bgp 64590 neighbor 2001:ded::20 description Peering:AS46489-SID12202
router bgp 64590 neighbor 2001:ded::20 address-family ipv6 unicast
router bgp 64590 neighbor 2001:ded::20 address-family ipv6 unicast maximum-prefix 12 90 restart 60
router bgp 64590 neighbor 2001:ded::20 address-family ipv6 unicast send-community-ebgp
router bgp 64590 neighbor 2001:ded::20 address-family ipv6 unicast route-policy AS46489-12202-IPV6-IN in
router bgp 64590 neighbor 2001:ded::20 address-family ipv6 unicast route-policy AS46489-12202-IPV6-OUT out
router bgp 64590 neighbor 2001:ded::20 address-family ipv6 unicast next-hop-self
router bgp 64590 neighbor 2001:ded::20 address-family ipv6 unicast remove-private-AS
router bgp 64590 neighbor 2001:ded::20 address-family ipv6 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 2001:ded::23 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 2001:ded::23 remote-as 199524
router bgp 64590 neighbor 2001:ded::23 timers 10 30
router bgp 64590 neighbor 2001:ded::23 description Peering:AS199524-SID12202
router bgp 64590 neighbor 2001:ded::23 address-family ipv6 unicast
router bgp 64590 neighbor 2001:ded::23 address-family ipv6 unicast maximum-prefix 63 90 restart 60
router bgp 64590 neighbor 2001:ded::23 address-family ipv6 unicast send-community-ebgp
router bgp 64590 neighbor 2001:ded::23 address-family ipv6 unicast route-policy AS199524-12202-IPV6-IN in
router bgp 64590 neighbor 2001:ded::23 address-family ipv6 unicast route-policy AS199524-12202-IPV6-OUT out
router bgp 64590 neighbor 2001:ded::23 address-family ipv6 unicast next-hop-self
router bgp 64590 neighbor 2001:ded::23 address-family ipv6 unicast remove-private-AS
router bgp 64590 neighbor 2001:ded::23 address-family ipv6 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 2001:ded::29 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 2001:ded::29 remote-as 55256
router bgp 64590 neighbor 2001:ded::29 timers 10 30
router bgp 64590 neighbor 2001:ded::29 description Peering:AS55256-SID12202
router bgp 64590 neighbor 2001:ded::29 address-family ipv6 unicast
router bgp 64590 neighbor 2001:ded::29 address-family ipv6 unicast maximum-prefix 25 90 restart 60
router bgp 64590 neighbor 2001:ded::29 address-family ipv6 unicast send-community-ebgp
router bgp 64590 neighbor 2001:ded::29 address-family ipv6 unicast route-policy AS55256-12202-IPV6-IN in
router bgp 64590 neighbor 2001:ded::29 address-family ipv6 unicast route-policy AS55256-12202-IPV6-OUT out
router bgp 64590 neighbor 2001:ded::29 address-family ipv6 unicast next-hop-self
router bgp 64590 neighbor 2001:ded::29 address-family ipv6 unicast remove-private-AS
router bgp 64590 neighbor 2001:ded::29 address-family ipv6 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 2001:ded::2c local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 2001:ded::2c remote-as 20940
router bgp 64590 neighbor 2001:ded::2c timers 10 30
router bgp 64590 neighbor 2001:ded::2c description Peering:AS20940-SID12202
router bgp 64590 neighbor 2001:ded::2c address-family ipv6 unicast
router bgp 64590 neighbor 2001:ded::2c address-family ipv6 unicast maximum-prefix 6250 90 restart 60
router bgp 64590 neighbor 2001:ded::2c address-family ipv6 unicast send-community-ebgp
router bgp 64590 neighbor 2001:ded::2c address-family ipv6 unicast route-policy AS20940-12202-IPV6-IN in
router bgp 64590 neighbor 2001:ded::2c address-family ipv6 unicast route-policy AS20940-12202-IPV6-OUT out
router bgp 64590 neighbor 2001:ded::2c address-family ipv6 unicast next-hop-self
router bgp 64590 neighbor 2001:ded::2c address-family ipv6 unicast remove-private-AS
router bgp 64590 neighbor 2001:ded::2c address-family ipv6 unicast soft-reconfiguration inbound always
router bgp 64590 neighbor 2001:ded::2d local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 2001:ded::2d remote-as 49544
router bgp 64590 neighbor 2001:ded::2d timers 10 30
router bgp 64590 neighbor 2001:ded::2d description Peering:AS49544-SID12202
router bgp 64590 neighbor 2001:ded::2d address-family ipv6 unicast
router bgp 64590 neighbor 2001:ded::2d address-family ipv6 unicast maximum-prefix 250 90 restart 60
router bgp 64590 neighbor 2001:ded::2d address-family ipv6 unicast send-community-ebgp
router bgp 64590 neighbor 2001:ded::2d address-family ipv6 unicast route-policy AS49544-12202-IPV6-IN in
router bgp 64590 neighbor 2001:ded::2d address-family ipv6 unicast route-policy AS49544-12202-IPV6-OUT out
router bgp 64590 neighbor 2001:ded::2d address-family ipv6 unicast next-hop-self
router bgp 64590 neighbor 2001:ded::2d address-family ipv6 unicast remove-private-AS
router bgp 64590 neighbor 2001:ded::2d address-family ipv6 unicast soft-reconfiguration inbound always




route-policy AS6939-12202-IPV4-IN-10
  if (destination in AS6939-IPV4-PREFIX-DROP-IN) or (as-path in AS6939-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS6939-12202-IPV4-IN-20
  if (destination in AS38195-SLCAU-LE24) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9705, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    if (community matches-any AS6939-VIA-SINGAPORE) then
      set local-preference 152
    else
      set local-preference 110
    endif
    done
  else
    drop
  endif
end-policy
!
route-policy AS6939-12202-IPV6-IN-10
  if (destination in AS6939-IPV6-PREFIX-DROP-IN) or (as-path in AS6939-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS6939-12202-IPV6-IN-20
  if (destination in AS38195-SLCAU-LE48) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9705, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    if (community matches-any AS6939-VIA-SINGAPORE) then
      set local-preference 152
    else
      set local-preference 110
    endif
    done
  else
    drop
  endif
end-policy
!
route-policy AS714-12202-IPV4-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS714-IPV4-PREFIX-DROP-OUT) or (as-path in AS714-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS714-12202-IPV4-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS714-12202-IPV4-OUT-20
  if (large-community matches-any (38195:714:0)) or (large-community matches-any (38195:714:68000)) or (large-community matches-any (38195:714:80100)) then
    drop
  endif
end-policy
!
route-policy AS714-12202-IPV4-OUT-21
  if (large-community matches-any (38195:714:1)) or (large-community matches-any (38195:714:68001)) or (large-community matches-any (38195:714:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:714:2)) or (large-community matches-any (38195:714:68002)) or (large-community matches-any (38195:714:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:714:3)) or (large-community matches-any (38195:714:68003)) or (large-community matches-any (38195:714:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS714-12202-IPV4-OUT-29
  if (large-community matches-any (38195:714:9)) or (large-community matches-any (38195:714:68009)) or (large-community matches-any (38195:714:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS714-12202-IPV4-OUT-90
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CORE) then
    done
  endif
end-policy
!
route-policy AS714-12202-IPV4-OUT-91
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    done
  endif
end-policy
!
route-policy AS714-12202-IPV4-OUT-99
  drop
end-policy
!
route-policy AS714-12202-IPV6-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS714-IPV6-PREFIX-DROP-OUT) or (as-path in AS714-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS714-12202-IPV6-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS714-12202-IPV6-OUT-20
  if (large-community matches-any (38195:714:0)) or (large-community matches-any (38195:714:68000)) or (large-community matches-any (38195:714:80100)) then
    drop
  endif
end-policy
!
route-policy AS714-12202-IPV6-OUT-21
  if (large-community matches-any (38195:714:1)) or (large-community matches-any (38195:714:68001)) or (large-community matches-any (38195:714:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:714:2)) or (large-community matches-any (38195:714:68002)) or (large-community matches-any (38195:714:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:714:3)) or (large-community matches-any (38195:714:68003)) or (large-community matches-any (38195:714:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS714-12202-IPV6-OUT-29
  if (large-community matches-any (38195:714:9)) or (large-community matches-any (38195:714:68009)) or (large-community matches-any (38195:714:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS714-12202-IPV6-OUT-90
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CORE) then
    done
  endif
end-policy
!
route-policy AS714-12202-IPV6-OUT-91
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    done
  endif
end-policy
!
route-policy AS714-12202-IPV6-OUT-99
  drop
end-policy
!
route-policy AS8075-12202-IPV4-IN-10
  if (destination in AS8075-IPV4-PREFIX-DROP-IN) or (as-path in AS8075-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS8075-12202-IPV4-IN-20
  if (destination in AS38195-SLCAU-LE24) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9705, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    set local-preference 152
    done
  else
    drop
  endif
end-policy
!
route-policy AS8075-12202-IPV6-IN-10
  if (destination in AS8075-IPV6-PREFIX-DROP-IN) or (as-path in AS8075-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS8075-12202-IPV6-IN-20
  if (destination in AS38195-SLCAU-LE48) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9705, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    set local-preference 152
    done
  else
    drop
  endif
end-policy
!
route-policy AS10310-12202-IPV4-IN-10
  if (destination in AS10310-IPV4-PREFIX-DROP-IN) or (as-path in AS10310-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS10310-12202-IPV4-IN-20
  if (destination in AS38195-SLCAU-LE24) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9705, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    set local-preference 152
    done
  else
    drop
  endif
end-policy
!
route-policy AS10310-12202-IPV6-IN-10
  if (destination in AS10310-IPV6-PREFIX-DROP-IN) or (as-path in AS10310-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS10310-12202-IPV6-IN-20
  if (destination in AS38195-SLCAU-LE48) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9705, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    set local-preference 152
    done
  else
    drop
  endif
end-policy
!
route-policy AS13335-12202-IPV4-IN-10
  if (destination in AS13335-IPV4-PREFIX-DROP-IN) or (as-path in AS13335-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS13335-12202-IPV4-IN-20
  if (destination in AS38195-SLCAU-LE24) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9705, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    set local-preference 152
    done
  else
    drop
  endif
end-policy
!
route-policy AS13335-12202-IPV6-IN-10
  if (destination in AS13335-IPV6-PREFIX-DROP-IN) or (as-path in AS13335-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS13335-12202-IPV6-IN-20
  if (destination in AS38195-SLCAU-LE48) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9705, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    set local-preference 152
    done
  else
    drop
  endif
end-policy
!
route-policy AS15133-12202-IPV4-IN-10
  if (destination in AS15133-IPV4-PREFIX-DROP-IN) or (as-path in AS15133-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS15133-12202-IPV4-IN-20
  if (destination in AS38195-SLCAU-LE24) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9705, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    set local-preference 152
    done
  else
    drop
  endif
end-policy
!
route-policy AS15133-12202-IPV6-IN-10
  if (destination in AS15133-IPV6-PREFIX-DROP-IN) or (as-path in AS15133-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS15133-12202-IPV6-IN-20
  if (destination in AS38195-SLCAU-LE48) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9705, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    set local-preference 152
    done
  else
    drop
  endif
end-policy
!
route-policy AS20940-12202-IPV4-IN-10
  if (destination in AS20940-IPV4-PREFIX-DROP-IN) or (as-path in AS20940-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS20940-12202-IPV4-IN-20
  if (destination in AS38195-SLCAU-LE24) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9705, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    set local-preference 152
    done
  else
    drop
  endif
end-policy
!
route-policy AS20940-12202-IPV6-IN-10
  if (destination in AS20940-IPV6-PREFIX-DROP-IN) or (as-path in AS20940-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS20940-12202-IPV6-IN-20
  if (destination in AS38195-SLCAU-LE48) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9705, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    set local-preference 152
    done
  else
    drop
  endif
end-policy
!
route-policy AS22822-12202-IPV4-IN-10
  if (destination in AS22822-IPV4-PREFIX-DROP-IN) or (as-path in AS22822-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS22822-12202-IPV4-IN-20
  if (destination in AS38195-SLCAU-LE24) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9705, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    set local-preference 152
    done
  else
    drop
  endif
end-policy
!
route-policy AS22822-12202-IPV6-IN-10
  if (destination in AS22822-IPV6-PREFIX-DROP-IN) or (as-path in AS22822-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS22822-12202-IPV6-IN-20
  if (destination in AS38195-SLCAU-LE48) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9705, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    set local-preference 152
    done
  else
    drop
  endif
end-policy
!
route-policy AS46489-12202-IPV4-IN-10
  if (destination in AS46489-IPV4-PREFIX-DROP-IN) or (as-path in AS46489-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS46489-12202-IPV4-IN-20
  if (destination in AS38195-SLCAU-LE24) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9705, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    set local-preference 152
    done
  else
    drop
  endif
end-policy
!
route-policy AS46489-12202-IPV6-IN-10
  if (destination in AS46489-IPV6-PREFIX-DROP-IN) or (as-path in AS46489-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS46489-12202-IPV6-IN-20
  if (destination in AS38195-SLCAU-LE48) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9705, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    set local-preference 152
    done
  else
    drop
  endif
end-policy
!
route-policy AS49544-12202-IPV4-IN-10
  if (destination in AS49544-IPV4-PREFIX-DROP-IN) or (as-path in AS49544-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS49544-12202-IPV4-IN-20
  if (destination in AS38195-SLCAU-LE24) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9705, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    set local-preference 152
    done
  else
    drop
  endif
end-policy
!
route-policy AS49544-12202-IPV6-IN-10
  if (destination in AS49544-IPV6-PREFIX-DROP-IN) or (as-path in AS49544-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS49544-12202-IPV6-IN-20
  if (destination in AS38195-SLCAU-LE48) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9705, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    set local-preference 152
    done
  else
    drop
  endif
end-policy
!
route-policy AS55256-12202-IPV4-IN-10
  if (destination in AS55256-IPV4-PREFIX-DROP-IN) or (as-path in AS55256-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS55256-12202-IPV4-IN-20
  if (destination in AS38195-SLCAU-LE24) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9705, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    set local-preference 152
    done
  else
    drop
  endif
end-policy
!
route-policy AS55256-12202-IPV6-IN-10
  if (destination in AS55256-IPV6-PREFIX-DROP-IN) or (as-path in AS55256-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS55256-12202-IPV6-IN-20
  if (destination in AS38195-SLCAU-LE48) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9705, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    set local-preference 152
    done
  else
    drop
  endif
end-policy
!
route-policy AS63832-12202-IPV4-IN-10
  if (destination in AS63832-IPV4-PREFIX-DROP-IN) or (as-path in AS63832-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS63832-12202-IPV4-IN-20
  if (destination in AS38195-SLCAU-LE24) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9706, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    set local-preference 152
    done
  else
    drop
  endif
end-policy
!
route-policy AS63832-12202-IPV6-IN-10
  if (destination in AS63832-IPV6-PREFIX-DROP-IN) or (as-path in AS63832-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS63832-12202-IPV6-IN-20
  if (destination in AS38195-SLCAU-LE48) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9706, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    set local-preference 152
    done
  else
    drop
  endif
end-policy
!
route-policy AS6939-12202-IPV4-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS6939-IPV4-PREFIX-DROP-OUT) or (as-path in AS6939-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS6939-12202-IPV4-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS6939-12202-IPV4-OUT-20
  if (large-community matches-any (38195:6939:0)) or (large-community matches-any (38195:6939:68000)) or (large-community matches-any (38195:6939:80100)) then
    drop
  endif
end-policy
!
route-policy AS6939-12202-IPV4-OUT-21
  if (large-community matches-any (38195:6939:1)) or (large-community matches-any (38195:6939:68001)) or (large-community matches-any (38195:6939:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:6939:2)) or (large-community matches-any (38195:6939:68002)) or (large-community matches-any (38195:6939:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:6939:3)) or (large-community matches-any (38195:6939:68003)) or (large-community matches-any (38195:6939:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS6939-12202-IPV4-OUT-29
  if (large-community matches-any (38195:6939:9)) or (large-community matches-any (38195:6939:68009)) or (large-community matches-any (38195:6939:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS6939-12202-IPV4-OUT-30
  if (community matches-any (38195:18060)) then
    drop
  endif
end-policy
!
route-policy AS6939-12202-IPV4-OUT-31
  if (community matches-any (38195:18061)) then
    prepend as-path 38195
    pass
  endif
  if (community matches-any (38195:18062)) then
    prepend as-path 38195 2
    pass
  endif
  if (community matches-any (38195:18063)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS6939-12202-IPV4-OUT-39
  if (community matches-any (38195:18069)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS6939-12202-IPV4-OUT-90
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CORE) then
    done
  endif
end-policy
!
route-policy AS6939-12202-IPV4-OUT-91
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    done
  endif
end-policy
!
route-policy AS6939-12202-IPV4-OUT-99
  drop
end-policy
!
route-policy AS6939-12202-IPV6-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS6939-IPV6-PREFIX-DROP-OUT) or (as-path in AS6939-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS6939-12202-IPV6-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS6939-12202-IPV6-OUT-20
  if (large-community matches-any (38195:6939:0)) or (large-community matches-any (38195:6939:68000)) or (large-community matches-any (38195:6939:80100)) then
    drop
  endif
end-policy
!
route-policy AS6939-12202-IPV6-OUT-21
  if (large-community matches-any (38195:6939:1)) or (large-community matches-any (38195:6939:68001)) or (large-community matches-any (38195:6939:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:6939:2)) or (large-community matches-any (38195:6939:68002)) or (large-community matches-any (38195:6939:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:6939:3)) or (large-community matches-any (38195:6939:68003)) or (large-community matches-any (38195:6939:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS6939-12202-IPV6-OUT-29
  if (large-community matches-any (38195:6939:9)) or (large-community matches-any (38195:6939:68009)) or (large-community matches-any (38195:6939:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS6939-12202-IPV6-OUT-30
  if (community matches-any (38195:18060)) then
    drop
  endif
end-policy
!
route-policy AS6939-12202-IPV6-OUT-31
  if (community matches-any (38195:18061)) then
    prepend as-path 38195
    pass
  endif
  if (community matches-any (38195:18062)) then
    prepend as-path 38195 2
    pass
  endif
  if (community matches-any (38195:18063)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS6939-12202-IPV6-OUT-39
  if (community matches-any (38195:18069)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS6939-12202-IPV6-OUT-90
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CORE) then
    done
  endif
end-policy
!
route-policy AS6939-12202-IPV6-OUT-91
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    done
  endif
end-policy
!
route-policy AS6939-12202-IPV6-OUT-99
  drop
end-policy
!
route-policy AS8075-12202-IPV4-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS8075-IPV4-PREFIX-DROP-OUT) or (as-path in AS8075-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS8075-12202-IPV4-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS8075-12202-IPV4-OUT-20
  if (large-community matches-any (38195:8075:0)) or (large-community matches-any (38195:8075:68000)) or (large-community matches-any (38195:8075:80100)) then
    drop
  endif
end-policy
!
route-policy AS8075-12202-IPV4-OUT-21
  if (large-community matches-any (38195:8075:1)) or (large-community matches-any (38195:8075:68001)) or (large-community matches-any (38195:8075:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:8075:2)) or (large-community matches-any (38195:8075:68002)) or (large-community matches-any (38195:8075:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:8075:3)) or (large-community matches-any (38195:8075:68003)) or (large-community matches-any (38195:8075:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS8075-12202-IPV4-OUT-29
  if (large-community matches-any (38195:8075:9)) or (large-community matches-any (38195:8075:68009)) or (large-community matches-any (38195:8075:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS8075-12202-IPV4-OUT-30
  if (community matches-any (38195:11910)) then
    drop
  endif
end-policy
!
route-policy AS8075-12202-IPV4-OUT-31
  if (community matches-any (38195:11911)) then
    prepend as-path 38195
    pass
  endif
  if (community matches-any (38195:11912)) then
    prepend as-path 38195 2
    pass
  endif
  if (community matches-any (38195:11913)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS8075-12202-IPV4-OUT-39
  if (community matches-any (38195:11919)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS8075-12202-IPV4-OUT-90
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CORE) then
    done
  endif
end-policy
!
route-policy AS8075-12202-IPV4-OUT-91
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    done
  endif
end-policy
!
route-policy AS8075-12202-IPV4-OUT-99
  drop
end-policy
!
route-policy AS8075-12202-IPV6-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS8075-IPV6-PREFIX-DROP-OUT) or (as-path in AS8075-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS8075-12202-IPV6-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS8075-12202-IPV6-OUT-20
  if (large-community matches-any (38195:8075:0)) or (large-community matches-any (38195:8075:68000)) or (large-community matches-any (38195:8075:80100)) then
    drop
  endif
end-policy
!
route-policy AS8075-12202-IPV6-OUT-21
  if (large-community matches-any (38195:8075:1)) or (large-community matches-any (38195:8075:68001)) or (large-community matches-any (38195:8075:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:8075:2)) or (large-community matches-any (38195:8075:68002)) or (large-community matches-any (38195:8075:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:8075:3)) or (large-community matches-any (38195:8075:68003)) or (large-community matches-any (38195:8075:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS8075-12202-IPV6-OUT-29
  if (large-community matches-any (38195:8075:9)) or (large-community matches-any (38195:8075:68009)) or (large-community matches-any (38195:8075:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS8075-12202-IPV6-OUT-30
  if (community matches-any (38195:11910)) then
    drop
  endif
end-policy
!
route-policy AS8075-12202-IPV6-OUT-31
  if (community matches-any (38195:11911)) then
    prepend as-path 38195
    pass
  endif
  if (community matches-any (38195:11912)) then
    prepend as-path 38195 2
    pass
  endif
  if (community matches-any (38195:11913)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS8075-12202-IPV6-OUT-39
  if (community matches-any (38195:11919)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS8075-12202-IPV6-OUT-90
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CORE) then
    done
  endif
end-policy
!
route-policy AS8075-12202-IPV6-OUT-91
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    done
  endif
end-policy
!
route-policy AS8075-12202-IPV6-OUT-99
  drop
end-policy
!
route-policy AS10310-12202-IPV4-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS10310-IPV4-PREFIX-DROP-OUT) or (as-path in AS10310-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS10310-12202-IPV4-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS10310-12202-IPV4-OUT-20
  if (large-community matches-any (38195:10310:0)) or (large-community matches-any (38195:10310:68000)) or (large-community matches-any (38195:10310:80100)) then
    drop
  endif
end-policy
!
route-policy AS10310-12202-IPV4-OUT-21
  if (large-community matches-any (38195:10310:1)) or (large-community matches-any (38195:10310:68001)) or (large-community matches-any (38195:10310:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:10310:2)) or (large-community matches-any (38195:10310:68002)) or (large-community matches-any (38195:10310:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:10310:3)) or (large-community matches-any (38195:10310:68003)) or (large-community matches-any (38195:10310:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS10310-12202-IPV4-OUT-29
  if (large-community matches-any (38195:10310:9)) or (large-community matches-any (38195:10310:68009)) or (large-community matches-any (38195:10310:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS10310-12202-IPV4-OUT-30
  if (community matches-any (38195:11950)) then
    drop
  endif
end-policy
!
route-policy AS10310-12202-IPV4-OUT-31
  if (community matches-any (38195:11951)) then
    prepend as-path 38195
    pass
  endif
  if (community matches-any (38195:11952)) then
    prepend as-path 38195 2
    pass
  endif
  if (community matches-any (38195:11953)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS10310-12202-IPV4-OUT-39
  if (community matches-any (38195:11959)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS10310-12202-IPV4-OUT-90
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CORE) then
    done
  endif
end-policy
!
route-policy AS10310-12202-IPV4-OUT-91
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    done
  endif
end-policy
!
route-policy AS10310-12202-IPV4-OUT-99
  drop
end-policy
!
route-policy AS10310-12202-IPV6-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS10310-IPV6-PREFIX-DROP-OUT) or (as-path in AS10310-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS10310-12202-IPV6-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS10310-12202-IPV6-OUT-20
  if (large-community matches-any (38195:10310:0)) or (large-community matches-any (38195:10310:68000)) or (large-community matches-any (38195:10310:80100)) then
    drop
  endif
end-policy
!
route-policy AS10310-12202-IPV6-OUT-21
  if (large-community matches-any (38195:10310:1)) or (large-community matches-any (38195:10310:68001)) or (large-community matches-any (38195:10310:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:10310:2)) or (large-community matches-any (38195:10310:68002)) or (large-community matches-any (38195:10310:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:10310:3)) or (large-community matches-any (38195:10310:68003)) or (large-community matches-any (38195:10310:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS10310-12202-IPV6-OUT-29
  if (large-community matches-any (38195:10310:9)) or (large-community matches-any (38195:10310:68009)) or (large-community matches-any (38195:10310:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS10310-12202-IPV6-OUT-30
  if (community matches-any (38195:11950)) then
    drop
  endif
end-policy
!
route-policy AS10310-12202-IPV6-OUT-31
  if (community matches-any (38195:11951)) then
    prepend as-path 38195
    pass
  endif
  if (community matches-any (38195:11952)) then
    prepend as-path 38195 2
    pass
  endif
  if (community matches-any (38195:11953)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS10310-12202-IPV6-OUT-39
  if (community matches-any (38195:11959)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS10310-12202-IPV6-OUT-90
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CORE) then
    done
  endif
end-policy
!
route-policy AS10310-12202-IPV6-OUT-91
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    done
  endif
end-policy
!
route-policy AS10310-12202-IPV6-OUT-99
  drop
end-policy
!
route-policy AS13335-12202-IPV4-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS13335-IPV4-PREFIX-DROP-OUT) or (as-path in AS13335-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS13335-12202-IPV4-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS13335-12202-IPV4-OUT-20
  if (large-community matches-any (38195:13335:0)) or (large-community matches-any (38195:13335:68000)) or (large-community matches-any (38195:13335:80100)) then
    drop
  endif
end-policy
!
route-policy AS13335-12202-IPV4-OUT-21
  if (large-community matches-any (38195:13335:1)) or (large-community matches-any (38195:13335:68001)) or (large-community matches-any (38195:13335:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:13335:2)) or (large-community matches-any (38195:13335:68002)) or (large-community matches-any (38195:13335:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:13335:3)) or (large-community matches-any (38195:13335:68003)) or (large-community matches-any (38195:13335:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS13335-12202-IPV4-OUT-29
  if (large-community matches-any (38195:13335:9)) or (large-community matches-any (38195:13335:68009)) or (large-community matches-any (38195:13335:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS13335-12202-IPV4-OUT-30
  if (community matches-any (38195:10940)) then
    drop
  endif
end-policy
!
route-policy AS13335-12202-IPV4-OUT-31
  if (community matches-any (38195:10941)) then
    prepend as-path 38195
    pass
  endif
  if (community matches-any (38195:10942)) then
    prepend as-path 38195 2
    pass
  endif
  if (community matches-any (38195:10943)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS13335-12202-IPV4-OUT-39
  if (community matches-any (38195:10949)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS13335-12202-IPV4-OUT-90
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CORE) then
    done
  endif
end-policy
!
route-policy AS13335-12202-IPV4-OUT-91
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    done
  endif
end-policy
!
route-policy AS13335-12202-IPV4-OUT-99
  drop
end-policy
!
route-policy AS13335-12202-IPV6-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS13335-IPV6-PREFIX-DROP-OUT) or (as-path in AS13335-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS13335-12202-IPV6-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS13335-12202-IPV6-OUT-20
  if (large-community matches-any (38195:13335:0)) or (large-community matches-any (38195:13335:68000)) or (large-community matches-any (38195:13335:80100)) then
    drop
  endif
end-policy
!
route-policy AS13335-12202-IPV6-OUT-21
  if (large-community matches-any (38195:13335:1)) or (large-community matches-any (38195:13335:68001)) or (large-community matches-any (38195:13335:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:13335:2)) or (large-community matches-any (38195:13335:68002)) or (large-community matches-any (38195:13335:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:13335:3)) or (large-community matches-any (38195:13335:68003)) or (large-community matches-any (38195:13335:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS13335-12202-IPV6-OUT-29
  if (large-community matches-any (38195:13335:9)) or (large-community matches-any (38195:13335:68009)) or (large-community matches-any (38195:13335:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS13335-12202-IPV6-OUT-30
  if (community matches-any (38195:10940)) then
    drop
  endif
end-policy
!
route-policy AS13335-12202-IPV6-OUT-31
  if (community matches-any (38195:10941)) then
    prepend as-path 38195
    pass
  endif
  if (community matches-any (38195:10942)) then
    prepend as-path 38195 2
    pass
  endif
  if (community matches-any (38195:10943)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS13335-12202-IPV6-OUT-39
  if (community matches-any (38195:10949)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS13335-12202-IPV6-OUT-90
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CORE) then
    done
  endif
end-policy
!
route-policy AS13335-12202-IPV6-OUT-91
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    done
  endif
end-policy
!
route-policy AS13335-12202-IPV6-OUT-99
  drop
end-policy
!
route-policy AS15133-12202-IPV4-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS15133-IPV4-PREFIX-DROP-OUT) or (as-path in AS15133-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS15133-12202-IPV4-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS15133-12202-IPV4-OUT-20
  if (large-community matches-any (38195:15133:0)) or (large-community matches-any (38195:15133:68000)) or (large-community matches-any (38195:15133:80100)) then
    drop
  endif
end-policy
!
route-policy AS15133-12202-IPV4-OUT-21
  if (large-community matches-any (38195:15133:1)) or (large-community matches-any (38195:15133:68001)) or (large-community matches-any (38195:15133:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:15133:2)) or (large-community matches-any (38195:15133:68002)) or (large-community matches-any (38195:15133:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:15133:3)) or (large-community matches-any (38195:15133:68003)) or (large-community matches-any (38195:15133:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS15133-12202-IPV4-OUT-29
  if (large-community matches-any (38195:15133:9)) or (large-community matches-any (38195:15133:68009)) or (large-community matches-any (38195:15133:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS15133-12202-IPV4-OUT-90
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CORE) then
    done
  endif
end-policy
!
route-policy AS15133-12202-IPV4-OUT-91
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    done
  endif
end-policy
!
route-policy AS15133-12202-IPV4-OUT-99
  drop
end-policy
!
route-policy AS15133-12202-IPV6-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS15133-IPV6-PREFIX-DROP-OUT) or (as-path in AS15133-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS15133-12202-IPV6-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS15133-12202-IPV6-OUT-20
  if (large-community matches-any (38195:15133:0)) or (large-community matches-any (38195:15133:68000)) or (large-community matches-any (38195:15133:80100)) then
    drop
  endif
end-policy
!
route-policy AS15133-12202-IPV6-OUT-21
  if (large-community matches-any (38195:15133:1)) or (large-community matches-any (38195:15133:68001)) or (large-community matches-any (38195:15133:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:15133:2)) or (large-community matches-any (38195:15133:68002)) or (large-community matches-any (38195:15133:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:15133:3)) or (large-community matches-any (38195:15133:68003)) or (large-community matches-any (38195:15133:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS15133-12202-IPV6-OUT-29
  if (large-community matches-any (38195:15133:9)) or (large-community matches-any (38195:15133:68009)) or (large-community matches-any (38195:15133:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS15133-12202-IPV6-OUT-90
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CORE) then
    done
  endif
end-policy
!
route-policy AS15133-12202-IPV6-OUT-91
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    done
  endif
end-policy
!
route-policy AS15133-12202-IPV6-OUT-99
  drop
end-policy
!
route-policy AS199524-12202-IPV4-IN-10
  if (destination in AS199524-IPV4-PREFIX-DROP-IN) or (as-path in AS199524-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS199524-12202-IPV4-IN-20
  if (destination in AS38195-SLCAU-LE24) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9705, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    set local-preference 152
    done
  else
    drop
  endif
end-policy
!
route-policy AS199524-12202-IPV6-IN-10
  if (destination in AS199524-IPV6-PREFIX-DROP-IN) or (as-path in AS199524-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS199524-12202-IPV6-IN-20
  if (destination in AS38195-SLCAU-LE48) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9705, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    set local-preference 152
    done
  else
    drop
  endif
end-policy
!
route-policy AS20940-12202-IPV4-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS20940-IPV4-PREFIX-DROP-OUT) or (as-path in AS20940-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS20940-12202-IPV4-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS20940-12202-IPV4-OUT-20
  if (large-community matches-any (38195:20940:0)) or (large-community matches-any (38195:20940:68000)) or (large-community matches-any (38195:20940:80100)) then
    drop
  endif
end-policy
!
route-policy AS20940-12202-IPV4-OUT-21
  if (large-community matches-any (38195:20940:1)) or (large-community matches-any (38195:20940:68001)) or (large-community matches-any (38195:20940:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:20940:2)) or (large-community matches-any (38195:20940:68002)) or (large-community matches-any (38195:20940:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:20940:3)) or (large-community matches-any (38195:20940:68003)) or (large-community matches-any (38195:20940:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS20940-12202-IPV4-OUT-29
  if (large-community matches-any (38195:20940:9)) or (large-community matches-any (38195:20940:68009)) or (large-community matches-any (38195:20940:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS20940-12202-IPV4-OUT-90
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CORE) then
    done
  endif
end-policy
!
route-policy AS20940-12202-IPV4-OUT-91
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    done
  endif
end-policy
!
route-policy AS20940-12202-IPV4-OUT-99
  drop
end-policy
!
route-policy AS20940-12202-IPV6-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS20940-IPV6-PREFIX-DROP-OUT) or (as-path in AS20940-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS20940-12202-IPV6-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS20940-12202-IPV6-OUT-20
  if (large-community matches-any (38195:20940:0)) or (large-community matches-any (38195:20940:68000)) or (large-community matches-any (38195:20940:80100)) then
    drop
  endif
end-policy
!
route-policy AS20940-12202-IPV6-OUT-21
  if (large-community matches-any (38195:20940:1)) or (large-community matches-any (38195:20940:68001)) or (large-community matches-any (38195:20940:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:20940:2)) or (large-community matches-any (38195:20940:68002)) or (large-community matches-any (38195:20940:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:20940:3)) or (large-community matches-any (38195:20940:68003)) or (large-community matches-any (38195:20940:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS20940-12202-IPV6-OUT-29
  if (large-community matches-any (38195:20940:9)) or (large-community matches-any (38195:20940:68009)) or (large-community matches-any (38195:20940:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS20940-12202-IPV6-OUT-90
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CORE) then
    done
  endif
end-policy
!
route-policy AS20940-12202-IPV6-OUT-91
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    done
  endif
end-policy
!
route-policy AS20940-12202-IPV6-OUT-99
  drop
end-policy
!
route-policy AS22822-12202-IPV4-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS22822-IPV4-PREFIX-DROP-OUT) or (as-path in AS22822-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS22822-12202-IPV4-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS22822-12202-IPV4-OUT-20
  if (large-community matches-any (38195:22822:0)) or (large-community matches-any (38195:22822:68000)) or (large-community matches-any (38195:22822:80100)) then
    drop
  endif
end-policy
!
route-policy AS22822-12202-IPV4-OUT-21
  if (large-community matches-any (38195:22822:1)) or (large-community matches-any (38195:22822:68001)) or (large-community matches-any (38195:22822:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:22822:2)) or (large-community matches-any (38195:22822:68002)) or (large-community matches-any (38195:22822:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:22822:3)) or (large-community matches-any (38195:22822:68003)) or (large-community matches-any (38195:22822:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS22822-12202-IPV4-OUT-29
  if (large-community matches-any (38195:22822:9)) or (large-community matches-any (38195:22822:68009)) or (large-community matches-any (38195:22822:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS22822-12202-IPV4-OUT-90
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CORE) then
    done
  endif
end-policy
!
route-policy AS22822-12202-IPV4-OUT-91
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    done
  endif
end-policy
!
route-policy AS22822-12202-IPV4-OUT-99
  drop
end-policy
!
route-policy AS22822-12202-IPV6-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS22822-IPV6-PREFIX-DROP-OUT) or (as-path in AS22822-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS22822-12202-IPV6-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS22822-12202-IPV6-OUT-20
  if (large-community matches-any (38195:22822:0)) or (large-community matches-any (38195:22822:68000)) or (large-community matches-any (38195:22822:80100)) then
    drop
  endif
end-policy
!
route-policy AS22822-12202-IPV6-OUT-21
  if (large-community matches-any (38195:22822:1)) or (large-community matches-any (38195:22822:68001)) or (large-community matches-any (38195:22822:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:22822:2)) or (large-community matches-any (38195:22822:68002)) or (large-community matches-any (38195:22822:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:22822:3)) or (large-community matches-any (38195:22822:68003)) or (large-community matches-any (38195:22822:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS22822-12202-IPV6-OUT-29
  if (large-community matches-any (38195:22822:9)) or (large-community matches-any (38195:22822:68009)) or (large-community matches-any (38195:22822:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS22822-12202-IPV6-OUT-90
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CORE) then
    done
  endif
end-policy
!
route-policy AS22822-12202-IPV6-OUT-91
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    done
  endif
end-policy
!
route-policy AS22822-12202-IPV6-OUT-99
  drop
end-policy
!
route-policy AS46489-12202-IPV4-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS46489-IPV4-PREFIX-DROP-OUT) or (as-path in AS46489-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS46489-12202-IPV4-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS46489-12202-IPV4-OUT-20
  if (large-community matches-any (38195:46489:0)) or (large-community matches-any (38195:46489:68000)) or (large-community matches-any (38195:46489:80100)) then
    drop
  endif
end-policy
!
route-policy AS46489-12202-IPV4-OUT-21
  if (large-community matches-any (38195:46489:1)) or (large-community matches-any (38195:46489:68001)) or (large-community matches-any (38195:46489:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:46489:2)) or (large-community matches-any (38195:46489:68002)) or (large-community matches-any (38195:46489:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:46489:3)) or (large-community matches-any (38195:46489:68003)) or (large-community matches-any (38195:46489:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS46489-12202-IPV4-OUT-29
  if (large-community matches-any (38195:46489:9)) or (large-community matches-any (38195:46489:68009)) or (large-community matches-any (38195:46489:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS46489-12202-IPV4-OUT-30
  if (community matches-any (38195:11980)) then
    drop
  endif
end-policy
!
route-policy AS46489-12202-IPV4-OUT-31
  if (community matches-any (38195:11981)) then
    prepend as-path 38195
    pass
  endif
  if (community matches-any (38195:11982)) then
    prepend as-path 38195 2
    pass
  endif
  if (community matches-any (38195:11983)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS46489-12202-IPV4-OUT-39
  if (community matches-any (38195:11989)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS46489-12202-IPV4-OUT-90
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CORE) then
    done
  endif
end-policy
!
route-policy AS46489-12202-IPV4-OUT-91
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    done
  endif
end-policy
!
route-policy AS46489-12202-IPV4-OUT-99
  drop
end-policy
!
route-policy AS46489-12202-IPV6-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS46489-IPV6-PREFIX-DROP-OUT) or (as-path in AS46489-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS46489-12202-IPV6-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS46489-12202-IPV6-OUT-20
  if (large-community matches-any (38195:46489:0)) or (large-community matches-any (38195:46489:68000)) or (large-community matches-any (38195:46489:80100)) then
    drop
  endif
end-policy
!
route-policy AS46489-12202-IPV6-OUT-21
  if (large-community matches-any (38195:46489:1)) or (large-community matches-any (38195:46489:68001)) or (large-community matches-any (38195:46489:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:46489:2)) or (large-community matches-any (38195:46489:68002)) or (large-community matches-any (38195:46489:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:46489:3)) or (large-community matches-any (38195:46489:68003)) or (large-community matches-any (38195:46489:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS46489-12202-IPV6-OUT-29
  if (large-community matches-any (38195:46489:9)) or (large-community matches-any (38195:46489:68009)) or (large-community matches-any (38195:46489:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS46489-12202-IPV6-OUT-30
  if (community matches-any (38195:11980)) then
    drop
  endif
end-policy
!
route-policy AS46489-12202-IPV6-OUT-31
  if (community matches-any (38195:11981)) then
    prepend as-path 38195
    pass
  endif
  if (community matches-any (38195:11982)) then
    prepend as-path 38195 2
    pass
  endif
  if (community matches-any (38195:11983)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS46489-12202-IPV6-OUT-39
  if (community matches-any (38195:11989)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS46489-12202-IPV6-OUT-90
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CORE) then
    done
  endif
end-policy
!
route-policy AS46489-12202-IPV6-OUT-91
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    done
  endif
end-policy
!
route-policy AS46489-12202-IPV6-OUT-99
  drop
end-policy
!
route-policy AS49544-12202-IPV4-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS49544-IPV4-PREFIX-DROP-OUT) or (as-path in AS49544-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS49544-12202-IPV4-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS49544-12202-IPV4-OUT-20
  if (large-community matches-any (38195:49544:0)) or (large-community matches-any (38195:49544:68000)) or (large-community matches-any (38195:49544:80100)) then
    drop
  endif
end-policy
!
route-policy AS49544-12202-IPV4-OUT-21
  if (large-community matches-any (38195:49544:1)) or (large-community matches-any (38195:49544:68001)) or (large-community matches-any (38195:49544:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:49544:2)) or (large-community matches-any (38195:49544:68002)) or (large-community matches-any (38195:49544:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:49544:3)) or (large-community matches-any (38195:49544:68003)) or (large-community matches-any (38195:49544:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS49544-12202-IPV4-OUT-29
  if (large-community matches-any (38195:49544:9)) or (large-community matches-any (38195:49544:68009)) or (large-community matches-any (38195:49544:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS49544-12202-IPV4-OUT-90
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CORE) then
    done
  endif
end-policy
!
route-policy AS49544-12202-IPV4-OUT-91
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    done
  endif
end-policy
!
route-policy AS49544-12202-IPV4-OUT-99
  drop
end-policy
!
route-policy AS49544-12202-IPV6-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS49544-IPV6-PREFIX-DROP-OUT) or (as-path in AS49544-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS49544-12202-IPV6-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS49544-12202-IPV6-OUT-20
  if (large-community matches-any (38195:49544:0)) or (large-community matches-any (38195:49544:68000)) or (large-community matches-any (38195:49544:80100)) then
    drop
  endif
end-policy
!
route-policy AS49544-12202-IPV6-OUT-21
  if (large-community matches-any (38195:49544:1)) or (large-community matches-any (38195:49544:68001)) or (large-community matches-any (38195:49544:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:49544:2)) or (large-community matches-any (38195:49544:68002)) or (large-community matches-any (38195:49544:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:49544:3)) or (large-community matches-any (38195:49544:68003)) or (large-community matches-any (38195:49544:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS49544-12202-IPV6-OUT-29
  if (large-community matches-any (38195:49544:9)) or (large-community matches-any (38195:49544:68009)) or (large-community matches-any (38195:49544:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS49544-12202-IPV6-OUT-90
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CORE) then
    done
  endif
end-policy
!
route-policy AS49544-12202-IPV6-OUT-91
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    done
  endif
end-policy
!
route-policy AS49544-12202-IPV6-OUT-99
  drop
end-policy
!
route-policy AS55256-12202-IPV4-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS55256-IPV4-PREFIX-DROP-OUT) or (as-path in AS55256-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS55256-12202-IPV4-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS55256-12202-IPV4-OUT-20
  if (large-community matches-any (38195:55256:0)) or (large-community matches-any (38195:55256:68000)) or (large-community matches-any (38195:55256:80100)) then
    drop
  endif
end-policy
!
route-policy AS55256-12202-IPV4-OUT-21
  if (large-community matches-any (38195:55256:1)) or (large-community matches-any (38195:55256:68001)) or (large-community matches-any (38195:55256:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:55256:2)) or (large-community matches-any (38195:55256:68002)) or (large-community matches-any (38195:55256:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:55256:3)) or (large-community matches-any (38195:55256:68003)) or (large-community matches-any (38195:55256:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS55256-12202-IPV4-OUT-29
  if (large-community matches-any (38195:55256:9)) or (large-community matches-any (38195:55256:68009)) or (large-community matches-any (38195:55256:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS55256-12202-IPV4-OUT-90
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CORE) then
    done
  endif
end-policy
!
route-policy AS55256-12202-IPV4-OUT-91
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    done
  endif
end-policy
!
route-policy AS55256-12202-IPV4-OUT-99
  drop
end-policy
!
route-policy AS55256-12202-IPV6-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS55256-IPV6-PREFIX-DROP-OUT) or (as-path in AS55256-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS55256-12202-IPV6-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS55256-12202-IPV6-OUT-20
  if (large-community matches-any (38195:55256:0)) or (large-community matches-any (38195:55256:68000)) or (large-community matches-any (38195:55256:80100)) then
    drop
  endif
end-policy
!
route-policy AS55256-12202-IPV6-OUT-21
  if (large-community matches-any (38195:55256:1)) or (large-community matches-any (38195:55256:68001)) or (large-community matches-any (38195:55256:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:55256:2)) or (large-community matches-any (38195:55256:68002)) or (large-community matches-any (38195:55256:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:55256:3)) or (large-community matches-any (38195:55256:68003)) or (large-community matches-any (38195:55256:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS55256-12202-IPV6-OUT-29
  if (large-community matches-any (38195:55256:9)) or (large-community matches-any (38195:55256:68009)) or (large-community matches-any (38195:55256:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS55256-12202-IPV6-OUT-90
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CORE) then
    done
  endif
end-policy
!
route-policy AS55256-12202-IPV6-OUT-91
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    done
  endif
end-policy
!
route-policy AS55256-12202-IPV6-OUT-99
  drop
end-policy
!
route-policy AS63832-12202-IPV4-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS63832-IPV4-PREFIX-DROP-OUT) or (as-path in AS63832-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS63832-12202-IPV4-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS63832-12202-IPV4-OUT-20
  if (large-community matches-any (38195:63832:0)) or (large-community matches-any (38195:63832:68000)) or (large-community matches-any (38195:63832:80100)) then
    drop
  endif
end-policy
!
route-policy AS63832-12202-IPV4-OUT-21
  if (large-community matches-any (38195:63832:1)) or (large-community matches-any (38195:63832:68001)) or (large-community matches-any (38195:63832:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:63832:2)) or (large-community matches-any (38195:63832:68002)) or (large-community matches-any (38195:63832:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:63832:3)) or (large-community matches-any (38195:63832:68003)) or (large-community matches-any (38195:63832:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS63832-12202-IPV4-OUT-29
  if (large-community matches-any (38195:63832:9)) or (large-community matches-any (38195:63832:68009)) or (large-community matches-any (38195:63832:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS63832-12202-IPV4-OUT-30
  if (community matches-any (38195:10100)) or (community matches-any (38195:10120)) then
    drop
  endif
end-policy
!
route-policy AS63832-12202-IPV4-OUT-31
  if (community matches-any (38195:10101)) or (community matches-any (38195:10121)) then
    prepend as-path 38195
    pass
  endif
  if (community matches-any (38195:10102)) or (community matches-any (38195:10122)) then
    prepend as-path 38195 2
    pass
  endif
  if (community matches-any (38195:10103)) or (community matches-any (38195:10123)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS63832-12202-IPV4-OUT-39
  if (community matches-any (38195:10109)) or (community matches-any (38195:10129)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS63832-12202-IPV4-OUT-90
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CORE) then
    set community (65000:6939) additive
    done
  endif
end-policy
!
route-policy AS63832-12202-IPV4-OUT-91
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    set community (65000:6939) additive
    done
  endif
end-policy
!
route-policy AS63832-12202-IPV4-OUT-99
  drop
end-policy
!
route-policy AS63832-12202-IPV6-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS63832-IPV6-PREFIX-DROP-OUT) or (as-path in AS63832-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS63832-12202-IPV6-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS63832-12202-IPV6-OUT-20
  if (large-community matches-any (38195:63832:0)) or (large-community matches-any (38195:63832:68000)) or (large-community matches-any (38195:63832:80100)) then
    drop
  endif
end-policy
!
route-policy AS63832-12202-IPV6-OUT-21
  if (large-community matches-any (38195:63832:1)) or (large-community matches-any (38195:63832:68001)) or (large-community matches-any (38195:63832:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:63832:2)) or (large-community matches-any (38195:63832:68002)) or (large-community matches-any (38195:63832:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:63832:3)) or (large-community matches-any (38195:63832:68003)) or (large-community matches-any (38195:63832:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS63832-12202-IPV6-OUT-29
  if (large-community matches-any (38195:63832:9)) or (large-community matches-any (38195:63832:68009)) or (large-community matches-any (38195:63832:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS63832-12202-IPV6-OUT-30
  if (community matches-any (38195:10100)) or (community matches-any (38195:10120)) then
    drop
  endif
end-policy
!
route-policy AS63832-12202-IPV6-OUT-31
  if (community matches-any (38195:10101)) or (community matches-any (38195:10121)) then
    prepend as-path 38195
    pass
  endif
  if (community matches-any (38195:10102)) or (community matches-any (38195:10122)) then
    prepend as-path 38195 2
    pass
  endif
  if (community matches-any (38195:10103)) or (community matches-any (38195:10123)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS63832-12202-IPV6-OUT-39
  if (community matches-any (38195:10109)) or (community matches-any (38195:10129)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS63832-12202-IPV6-OUT-90
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CORE) then
    set community (65000:6939) additive
    done
  endif
end-policy
!
route-policy AS63832-12202-IPV6-OUT-91
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    set community (65000:6939) additive
    done
  endif
end-policy
!
route-policy AS63832-12202-IPV6-OUT-99
  drop
end-policy
!
route-policy AS199524-12202-IPV4-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS199524-IPV4-PREFIX-DROP-OUT) or (as-path in AS199524-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS199524-12202-IPV4-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS199524-12202-IPV4-OUT-20
  if (large-community matches-any (38195:199524:0)) or (large-community matches-any (38195:199524:68000)) or (large-community matches-any (38195:199524:80100)) then
    drop
  endif
end-policy
!
route-policy AS199524-12202-IPV4-OUT-21
  if (large-community matches-any (38195:199524:1)) or (large-community matches-any (38195:199524:68001)) or (large-community matches-any (38195:199524:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:199524:2)) or (large-community matches-any (38195:199524:68002)) or (large-community matches-any (38195:199524:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:199524:3)) or (large-community matches-any (38195:199524:68003)) or (large-community matches-any (38195:199524:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS199524-12202-IPV4-OUT-29
  if (large-community matches-any (38195:199524:9)) or (large-community matches-any (38195:199524:68009)) or (large-community matches-any (38195:199524:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS199524-12202-IPV4-OUT-90
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CORE) then
    done
  endif
end-policy
!
route-policy AS199524-12202-IPV4-OUT-91
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    done
  endif
end-policy
!
route-policy AS199524-12202-IPV4-OUT-99
  drop
end-policy
!
route-policy AS199524-12202-IPV6-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS199524-IPV6-PREFIX-DROP-OUT) or (as-path in AS199524-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS199524-12202-IPV6-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS199524-12202-IPV6-OUT-20
  if (large-community matches-any (38195:199524:0)) or (large-community matches-any (38195:199524:68000)) or (large-community matches-any (38195:199524:80100)) then
    drop
  endif
end-policy
!
route-policy AS199524-12202-IPV6-OUT-21
  if (large-community matches-any (38195:199524:1)) or (large-community matches-any (38195:199524:68001)) or (large-community matches-any (38195:199524:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:199524:2)) or (large-community matches-any (38195:199524:68002)) or (large-community matches-any (38195:199524:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:199524:3)) or (large-community matches-any (38195:199524:68003)) or (large-community matches-any (38195:199524:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS199524-12202-IPV6-OUT-29
  if (large-community matches-any (38195:199524:9)) or (large-community matches-any (38195:199524:68009)) or (large-community matches-any (38195:199524:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS199524-12202-IPV6-OUT-90
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CORE) then
    done
  endif
end-policy
!
route-policy AS199524-12202-IPV6-OUT-91
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    done
  endif
end-policy
!
route-policy AS199524-12202-IPV6-OUT-99
  drop
end-policy
!


prefix-set AS174-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS714-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS714-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS984-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS984-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS174-IPV4-PREFIX-DROP-OUT
  45.67.96.0/24
end-set
prefix-set AS174-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS1828-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS1828-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS2603-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS2603-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS2635-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS2635-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS2687-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS2687-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS3223-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS3223-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS3300-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS3300-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS3356-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS3491-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32,
  # Ticket 56911 - Mammoth Packet Loss
  43.228.96.0/24
end-set
prefix-set AS3491-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS38195-SLCAU-DEFAULT-ONLY
  0.0.0.0/0 eq 0,
  ::/0 eq 0
end-set
prefix-set AS4657-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS4657-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS4788-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32,
  175.139.192.0/18
end-set
prefix-set AS4788-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS6447-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS6447-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS6507-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS6507-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS6939-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS6939-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS714-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS714-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS7387-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS7387-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS7473-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS7473-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS8075-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS8075-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS8220-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS8220-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS8529-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS8529-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS8674-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS8674-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS8966-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS8966-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS9009-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS9009-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS9498-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS9498-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS984-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS984-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS10075-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS10075-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS10310-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS10310-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS13150-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS13335-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS13335-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS13414-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS13414-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS13445-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS13445-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS14061-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS14061-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS14907-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS14907-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS15133-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS15133-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS15169-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS15169-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS15695-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS15695-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS16276-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS16276-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS16509-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS16509-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS17639-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS17639-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS18106-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS18106-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS1828-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS1828-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS18403-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS18403-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS19551-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS19551-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS19679-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS19679-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS20473-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS20473-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS20940-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  100.64.0.0/10 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS20940-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS21859-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS21859-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS22697-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS22697-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS22822-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  100.64.0.0/10 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS22822-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS23764-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS23764-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS24115-IPV4-PREFIX-DROP-IN
  # OOB Transit via EquinixIX
  27.111.208.0/20,
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS24115-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS24224-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS24224-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS2603-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS2603-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS2635-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS2635-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS26415-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS26415-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS2687-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS2687-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS3223-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS3223-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS32590-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS32590-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS32787-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS32787-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS32934-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS32934-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS3300-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS3300-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS33438-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS33438-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS3356-IPV4-PREFIX-DROP-OUT
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS3356-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS34309-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS34309-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS3491-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS3491-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS36351-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS36351-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS36692-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS36692-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS37468-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS37468-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS45102-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS45489-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS45489-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS46489-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS46489-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS4657-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS4657-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS4788-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS4788-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS49544-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS49544-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS54113-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS54113-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS54994-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS54994-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS55256-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS55256-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS55518-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32,
  175.139.192.0/18
end-set
prefix-set AS55518-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS57976-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS57976-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS62955-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS62955-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS63832-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS63832-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS63949-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS63949-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS64096-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS64096-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS6447-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS6447-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS6507-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS6507-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS6939-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS6939-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS7387-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS7387-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS7473-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS7473-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS8075-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS8075-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS8220-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS8220-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS8529-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS8529-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS8674-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS8674-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS8966-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS8966-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS9009-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS9009-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS9498-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS9498-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS10075-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS10075-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS10310-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS10310-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS13150-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS132203-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS132203-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS13335-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS13335-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS13414-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS13414-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS13445-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS13445-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS136907-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS136907-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS137409-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS137409-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS139225-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS139225-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS139341-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS139341-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS14061-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS14061-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS14907-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS14907-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS151326-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS151326-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS15133-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS15133-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS15169-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS15169-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS15695-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS15695-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS16276-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS16276-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS16509-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS16509-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS17639-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS17639-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS18106-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS18106-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS18403-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS18403-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS19551-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS19551-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS19679-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS19679-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS199524-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS199524-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS20473-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS20473-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS20940-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS20940-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS21859-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS21859-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS22697-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS22697-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS22822-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS22822-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS23764-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS23764-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS24115-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS24115-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS24224-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS24224-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS26415-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS26415-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS32590-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS32590-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS32787-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS32787-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS32934-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS32934-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS33438-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS33438-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS34309-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS34309-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS36351-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS36351-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS36692-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS36692-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS37468-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS37468-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS396986-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS396986-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS396998-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS396998-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS397601-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
prefix-set AS397601-IPV6-PREFIX-DROP-IN
end-set
prefix-set AS45102-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS45489-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS45489-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS46489-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS46489-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS49544-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS49544-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS54113-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS54113-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS54994-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS54994-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS55256-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS55256-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS55518-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS55518-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS57976-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS57976-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS62955-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS62955-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS63832-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS63832-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS63949-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS63949-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS64096-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS64096-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS132203-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS132203-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS136907-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS136907-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS137409-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS137409-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS139225-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS139225-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS139341-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS139341-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS151326-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS151326-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS199524-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS199524-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS396986-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS396986-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS396998-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS396998-IPV6-PREFIX-DROP-OUT
end-set
prefix-set AS397601-IPV4-PREFIX-DROP-OUT
end-set
prefix-set AS397601-IPV6-PREFIX-DROP-OUT
end-set


as-path-set AS3356-DEPREF-IN
end-set
!
as-path-set AS7473-DEPREF-IN
end-set
!
as-path-set AS174-ASPATH-DROP-IN
  ios-regex '_1221_',
  ios-regex '_1851_',
  ios-regex '_4826_',
  ios-regex '_7545_',
  ios-regex '_7575_',
  ios-regex '_15133$',
  ios-regex '_17819_',
  ios-regex '_27435$',
  ios-regex '_32934_',
  ios-regex '_3257_',
  # Reject china routes due to latency
  ios-regex '_58453_',
  ios-regex '_4134_',
  ios-regex '_4837_'
end-set
!
as-path-set AS714-ASPATH-DROP-IN
end-set
!
as-path-set AS984-ASPATH-DROP-IN
end-set
!
as-path-set AS174-ASPATH-DROP-OUT
end-set
!
as-path-set AS1828-ASPATH-DROP-IN
end-set
!
as-path-set AS2603-ASPATH-DROP-IN
end-set
!
as-path-set AS2635-ASPATH-DROP-IN
end-set
!
as-path-set AS2687-ASPATH-DROP-IN
end-set
!
as-path-set AS3223-ASPATH-DROP-IN
end-set
!
as-path-set AS3300-ASPATH-DROP-IN
end-set
!
as-path-set AS3356-ASPATH-DROP-IN
  #Australian ASes
  ios-regex '_1221_',
  ios-regex '_1851_',
  ios-regex '_4826_',
  ios-regex '_7545_',
  ios-regex '_7575_',
  ios-regex '_3491_',
  # Lumen US
  ios-regex '_202_'
end-set
!
as-path-set AS3491-ASPATH-DROP-IN
  ios-regex '_174_',
  ios-regex '_1221_',
  ios-regex '_1851_',
  ios-regex '_4826_',
  ios-regex '_7545_',
  ios-regex '_7575_',
  ios-regex '_15133$',
  ios-regex '_17819_',
  ios-regex '_27435$',
  ios-regex '_32934_',
  ios-regex '_7474_',
  ios-regex '_4804_',
  # Internet Initiative Japan Inc via PCCW SG
  ios-regex '_2497_',
  # defence.net
  ios-regex '_55002_',
  # Drop Gameloft via PCCW - CHG00945
  ios-regex '_11807_',
  # CHG01293 - Optimise china routes
  ios-regex '_58453_',
  # Lumen US
  ios-regex '_202_'
end-set
!
as-path-set AS4657-ASPATH-DROP-IN
end-set
!
as-path-set AS4788-ASPATH-DROP-IN
end-set
!
as-path-set AS6447-ASPATH-DROP-IN
end-set
!
as-path-set AS6507-ASPATH-DROP-IN
end-set
!
as-path-set AS6939-ASPATH-DROP-IN
  # Block Equinix via HE
  ios-regex '_9989_',
  #starhub via HE
  originates-from '4657',
  #omaha mutual block - 20190509 tw
  originates-from '17094',
  #Mythic-Beasts (UK) - 20190821 tw
  originates-from '44684',
  # Shinjiru - Known Malaysian ASN
  originates-from '45839',
  # Abavia, NL
  originates-from '48345',
  # TPG via HE - SL132758
  ios-regex '_7545_',
  ios-regex '_1221_',
  ios-regex '_2764_',
  ios-regex '_7474_',
  ios-regex '_10143_',
  ios-regex '_9280_',
  ios-regex '_45671_'
end-set
!
as-path-set AS714-ASPATH-DROP-OUT
end-set
!
as-path-set AS7387-ASPATH-DROP-IN
end-set
!
as-path-set AS7473-ASPATH-DROP-IN
end-set
!
as-path-set AS8075-ASPATH-DROP-IN
end-set
!
as-path-set AS8220-ASPATH-DROP-IN
end-set
!
as-path-set AS8529-ASPATH-DROP-IN
end-set
!
as-path-set AS8674-ASPATH-DROP-IN
end-set
!
as-path-set AS8966-ASPATH-DROP-IN
end-set
!
as-path-set AS9009-ASPATH-DROP-IN
end-set
!
as-path-set AS9498-ASPATH-DROP-IN
  # Drop Deluxe.in via Bharti
  ios-regex '_139911_'
end-set
!
as-path-set AS984-ASPATH-DROP-OUT
end-set
!
as-path-set AS10075-ASPATH-DROP-IN
end-set
!
as-path-set AS10310-ASPATH-DROP-IN
end-set
!
as-path-set AS13150-ASPATH-DROP-IN
end-set
!
as-path-set AS13335-ASPATH-DROP-IN
end-set
!
as-path-set AS13414-ASPATH-DROP-IN
end-set
!
as-path-set AS13445-ASPATH-DROP-IN
end-set
!
as-path-set AS14061-ASPATH-DROP-IN
end-set
!
as-path-set AS14907-ASPATH-DROP-IN
end-set
!
as-path-set AS15133-ASPATH-DROP-IN
end-set
!
as-path-set AS15169-ASPATH-DROP-IN
end-set
!
as-path-set AS15695-ASPATH-DROP-IN
end-set
!
as-path-set AS16276-ASPATH-DROP-IN
end-set
!
as-path-set AS16509-ASPATH-DROP-IN
end-set
!
as-path-set AS17639-ASPATH-DROP-IN
end-set
!
as-path-set AS18106-ASPATH-DROP-IN
end-set
!
as-path-set AS1828-ASPATH-DROP-OUT
end-set
!
as-path-set AS18403-ASPATH-DROP-IN
end-set
!
as-path-set AS19551-ASPATH-DROP-IN
end-set
!
as-path-set AS19679-ASPATH-DROP-IN
end-set
!
as-path-set AS20473-ASPATH-DROP-IN
end-set
!
as-path-set AS20940-ASPATH-DROP-IN
end-set
!
as-path-set AS21859-ASPATH-DROP-IN
end-set
!
as-path-set AS22697-ASPATH-DROP-IN
end-set
!
as-path-set AS22822-ASPATH-DROP-IN
end-set
!
as-path-set AS23764-ASPATH-DROP-IN
end-set
!
as-path-set AS24115-ASPATH-DROP-IN
  ios-regex '_1851_',
  ios-regex '_2687_',
  ios-regex '_3549_',
  ios-regex '_3908_',
  ios-regex '_4826_',
  ios-regex '_7545_',
  ios-regex '_7575_',
  ios-regex '_7632_',
  ios-regex '_32097_',
  ios-regex '_39386_',
  ios-regex '_45458_',
  #Blocking ASNs due to bilat
  ios-regex '_6939_',
  ios-regex '_8403_',
  ios-regex '_13335_',
  ios-regex '_16509_',
  ios-regex '_16276_',
  ios-regex '_32934_',
  ios-regex '_137409_'
end-set
!
as-path-set AS24224-ASPATH-DROP-IN
  ios-regex '_1221_',
  ios-regex '_7474_',
  ios-regex '_2687_',
  ios-regex '_3549_',
  ios-regex '_4826_',
  ios-regex '_7545_',
  ios-regex '_7575_',
  ios-regex '_7632_',
  ios-regex '_32097_',
  ios-regex '_39386_',
  ios-regex '_45458_',
  ios-regex '_4764_',
  #Blocking ASNs due to bilat
  ios-regex '_6939_',
  ios-regex '_8403_',
  ios-regex '_13335_',
  ios-regex '_16509_',
  ios-regex '_16276_',
  ios-regex '_32934_',
  ios-regex '_137409_'
end-set
!
as-path-set AS2603-ASPATH-DROP-OUT
end-set
!
as-path-set AS2635-ASPATH-DROP-OUT
end-set
!
as-path-set AS26415-ASPATH-DROP-IN
end-set
!
as-path-set AS2687-ASPATH-DROP-OUT
end-set
!
as-path-set AS3223-ASPATH-DROP-OUT
end-set
!
as-path-set AS32590-ASPATH-DROP-IN
end-set
!
as-path-set AS32787-ASPATH-DROP-IN
end-set
!
as-path-set AS32934-ASPATH-DROP-IN
end-set
!
as-path-set AS3300-ASPATH-DROP-OUT
end-set
!
as-path-set AS33438-ASPATH-DROP-IN
end-set
!
as-path-set AS3356-ASPATH-DROP-OUT
  ios-regex '_4749_'
end-set
!
as-path-set AS34309-ASPATH-DROP-IN
end-set
!
as-path-set AS3491-ASPATH-DROP-OUT
end-set
!
as-path-set AS3491-DOMESTIC-DEPREF
  # Cogent via PCCW
  ios-regex '_174_',
  # Vodafone Global
  ios-regex '_1273_',
  # Telstra Global via PCCW
  ios-regex '_4637_',
  # HE.net via PCCW
  ios-regex '_6939_',
  # Depreference EA via PCCW
  ios-regex '_19541_'
end-set
!
as-path-set AS36351-ASPATH-DROP-IN
end-set
!
as-path-set AS36692-ASPATH-DROP-IN
end-set
!
as-path-set AS37468-ASPATH-DROP-IN
end-set
!
as-path-set AS45102-ASPATH-DROP-IN
end-set
!
as-path-set AS45489-ASPATH-DROP-IN
end-set
!
as-path-set AS46489-ASPATH-DROP-IN
end-set
!
as-path-set AS4657-ASPATH-DROP-OUT
end-set
!
as-path-set AS4788-ASPATH-DROP-OUT
end-set
!
as-path-set AS49544-ASPATH-DROP-IN
end-set
!
as-path-set AS54113-ASPATH-DROP-IN
end-set
!
as-path-set AS54994-ASPATH-DROP-IN
end-set
!
as-path-set AS55256-ASPATH-DROP-IN
end-set
!
as-path-set AS55518-ASPATH-DROP-IN
  ios-regex '_1221_',
  ios-regex '_1851_',
  ios-regex '_2687_',
  ios-regex '_3549_',
  ios-regex '_3908_',
  ios-regex '_4826_',
  ios-regex '_7545_',
  ios-regex '_7575_',
  ios-regex '_7632_',
  ios-regex '_32097_',
  ios-regex '_45430_',
  ios-regex '_45458_',
  #Blocking ASNs due to bilat
  ios-regex '_6939_',
  ios-regex '_8403_',
  ios-regex '_13335_',
  ios-regex '_16509_',
  ios-regex '_16276_',
  ios-regex '_32934_',
  ios-regex '_137409_'
end-set
!
as-path-set AS57976-ASPATH-DROP-IN
end-set
!
as-path-set AS62955-ASPATH-DROP-IN
end-set
!
as-path-set AS63832-ASPATH-DROP-IN
  #Blocking ASNs due to bilat
  ios-regex '_6939_',
  ios-regex '_8403_',
  ios-regex '_13335_',
  ios-regex '_16509_',
  ios-regex '_16276_',
  ios-regex '_32590_',
  ios-regex '_32934_',
  ios-regex '_36351_',
  ios-regex '_38622_',
  ios-regex '_57976_',
  ios-regex '_137409_',
  #Blocking KordaMentha ASN due to fault escalation
  ios-regex '_135526_'
end-set
!
as-path-set AS63949-ASPATH-DROP-IN
end-set
!
as-path-set AS64096-ASPATH-DROP-IN
end-set
!
as-path-set AS6447-ASPATH-DROP-OUT
end-set
!
as-path-set AS6507-ASPATH-DROP-OUT
end-set
!
as-path-set AS6939-ASPATH-DROP-OUT
end-set
!
as-path-set AS7387-ASPATH-DROP-OUT
end-set
!
as-path-set AS7473-ASPATH-DROP-OUT
  ios-regex '_54994_'
end-set
!
as-path-set AS7473-DOMESTIC-DEPREF
end-set
!
as-path-set AS8075-ASPATH-DROP-OUT
end-set
!
as-path-set AS8220-ASPATH-DROP-OUT
end-set
!
as-path-set AS8529-ASPATH-DROP-OUT
end-set
!
as-path-set AS8674-ASPATH-DROP-OUT
end-set
!
as-path-set AS8966-ASPATH-DROP-OUT
end-set
!
as-path-set AS9009-ASPATH-DROP-OUT
end-set
!
as-path-set AS9498-ASPATH-DROP-OUT
end-set
!
as-path-set AS10075-ASPATH-DROP-OUT
end-set
!
as-path-set AS10310-ASPATH-DROP-OUT
end-set
!
as-path-set AS13150-ASPATH-DROP-OUT
end-set
!
as-path-set AS132203-ASPATH-DROP-IN
end-set
!
as-path-set AS13335-ASPATH-DROP-OUT
end-set
!
as-path-set AS13414-ASPATH-DROP-OUT
end-set
!
as-path-set AS13445-ASPATH-DROP-OUT
end-set
!
as-path-set AS136907-ASPATH-DROP-IN
end-set
!
as-path-set AS137409-ASPATH-DROP-IN
end-set
!
as-path-set AS139225-ASPATH-DROP-IN
end-set
!
as-path-set AS139341-ASPATH-DROP-IN
end-set
!
as-path-set AS14061-ASPATH-DROP-OUT
end-set
!
as-path-set AS14907-ASPATH-DROP-OUT
end-set
!
as-path-set AS151326-ASPATH-DROP-IN
end-set
!
as-path-set AS15133-ASPATH-DROP-OUT
end-set
!
as-path-set AS15169-ASPATH-DROP-OUT
end-set
!
as-path-set AS15695-ASPATH-DROP-OUT
end-set
!
as-path-set AS16276-ASPATH-DROP-OUT
end-set
!
as-path-set AS16509-ASPATH-DROP-OUT
end-set
!
as-path-set AS17639-ASPATH-DROP-OUT
end-set
!
as-path-set AS18106-ASPATH-DROP-OUT
end-set
!
as-path-set AS18403-ASPATH-DROP-OUT
end-set
!
as-path-set AS19551-ASPATH-DROP-OUT
end-set
!
as-path-set AS19679-ASPATH-DROP-OUT
end-set
!
as-path-set AS199524-ASPATH-DROP-IN
end-set
!
as-path-set AS20473-ASPATH-DROP-OUT
end-set
!
as-path-set AS20940-ASPATH-DROP-OUT
end-set
!
as-path-set AS21859-ASPATH-DROP-OUT
end-set
!
as-path-set AS22697-ASPATH-DROP-OUT
end-set
!
as-path-set AS22822-ASPATH-DROP-OUT
end-set
!
as-path-set AS23764-ASPATH-DROP-OUT
end-set
!
as-path-set AS24115-ASPATH-DROP-OUT
  ios-regex '_9661_'
end-set
!
as-path-set AS24224-ASPATH-DROP-OUT
end-set
!
as-path-set AS26415-ASPATH-DROP-OUT
end-set
!
as-path-set AS32590-ASPATH-DROP-OUT
end-set
!
as-path-set AS32787-ASPATH-DROP-OUT
end-set
!
as-path-set AS32934-ASPATH-DROP-OUT
end-set
!
as-path-set AS33438-ASPATH-DROP-OUT
end-set
!
as-path-set AS34309-ASPATH-DROP-OUT
end-set
!
as-path-set AS36351-ASPATH-DROP-OUT
end-set
!
as-path-set AS36692-ASPATH-DROP-OUT
end-set
!
as-path-set AS37468-ASPATH-DROP-OUT
end-set
!
as-path-set AS396986-ASPATH-DROP-IN
end-set
!
as-path-set AS396998-ASPATH-DROP-IN
end-set
!
as-path-set AS397601-ASPATH-DROP-IN
end-set
!
as-path-set AS45102-ASPATH-DROP-OUT
end-set
!
as-path-set AS45489-ASPATH-DROP-OUT
end-set
!
as-path-set AS46489-ASPATH-DROP-OUT
end-set
!
as-path-set AS49544-ASPATH-DROP-OUT
end-set
!
as-path-set AS54113-ASPATH-DROP-OUT
end-set
!
as-path-set AS54994-ASPATH-DROP-OUT
end-set
!
as-path-set AS55256-ASPATH-DROP-OUT
end-set
!
as-path-set AS55518-ASPATH-DROP-OUT
  ios-regex '_9661_'
end-set
!
as-path-set AS57976-ASPATH-DROP-OUT
end-set
!
as-path-set AS62955-ASPATH-DROP-OUT
end-set
!
as-path-set AS63832-ASPATH-DROP-OUT
end-set
!
as-path-set AS63949-ASPATH-DROP-OUT
end-set
!
as-path-set AS64096-ASPATH-DROP-OUT
end-set
!
as-path-set AS132203-ASPATH-DROP-OUT
end-set
!
as-path-set AS136907-ASPATH-DROP-OUT
end-set
!
as-path-set AS137409-ASPATH-DROP-OUT
end-set
!
as-path-set AS139225-ASPATH-DROP-OUT
end-set
!
as-path-set AS139341-ASPATH-DROP-OUT
end-set
!
as-path-set AS151326-ASPATH-DROP-OUT
end-set
!
as-path-set AS199524-ASPATH-DROP-OUT
end-set
!
as-path-set AS396986-ASPATH-DROP-OUT
end-set
!
as-path-set AS396998-ASPATH-DROP-OUT
end-set
!
as-path-set AS397601-ASPATH-DROP-OUT
end-set
!
as-path-set AS7473-CHINA-NO-PREPENDS
  ios-regex '_151964_',
  ios-regex '_38008_'
end-set
!
as-path-set AS3356-DOMESTIC-DEPREF-IN
end-set
!
as-path-set AS174-30280-ASPATH-DEPREF-OUT
  ios-regex '_10143_'
end-set
!
as-path-set AS174-ASPATH-LOCALPREF-131-IN
end-set
!
as-path-set AS3356-ASPATH-LOCALPREF-131-IN
end-set
!
as-path-set AS3356-ASPATH-LOCALPREF-143-IN
end-set
!
as-path-set AS3491-ASPATH-LOCALPREF-129-IN
  # Depreference Jagex via PCCW/GTT
  ios-regex '_3257_44521_',
  # Depreference EA via PCCW
  ios-regex '_19541_',
  # path.net via USA - tw 20200308
  originates-from '396998',
  # Lumen USA
  ios-regex '_3356_202_'
end-set
!
as-path-set AS3491-ASPATH-LOCALPREF-131-IN
end-set
!
as-path-set AS3491-ASPATH-LOCALPREF-143-IN
end-set
!
as-path-set AS7473-ASPATH-LOCALPREF-143-IN
end-set
!


route-policy AS63832-12202-IPV4-IN-20
  if (destination in AS38195-SLCAU-LE24) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9706, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    set local-preference 152
    done
  else
    drop
  endif
end-policy
!
route-policy AS714-12202-IPV6-IN-20
  if (destination in AS38195-SLCAU-LE48) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9705, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    set local-preference 152
    done
  else
    drop
  endif
end-policy
!

route-policy AS714-12202-IPV4-IN-20
  if (destination in AS38195-SLCAU-LE24) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    set community (38195:9705, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    set local-preference 152
    done
  else
    drop
  endif
end-policy
