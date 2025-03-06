generate_env_example() {
    local env_file=$1
    local env_example_file=$2

    # Check if the .env file exists
    if [[ ! -f "$env_file" ]]; then
        echo "The .env file does not exist at the specified path: $env_file"
        return 1
    fi

    # Create a temporary file to store the new .env.example content
    local temp_file=$(mktemp)

    # Read the .env file line by line
    while IFS= read -r line; do
        # Skip empty lines and comments
        if [[ -z "$line" || "$line" =~ ^# ]]; then
            continue
        fi

        # Extract the key from the line
        key=$(echo "$line" | cut -d '=' -f 1)

        # Check if the key already exists in the .env.example file
        if grep -q "^$key=" "$env_example_file"; then
            # If the key exists, keep the existing line
            grep "^$key=" "$env_example_file" >> "$temp_file"
        else
            # If the key does not exist, add it with an empty value
            echo "$key=" >> "$temp_file"
        fi
    done < "$env_file"

    # Append any keys from the existing .env.example file that are not in the .env file
    while IFS= read -r line; do
        # Skip empty lines and comments
        if [[ -z "$line" || "$line" =~ ^# ]]; then
            continue
        fi

        # Extract the key from the line
        key=$(echo "$line" | cut -d '=' -f 1)

        # Check if the key is already in the temp file
        if ! grep -q "^$key=" "$temp_file"; then
            echo "$line" >> "$temp_file"
        fi
    done < "$env_example_file"

    # Replace the .env.example file with the new content
    mv "$temp_file" "$env_example_file"

    echo ".env.example file has been generated/updated at: $env_example_file"
}

# Usage example:
# generate_env_example /path/to/.env /path/to/.env.example

# Function to generate .env.example in the backend folder
generate_backend_env_example() {
    local backend_folder="./backend"
    local env_file="$backend_folder/.env"
    local env_example_file="$backend_folder/.env.example"

    generate_env_example "$env_file" "$env_example_file"
}
generate_frontend_env_example() {
    local frontend_folder="./frontend"
    local env_file="$frontend_folder/.env"
    local env_example_file="$frontend_folder/.env.example"

    generate_env_example "$env_file" "$env_example_file"
}

generate_devops_env_example() {
    local devops_folder="./devops"
    local env_file="$devops_folder/.env"
    local env_example_file="$devops_folder/.env.example"

    generate_env_example "$env_file" "$env_example_file"
}

generate_root_dir_env_example() {

    generate_env_example ./.env ./.env.example
}
generate_all_env_examples() {
    generate_root_dir_env_example
    generate_backend_env_example
    generate_frontend_env_example
    generate_devops_env_example
}
