#!/bin/bash

tag=$1
iteration=$2
version=${tag:1}

echo $tag
echo $iteration
echo $version

sudo apt-get update -y
sudo apt-get install ruby-dev build-essential rpm -y
gem install fpm

echo "Create deb package..."
fpm --verbose -s dir -t deb -n frontier --version $version --iteration $iteration --description "DigitalBits Frontier Server" $GOPATH/bin/frontier=/usr/local/bin/frontier
fpm --verbose -s dir -t deb -n bifrost --version $version --iteration $iteration --description "DigitalBits BiFrost Server" $GOPATH/bin/bifrost=/usr/local/bin/bifrost
fpm --verbose -s dir -t deb -n federation --version $version --iteration $iteration --description "DigitalBits Federation Server" $GOPATH/bin/federation=/usr/local/bin/federation

echo "Create rpm package..."
fpm --verbose -s dir -t rpm -n frontier --version $version --iteration $iteration --description "DigitalBits Frontier Server" $GOPATH/bin/frontier=/usr/local/bin/frontier
fpm --verbose -s dir -t rpm -n bifrost --version $version --iteration $iteration --description "DigitalBits BiFrost Server" $GOPATH/bin/bifrost=/usr/local/bin/bifrost
fpm --verbose -s dir -t rpm -n federation --version $version --iteration $iteration --description "DigitalBits Federation Server" $GOPATH/bin/federation=/usr/local/bin/federation

echo "deploying to Cloudsmith with cloudsmith-cli"

ls
cloudsmith push deb digitalbits/dbtest/ubuntu/trusty frontier-$version-$iteration.amd64.deb
cloudsmith push deb digitalbits/dbtest/ubuntu/trusty bifrost-$version-$iteration.amd64.deb
cloudsmith push deb digitalbits/dbtest/ubuntu/trusty federation-$version-$iteration.amd64.deb

cloudsmith push rpm digitalbits/dbtest/el/7 frontier-$version-$iteration.x86_64.rpm
cloudsmith push rpm digitalbits/dbtest/el/7 bifrost-$version-$iteration.x86_64.rpm
cloudsmith push rpm digitalbits/dbtest/el/7 federation-$version-$iteration.x86_64.rpm
