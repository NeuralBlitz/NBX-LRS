# NeuralBlitz v50 - Docker Container
FROM python:3.11-slim

# Set working directory
WORKDIR /app

# Install system dependencies
RUN apt-get update && apt-get install -y \
    gcc \
    g++ \
    libopenblas-dev \
    && rm -rf /var/lib/apt/lists/*

# Copy requirements
COPY requirements.txt .

# Install Python dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Copy application code
COPY . .

# Set Python path
ENV PYTHONPATH=/app/NB-Ecosystem/lib/python3.11/site-packages:/app/opencode-lrs-agents-nbx/neuralblitz-v50:/app/opencode-lrs-agents-nbx/neuralblitz-v50/Advanced-Research/production

# Expose ports for web dashboard and APIs
EXPOSE 8080 8888

# Health check
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD python3 -c "import sys; sys.exit(0)"

# Default command
CMD ["python3", "opencode-lrs-agents-nbx/neuralblitz-v50/applications/web_dashboard.py"]
