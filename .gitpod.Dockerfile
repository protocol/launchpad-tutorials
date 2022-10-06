FROM gitpod/workspace-full:2022-10-05-23-28-15

RUN sudo wget https://dist.ipfs.tech/kubo/v0.16.0/kubo_v0.16.0_linux-amd64.tar.gz
RUN tar -xvzf kubo_v0.16.0_linux-amd64.tar.gz
RUN cd kubo \
    sudo bash install.sh
