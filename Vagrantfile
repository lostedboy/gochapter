# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
  config.vm.box = "ubuntu/xenial64"

  config.vm.network "private_network", ip: "192.168.33.10"

  config.vm.provision "shell", path: "dist/scripts/apt-install.sh"

  config.vm.synced_folder ".", "/vagrant", type: "nfs"

  config.vm.hostname = "vagrant.loc"
end
