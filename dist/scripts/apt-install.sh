#!/bin/bash

# Repair "==> default: stdin: is not a tty" message
sudo sed -i '/tty/!s/mesg n/tty -s \\&\\& mesg n/' /root/.profile

export DEBIAN_FRONTEND=noninteractive
export LANGUAGE=en_US.UTF-8
export LANG=en_US.UTF-8
export LC_ALL=en_US.UTF-8
locale-gen en_US.UTF-8
dpkg-reconfigure locales

echo -e "\n--- APT update ---\n"
sudo apt-get update > /dev/null 2>&1

echo -e "\n--- Install utilities ---\n"
sudo apt-get install -y curl git
sudo apt-get install -y zip gzip tar

echo -e "\n--- Install go ---\n"
sudo wget https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.8.3.linux-amd64.tar.gz
sudo echo "export export GOROOT=/usr/local/go" >> /root/.bashrc
sudo echo "export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin" >> /root/.bashrc
sudo source /root/.bashrc
echo -e "\n--- System install complete ---\n"
