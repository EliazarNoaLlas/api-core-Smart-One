#!/bin/bash

# Get the current branch name
current_branch=$CI_COMMIT_BRANCH

# Verify if the current branch is 'develop' or 'master'
#if [ "$current_branch" != "develop" ] && [ "$current_branch" != "master" ]; then
#    echo "Error: This script should only be run on 'develop' or 'master' branches. branch: $current_branch"
#    exit 0
#fi

# Get the latest tag from the repository and clean newline and whitespace characters
latest_tag=$(git describe --tags $(git rev-list --tags --max-count=1) | sed 's/v//g')


# Parse version components from the latest tag
IFS='.' read -r -a version_components <<< "$latest_tag"
major=${version_components[0]}
minor=${version_components[1]}
patch=${version_components[2]}

# Function to increment version number based on the branch
increment_version() {
    local major=$1
    local minor=$2
    local patch=$3
    local branch=$4
    local new_minor

    if [ "$branch" == "develop" ]; then
        next_version="v$major.$minor.$((patch + 1))"
    elif [ "$branch" == "master" ]; then
        next_version="v$major.$((minor + 1)).0"
    else
        # Unknown branch, do not increment version
        next_version="v$major.$minor.$patch"
    fi

    # Print the new version
    echo "$next_version"
}

# If there is no previous tag, set initial version as v1.0.0
if [ -z "$latest_tag" ]; then
    next_version="v1.0.0"
else
    # Calculate the new version according to the specified rules
    next_version=$(increment_version "$major" "$minor" "$patch" "$current_branch")
fi

# Print the next tag
echo "$next_version"
