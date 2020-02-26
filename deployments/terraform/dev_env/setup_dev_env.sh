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

