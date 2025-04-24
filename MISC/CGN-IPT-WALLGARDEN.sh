config system interface
    edit "Be71.4089-Local"
        set vdom "WALL_GARDEN"
        set ip 202.90.207.233 255.255.255.254
        set allowaccess ping
        set description "NNI-IPT: bdr03-ipt-454stpau-bne - Be21.4089 [20000Mbit]{1346051}"
        set alias "WALL to IPT Local Peer"
        set device-identification enable
        set monitor-bandwidth enable
        set role lan
        set snmp-index 72
        set interface "Bunde-Ether30"
        set vlanid 4089
    next
end
end

config vdom
edit WALL_GARDEN
config router route-map
edit "AS38195_V4_IPT_LOCAL_PEER_OUT"
    config rule
        edit 1
            set match-ip-address "AS38195_V4_IPT_LOCAL_PEER_OUT"
            set set-community "no-export"
            set set-community-additive enable
        next
    end
next
end


config router prefix-list
    edit "AS38195_V4_IPT_LOCAL_PEER_OUT"
        config rule
            edit 1
                set prefix 202.90.207.232 255.255.255.255
                unset ge
                unset le
            next
            edit 2
                set action deny
                set prefix any
                unset ge
                unset le
            next
        end
    next
end

config router bgp
    config neighbor
        edit "202.90.207.232"
            set capability-graceful-restart enable
            set soft-reconfiguration enable
            set remote-as 38195
            set route-map-out "AS38195_V4_IPT_LOCAL_PEER_OUT"
        next



config system zone
    edit "OUTSIDE"
        set description "Outside Interface"
        set interface "Be30.4089-Local"
    next
end











#IPT:

interface Bundle-Ether41.4089
interface Bundle-Ether41.4089 description Cust: Superloop Australia (Wall Garden)(AS64595)[200000Mbit]{1346051}
interface Bundle-Ether41.4089 service-policy input AS38195-DSCP-REWRITE-IN
interface Bundle-Ether41.4089 ipv4 mtu 1500
interface Bundle-Ether41.4089 ipv4 address 202.90.207.232 255.255.255.254
interface Bundle-Ether41.4089 ipv4 unreachables disable
interface Bundle-Ether41.4089 load-interval 30
interface Bundle-Ether41.4089 flow ipv4 monitor netflow-ipv4 sampler 1in1k ingress
interface Bundle-Ether41.4089 encapsulation dot1q 4089

router bgp 38195 neighbor 202.90.207.233
router bgp 38195 neighbor 202.90.207.233 remote-as 64595
router bgp 38195 neighbor 202.90.207.233 timers 10 30
router bgp 38195 neighbor 202.90.207.233 description Cust:CGN2-WG-1346051
router bgp 38195 neighbor 202.90.207.233 ignore-connected-check
router bgp 38195 neighbor 202.90.207.233 address-family ipv4 unicast
router bgp 38195 neighbor 202.90.207.233 address-family ipv4 unicast maximum-prefix 100 90 restart 15
router bgp 38195 neighbor 202.90.207.233 address-family ipv4 unicast send-community-ebgp
router bgp 38195 neighbor 202.90.207.233 address-family ipv4 unicast route-policy SLCAU-ENTERPRISE-IPV4-IN(AS64595-1346051-IPV4-IN) in
router bgp 38195 neighbor 202.90.207.233 address-family ipv4 unicast route-policy SLCAU-CUST-IPV4-DEFAULTONLY-OUT out
router bgp 38195 neighbor 202.90.207.233 address-family ipv4 unicast default-originate
router bgp 38195 neighbor 202.90.207.233 address-family ipv4 unicast remove-private-AS
router bgp 38195 neighbor 202.90.207.233 address-family ipv4 unicast soft-reconfiguration inbound always

prefix-set AS64595-1346051-IPV4-IN
  27.122.116.160/28 le 32
end-set
!


#cor01-etp-454stpau-bne.au:

interface Bundle-Ether41.1346051 l2transport
interface Bundle-Ether41.1346051 l2transport description Prov: Core: bdr01-cgn-54alfred-bne Be71.4089 (IPT-CGN-NNI-SECONDARY)[200000Mbit]{1346051}
interface Bundle-Ether41.1346051 l2transport encapsulation dot1q 4089
l2vpn xconnect group xc_1346051
l2vpn xconnect group xc_1346051 p2p p2p_1346051
l2vpn xconnect group xc_1346051 p2p p2p_1346051 interface Bundle-Ether41.1346051
l2vpn xconnect group xc_1346051 p2p p2p_1346051 neighbor ipv4 172.21.119.52 pw-id 1346051
l2vpn xconnect group xc_1346051 p2p p2p_1346051 neighbor ipv4 172.21.119.52 pw-id 1346051 pw-class CW-enabled


#agg01-sat-54alfred-bne.au:
interface Bundle-Ether71.1346051 l2transport
interface Bundle-Ether71.1346051 l2transport description Core: bdr01-cgn-54alfred-bne Be71.4089 (IPT-CGN-WALLGARDEN)[200000Mbit]{1346051}
interface Bundle-Ether71.1346051 l2transport encapsulation dot1q 4089
l2vpn xconnect group xc_1346051
l2vpn xconnect group xc_1346051 p2p p2p_1346051
l2vpn xconnect group xc_1346051 p2p p2p_1346051 interface Bundle-Ether71.1346051
l2vpn xconnect group xc_1346051 p2p p2p_1346051 neighbor ipv4 172.21.0.5 pw-id 1346051
l2vpn xconnect group xc_1346051 p2p p2p_1346051 neighbor ipv4 172.21.0.5 pw-id 1346051 pw-class CW-enabled

