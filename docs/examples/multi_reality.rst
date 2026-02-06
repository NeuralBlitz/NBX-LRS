Multi-Reality Neural Networks Example
=====================================

This example demonstrates how to use multi-reality neural networks
for parallel processing across different reality contexts.

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
       // Create a multi-reality network
       // 4 realities with 25 nodes each
       mrnn := systems.NewMultiRealityNeuralNetwork(4, 25)
       
       // Configure with different reality types
       config := systems.RealityNetworkConfig{
           RealityTypes: []systems.RealityType{
               systems.BASE_REALITY,
               systems.QUANTUM_DIVERGENT,
               systems.TEMPORAL_INVERTED,
               systems.CONSCIOUSNESS_AMPLIFIED,
           },
           ConnectionProbability: 0.3,
           EntropyBudget:        0.15,
           EthicsMode:          systems.EthicsBalanced,
       }
       
       if err := mrnn.Initialize(config); err != nil {
           log.Fatal(err)
       }
       
       // Evolve the network
       fmt.Println("Evolving multi-reality network...")
       
       for cycle := 0; cycle < 20; cycle++ {
           mrnn.Evolve()
           
           if cycle % 5 == 0 {
               metrics := mrnn.GetMetrics()
               fmt.Printf("Cycle %d: Consciousness=%.3f, Coherence=%.3f\n",
                   cycle, metrics.GlobalConsciousness, metrics.CrossRealityCoherence)
           }
       }
       
       // Get final metrics
       finalMetrics := mrnn.GetMetrics()
       fmt.Printf("\nFinal Results:\n")
       fmt.Printf("Global Consciousness: %.4f\n", finalMetrics.GlobalConsciousness)
       fmt.Printf("Cross-Reality Coherence: %.4f\n", finalMetrics.CrossRealityCoherence)
       fmt.Printf("Information Flow Rate: %.2f bits/s\n", finalMetrics.InformationFlowRate)
   }

Parallel Realities
------------------

.. code-block:: go

   package main

   import (
       "fmt"
       "sync"
       "neuralblitz/pkg/systems"
   )

   func main() {
       // Create multiple independent reality networks
       networks := make([]*systems.MultiRealityNeuralNetwork, 3)
       
       // Different configurations for each network
       configs := []struct {
           name       string
           realities  int
           nodes      int
           types      []systems.RealityType
       }{
           {
               name:      "Classical",
               realities: 2,
               nodes:     20,
               types: []systems.RealityType{
                   systems.BASE_REALITY,
                   systems.INFORMATION_DENSE,
               },
           },
           {
               name:      "Quantum",
               realities: 4,
               nodes:     15,
               types: []systems.RealityType{
                   systems.BASE_REALITY,
                   systems.QUANTUM_DIVERGENT,
                   systems.CAUSAL_BROKEN,
                   systems.VOID_REALITY,
               },
           },
           {
               name:      "Conscious",
               realities: 3,
               nodes:     30,
               types: []systems.RealityType{
                   systems.BASE_REALITY,
                   systems.CONSCIOUSNESS_AMPLIFIED,
                   systems.SINGULARITY_REALITY,
               },
           },
       }
       
       // Initialize networks in parallel
       var wg sync.WaitGroup
       
       for i, cfg := range configs {
           wg.Add(1)
           go func(idx int, config struct {
               name       string
               realities  int
               nodes      int
               types      []systems.RealityType
           }) {
               defer wg.Done()
               
               networks[idx] = systems.NewMultiRealityNeuralNetwork(
                   config.realities, config.nodes,
               )
               
               networks[idx].Initialize(systems.RealityNetworkConfig{
                   RealityTypes: config.types,
                   ConnectionProbability: 0.25,
               })
               
               // Evolve
               for cycle := 0; cycle < 10; cycle++ {
                   networks[idx].Evolve()
               }
               
               metrics := networks[idx].GetMetrics()
               fmt.Printf("%s: Consciousness=%.3f, Coherence=%.3f\n",
                   config.name, metrics.GlobalConsciousness, metrics.CrossRealityCoherence)
           }(i, cfg)
       }
       
       wg.Wait()
       fmt.Println("\nAll networks evolved successfully")
   }

Reality Analysis
----------------

.. code-block:: go

   package main

   import (
       "fmt"
       "sort"
       "neuralblitz/pkg/systems"
   )

   type RealityAnalysis struct {
       Type               systems.RealityType
       Consciousness      float64
       InformationDensity float64
       QuantumCoherence   float64
       Performance        float64
   }

   func analyzeRealities(mrnn *systems.MultiRealityNeuralNetwork) []RealityAnalysis {
       realities := mrnn.GetRealities()
       analyses := make([]RealityAnalysis, len(realities))
       
       for i, reality := range realities {
           analyses[i] = RealityAnalysis{
               Type:               reality.Type,
               Consciousness:      reality.ConsciousnessLevel,
               InformationDensity: reality.InformationDensity,
               QuantumCoherence:   reality.QuantumCoherence,
               Performance:        reality.InformationDensity * reality.QuantumCoherence,
           }
       }
       
       // Sort by performance
       sort.Slice(analyses, func(i, j int) bool {
           return analyses[i].Performance > analyses[j].Performance
       })
       
       return analyses
   }

   func main() {
       mrnn := systems.NewMultiRealityNeuralNetwork(8, 20)
       mrnn.Initialize(systems.RealityNetworkConfig{
           RealityTypes: []systems.RealityType{
               systems.BASE_REALITY,
               systems.QUANTUM_DIVERGENT,
               systems.TEMPORAL_INVERTED,
               systems.ENTROPIC_REVERSED,
               systems.CONSCIOUSNESS_AMPLIFIED,
               systems.DIMENSIONAL_SHIFTED,
               systems.INFORMATION_DENSE,
               systems.SINGULARITY_REALITY,
           },
       })
       
       // Run evolution
       for cycle := 0; cycle < 50; cycle++ {
           mrnn.Evolve()
       }
       
       // Analyze realities
       analyses := analyzeRealities(mrnn)
       
       fmt.Println("Reality Performance Analysis:")
       fmt.Println("=============================")
       
       for i, analysis := range analyses {
           fmt.Printf("%d. %s\n", i+1, realityTypeString(analysis.Type))
           fmt.Printf("   Consciousness: %.3f\n", analysis.Consciousness)
           fmt.Printf("   Information Density: %.2f\n", analysis.InformationDensity)
           fmt.Printf("   Quantum Coherence: %.4f\n", analysis.QuantumCoherence)
           fmt.Printf("   Performance Score: %.4f\n", analysis.Performance)
           fmt.Println()
       }
   }

   func realityTypeString(t systems.RealityType) string {
       names := []string{
           "Base Reality",
           "Quantum Divergent",
           "Temporal Inverted",
           "Entropic Reversed",
           "Consciousness Amplified",
           "Dimensional Shifted",
           "Causal Broken",
           "Information Dense",
           "Void Reality",
           "Singularity Reality",
       }
       if int(t) < len(names) {
           return names[t]
       }
       return "Unknown"
   }

Cross-Reality Communication
---------------------------

.. code-block:: go

   package main

   import (
       "fmt"
       "time"
       "neuralblitz/pkg/systems"
   )

   func main() {
       mrnn := systems.NewMultiRealityNeuralNetwork(4, 15)
       mrnn.Initialize(systems.RealityNetworkConfig{
           RealityTypes: []systems.RealityType{
               systems.BASE_REALITY,
               systems.QUANTUM_DIVERGENT,
               systems.CONSCIOUSNESS_AMPLIFIED,
               systems.INFORMATION_DENSE,
           },
       })
       
       // Create signals between realities
       fmt.Println("Creating cross-reality signals...")
       
       mrnn.CreateSignal(systems.CrossRealitySignal{
           SourceReality:      0, // Base
           TargetReality:      1, // Quantum
           SignalType:         systems.SignalQuantum,
           Amplitude:          0.8,
           EntanglementType:   systems.EntanglementQuantum,
       })
       
       mrnn.CreateSignal(systems.CrossRealitySignal{
           SourceReality:      1, // Quantum
           TargetReality:      2, // Consciousness
           SignalType:         systems.SignalConsciousness,
           Amplitude:          0.6,
           EntanglementType:   systems.EntanglementConsciousness,
       })
       
       mrnn.CreateSignal(systems.CrossRealitySignal{
           SourceReality:      2, // Consciousness
           TargetReality:      3, // Information
           SignalType:         systems.SignalInformation,
           Amplitude:          0.9,
           EntanglementType:   systems.EntanglementInformational,
       })
       
       // Evolve with signals
       fmt.Println("Evolving with cross-reality communication...")
       
       start := time.Now()
       
       for cycle := 0; cycle < 30; cycle++ {
           mrnn.Evolve()
           mrnn.ProcessSignals()
           
           if cycle % 10 == 0 {
               metrics := mrnn.GetMetrics()
               activeSignals := len(mrnn.GetActiveSignals())
               
               fmt.Printf("Cycle %d: Signals=%d, Coherence=%.3f, Time=%.2fms\n",
                   cycle, activeSignals, metrics.CrossRealityCoherence,
                   float64(time.Since(start).Milliseconds())/float64(cycle+1))
           }
       }
       
       finalMetrics := mrnn.GetMetrics()
       fmt.Printf("\nFinal: Global Consciousness=%.4f, Coherence=%.4f\n",
           finalMetrics.GlobalConsciousness, finalMetrics.CrossRealityCoherence)
   }

Running the Examples
--------------------

Save as ``main.go`` and run:

.. code-block:: bash

   go run main.go

Or build:

.. code-block:: bash

   go build -o multi_reality main.go
   ./multi_reality

Expected Output
---------------

.. code-block:: text

   Evolving multi-reality network...
   Cycle 0: Consciousness=0.382, Coherence=0.845
   Cycle 5: Consciousness=0.487, Coherence=0.892
   Cycle 10: Consciousness=0.534, Coherence=0.923
   Cycle 15: Consciousness=0.567, Coherence=0.945
   
   Final Results:
   Global Consciousness: 0.5892
   Cross-Reality Coherence: 0.9583
   Information Flow Rate: 47.23 bits/s

Performance Characteristics
---------------------------

- **Base Reality**: Standard neural processing
- **Quantum Divergent**: High uncertainty, parallel paths
- **Temporal Inverted**: Reversed causality, predictive processing
- **Consciousness Amplified**: High awareness, ethical constraints
- **Information Dense**: High-capacity storage and retrieval
- **Singularity Reality**: Maximum density, extreme processing

Next Steps
----------

- See :doc:`/examples/quantum_neurons` for quantum neuron examples
- See :doc:`/examples/consciousness` for consciousness integration
- See :doc:`/examples/agents` for autonomous agent examples
