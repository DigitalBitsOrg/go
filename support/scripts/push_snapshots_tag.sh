#!/bin/bash
#this script will delete and recreate the snapshots tag, unless it's a tagged release

if [ "$TRAVIS_TAG" != "" ]
then
    echo "tagged release - doing nothing"
else
    echo "snapshots release - overwriting snapshots tag"
    git config --global user.email "devops@digitalbits.io"
    git config --global user.name "Travis CI"
    git tag -d snapshots || true
    git push --delete -q https://$GITHUB_OAUTH_TOKEN@github.com/digitalbitsorg/go snapshots || true
    git tag snapshots -a -m "Generated snapshots from TravisCI for build $TRAVIS_BUILD_NUMBER on $(date --utc +'%F-%T')"
    git push -q https://$GITHUB_OAUTH_TOKEN@github.com/digitalbitsorg/go --tags
fi
