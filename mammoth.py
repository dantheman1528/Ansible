#Mammoth parallel build:


bdr02-etp-26aayer:

    interface Te0/0/0/1 description Cust: Mammoth Media Pty Ltd - IP-PT [10000Mbit]{545022}
    interface Te0/0/0/1 mtu 9216
    interface Te0/0/0/1 load-interval 30
    interface Te0/0/0/1 l2transport


    l2vpn xconnect group xc_545022
    l2vpn xconnect group xc_545022 p2p p2p_545022
    l2vpn xconnect group xc_545022 p2p p2p_545022 interface Te0/0/0/1
    l2vpn xconnect group xc_545022 p2p p2p_545022 neighbor ipv4 202.177.40.3 pw-id 545022
    l2vpn xconnect group xc_545022 p2p p2p_545022 neighbor ipv4 202.177.40.3 pw-id 545022 pw-class CW-enabled


bdr01-eip-26aayer:

    l2vpn xconnect group xc_545022
    l2vpn xconnect group xc_545022 p2p p2p_545022
    l2vpn xconnect group xc_545022 p2p p2p_545022 interface HundredGigE0/0/0/22.545022
    l2vpn xconnect group xc_545022 p2p p2p_545022 neighbor ipv4 172.21.55.14 pw-id 545022
    l2vpn xconnect group xc_545022 p2p p2p_545022 neighbor ipv4 172.21.55.14 pw-id 545022 pw-class CW-enabled

    interface HundredGigE0/0/0/22.545022 l2transport
    interface HundredGigE0/0/0/22.545022 l2transport description Loop-L2: Mammoth Media Pty Ltd - IP-PT (AS39855)[5000Mbps]{545022}
    interface HundredGigE0/0/0/22.545022 l2transport encapsulation dot1q 101
    interface HundredGigE0/0/0/22.545022 l2transport rewrite ingress tag pop 1 symmetric
    interface HundredGigE0/0/0/22.545022 l2transport mtu 9216


    interface HundredGigE0/0/0/23.103
    interface HundredGigE0/0/0/23.103 description Cust: Mammoth Media Pty - IP-PT (AS39855)[5000Mbps]{545022}
    interface HundredGigE0/0/0/23.103 service-policy input POLICE-5000MBPS
    interface HundredGigE0/0/0/23.103 ipv4 mtu 1500
    interface HundredGigE0/0/0/23.103 ipv4 address 202.177.42.8 255.255.255.254
    interface HundredGigE0/0/0/23.103 ipv4 unreachables disable
    interface HundredGigE0/0/0/23.103 load-interval 30
    interface HundredGigE0/0/0/23.103 flow ipv4 monitor netflow-ipv4 sampler 1in1k ingress
    interface HundredGigE0/0/0/23.103 ipv6 nd suppress-ra
    interface HundredGigE0/0/0/23.103 ipv6 mtu 1500
    interface HundredGigE0/0/0/23.103 ipv6 address 2401:d001:1101::81/124
    interface HundredGigE0/0/0/23.103 ipv6 unreachables disable
    interface HundredGigE0/0/0/23.103 load-interval 30
    interface HundredGigE0/0/0/23.103 flow ipv6 monitor netflow-ipv6 sampler 1in1k ingress
    interface HundredGigE0/0/0/23.103 encapsulation dot1q 101



    router bgp 64590 neighbor 202.177.42.9 local-as 38195 no-prepend replace-as
    router bgp 64590 neighbor 202.177.42.9 remote-as 133159
    router bgp 64590 neighbor 202.177.42.9 timers 10 180
    router bgp 64590 neighbor 202.177.42.9 description Cust:Mammoth Media Pty Ltd-545022
    router bgp 64590 neighbor 202.177.42.9 address-family ipv4 unicast
    router bgp 64590 neighbor 202.177.42.9 address-family ipv4 unicast maximum-prefix 100 90 restart 15
    router bgp 64590 neighbor 202.177.42.9 address-family ipv4 unicast send-community-ebgp
    router bgp 64590 neighbor 202.177.42.9 address-family ipv4 unicast route-policy SLCAU-CUST-IPV4-IN(AS133159-545022-IPV4-IN) in
    router bgp 64590 neighbor 202.177.42.9 address-family ipv4 unicast route-policy SLCAU-CUST-IPV4-FULLPLUSDEFAULT-OUT out
    router bgp 64590 neighbor 202.177.42.9 address-family ipv4 unicast default-originate
    router bgp 64590 neighbor 202.177.42.9 address-family ipv4 unicast remove-private-AS
    router bgp 64590 neighbor 202.177.42.9 address-family ipv4 unicast soft-reconfiguration inbound always
    router bgp 64590 neighbor 2401:d001:1101::82 local-as 38195 no-prepend replace-as
    router bgp 64590 neighbor 2401:d001:1101::82 remote-as 133159
    router bgp 64590 neighbor 2401:d001:1101::82 timers 10 180
    router bgp 64590 neighbor 2401:d001:1101::82 description Cust:Mammoth Media Pty Ltd-545022
    router bgp 64590 neighbor 2401:d001:1101::82 address-family ipv6 unicast
    router bgp 64590 neighbor 2401:d001:1101::82 address-family ipv6 unicast maximum-prefix 100 90 restart 15
    router bgp 64590 neighbor 2401:d001:1101::82 address-family ipv6 unicast send-community-ebgp
    router bgp 64590 neighbor 2401:d001:1101::82 address-family ipv6 unicast route-policy SLCAU-CUST-IPV6-IN(AS133159-545022-IPV6-IN) in
    router bgp 64590 neighbor 2401:d001:1101::82 address-family ipv6 unicast route-policy SLCAU-CUST-IPV6-FULLPLUSDEFAULT-OUT out
    router bgp 64590 neighbor 2401:d001:1101::82 address-family ipv6 unicast default-originate
    router bgp 64590 neighbor 2401:d001:1101::82 address-family ipv6 unicast remove-private-AS
    router bgp 64590 neighbor 2401:d001:1101::82 address-family ipv6 unicast soft-reconfiguration inbound always


    prefix-set AS133159-545022-IPV4-IN
    # Company: Mammoth Media
    # Service ID: 545022
    # IRT: https://apps-networks.au.superloop.systems/irt/company/381/
    103.1.184.0/22 le 32,
    103.100.36.0/22 le 32,
    103.153.18.0/23 le 32,
    103.16.128.0/22 le 32,
    103.17.56.0/23 le 32,
    103.190.129.0/24 le 32,
    103.225.167.0/24 le 32,
    103.230.156.0/22 le 32,
    103.236.162.0/23 le 32,
    103.249.236.0/22 le 32,
    103.4.234.0/23 le 32,
    110.232.112.0/22 le 32,
    112.213.32.0/21 le 32,
    119.42.52.0/22 le 32,
    150.107.72.0/22 le 32,
    175.45.180.0/22 le 32,
    203.18.245.0/24 le 32,
    203.18.30.0/24 le 32,
    203.25.132.0/24 le 32,
    203.25.238.0/24 le 32,
    203.29.240.0/22 le 32,
    203.56.35.0/24 le 32,
    203.57.114.0/23 le 32,
    203.57.50.0/23 le 32,
    43.224.180.0/22 le 32,
    43.229.60.0/22 le 32,
    44.31.48.0/24 le 32,
    45.124.52.0/22 le 32
    end-set
    !
    prefix-set AS133159-545022-IPV6-IN
    # Company: Mammoth Media
    # Service ID: 545022
    # IRT: https://apps-networks.au.superloop.systems/irt/company/381/
    2404:9400::/32 le 128
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


RP/0/RP0/CPU0:bdr01-ipt-26aayerr-sin.sg#sh bgp neighbor 202.177.42.51 routes
Tue Mar 25 03:11:47.285 UTC
BGP router identifier 202.177.40.2, local AS number 38195
BGP generic scan interval 60 secs
Non-stop routing is enabled
BGP table state: Active
Table ID: 0xe0000000   RD version: 3673682009
BGP main routing table version 3673682009
BGP NSR Initial initsync version 682689 (Reached)
BGP NSR/ISSU Sync-Group versions 0/0
BGP scan interval 60 secs

Status codes: s suppressed, d damped, h history, * valid, > best
              i - internal, r RIB-failure, S stale, N Nexthop-discard
Origin codes: i - IGP, e - EGP, ? - incomplete
   Network            Next Hop            Metric LocPrf Weight Path
*> 103.17.56.0/24     202.177.42.51            0    194      0 133159 i
*> 103.100.36.0/22    202.177.42.51            0    194      0 133159 i
*> 103.100.36.0/23    202.177.42.51            0    194      0 133159 i

Processed 3 prefixes, 3 paths
RP/0/RP0/CPU0:bdr01-ipt-26aayerr-sin.sg#sh bgp neighbor 202.177.42.51 advertised-count
Tue Mar 25 03:11:56.013 UTC
No of prefixes Advertised: 977386





RP/0/RP0/CPU0:bdr01-ipt-26aayerr-sin.sg#sh bgp ipv6 unicast neighbors 2401:d001:1101::22 routes
Tue Mar 25 03:12:11.849 UTC
BGP router identifier 202.177.40.2, local AS number 38195
BGP generic scan interval 60 secs
Non-stop routing is enabled
BGP table state: Active
Table ID: 0xe0800000   RD version: 3116966102
BGP main routing table version 3116966102
BGP NSR Initial initsync version 177811 (Reached)
BGP NSR/ISSU Sync-Group versions 0/0
BGP scan interval 60 secs

Status codes: s suppressed, d damped, h history, * valid, > best
              i - internal, r RIB-failure, S stale, N Nexthop-discard
Origin codes: i - IGP, e - EGP, ? - incomplete
   Network            Next Hop            Metric LocPrf Weight Path
*> 2404:9400:5::/48   2401:d001:1101::22
                                               0    194      0 133159 i
*> 2404:9400:f::/48   2401:d001:1101::22
                                               0    194      0 133159 i
*> 2404:9400:5000::/36
                      2401:d001:1101::22
                                               0    194      0 133159 i

Processed 3 prefixes, 3 paths
RP/0/RP0/CPU0:bdr01-ipt-26aayerr-sin.sg#sh bgp ipv6 unicast neighbors 2401:d001:1101::22 advertised-count
Tue Mar 25 03:12:19.439 UTC
No of prefixes Advertised: 210445
