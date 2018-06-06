#!/bin/bash

DATADIR=/var/lib/etcd
IPV4_ADDR=$(hostname -I | cut -d ' ' -f1)

if [ ! -d $DATADIR ]; then
    mkdir -p $DATADIR &> /dev/null
    err=$?; if [ $err -ne 0 ]; then
        echo "Unable to create the data directory, try to run the script with sudo"
        exit 1
    fi
fi

# Startup script for Openshift Origin vagrant lab
oc cluster up \
    --public-hostname=${HOSTNAME} \
    --routing-suffix=${IPV4_ADDR}.xip.io \
    --use-existing-config=true \
    --host-data-dir=${DATADIR}
err=$?; if [ $err -ne 0 ]; then
    echo "Error starting OpenShift cluster"
    exit 1
fi

exit 0
