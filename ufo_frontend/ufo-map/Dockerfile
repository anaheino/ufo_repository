# Use an official Node.js image as the base image
FROM node:18

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install project dependencies
RUN npm install

# Copy the rest of the application code to the working directory
COPY . .

# Expose the port on which the Vue app will run (default is 8080)
EXPOSE 5173

# Command to run the Vue app
CMD ["npm", "run", "host"]