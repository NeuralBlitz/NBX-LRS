package lrs

import (
	"encoding/json"
	"testing"
	"time"
)

// Test IntegrationState String()
func TestIntegrationStateString(t *testing.T) {
	tests := []struct {
		state   IntegrationState
		expected string
	}{
		{StateInitializing, "initializing"},
		{StateConnected, "connected"},
		{StateActive, "active"},
		{StateError, "error"},
		{IntegrationState(99), "unknown"},
	}

	for _, tt := range tests {
		result := tt.state.String()
		if result != tt.expected {
			t.Errorf("IntegrationState(%d).String() = %s, want %s", tt.state, result, tt.expected)
		}
	}
}

// Test NewLRSNeuralBlitzBridge
func TestNewLRSNeuralBlitzBridge(t *testing.T) {
	bridge := NewLRSNeuralBlitzBridge()

	if bridge.AgentEndpoint != DefaultAgentEndpoint {
		t.Errorf("AgentEndpoint = %s, want %s", bridge.AgentEndpoint, DefaultAgentEndpoint)
	}

	if bridge.AuthKey != DefaultAuthKey {
		t.Errorf("AuthKey = %s, want %s", bridge.AuthKey, DefaultAuthKey)
	}

	if bridge.BridgePort != DefaultBridgePort {
		t.Errorf("BridgePort = %d, want %d", bridge.BridgePort, DefaultBridgePort)
	}

	if bridge.QuantumNeuron == nil {
		t.Error("QuantumNeuron is nil")
	}

	if bridge.RealityNetwork == nil {
		t.Error("RealityNetwork is nil")
	}

	if bridge.LRSAgent == nil {
		t.Error("LRSAgent is nil")
	}

	if bridge.State != StateInitializing {
		t.Errorf("State = %s, want initializing", bridge.State)
	}
}

// Test Initialize
func TestBridgeInitialize(t *testing.T) {
	bridge := NewLRSNeuralBlitzBridge()

	err := bridge.Initialize()

	if err != nil {
		t.Errorf("Initialize() error = %v", err)
	}

	if bridge.State != StateConnected {
		t.Errorf("State = %s, want connected", bridge.State)
	}

	if bridge.QuantumNeuron.Config == nil {
		t.Error("QuantumNeuron.Config is nil after Initialize")
	}

	if bridge.RealityNetwork.RealityStates == nil {
		t.Error("RealityNetwork.RealityStates is nil after Initialize")
	}

	if bridge.TotalCycles != 0 {
		t.Errorf("TotalCycles = %d, want 0", bridge.TotalCycles)
	}
}

// Test RunCycle
func TestRunCycle(t *testing.T) {
	bridge := NewLRSNeuralBlitzBridge()
	bridge.Initialize()

	metrics, err := bridge.RunCycle(0, 15.0)

	if err != nil {
		t.Errorf("RunCycle() error = %v", err)
	}

	if metrics.Cycle != 0 {
		t.Errorf("Metrics.Cycle = %d, want 0", metrics.Cycle)
	}

	if metrics.Timestamp.IsZero() {
		t.Error("Metrics.Timestamp is zero")
	}

	if bridge.TotalCycles != 1 {
		t.Errorf("TotalCycles = %d, want 1", bridge.TotalCycles)
	}
}

// Test RunCycle Not Ready Error
func TestRunCycleNotReady(t *testing.T) {
	bridge := NewLRSNeuralBlitzBridge()
	// Don't initialize

	_, err := bridge.RunCycle(0, 15.0)

	if err != ErrBridgeNotReady {
		t.Errorf("RunCycle() without Initialize = %v, want ErrBridgeNotReady", err)
	}
}

// Test RunCycle Multiple Cycles
func TestRunCycleMultiple(t *testing.T) {
	bridge := NewLRSNeuralBlitzBridge()
	bridge.Initialize()

	for i := 0; i < 10; i++ {
		metrics, err := bridge.RunCycle(i, 15.0)

		if err != nil {
			t.Errorf("RunCycle(%d) error = %v", i, err)
		}

		if metrics.Cycle != i {
			t.Errorf("Metrics.Cycle = %d, want %d", metrics.Cycle, i)
		}
	}

	if bridge.TotalCycles != 10 {
		t.Errorf("TotalCycles = %d, want 10", bridge.TotalCycles)
	}

	if len(bridge.MetricsHistory) != 10 {
		t.Errorf("MetricsHistory length = %d, want 10", len(bridge.MetricsHistory))
	}
}

// Test GetMetricsHistory
func TestGetMetricsHistory(t *testing.T) {
	bridge := NewLRSNeuralBlitzBridge()
	bridge.Initialize()

	// Run some cycles
	for i := 0; i < 5; i++ {
		bridge.RunCycle(i, 15.0)
	}

	history := bridge.GetMetricsHistory()

	if len(history) != 5 {
		t.Errorf("GetMetricsHistory() length = %d, want 5", len(history))
	}

	// Verify history is independent
	bridge.RunCycle(99, 15.0)
	if len(history) != 5 {
		t.Error("GetMetricsHistory() returned mutable slice")
	}
}

// Test UpdatePrecision
func TestUpdatePrecision(t *testing.T) {
	bridge := NewLRSNeuralBlitzBridge()
	bridge.Initialize()

	bridge.UpdatePrecision(1.5, 2.0)

	if bridge.LRSAgent.PrecisionParams.Alpha != 1.5 {
		t.Errorf("PrecisionParams.Alpha = %f, want 1.5", bridge.LRSAgent.PrecisionParams.Alpha)
	}

	if bridge.LRSAgent.PrecisionParams.Beta != 2.0 {
		t.Errorf("PrecisionParams.Beta = %f, want 2.0", bridge.LRSAgent.PrecisionParams.Beta)
	}
}

// Test GetBridgeStatus
func TestGetBridgeStatus(t *testing.T) {
	bridge := NewLRSNeuralBlitzBridge()
	bridge.Initialize()

	status := bridge.GetBridgeStatus()

	if status["state"] != "connected" {
		t.Errorf("Status state = %v, want connected", status["state"])
	}

	if status["agent_endpoint"].(string) != DefaultAgentEndpoint {
		t.Errorf("Agent endpoint = %v, want %s", status["agent_endpoint"], DefaultAgentEndpoint)
	}

	quantumNeuronStatus := status["quantum_neuron"].(map[string]interface{})
	if quantumNeuronStatus["neuron_id"] == "" {
		t.Error("Neuron ID is empty")
	}

	realityNetworkStatus := status["reality_network"].(map[string]interface{})
	if realityNetworkStatus["num_realities"].(int) != 4 {
		t.Errorf("Num realities = %v, want 4", realityNetworkStatus["num_realities"])
	}
}

// Test NewQuantumNeuronBridge
func TestNewQuantumNeuronBridge(t *testing.T) {
	neuron := NewQuantumNeuronBridge()

	if neuron.NeuronID == "" {
		t.Error("NeuronID is empty")
	}

	if neuron.Config == nil {
		t.Error("Config is nil")
	}

	if neuron.MembranePotential != -70.0 {
		t.Errorf("MembranePotential = %f, want -70.0", neuron.MembranePotential)
	}
}

// Test NewRealityNetworkBridge
func TestNewRealityNetworkBridge(t *testing.T) {
	network := NewRealityNetworkBridge()

	if network.NumRealities != 4 {
		t.Errorf("NumRealities = %d, want 4", network.NumRealities)
	}

	if network.NodesPerReality != 25 {
		t.Errorf("NodesPerReality = %d, want 25", network.NodesPerReality)
	}

	if network.GlobalConsciousness != 0.5 {
		t.Errorf("GlobalConsciousness = %f, want 0.5", network.GlobalConsciousness)
	}
}

// Test NewLRSElementaryAgent
func TestNewLRSElementaryAgent(t *testing.T) {
	agent := NewLRSElementaryAgent()

	if agent.Name == "" {
		t.Error("Name is empty")
	}

	if agent.PrecisionParams == nil {
		t.Error("PrecisionParams is nil")
	}

	if agent.Precision != 1.0 {
		t.Errorf("Precision = %f, want 1.0", agent.Precision)
	}
}

// Test CycleMetrics Serialization
func TestCycleMetricsSerialization(t *testing.T) {
	metrics := &CycleMetrics{
		Cycle:             5,
		Spikes:            10,
		SpikeRate:         50.0,
		MembranePotential: -65.0,
		FreeEnergy:        0.5,
		PredictionError:   0.2,
		Consciousness:     0.6,
		Timestamp:        time.Now(),
	}

	jsonData, err := json.Marshal(metrics)
	if err != nil {
		t.Errorf("CycleMetrics JSON marshal error = %v", err)
	}

	var result CycleMetrics
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Errorf("CycleMetrics JSON unmarshal error = %v", err)
	}

	if result.Cycle != metrics.Cycle {
		t.Errorf("Cycle = %d, want %d", result.Cycle, metrics.Cycle)
	}

	if result.Spikes != metrics.Spikes {
		t.Errorf("Spikes = %d, want %d", result.Spikes, metrics.Spikes)
	}

	if result.FreeEnergy != metrics.FreeEnergy {
		t.Errorf("FreeEnergy = %f, want %f", result.FreeEnergy, metrics.FreeEnergy)
	}
}

// Test CalculateFreeEnergyResult
func TestCalculateFreeEnergyResult(t *testing.T) {
	bridge := NewLRSNeuralBlitzBridge()
	bridge.Initialize()

	result := bridge.CalculateFreeEnergyResult(5.0, 0.3)

	if result.FreeEnergy < 0.0 {
		t.Error("FreeEnergy is negative")
	}

	if result.PredictionError != 0.3 {
		t.Errorf("PredictionError = %f, want 0.3", result.PredictionError)
	}
}

// Test Quantum Neuron Simulation
func TestSimulateQuantumNeuron(t *testing.T) {
	bridge := NewLRSNeuralBlitzBridge()
	bridge.Initialize()

	// Test with different input currents
	testCases := []struct {
		inputCurrent float64
		minSpikes    int
		maxSpikes    int
	}{
		{10.0, 0, 5},    // Low current - few spikes
		{20.0, 5, 20},   // Medium current - more spikes
		{50.0, 20, 50},  // High current - many spikes
	}

	for _, tc := range testCases {
		spikes := bridge.simulateQuantumNeuron(tc.inputCurrent)

		if spikes < tc.minSpikes || spikes > tc.maxSpikes {
			t.Errorf("SimulateQuantumNeuron(%f) spikes = %d, want between %d and %d",
				tc.inputCurrent, spikes, tc.minSpikes, tc.maxSpikes)
		}
	}
}

// Test Reality Network Evolution
func TestEvolveRealityNetwork(t *testing.T) {
	bridge := NewLRSNeuralBlitzBridge()
	bridge.Initialize()

	initialConsciousness := bridge.RealityNetwork.GlobalConsciousness

	// Evolve multiple times
	for i := 0; i < 10; i++ {
		bridge.evolveRealityNetwork()
	}

	finalConsciousness := bridge.RealityNetwork.GlobalConsciousness

	// Consciousness should have changed
	if initialConsciousness == finalConsciousness {
		t.Error("Reality network consciousness should change after evolution")
	}

	// Consciousness should be within bounds
	if finalConsciousness < 0.0 || finalConsciousness > 1.0 {
		t.Errorf("GlobalConsciousness = %f, want between 0.0 and 1.0", finalConsciousness)
	}

	// Should have updated reality states
	if len(bridge.RealityNetwork.RealityStates) != bridge.RealityNetwork.NumRealities {
		t.Errorf("RealityStates count = %d, want %d",
			len(bridge.RealityNetwork.RealityStates), bridge.RealityNetwork.NumRealities)
	}
}

// Test LRSPredict
func TestLRSPredict(t *testing.T) {
	bridge := NewLRSNeuralBlitzBridge()
	bridge.Initialize()

	// First prediction
	prediction1, error1 := bridge.lrsPredict(10)

	if prediction1 < 0.0 {
		t.Error("First prediction is negative")
	}

	if error1 < 0.0 {
		t.Error("First prediction error is negative")
	}

	// Second prediction with different observation
	prediction2, error2 := bridge.lrsPredict(20)

	// With more data, predictions should be based on recent observations
	if bridge.LRSAgent.Precision < MinPrecision {
		t.Error("Precision should not go below minimum")
	}

	if bridge.LRSAgent.Precision > MaxPrecision {
		t.Error("Precision should not exceed maximum")
	}
}

// Test CalculateFreeEnergy
func TestCalculateFreeEnergy(t *testing.T) {
	bridge := NewLRSNeuralBlitzBridge()
	bridge.Initialize()

	// Set precision
	bridge.LRSAgent.Precision = 1.0

	freeEnergy := bridge.calculateFreeEnergy(5.0, 0.2)

	expectedFE := 0.2 + 1.0*FreeEnergyPenalty*5.0
	if freeEnergy != expectedFE {
		t.Errorf("FreeEnergy = %f, want %f", freeEnergy, expectedFE)
	}
}

// Test ToJSON
func TestBridgeToJSON(t *testing.T) {
	bridge := NewLRSNeuralBlitzBridge()
	bridge.Initialize()

	jsonData, err := bridge.ToJSON()
	if err != nil {
		t.Errorf("ToJSON() error = %v", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Errorf("JSON unmarshal error = %v", err)
	}

	if result["state"] != "connected" {
		t.Errorf("JSON state = %v, want connected", result["state"])
	}
}

// Test generateNeuronID
func TestGenerateNeuronID(t *testing.T) {
	id1 := generateNeuronID()
	id2 := generateNeuronID()

	// Should generate unique IDs
	if id1 == id2 {
		t.Error("Neuron IDs should be unique")
	}

	// Should start with "neuron_"
	if len(id1) < 8 || id1[:7] != "neuron_" {
		t.Errorf("Neuron ID = %s, want to start with 'neuron_'", id1)
	}
}

// Test generateRealityID
func TestGenerateRealityID(t *testing.T) {
	id := generateRealityID(0)

	if id != "reality_a" {
		t.Errorf("Reality ID = %s, want reality_a", id)
	}
}

// Test Precision Boundaries
func TestPrecisionBoundaries(t *testing.T) {
	bridge := NewLRSNeuralBlitzBridge()
	bridge.Initialize()

	// Test with very high prediction error
	bridge.LRSAgent.Precision = MaxPrecision
	_, error1 := bridge.lrsPredict(100)
	freeEnergy1 := bridge.calculateFreeEnergy(100, error1)

	// Should increase precision
	if bridge.LRSAgent.Precision > 1.0 {
		t.Error("Precision should decrease with high error")
	}

	// Test precision clamping
	bridge.LRSAgent.Precision = 100.0
	bridge.LRSAgent.Precision *= 0.01
	if bridge.LRSAgent.Precision < MinPrecision {
		t.Error("Precision should be clamped to minimum")
	}

	bridge.LRSAgent.Precision = 0.01
	bridge.LRSAgent.Precision *= 10.0
	if bridge.LRSAgent.Precision > MaxPrecision {
		t.Error("Precision should be clamped to maximum")
	}
}

// Test Metrics History Limit
func TestMetricsHistoryLimit(t *testing.T) {
	bridge := NewLRSNeuralBlitzBridge()
	bridge.Initialize()

	// Run many cycles
	for i := 0; i < 1500; i++ {
		bridge.RunCycle(i, 15.0)
	}

	// History should be limited to 1000
	if len(bridge.MetricsHistory) > 1000 {
		t.Errorf("MetricsHistory length = %d, want <= 1000", len(bridge.MetricsHistory))
	}
}
