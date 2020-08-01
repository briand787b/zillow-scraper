# install application
FROM node:14.4 as installer
WORKDIR /usr/local/src
COPY package*.json ./
RUN npm install
COPY . ./

# install watchman
FROM node:14.4 as watcher
COPY --from=installer /usr/local/src /usr/local/src
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


# run watchman
RUN watchman -- trigger /usr/local/src/src rebuild '*' -- npm run build
CMD [ "watchman", "--foreground", "--logfile=/dev/stdout", "watch-project", "/usr/local/src/src"]