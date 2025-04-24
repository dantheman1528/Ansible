TE1 = SG2 - P1
    #bdr01-eip-15pioneer
    explicit-path name cor01-tpt-4millros-IndigoWest-01
    explicit-path name cor01-tpt-4millros-IndigoWest-01 index 10 next-address strict ipv4 unicast 172.21.52.81
    explicit-path name cor01-tpt-4millros-IndigoWest-01 index 20 next-address strict ipv4 unicast 172.21.16.13
    explicit-path name cor01-tpt-4millros-IndigoWest-01 index 30 next-address strict ipv4 unicast 172.21.0.41

    interface tunnel-te151
    
    interface tunnel-te151 description TE-Tun: cor01-tpt-4millros-Indigo West-01
    interface tunnel-te151 ipv4 unnumbered Loopback0
    interface tunnel-te151 mpls mtu 17916
    interface tunnel-te151 load-interval 30
    interface tunnel-te151 autoroute destination 172.21.0.41
    interface tunnel-te151 destination 172.21.0.41
    interface tunnel-te151 path-option 10 explicit name cor01-tpt-4millros-IndigoWest-01

    #cor01-tpt-4millros
    explicit-path name bdr01-eip-15pioneer-IndigoWest-1
    explicit-path name bdr01-eip-15pioneer-IndigoWest-1 index 1 next-address strict ipv4 unicast 172.21.16.12
    explicit-path name bdr01-eip-15pioneer-IndigoWest-1 index 2 next-address strict ipv4 unicast 172.21.52.80
    explicit-path name bdr01-eip-15pioneer-IndigoWest-1 index 3 next-address strict ipv4 unicast 202.177.40.4

    interface tunnel-te151
    interface tunnel-te151 description bdr01-eip-15pioneer-IndigoWest-01
    interface tunnel-te151 ipv4 unnumbered Loopback0
    interface tunnel-te151 signalled-name bdr01-eip-15pioneer-IndigoWest-01
    interface tunnel-te151 logging events lsp-status state
    interface tunnel-te151 logging events lsp-status reroute
    interface tunnel-te151 logging events lsp-status switchover
    interface tunnel-te151 logging events lsp-status record-route
    interface tunnel-te151 logging events pcalc-failure
    interface tunnel-te151 logging events lsp-status insufficient-bandwidth
    interface tunnel-te151 priority 6 6
    interface tunnel-te151 destination 202.177.40.4
    interface tunnel-te151 record-route
    interface tunnel-te151 affinity exclude DEDICATED MAINTENANCE
    interface tunnel-te151 path-option 10 explicit name bdr01-eip-15pioneer-IndigoWest-1
    interface tunnel-te151 logging events link-status

    l2vpn pw-class bdr01-eip-15pioneer-IndigoWest-1
    l2vpn pw-class bdr01-eip-15pioneer-IndigoWest-1 encapsulation mpls
    l2vpn pw-class bdr01-eip-15pioneer-IndigoWest-1 encapsulation mpls transport-mode ethernet
    l2vpn pw-class bdr01-eip-15pioneer-IndigoWest-1 encapsulation mpls preferred-path interface tunnel-te 171 fallback disable
TE2 = SG2 - P1
     #bdr02-eip-15pioneer
    explicit-path name cor02-tpt-4millros-IndigoWest-02
    explicit-path name cor02-tpt-4millros-IndigoWest-02 index 10 next-address strict ipv4 unicast 172.21.52.81
    explicit-path name cor02-tpt-4millros-IndigoWest-02 index 20 next-address strict ipv4 unicast 172.21.16.13
    explicit-path name cor02-tpt-4millros-IndigoWest-02 index 30 next-address strict ipv4 unicast 172.21.0.41

    interface tunnel-te152
    
    interface tunnel-te152 description TE-Tun: cor02-tpt-4millros-Indigo West-02
    interface tunnel-te152 ipv4 unnumbered Loopback0
    interface tunnel-te152 mpls mtu 17916
    interface tunnel-te152 load-interval 30
    interface tunnel-te152 autoroute destination 172.21.0.41
    interface tunnel-te152 destination 172.21.0.41
    interface tunnel-te152 path-option 10 explicit name cor02-tpt-4millros-IndigoWest-02

    #cor02-tpt-4millros
    explicit-path name bdr02-eip-15pioneer-IndigoWest-2
    explicit-path name bdr02-eip-15pioneer-IndigoWest-2 index 1 next-address strict ipv4 unicast 172.21.16.12
    explicit-path name bdr02-eip-15pioneer-IndigoWest-2 index 2 next-address strict ipv4 unicast 172.21.52.80
    explicit-path name bdr02-eip-15pioneer-IndigoWest-2 index 3 next-address strict ipv4 unicast 202.177.40.4

    interface tunnel-te152
    interface tunnel-te152 description bdr02-eip-15pioneer-IndigoWest-02
    interface tunnel-te152 ipv4 unnumbered Loopback0
    interface tunnel-te152 signalled-name bdr02-eip-15pioneer-IndigoWest-02
    interface tunnel-te152 logging events lsp-status state
    interface tunnel-te152 logging events lsp-status reroute
    interface tunnel-te152 logging events lsp-status switchover
    interface tunnel-te152 logging events lsp-status record-route
    interface tunnel-te152 logging events pcalc-failure
    interface tunnel-te152 logging events lsp-status insufficient-bandwidth
    interface tunnel-te152 priority 6 6
    interface tunnel-te152 destination 202.177.40.4
    interface tunnel-te152 record-route
    interface tunnel-te152 affinity exclude DEDICATED MAINTENANCE
    interface tunnel-te152 path-option 10 explicit name bdr02-eip-15pioneer-IndigoWest-2
    interface tunnel-te152 logging events link-status

    l2vpn pw-class bdr02-eip-15pioneer-IndigoWest-2
    l2vpn pw-class bdr02-eip-15pioneer-IndigoWest-2 encapsulation mpls
    l2vpn pw-class bdr02-eip-15pioneer-IndigoWest-2 encapsulation mpls transport-mode ethernet
    l2vpn pw-class bdr02-eip-15pioneer-IndigoWest-2 encapsulation mpls preferred-path interface tunnel-te 152 fallback disable
TE3 = SG2 - S1
#bdr01-eip-15pioneer
    explicit-path name cor01-etp-4edenpar-syd.au-IndigoCentral-1
    explicit-path name cor01-etp-4edenpar-syd.au-IndigoCentral-1 index 10 next-address strict ipv4 unicast 172.21.52.81
    explicit-path name cor01-etp-4edenpar-syd.au-IndigoCentral-1 index 20 next-address strict ipv4 unicast 172.21.16.13
    explicit-path name cor01-etp-4edenpar-syd.au-IndigoCentral-1 index 30 next-address strict ipv4 unicast 172.21.6.3
    explicit-path name cor01-etp-4edenpar-syd.au-IndigoCentral-1 index 40 next-address strict ipv4 unicast 172.21.70.128
    explicit-path name cor01-etp-4edenpar-syd.au-IndigoCentral-1 index 50 next-address strict ipv4 unicast 172.21.0.13

    interface tunnel-te171
    
    interface tunnel-te171 description TE-Tun: cor01-etp-4edenpar-syd.au-IndigoCentral-1
    interface tunnel-te171 ipv4 unnumbered Loopback0
    interface tunnel-te171 mpls mtu 17916
    interface tunnel-te171 load-interval 30
    interface tunnel-te171 autoroute destination 172.21.0.13
    interface tunnel-te171 destination 172.21.0.13
    interface tunnel-te171 path-option 10 explicit name cor01-etp-4edenpar-syd.au-IndigoCentral-1

    #cor01-etp-4edenpar-syd
    explicit-path name bdr01-eip-15pioneer-IndigoCentral-1
    explicit-path name bdr01-eip-15pioneer-IndigoCentral-1 index 1 next-address strict ipv4 unicast 172.21.70.129
    explicit-path name bdr01-eip-15pioneer-IndigoCentral-1 index 2 next-address strict ipv4 unicast 172.21.6.2
    explicit-path name bdr01-eip-15pioneer-IndigoCentral-1 index 3 next-address strict ipv4 unicast 172.21.16.12
    explicit-path name bdr01-eip-15pioneer-IndigoCentral-1 index 4 next-address strict ipv4 unicast 172.21.52.80
    explicit-path name bdr01-eip-15pioneer-IndigoCentral-1 index 5 next-address strict ipv4 unicast 202.177.40.4

    interface tunnel-te171
    interface tunnel-te171 description bdr01-eip-15pioneer-IndigoCentral-1
    interface tunnel-te171 ipv4 unnumbered Loopback0
    interface tunnel-te171 signalled-name bdr01-eip-15pioneer-IndigoCentral-1
    interface tunnel-te171 logging events lsp-status state
    interface tunnel-te171 logging events lsp-status reroute
    interface tunnel-te171 logging events lsp-status switchover
    interface tunnel-te171 logging events lsp-status record-route
    interface tunnel-te171 logging events pcalc-failure
    interface tunnel-te171 logging events lsp-status insufficient-bandwidth
    interface tunnel-te171 priority 6 6
    interface tunnel-te171 destination 202.177.40.4
    interface tunnel-te171 record-route
    interface tunnel-te171 affinity exclude DEDICATED MAINTENANCE
    interface tunnel-te171 path-option 10 explicit name bdr01-eip-15pioneer-IndigoCentral-1
    interface tunnel-te171 logging events link-status

    l2vpn pw-class bdr01-eip-15pioneer-IndigoCentral-1
    l2vpn pw-class bdr01-eip-15pioneer-IndigoCentral-1 encapsulation mpls
    l2vpn pw-class bdr01-eip-15pioneer-IndigoCentral-1 encapsulation mpls transport-mode ethernet
    l2vpn pw-class bdr01-eip-15pioneer-IndigoCentral-1 encapsulation mpls preferred-path interface tunnel-te 171 fallback disable
TE4 = SG2 - S1
#bdr01-eip-15pioneer
    explicit-path name cor01-etp-4edenpar-syd.au-IndigoCentral-2
    explicit-path name cor01-etp-4edenpar-syd.au-IndigoCentral-2 index 10 next-address strict ipv4 unicast 172.21.52.81
    explicit-path name cor01-etp-4edenpar-syd.au-IndigoCentral-2 index 20 next-address strict ipv4 unicast 172.21.16.13
    explicit-path name cor01-etp-4edenpar-syd.au-IndigoCentral-2 index 30 next-address strict ipv4 unicast 172.21.6.3
    explicit-path name cor01-etp-4edenpar-syd.au-IndigoCentral-2 index 40 next-address strict ipv4 unicast 172.21.70.130
    explicit-path name cor01-etp-4edenpar-syd.au-IndigoCentral-2 index 50 next-address strict ipv4 unicast 172.21.0.13

    interface tunnel-te172
    
    interface tunnel-te172 description TE-Tun: cor01-etp-4edenpar-syd.au-IndigoCentral-2
    interface tunnel-te172 ipv4 unnumbered Loopback0
    interface tunnel-te172 mpls mtu 17916
    interface tunnel-te172 load-interval 30
    interface tunnel-te172 autoroute destination 172.21.0.13
    interface tunnel-te172 destination 172.21.0.13
    interface tunnel-te172 path-option 10 explicit name cor01-etp-4edenpar-syd.au-IndigoCentral-2

    #cor01-etp-4edenpar-syd
    explicit-path name bdr01-eip-15pioneer-IndigoCentral-2
    explicit-path name bdr01-eip-15pioneer-IndigoCentral-2 index 1 next-address strict ipv4 unicast 172.21.70.131
    explicit-path name bdr01-eip-15pioneer-IndigoCentral-2 index 2 next-address strict ipv4 unicast 172.21.6.2
    explicit-path name bdr01-eip-15pioneer-IndigoCentral-2 index 3 next-address strict ipv4 unicast 172.21.16.12
    explicit-path name bdr01-eip-15pioneer-IndigoCentral-2 index 4 next-address strict ipv4 unicast 172.21.52.80
    explicit-path name bdr01-eip-15pioneer-IndigoCentral-2 index 5 next-address strict ipv4 unicast 202.177.40.4

    interface tunnel-te172
    interface tunnel-te172 description bdr01-eip-15pioneer-IndigoCentral-2
    interface tunnel-te172 ipv4 unnumbered Loopback0
    interface tunnel-te172 signalled-name bdr01-eip-15pioneer-IndigoCentral-2
    interface tunnel-te172 logging events lsp-status state
    interface tunnel-te172 logging events lsp-status reroute
    interface tunnel-te172 logging events lsp-status switchover
    interface tunnel-te172 logging events lsp-status record-route
    interface tunnel-te172 logging events pcalc-failure
    interface tunnel-te172 logging events lsp-status insufficient-bandwidth
    interface tunnel-te172 priority 6 6
    interface tunnel-te172 destination 202.177.40.4
    interface tunnel-te172 record-route
    interface tunnel-te172 affinity exclude DEDICATED MAINTENANCE
    interface tunnel-te172 path-option 10 explicit name bdr01-eip-15pioneer-IndigoCentral-2
    interface tunnel-te172 logging events link-status

    l2vpn pw-class bdr01-eip-15pioneer-IndigoCentral-2
    l2vpn pw-class bdr01-eip-15pioneer-IndigoCentral-2 encapsulation mpls
    l2vpn pw-class bdr01-eip-15pioneer-IndigoCentral-2 encapsulation mpls transport-mode ethernet
    l2vpn pw-class bdr01-eip-15pioneer-IndigoCentral-2 encapsulation mpls preferred-path interface tunnel-te 172 fallback disable
TE5 = SG3 - PIX
TE6 = SG3 - PIX
TE7 = SG3 - SY3
TE8 = SG3 - SY3


