FROM node:14.4

# install watchman
# TODO: shrink number of stages in this build
WORKDIR /tmp
RUN apt-get update
RUN apt-get install -y  autoconf automake build-essential python-setuptools python-dev libssl-dev libtool
RUN git clone https://github.com/facebook/watchman.git
WORKDIR /tmp/watchman
RUN git checkout v4.9.0
RUN ./autogen.sh
RUN ./configure
RUN make
RUN make install

# load application
WORKDIR /usr/local/src
COPY package*.json ./
RUN npm install
COPY . ./

# execute watchman
# ENTRYPOINT [ "watchman " ]