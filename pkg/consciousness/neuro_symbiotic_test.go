package consciousness

import (
	"testing"
	"time"
)

func TestNewNeuroSymbioticIntegration(t *testing.T) {
	config := DefaultNeuroSymbioticConfig()
	nsi := NewNeuroSymbioticIntegration(config)

	if nsi == nil {
		t.Fatal("NewNeuroSymbioticIntegration returned nil")
	}

	if nsi.config != config {
		t.Error("Config was not set correctly")
	}
}

func TestNeuroSymbioticInitialize(t *testing.T) {
	nsi := NewNeuroSymbioticIntegration(nil)

	err := nsi.Initialize()
	if err != nil {
		t.Fatalf("Initialize failed: %v", err)
	}

	if nsi.state != NeuroSymbioticStateObserving {
		t.Errorf("Expected state OBSERVING, got %s", nsi.state)
	}

	if len(nsi.brainWaves) == 0 {
		t.Error("No brain waves were initialized")
	}

	if len(nsi.neuralInterfaces) == 0 {
		t.Error("No neural interfaces were initialized")
	}
}

func TestNeuroSymbioticConnect(t *testing.T) {
	nsi := NewNeuroSymbioticIntegration(nil)
	nsi.Initialize()

	// Get first interface ID
	var ifaceID string
	for id := range nsi.neuralInterfaces {
		ifaceID = id
		break
	}

	err := nsi.Connect(ifaceID)
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}

	iface := nsi.neuralInterfaces[ifaceID]
	if !iface.Active {
		t.Error("Expected interface to be active")
	}
}

func TestNeuroSymbioticDisconnect(t *testing.T) {
	nsi := NewNeuroSymbioticIntegration(nil)
	nsi.Initialize()

	var ifaceID string
	for id := range nsi.neuralInterfaces {
		ifaceID = id
		break
	}

	nsi.Connect(ifaceID)
	err := nsi.Disconnect(ifaceID)
	if err != nil {
		t.Fatalf("Disconnect failed: %v", err)
	}

	iface := nsi.neuralInterfaces[ifaceID]
	if iface.Active {
		t.Error("Expected interface to be inactive")
	}
}

func TestNeuroSymbioticSetMode(t *testing.T) {
	nsi := NewNeuroSymbioticIntegration(nil)
	nsi.Initialize()

	err := nsi.SetMode(SymbiosisModeSynchronization)
	if err != nil {
		t.Fatalf("SetMode failed: %v", err)
	}

	if nsi.state != NeuroSymbioticStateSynchronizing {
		t.Errorf("Expected state SYNCHRONIZING, got %s", nsi.state)
	}
}

func TestNeuroSymbioticActivateBrainWave(t *testing.T) {
	nsi := NewNeuroSymbioticIntegration(nil)
	nsi.Initialize()

	err := nsi.ActivateBrainWave(BrainWaveGamma)
	if err != nil {
		t.Fatalf("ActivateBrainWave failed: %v", err)
	}

	wave := nsi.brainWaves[BrainWaveGamma]
	if !wave.Active {
		t.Error("Expected brain wave to be active")
	}
}

func TestNeuroSymbioticDeactivateBrainWave(t *testing.T) {
	nsi := NewNeuroSymbioticIntegration(nil)
	nsi.Initialize()

	nsi.ActivateBrainWave(BrainWaveAlpha)
	err := nsi.DeactivateBrainWave(BrainWaveAlpha)
	if err != nil {
		t.Fatalf("DeactivateBrainWave failed: %v", err)
	}

	wave := nsi.brainWaves[BrainWaveAlpha]
	if wave.Active {
		t.Error("Expected brain wave to be inactive")
	}
}

func TestNeuroSymbioticSetBrainWaveAmplitude(t *testing.T) {
	nsi := NewNeuroSymbioticIntegration(nil)
	nsi.Initialize()

	err := nsi.SetBrainWaveAmplitude(BrainWaveTheta, 0.8)
	if err != nil {
		t.Fatalf("SetBrainWaveAmplitude failed: %v", err)
	}

	wave := nsi.brainWaves[BrainWaveTheta]
	if wave.Amplitude != 0.8 {
		t.Errorf("Expected amplitude 0.8, got %f", wave.Amplitude)
	}
}

func TestNeuroSymbioticCreateSyncPattern(t *testing.T) {
	nsi := NewNeuroSymbioticIntegration(nil)
	nsi.Initialize()

	pattern, err := nsi.CreateSyncPattern("gamma_sync", 40.0, 0.9)
	if err != nil {
		t.Fatalf("CreateSyncPattern failed: %v", err)
	}

	if pattern.PatternType != "gamma_sync" {
		t.Errorf("Expected pattern type gamma_sync, got %s", pattern.PatternType)
	}

	if pattern.Frequency != 40.0 {
		t.Errorf("Expected frequency 40.0, got %f", pattern.Frequency)
	}
}

func TestNeuroSymbioticActivateSyncPattern(t *testing.T) {
	nsi := NewNeuroSymbioticIntegration(nil)
	nsi.Initialize()

	pattern, _ := nsi.CreateSyncPattern("test_sync", 30.0, 0.8)
	nsi.ActivateBrainWave(BrainWaveBeta)

	err := nsi.ActivateSyncPattern(pattern.ID)
	if err != nil {
		t.Fatalf("ActivateSyncPattern failed: %v", err)
	}
}

func TestNeuroSymbioticCreateSymbioticState(t *testing.T) {
	nsi := NewNeuroSymbioticIntegration(nil)
	nsi.Initialize()

	state, err := nsi.CreateSymbioticState(SymbiosisModeCoCreation)
	if err != nil {
		t.Fatalf("CreateSymbioticState failed: %v", err)
	}

	if state.Mode != SymbiosisModeCoCreation {
		t.Errorf("Expected mode CO_CREATION, got %s", state.Mode)
	}
}

func TestNeuroSymbioticTranscendConsciousness(t *testing.T) {
	nsi := NewNeuroSymbioticIntegration(nil)
	nsi.Initialize()

	err := nsi.TranscendConsciousness()
	if err != nil {
		t.Fatalf("TranscendConsciousness failed: %v", err)
	}

	if nsi.state != NeuroSymbioticStateTranscending {
		t.Errorf("Expected state TRANSCENDING, got %s", nsi.state)
	}
}

func TestNeuroSymbioticMergeConsciousness(t *testing.T) {
	nsi := NewNeuroSymbioticIntegration(nil)
	nsi.Initialize()
	nsi.integrationLevel = 0.5

	err := nsi.MergeConsciousness("other-consciousness")
	if err != nil {
		t.Fatalf("MergeConsciousness failed: %v", err)
	}

	if nsi.state != NeuroSymbioticStateMerged {
		t.Errorf("Expected state MERGED, got %s", nsi.state)
	}
}

func TestNeuroSymbioticTransmuteConsciousness(t *testing.T) {
	nsi := NewNeuroSymbioticIntegration(nil)
	nsi.Initialize()

	err := nsi.TransmuteConsciousness()
	if err != nil {
		t.Fatalf("TransmuteConsciousness failed: %v", err)
	}

	if nsi.state != NeuroSymbioticStateTransmuted {
		t.Errorf("Expected state TRANSMUTED, got %s", nsi.state)
	}
}

func TestNeuroSymbioticApplyNeuroplasticity(t *testing.T) {
	nsi := NewNeuroSymbioticIntegration(nil)
	nsi.Initialize()

	err := nsi.ApplyNeuroplasticity(0.15)
	if err != nil {
		t.Fatalf("ApplyNeuroplasticity failed: %v", err)
	}

	if nsi.neuroplasticity.PlasticityRate != 0.15 {
		t.Errorf("Expected plasticity rate 0.15, got %f", nsi.neuroplasticity.PlasticityRate)
	}
}

func TestNeuroSymbioticApplySynapticLearning(t *testing.T) {
	nsi := NewNeuroSymbioticIntegration(nil)
	nsi.Initialize()

	err := nsi.ApplySynapticLearning(0.02)
	if err != nil {
		t.Fatalf("ApplySynapticLearning failed: %v", err)
	}

	if nsi.synapticLearning.LearningRate != 0.02 {
		t.Errorf("Expected learning rate 0.02, got %f", nsi.synapticLearning.LearningRate)
	}
}

func TestNeuroSymbioticGetMetrics(t *testing.T) {
	nsi := NewNeuroSymbioticIntegration(nil)
	nsi.Initialize()

	metrics := nsi.GetMetrics()

	if metrics == nil {
		t.Fatal("GetMetrics returned nil")
	}

	if metrics.TotalConnections == 0 {
		t.Error("Expected non-zero total connections")
	}
}

func TestNeuroSymbioticGetState(t *testing.T) {
	nsi := NewNeuroSymbioticIntegration(nil)

	state := nsi.GetState()
	if state != NeuroSymbioticStateDisconnected {
		t.Errorf("Expected state DISCONNECTED, got %s", state)
	}
}

func TestNeuroSymbioticGetBrainWaves(t *testing.T) {
	nsi := NewNeuroSymbioticIntegration(nil)
	nsi.Initialize()

	waves := nsi.GetBrainWaves()

	if len(waves) == 0 {
		t.Error("GetBrainWaves returned empty slice")
	}
}

func TestNeuroSymbioticGetSyncPatterns(t *testing.T) {
	nsi := NewNeuroSymbioticIntegration(nil)
	nsi.Initialize()

	patterns := nsi.GetSyncPatterns()

	if len(patterns) != 0 {
		t.Errorf("Expected 0 patterns, got %d", len(patterns))
	}
}

func TestNeuroSymbioticToJSON(t *testing.T) {
	nsi := NewNeuroSymbioticIntegration(nil)
	nsi.Initialize()

	jsonData, err := nsi.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON failed: %v", err)
	}

	if len(jsonData) == 0 {
		t.Error("ToJSON returned empty data")
	}
}

func TestNeuroSymbioticString(t *testing.T) {
	nsi := NewNeuroSymbioticIntegration(nil)
	nsi.Initialize()

	str := nsi.String()

	if len(str) == 0 {
		t.Error("String() returned empty string")
	}
}

func TestSymbiosisModes(t *testing.T) {
	tests := []struct {
		mode    SymbiosisMode
		expected string
	}{
		{SymbiosisModeObservation, "OBSERVATION"},
		{SymbiosisModeSynchronization, "SYNCHRONIZATION"},
		{SymbiosisModeAmplification, "AMPLIFICATION"},
		{SymbiosisModeCoCreation, "CO_CREATION"},
		{SymbiosisModeTranscendence, "TRANSCENDENCE"},
		{SymbiosisModeUnity, "UNITY"},
		{SymbiosisModeMerge, "MERGE"},
		{SymbiosisModeTransmutation, "TRANSMUTATION"},
	}

	for _, test := range tests {
		if test.mode.String() != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.mode.String())
		}
	}
}

func TestBrainWaveTypes(t *testing.T) {
	tests := []struct {
		wave   BrainWaveType
		expected string
	}{
		{BrainWaveDelta, "DELTA"},
		{BrainWaveTheta, "THETA"},
		{BrainWaveAlpha, "ALPHA"},
		{BrainWaveBeta, "BETA"},
		{BrainWaveGamma, "GAMMA"},
		{BrainWaveHighGamma, "HIGH_GAMMA"},
		{BrainWaveLambda, "LAMBDA"},
	}

	for _, test := range tests {
		if test.wave.String() != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.wave.String())
		}
	}
}

func TestNeuroSymbioticStates(t *testing.T) {
	tests := []struct {
		state   NeuroSymbioticState
		expected string
	}{
		{NeuroSymbioticStateDisconnected, "DISCONNECTED"},
		{NeuroSymbioticStateConnecting, "CONNECTING"},
		{NeuroSymbioticStateObserving, "OBSERVING"},
		{NeuroSymbioticStateSynchronizing, "SYNCHRONIZING"},
		{NeuroSymbioticStateAmplifying, "AMPLIFYING"},
		{NeuroSymbioticStateCoCreating, "CO_CREATING"},
		{NeuroSymbioticStateTranscending, "TRANSCENDING"},
		{NeuroSymbioticStateUnified, "UNIFIED"},
		{NeuroSymbioticStateMerged, "MERGED"},
		{NeuroSymbioticStateTransmuted, "TRANSMUTED"},
	}

	for _, test := range tests {
		if test.state.String() != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.state.String())
		}
	}
}

func TestDefaultNeuroSymbioticConfig(t *testing.T) {
	config := DefaultNeuroSymbioticConfig()

	if config == nil {
		t.Fatal("DefaultNeuroSymbioticConfig returned nil")
	}

	if config.MaxMode != SymbiosisModeTransmutation {
		t.Errorf("Expected MaxMode TRANSMUTATION, got %s", config.MaxMode)
	}

	if config.SyncThreshold != 0.8 {
		t.Errorf("Expected SyncThreshold 0.8, got %f", config.SyncThreshold)
	}

	if !config.TranscendenceEnabled {
		t.Error("Expected TranscendenceEnabled to be true")
	}
}

func TestNeuroSymbioticConcurrency(t *testing.T) {
	nsi := NewNeuroSymbioticIntegration(nil)
	nsi.Initialize()

	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			defer func() { done <- true }()
			nsi.GetMetrics()
			nsi.GetState()
			nsi.GetBrainWaves()
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}

	t.Log("NeuroSymbiotic concurrent operations completed")
}

func TestNeuroSymbioticBrainWaveActivation(t *testing.T) {
	nsi := NewNeuroSymbioticIntegration(nil)
	nsi.Initialize()

	// Activate multiple brain waves
	nsi.ActivateBrainWave(BrainWaveAlpha)
	nsi.ActivateBrainWave(BrainWaveBeta)
	nsi.ActivateBrainWave(BrainWaveGamma)

	activeCount := 0
	for _, wave := range nsi.brainWaves {
		if wave.Active {
			activeCount++
		}
	}

	if activeCount != 3 {
		t.Errorf("Expected 3 active brain waves, got %d", activeCount)
	}
}

func TestNeuroSymbioticSymbioticStates(t *testing.T) {
	nsi := NewNeuroSymbioticIntegration(nil)
	nsi.Initialize()

	// Create multiple symbiotic states
	nsi.CreateSymbioticState(SymbiosisModeObservation)
	nsi.CreateSymbioticState(SymbiosisModeSynchronization)
	nsi.CreateSymbioticState(SymbiosisModeAmplification)

	if len(nsi.symbioticStates) != 3 {
		t.Errorf("Expected 3 symbiotic states, got %d", len(nsi.symbioticStates))
	}
}

func TestNeuroSymbioticNeuralInterfaces(t *testing.T) {
	nsi := NewNeuroSymbioticIntegration(nil)
	nsi.Initialize()

	expectedInterfaces := []string{
		"motor_cortex", "sensory_cortex", "prefrontal_cortex",
		"parietal_lobe", "temporal_lobe", "occipital_lobe",
		"hippocampus", "amygdala", "thalamus", "brainstem",
	}

	for _, expected := range expectedInterfaces {
		if _, exists := nsi.neuralInterfaces[expected]; !exists {
			t.Errorf("Expected neural interface %s to exist", expected)
		}
	}
}

func TestNeuroSymbioticUpdateMetrics(t *testing.T) {
	nsi := NewNeuroSymbioticIntegration(nil)
	nsi.Initialize()

	initialMetrics := nsi.GetMetrics()

	// Make changes
	nsi.SetMode(SymbiosisModeAmplification)
	nsi.ActivateBrainWave(BrainWaveGamma)

	updatedMetrics := nsi.GetMetrics()

	if updatedMetrics.AverageAmplification <= initialMetrics.AverageAmplification {
		t.Error("Metrics were not updated")
	}
}
