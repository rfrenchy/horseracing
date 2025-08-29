# Use official Node.js image as the base image
FROM node:20-alpine AS builder

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json (if available)
COPY package.json .
COPY package-lock.json .

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Build the Next.js app
RUN npm run build

# Production image, copy built assets and install only production dependencies
FROM node:20-alpine AS runner
WORKDIR /app

ENV NODE_ENV=production

# Copy package.json and package-lock.json
COPY package.json .
COPY package-lock.json .

# Install only production dependencies
RUN npm install --omit=dev

# Copy built app from builder
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public
COPY --from=builder /app/src ./src
COPY --from=builder /app/next.config.* ./
COPY --from=builder /app/next-env.d.ts ./
COPY --from=builder /app/tsconfig.json ./

# Expose port 3000
EXPOSE 3000

# Start the Next.js app
CMD ["npm", "start"]
