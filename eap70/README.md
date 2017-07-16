## JBoss EAP 7.0.0 Vagrantfile
This Vagrantfile provides a simple way to build a JBoss EAP 7.0 lab.

To install Vagrant please refer to the documentation on https://www.vagrantup.com/docs/installation/

To successfully build the machine you need to download the JBoss EAP zip file from the JBoss Developer Network and place it under this directory.
https://developers.redhat.com/downloads/?referrer=jbd

Otherwise, a jboss-eap-7.0.0.zip file will be downloaded from a common repository during the virtual machine provisioning.

For building purposes the Apache Maven has been installed. Please download the JBoss EAP 7 Quickstarts from GitHub:
https://github.com/jboss-developer/jboss-eap-quickstarts

To run the vagrant machine:

```shell
# vagrant up
```

To access the vagrant machine:

```shell
# vagrant ssh
```

To provision the vagrant machine again:

```shell
# vagrant provision
```

### Windows users only:
Please install rsync to enable synced folders. Otherwise, uncomment the following line in the vagrant file to disable synced folders in Vagrant.

```shell
# config.vm.synced_folder ".", "/vagrant", disabled: true
```

Missing SSH client: if no ssh client is installed on the Windows hypervisor, then the vagrant ssh command will produce an error. 

The easier way to fix this issue is to install the latest git client from the following URL: https://git-scm.com/download/win

