FROM mcr.microsoft.com/devcontainers/base:ubuntu

# Install vim for git commit messages
RUN apt-get update && apt-get install -y vim

# Install Kurtosis CLI
RUN echo "deb [trusted=yes] https://apt.fury.io/kurtosis-tech/ /" | sudo tee /etc/apt/sources.list.d/kurtosis.list
RUN sudo apt update
RUN sudo apt install kurtosis-cli