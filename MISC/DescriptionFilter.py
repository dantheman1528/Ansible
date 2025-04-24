# \home\daname\Ansible\MISC\DescriptionFilter.py

def extract_interface_descriptions(input_lines):
    descriptions = []
    for line in input_lines:
        # Split the line and extract the description part
        parts = line.split("          ")
        if len(parts) > 3:
            descriptions.append(parts[3].strip())
    return descriptions

if __name__ == "__main__":
    input_data = []
    
    print("Enter interface descriptions (type 'exit' to finish):")
    while True:
        line = input()
        if line.lower() == 'exit':
            break
        input_data.append(line)
    
    # Extract and print descriptions
    descriptions = extract_interface_descriptions(input_data)
    for description in descriptions:
        print(description)