Systems API Reference
=====================

This document provides the complete API reference for the ``pkg/systems``
package, which contains the core AI and cognitive modules of NeuralBlitz
v50 Go.

.. contents::
   :local:
   :depth: 2

Agent Framework
---------------

The agent framework provides capabilities for creating and managing
autonomous AI agents.

Types
~~~~~

CapabilityType
^^^^^^^^^^^^^^

.. code-block:: go

   type CapabilityType int

   const (
       CapabilityPerception CapabilityType = iota
       CapabilityReasoning
       CapabilityPlanning
       CapabilityActionExecution
       CapabilityLearning
       CapabilityCommunication
       CapabilitySelfModification
       CapabilityEthicalReasoning
       CapabilityCreativity
       CapabilityMemory
       CapabilityAdaptation
       CapabilityMetacognition
   )

Enumeration of agent capabilities.

Functions
~~~~~~~~~

NewAgentCapabilities
^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func NewAgentCapabilities() *AgentCapabilities

Creates a new AgentCapabilities instance with all capabilities initialized
to disabled state.

**Returns:**
   - ``*AgentCapabilities``: New capabilities instance

**Example:**

.. code-block:: go

   caps := systems.NewAgentCapabilities()
   caps.EnableCapability(systems.CapabilityPerception)
   caps.SetLevel(systems.CapabilityPerception, 0.8)

NewAdvancedAutonomousAgent
^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func NewAdvancedAutonomousAgent(agentID string) *AdvancedAutonomousAgent

Creates a new autonomous agent with the specified ID.

**Parameters:**
   - ``agentID`` (string): Unique identifier for the agent

**Returns:**
   - ``*AdvancedAutonomousAgent``: New agent instance

**Example:**

.. code-block:: go

   agent := systems.NewAdvancedAutonomousAgent("agent_001")
   agent.Status = systems.StatusActive

NewAgentRegistry
^^^^^^^^^^^^^^^^

.. code-block:: go

   func NewAgentRegistry() *AgentRegistry

Creates a new agent registry for managing multiple agents.

**Returns:**
   - ``*AgentRegistry``: New registry instance

Methods
~~~~~~~

AgentCapabilities.EnableCapability
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func (a *AgentCapabilities) EnableCapability(cap CapabilityType) error

Enables a specific capability.

**Parameters:**
   - ``cap`` (CapabilityType): Capability to enable

**Returns:**
   - ``error``: Error if capability invalid

AgentCapabilities.SetLevel
^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func (a *AgentCapabilities) SetLevel(cap CapabilityType, level float64)

Sets the proficiency level for a capability (0.0 to 1.0).

**Parameters:**
   - ``cap`` (CapabilityType): Capability to set
   - ``level`` (float64): Proficiency level

AgentCapabilities.IsEnabled
^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func (a *AgentCapabilities) IsEnabled(cap CapabilityType) bool

Checks if a capability is enabled.

**Parameters:**
   - ``cap`` (CapabilityType): Capability to check

**Returns:**
   - ``bool``: True if enabled

Autonomous Self-Evolution
-------------------------

Provides self-improving capabilities for code analysis and generation.

Types
~~~~~

EvolutionConfig
^^^^^^^^^^^^^^^

.. code-block:: go

   type EvolutionConfig struct {
       Enabled        bool
       EvolutionID    string
       MaxIterations  int
       SafetyLevel    int
   }

Configuration for autonomous evolution.

Functions
~~~~~~~~~

NewAutonomousSelfEvolution
^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func NewAutonomousSelfEvolution() *AutonomousSelfEvolution

Creates a new autonomous self-evolution system.

**Returns:**
   - ``*AutonomousSelfEvolution``: New evolution system

Methods
~~~~~~~

AutonomousSelfEvolution.Evolve
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func (ase *AutonomousSelfEvolution) Evolve(code string) (*EvolutionResult, error)

Analyzes and evolves the provided code.

**Parameters:**
   - ``code`` (string): Source code to evolve

**Returns:**
   - ``*EvolutionResult``: Evolution results
   - ``error``: Error if evolution fails

Code Generation
---------------

AI-powered code generation with syntax analysis and validation.

Types
~~~~~

CodeGenConfig
^^^^^^^^^^^^^

.. code-block:: go

   type CodeGenConfig struct {
       GenerationID   string
       Enabled        bool
       TargetLanguage string
       SafetyChecks   bool
   }

Configuration for code generation.

Functions
~~~~~~~~~

NewSelfImprovingCodeGenerator
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func NewSelfImprovingCodeGenerator(config CodeGenConfig) *SelfImprovingCodeGenerator

Creates a new code generator with the specified configuration.

**Parameters:**
   - ``config`` (CodeGenConfig): Generator configuration

**Returns:**
   - ``*SelfImprovingCodeGenerator``: New generator instance

Consciousness Integration
-------------------------

Hierarchical consciousness models for AI agents.

Types
~~~~~

ConsciousnessLevel
^^^^^^^^^^^^^^^^^^

.. code-block:: go

   type ConsciousnessLevel int

   const (
       ConsciousnessIndividual ConsciousnessLevel = iota
       ConsciousnessCollective
       ConsciousnessPlanetary
       ConsciousnessSolar
       ConsciousnessGalactic
       ConsciousnessUniversal
       ConsciousnessMultiversal
       ConsciousnessAbsolute
   )

Enumeration of consciousness levels.

IntegrationMode
^^^^^^^^^^^^^^^

.. code-block:: go

   type IntegrationMode int

   const (
       IntegrationObservation IntegrationMode = iota
       IntegrationSynchronization
       IntegrationAmplification
       IntegrationCoCreation
       IntegrationTranscendence
   )

Enumeration of integration modes.

Functions
~~~~~~~~~

NewConsciousnessIntegration
^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func NewConsciousnessIntegration(
       id string,
       participants []string,
       mode IntegrationMode,
   ) *ConsciousnessIntegration

Creates a new consciousness integration system.

**Parameters:**
   - ``id`` (string): Integration ID
   - ``participants`` ([]string): List of participant IDs
   - ``mode`` (IntegrationMode): Integration mode

**Returns:**
   - ``*ConsciousnessIntegration``: New integration instance

NewConsciousnessState
^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func NewConsciousnessState(
       id string,
       currentLevel ConsciousnessLevel,
       targetLevel ConsciousnessLevel,
       mode IntegrationMode,
   ) *ConsciousnessState

Creates a new consciousness state.

**Parameters:**
   - ``id`` (string): State ID
   - ``currentLevel`` (ConsciousnessLevel): Current consciousness level
   - ``targetLevel`` (ConsciousnessLevel): Target consciousness level
   - ``mode`` (IntegrationMode): Integration mode

**Returns:**
   - ``*ConsciousnessState``: New state instance

Cross-Reality Entanglement
--------------------------

Quantum-like entanglement between different reality states.

Types
~~~~~

EntanglementType
^^^^^^^^^^^^^^^^

.. code-block:: go

   type EntanglementType int

   const (
       EntanglementSpatial EntanglementType = iota
       EntanglementTemporal
       EntanglementCausal
       EntanglementInformational
       EntanglementConsciousness
       EntanglementEmotional
       EntanglementPurpose
       EntanglementTranscendent
   )

Enumeration of entanglement types.

EntanglementStrength
^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   type EntanglementStrength int

   const (
       StrengthWeak EntanglementStrength = iota
       StrengthModerate
       StrengthStrong
   )

Enumeration of entanglement strengths.

Functions
~~~~~~~~~

NewQuantumEntanglementSystem
^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func NewQuantumEntanglementSystem(id string) *QuantumEntanglementSystem

Creates a new quantum entanglement system.

**Parameters:**
   - ``id`` (string): System ID

**Returns:**
   - ``*QuantumEntanglementSystem``: New entanglement system

NewRealityBridge
^^^^^^^^^^^^^^^^

.. code-block:: go

   func NewRealityBridge(
       id, realityA, realityB string,
       entType EntanglementType,
       strength EntanglementStrength,
   ) *RealityBridge

Creates a new reality bridge.

**Parameters:**
   - ``id`` (string): Bridge ID
   - ``realityA`` (string): First reality ID
   - ``realityB`` (string): Second reality ID
   - ``entType`` (EntanglementType): Type of entanglement
   - ``strength`` (EntanglementStrength): Entanglement strength

**Returns:**
   - ``*RealityBridge``: New bridge instance

NewEntanglementState
^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func NewEntanglementState(
       id, bridgeID string,
       entType EntanglementType,
       numQubits int,
   ) *EntanglementState

Creates a new entanglement state.

**Parameters:**
   - ``id`` (string): State ID
   - ``bridgeID`` (string): Associated bridge ID
   - ``entType`` (EntanglementType): Type of entanglement
   - ``numQubits`` (int): Number of qubits

**Returns:**
   - ``*EntanglementState``: New state instance

Dimensional Neural Processing
-----------------------------

Neural networks operating in 11 dimensions.

Types
~~~~~

DimensionalConfig
^^^^^^^^^^^^^^^^^

.. code-block:: go

   type DimensionalConfig struct {
       NetworkID      string
       Enabled        bool
       NumDimensions  int
       SemanticDepth  int
   }

Configuration for dimensional neural processing.

Functions
~~~~~~~~~

NewDimensionalNeuralProcessor
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func NewDimensionalNeuralProcessor(config DimensionalConfig) *DimensionalNeuralProcessor

Creates a new dimensional neural processor.

**Parameters:**
   - ``config`` (DimensionalConfig): Processor configuration

**Returns:**
   - ``*DimensionalNeuralProcessor``: New processor instance

Emergent Purpose Discovery
--------------------------

Systems for discovering and evolving purposes.

Functions
~~~~~~~~~

NewPurposeDiscovery
^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func NewPurposeDiscovery() *PurposeDiscovery

Creates a new purpose discovery system.

**Returns:**
   - ``*PurposeDiscovery``: New discovery system

Multi-Reality Neural Networks
-----------------------------

Neural networks operating across multiple realities.

Types
~~~~~

RealityType
^^^^^^^^^^^

.. code-block:: go

   type RealityType int

   const (
       BASE_REALITY RealityType = iota
       QUANTUM_DIVERGENT
       TEMPORAL_INVERTED
       ENTROPIC_REVERSED
       CONSCIOUSNESS_AMPLIFIED
       DIMENSIONAL_SHIFTED
       CAUSAL_BROKEN
       INFORMATION_DENSE
       VOID_REALITY
       SINGULARITY_REALITY
   )

Enumeration of reality types.

Functions
~~~~~~~~~

NewMultiRealityNeuralNetwork
^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func NewMultiRealityNeuralNetwork(
       numRealities int,
       nodesPerReality int,
   ) *MultiRealityNeuralNetwork

Creates a new multi-reality neural network.

**Parameters:**
   - ``numRealities`` (int): Number of realities
   - ``nodesPerReality`` (int): Nodes per reality

**Returns:**
   - ``*MultiRealityNeuralNetwork``: New network instance

Neuro-BCI
---------

Brain-computer interface inspired neural processing.

Types
~~~~~

BCIBackend
^^^^^^^^^^

.. code-block:: go

   type BCIBackend struct {
       ID             string
       Status         BCIStatus
       UseSimulator   bool
       SamplingRate   int
   }

BCI backend structure.

Functions
~~~~~~~~~

NewBCIBackend
^^^^^^^^^^^^^

.. code-block:: go

   func NewBCIBackend(useSimulator bool, samplingRate int) *BCIBackend

Creates a new BCI backend.

**Parameters:**
   - ``useSimulator`` (bool): Use simulator instead of real hardware
   - ``samplingRate`` (int): Sampling rate in Hz

**Returns:**
   - ``*BCIBackend``: New BCI backend

NewEEGSimulator
^^^^^^^^^^^^^^^

.. code-block:: go

   func NewEEGSimulator(samplingRate int, numChannels int) *EEGSimulator

Creates a new EEG simulator.

**Parameters:**
   - ``samplingRate`` (int): Sampling rate in Hz
   - ``numChannels`` (int): Number of EEG channels

**Returns:**
   - ``*EEGSimulator``: New EEG simulator

Quantum Spiking Neurons
-----------------------

Neural networks with quantum-inspired dynamics.

Types
~~~~~

NeuronConfiguration
^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   type NeuronConfiguration struct {
       RestingPotential     float64
       ThresholdPotential   float64
       MembraneResistance   float64
       MembraneTimeConstant float64
       MaxHistory          int
   }

Configuration for quantum spiking neurons.

Functions
~~~~~~~~~

NewQuantumSpikingNeuron
^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func NewQuantumSpikingNeuron(neuronID string) *QuantumSpikingNeuron

Creates a new quantum spiking neuron.

**Parameters:**
   - ``neuronID`` (string): Unique neuron identifier

**Returns:**
   - ``*QuantumSpikingNeuron``: New neuron instance

Methods
~~~~~~~

QuantumSpikingNeuron.Initialize
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func (qsn *QuantumSpikingNeuron) Initialize(config NeuronConfiguration) error

Initializes the neuron with configuration.

**Parameters:**
   - ``config`` (NeuronConfiguration): Neuron configuration

**Returns:**
   - ``error``: Error if initialization fails

QuantumSpikingNeuron.Step
^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func (qsn *QuantumSpikingNeuron) Step(inputCurrent float64)

Advances the neuron simulation by one time step.

**Parameters:**
   - ``inputCurrent`` (float64): Input current in nA

NeurochemicalState
------------------

Functions
~~~~~~~~~

NewNeurochemicalState
^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func NewNeurochemicalState() *NeurochemicalState

Creates a new neurochemical state.

**Returns:**
   - ``*NeurochemicalState``: New neurochemical state

CognitiveState
--------------

Functions
~~~~~~~~~

NewCognitiveState
^^^^^^^^^^^^^^^^^

.. code-block:: go

   func NewCognitiveState() *CognitiveState

Creates a new cognitive state.

**Returns:**
   - ``*CognitiveState``: New cognitive state
