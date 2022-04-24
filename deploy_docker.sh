#!/bin/bash

version=$1 #e.g. v0.1.1

if [[ -z $version ]]; then
  echo "ERROR: specify a version e.g v0.1.1"
  exit 1
fi

echo "Building docker image..."
docker build -t pwdgen:$version --build-arg version=$version .

read -p "Deploy to docker hub? (y/n)" -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]
then
    exit 0
fi


