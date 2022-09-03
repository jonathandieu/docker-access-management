#! /usr/bin/bash


# find . | grep versions.tf | xargs -L1 sed -i 's/"~> 1.0.7"/"~> 1.1.9"/

# find . -type f | grep 'versions.tf' | grep -v '.terraform' | xargs grep 1\.1\.9   # lists versions.tf files -> with 1.1.9 version

# Find all the directories with "main.tf" files in them
echo "Running script..."

echo "What is the OLD version of terraform you want to replace?"
read OLD_VER
OLD_VER_STR="\">= $OLD_VER\""

echo "What is the NEW version of terraform you want to replace the OLD one with?"
read NEW_VER
NEW_VER_STR="\">= $NEW_VER\""

# echo "Replace $OLD_VER_STR with $NEW_VER_STR ?"
echo replacing Terraform v$OLD_VER with Terraform v$NEW_VER

# Replace old version with new one
# find . | grep versions.tf | xargs -L1 sed -i 's/"~> 1.0.0"/"=> 1.1.9"/'
find . | grep versions.tf | xargs -L1 sed -i "s/$OLD_VER_STR/$NEW_VER_STR/"

# TERRAFORM_DIRECTORIES=$(find . -name "main.tf" | cut -c 2- | rev | cut -c 9- | rev)
TERRAFORM_DIRECTORIES=$(rg required_version | grep 1.1.9 | rev | cut -c45- | rev )
# echo $TERRAFORM_DIRECTORIES

echo Running git status:
git status

################

    BASE_PATH=$(pwd)
    echo "Base Path: $BASE_PATH"
    echo $TERRAFORM_DIRECTORIES

    for d in $TERRAFORM_DIRECTORIES; do
        cd $BASE_PATH/$d
        pwd
        terraform init -upgrade=true


# Run terraform init in all changed directories
echo "Would you like to run terraform init -upgrade=true in all changed directories?"
# echo "Enter 'y' for yes or 'n' for no."
read -p "Enter 'y' for yes or 'n' for no: " TO_INIT
if ["$TO_INIT" == "y"];
then
    BASE_PATH=$(pwd)
    echo "Base Path: $BASE_PATH"
    echo $TERRAFORM_DIRECTORIES

    for d in $TERRAFORM_DIRECTORIES; do
        cd $BASE_PATH/$d
        pwd
        terraform init -upgrade=true

    done
else
    echo "OK."

fi