#bin/bash

set -e

printf "Welcome to the Transparency in Coverage app!! Checking your software to make sure you have the dependencies you need to run this application...\n"
sleep 1;

if [[ $(which docker) ]]; then
  echo "Docker is present..."
else
  echo "Docker desktop NOT FOUND. It is required to run this project. You can download it at https://www.docker.com/products/docker-desktop/"
  exit 1;
fi

printf ".env file is ignored in the public repo as a general best practice although there is no sensitive content /nthe contents  copying contents of .example.env are what you need for the .env file so copying the content..."
sleep 1;

cp .example.env .env
echo "content copied to .env successfully!!";
