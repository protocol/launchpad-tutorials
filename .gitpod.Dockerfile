FROM gitpod/workspace-full:2022-10-05-23-28-15

USER gitpod
RUN wget https://dist.ipfs.tech/kubo/v0.16.0/kubo_v0.16.0_linux-amd64.tar.gz \
    && tar -xvzf kubo_v0.16.0_linux-amd64.tar.gz \
    && cd kubo && sudo cp ipfs /usr/bin/ipfs

RUN sudo apt update -y && sudo apt upgrade -y \
    && sudo apt -y install fuse \
    && wget https://github.com/filecoin-project/lotus/releases/download/v1.17.1/Lotus-v1.17.1-x86_64.AppImage \
    && chmod +x Lotus-v1.17.1-x86_64.AppImage \
    && sudo mv Lotus-v1.17.1-x86_64.AppImage /usr/local/bin/lotus
