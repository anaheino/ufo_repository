# Use the official MongoDB image from Docker Hub
FROM mongo
# Expose MongoDB port
EXPOSE 27017
# Copy the MongoDB script to the container
COPY init-mongo.js /docker-entrypoint-initdb.d/
# Set the default command to run MongoDB
CMD ["mongod", "--bind_ip", "0.0.0.0"]