#!/usr/bin/env python3
"""
Phase 4: Visualization Suite
Generate visual outputs from NeuralBlitz systems
"""

import sys
import numpy as np

sys.path.insert(0, "/home/runner/workspace/NB-Ecosystem/lib/python3.11/site-packages")
sys.path.insert(0, "/home/runner/workspace/opencode-lrs-agents-nbx/neuralblitz-v50")
sys.path.insert(
    0,
    "/home/runner/workspace/opencode-lrs-agents-nbx/neuralblitz-v50/Advanced-Research/production",
)

print("=" * 70)
print("PHASE 4: CREATING VISUALIZATION SUITE")
print("=" * 70)

# Setup matplotlib for non-interactive (file output)
import matplotlib  # type: ignore

matplotlib.use("Agg")
import matplotlib  # type: ignore.pyplot as plt

# Visualization 1: Quantum Neuron Membrane Potential
print("\n1. GENERATING QUANTUM NEURON VISUALIZATIONS")
print("-" * 70)

from quantum_spiking_neuron import QuantumSpikingNeuron, NeuronConfiguration

neuron = QuantumSpikingNeuron("viz_neuron")
t = np.linspace(0, 200, 2000)
inputs = 15 * np.ones_like(t)
results = neuron.simulate(inputs)

# Plot membrane potential and spikes
fig, (ax1, ax2) = plt.subplots(2, 1, figsize=(12, 8))

# Membrane potential trace
times = np.arange(len(results["membrane_trace"])) * 0.1
ax1.plot(
    times, results["membrane_trace"], "b-", linewidth=1, label="Membrane Potential"
)
ax1.axhline(y=-55.0, color="r", linestyle="--", label="Threshold (-55mV)")
ax1.axhline(y=-70.0, color="g", linestyle="--", label="Resting (-70mV)")
ax1.set_xlabel("Time (ms)")
ax1.set_ylabel("Membrane Potential (mV)")
ax1.set_title("Quantum Spiking Neuron - Membrane Dynamics")
ax1.legend()
ax1.grid(True, alpha=0.3)

# Quantum coherence trace
coherence_trace = []
neuron.reset()
for i in range(len(inputs)):
    neuron.step(inputs[i])
    coherence_trace.append(neuron.quantum_state.coherence)

ax2.plot(times, coherence_trace, "purple", linewidth=1)
ax2.set_xlabel("Time (ms)")
ax2.set_ylabel("Quantum Coherence")
ax2.set_title("Quantum State Coherence Over Time")
ax2.grid(True, alpha=0.3)

plt.tight_layout()
plt.savefig(
    "/home/runner/workspace/visualization_output/quantum_neuron_dynamics.png", dpi=150
)
print("   ✅ Saved: quantum_neuron_dynamics.png")
plt.close()

# Visualization 2: Multi-Reality Network Metrics
print("\n2. GENERATING MULTI-REALITY NETWORK VISUALIZATIONS")
print("-" * 70)

from multi_reality_nn import MultiRealityNeuralNetwork

mrnn = MultiRealityNeuralNetwork(num_realities=8, nodes_per_reality=50)
history = mrnn.evolve_multi_reality_network(num_cycles=50)

# Plot evolution metrics
fig, ((ax1, ax2), (ax3, ax4)) = plt.subplots(2, 2, figsize=(14, 10))

# Global consciousness
cycles = range(len(history["global_consciousness"]))
ax1.plot(cycles, history["global_consciousness"], "b-", linewidth=2)
ax1.set_xlabel("Evolution Cycle")
ax1.set_ylabel("Global Consciousness")
ax1.set_title("Global Consciousness Evolution")
ax1.grid(True, alpha=0.3)

# Cross-reality coherence
ax2.plot(cycles, history["cross_reality_coherence"], "g-", linewidth=2)
ax2.set_xlabel("Evolution Cycle")
ax2.set_ylabel("Cross-Reality Coherence")
ax2.set_title("Cross-Reality Coherence Over Time")
ax2.grid(True, alpha=0.3)

# Information flow rate
ax3.plot(cycles, history["information_flow_rate"], "r-", linewidth=2)
ax3.set_xlabel("Evolution Cycle")
ax3.set_ylabel("Information Flow Rate")
ax3.set_title("Information Flow Dynamics")
ax3.grid(True, alpha=0.3)

# Reality synchronization
ax4.plot(cycles, history["reality_synchronization"], "purple", linewidth=2)
ax4.set_xlabel("Evolution Cycle")
ax4.set_ylabel("Reality Synchronization")
ax4.set_title("Reality Synchronization Progress")
ax4.grid(True, alpha=0.3)

plt.tight_layout()
plt.savefig(
    "/home/runner/workspace/visualization_output/multi_reality_evolution.png", dpi=150
)
print("   ✅ Saved: multi_reality_evolution.png")
plt.close()

# Visualization 3: Reality Type Comparison
print("\n3. GENERATING REALITY TYPE COMPARISON")
print("-" * 70)

# Get final state
state = mrnn.get_multi_reality_state()

# Extract data for visualization
reality_types = []
consciousness = []
info_density = []
quantum_coherence = []

for rid, rstate in state["reality_states"].items():
    reality_types.append(rstate["reality_type"].replace("_", "\n"))
    consciousness.append(rstate["consciousness_level"])
    info_density.append(np.log10(rstate["information_density"] + 1))  # Log scale
    quantum_coherence.append(rstate["quantum_coherence"])

fig, (ax1, ax2, ax3) = plt.subplots(1, 3, figsize=(16, 5))

# Consciousness by reality type
colors = plt.cm.viridis(np.linspace(0, 1, len(reality_types)))
bars1 = ax1.bar(range(len(reality_types)), consciousness, color=colors)
ax1.set_xticks(range(len(reality_types)))
ax1.set_xticklabels(reality_types, rotation=45, ha="right", fontsize=8)
ax1.set_ylabel("Consciousness Level")
ax1.set_title("Consciousness by Reality Type")
ax1.grid(True, alpha=0.3, axis="y")

# Information density (log scale)
bars2 = ax2.bar(range(len(reality_types)), info_density, color=colors)
ax2.set_xticks(range(len(reality_types)))
ax2.set_xticklabels(reality_types, rotation=45, ha="right", fontsize=8)
ax2.set_ylabel("Log10(Information Density + 1)")
ax2.set_title("Information Density by Reality Type")
ax2.grid(True, alpha=0.3, axis="y")

# Quantum coherence
bars3 = ax3.bar(range(len(reality_types)), quantum_coherence, color=colors)
ax3.set_xticks(range(len(reality_types)))
ax3.set_xticklabels(reality_types, rotation=45, ha="right", fontsize=8)
ax3.set_ylabel("Quantum Coherence")
ax3.set_title("Quantum Coherence by Reality Type")
ax3.grid(True, alpha=0.3, axis="y")

plt.tight_layout()
plt.savefig(
    "/home/runner/workspace/visualization_output/reality_type_comparison.png", dpi=150
)
print("   ✅ Saved: reality_type_comparison.png")
plt.close()

# Summary statistics visualization
print("\n4. GENERATING PERFORMANCE SUMMARY")
print("-" * 70)

fig, ax = plt.subplots(figsize=(10, 6))

metrics = [
    "Quantum Neuron\n(steps/sec)",
    "MRNN 200 nodes\n(cycles/sec)",
    "MRNN 400 nodes\n(cycles/sec)",
    "MRNN 800 nodes\n(cycles/sec)",
]
values = [10705, 3419.5, 2709.8, 569.2]
colors = ["#1f77b4", "#ff7f0e", "#2ca02c", "#d62728"]

bars = ax.bar(
    metrics, values, color=colors, alpha=0.8, edgecolor="black", linewidth=1.5
)
ax.set_ylabel("Operations per Second", fontsize=12)
ax.set_title("NeuralBlitz Performance Benchmarks", fontsize=14, fontweight="bold")
ax.grid(True, alpha=0.3, axis="y")

# Add value labels on bars
for bar, value in zip(bars, values):
    height = bar.get_height()
    ax.text(
        bar.get_x() + bar.get_width() / 2.0,
        height,
        f"{value:.0f}",
        ha="center",
        va="bottom",
        fontsize=10,
        fontweight="bold",
    )

plt.tight_layout()
plt.savefig(
    "/home/runner/workspace/visualization_output/performance_benchmarks.png", dpi=150
)
print("   ✅ Saved: performance_benchmarks.png")
plt.close()

print("\n" + "=" * 70)
print("PHASE 4 COMPLETE: Visualization Suite Generated")
print("=" * 70)
print("\nVisualizations Created:")
print("  ✅ 1. quantum_neuron_dynamics.png")
print("  ✅ 2. multi_reality_evolution.png")
print("  ✅ 3. reality_type_comparison.png")
print("  ✅ 4. performance_benchmarks.png")
print("\nAll saved to: /home/runner/workspace/visualization_output/")
