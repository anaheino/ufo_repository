# ufo_repository
Containing all the ufo stuff:

A functioning ufo sighting map with queryable sightings, and markers placed on the map based on query results.

Split into multiple modules by their purpose:

# rust

- Contains web-scraper and all other data formatting and handling implemented in Rust for parsing sightings to mongoDB

# ufo_backend

- contains REST-service that uses MongoDb driver, and offers REST API implemented in golang and gin framework.

# ufo_frontend

- Contains a Vue3 frontend, that communicates with the ufo_backend.

# deeplearning

- Still a work in progress, but will contain all the python code used to create ML models from sighting data.
