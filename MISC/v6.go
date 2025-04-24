v6:


SG2 - PIX
SG3 - P1
SG2 - S1
SG3 - SY3

===========================================================================

SG2 - PIX:

build layer 2:
    bdr01-eip-15pioneer:
        172.21.16.13
        172.21.0.41

    cor01-tpt-4milros:
        172.21.16.12
        172.21.55.11

Build layer 3:
    bdr01-eip-15pioneer:
        bfd interface HundredGigE0/0/0/23.772
        bfd interface HundredGigE0/0/0/23.772 echo disable
        interface HundredGigE0/0/0/23.772
        interface HundredGigE0/0/0/23.772 description NNI-IPT-L3: bdr01-eip-15pionee-sin.sg 0/0/0/23.772 ()[100Gbit]{1345949}
        interface HundredGigE0/0/0/23.772 mtu 9212
        interface HundredGigE0/0/0/23.772 ipv6 mtu 4470
        interface HundredGigE0/0/0/23.772 ipv6 address 2401:d000:10:800::62/124
        interface HundredGigE0/0/0/23.772 ipv6 unreachables disable
        interface HundredGigE0/0/0/23.772 load-interval 30
        interface HundredGigE0/0/0/23.772 encapsulation dot1q 772
        router isis core interface HundredGigE0/0/0/23.772
        router isis core interface HundredGigE0/0/0/23.772 circuit-type level-2-only
        router isis core interface HundredGigE0/0/0/23.772 bfd minimum-interval 200
        router isis core interface HundredGigE0/0/0/23.772 bfd multiplier 5
        router isis core interface HundredGigE0/0/0/23.772 bfd fast-detect ipv6
        router isis core interface HundredGigE0/0/0/23.772 point-to-point
        router isis core interface HundredGigE0/0/0/23.772 hello-padding disable
        router isis core interface HundredGigE0/0/0/23.772 hello-password hmac-md5 encrypted 070E321F165840505A1B18410D39
        router isis core interface HundredGigE0/0/0/23.772 address-family ipv6 unicast
        router isis core interface HundredGigE0/0/0/23.772 address-family ipv6 unicast metric 30
        mpls traffic-eng interface HundredGigE0/0/0/23.772
        mpls ldp interface HundredGigE0/0/0/23.772


