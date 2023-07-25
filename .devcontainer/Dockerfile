FROM public.ecr.aws/lts/ubuntu:22.04_stable

RUN apt-get update && apt-get install -y \
    curl \
    wget \
    libc6 \
    libstdc++6 \
    unzip \
    git \
    ca-certificates \
    tar && \
    wget https://rclone.org/install.sh && \
    bash install.sh && \ 
    wget https://go.dev/dl/go1.20.6.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.20.6.linux-amd64.tar.gz

