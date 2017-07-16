Vagrant.configure(2) do |config|
  config.vm.box = "centos/7"

  # config.vm.box_check_update = false

  config.vm.hostname = "rtd.extraordy.com"
  config.vm.network "forwarded_port", guest: 8000, host: 8000

  config.vm.provision "ansible" do |ansible|
    ansible.playbook = "rtd_setup.yml"
  end

end
