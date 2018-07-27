#!/bin/sh

#if [ $BUILD_FROM_SRC -gt 0 ]
#then
  # deploy horizon binary
sudo mkdir -p $GOPATH/src/github.com/digitalbits/go \
    && git clone --depth 1 --branch master git@github.com:DigitalBitsOrg/go.git $GOPATH/src/github.com/digitalbits/go \
    && cd $GOPATH/src/github.com/digitalbits/go \
    && curl https://glide.sh/get | sh \
    && glide install \
    && go install github.com/digitalbits/go/services/bifrost \
    && go install github.com/digitalbits/go/services/federation \
    && go install github.com/digitalbits/go/services/horizon 

sudo mv $GOPATH/bin/bifrost /usr/local/bin \
  && sudo chmod +x /usr/local/bin/bifrost 

echo sudo ls -lah /usr/local/bin && echo sudo ls -lah $GOPATH/bin 

sudo apt-get update -y
sudo apt-get install ruby-dev build-essential -y
gem install fpm

echo "Create deb package..."
fpm --verbose -s dir -t deb -C /usr/local/bin/bifrost --name bifrost --version 0.1.0 --iteration 1 --depends debian_dependency1 --description "Digitalbits-bifrost" .

echo "Create rpm package..."

fpm --verbose -s dir -t rpm -C ~/bifrost --name bifrost --version 0.1.0 --iteration 1 --depends  redhat_dependency1 --description "digitalbits-bifrost" .

echo "deploying to Cloudsmith with cloudsmith-cli"

ls
cloudsmith push deb digitalbits/dbtest/ubuntu/trusty bifrost_0.1.0-1_amd64.deb
cloudsmith push rpm digitalbits/dbtest/el/7 bifrost-0.1.0-1.x86_64.rpm 
