#!/bin/bash

apt-get update -y
apt-get install ruby-dev build-essential -y
gem install fpm

echo "Create deb package..."
fpm --verbose -s dir -t deb -n bifrost --version 0.1.0 --iteration 1 --description "Digitalbits-bifrost" $GOPATH/bin/bifrost=/usr/local/bin

echo "Create rpm package..."
fpm --verbose -s dir -t rpm -n bifrost --version 0.1.0 --iteration 1 --description "Digitalbits-bifrost" $GOPATH/bin/bifrost=/usr/local/bin

echo "deploying to Cloudsmith with cloudsmith-cli"

ls
cloudsmith push deb digitalbits/dbtest/ubuntu/trusty bifrost_0.1.0-1_amd64.deb
cloudsmith push rpm digitalbits/dbtest/el/7 bifrost-0.1.0-1.x86_64.rpm 
