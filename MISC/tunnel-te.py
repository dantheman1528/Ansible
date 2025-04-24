def get_input(prompt):
    return input(prompt).strip()

def get_device_details(device_label):
    print(f"\nEnter details for {device_label}:")
    far_end_device = get_input("Enter the far-end device name: ")
    cable_system = get_input("Enter the cable system name: ")
    identifier = get_input("Enter the identifier: ")
    tunnel_num = get_input("Enter the tunnel number: ")
    far_end_loopback0 = get_input("Enter the far-end Loopback0 IP: ")
    explicit_path_name = f"{far_end_device}-{cable_system}-{identifier}"

    # Get the number of hops for the explicit path
    num_hops = int(get_input("Enter the number of hops for the explicit path: "))

    # Generate the explicit path
    explicit_path = f"\n{device_label}:\nexplicit-path name {explicit_path_name}\n"
    for i in range(1, num_hops + 1):
        next_address = get_input(f"Enter next-address for index {i}: ")
        explicit_path += (
            f"explicit-path name {explicit_path_name} index {i * 10} next-address strict ipv4 unicast {next_address}\n"
        )

    # Generate tunnel config
    tunnel_config = f"""{explicit_path}
interface tunnel-te{tunnel_num}
 description TE-Tun: {explicit_path_name}
 ipv4 unnumbered Loopback0
 mpls mtu 17916
 load-interval 30
 autoroute destination {far_end_loopback0}
 destination {far_end_loopback0}
 path-option 10 explicit name {explicit_path_name}
"""

    return tunnel_config

def generate_template():
    config_a = get_device_details("Device A")
    config_b = get_device_details("Device B")

    # Print the generated configurations
    print("\n--- Device A Configuration ---")
    print(config_a)
    print("--- Device B Configuration ---")
    print(config_b)

if __name__ == "__main__":
    generate_template()
