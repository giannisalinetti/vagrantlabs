## origin_aio vagrant lab

This Vagrantfile provides a simple way to build an Origin All-in-One  lab.

To install Vagrant please refer to the documentation on https://www.vagrantup.com/docs/installation/

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

### Lab features
Provisioning is automated using shell syntax and ansible plabooks.
The docker daemon is automatically started at boot.
The latest OpenShift Origin client is installed. To run a complete Origin cluster on the single machine run:

```shell
# /home/vagrant/oc-cluster-up.sh
```
