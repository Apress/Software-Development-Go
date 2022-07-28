#!/usr/bin/env bash

# Destroy everything
ip netns delete ns1
ip netns delete ns2
ip link delete br0

sleep 1

# create namespaces
ip netns add ns1
ip netns add ns2

# setup local interface inside namespace
ip netns exec ns1 ip link set lo up
ip netns exec ns1 ip link

ip netns exec ns2 ip link set lo up
ip netns exec ns2 ip link

# setup bridge
ip link add br0 type bridge
ip link set br0 up
# setup bridge IP
ip addr add 10.0.0.1/8 dev br0

# setup virtual ethernet and link it to namespace
ip link add v0 type veth peer name virt0
ip link set v0 master br0
ip link set v0 up
ip link set virt0 netns ns1
# bring up the virtual ethernet
ip netns exec ns1 ip link set virt0 up
# print out info about the network link
ip netns exec ns1 ip link

# setup virtual ethernet and link it to namespace
ip link add v1 type veth peer name virt1
ip link set v1 master br0
ip link set v1 up
ip link set virt1 netns ns2
# bring up the virtual ethernet
ip netns exec ns2 ip link set virt1 up
# print out info about the network link
ip netns exec ns2 ip link

# Set IP address to the different virtual interfaces
ip netns exec ns1 ip addr add 10.0.0.10/8 dev virt0
ip netns exec ns2 ip addr add 10.0.0.11/8 dev virt1

# register the bridge in iptables to allow forwarding
iptables -I FORWARD -i br0 -o br0 -j ACCEPT

# Check network connection inside the different network namespaces
ip netns exec ns1 ping -c 5 10.0.0.1
ip netns exec ns2 ping -c 5 10.0.0.1

ip netns exec ns1 ping -c 5 10.0.0.10
ip netns exec ns2 ping -c 5 10.0.0.10

ip netns exec ns1 ping -c 5 10.0.0.11
ip netns exec ns2 ping -c 5 10.0.0.11

# Destroy everything
ip netns delete ns1
ip netns delete ns2
ip link delete br0


# delete the iptables information
# run the following
#   iptables  -v --list FORWARD  --line-number
# get the chain num and use the following to delete
#   iptables  -v --delete FORWARD  <chainnum>



