# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.16, 1.17,1.18, 1-bullseye, 1.16-bullseye, 1.17-bullseye, 1.18-bullseye, 1-buster, 1.16-buster, 1.17-buster
ARG VARIANT=1.18-bullseye
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Choice] Node.js version: none, lts/*, 16, 14, 12, 10
ARG NODE_VERSION="none"
RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next lines to use go get to install anything else you need
# USER vscode
# RUN go get -x <your-dependency-or-tool>
# USER root

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
RUN apt update && apt upgrade -y

# use localhost ssl
# RUN apt install -y mkcert
RUN apt install -y libnss3-tools

RUN wget https://github.com/FiloSottile/mkcert/releases/download/v1.4.3/mkcert-v1.4.3-linux-amd64 \
    && mv mkcert* mkcert \
    && install ./mkcert /usr/bin/ \
    && rm ./mkcert

# install htpasswd for bcrypt
RUN apt-get install apache2-utils

# golang install tools
RUN go install golang.org/x/tools/cmd/godoc@latest

RUN echo 'export GOPATH=/home/vscode/go' >> /home/vscode/.bash_profile
# vscode の拡張が環境変数を設定していないと理解できない構成なので、設定しておく。
RUN echo 'export GOROOT=/usr/local/go' >> /home/vscode/.bash_profile

RUN echo 'export PORT=8080' >> /home/vscode/.bash_profile
