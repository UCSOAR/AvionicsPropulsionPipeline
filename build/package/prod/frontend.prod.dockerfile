# Stage 1: Build the application
FROM node:18 AS build

# Set the working directory inside the container
WORKDIR /app

# Install dependencies
COPY ./web/package*.json ./
RUN npm install

# Copy the rest of the frontend application files
COPY ./web .

# Build the application
RUN npm run build

# Stage 2: Serve the application with Nginx
FROM nginx:latest

# Copy the build output from the previous stage
COPY --from=build /app/dist /usr/share/nginx/html

# Copy custom Nginx configuration file
COPY ./web/nginx.conf /etc/nginx/nginx.conf

# Expose port 80
EXPOSE 80

# Start Nginx server
CMD ["nginx", "-g", "daemon off;"]