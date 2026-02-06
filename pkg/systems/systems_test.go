package systems

import (
	"testing"
)

// TestCapabilityTypes tests basic capability types
func TestCapabilityTypes(t *testing.T) {
	caps := NewAgentCapabilities()
	if caps == nil {
		t.Fatal("Failed to create AgentCapabilities")
	}

	// Test perception capability
	caps.SetLevel(CapabilityPerception, 0.8)
	caps.EnableCapability(CapabilityPerception)
	if !caps.IsEnabled(CapabilityPerception) {
		t.Error("Perception should be enabled")
	}
	if caps.GetLevel(CapabilityPerception) != 0.8 {
		t.Errorf("Expected level 0.8, got %f", caps.GetLevel(CapabilityPerception))
	}
}

// TestBasicAgent tests basic agent creation
func TestBasicAgent(t *testing.T) {
	agent := NewAdvancedAutonomousAgent("test_agent")
	if agent.AgentID != "test_agent" {
		t.Errorf("Expected agent ID 'test_agent', got '%s'", agent.AgentID)
	}
}

// TestSelfImprovingCodeGenerator tests code generation
func TestSelfImprovingCodeGenerator(t *testing.T) {
	gen := NewSelfImprovingCodeGenerator(CodeGenConfig{
		GenerationID: "test_gen",
		Enabled:      true,
	})
	if gen == nil {
		t.Fatal("Failed to create code generator")
	}
}

// TestPurposeDiscovery tests purpose discovery
func TestPurposeDiscovery(t *testing.T) {
	pd := NewPurposeDiscovery()
	if pd == nil {
		t.Fatal("Failed to create purpose discovery")
	}
}

// TestAutonomousSelfEvolution tests self evolution
func TestAutonomousSelfEvolution(t *testing.T) {
	evo := NewAutonomousSelfEvolution()
	if evo == nil {
		t.Fatal("Failed to create self evolution")
	}
}

// TestMultiRealityNeuralNetwork tests multi-reality network
func TestMultiRealityNeuralNetwork(t *testing.T) {
	mrnn := NewMultiRealityNeuralNetwork(4, 10)
	if mrnn == nil {
		t.Fatal("Failed to create multi-reality network")
	}
}

// TestDimensionalNeuralProcessor tests dimensional processing
func TestDimensionalNeuralProcessor(t *testing.T) {
	proc := NewDimensionalNeuralProcessor(DimensionalConfig{
		NetworkID: "test_dim",
		Enabled:   true,
	})
	if proc == nil {
		t.Fatal("Failed to create dimensional processor")
	}
}

// TestNeuroBCI tests neuro-BCI
func TestNeuroBCI(t *testing.T) {
	bci := NewBCIBackend(false, 256)
	if bci == nil {
		t.Fatal("Failed to create neuro-BCI")
	}
}

// TestConsciousnessIntegration tests consciousness integration
func TestConsciousnessIntegration(t *testing.T) {
	participants := []string{"alice", "bob"}
	ci := NewConsciousnessIntegration("test_ci", participants, IntegrationObservation)
	if ci == nil {
		t.Fatal("Failed to create consciousness integration")
	}
}

// TestCrossRealityEntanglement tests cross-reality entanglement
func TestCrossRealityEntanglement(t *testing.T) {
	ent := NewQuantumEntanglementSystem("test_ent")
	if ent == nil {
		t.Fatal("Failed to create cross-reality entanglement")
	}
}

// TestQuantumSpikingNeuron tests quantum spiking neuron
func TestQuantumSpikingNeuron(t *testing.T) {
	neuron := NewQuantumSpikingNeuron("test_neuron")
	if neuron == nil {
		t.Fatal("Failed to create quantum spiking neuron")
	}
}

// TestConsciousnessState tests consciousness state
func TestConsciousnessState(t *testing.T) {
	state := NewConsciousnessState("test_state", ConsciousnessIndividual, ConsciousnessCollective, IntegrationObservation)
	if state == nil {
		t.Fatal("Failed to create consciousness state")
	}
}

// TestRealityBridge tests reality bridge
func TestRealityBridge(t *testing.T) {
	bridge := NewRealityBridge("bridge_1", "reality_a", "reality_b", EntanglementSpatial, StrengthStrong)
	if bridge == nil {
		t.Fatal("Failed to create reality bridge")
	}
}

// TestEntanglementState tests entanglement state
func TestEntanglementState(t *testing.T) {
	state := NewEntanglementState("state_1", "bridge_1", EntanglementTemporal, 8)
	if state == nil {
		t.Fatal("Failed to create entanglement state")
	}
}

// TestAgentRegistry tests agent registry
func TestAgentRegistry(t *testing.T) {
	registry := NewAgentRegistry()
	if registry == nil {
		t.Fatal("Failed to create agent registry")
	}
}

// TestEEGSimulator tests EEG simulator
func TestEEGSimulator(t *testing.T) {
	sim := NewEEGSimulator(256, 4)
	if sim == nil {
		t.Fatal("Failed to create EEG simulator")
	}
}

// TestNeurochemicalState tests neurochemical state
func TestNeurochemicalState(t *testing.T) {
	state := NewNeurochemicalState()
	if state == nil {
		t.Fatal("Failed to create neurochemical state")
	}
}

// TestCognitiveState tests cognitive state
func TestCognitiveState(t *testing.T) {
	state := NewCognitiveState()
	if state == nil {
		t.Fatal("Failed to create cognitive state")
	}
}

