#!/bin/bash

set -ueo pipefail

# Define services and their directories
declare -A services
services=( ["service-a"]="service-a" ["service-b"]="service-b" ["service-c"]="service-c" )

# Check for changes and tag accordingly
for service in "${!services[@]}"; do
    dir=${services[$service]}
    if [[ $(git diff --name-only HEAD HEAD~1 2>/dev/null) =~ $dir/ ]]; then
        # Extract the current version from the latest tag for this service
        current_version=$(git tag --list "${service}-v*" | sort -V | tail -1)
        if [[ -z "$current_version" ]]; then
            new_version="v0.1.0"
        else
            new_version=$(echo "$current_version" | awk -F. -v OFS=. '{$NF += 1; print}')
        fi
        # git tag "${service}-${new_version}"
        echo "Tagged ${service} with ${service}-${new_version}"
    fi
done
