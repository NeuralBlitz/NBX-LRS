#!/bin/bash
# Setup script for NeuralBlitz exploration environment

echo "🚀 Setting up NeuralBlitz exploration environment..."

# Create virtual environment
python3 -m venv venv_neuralblitz
source venv_neuralblitz/bin/activate

# Upgrade pip
pip install --upgrade pip

# Install core dependencies
echo "📦 Installing core dependencies..."
pip install numpy scipy pandas matplotlib

# Install AI/ML dependencies
echo "🧠 Installing AI dependencies..."
pip install torch torchvision --index-url https://download.pytorch.org/whl/cpu

# Install web framework dependencies
echo "🌐 Installing web framework..."
pip install fastapi uvicorn pydantic

# Install testing dependencies
echo "🧪 Installing testing tools..."
pip install pytest pytest-asyncio httpx

# Install documentation dependencies
echo "📝 Installing documentation tools..."
pip install jupyter notebook

echo "✅ Environment setup complete!"
echo ""
echo "To activate: source venv_neuralblitz/bin/activate"
echo "To test: python opencode-lrs-agents-nbx/neuralblitz-v50/Advanced-Research/production/quantum_spiking_neuron.py"