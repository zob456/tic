# Transparency in Coverage File Location Fetcher
This application fetches the file URLs for the `.gz` files for the `Anthem NY` plans
it logs the URL locations for all the pricing files

### Set-up
1. run `make init`
   1. This runs an init script that will set up you `.env`file & check dependencies
2. run `make start`
   1. This starts the docker container
3. run `make teardown`
   1. This stops the container & removes the volumes
