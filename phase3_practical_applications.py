#!/usr/bin/env python3
"""
Phase 3: Practical Applications
Build working tools using validated NeuralBlitz components
"""

import sys
import numpy as np
import time

sys.path.insert(0, "/home/runner/workspace/NB-Ecosystem/lib/python3.11/site-packages")
sys.path.insert(0, "/home/runner/workspace/opencode-lrs-agents-nbx/neuralblitz-v50")
sys.path.insert(
    0,
    "/home/runner/workspace/opencode-lrs-agents-nbx/neuralblitz-v50/Advanced-Research/production",
)

print("=" * 70)
print("PHASE 3: BUILDING PRACTICAL APPLICATIONS")
print("=" * 70)

# Application 1: Quantum Pattern Recognizer
print("\n1. QUANTUM PATTERN RECOGNIZER")
print("-" * 70)

from quantum_spiking_neuron import QuantumSpikingNeuron, NeuronConfiguration


class QuantumPatternRecognizer:
    """Pattern recognition using quantum spiking neurons"""

    def __init__(self, num_neurons=5):
        self.neurons = []
        for i in range(num_neurons):
            config = NeuronConfiguration(
                quantum_tunneling=0.05 + i * 0.03, coherence_time=50 + i * 25
            )
            neuron = QuantumSpikingNeuron(f"pattern_{i}", config)
            self.neurons.append(neuron)

    def encode_pattern(self, pattern):
        """Encode pattern into spike signatures"""
        signatures = []
        for neuron in self.neurons:
            neuron.reset()
            results = neuron.simulate(pattern * 15.0)  # Scale to input current
            signatures.append(
                {
                    "spike_count": results["spike_count"],
                    "spike_rate": results["spike_rate"],
                    "coherence": results["final_state"].coherence,
                }
            )
        return signatures

    def recognize(self, test_pattern, reference_patterns):
        """Compare pattern against references"""
        test_sig = self.encode_pattern(test_pattern)

        best_match = None
        best_score = float("inf")

        for name, ref_pattern in reference_patterns.items():
            ref_sig = self.encode_pattern(ref_pattern)

            # Calculate distance
            distance = 0
            for t, r in zip(test_sig, ref_sig):
                distance += abs(t["spike_count"] - r["spike_count"])
                distance += abs(t["spike_rate"] - r["spike_rate"]) * 10

            if distance < best_score:
                best_score = distance
                best_match = name

        confidence = 1.0 / (1.0 + best_score / len(test_sig))
        return best_match, confidence


# Test pattern recognition
recognizer = QuantumPatternRecognizer(num_neurons=3)

# Define reference patterns (binary encoded)
patterns = {
    "VERTICAL": np.array([0, 1, 1, 1, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1, 0]),
    "HORIZONTAL": np.array([0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0]),
    "DIAGONAL": np.array([1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0]),
}

# Test with noisy vertical pattern
test = np.array([0, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 0])  # Noisy vertical
match, conf = recognizer.recognize(test, patterns)
print(f"   Pattern recognized: {match} (confidence: {conf:.2%})")

# Application 2: Multi-Reality Optimizer
print("\n2. MULTI-REALITY OPTIMIZER")
print("-" * 70)

from multi_reality_nn import MultiRealityNeuralNetwork


class MultiRealityOptimizer:
    """Use parallel realities to optimize solutions"""

    def __init__(self, objective_func, bounds):
        self.objective_func = objective_func
        self.bounds = bounds
        self.mrnn = MultiRealityNeuralNetwork(
            num_realities=4, nodes_per_reality=len(bounds)
        )
        self.best_solution = None
        self.best_score = float("inf")

    def optimize(self, iterations=30):
        """Run optimization across parallel realities"""
        print("   Optimizing across 4 parallel realities...")

        for i in range(iterations):
            # Get state from each reality
            for rid, reality in self.mrnn.realities.items():
                # Map node states to solution space
                solution = np.clip(reality.node_states[: len(self.bounds)], -5, 5)

                # Evaluate
                score = self.objective_func(solution)

                if score < self.best_score:
                    self.best_score = score
                    self.best_solution = solution.copy()
                    print(f"   Iteration {i}: New best score = {score:.4f}")

            # Evolve realities
            self.mrnn.process_multi_reality_computation({})

        return self.best_solution, self.best_score


# Test optimization
def rastrigin(x):
    """Test function - global minimum at 0"""
    return 10 * len(x) + sum(xi**2 - 10 * np.cos(2 * np.pi * xi) for xi in x)


optimizer = MultiRealityOptimizer(rastrigin, [(-5, 5)] * 3)
solution, score = optimizer.optimize(iterations=20)
print(f"\n   Best solution found: {solution}")
print(f"   Best score: {score:.4f} (theoretical min: 0.0)")

# Application 3: Consciousness Monitor
print("\n3. CONSCIOUSNESS MONITOR")
print("-" * 70)

mrnn = MultiRealityNeuralNetwork(num_realities=4, nodes_per_reality=30)
history = mrnn.evolve_multi_reality_network(num_cycles=15)
state = mrnn.get_multi_reality_state()

print(f"   Global Consciousness: {state['global_consciousness']:.4f}")
print(f"   Cross-Reality Coherence: {state['cross_reality_coherence']:.4f}")
print(f"   Reality Synchronization: {state['reality_synchronization']:.3f}")
print(f"   Active Signals: {state['active_signals']}")

print("\n   Reality-by-Reality Breakdown:")
for rid, rstate in state["reality_states"].items():
    print(
        f"     {rid}: consciousness={rstate['consciousness_level']:.3f}, "
        f"density={rstate['information_density']:.1f}"
    )

print("\n" + "=" * 70)
print("PHASE 3 COMPLETE: 3 Practical Applications Built & Tested")
print("=" * 70)
print("\nApplications Created:")
print("  ✅ 1. Quantum Pattern Recognizer")
print("  ✅ 2. Multi-Reality Optimizer")
print("  ✅ 3. Consciousness Monitor")
