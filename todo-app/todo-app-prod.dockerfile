# Use the official Nginx image
FROM nginx:alpine

# Copy the build output to the Nginx HTML directory
COPY dist /usr/share/nginx/html

# Copy the custom Nginx configuration
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Expose port 80
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]