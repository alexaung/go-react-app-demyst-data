# Use the official Node.js base image
FROM node:16

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json to the container's working directory
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the source code to the container's working directory
COPY . .

# Build the React app
RUN npm run build

# Use the official Nginx base image
FROM nginx:1.21

# Remove the default Nginx configuration file
RUN rm /etc/nginx/conf.d/default.conf

# Copy the built React app to the Nginx container
COPY --from=0 /app/build /usr/share/nginx/html

# Copy the Nginx configuration file
COPY nginx.conf /etc/nginx/conf.d

# Expose the port Nginx will run on
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]
