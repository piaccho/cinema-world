# Use an official Node.js runtime as the base image
FROM node:18-alpine

# Set the working directory in the container to /app
WORKDIR /app

# Copy package.json and pnpm-lock.yaml (if available) into the working directory
COPY package.json pnpm-lock.yaml* ./

# Install any needed packages specified in package.json
RUN npm install -g pnpm && pnpm install

# Bundle the app source inside the Docker image
COPY . .

# Make port 5000 available to the outside world
EXPOSE 5000

# Run the app when the container launches
CMD ["pnpm", "run", "dev"]