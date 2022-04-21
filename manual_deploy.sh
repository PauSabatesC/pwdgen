#!/bin/bash

version=$1 #e.g. v0.1.1

if [[ -z $version ]]; then
  echo "ERROR: specify a version e.g v0.1.1"
fi

echo "Did you push the local commit?"
echo "Did you export your github token? export GITHUB_TOKEN=\"YOUR_GH_TOKEN\""
echo "Do nothing if you did all... (ctrl+c to abort)"
sleep 15

git tag -a $version -m "Version $version"
git push origin $version

echo "Tag created and pushed. Doing the release in 10 seconds..."
sleep 10
goreleaser release --rm-dist

