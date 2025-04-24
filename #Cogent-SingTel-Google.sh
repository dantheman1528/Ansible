#Cogent SG

bdr01-eip-15pioneer

prefix-set AS174-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32,
  85.184.96.0/20 le 32
end-set
!
prefix-set AS38195-SLCAU-DEFAULT-ONLY
  0.0.0.0/0 eq 0,
  ::/0 eq 0
end-set
!
prefix-set AS38195-SLCAU-LE24
  0.0.0.0/0 le 24
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
as-path-set AS174-ASPATH-LOCALPREF-131-IN
end-set
!
community-set AS38195-SLCAU-EXTERNAL-DELETE
  38195:*
end-set
!
route-policy AS174-30280-IPV4-IN-10
  if (destination in AS174-IPV4-PREFIX-DROP-IN) or (as-path in AS174-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS174-30280-IPV4-IN-20
  if (community matches-any (174:21201)) then
    set community (38195:9708, 38195:7350, 38195:7662, 38195:8010, 38195:6800) additive
    set med 0
    if (community matches-any (174:22038)) then
      set local-preference 142
      done
    endif
  endif
end-policy
!
route-policy AS174-30280-IPV4-IN-30
  set community (38195:9709, 38195:7350, 38195:7662, 38195:8010, 38195:6800) additive
  if (as-path in AS174-ASPATH-LOCALPREF-131-IN) then
    set local-preference 131
    done
  endif
  if (community matches-any (174:21201, 174:22005, 174:21100, 174:21101)) and (not community matches-any (174:22032)) then
    set local-preference 130
    done
  else
    set local-preference 110
    done
  endif
end-policy
!
route-policy AS174-30280-IPV4-IN
  if (destination in AS38195-SLCAU-LE24) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    apply AS174-30280-IPV4-IN-10
    apply AS174-30280-IPV4-IN-20
    apply AS174-30280-IPV4-IN-30
  endif
end-policy
!
prefix-set AS174-IPV6-PREFIX-DROP-IN
end-set
!
prefix-set AS38195-SLCAU-DEFAULT-ONLY
  0.0.0.0/0 eq 0,
  ::/0 eq 0
end-set
!
prefix-set AS38195-SLCAU-LE48
  ::/0 le 48
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
as-path-set AS174-ASPATH-LOCALPREF-131-IN
end-set
!
community-set AS38195-SLCAU-EXTERNAL-DELETE
  38195:*
end-set
!
route-policy AS174-30280-IPV6-IN-10
  if (destination in AS174-IPV6-PREFIX-DROP-IN) or (as-path in AS174-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS174-30280-IPV6-IN-20
  if (community matches-any (174:21201)) then
    set community (38195:9708, 38195:7350, 38195:7662, 38195:8010, 38195:6800) additive
    set med 0
    if (community matches-any (174:22038)) then
      set local-preference 142
      done
    endif
  endif
end-policy
!
route-policy AS174-30280-IPV6-IN-30
  set community (38195:9709, 38195:7350, 38195:7662, 38195:8010, 38195:6800) additive
  if (as-path in AS174-ASPATH-LOCALPREF-131-IN) then
    set local-preference 131
    done
  endif
  if (community matches-any (174:21201, 174:22005)) and (not community matches-any (174:22032)) then
    set local-preference 130
    done
  else
    set local-preference 110
    done
  endif
end-policy
!
route-policy AS174-30280-IPV6-IN
  if (destination in AS38195-SLCAU-LE48) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    apply AS174-30280-IPV6-IN-10
    apply AS174-30280-IPV6-IN-20
    apply AS174-30280-IPV6-IN-30
  endif
end-policy
!
prefix-set AS174-IPV4-PREFIX-DROP-OUT
  45.67.96.0/24
end-set
!
prefix-set AS38195-SLCAU-LE24
  0.0.0.0/0 le 24
end-set
!
as-path-set AS174-ASPATH-DROP-OUT
end-set
!
as-path-set COGENT-SYD-CUST
  ios-regex '_18200_'
end-set
!
community-set AS38195-SLCAU-CORE
  38195:9700
end-set
!
community-set AS38195-SLCAU-CUSTOMER
  38195:9701
end-set
!
community-set AS38195-SLCAU-DOMESTIC-ONLY
  38195:800
end-set
!
community-set AS38195-SLCAU-IPTRANSIT
  38195:9708,
  38195:9709
end-set
!
community-set AS38195-SLCAU-PEERING
  38195:9705,
  38195:9706
end-set
!
route-policy AS174-30280-IPV4-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS174-IPV4-PREFIX-DROP-OUT) or (as-path in AS174-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS174-30280-IPV4-OUT-11
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS174-30280-IPV4-OUT-20
  if (large-community matches-any (38195:174:0)) or (large-community matches-any (38195:174:68000)) or (large-community matches-any (38195:174:80100)) then
    drop
  endif
end-policy
!
route-policy AS174-30280-IPV4-OUT-21
  if (large-community matches-any (38195:174:1)) or (large-community matches-any (38195:174:68001)) or (large-community matches-any (38195:174:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:174:2)) or (large-community matches-any (38195:174:68002)) or (large-community matches-any (38195:174:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:174:3)) or (large-community matches-any (38195:174:68003)) or (large-community matches-any (38195:174:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS174-30280-IPV4-OUT-29
  if (large-community matches-any (38195:174:9)) or (large-community matches-any (38195:174:68009)) or (large-community matches-any (38195:174:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS174-30280-IPV4-OUT-30
  if (community matches-any (38195:18800)) or (community matches-any (38195:18830)) then
    drop
  endif
end-policy
!
route-policy AS174-30280-IPV4-OUT-31
  if (community matches-any (38195:18801)) or (community matches-any (38195:18831)) then
    prepend as-path 38195
    pass
  endif
  if (community matches-any (38195:18802)) or (community matches-any (38195:18832)) then
    prepend as-path 38195 2
    pass
  endif
  if (community matches-any (38195:18803)) or (community matches-any (38195:18833)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS174-30280-IPV4-OUT-39
  if (community matches-any (38195:18809)) or (community matches-any (38195:18839)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS174-30280-IPV4-OUT-90
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CORE) then
    set community (174:940, 174:950, 174:3110, 174:980) additive
    done
  endif
end-policy
!
route-policy AS174-30280-IPV4-OUT-91
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    set community (174:940, 174:950, 174:3110) additive
    if (as-path in COGENT-SYD-CUST) then
      set community (174:140) additive
    endif
    done
  endif
end-policy
!
route-policy AS174-30280-IPV4-OUT-99
  drop
end-policy
!
route-policy AS174-30280-IPV4-OUT
  apply AS174-30280-IPV4-OUT-10
  apply AS174-30280-IPV4-OUT-11
  apply AS174-30280-IPV4-OUT-20
  apply AS174-30280-IPV4-OUT-21
  apply AS174-30280-IPV4-OUT-29
  apply AS174-30280-IPV4-OUT-30
  apply AS174-30280-IPV4-OUT-31
  apply AS174-30280-IPV4-OUT-39
  apply AS174-30280-IPV4-OUT-90
  apply AS174-30280-IPV4-OUT-91
  apply AS174-30280-IPV4-OUT-99
end-policy
!
prefix-set AS174-IPV6-PREFIX-DROP-OUT
end-set
!
prefix-set AS38195-SLCAU-LE48
  ::/0 le 48
end-set
!
as-path-set AS174-ASPATH-DROP-OUT
end-set
!
as-path-set COGENT-SYD-CUST
  ios-regex '_18200_'
end-set
!
community-set AS38195-SLCAU-CORE
  38195:9700
end-set
!
community-set AS38195-SLCAU-CUSTOMER
  38195:9701
end-set
!
community-set AS38195-SLCAU-DOMESTIC-ONLY
  38195:800
end-set
!
community-set AS38195-SLCAU-IPTRANSIT
  38195:9708,
  38195:9709
end-set
!
community-set AS38195-SLCAU-PEERING
  38195:9705,
  38195:9706
end-set
!
route-policy AS174-30280-IPV6-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS174-IPV6-PREFIX-DROP-OUT) or (as-path in AS174-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS174-30280-IPV6-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS174-30280-IPV6-OUT-20
  if (large-community matches-any (38195:174:0)) or (large-community matches-any (38195:174:68000)) or (large-community matches-any (38195:174:80100)) then
    drop
  endif
end-policy
!
route-policy AS174-30280-IPV6-OUT-21
  if (large-community matches-any (38195:174:1)) or (large-community matches-any (38195:174:68001)) or (large-community matches-any (38195:174:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:174:2)) or (large-community matches-any (38195:174:68002)) or (large-community matches-any (38195:174:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:174:3)) or (large-community matches-any (38195:174:68003)) or (large-community matches-any (38195:174:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS174-30280-IPV6-OUT-29
  if (large-community matches-any (38195:174:9)) or (large-community matches-any (38195:174:68009)) or (large-community matches-any (38195:174:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS174-30280-IPV6-OUT-30
  if (community matches-any (38195:18800)) or (community matches-any (38195:18830)) then
    drop
  endif
end-policy
!
route-policy AS174-30280-IPV6-OUT-31
  if (community matches-any (38195:18801)) or (community matches-any (38195:18831)) then
    prepend as-path 38195
    pass
  endif
  if (community matches-any (38195:18802)) or (community matches-any (38195:18832)) then
    prepend as-path 38195 2
    pass
  endif
  if (community matches-any (38195:18803)) or (community matches-any (38195:18833)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS174-30280-IPV6-OUT-39
  if (community matches-any (38195:18809)) or (community matches-any (38195:18839)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS174-30280-IPV6-OUT-90
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CORE) then
    set community (174:940, 174:950, 174:3110) additive
    done
  endif
end-policy
!
route-policy AS174-30280-IPV6-OUT-91
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    set community (174:940, 174:950, 174:3110) additive
    if (as-path in COGENT-SYD-CUST) then
      set community (174:140) additive
    endif
    done
  endif
end-policy
!
route-policy AS174-30280-IPV6-OUT-99
  drop
end-policy
!
route-policy AS174-30280-IPV6-OUT
  apply AS174-30280-IPV6-OUT-10
  apply AS174-30280-IPV6-OUT-19
  apply AS174-30280-IPV6-OUT-20
  apply AS174-30280-IPV6-OUT-21
  apply AS174-30280-IPV6-OUT-29
  apply AS174-30280-IPV6-OUT-30
  apply AS174-30280-IPV6-OUT-31
  apply AS174-30280-IPV6-OUT-39
  apply AS174-30280-IPV6-OUT-90
  apply AS174-30280-IPV6-OUT-91
  apply AS174-30280-IPV6-OUT-99
end-policy
!
interface HundredGigE0/0/0/18 description Transit: Cogent SG2 AS174 (3-001628364)[15000Mbit]{30280}
interface HundredGigE0/0/0/18 bandwidth 15000000
interface HundredGigE0/0/0/18 mtu 9212
interface HundredGigE0/0/0/18 service-policy input AS38195-DSCP-REWRITE-IN
interface HundredGigE0/0/0/18 ipv4 mtu 1500
interface HundredGigE0/0/0/18 ipv4 address 154.18.3.90 255.255.255.248
interface HundredGigE0/0/0/18 ipv4 unreachables disable
interface HundredGigE0/0/0/18 ipv6 nd suppress-ra
interface HundredGigE0/0/0/18 ipv6 mtu 1500
interface HundredGigE0/0/0/18 ipv6 address 2402:4480:2:2::2/112
interface HundredGigE0/0/0/18 ipv6 unreachables disable
interface HundredGigE0/0/0/18 load-interval 30
interface HundredGigE0/0/0/18 flow ipv4 monitor netflow-ipv4 sampler 1in1k ingress
interface HundredGigE0/0/0/18 flow ipv6 monitor netflow-ipv6 sampler 1in1k ingress
interface HundredGigE0/0/0/18 ipv4 access-group AS38195-IPV4-COGENT-DENY ingress
interface HundredGigE0/0/0/18 ipv6 access-group AS38195-IPV6-TRANSIT-DENY ingress


ipv4 access-list AS38195-IPV4-COGENT-DENY 10 remark security filtering
ipv4 access-list AS38195-IPV4-COGENT-DENY 20 remark drop RFC1918
ipv4 access-list AS38195-IPV4-COGENT-DENY 30 deny ipv4 10.0.0.0 0.255.255.255 any
ipv4 access-list AS38195-IPV4-COGENT-DENY 40 deny ipv4 172.16.0.0 0.15.255.255 any
ipv4 access-list AS38195-IPV4-COGENT-DENY 50 deny ipv4 192.168.0.0 0.0.255.255 any
ipv4 access-list AS38195-IPV4-COGENT-DENY 60 deny ipv4 any 10.0.0.0 0.255.255.255
ipv4 access-list AS38195-IPV4-COGENT-DENY 70 deny ipv4 any 172.16.0.0 0.15.255.255
ipv4 access-list AS38195-IPV4-COGENT-DENY 80 deny ipv4 any 192.168.0.0 0.0.255.255
ipv4 access-list AS38195-IPV4-COGENT-DENY 90 remark drop RFC2544
ipv4 access-list AS38195-IPV4-COGENT-DENY 100 deny ipv4 198.18.0.0 0.1.255.255 any
ipv4 access-list AS38195-IPV4-COGENT-DENY 110 deny ipv4 any 198.18.0.0 0.1.255.255
ipv4 access-list AS38195-IPV4-COGENT-DENY 120 remark drop RFC5735
ipv4 access-list AS38195-IPV4-COGENT-DENY 130 deny ipv4 0.0.0.0 0.255.255.255 any
ipv4 access-list AS38195-IPV4-COGENT-DENY 140 deny ipv4 127.0.0.0 0.255.255.255 any
ipv4 access-list AS38195-IPV4-COGENT-DENY 150 deny ipv4 169.254.0.0 0.0.255.255 any
ipv4 access-list AS38195-IPV4-COGENT-DENY 160 deny ipv4 any 0.0.0.0 0.255.255.255
ipv4 access-list AS38195-IPV4-COGENT-DENY 170 deny ipv4 any 127.0.0.0 0.255.255.255
ipv4 access-list AS38195-IPV4-COGENT-DENY 180 deny ipv4 any 169.254.0.0 0.0.255.255
ipv4 access-list AS38195-IPV4-COGENT-DENY 190 remark drop RFC6598
ipv4 access-list AS38195-IPV4-COGENT-DENY 200 deny ipv4 100.64.0.0 0.63.255.255 any
ipv4 access-list AS38195-IPV4-COGENT-DENY 210 deny ipv4 any 100.64.0.0 0.63.255.255
ipv4 access-list AS38195-IPV4-COGENT-DENY 220 remark drop TEST-NETS
ipv4 access-list AS38195-IPV4-COGENT-DENY 230 deny ipv4 192.0.2.0 0.0.0.255 any
ipv4 access-list AS38195-IPV4-COGENT-DENY 240 deny ipv4 198.51.100.0 0.0.0.255 any
ipv4 access-list AS38195-IPV4-COGENT-DENY 250 deny ipv4 203.0.113.0 0.0.0.255 any
ipv4 access-list AS38195-IPV4-COGENT-DENY 260 deny ipv4 any 192.0.2.0 0.0.0.255
ipv4 access-list AS38195-IPV4-COGENT-DENY 270 deny ipv4 any 198.51.100.0 0.0.0.255
ipv4 access-list AS38195-IPV4-COGENT-DENY 280 deny ipv4 any 203.0.113.0 0.0.0.255
ipv4 access-list AS38195-IPV4-COGENT-DENY 500 remark PDK DDoS CHG-000015779
ipv4 access-list AS38195-IPV4-COGENT-DENY 510 deny udp 91.246.49.0/24 95.82.0.0/22
ipv4 access-list AS38195-IPV4-COGENT-DENY 511 deny udp 91.246.49.0/24 95.82.4.0/22
ipv4 access-list AS38195-IPV4-COGENT-DENY 512 deny udp 91.246.49.0/24 95.82.32.0/22
ipv4 access-list AS38195-IPV4-COGENT-DENY 513 deny udp 91.246.49.0/24 95.82.36.0/22
ipv4 access-list AS38195-IPV4-COGENT-DENY 520 deny udp 109.122.242.0/24 95.82.0.0/22
ipv4 access-list AS38195-IPV4-COGENT-DENY 521 deny udp 109.122.242.0/24 95.82.4.0/22
ipv4 access-list AS38195-IPV4-COGENT-DENY 522 deny udp 109.122.242.0/24 95.82.32.0/22
ipv4 access-list AS38195-IPV4-COGENT-DENY 523 deny udp 109.122.242.0/24 95.82.36.0/22
ipv4 access-list AS38195-IPV4-COGENT-DENY 530 deny udp 85.9.112.0/24 95.82.0.0/22
ipv4 access-list AS38195-IPV4-COGENT-DENY 531 deny udp 85.9.112.0/24 95.82.4.0/22
ipv4 access-list AS38195-IPV4-COGENT-DENY 532 deny udp 85.9.112.0/24 95.82.32.0/22
ipv4 access-list AS38195-IPV4-COGENT-DENY 533 deny udp 85.9.112.0/24 95.82.36.0/22
ipv4 access-list AS38195-IPV4-COGENT-DENY 9990 remark permit the rest
ipv4 access-list AS38195-IPV4-COGENT-DENY 9999 permit ipv4 any any
interface HundredGigE0/0/0/18 ipv4 access-group AS38195-IPV4-COGENT-DENY ingress


router static address-family ipv4 unicast 66.28.8.2/32 154.18.3.89 description COGENT-AS174-BLACKHOLEV4
router bgp 64590 neighbor 154.18.3.89 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 154.18.3.89 remote-as 174
router bgp 64590 neighbor 154.18.3.89 timers 10 30
router bgp 64590 neighbor 154.18.3.89 description Transit:AS174-SID30280
router bgp 64590 neighbor 154.18.3.89 address-family ipv4 unicast
router bgp 64590 neighbor 154.18.3.89 address-family ipv4 unicast maximum-prefix 1048576 90 restart 10
router bgp 64590 neighbor 154.18.3.89 address-family ipv4 unicast send-community-ebgp
router bgp 64590 neighbor 154.18.3.89 address-family ipv4 unicast route-policy AS174-30280-IPV4-IN in
router bgp 64590 neighbor 154.18.3.89 address-family ipv4 unicast route-policy AS174-30280-IPV4-OUT out
router bgp 64590 neighbor 154.18.3.89 address-family ipv4 unicast next-hop-self
router bgp 64590 neighbor 154.18.3.89 address-family ipv4 unicast remove-private-AS
router bgp 64590 neighbor 154.18.3.89 address-family ipv4 unicast soft-reconfiguration inbound always

router static address-family ipv6 unicast 2001:550:0:1000::421c:802/128 2402:4480:2:2::1 description COGENT-AS174-BLACKHOLEV6
router bgp 64590 neighbor 2402:4480:2:2::1 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 2402:4480:2:2::1 remote-as 174
router bgp 64590 neighbor 2402:4480:2:2::1 timers 10 30
router bgp 64590 neighbor 2402:4480:2:2::1 description Transit:AS174-SID30280
router bgp 64590 neighbor 2402:4480:2:2::1 address-family ipv6 unicast
router bgp 64590 neighbor 2402:4480:2:2::1 address-family ipv6 unicast send-community-ebgp
router bgp 64590 neighbor 2402:4480:2:2::1 address-family ipv6 unicast route-policy AS174-30280-IPV6-IN in
router bgp 64590 neighbor 2402:4480:2:2::1 address-family ipv6 unicast route-policy AS174-30280-IPV6-OUT out
router bgp 64590 neighbor 2402:4480:2:2::1 address-family ipv6 unicast next-hop-self
router bgp 64590 neighbor 2402:4480:2:2::1 address-family ipv6 unicast remove-private-AS
router bgp 64590 neighbor 2402:4480:2:2::1 address-family ipv6 unicast soft-reconfiguration inbound always



#SingTel SG

bdr01-eip-26aayer


prefix-set AS7473-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
!
as-path-set AS7473-ASPATH-DROP-IN
  ios-regex '_1221_',
  ios-regex '_7474_',
  ios-regex '_4804_',
  ios-regex '_6939_',
  ios-regex '_7545_',
  ios-regex '_7575_',
  ios-regex '_2764_',
  ios-regex '_4826_',
  ios-regex '_6939_'
end-set
!
as-path-set AS7473-ASPATH-LOCALPREF-143-IN
end-set
!
as-path-set AS7473-DEPREF-IN
end-set
!
as-path-set AS7473-DOMESTIC-DEPREF
end-set
!
community-set AS38195-SLCAU-EXTERNAL-DELETE
  38195:*
end-set
!
community-set AS7473-DOMESTIC-IN
  # Singapore
  7473:41101
end-set
!
community-set AS7473-ROUTE-SOURCE-APAC
  # APAC customers
  7473:41103,
  7473:41104,
  7473:41106,
  7473:21075
end-set
!
route-policy AS7473-38524-IPV4-IN-10
  if (destination in AS7473-IPV4-PREFIX-DROP-IN) or (as-path in AS7473-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS7473-38524-IPV4-IN-20
  if (community matches-any AS7473-DOMESTIC-IN) and not (as-path in AS7473-DOMESTIC-DEPREF) then
    set community (38195:9708, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    if (as-path in AS7473-ASPATH-LOCALPREF-143-IN) then
      set local-preference 143
      done
    endif
    set local-preference 142
    done
  endif
end-policy
!
route-policy AS7473-38524-IPV4-IN-30
  if (community matches-any AS7473-ROUTE-SOURCE-APAC) then
    set community (38195:9709, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    if not (as-path in AS7473-DEPREF-IN) then
      set local-preference 130
      done
    endif
    set local-preference 129
  else
    drop
  endif
end-policy
!
route-policy AS7473-38524-IPV4-IN
  if (destination in AS38195-SLCAU-LE24) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    apply AS7473-38524-IPV4-IN-10
    apply AS7473-38524-IPV4-IN-20
    apply AS7473-38524-IPV4-IN-30
  else
    drop
  endif
end-policy
!
prefix-set AS7473-IPV4-PREFIX-DROP-IN
  10.0.0.0/8 le 32,
  172.16.0.0/12 le 32,
  192.168.0.0/16 le 32
end-set
!
as-path-set AS7473-ASPATH-DROP-IN
  ios-regex '_1221_',
  ios-regex '_7474_',
  ios-regex '_4804_',
  ios-regex '_6939_',
  ios-regex '_7545_',
  ios-regex '_7575_',
  ios-regex '_2764_',
  ios-regex '_4826_',
  ios-regex '_6939_'
end-set
!
as-path-set AS7473-ASPATH-LOCALPREF-143-IN
end-set
!
as-path-set AS7473-DEPREF-IN
end-set
!
as-path-set AS7473-DOMESTIC-DEPREF
end-set
!
community-set AS38195-SLCAU-EXTERNAL-DELETE
  38195:*
end-set
!
community-set AS7473-DOMESTIC-IN
  # Singapore
  7473:41101
end-set
!
community-set AS7473-ROUTE-SOURCE-APAC
  # APAC customers
  7473:41103,
  7473:41104,
  7473:41106,
  7473:21075
end-set
!
route-policy AS7473-38524-IPV4-IN-10
  if (destination in AS7473-IPV4-PREFIX-DROP-IN) or (as-path in AS7473-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS7473-38524-IPV4-IN-20
  if (community matches-any AS7473-DOMESTIC-IN) and not (as-path in AS7473-DOMESTIC-DEPREF) then
    set community (38195:9708, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    if (as-path in AS7473-ASPATH-LOCALPREF-143-IN) then
      set local-preference 143
      done
    endif
    set local-preference 142
    done
  endif
end-policy
!
route-policy AS7473-38524-IPV4-IN-30
  if (community matches-any AS7473-ROUTE-SOURCE-APAC) then
    set community (38195:9709, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    if not (as-path in AS7473-DEPREF-IN) then
      set local-preference 130
      done
    endif
    set local-preference 129
  else
    drop
  endif
end-policy
!
route-policy AS7473-38524-IPV4-IN
  if (destination in AS38195-SLCAU-LE24) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    apply AS7473-38524-IPV4-IN-10
    apply AS7473-38524-IPV4-IN-20
    apply AS7473-38524-IPV4-IN-30
  else
    drop
  endif
end-policy
!
RP/0/RP0/CPU0:bdr01-ipt-26aayerr-sin.sg#sh rpl route-policy AS7473-38524-IPV6-IN detail
Wed Apr  2 11:52:39.520 UTC
prefix-set AS38195-SLCAU-DEFAULT-ONLY
  0.0.0.0/0 eq 0,
  ::/0 eq 0
end-set
!
prefix-set AS38195-SLCAU-LE48
  ::/0 le 48
end-set
!
prefix-set AS7473-IPV6-PREFIX-DROP-IN
end-set
!
as-path-set AS7473-ASPATH-DROP-IN
  ios-regex '_1221_',
  ios-regex '_7474_',
  ios-regex '_4804_',
  ios-regex '_6939_',
  ios-regex '_7545_',
  ios-regex '_7575_',
  ios-regex '_2764_',
  ios-regex '_4826_',
  ios-regex '_6939_'
end-set
!
as-path-set AS7473-ASPATH-LOCALPREF-143-IN
end-set
!
as-path-set AS7473-DEPREF-IN
end-set
!
as-path-set AS7473-DOMESTIC-DEPREF
end-set
!
community-set AS38195-SLCAU-EXTERNAL-DELETE
  38195:*
end-set
!
community-set AS7473-DOMESTIC-IN
  # Singapore
  7473:41101
end-set
!
community-set AS7473-ROUTE-SOURCE-APAC
  # APAC customers
  7473:41103,
  7473:41104,
  7473:41106,
  7473:21075
end-set
!
route-policy AS7473-38524-IPV6-IN-10
  if (destination in AS7473-IPV6-PREFIX-DROP-IN) or (as-path in AS7473-ASPATH-DROP-IN) then
    drop
  endif
  if (destination in AS38195-SLCAU-DEFAULT-ONLY) or (validation-state is invalid) then
    drop
  endif
end-policy
!
route-policy AS7473-38524-IPV6-IN-20
  if (community matches-any AS7473-DOMESTIC-IN) and not (as-path in AS7473-DOMESTIC-DEPREF) then
    set community (38195:9708, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    if (as-path in AS7473-ASPATH-LOCALPREF-143-IN) then
      set local-preference 143
      done
    endif
    set local-preference 142
    done
  endif
end-policy
!
route-policy AS7473-38524-IPV6-IN-30
  if (community matches-any AS7473-ROUTE-SOURCE-APAC) then
    set community (38195:9709, 38195:7349, 38195:7663, 38195:8010, 38195:6800) additive
    if not (as-path in AS7473-DEPREF-IN) then
      set local-preference 130
      done
    endif
    set local-preference 129
  else
    drop
  endif
end-policy
!
route-policy AS7473-38524-IPV6-IN
  if (destination in AS38195-SLCAU-LE48) then
    delete community in AS38195-SLCAU-EXTERNAL-DELETE
    apply AS7473-38524-IPV6-IN-10
    apply AS7473-38524-IPV6-IN-20
    apply AS7473-38524-IPV6-IN-30
  else
    drop
  endif
end-policy
!
prefix-set AS7473-IPV4-PREFIX-DROP-OUT
end-set
!
as-path-set AS7473-ASPATH-DROP-OUT
  ios-regex '_54994_'
end-set
!
community-set AS38195-SLCAU-BLACKHOLE
  38195:666
end-set
!
community-set AS38195-SLCAU-CORE
  38195:9700
end-set
!
community-set AS38195-SLCAU-CUSTOMER
  38195:9701
end-set
!
community-set AS38195-SLCAU-DOMESTIC-ONLY
  38195:800
end-set
!
community-set AS38195-SLCAU-IPTRANSIT
  38195:9708,
  38195:9709
end-set
!
community-set AS38195-SLCAU-PEERING
  38195:9705,
  38195:9706
end-set
!
route-policy AS7473-38524-IPV4-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS7473-IPV4-PREFIX-DROP-OUT) or (as-path in AS7473-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS7473-38524-IPV4-OUT-15
  if (destination in AS38195-SLCAU-GE32) and (community matches-any AS38195-SLCAU-BLACKHOLE) then
    set community (7473:995) additive
    done
  endif
end-policy
!
route-policy AS7473-38524-IPV4-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS7473-38524-IPV4-OUT-20
  if (large-community matches-any (38195:7473:0)) or (large-community matches-any (38195:7473:68000)) or (large-community matches-any (38195:7473:80100)) then
    drop
  endif
end-policy
!
route-policy AS7473-38524-IPV4-OUT-21
  if (large-community matches-any (38195:7473:1)) or (large-community matches-any (38195:7473:68001)) or (large-community matches-any (38195:7473:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:7473:2)) or (large-community matches-any (38195:7473:68002)) or (large-community matches-any (38195:7473:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:7473:3)) or (large-community matches-any (38195:7473:68003)) or (large-community matches-any (38195:7473:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS7473-38524-IPV4-OUT-29
  if (large-community matches-any (38195:7473:9)) or (large-community matches-any (38195:7473:68009)) or (large-community matches-any (38195:7473:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS7473-38524-IPV4-OUT-30
  if (community matches-any (38195:17800)) then
    drop
  endif
end-policy
!
route-policy AS7473-38524-IPV4-OUT-31
  if (community matches-any (38195:17801)) then
    prepend as-path 38195
    pass
  endif
  if (community matches-any (38195:17802)) then
    prepend as-path 38195 2
    pass
  endif
  if (community matches-any (38195:17803)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS7473-38524-IPV4-OUT-39
  if (community matches-any (38195:17809)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS7473-38524-IPV4-OUT-90
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CORE) then
    set community (7473:22009, 7473:23009, 7473:21029, 7473:21129, 7473:21073, 7473:21032, 7473:21042, 9:6939, 9:8075) additive
    done
  endif
end-policy
!
route-policy AS7473-38524-IPV4-OUT-91
  if (destination in AS38195-SLCAU-LE24) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    set community (7473:22009, 7473:23009, 7473:21029, 7473:21129, 7473:21073, 7473:21032, 7473:21042, 9:6939, 9:8075) additive
    done
  endif
end-policy
!
route-policy AS7473-38524-IPV4-OUT-99
  drop
end-policy
!
route-policy AS7473-38524-IPV4-OUT
  apply AS7473-38524-IPV4-OUT-10
  apply AS7473-38524-IPV4-OUT-15
  apply AS7473-38524-IPV4-OUT-19
  apply AS7473-38524-IPV4-OUT-20
  apply AS7473-38524-IPV4-OUT-21
  apply AS7473-38524-IPV4-OUT-29
  apply AS7473-38524-IPV4-OUT-30
  apply AS7473-38524-IPV4-OUT-31
  apply AS7473-38524-IPV4-OUT-39
  apply AS7473-38524-IPV4-OUT-90
  apply AS7473-38524-IPV4-OUT-91
  apply AS7473-38524-IPV4-OUT-99
end-policy
!
prefix-set AS7473-IPV6-PREFIX-DROP-OUT
end-set
!
as-path-set AS7473-ASPATH-DROP-OUT
  ios-regex '_54994_'
end-set
!
community-set AS38195-SLCAU-BLACKHOLE
  38195:666
end-set
!
community-set AS38195-SLCAU-CORE
  38195:9700
end-set
!
community-set AS38195-SLCAU-CUSTOMER
  38195:9701
end-set
!
community-set AS38195-SLCAU-DOMESTIC-ONLY
  38195:800
end-set
!
community-set AS38195-SLCAU-IPTRANSIT
  38195:9708,
  38195:9709
end-set
!
community-set AS38195-SLCAU-PEERING
  38195:9705,
  38195:9706
end-set
!
route-policy AS7473-38524-IPV6-OUT-10
  if (community matches-any AS38195-SLCAU-PEERING) or (community matches-any AS38195-SLCAU-IPTRANSIT) then
    drop
  endif
  if (destination in AS7473-IPV6-PREFIX-DROP-OUT) or (as-path in AS7473-ASPATH-DROP-OUT) then
    drop
  endif
end-policy
!
route-policy AS7473-38524-IPV6-OUT-15
  if (destination in AS38195-SLCAU-GE128) and (community matches-any AS38195-SLCAU-BLACKHOLE) then
    set community (7473:995) additive
    done
  endif
end-policy
!
route-policy AS7473-38524-IPV6-OUT-19
  if (community matches-any AS38195-SLCAU-DOMESTIC-ONLY) then
    if (community matches-any (38195:6800)) then
      pass
    else
      drop
    endif
  endif
end-policy
!
route-policy AS7473-38524-IPV6-OUT-20
  if (large-community matches-any (38195:7473:0)) or (large-community matches-any (38195:7473:68000)) or (large-community matches-any (38195:7473:80100)) then
    drop
  endif
end-policy
!
route-policy AS7473-38524-IPV6-OUT-21
  if (large-community matches-any (38195:7473:1)) or (large-community matches-any (38195:7473:68001)) or (large-community matches-any (38195:7473:80101)) then
    prepend as-path 38195
    pass
  endif
  if (large-community matches-any (38195:7473:2)) or (large-community matches-any (38195:7473:68002)) or (large-community matches-any (38195:7473:80102)) then
    prepend as-path 38195 2
    pass
  endif
  if (large-community matches-any (38195:7473:3)) or (large-community matches-any (38195:7473:68003)) or (large-community matches-any (38195:7473:80103)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS7473-38524-IPV6-OUT-29
  if (large-community matches-any (38195:7473:9)) or (large-community matches-any (38195:7473:68009)) or (large-community matches-any (38195:7473:80109)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS7473-38524-IPV6-OUT-30
  if (community matches-any (38195:17800)) then
    drop
  endif
end-policy
!
route-policy AS7473-38524-IPV6-OUT-31
  if (community matches-any (38195:17801)) then
    prepend as-path 38195
    pass
  endif
  if (community matches-any (38195:17802)) then
    prepend as-path 38195 2
    pass
  endif
  if (community matches-any (38195:17803)) then
    prepend as-path 38195 3
    pass
  endif
end-policy
!
route-policy AS7473-38524-IPV6-OUT-39
  if (community matches-any (38195:17809)) then
    set community (no-export)
    pass
  endif
end-policy
!
route-policy AS7473-38524-IPV6-OUT-90
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CORE) then
    set community (7473:22009, 7473:23009, 7473:21029, 7473:21129, 7473:21073, 7473:21032, 7473:21042, 9:6939, 9:8075) additive
    done
  endif
end-policy
!
route-policy AS7473-38524-IPV6-OUT-91
  if (destination in AS38195-SLCAU-LE48) and (community matches-any AS38195-SLCAU-CUSTOMER) then
    set community (7473:22009, 7473:23009, 7473:21029, 7473:21129, 7473:21073, 7473:21032, 7473:21042, 9:6939, 9:8075) additive
    done
  endif
end-policy
!
route-policy AS7473-38524-IPV6-OUT-99
  drop
end-policy
!
route-policy AS7473-38524-IPV6-OUT
  apply AS7473-38524-IPV6-OUT-10
  apply AS7473-38524-IPV6-OUT-15
  apply AS7473-38524-IPV6-OUT-19
  apply AS7473-38524-IPV6-OUT-20
  apply AS7473-38524-IPV6-OUT-21
  apply AS7473-38524-IPV6-OUT-29
  apply AS7473-38524-IPV6-OUT-30
  apply AS7473-38524-IPV6-OUT-31
  apply AS7473-38524-IPV6-OUT-39
  apply AS7473-38524-IPV6-OUT-90
  apply AS7473-38524-IPV6-OUT-91
  apply AS7473-38524-IPV6-OUT-99
end-policy
!


router bgp 64590 neighbor 203.208.168.249 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 203.208.168.249 remote-as 7473
router bgp 64590 neighbor 203.208.168.249 timers 10 30
router bgp 64590 neighbor 203.208.168.249 password encrypted 04683F57177D1B1A5E4A5B24071B5F1606242B386F667A425E435D
router bgp 64590 neighbor 203.208.168.249 description Transit:AS7473-SID38524
router bgp 64590 neighbor 203.208.168.249 address-family ipv4 unicast
router bgp 64590 neighbor 203.208.168.249 address-family ipv4 unicast send-community-ebgp
router bgp 64590 neighbor 203.208.168.249 address-family ipv4 unicast route-policy AS7473-38524-IPV4-IN in
router bgp 64590 neighbor 203.208.168.249 address-family ipv4 unicast route-policy AS7473-38524-IPV4-OUT out
router bgp 64590 neighbor 203.208.168.249 address-family ipv4 unicast next-hop-self
router bgp 64590 neighbor 203.208.168.249 address-family ipv4 unicast remove-private-AS
router bgp 64590 neighbor 203.208.168.249 address-family ipv4 unicast soft-reconfiguration inbound always

router bgp 64590 neighbor 2001:c10:80:2::ac9 local-as 38195 no-prepend replace-as
router bgp 64590 neighbor 2001:c10:80:2::ac9 remote-as 7473
router bgp 64590 neighbor 2001:c10:80:2::ac9 timers 10 30
router bgp 64590 neighbor 2001:c10:80:2::ac9 password encrypted 0538325E3910195D4E5649211E1C5738072B272369714B564F565F
router bgp 64590 neighbor 2001:c10:80:2::ac9 description Transit:AS7473-SID38524
router bgp 64590 neighbor 2001:c10:80:2::ac9 address-family ipv6 unicast
router bgp 64590 neighbor 2001:c10:80:2::ac9 address-family ipv6 unicast send-community-ebgp
router bgp 64590 neighbor 2001:c10:80:2::ac9 address-family ipv6 unicast route-policy AS7473-38524-IPV6-IN in
router bgp 64590 neighbor 2001:c10:80:2::ac9 address-family ipv6 unicast route-policy AS7473-38524-IPV6-OUT out
router bgp 64590 neighbor 2001:c10:80:2::ac9 address-family ipv6 unicast next-hop-self
router bgp 64590 neighbor 2001:c10:80:2::ac9 address-family ipv6 unicast remove-private-AS
router bgp 64590 neighbor 2001:c10:80:2::ac9 address-family ipv6 unicast soft-reconfiguration inbound always


l2vpn xconnect group xc_38524
l2vpn xconnect group xc_38524 p2p p2p_38524
l2vpn xconnect group xc_38524 p2p p2p_38524 interface HundredGigE0/0/0/22.38524
l2vpn xconnect group xc_38524 p2p p2p_38524 neighbor ipv4 172.21.55.14 pw-id 38524
l2vpn xconnect group xc_38524 p2p p2p_38524 neighbor ipv4 172.21.55.14 pw-id 38524 pw-class CW-enabled

interface HundredGigE0/0/0/22.38524 l2transport
interface HundredGigE0/0/0/22.38524 l2transport description Loop-L2: SingTel SG3 AS7473 - SP-SGT-IP10 [10000Mbit]{38524}
interface HundredGigE0/0/0/22.38524 l2transport encapsulation dot1q 55
interface HundredGigE0/0/0/22.38524 l2transport rewrite ingress tag pop 1 symmetric
interface HundredGigE0/0/0/22.38524 l2transport mtu 9216


interface HundredGigE0/0/0/23.55 description Transit: SingTel SG3 AS7473 - SP-SGT-IP10 [10000Mbit]{38524}
interface HundredGigE0/0/0/23.55 mtu 9212
interface HundredGigE0/0/0/23.55 service-policy input AS38195-DSCP-REWRITE-IN
interface HundredGigE0/0/0/23.55 ipv4 mtu 1500
interface HundredGigE0/0/0/23.55 ipv4 address 203.208.168.250 255.255.255.252
interface HundredGigE0/0/0/23.55 ipv4 unreachables disable
interface HundredGigE0/0/0/23.55 ipv6 nd suppress-ra
interface HundredGigE0/0/0/23.55 ipv6 mtu 1500
interface HundredGigE0/0/0/23.55 ipv6 address 2001:c10:80:2::aca/126
interface HundredGigE0/0/0/23.55 ipv6 unreachables disable
interface HundredGigE0/0/0/23.55 load-interval 30
interface HundredGigE0/0/0/23.55 flow ipv4 monitor netflow-ipv4 sampler 1in1k ingress
interface HundredGigE0/0/0/23.55 flow ipv6 monitor netflow-ipv6 sampler 1in1k ingress
interface HundredGigE0/0/0/23.55 encapsulation dot1q 55
interface HundredGigE0/0/0/23.55 ipv4 access-group AS38195-IPV4-TRANSIT-DENY ingress
interface HundredGigE0/0/0/23.55 ipv6 access-group AS38195-IPV6-TRANSIT-DENY ingress



#bdr02-etp-26aayer:

l2vpn xconnect group xc_38524
l2vpn xconnect group xc_38524 p2p p2p_38524
l2vpn xconnect group xc_38524 p2p p2p_38524 interface Te0/0/0/8
l2vpn xconnect group xc_38524 p2p p2p_38524 neighbor ipv4 202.177.40.4 pw-id 38524
l2vpn xconnect group xc_38524 p2p p2p_38524 neighbor ipv4 202.177.40.4 pw-id 38524 pw-class CW-enabled


interface Te0/0/0/8 description Transit: SingTel SG3 AS7473 - SP-SGT-IP10 [10000Mbit]{38524}
interface Te0/0/0/8 mtu 9216
interface Te0/0/0/8 load-interval 30
interface Te0/0/0/8 l2transport

