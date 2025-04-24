import paramiko
import time
import getpass
import re

def ssh_connect(hostname, username, password):
    print(f"[DEBUG] Attempting to connect to {hostname} as {username}...")
    ssh = paramiko.SSHClient()
    ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    ssh.connect(hostname, username=username, password=password)
    print(f"[DEBUG] Successfully connected to {hostname}.")
    return ssh

def ssh_command(ssh, command):
    stdin, stdout, stderr = ssh.exec_command(command)
    return stdout.read().decode()

def filter_output(output):
    # Use regex to filter out unwanted lines
    filtered_lines = []
    for line in output.splitlines():
        # Exclude lines that contain the command, timestamps, or "Building configuration..."
        if not re.match(r'^(RP/0/RP0/CPU0:.*|.*UTC|Building configuration...)', line):
            filtered_lines.append(line)
    return "\n".join(filtered_lines)

def main():
    # Prompt for username and password once
    username = input("Enter your username: ")
    password = getpass.getpass("Enter your password: ")

    # Read inventory file
    with open('inventory.txt', 'r') as f:
        devices = f.read().splitlines()  # Read all lines as strings

    # Open file to write output
    with open('output.txt', 'w') as f:
        for device in devices:  # Iterate over all devices
            print(f"\n[DEBUG] Connecting to device: {device}")
            try:
                # Connect to the device
                ssh = ssh_connect(device, username, password)

                # Set terminal length to 0
                terminal_command = "terminal length 0"
                ssh_command(ssh, terminal_command)
                time.sleep(2)  # Wait for the command to take effect

                # Execute command for each item in the inventory
                for item in items:  # Assuming 'items' is defined elsewhere
                    command = f"show run formal | i {item}"
                    output = ssh_command(ssh, command)  # Execute command
                    filtered_output = filter_output(output)  # Filter the output
                    f.write(f"Output for {item} on {device}:\n{filtered_output}\n\n")  # Add a newline for separation

                # Close the SSH connection
                ssh.close()
                print(f"[DEBUG] Disconnected from {device}.")
            except Exception as e:
                print(f"[ERROR] Failed to connect to {device}: {e}")

    print("Command outputs have been saved to output.txt")

if __name__ == "__main__":
    main()