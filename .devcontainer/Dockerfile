FROM mcr.microsoft.com/devcontainers/base:ubuntu

# Install vim for git commit messages
RUN apt-get update && apt-get install -y vim

# Install Docker CLI and Daemon
RUN apt-get update && apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    software-properties-common && \
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg && \
    echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" > /etc/apt/sources.list.d/docker.list && \
    apt-get update && apt-get install -y docker-ce docker-ce-cli containerd.io

# Enable Docker service
RUN systemctl enable docker.service && systemctl enable containerd.service

# Add the default user to the docker group
RUN usermod -aG docker vscode

# Install Kurtosis CLI
RUN echo "deb [trusted=yes] https://apt.fury.io/kurtosis-tech/ /" | sudo tee /etc/apt/sources.list.d/kurtosis.list
RUN apt update
RUN apt install kurtosis-cli

# Install Foundry
ENV FOUNDRY_DIR=/usr/local
RUN curl -L https://foundry.paradigm.xyz | bash
RUN foundryup