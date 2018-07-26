#!/bin/sh

#if [ $BUILD_FROM_SRC -gt 0 ]
#then
  # deploy horizon binary
sudo mkdir -p /go/src/github.com/digitalbits/go \
    && git clone --depth 1 --branch master git@github.com:DigitalBitsOrg/go.git /go/src/github.com/digitalbits/go \
    && cd /go/src/github.com/digitalbits/go \
    && curl https://glide.sh/get | sh \
    && glide install \
    && go install github.com/digitalbits/go/services/bifrost \
    && go install github.com/digitalbits/go/services/federation \
    && go install github.com/digitalbits/go/services/horizon 

sudo mv /go/bin/bifrost /usr/local/bin \
  && sudo chmod +x /usr/local/bin/bifrost 
  
sudo mv /go/bin/federation /usr/local/bin \
  && sudo chmod +x /usr/local/bin/federation 

sudo mv /go/bin/horizon /usr/local/bin \
  && chmod +x /usr/local/bin/horizon 
  
sudo apt-get update -y
sudo apt-get install ruby-dev build-essential -y
gem install fpm

echo "Create deb package..."
fpm --verbose -s dir -t deb -C /usr/local/bin/bifrost --name bifrost --version 0.1.0 --iteration 1 --depends debian_dependency1 --description "Digitalbits-bifrost" .
fpm -s dir -t deb -C /usr/local/bin/federation --name federation --version 0.1.0 --iteration 1 --depends debian_dependency1 --description "Digitalbits-federation" .
fpm -s dir -t deb -C /usr/local/bin/horizon --name horizon --version 0.1.0 --iteration 1 --depends debian_dependency1 --description "Digitalbits-horizon" .
echo "Create rpm package..."

fpm --verbose -s dir -t rpm -C ~/bifrost --name bifrost --version 0.1.0 --iteration 1 --depends  redhat_dependency1 --description "digitalbits-bifrost" .
fpm -s dir -t rpm -C /usr/local/bin/federation --name federation --version 0.1.0 --iteration 1 --depends  redhat_dependency1 --description "digitalbits-federation" .
fpm -s dir -t rpm -C /usr/local/bin/horizon --name horizon --version 0.1.0 --iteration 1 --depends  redhat_dependency1 --description "digitalbits-horizon" .
echo "deploying to Cloudsmith with cloudsmith-cli"

ls
cloudsmith push deb digitalbits/dbtest/ubuntu/trusty bifrost_0.1.0-1_amd64.deb
cloudsmith push rpm digitalbits/dbtest/el/7 bifrost-0.1.0-1.x86_64.rpm 

cloudsmith push deb digitalbits/dbtest/ubuntu/trusty federation_0.1.0-1_amd64.deb
cloudsmith push rpm digitalbits/dbtest/el/7 federation-0.1.0-1.x86_64.rpm 

cloudsmith push deb digitalbits/dbtest/ubuntu/trusty horizon_0.1.0-1_amd64.deb
cloudsmith push rpm digitalbits/dbtest/el/7 horizon-0.1.0-1.x86_64.rpm 
