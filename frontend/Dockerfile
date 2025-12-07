FROM node:22.11.0-alpine

WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy source code
COPY . .

EXPOSE 5173

# Vite dev server will be run via docker-compose command
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]
