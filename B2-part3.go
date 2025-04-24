1. start at 11:45pm brisbane time
2. mention CHG in channel
3. turn off the box alerting in obs
4. do prechecks:
	-save config of bdr01-ipt BE2
	-show bgp all all summ wide >>check which live services running BGP
	-show run formal interface | i address >>check live IP address to which interface
5. migrate first the 5 logical services
	- remove ip address and shutdown bgp on bdr01
6. migrate the 1 physical service - Nexus
	- remove ip address and shutdown bgp on bdr01
7. check the l2circuit is up on bdr03
8. check the BGP is up on bdr03
9. Once all services migrated, Decom the ISIS adjacency
	-make sure to login via console:
		ssh oob01-mgt-454stpau-bne.au.superloop.net.co -p 3000


		router isis core 
		no  interface HundredGigE0/0/1/2
		no interface Bundle-Ether33
		no interface HundredGigE0/0/1/2.402
10. Decom the iBGP
router bgp 38195
neighbor 103.200.13.7 shutdown
neighbor 103.200.13.8 shutdown
neighbor 103.200.13.32 shutdown
neighbor 192.168.255.11 shutdown
neighbor 192.168.255.12 shutdown
neighbor 192.168.255.13 shutdown
neighbor 208.76.14.223 shutdown
neighbor 2401:d000:7100::1 shutdown
neighbor 2401:d000:7100::3 shutdown
neighbor 2401:d000:7200::3  shutdown
neighbor 2620:129:1:2::1 shutdown
neighbor fdff:dd:ddaa:1::1 shutdown
neighbor fdff:dd:ddaa:2::1 shutdown
neighbor fdff:dd:ddaa:3::1 shutdown
11. shutdown and Decom BE2 and its physical ports on bdr01
12. shutdown and Decom BE33 and its physical ports on bdr01
13. shutdown and Decom the 0/0/1/2 on bdr01
14. remove all patches and de-rack the device
15. ensure traffic and BGP is returned to normal for each customer:
https://obs-carrier.au.superloop.systems/device/device=400/tab=port/port=31144/
https://obs-carrier.au.superloop.systems/device/device=400/tab=port/port=33183/
https://obs-carrier.au.superloop.systems/device/device=400/tab=port/port=57634/
https://obs-carrier.au.superloop.systems/device/device=400/tab=port/port=20735/
https://obs-carrier.au.superloop.systems/device/device=400/tab=port/port=167288/
https://obs-carrier.au.superloop.systems/device/device=400/tab=port/port=20724/

16. update accounting


for 18621
	on bdr01-met-54alfred-bne.au

		delete protocols l2circuit neighbor 103.200.12.24 interface ge-0/0/12.1

		set protocols l2circuit neighbor 172.21.119.11 interface ge-0/0/12.1 apply-groups L2CIRCUIT-TAGGED
		set protocols l2circuit neighbor 172.21.119.11 interface ge-0/0/12.1 virtual-circuit-id 18621
		set protocols l2circuit neighbor 172.21.119.11 interface ge-0/0/12.1 description SLC-IPT-18621
		set protocols l2circuit neighbor 172.21.119.11 interface ge-0/0/12.1 community evc-core-protected
		set protocols l2circuit neighbor 172.21.119.11 interface ge-0/0/12.1 ignore-mtu-mismatch
		set protocols l2circuit neighbor 172.21.119.11 interface ge-0/0/12.1 pseudowire-status-tlv
		set protocols l2circuit neighbor 172.21.119.11 interface ge-0/0/12.1 encapsulation-type ethernet


for 21278
	on bdr01-met-454stpau-bne.au:

		delete interfaces xe-0/0/42 unit 15999
		delete interfaces ae2 unit 15999
		delete vlans S21278

		set interfaces xe-0/0/42 apply-groups INTERFACE-PARAMETERS-TAGGED
		set interfaces xe-0/0/42 flexible-vlan-tagging
		set interfaces xe-0/0/42 
		set interfaces xe-0/0/42 mtu 9216
		set interfaces xe-0/0/42 encapsulation flexible-ethernet-services
		set interfaces xe-0/0/42 unit 0 description "Cust: Professional Data Kinetics Pty Ltd - SLAUIN-IPTP [10000Mbit]{21278}"
		set interfaces xe-0/0/42 unit 0 encapsulation vlan-ccc 
		set interfaces xe-0/0/42 unit 0 family ccc 
	
		set protocols l2circuit neighbor 172.21.119.11 interface xe-0/0/42.0 virtual-circuit-id 21278
		set protocols l2circuit neighbor 172.21.119.11 interface xe-0/0/42.0 community evc-core-protected
		set protocols l2circuit neighbor 172.21.119.11 interface xe-0/0/42.0 encapsulation-type ethernet
		set protocols l2circuit neighbor 172.21.119.11 interface xe-0/0/42.0 ignore-mtu-mismatch
		set protocols l2circuit neighbor 172.21.119.11 interface xe-0/0/42.0 pseudowire-status-tlv


		27.122.113.164                            0      134143    6676297  202382284 16457416433          0          0       1w6d         59
		2401:d000:7201::32                        0      134143    1433564   33613879 24280521811          0          0       1w6d          1
		


		interface Bundle-Ether2.21278 l2transport
		interface Bundle-Ether2.21278 l2transport description Cust: Professional Data Kinetics Pty Ltd - SLAUIN-IPTP [10000Mbit]{21278}
		interface Bundle-Ether2.21278 l2transport encapsulation dot1q 112
		interface Bundle-Ether2.21278 l2transport rewrite ingress tag pop 1 symmetric
		interface Bundle-Ether2.21278 l2transport mtu 9216
		interface TenGigE0/0/0/27 description Cust: Professional Data Kinetics Pty Ltd - SLAUIN-IPTP [10000Mbit]{21278}
		l2vpn xconnect group xc_21278
		l2vpn xconnect group xc_21278 p2p p2p_21278
		l2vpn xconnect group xc_21278 p2p p2p_21278 interface TenGigE0/0/0/27
		l2vpn xconnect group xc_21278 p2p p2p_21278 interface Bundle-Ether2.21278

		interface TenGig0/0/0/27.21278 l2transport encapsulation untagged 
interface TenGig0/0/0/27.21278 l2transport rewrite ingress tag push dot1q 112 symmetric 
		
		
		interface TenGigE0/0/0/27 mtu 9216
		interface TenGigE0/0/0/27 load-interval 30
		interface TenGigE0/0/0/27 l2transport
		



for 34164
	on bdr02-met-454stpau-bne.au:

		delete protocols l2circuit neighbor 103.200.12.24 interface xe-0/0/5.16383

		set protocols l2circuit neighbor 172.21.119.11 interface xe-0/0/5.16383 virtual-circuit-id 34164
		set protocols l2circuit neighbor 172.21.119.11 interface xe-0/0/5.16383 description "Prov: Cust Alt VFX Pty Ltd - SLAUIN-IPTP [800Mbit]{34164}"
		set protocols l2circuit neighbor 172.21.119.11 interface xe-0/0/5.16383 community evc-core-protected
		set protocols l2circuit neighbor 172.21.119.11 interface xe-0/0/5.16383 encapsulation-type ethernet
		set protocols l2circuit neighbor 172.21.119.11 interface xe-0/0/5.16383 ignore-mtu-mismatch
		set protocols l2circuit neighbor 172.21.119.11 interface xe-0/0/5.16383 pseudowire-status-tlv


		27.122.120.2                              0       65153          0          0           0          0          0   00:00:00 Idle
		203.153.16.93                             0       65153    6385442    5560833 16457438336          0          0       1w6d          2
		2401:d000:7201::92                        0       65153          0          0           0          0          0   00:00:00 Active
		


for 39985
	on bdr01-met-54alfred-bne.au:

		delete  protocols l2circuit neighbor 172.21.119.2 interface xe-0/0/4.0

		set protocols l2circuit neighbor 172.21.119.11 interface xe-0/0/4.0 virtual-circuit-id 39985
		set protocols l2circuit neighbor 172.21.119.11 interface xe-0/0/4.0 description "Prov: Cust Macquarie Telecom Pty Ltd - SLAUIN-IPTP [8000Mbit]{39985}"
		set protocols l2circuit neighbor 172.21.119.11 interface xe-0/0/4.0 community evc-core-protected
		set protocols l2circuit neighbor 172.21.119.11 interface xe-0/0/4.0 encapsulation-type ethernet
		set protocols l2circuit neighbor 172.21.119.11 interface xe-0/0/4.0 ignore-mtu-mismatch
		set protocols l2circuit neighbor 172.21.119.11 interface xe-0/0/4.0 pseudowire-status-tlv

		27.122.114.230                            0       17477    6122261  201843204 16457442038          0          0       1w6d          6
		2401:d000:7201::c2                        0       17477    6122133  146529982 24280533511          0          0       1w6d          1
		


for 480157
	on bdr02-met-454stpau-bne.au:

		delete protocols l2circuit neighbor 103.200.12.24 interface xe-0/0/47.16271

		set protocols l2circuit neighbor 172.21.119.11 interface xe-0/0/47.16271 virtual-circuit-id 480157
		set protocols l2circuit neighbor 172.21.119.11 interface xe-0/0/47.16271 community evc-core-protected
		set protocols l2circuit neighbor 172.21.119.11 interface xe-0/0/47.16271 encapsulation-type ethernet
		set protocols l2circuit neighbor 172.21.119.11 interface xe-0/0/47.16271 ignore-mtu-mismatch
		set protocols l2circuit neighbor 172.21.119.11 interface xe-0/0/47.16271 pseudowire-status-tlv


		27.122.113.89                             0      131268    3739173    3559749 16457450624          0          0       1w6d          1
		2401:d000:7201::222                       0      131268          0          0           0          0          0   00:00:00 Active
		

for 225134
	on bdr02-etp-454stpau-bne.au:

		interface Bundle-Ether2.225134 l2transport
		interface Bundle-Ether2.225134 l2transport description  Cust Nexus One Pty Ltd - SLAUIN-IPTP (AS64006)[500Mbit]{225134}
		interface Bundle-Ether2.225134 l2transport encapsulation dot1q 304
		interface Bundle-Ether2.225134 l2transport rewrite ingress tag pop 1 symmetric
		interface Bundle-Ether2.225134 l2transport mtu 9216
		interface Te0/0/0/12 description  Cust Nexus One Pty Ltd - SLAUIN-IPTP (AS64006)[500Mbit]{225134}
		interface Te0/0/0/12 mtu 9216
		interface Te0/0/0/12 load-interval 30
		interface Te0/0/0/12 l2transport 
		l2vpn xconnect group xc_225134
		l2vpn xconnect group xc_225134 p2p p2p_225134
		l2vpn xconnect group xc_225134 p2p p2p_225134 interface Te0/0/0/12
		l2vpn xconnect group xc_225134 p2p p2p_225134 interface Bundle-Ether2.225134

	on bdr03-ipt-454stpau-bne.au:

		interface BE2.304 description  Cust Nexus One Pty Ltd - SLAUIN-IPTP (AS64006)[500Mbit]{225134}
		interface BE2.304 service-policy input AS38195-DSCP-REWRITE-IN
		interface BE2.304 ipv4 mtu 1500
		interface BE2.304 ipv4 address 203.153.16.201 255.255.255.248
		interface BE2.304 ipv4 unreachables disable
		interface BE2.304 ipv6 nd suppress-ra
		interface BE2.304 ipv6 mtu 1500
		interface BE2.304 ipv6 address 2401:d000:7201::21/124
		interface BE2.304 ipv6 unreachables disable
		interface BE2.304 load-interval 30
		interface BE2.304 flow ipv4 monitor netflow-ipv4 sampler 1in1k ingress
		interface BE2.304 encapsulation dot1q 304
		interface BE2.304 flow ipv6 monitor netflow-ipv6 sampler 1in1k ingress


		router bgp 38195 neighbor 2401:d000:7201::22
		no shutdown

		router bgp 38195 neighbor 2401:d000:7201::23
		no shutdown

		router bgp 38195 neighbor 203.153.16.205
		no shutdown

		router bgp 38195 neighbor 203.153.16.206
		no shutdown

		router bgp 38195 neighbor 27.122.114.230 
		no shutdown

		router bgp 38195 neighbor 27.122.113.89 
		no shutdown

		router bgp 38195 neighbor 203.153.16.93
		no shutdown


		203.153.16.205                            0       64006    5844443  433908043 16457451458          0          0       2w4d          2
		203.153.16.206                            0       64006    5843605  433401219 16457450001          0          0       2w4d          2
		2401:d000:7201::22                        0       64006    5837223  227296002 24280539604          0          0       2w4d          1
		2401:d000:7201::23                        0       64006    5837491  227387788 24280539604          0          0       2w4d          1
		





	on bdr03-ipt-454stpau-bne.au:

		router static address-family ipv4 unicast 27.122.112.0/20 Null0 description BGP_NAILDOWN_SUMMARY_AS38195
		router static address-family ipv4 unicast 27.122.112.0/24 Null0 description BGP_NAILDOWN_AS38195
		router static address-family ipv4 unicast 27.122.113.0/24 Null0 description BGP_NAILDOWN_AS38195
		router static address-family ipv4 unicast 27.122.114.0/24 Null0 description BGP_NAILDOWN_AS38195
		router static address-family ipv4 unicast 27.122.116.0/24 Null0 description BGP_NAILDOWN_AS38195
		router static address-family ipv4 unicast 103.200.12.0/24 Null0 description BGP_NAILDOWN_AS38195
		router static address-family ipv4 unicast 125.253.111.0/24 Null0 description BGP_NAILDOWN_SUMMARY_PEX
		router static address-family ipv4 unicast 192.0.2.1/32 Null0
		router static address-family ipv4 unicast 202.130.197.0/24 Null0 description BGP_NAILDOWN_SUMMARY_PEX
		router static address-family ipv4 unicast 202.137.162.0/24 Null0 description BGP_NAILDOWN_SUMMARY_PEX
		router static address-family ipv4 unicast 202.171.174.0/24 Null0 description BGP_NAILDOWN_SUMMARY_PEX
		router static address-family ipv4 unicast 203.153.16.0/24 Null0 description BGP_NAILDOWN_AS38195


	
#Decom 1G PEX
		router bgp 38195 neighbor 103.200.14.245
		shutdown

		router bgp 38195 neighbor 2401:d000:7201::82
		shutdown











migration part 4! >.<

action plan:
migrate directly the customer to port 0/0/18/0 of bdr03, 
if fails;
	apply mtu 
	keep patching on met then haul to bdr01 B1
	repatch in bdr02-etp then haul to bdr01 B1

		interface Te0/0/0/18/0 mtu 9212
		interface Te0/0/0/18/0 description Cust: Professional Data Kinetics - SLAUIN-IPTP (AS134143)[10000Mbps]{21278}
		interface Te0/0/0/18/0 service-policy input AS38195-DSCP-REWRITE-IN
		interface Te0/0/0/18/0 ipv4 mtu 1500
		interface Te0/0/0/18/0 ipv4 unreachables disable
		interface Te0/0/0/18/0 ipv4 address 27.122.113.165 255.255.255.254
		interface Te0/0/0/18/0 ipv4 unreachables disable
		interface Te0/0/0/18/0 ipv6 mtu 1500
		interface Te0/0/0/18/0 ipv6 address 2401:d000:7201::31/124
		interface Te0/0/0/18/0 ipv6 nd suppress-ra
		interface Te0/0/0/18/0 ipv6 unreachables disable
		interface Te0/0/0/18/0 load-interval 30
		interface Te0/0/0/18/0 flow ipv4 monitor netflow-ipv4 sampler 1in1k ingress
		interface Te0/0/0/18/0 flow ipv6 monitor netflow-ipv6 sampler 1in1k ingress



		router static address-family ipv4 unicast 27.122.112.0/20 Null0 description BGP_NAILDOWN_SUMMARY_AS38195
		router static address-family ipv4 unicast 27.122.112.0/24 Null0 description BGP_NAILDOWN_AS38195
		router static address-family ipv4 unicast 27.122.113.0/24 Null0 description BGP_NAILDOWN_AS38195
		router static address-family ipv4 unicast 27.122.114.0/24 Null0 description BGP_NAILDOWN_AS38195
		router static address-family ipv4 unicast 27.122.116.0/24 Null0 description BGP_NAILDOWN_AS38195
		router static address-family ipv4 unicast 103.200.12.0/24 Null0 description BGP_NAILDOWN_AS38195
		router static address-family ipv4 unicast 125.253.111.0/24 Null0 description BGP_NAILDOWN_SUMMARY_PEX
		router static address-family ipv4 unicast 192.0.2.1/32 Null0
		router static address-family ipv4 unicast 202.130.197.0/24 Null0 description BGP_NAILDOWN_SUMMARY_PEX
		router static address-family ipv4 unicast 202.137.162.0/24 Null0 description BGP_NAILDOWN_SUMMARY_PEX
		router static address-family ipv4 unicast 202.171.174.0/24 Null0 description BGP_NAILDOWN_SUMMARY_PEX
		router static address-family ipv4 unicast 203.153.16.0/24 Null0 description BGP_NAILDOWN_AS38195
		router static address-family ipv6 unicast fc00::192:0:2:1/128 Null0 description RTBH_DEST
		router static vrf MGMT
		router static vrf MGMT address-family ipv4 unicast 0.0.0.0/0 10.61.22.1
		router static vrf MGMT address-family ipv4 unicast 0.0.0.0/0 10.205.104.97
		

		