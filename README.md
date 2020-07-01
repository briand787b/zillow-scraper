# Zillow-Scraper
Scrapes Zillow for favorited properties

# Cool Features
1. Chrome Extension: Runs watchman in docker container with node.js runtime and a local filesystem bind mount so that every modification to the chrome react app triggers a new `npm run build`.  The execution of this command is required for packaging up the static bundle that chrome runs.  Once this has happened, however, subsequent invocations of `npm run build` will allow updates to the option page to be visible by refreshing the options page, rather than re-uploading the packaged assets (takes much longer); -- Major Downside -> does not work correctly on MacOS because virtual box (which the Docker containers are run inside of) does not forward inotify events when the bind mounted files on the host are modified (https://github.com/moby/moby/issues/18246) (https://blog.codecentric.de/en/2017/08/fix-webpack-watch-virtualbox/).

# Gotchas
1. This app will not work in a chrome extension without this: https://stackoverflow.com/a/52498189