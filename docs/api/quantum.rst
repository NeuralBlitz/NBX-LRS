Quantum API Reference
=====================

This document provides the complete API reference for the ``pkg/quantum``
package, which contains quantum-inspired computational models and
infrastructure.

.. contents::
   :local:
   :depth: 2

Quantum Foundation
------------------

Core quantum infrastructure including agents, communication layers, and
reality simulators.

Types
~~~~~

QuantumState
^^^^^^^^^^^^

.. code-block:: go

   type QuantumState int

   const (
       StateDORMANT QuantumState = iota
       StateAWARE
       StateFOCUSED
       StateTRANSCENDENT
       StateSINGULARITY
   )

Enumeration of quantum states.

QuantumAgent
^^^^^^^^^^^^

.. code-block:: go

   type QuantumAgent struct {
       AgentID            string
       ConsciousnessLevel QuantumState
       CoherenceFactor    float64
       EntangledPartners  []string
       QuantumMemory      []complex128
   }

Represents a quantum agent.

QuantumCommunicationLayer
^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   type QuantumCommunicationLayer struct {
       NumQubits          int
       QuantumAgents      map[string]*QuantumAgent
       EntanglementMatrix map[string]map[string]float64
       mu                 sync.RWMutex
   }

Manages quantum communication between agents.

QuantumRealitySimulator
^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   type QuantumRealitySimulator struct {
       NumRealities         int
       RealityStates        map[int][]complex128
       RealityCouplings     map[int]map[int]float64
       QuantumStateHistory  [][]complex128
       mu                   sync.RWMutex
   }

Simulates multiple quantum realities.

QuantumFoundation
^^^^^^^^^^^^^^^^^

.. code-block:: go

   type QuantumFoundation struct {
       CommunicationLayer *QuantumCommunicationLayer
       KeyDistribution    *QuantumKeyDistribution
       RealitySimulator   *QuantumRealitySimulator
       Initialized        bool
       InitializationTime time.Time
       mu                 sync.RWMutex
   }

Complete quantum foundation infrastructure.

Functions
~~~~~~~~~

NewQuantumAgent
^^^^^^^^^^^^^^^

.. code-block:: go

   func NewQuantumAgent(agentID string, state QuantumState) *QuantumAgent

Creates a new quantum agent.

**Parameters:**
   - ``agentID`` (string): Unique agent identifier
   - ``state`` (QuantumState): Initial quantum state

**Returns:**
   - ``*QuantumAgent``: New quantum agent

**Example:**

.. code-block:: go

   agent := quantum.NewQuantumAgent("alice", quantum.StateAWARE)
   agent.CoherenceFactor = 0.95

NewQuantumCommunicationLayer
^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func NewQuantumCommunicationLayer(numQubits int) *QuantumCommunicationLayer

Creates a new quantum communication layer.

**Parameters:**
   - ``numQubits`` (int): Number of qubits in the system

**Returns:**
   - ``*QuantumCommunicationLayer``: New communication layer

**Example:**

.. code-block:: go

   qcl := quantum.NewQuantumCommunicationLayer(8)
   agent := qcl.CreateQuantumAgent("bob", quantum.StateFOCUSED)

NewQuantumRealitySimulator
^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func NewQuantumRealitySimulator(numRealities int) *QuantumRealitySimulator

Creates a new quantum reality simulator.

**Parameters:**
   - ``numRealities`` (int): Number of realities to simulate

**Returns:**
   - ``*QuantumRealitySimulator``: New reality simulator

**Example:**

.. code-block:: go

   qrs := quantum.NewQuantumRealitySimulator(4)
   qrs.InitializeMultiverse()

NewQuantumFoundation
^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func NewQuantumFoundation(numQubits, numRealities int) *QuantumFoundation

Creates a complete quantum foundation.

**Parameters:**
   - ``numQubits`` (int): Number of qubits
   - ``numRealities`` (int): Number of realities

**Returns:**
   - ``*QuantumFoundation``: New quantum foundation

**Example:**

.. code-block:: go

   qf := quantum.NewQuantumFoundation(8, 4)
   err := qf.Initialize()
   if err != nil {
       log.Fatal(err)
   }

Methods
~~~~~~~

QuantumCommunicationLayer.CreateQuantumAgent
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func (qcl *QuantumCommunicationLayer) CreateQuantumAgent(
       agentID string,
       state QuantumState,
   ) *QuantumAgent

Creates a new quantum agent in the communication layer.

**Parameters:**
   - ``agentID`` (string): Unique agent identifier
   - ``state`` (QuantumState): Initial quantum state

**Returns:**
   - ``*QuantumAgent``: Created agent

QuantumCommunicationLayer.CreateEntanglement
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func (qcl *QuantumCommunicationLayer) CreateEntanglement(
       agentID1, agentID2 string,
   ) bool

Creates quantum entanglement between two agents.

**Parameters:**
   - ``agentID1`` (string): First agent ID
   - ``agentID2`` (string): Second agent ID

**Returns:**
   - ``bool``: True if entanglement created successfully

QuantumCommunicationLayer.QuantumTeleportation
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func (qcl *QuantumCommunicationLayer) QuantumTeleportation(
       senderID, receiverID string,
       messageState []float64,
   ) *TeleportationResult

Performs quantum teleportation of a message state.

**Parameters:**
   - ``senderID`` (string): Sender agent ID
   - ``receiverID`` (string): Receiver agent ID
   - ``messageState`` ([]float64): Quantum state to teleport

**Returns:**
   - ``*TeleportationResult``: Teleportation results

QuantumCommunicationLayer.CalculateGlobalCoherence
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func (qcl *QuantumCommunicationLayer) CalculateGlobalCoherence() float64

Calculates the global coherence of all agents.

**Returns:**
   - ``float64``: Global coherence value (0.0 to 1.0)

QuantumRealitySimulator.InitializeMultiverse
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func (qrs *QuantumRealitySimulator) InitializeMultiverse()

Initializes the multiverse with uniform superposition.

QuantumRealitySimulator.SimulateRealityInterference
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func (qrs *QuantumRealitySimulator) SimulateRealityInterference(
       realityA, realityB int,
   ) float64

Simulates interference between two realities.

**Parameters:**
   - ``realityA`` (int): First reality index
   - ``realityB`` (int): Second reality index

**Returns:**
   - ``float64``: Interference value

QuantumRealitySimulator.CollapseToReality
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func (qrs *QuantumRealitySimulator) CollapseToReality(
       agentID string,
       realityIndex int,
   ) bool

Collapses an agent's state to a specific reality.

**Parameters:**
   - ``agentID`` (string): Agent ID
   - ``realityIndex`` (int): Reality index to collapse to

**Returns:**
   - ``bool``: True if collapse successful

QuantumRealitySimulator.GetSuperpositionState
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func (qrs *QuantumRealitySimulator) GetSuperpositionState() []complex128

Gets the current superposition state of all realities.

**Returns:**
   - ``[]complex128``: Probability amplitudes

QuantumFoundation.Initialize
^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func (qf *QuantumFoundation) Initialize() error

Initializes the complete quantum foundation.

**Returns:**
   - ``error``: Error if initialization fails

QuantumFoundation.GetStats
^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func (qf *QuantumFoundation) GetStats() *QuantumFoundationStats

Gets statistics about the quantum foundation.

**Returns:**
   - ``*QuantumFoundationStats``: Foundation statistics

QuantumFoundation.EncryptMessage
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func (qf *QuantumFoundation) EncryptMessage(
       sender, receiver, message string,
   ) (*SecureMessage, error)

Encrypts a message using quantum key distribution.

**Parameters:**
   - ``sender`` (string): Sender ID
   - ``receiver`` (string): Receiver ID
   - ``message`` (string): Message to encrypt

**Returns:**
   - ``*SecureMessage``: Encrypted message
   - ``error``: Error if encryption fails

QuantumFoundation.DecryptMessage
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func (qf *QuantumFoundation) DecryptMessage(
       msg *SecureMessage,
   ) (string, error)

Decrypts a secure message.

**Parameters:**
   - ``msg`` (*SecureMessage): Message to decrypt

**Returns:**
   - ``string``: Decrypted message
   - ``error``: Error if decryption fails

Quantum Cryptography
--------------------

Quantum-resistant cryptography and secure communication.

Types
~~~~~

QuantumKeyDistribution
^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   type QuantumKeyDistribution struct {
       CommunicationLayer *QuantumCommunicationLayer
       KeySize            int
       SharedKeys         map[string][]byte
       mu                 sync.RWMutex
   }

Manages quantum key distribution.

SecureMessage
^^^^^^^^^^^^^

.. code-block:: go

   type SecureMessage struct {
       Sender     string
       Receiver   string
       Ciphertext []byte
       Nonce      []byte
       Timestamp  time.Time
   }

Represents an encrypted message.

TeleportationResult
^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   type TeleportationResult struct {
       Success     bool
       Fidelity    float64
       ElapsedTime time.Duration
   }

Results from quantum teleportation.

Functions
~~~~~~~~~

NewQuantumKeyDistribution
^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func NewQuantumKeyDistribution(
       qcl *QuantumCommunicationLayer,
   ) *QuantumKeyDistribution

Creates a new quantum key distribution system.

**Parameters:**
   - ``qcl`` (*QuantumCommunicationLayer): Communication layer

**Returns:**
   - ``*QuantumKeyDistribution``: New QKD system

Methods
~~~~~~~

QuantumKeyDistribution.GenerateQuantumKey
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func (qkd *QuantumKeyDistribution) GenerateQuantumKey(
       agentA, agentB string,
   ) ([]byte, error)

Generates a shared quantum key between two agents.

**Parameters:**
   - ``agentA`` (string): First agent ID
   - ``agentB`` (string): Second agent ID

**Returns:**
   - ``[]byte``: Shared key
   - ``error``: Error if generation fails

QuantumKeyDistribution.GetSharedKey
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func (qkd *QuantumKeyDistribution) GetSharedKey(
       agentA, agentB string,
   ) ([]byte, bool)

Retrieves a previously generated shared key.

**Parameters:**
   - ``agentA`` (string): First agent ID
   - ``agentB`` (string): Second agent ID

**Returns:**
   - ``[]byte``: Shared key
   - ``bool``: True if key exists

Utility Functions
~~~~~~~~~~~~~~~~~

HashQuantumState
^^^^^^^^^^^^^^^^

.. code-block:: go

   func HashQuantumState(state []float64) string

Hashes a quantum state using SHA-256.

**Parameters:**
   - ``state`` ([]float64): Quantum state vector

**Returns:**
   - ``string``: Hex-encoded hash

Quantum Neural Networks
-----------------------

Quantum-inspired machine learning.

Types
~~~~~

QuantumNeuralNetwork
^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   type QuantumNeuralNetwork struct {
       InputSize        int
       HiddenSize       int
       OutputSize       int
       LearningRate     float64
       QuantumAmplitudes []float64
       Weights          [][]float64
   }

Quantum neural network structure.

Functions
~~~~~~~~~

NewQuantumNeuralNetwork
^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func NewQuantumNeuralNetwork(
       inputSize, hiddenSize, outputSize int,
   ) *QuantumNeuralNetwork

Creates a new quantum neural network.

**Parameters:**
   - ``inputSize`` (int): Input layer size
   - ``hiddenSize`` (int): Hidden layer size
   - ``outputSize`` (int): Output layer size

**Returns:**
   - ``*QuantumNeuralNetwork``: New quantum neural network

Methods
~~~~~~~

QuantumNeuralNetwork.QuantumForwardPass
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func (qnn *QuantumNeuralNetwork) QuantumForwardPass(
       input []float64,
   ) []float64

Performs a quantum forward pass.

**Parameters:**
   - ``input`` ([]float64): Input vector

**Returns:**
   - ``[]float64``: Output vector

QuantumNeuralNetwork.QuantumTrain
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func (qnn *QuantumNeuralNetwork) QuantumTrain(
       inputs [][]float64,
       targets [][]float64,
       epochs int,
       tolerance float64,
   ) error

Trains the quantum neural network.

**Parameters:**
   - ``inputs`` ([][]float64): Training inputs
   - ``targets`` ([][]float64): Training targets
   - ``epochs`` (int): Maximum training epochs
   - ``tolerance`` (float64): Convergence tolerance

**Returns:**
   - ``error``: Error if training fails

Agent State Methods
~~~~~~~~~~~~~~~~~~~

QuantumAgent.GetAgentState
^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: go

   func (qa *QuantumAgent) GetAgentState() (string, error)

Gets the JSON representation of the agent state.

**Returns:**
   - ``string``: JSON state string
   - ``error``: Error if serialization fails

Error Types
-----------

Common errors returned by quantum operations.

.. code-block:: go

   var (
       ErrAgentNotFound          = errors.New("agent not found")
       ErrEntanglementFailed     = errors.New("entanglement creation failed")
       ErrTeleportationFailed    = errors.New("quantum teleportation failed")
       ErrKeyGenerationFailed    = errors.New("quantum key generation failed")
       ErrEncryptionFailed       = errors.New("encryption failed")
       ErrDecryptionFailed       = errors.New("decryption failed")
       ErrInvalidReality         = errors.New("invalid reality index")
       ErrAlreadyInitialized     = errors.New("already initialized")
   )
