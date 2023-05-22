#!/bin/bash

BASEDIR=$(dirname "$0")
sed -i "s/username:/username: ${DDNS_USERNAME}/g" ${BASEDIR}/config_test.yaml
sed -i "s/password:/password: ${DDNS_PASSWORD}/g" ${BASEDIR}/config_test.yaml
