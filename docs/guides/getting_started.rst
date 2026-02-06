Getting Started
===============

This guide will help you get started with NeuralBlitz v50 Go quickly.

Prerequisites
-------------

Before you begin, ensure you have the following installed:

- **Go 1.21+**: The Go programming language
- **Git**: Version control system
- **Make**: Build automation tool
- **Docker** (optional): For containerized deployment

Installation
------------

1. **Clone the Repository**

   .. code-block:: bash

      git clone https://github.com/neuralblitz/neuralblitz-v50-go.git
      cd neuralblitz-v50-go

2. **Install Dependencies**

   .. code-block:: bash

      go mod download

3. **Verify Installation**

   .. code-block:: bash

      go version
      go env

4. **Build the Project**

   .. code-block:: bash

      make build

   Or manually:

   .. code-block:: bash

      go build ./pkg/systems/... ./pkg/quantum/...

5. **Run Tests**

   .. code-block:: bash

      make test

   Or manually:

   .. code-block:: bash

      go test ./pkg/systems/... ./pkg/quantum/... -v

Your First NeuralBlitz Program
------------------------------

Create a simple program that demonstrates quantum spiking neurons:

.. code-block:: go

   package main

   import (
       "fmt"
       "log"
       "time"

       "neuralblitz/pkg/systems"
   )

   func main() {
       // Create a quantum spiking neuron
       neuron := systems.NewQuantumSpikingNeuron("neuron_1")
       
       // Configure the neuron
       config := systems.NeuronConfiguration{
           RestingPotential:    -70.0,
           ThresholdPotential:  -55.0,
           MembraneResistance:  1.0,
           MembraneTimeConstant: 20.0,
           MaxHistory:          1000,
       }
       
       // Initialize with configuration
       neuron.Initialize(config)
       
       // Run simulation for 100ms
       fmt.Println("Running quantum neuron simulation...")
       
       for t := 0; t < 1000; t++ {
           // Apply input current
           current := 20.0 // nA
           neuron.Step(current)
           
           // Print every 100 steps
           if t % 100 == 0 {
               state := neuron.GetState()
               fmt.Printf("Time: %dms, Membrane Potential: %.2f mV, Spikes: %d\n",
                   t/10, state.MembranePotential, len(state.SpikeTimes))
           }
       }
       
       // Get final statistics
       stats := neuron.GetStats()
       fmt.Printf("\nFinal Statistics:\n")
       fmt.Printf("Total Spikes: %d\n", stats.TotalSpikes)
       fmt.Printf("Average Spike Rate: %.2f Hz\n", stats.SpikeRate)
       fmt.Printf("Quantum Coherence: %.4f\n", stats.QuantumCoherence)
   }

Save this as ``main.go`` and run:

.. code-block:: bash

   go run main.go

Working with Multi-Reality Networks
-----------------------------------

Example of creating a multi-reality neural network:

.. code-block:: go

   package main

   import (
       "fmt"
       "neuralblitz/pkg/systems"
   )

   func main() {
       // Create a multi-reality network with 4 realities
       network := systems.NewMultiRealityNeuralNetwork(4, 25)
       
       // Configure reality types
       config := systems.RealityNetworkConfig{
           RealityTypes: []systems.RealityType{
               systems.BASE_REALITY,
               systems.QUANTUM_DIVERGENT,
               systems.TEMPORAL_INVERTED,
               systems.CONSCIOUSNESS_AMPLIFIED,
           },
           ConnectionProbability: 0.3,
           EntropyBudget:        0.1,
       }
       
       // Initialize network
       network.Initialize(config)
       
       // Run evolution cycles
       fmt.Println("Evolving multi-reality network...")
       
       for cycle := 0; cycle < 10; cycle++ {
           network.Evolve()
           
           metrics := network.GetMetrics()
           fmt.Printf("Cycle %d: Consciousness=%.3f, Coherence=%.3f\n",
               cycle, metrics.GlobalConsciousness, metrics.CrossRealityCoherence)
       }
       
       // Get final state
       finalMetrics := network.GetMetrics()
       fmt.Printf("\nFinal Metrics:\n")
       fmt.Printf("Global Consciousness: %.4f\n", finalMetrics.GlobalConsciousness)
       fmt.Printf("Cross-Reality Coherence: %.4f\n", finalMetrics.CrossRealityCoherence)
       fmt.Printf("Information Flow Rate: %.2f bits/s\n", finalMetrics.InformationFlowRate)
   }

Understanding the Architecture
------------------------------

NeuralBlitz v50 Go consists of three main architectural layers:

1. **Systems Layer** (``pkg/systems/``)
   - High-level cognitive modules
   - Agent frameworks and self-evolution
   - Consciousness and reality models

2. **Quantum Layer** (``pkg/quantum/``)
   - Quantum-inspired computational models
   - Entanglement and cryptography
   - Reality simulation

3. **Integration Layer** (``pkg/core/``, ``pkg/bridge/``)
   - Connects systems and quantum layers
   - API and protocol implementations
   - Data flow management

Next Steps
----------

Now that you've seen the basics, explore more advanced topics:

- :doc:`/guides/configuration` - Configure NeuralBlitz for your use case
- :doc:`/examples/quantum_neurons` - Deep dive into quantum neurons
- :doc:`/examples/multi_reality` - Build multi-reality applications
- :doc:`/architecture/overview` - Understand the system architecture

Common Issues
-------------

**Build Errors**

If you encounter build errors, ensure you have Go 1.21+:

.. code-block:: bash

   go version  # Should show go1.21 or higher

**Import Errors**

Make sure your ``GOPATH`` is set correctly:

.. code-block:: bash

   export PATH=$PATH:$(go env GOPATH)/bin

**Test Failures**

Some tests may require specific environment variables:

.. code-block:: bash

   export NEURALBLITZ_ENV=test
   export NEURALBLITZ_LOG_LEVEL=debug

Getting Help
------------

- **Documentation**: https://neuralblitz.readthedocs.io
- **GitHub Issues**: https://github.com/neuralblitz/neuralblitz-v50-go/issues
- **Community Discord**: https://discord.gg/neuralblitz
- **Stack Overflow**: Tag your questions with ``neuralblitz``

Contributing
------------

We welcome contributions! See :doc:`/reference/contributing` for guidelines.

License
-------

NeuralBlitz v50 Go is released under the MIT License. See LICENSE file
for details.
