#!/bin/sh

if [ $BUILD_FROM_SRC -gt 0 ]
then
  # deploy horizon binary
  mkdir -p /go/src/github.com/digitalbits/go \
    && git clone --depth 1 --branch master git@github.com:DigitalBitsOrg/go.git /go/src/github.com/digitalbits/go \
    && cd /go/src/github.com/digitalbits/go \
    && curl https://glide.sh/get | sh \
    && glide install \
    && go install github.com/digitalbits/go/services/bifrost

  mv /go/bin/bifrost /usr/local/bin \
  && chmod +x /usr/local/bin/bifrost

apt-get update -y 
apt-get install ruby-dev build-essential -y
gem install fpm

echo "Create deb package..."

fpm -s dir -t deb -C /usr/local/bin/bifrost --name bifrost --version 0.1.0 --iteration 1 --depends debian_dependency1 --description "Digitalbits-bifrost" .

echo "Create rpm package..."

fpm -s dir -t rpm -C /usr/local/bin/bifrost --name bifrost --version 0.1.0 --iteration 1 --depends  redhat_dependency1 --description "digitalbits-bifrost" .

echo "deploying to Cloudsmith with cloudsmith-cli"

ls
cloudsmith push deb digitalbits/dbtest/ubuntu/trusty bifrost_0.1.0-1_amd64.deb
cloudsmith push rpm digitalbits/dbtest/el/7 bifrost-0.1.0-1.x86_64.rpm 
