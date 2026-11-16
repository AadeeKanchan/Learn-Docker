# Step 1: Base image
FROM python:3.12-slim

# Step 2: Set working directory
WORKDIR /app

# Step 3: Set environment variables
ENV PYTHONUNBUFFERED=1 \
    PORT=5000 \
    APP_ENV=production

# Step 4: Copy dependency definitions first (to leverage layer caching)
COPY requirements.txt .

# Step 5: Install dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Step 6: Create a non-root user for security
RUN useradd -m -u 1000 appuser && chown -R appuser:appuser /app
USER appuser

# Step 7: Copy the application source code
COPY . .

# Step 8: Document port
EXPOSE 5000

# Step 9: Define entrypoint & default command
ENTRYPOINT ["python"]
CMD ["app.py"]
