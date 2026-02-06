Quantum Spiking Neurons Example
===============================

This example demonstrates how to use quantum spiking neurons for
neural computation with quantum-inspired dynamics.

Basic Usage
-----------

.. code-block:: go

   package main

   import (
       "fmt"
       "log"
       "neuralblitz/pkg/systems"
   )

   func main() {
       // Create a quantum spiking neuron
       neuron := systems.NewQuantumSpikingNeuron("neuron_1")
       
       // Configure with custom parameters
       config := systems.NeuronConfiguration{
           RestingPotential:     -70.0,
           ThresholdPotential:   -55.0,
           MembraneResistance:   1.0,
           MembraneTimeConstant: 20.0,
           MaxHistory:          1000,
       }
       
       if err := neuron.Initialize(config); err != nil {
           log.Fatal(err)
       }
       
       // Run simulation
       fmt.Println("Running 1000 time steps...")
       
       for t := 0; t < 1000; t++ {
           // Apply varying input current
           current := 15.0 + 5.0*float64(t%100)/100.0
           neuron.Step(current)
       }
       
       // Get results
       stats := neuron.GetStats()
       fmt.Printf("Total Spikes: %d\n", stats.TotalSpikes)
       fmt.Printf("Spike Rate: %.2f Hz\n", stats.SpikeRate)
       fmt.Printf("Quantum Coherence: %.4f\n", stats.QuantumCoherence)
   }

Multiple Neurons
----------------

.. code-block:: go

   package main

   import (
       "fmt"
       "sync"
       "neuralblitz/pkg/systems"
   )

   func main() {
       // Create a network of quantum neurons
       numNeurons := 10
       neurons := make([]*systems.QuantumSpikingNeuron, numNeurons)
       
       for i := 0; i < numNeurons; i++ {
           neurons[i] = systems.NewQuantumSpikingNeuron(
               fmt.Sprintf("neuron_%d", i),
           )
           neurons[i].Initialize(systems.NeuronConfiguration{
               RestingPotential:   -70.0,
               ThresholdPotential: -55.0,
           })
       }
       
       // Run parallel simulation
       var wg sync.WaitGroup
       results := make([]*systems.NeuronStats, numNeurons)
       
       for i, neuron := range neurons {
           wg.Add(1)
           go func(idx int, n *systems.QuantumSpikingNeuron) {
               defer wg.Done()
               
               // Each neuron gets different input
               for t := 0; t < 500; t++ {
                   current := 10.0 + float64(idx)*2.0
                   n.Step(current)
               }
               
               results[idx] = n.GetStats()
           }(i, neuron)
       }
       
       wg.Wait()
       
       // Print results
       for i, stats := range results {
           fmt.Printf("Neuron %d: %d spikes, %.2f Hz\n",
               i, stats.TotalSpikes, stats.SpikeRate)
       }
   }

Neuron Population Analysis
--------------------------

.. code-block:: go

   package main

   import (
       "fmt"
       "math"
       "neuralblitz/pkg/systems"
   )

   type PopulationStats struct {
       TotalSpikes      int
       AverageRate      float64
       MaxCoherence     float64
       MinCoherence     float64
       Synchronization  float64
   }

   func analyzePopulation(neurons []*systems.QuantumSpikingNeuron) PopulationStats {
       var totalSpikes int
       var totalRate float64
       maxCoherence := 0.0
       minCoherence := 1.0
       
       for _, neuron := range neurons {
           stats := neuron.GetStats()
           totalSpikes += stats.TotalSpikes
           totalRate += stats.SpikeRate
           
           if stats.QuantumCoherence > maxCoherence {
               maxCoherence = stats.QuantumCoherence
           }
           if stats.QuantumCoherence < minCoherence {
               minCoherence = stats.QuantumCoherence
           }
       }
       
       n := float64(len(neurons))
       return PopulationStats{
           TotalSpikes:     totalSpikes,
           AverageRate:     totalRate / n,
           MaxCoherence:    maxCoherence,
           MinCoherence:    minCoherence,
           Synchronization: maxCoherence - minCoherence,
       }
   }

   func main() {
       // Create population
       population := make([]*systems.QuantumSpikingNeuron, 100)
       
       for i := range population {
           population[i] = systems.NewQuantumSpikingNeuron(
               fmt.Sprintf("pop_neuron_%d", i),
           )
           population[i].Initialize(systems.NeuronConfiguration{
               RestingPotential:     -65.0,
               ThresholdPotential:   -50.0,
               MembraneResistance:   1.5,
               MembraneTimeConstant: 25.0,
           })
       }
       
       // Stimulate population
       fmt.Println("Stimulating neuron population...")
       
       for t := 0; t < 1000; t++ {
           for i, neuron := range population {
               // Create patterned input
               phase := float64(t) / 100.0
               baseCurrent := 12.0
               oscillation := 8.0 * math.Sin(phase+float64(i)*0.1)
               
               neuron.Step(baseCurrent + oscillation)
           }
       }
       
       // Analyze results
       stats := analyzePopulation(population)
       
       fmt.Printf("\nPopulation Analysis:\n")
       fmt.Printf("Total Spikes: %d\n", stats.TotalSpikes)
       fmt.Printf("Average Rate: %.2f Hz\n", stats.AverageRate)
       fmt.Printf("Max Coherence: %.4f\n", stats.MaxCoherence)
       fmt.Printf("Min Coherence: %.4f\n", stats.MinCoherence)
       fmt.Printf("Synchronization: %.4f\n", stats.Synchronization)
   }

Advanced Configuration
----------------------

.. code-block:: go

   package main

   import (
       "fmt"
       "neuralblitz/pkg/systems"
   )

   func createOptimizedNeuron(id string, neuronType string) *systems.QuantumSpikingNeuron {
       neuron := systems.NewQuantumSpikingNeuron(id)
       
       var config systems.NeuronConfiguration
       
       switch neuronType {
       case "fast_spiking":
           config = systems.NeuronConfiguration{
               RestingPotential:     -70.0,
               ThresholdPotential:   -55.0,
               MembraneResistance:   0.5,
               MembraneTimeConstant: 10.0,
               MaxHistory:          500,
           }
       case "regular_spiking":
           config = systems.NeuronConfiguration{
               RestingPotential:     -70.0,
               ThresholdPotential:   -55.0,
               MembraneResistance:   1.0,
               MembraneTimeConstant: 20.0,
               MaxHistory:          1000,
           }
       case "bursting":
           config = systems.NeuronConfiguration{
               RestingPotential:     -75.0,
               ThresholdPotential:   -50.0,
               MembraneResistance:   2.0,
               MembraneTimeConstant: 30.0,
               MaxHistory:          2000,
           }
       default:
           config = systems.NeuronConfiguration{
               RestingPotential:     -70.0,
               ThresholdPotential:   -55.0,
               MembraneResistance:   1.0,
               MembraneTimeConstant: 20.0,
               MaxHistory:          1000,
           }
       }
       
       neuron.Initialize(config)
       return neuron
   }

   func main() {
       // Create different neuron types
       fast := createOptimizedNeuron("fast_1", "fast_spiking")
       regular := createOptimizedNeuron("regular_1", "regular_spiking")
       bursting := createOptimizedNeuron("bursting_1", "bursting")
       
       // Test each type
       neurons := []*systems.QuantumSpikingNeuron{fast, regular, bursting}
       names := []string{"Fast Spiking", "Regular Spiking", "Bursting"}
       
       for i, neuron := range neurons {
           // Run 500ms simulation
           for t := 0; t < 500; t++ {
               neuron.Step(20.0)
           }
           
           stats := neuron.GetStats()
           fmt.Printf("%s: %d spikes (%.1f Hz), Coherence: %.4f\n",
               names[i], stats.TotalSpikes, stats.SpikeRate, stats.QuantumCoherence)
       }
   }

Running the Examples
--------------------

Save the examples as ``main.go`` and run:

.. code-block:: bash

   go run main.go

Or build and run:

.. code-block:: bash

   go build -o quantum_neurons main.go
   ./quantum_neurons

Expected Output
---------------

Typical output for the basic example:

.. code-block:: text

   Running 1000 time steps...
   Total Spikes: 47
   Spike Rate: 47.00 Hz
   Quantum Coherence: 0.9234

Typical output for the population analysis:

.. code-block:: text

   Stimulating neuron population...
   
   Population Analysis:
   Total Spikes: 3247
   Average Rate: 32.47 Hz
   Max Coherence: 0.9876
   Min Coherence: 0.7823
   Synchronization: 0.2053

Performance Tips
----------------

1. **Use Parallel Processing**: For large populations, use goroutines
2. **Adjust Time Step**: Smaller time steps increase accuracy but slow simulation
3. **Limit History**: Set ``MaxHistory`` based on your analysis needs
4. **Batch Operations**: Process neurons in batches for better cache utilization

Next Steps
----------

- See :doc:`/examples/multi_reality` for reality-based neural processing
- See :doc:`/examples/consciousness` for consciousness integration
- See :doc:`/examples/agents` for autonomous agent examples
