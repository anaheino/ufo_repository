#!/bin/bash
cd ..
docker build -t gin-example .
docker run -p 8080:8080 gin-example
