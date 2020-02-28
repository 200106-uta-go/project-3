#!/bin/bash

#Start Kubelet for overlay connections
systemctl enable kubelet
systemctl start kubelet

kubeadm init

mkdir /home/ubuntu/.kube
cp /etc/kubernetes/admin.conf /home/ubuntu/.kube/config
chown -R ubuntu:ubuntu /home/ubuntu/.kube

kubectl create -f https://docs.projectcalico.org/v3.11/manifests/calico.yaml

iptables -P FORWARD ACCEPT

# get file for workers to join kubernetes cluster
touch run.sh
chmod 777 join.sh
echo "#! /bin/bash -xe" >> join.sh
kubeadm token create --print-join-command >> join.sh

# run terraform to launch workers
terraform init
terraform apply -auto-approve