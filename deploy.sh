#!/bin/bash

sudo apt-get update -y
sudo apt-get install ruby-dev build-essential rpm -y
gem install fpm

mkdir -p $GOPATH/fpm && cd $GOPATH/fpm
echo "Create deb package..."
fpm --verbose -s dir -t deb -n frontier --version 0.0.3 --iteration 1 --description "Digitalbits-frontier" $GOPATH/bin/frontier=/usr/local/bin/frontier
fpm --verbose -s dir -t deb -n bifrost --version 0.0.3 --iteration 1 --description "Digitalbits-bifrost" $GOPATH/bin/bifrost=/usr/local/bin/bifrost
fpm --verbose -s dir -t deb -n federation --version 0.0.3 --iteration 1 --description "Digitalbits-federation" $GOPATH/bin/federation=/usr/local/bin/federation

echo "Create rpm package..."
fpm --verbose -s dir -t rpm -n frontier --version 0.0.3 --iteration 1 --description "Digitalbits-frontier" $GOPATH/bin/frontier=/usr/local/bin/frontier
fpm --verbose -s dir -t rpm -n bifrost --version 0.0.3 --iteration 1 --description "Digitalbits-bifrost" $GOPATH/bin/bifrost=/usr/local/bin/bifrost
fpm --verbose -s dir -t rpm -n federation --version 0.0.3 --iteration 1 --description "Digitalbits-federation" $GOPATH/bin/federation=/usr/local/bin/federation

echo "deploying to Cloudsmith with cloudsmith-cli"

ls
cloudsmith push deb digitalbits/dbtest/ubuntu/trusty frontier.0.3-1_amd64.deb
cloudsmith push deb digitalbits/dbtest/ubuntu/trusty bifrost_0.0.3-1_amd64.deb
cloudsmith push deb digitalbits/dbtest/ubuntu/trusty federation_0.0.3-1_amd64.deb

cloudsmith push rpm digitalbits/dbtest/el/7 frontier-0.0.3-1.x86_64.rpm
cloudsmith push rpm digitalbits/dbtest/el/7 bifrost-0.0.3-1.x86_64.rpm
cloudsmith push rpm digitalbits/dbtest/el/7 federation-0.0.3-1.x86_64.rpm
