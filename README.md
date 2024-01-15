# UFO MAP
A hobby project for learning new technologies, implemented in various frameworks and languages.

The aim is to create a half-decent UFO -sighting map, purely because I think it's funny.

# Core technologies for the map are:

1. MongoDB as the main database
2. Golang gin framework REST API that uses MongoDB driver to perform queries for the front-end
3. Vue.js with typescript as the front-end. Uses OpenStreetMaps to display UFO-sightings on the browser on the map.

All 3 of these are in their own docker containers, and can be launched with "docker-compose up". mongo-folder also contains data.json and import-data.sh script to import the basic sighting data to the MongoDB docker-container.

# Other noteworthy things in the project:

## Additional semi-interesting implementations and technologies in this project are written in Rust and Python.

1. Rust was the language of choice to create a web-scraper to parse the UFO sighting data, mainly because I had not written Rust before.
2. Python is the language of choice for many machine learning tasks.

The python is still a very-much work in progress, but the idea will be to create a separate REST API with it's own docker container, that can call .pkl files (machine learning models) trained with this data.
Currently I've trained a random forest -classifier to provide me with a probability of a ufo sighting when it's given coordinates.

# Split into multiple folders by their purpose:

## rust

- Contains web-scraper and all other data formatting and handling implemented in Rust for parsing sightings to mongoDB

## mongo
- Contains basic mongoDB configurations

## ufo_backend

- contains REST-service that uses MongoDb driver, and offers REST API implemented in golang and gin framework.

## ufo_frontend

- Contains a Vue3 frontend, that communicates with the ufo_backend.

## deeplearning

- Still a work in progress, but will contain all the python code used to create ML models from sighting data.

# Other notes

Please note that this is very much a hobby project, and thus doesn't contain many things a real-life application would, such as proper tests.
Also the code can be very ugly in some parts of the application, because as I've said, I've purposefully used technologies that are a bit unfamiliar to me.