# Zillow-Scraper
Scrapes Zillow for favorited properties

# TODOs
1. Chrome Extension: Run watchman in docker container with node.js runtime and a local filesystem bind mount so that every modification to the chrome react app triggers a new `npm run build`.  The execution of this command is required for packaging up the static bundle that chrome runs.  Once this has happened, however, subsequent invocations of `npm run build` will allow updates to the option page to be visible by refreshing the options page, rather than re-uploading the packaged assets (takes much longer);