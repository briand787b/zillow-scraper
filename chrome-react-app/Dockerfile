FROM node:14.4

# install watchman
WORKDIR /tmp
RUN apt-get update && \
    apt-get install -y  autoconf automake build-essential python-setuptools python-dev libssl-dev libtool && \
    git clone https://github.com/facebook/watchman.git
WORKDIR /tmp/watchman
RUN git checkout v4.9.0 && \
    ./autogen.sh && \
    ./configure && \
    make && \
    make install

# install application
WORKDIR /usr/local/src
COPY package*.json ./
RUN npm install
COPY . ./

# run watchman
RUN watchman -- trigger /usr/local/src/src retest '*.js' -- npm run build
CMD [ "watchman", "--foreground", "--logfile=/dev/stdout", "watch", "/usr/local/src/src"]