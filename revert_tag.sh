#!/bin/bash

version=$1 #e.g. v0.1.1

if [[ -z $version ]]; then
  echo "ERROR: specify a version e.g v0.1.1"
  exit 1
fi

git tag -d $version
git push --delete origin $version
echo "Done. Tag $version is deleted, you can push it again"

