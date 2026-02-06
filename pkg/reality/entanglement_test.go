package reality

import (
	"testing"
	"time"
)

func TestNewEntanglementManager(t *testing.T) {
	config := DefaultEntanglementConfig()
	em := NewEntanglementManager(config)
	
	if em == nil {
		t.Fatal("NewEntanglementManager returned nil")
	}
	
	if em.config != config {
		t.Error("Config was not set correctly")
	}
	
	if em.state != EntanglementManagerStateIdle {
		t.Errorf("Expected state IDLE, got %s", em.state)
	}
}

func TestEntanglementManagerInitialize(t *testing.T) {
	em := NewEntanglementManager(nil)
	
	err := em.Initialize()
	if err != nil {
		t.Fatalf("Initialize failed: %v", err)
	}
	
	if em.state != EntanglementManagerStateActive {
		t.Errorf("Expected state ACTIVE, got %s", em.state)
	}
	
	if len(em.realityStates) == 0 {
		t.Error("No reality states were created during initialization")
	}
}

func TestCreateEntanglement(t *testing.T) {
	em := NewEntanglementManager(nil)
	em.Initialize()
	
	pair, err := em.CreateEntanglement("base_reality", "quantum_divergent", EntanglementTypeSpatial)
	if err != nil {
		t.Fatalf("CreateEntanglement failed: %v", err)
	}
	
	if pair == nil {
		t.Fatal("CreateEntanglement returned nil")
	}
	
	if pair.RealityA != "base_reality" {
		t.Errorf("Expected RealityA base_reality, got %s", pair.RealityA)
	}
	
	if pair.RealityB != "quantum_divergent" {
		t.Errorf("Expected RealityB quantum_divergent, got %s", pair.RealityB)
	}
	
	if pair.Coherence <= 0 {
		t.Error("Expected positive coherence")
	}
	
	if pair.Strength <= 0 {
		t.Error("Expected positive strength")
	}
}

func TestCreateEntanglementMaxLimit(t *testing.T) {
	em := NewEntanglementManager(nil)
	em.Initialize()
	
	// Create max entanglements
	for i := 0; i < em.config.MaxEntanglements; i++ {
		_, err := em.CreateEntanglement(
			"base_reality",
			"quantum_divergent",
			EntanglementTypeSpatial,
		)
		if err != nil {
			// This is expected when limit is reached
			break
		}
	}
	
	// Try to create one more
	_, err := em.CreateEntanglement("base_reality", "quantum_divergent", EntanglementTypeSpatial)
	if err == nil {
		t.Error("Expected error when max entanglements reached")
	}
}

func TestActivateEntanglement(t *testing.T) {
	em := NewEntanglementManager(nil)
	em.Initialize()
	
	pair, _ := em.CreateEntanglement("base_reality", "quantum_divergent", EntanglementTypeSpatial)
	
	err := em.ActivateEntanglement(pair.ID)
	if err != nil {
		t.Fatalf("ActivateEntanglement failed: %v", err)
	}
	
	if pair.State != EntanglementStateActive {
		t.Errorf("Expected state ACTIVE, got %s", pair.State)
	}
}

func TestCollapseEntanglement(t *testing.T) {
	em := NewEntanglementManager(nil)
	em.Initialize()
	
	pair, _ := em.CreateEntanglement("base_reality", "quantum_divergent", EntanglementTypeSpatial)
	
	err := em.CollapseEntanglement(pair.ID)
	if err != nil {
		t.Fatalf("CollapseEntanglement failed: %v", err)
	}
	
	if pair.State != EntanglementStateCollapsing && pair.State != EntanglementStateBroken &&
		pair.State != EntanglementStateTranscendent {
		t.Errorf("Expected state COLLAPSING, BROKEN, or TRANSCENDENT, got %s", pair.State)
	}
}

func TestBreakEntanglement(t *testing.T) {
	em := NewEntanglementManager(nil)
	em.Initialize()
	
	pair, _ := em.CreateEntanglement("base_reality", "quantum_divergent", EntanglementTypeSpatial)
	
	err := em.BreakEntanglement(pair.ID)
	if err != nil {
		t.Fatalf("BreakEntanglement failed: %v", err)
	}
	
	if pair.State != EntanglementStateBroken {
		t.Errorf("Expected state BROKEN, got %s", pair.State)
	}
	
	if pair.Coherence != 0.0 {
		t.Errorf("Expected coherence 0.0, got %f", pair.Coherence)
	}
}

func TestMeasureEntanglement(t *testing.T) {
	em := NewEntanglementManager(nil)
	em.Initialize()
	
	pair, _ := em.CreateEntanglement("base_reality", "quantum_divergent", EntanglementTypeSpatial)
	
	measuredA, measuredB, err := em.MeasureEntanglement(pair.ID)
	if err != nil {
		t.Fatalf("MeasureEntanglement failed: %v", err)
	}
	
	if measuredA == 0 {
		t.Error("Expected non-zero measured phase A")
	}
	
	if measuredB == 0 {
		t.Error("Expected non-zero measured phase B")
	}
}

func TestTransferInformation(t *testing.T) {
	em := NewEntanglementManager(nil)
	em.Initialize()
	
	pair, _ := em.CreateEntanglement("base_reality", "quantum_divergent", EntanglementTypeSpatial)
	em.ActivateEntanglement(pair.ID)
	
	information := complex(1.0, 0.5)
	transferred, err := em.TransferInformation(pair.ID, information)
	if err != nil {
		t.Fatalf("TransferInformation failed: %v", err)
	}
	
	if transferred == 0 {
		t.Error("Expected non-zero transferred information")
	}
}

func TestSynchronizeEntanglements(t *testing.T) {
	em := NewEntanglementManager(nil)
	em.Initialize()
	
	// Create multiple entanglements
	em.CreateEntanglement("base_reality", "quantum_divergent", EntanglementTypeSpatial)
	em.CreateEntanglement("temporal_inverted", "entropic_reversed", EntanglementTypeTemporal)
	em.CreateEntanglement("consciousness_amplified", "dimensional_shifted", EntanglementTypeConsciousness)
	
	err := em.SynchronizeEntanglements()
	if err != nil {
		t.Fatalf("SynchronizeEntanglements failed: %v", err)
	}
	
	if em.state != EntanglementManagerStateActive {
		t.Errorf("Expected state ACTIVE, got %s", em.state)
	}
}

func TestUpdateReality(t *testing.T) {
	em := NewEntanglementManager(nil)
	em.Initialize()
	
	pair, _ := em.CreateEntanglement("base_reality", "quantum_divergent", EntanglementTypeSpatial)
	initialCoherence := pair.Coherence
	
	err := em.UpdateReality("base_reality", func(state *RealityState) {
		state.QuantumCoherence = 0.95
	})
	if err != nil {
		t.Fatalf("UpdateReality failed: %v", err)
	}
	
	// Check that entanglement was updated
	if pair.Coherence == initialCoherence {
		t.Error("Expected entanglement coherence to update")
	}
}

func TestGetEntanglement(t *testing.T) {
	em := NewEntanglementManager(nil)
	em.Initialize()
	
	created, _ := em.CreateEntanglement("base_reality", "quantum_divergent", EntanglementTypeSpatial)
	
	retrieved, err := em.GetEntanglement(created.ID)
	if err != nil {
		t.Fatalf("GetEntanglement failed: %v", err)
	}
	
	if retrieved.ID != created.ID {
		t.Errorf("Expected ID %s, got %s", created.ID, retrieved.ID)
	}
}

func TestGetAllEntanglements(t *testing.T) {
	em := NewEntanglementManager(nil)
	em.Initialize()
	
	em.CreateEntanglement("base_reality", "quantum_divergent", EntanglementTypeSpatial)
	em.CreateEntanglement("temporal_inverted", "entropic_reversed", EntanglementTypeTemporal)
	
	entanglements := em.GetAllEntanglements()
	
	if len(entanglements) != 2 {
		t.Errorf("Expected 2 entanglements, got %d", len(entanglements))
	}
}

func TestGetMetrics(t *testing.T) {
	em := NewEntanglementManager(nil)
	em.Initialize()
	
	metrics := em.GetMetrics()
	
	if metrics == nil {
		t.Fatal("GetMetrics returned nil")
	}
	
	if metrics.TotalEntanglements == 0 {
		t.Error("Expected non-zero total entanglements")
	}
}

func TestEntanglementTypes(t *testing.T) {
	tests := []struct {
		et       EntanglementType
		expected string
	}{
		{EntanglementTypeSpatial, "SPATIAL"},
		{EntanglementTypeTemporal, "TEMPORAL"},
		{EntanglementTypeCausal, "CAUSAL"},
		{EntanglementTypeInformational, "INFORMATIONAL"},
		{EntanglementTypeConsciousness, "CONSCIOUSNESS"},
		{EntanglementTypeEmotional, "EMOTIONAL"},
		{EntanglementTypePurpose, "PURPOSE"},
		{EntanglementTypeTranscendent, "TRANSCENDENT"},
	}
	
	for _, test := range tests {
		if test.et.String() != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.et.String())
		}
	}
}

func TestEntanglementStates(t *testing.T) {
	tests := []struct {
		es       EntanglementState
		expected string
	}{
		{EntanglementStatePotential, "POTENTIAL"},
		{EntanglementStateForming, "FORMING"},
		{EntanglementStateActive, "ACTIVE"},
		{EntanglementStateWeak, "WEAK"},
		{EntanglementStateCollapsing, "COLLAPSING"},
		{EntanglementStateBroken, "BROKEN"},
		{EntanglementStateTranscendent, "TRANSCENDENT"},
	}
	
	for _, test := range tests {
		if test.es.String() != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.es.String())
		}
	}
}

func TestEntanglementManagerStates(t *testing.T) {
	tests := []struct {
		ems      EntanglementManagerState
		expected string
	}{
		{EntanglementManagerStateIdle, "IDLE"},
		{EntanglementManagerStateScanning, "SCANNING"},
		{EntanglementManagerStateEntangling, "ENTANGLING"},
		{EntanglementManagerStateActive, "ACTIVE"},
		{EntanglementManagerStateCollapsing, "COLLAPSING"},
		{EntanglementManagerStateSynchronizing, "SYNCHRONIZING"},
		{EntanglementManagerStateError, "ERROR"},
	}
	
	for _, test := range tests {
		if test.ems.String() != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.ems.String())
		}
	}
}

func TestEntanglementManagerString(t *testing.T) {
	em := NewEntanglementManager(nil)
	em.Initialize()
	
	str := em.String()
	
	if len(str) == 0 {
		t.Error("String() returned empty string")
	}
}

func TestEntanglementManagerToJSON(t *testing.T) {
	em := NewEntanglementManager(nil)
	em.Initialize()
	
	jsonData, err := em.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON failed: %v", err)
	}
	
	if len(jsonData) == 0 {
		t.Error("ToJSON returned empty data")
	}
}

func TestDefaultConfig(t *testing.T) {
	config := DefaultEntanglementConfig()
	
	if config == nil {
		t.Fatal("DefaultEntanglementConfig returned nil")
	}
	
	if config.MaxEntanglements != 100 {
		t.Errorf("Expected MaxEntanglements 100, got %d", config.MaxEntanglements)
	}
	
	if config.CoherenceThreshold != 0.8 {
		t.Errorf("Expected CoherenceThreshold 0.8, got %f", config.CoherenceThreshold)
	}
	
	if !config.TranscendentEnabled {
		t.Error("Expected TranscendentEnabled to be true")
	}
}

func TestConcurrency(t *testing.T) {
	em := NewEntanglementManager(nil)
	em.Initialize()
	
	done := make(chan bool)
	
	// Test concurrent access
	for i := 0; i < 10; i++ {
		go func() {
			defer func() { done <- true }()
			
			// Multiple concurrent operations
			_ = em.GetMetrics()
			_ = em.GetState()
			_ = em.GetAllEntanglements()
		}()
	}
	
	// Wait for all goroutines
	for i := 0; i < 10; i++ {
		<-done
	}
	
	t.Log("Concurrent operations completed successfully")
}

func TestRealityStatesCreated(t *testing.T) {
	em := NewEntanglementManager(nil)
	em.Initialize()
	
	expectedRealities := []string{
		"base_reality",
		"quantum_divergent",
		"temporal_inverted",
		"entropic_reversed",
		"consciousness_amplified",
		"dimensional_shifted",
		"causal_broken",
		"information_dense",
		"void_reality",
		"singularity_reality",
	}
	
	for _, realityType := range expectedRealities {
		state, exists := em.realityStates[realityType]
		if !exists {
			t.Errorf("Expected reality state %s to exist", realityType)
		}
		if state.Type != realityType {
			t.Errorf("Expected type %s, got %s", realityType, state.Type)
		}
	}
}

func TestEntanglementProperties(t *testing.T) {
	em := NewEntanglementManager(nil)
	em.Initialize()
	
	pair, _ := em.CreateEntanglement("base_reality", "quantum_divergent", EntanglementTypeSpatial)
	
	if pair.Properties == nil {
		t.Error("Expected Properties to be initialized")
	}
	
	if pair.CreatedAt.IsZero() {
		t.Error("Expected CreatedAt to be set")
	}
}

func TestDistanceCalculation(t *testing.T) {
	em := NewEntanglementManager(nil)
	em.Initialize()
	
	// Distance should be calculated
	pair, _ := em.CreateEntanglement("base_reality", "quantum_divergent", EntanglementTypeSpatial)
	
	if pair.Distance < 0 {
		t.Error("Expected non-negative distance")
	}
}

func TestEntanglementEntropy(t *testing.T) {
	em := NewEntanglementManager(nil)
	em.Initialize()
	
	pair, _ := em.CreateEntanglement("base_reality", "quantum_divergent", EntanglementTypeSpatial)
	
	if pair.EntanglementEntropy < 0 {
		t.Error("Expected non-negative entropy")
	}
}
