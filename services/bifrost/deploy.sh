#!/bin/bash

mkdir -p /tmp/build
make DESTDIR=/tmp/build install


apt-get update -y 
apt-get install ruby-dev build-essential -y
gem install fpm

echo "Create deb package..."

fpm -s dir -t deb -C /tmp/build --name bifrost --version 0.1.0 --iteration 1 --depends debian_dependency1 --description "Digitalbits-bifrost" .

echo "Create rpm package..."

fpm -s dir -t rpm -C /tmp/build --name bifrost --version 0.1.0 --iteration 1 --depends  redhat_dependency1 --description "digitalbits-bifrost" .

echo "deploying to Cloudsmith with cloudsmith-cli"

ls
cloudsmith push deb digitalbits/dbtest/ubuntu/trusty bifrost_0.1.0-1_amd64.deb
cloudsmith push rpm digitalbits/dbtest/el/7 bifrost-0.1.0-1.x86_64.rpm 
