






#EIP-15pioneer:

    interface HundredGigE0/0/0/22.1345949 l2transport
    interface HundredGigE0/0/0/22.1345949 l2transport description Cust: Superloop IP Transit (15pionee)[100Gbit]{1345949}
    interface HundredGigE0/0/0/22.1345949 l2transport encapsulation dot1q 772
    interface HundredGigE0/0/0/22.1345949 l2transport mtu 9216
    interface HundredGigE0/0/0/23.772
    interface HundredGigE0/0/0/23.772 description NNI-IPT-L3: bdr01-eip-15pionee-sin.sg 0/0/0/23.772 ()[100Gbit]{1345949}
    interface HundredGigE0/0/0/23.772 mtu 9212
    interface HundredGigE0/0/0/23.772 ipv6 mtu 4470
    interface HundredGigE0/0/0/23.772 ipv6 address 2401:d000:10:800::61/124
    interface HundredGigE0/0/0/23.772 ipv6 unreachables disable
    interface HundredGigE0/0/0/23.772 load-interval 30
    interface HundredGigE0/0/0/23.772 encapsulation dot1q 772

    l2vpn xconnect group xc_1345949
    l2vpn xconnect group xc_1345949 p2p p2p_1345949
    l2vpn xconnect group xc_1345949 p2p p2p_1345949 interface HundredGigE0/0/0/22.1345949
    l2vpn xconnect group xc_1345949 p2p p2p_1345949 neighbor ipv4 202.177.40.4 pw-id 1345949



    router isis core interface HundredGigE0/0/0/23.772
    router isis core interface HundredGigE0/0/0/23.772 circuit-type level-2-only
    router isis core interface HundredGigE0/0/0/23.772 bfd minimum-interval 200
    router isis core interface HundredGigE0/0/0/23.772 bfd multiplier 5
    router isis core interface HundredGigE0/0/0/23.772 bfd fast-detect ipv6
    router isis core interface HundredGigE0/0/0/23.772 point-to-point
    router isis core interface HundredGigE0/0/0/23.772 hello-padding disable
    router isis core interface HundredGigE0/0/0/23.772 hello-password hmac-md5 as38195-is-is
    router isis core interface HundredGigE0/0/0/23.772 address-family ipv6 unicast
    router isis core interface HundredGigE0/0/0/23.772 address-family ipv6 unicast metric 30

    bfd interface HundredGigE0/0/0/23.772
    bfd interface HundredGigE0/0/0/23.772 echo disable

    mpls traffic-eng interface HundredGigE0/0/0/23.772
    mpls ldp interface HundredGigE0/0/0/23.772






#EIP-26aayerr:

    interface HundredGigE0/0/0/22.1345949 l2transport
    interface HundredGigE0/0/0/22.1345949 l2transport description Cust: Superloop IP Transit (26aayerr)[100Gbit]{1345949}
    interface HundredGigE0/0/0/22.1345949 l2transport encapsulation dot1q 772
    interface HundredGigE0/0/0/22.1345949 l2transport mtu 9216
    interface HundredGigE0/0/0/23.772
    interface HundredGigE0/0/0/23.772 description NNI-IPT-L3: bdr01-eip-15pionee-sin.sg 0/0/0/23.772 ()[100Gbit]{1345949}
    interface HundredGigE0/0/0/23.772 mtu 9212
    interface HundredGigE0/0/0/23.772 ipv6 mtu 4470
    interface HundredGigE0/0/0/23.772 ipv6 address 2401:d000:10:800::62/124
    interface HundredGigE0/0/0/23.772 ipv6 unreachables disable
    interface HundredGigE0/0/0/23.772 load-interval 30
    interface HundredGigE0/0/0/23.772 encapsulation dot1q 772

    l2vpn xconnect group xc_1345949
    l2vpn xconnect group xc_1345949 p2p p2p_1345949
    l2vpn xconnect group xc_1345949 p2p p2p_1345949 interface HundredGigE0/0/0/22.1345949
    l2vpn xconnect group xc_1345949 p2p p2p_1345949 neighbor ipv4 202.177.40.3 pw-id 1345949



    router isis core interface HundredGigE0/0/0/23.772
    router isis core interface HundredGigE0/0/0/23.772 circuit-type level-2-only
    router isis core interface HundredGigE0/0/0/23.772 bfd minimum-interval 200
    router isis core interface HundredGigE0/0/0/23.772 bfd multiplier 5
    router isis core interface HundredGigE0/0/0/23.772 bfd fast-detect ipv6
    router isis core interface HundredGigE0/0/0/23.772 point-to-point
    router isis core interface HundredGigE0/0/0/23.772 hello-padding disable
    router isis core interface HundredGigE0/0/0/23.772 hello-password hmac-md5 encrypted 094D5D5A41544E47460517677617
    router isis core interface HundredGigE0/0/0/23.772 address-family ipv6 unicast
    router isis core interface HundredGigE0/0/0/23.772 address-family ipv6 unicast metric 30

    bfd interface HundredGigE0/0/0/23.772
    bfd interface HundredGigE0/0/0/23.772 echo disable

    mpls traffic-eng interface HundredGigE0/0/0/23.772
    mpls ldp interface HundredGigE0/0/0/23.772
