#!/bin/bash

version=$1 #e.g. v0.1.1

if [[ -z $version ]]; then
  echo "ERROR: specify a version e.g v0.1.1"
fi

# Remember to push the commit first then:

git tag -a $version -m "Version $version"
git push origin $version

#Remember to export github token: export GITHUB_TOKEN="YOUR_GH_TOKEN"
goreleaser release --rm-dist

