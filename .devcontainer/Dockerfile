FROM public.ecr.aws/lts/ubuntu:22.04_stable

RUN apt-get update && apt-get install -y \
    curl \
    wget \
    libc6 \
    libstdc++6 \
    unzip \
    git \
    ca-certificates \
    tar \
    fontconfig && \
    adduser --disabled-password --gecos "" jordan && \
    usermod -aG sudo jordan && \
    wget https://rclone.org/install.sh && \
    bash install.sh && \ 
    wget https://go.dev/dl/go1.20.6.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.20.6.linux-amd64.tar.gz && \
    wget https://github.com/be5invis/Iosevka/releases/download/v25.1.1/ttf-iosevka-25.1.1.zip && \
    unzip ttf-iosevka-25.1.1.zip -d /usr/share/fonts && \
    fc-cache -f -v && \
    wget https://nodejs.org/dist/v18.17.0/node-v18.17.0-linux-x64.tar.xz && \
    tar -xf node-v18.17.0-linux-x64.tar.xz && \
    mv node-v18.17.0-linux-x64 /home/jordan/node


USER jordan

