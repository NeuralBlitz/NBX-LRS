package consciousness

import (
	"encoding/json"
	"testing"
	"time"
)

// Test EntrainmentMode String()
func TestEntrainmentModeString(t *testing.T) {
	tests := []struct {
		mode    EntrainmentMode
		expected string
	}{
		{ModeBinauralBeats, "binaural_beats"},
		{ModeIsochronicTones, "isochronic_tones"},
		{ModeMonauralBeats, "monaural_beats"},
		{ModePhoticStimulation, "photic_stimulation"},
		{ModeAudiovisual, "audiovisual"},
		{ModeNeurofeedback, "neurofeedback"},
		{EntrainmentMode(99), "unknown"},
	}

	for _, tt := range tests {
		result := tt.mode.String()
		if result != tt.expected {
			t.Errorf("EntrainmentMode(%d).String() = %s, want %s", tt.mode, result, tt.expected)
		}
	}
}

// Test TargetFrequency Description()
func TestTargetFrequencyDescription(t *testing.T) {
	tests := []struct {
		freq    TargetFrequency
		expected string
	}{
		{FrequencyDelta, "Deep Sleep/Healing"},
		{FrequencyTheta, "Meditation/Creativity"},
		{FrequencyAlpha, "Relaxed Focus"},
		{FrequencyBeta, "Active Thinking"},
		{FrequencyGamma, "Higher Consciousness"},
		{FrequencyLambda, "Advanced Processing"},
		{TargetFrequency(99), "Unknown"},
	}

	for _, tt := range tests {
		result := tt.freq.Description()
		if result != tt.expected {
			t.Errorf("TargetFrequency(%f).Description() = %s, want %s", tt.freq, result, tt.expected)
		}
	}
}

// Test BrainWaveGenerator Initialization
func TestNewBrainWaveGenerator(t *testing.T) {
	generator := NewBrainWaveGenerator(44100)

	if generator.SampleRate != 44100 {
		t.Errorf("SampleRate = %d, want 44100", generator.SampleRate)
	}

	// Test with default sample rate
	generatorDefault := NewBrainWaveGenerator(0)
	if generatorDefault.SampleRate != DefaultSampleRate {
		t.Errorf("Default SampleRate = %d, want %d", generatorDefault.SampleRate, DefaultSampleRate)
	}
}

// Test GenerateBinauralBeats
func TestGenerateBinauralBeats(t *testing.T) {
	generator := NewBrainWaveGenerator(44100)

	left, right, err := generator.GenerateBinauralBeats(10.0, 1.0, 200.0, 0.0)

	if err != nil {
		t.Errorf("GenerateBinauralBeats() error = %v", err)
	}

	if len(left) == 0 {
	BinauralBe	t.Error("Generateats() returned empty left signal")
	}

	if len(right) == 0 {
		t.Error("GenerateBinauralBeats() returned empty right signal")
	}

	if len(left) != len(right) {
		t.Errorf("Left and right signals have different lengths: %d vs %d", len(left), len(right))
	}

	// Test with invalid frequency
	_, _, err = generator.GenerateBinauralBeats(0, 1.0, 200.0, 0.0)
	if err != ErrInvalidFrequency {
		t.Errorf("GenerateBinauralBeats() with 0 frequency = %v, want ErrInvalidFrequency", err)
	}
}

// Test GenerateIsochronicTones
func TestGenerateIsochronicTones(t *testing.T) {
	generator := NewBrainWaveGenerator(44100)

	signal, err := generator.GenerateIsochronicTones(10.0, 1.0, 0.8)

	if err != nil {
		t.Errorf("GenerateIsochronicTones() error = %v", err)
	}

	if len(signal) == 0 {
		t.Error("GenerateIsochronicTones() returned empty signal")
	}

	// Test with invalid intensity
	_, err = generator.GenerateIsochronicTones(10.0, 1.0, 1.5)
	if err != ErrInvalidIntensity {
		t.Errorf("GenerateIsochronicTones() with 1.5 intensity = %v, want ErrInvalidIntensity", err)
	}
}

// Test GenerateMonauralBeats
func TestGenerateMonauralBeats(t *testing.T) {
	generator := NewBrainWaveGenerator(44100)

	signal, err := generator.GenerateMonauralBeats(10.0, 1.0, 440.0)

	if err != nil {
		t.Errorf("GenerateMonauralBeats() error = %v", err)
	}

	if len(signal) == 0 {
		t.Error("GenerateMonauralBeats() returned empty signal")
	}
}

// Test GeneratePhoticStimulation
func TestGeneratePhoticStimulation(t *testing.T) {
	generator := NewBrainWaveGenerator(44100)

	signal, err := generator.GeneratePhoticStimulation(10.0, 1.0, 0.8)

	if err != nil {
		t.Errorf("GeneratePhoticStimulation() error = %v", err)
	}

	if len(signal) == 0 {
		t.Error("GeneratePhoticStimulation() returned empty signal")
	}

	// Test with invalid intensity
	_, err = generator.GeneratePhoticStimulation(10.0, 1.0, -0.1)
	if err != ErrInvalidIntensity {
		t.Errorf("GeneratePhoticStimulation() with -0.1 intensity = %v, want ErrInvalidIntensity", err)
	}
}

// Test NeuroFeedbackProcessor Initialization
func TestNewNeuroFeedbackProcessor(t *testing.T) {
	processor := NewNeuroFeedbackProcessor(250)

	if processor.SampleRate != 250 {
		t.Errorf("SampleRate = %d, want 250", processor.SampleRate)
	}

	if processor.TargetFrequency != 10.0 {
		t.Errorf("TargetFrequency = %f, want 10.0", processor.TargetFrequency)
	}

	// Test with default sample rate
	processorDefault := NewNeuroFeedbackProcessor(0)
	if processorDefault.SampleRate != DefaultNeuroSampleRate {
		t.Errorf("Default SampleRate = %d, want %d", processorDefault.SampleRate, DefaultNeuroSampleRate)
	}
}

// Test ProcessEEGFeedback
func TestProcessEEGFeedback(t *testing.T) {
	processor := NewNeuroFeedbackProcessor(250)

	// Test with insufficient data
	feedback := processor.ProcessEEGFeedback([]float64{0.1, 0.2, 0.3})
	if feedback.FrequencyLock != 0.0 {
		t.Errorf("FrequencyLock with insufficient data = %f, want 0.0", feedback.FrequencyLock)
	}

	// Test with sufficient data
	eegData := make([]float64, 200)
	for i := 0; i < 200; i++ {
		eegData[i] = 0.1 * float64(i%10)
	}

	feedback = processor.ProcessEEGFeedback(eegData)
	if feedback.FrequencyLock < 0.0 || feedback.FrequencyLock > 1.0 {
		t.Errorf("FrequencyLock = %f, want between 0.0 and 1.0", feedback.FrequencyLock)
	}
}

// Test ProcessHeartRateFeedback
func TestProcessHeartRateFeedback(t *testing.T) {
	processor := NewNeuroFeedbackProcessor(250)

	// Test with insufficient data
	feedback := processor.ProcessHeartRateFeedback([]float64{60.0, 62.0})
	if feedback.HRV != 0.5 {
		t.Errorf("HRV with insufficient data = %f, want 0.5", feedback.HRV)
	}

	// Test with sufficient data
	hrData := []float64{60.0, 62.0, 65.0, 63.0, 61.0, 64.0, 66.0, 62.0, 63.0, 65.0, 61.0}
	feedback = processor.ProcessHeartRateFeedback(hrData)

	if feedback.HRV < 0.0 || feedback.HRV > 1.0 {
		t.Errorf("HRV = %f, want between 0.0 and 1.0", feedback.HRV)
	}

	if feedback.HeartRate < 50.0 || feedback.HeartRate > 100.0 {
		t.Errorf("HeartRate = %f, want between 50.0 and 100.0", feedback.HeartRate)
	}
}

// Test CalculateAdaptationParameters
func TestCalculateAdaptationParameters(t *testing.T) {
	processor := NewNeuroFeedbackProcessor(250)

	// Test with low effectiveness
	feedback := NeuroFeedbackData{
		FrequencyLock:  0.2,
		PhaseCoherence: 0.2,
		HRV:           0.2,
	}

	adaptations := processor.CalculateAdaptationParameters(feedback)

	if adaptations.Effectiveness < 0.0 || adaptations.Effectiveness >= 0.3 {
		t.Errorf("Low effectiveness = %f, want < 0.3", adaptations.Effectiveness)
	}

	// Test with moderate effectiveness
	feedback = NeuroFeedbackData{
		FrequencyLock:  0.4,
		PhaseCoherence: 0.4,
		HRV:           0.4,
	}

	adaptations = processor.CalculateAdaptationParameters(feedback)

	if adaptations.Effectiveness < 0.3 || adaptations.Effectiveness >= 0.6 {
		t.Errorf("Moderate effectiveness = %f, want between 0.3 and 0.6", adaptations.Effectiveness)
	}

	// Test with good effectiveness
	feedback = NeuroFeedbackData{
		FrequencyLock:  0.8,
		PhaseCoherence: 0.8,
		HRV:           0.8,
	}

	adaptations = processor.CalculateAdaptationParameters(feedback)

	if adaptations.Effectiveness < 0.6 || adaptations.Effectiveness > 1.0 {
		t.Errorf("Good effectiveness = %f, want >= 0.6", adaptations.Effectiveness)
	}
}

// Test BrainWaveEntrainmentSystem Initialization
func TestNewBrainWaveEntrainmentSystem(t *testing.T) {
	system := NewBrainWaveEntrainmentSystem()

	if system.SignalGenerator == nil {
		t.Error("SignalGenerator is nil")
	}

	if system.NeuroFeedback == nil {
		t.Error("NeuroFeedback is nil")
	}

	if len(system.ActiveSessions) != 0 {
		t.Error("ActiveSessions should be empty")
	}

	if system.State != EntrainmentStateInactive {
		t.Errorf("State = %s, want inactive", system.State)
	}
}

// Test CreateEntrainmentSession
func TestCreateEntrainmentSession(t *testing.T) {
	system := NewBrainWaveEntrainmentSystem()

	sessionID, err := system.CreateEntrainmentSession(
		ModeBinauralBeats,
		FrequencyAlpha,
		300.0, // 5 minutes
		0.7,
		true,
	)

	if err != nil {
		t.Errorf("CreateEntrainmentSession() error = %v", err)
	}

	if len(sessionID) == 0 {
		t.Error("CreateEntrainmentSession() returned empty session ID")
	}

	// Check session was added
	if len(system.ActiveSessions) != 1 {
		t.Errorf("ActiveSessions count = %d, want 1", len(system.ActiveSessions))
	}

	// Test with invalid intensity
	_, err = system.CreateEntrainmentSession(
		ModeBinauralBeats,
		FrequencyAlpha,
		300.0,
		1.5,
		true,
	)

	if err != ErrInvalidIntensity {
		t.Errorf("CreateEntrainmentSession() with 1.5 intensity = %v, want ErrInvalidIntensity", err)
	}
}

// Test StartEntrainment
func TestStartEntrainment(t *testing.T) {
	system := NewBrainWaveEntrainmentSystem()

	sessionID, _ := system.CreateEntrainmentSession(
		ModeBinauralBeats,
		FrequencyAlpha,
		300.0,
		0.7,
		true,
	)

	success, err := system.StartEntrainment(sessionID)

	if !success {
		t.Errorf("StartEntrainment() returned false, error = %v", err)
	}

	if system.State != EntrainmentStateActive {
		t.Errorf("State = %s, want active", system.State)
	}

	// Test with non-existent session
	_, err = system.StartEntrainment("non_existent_session")
	if err != ErrSessionNotFound {
		t.Errorf("StartEntrainment() with non-existent session = %v, want ErrSessionNotFound", err)
	}
}

// Test StopEntrainment
func TestStopEntrainment(t *testing.T) {
	system := NewBrainWaveEntrainmentSystem()

	sessionID, _ := system.CreateEntrainmentSession(
		ModeBinauralBeats,
		FrequencyAlpha,
		300.0,
		0.7,
		true,
	)

	system.StartEntrainment(sessionID)

	success, err := system.StopEntrainment(sessionID)

	if !success {
		t.Errorf("StopEntrainment() returned false, error = %v", err)
	}

	// Check session was moved to history
	if len(system.ActiveSessions) != 0 {
		t.Errorf("ActiveSessions count = %d, want 0", len(system.ActiveSessions))
	}

	if len(system.SessionHistory) != 1 {
		t.Errorf("SessionHistory count = %d, want 1", len(system.SessionHistory))
	}

	if system.State != EntrainmentStateInactive {
		t.Errorf("State = %s, want inactive", system.State)
	}
}

// Test GenerateEntrainmentSignal
func TestGenerateEntrainmentSignal(t *testing.T) {
	system := NewBrainWaveEntrainmentSystem()

	sessionID, _ := system.CreateEntrainmentSession(
		ModeBinauralBeats,
		FrequencyAlpha,
		300.0,
		0.7,
		true,
	)

	// Generate without starting
	_, _, err := system.GenerateEntrainmentSignal(sessionID)
	if err != ErrSessionInactive {
		t.Errorf("GenerateEntrainmentSignal() without starting = %v, want ErrSessionInactive", err)
	}

	system.StartEntrainment(sessionID)

	left, right, err := system.GenerateEntrainmentSignal(sessionID)

	if err != nil {
		t.Errorf("GenerateEntrainmentSignal() error = %v", err)
	}

	if len(left) == 0 || len(right) == 0 {
		t.Error("GenerateEntrainmentSignal() returned empty signals")
	}
}

// Test GenerateIsochronicSignal
func TestGenerateIsochronicSignal(t *testing.T) {
	system := NewBrainWaveEntrainmentSystem()

	sessionID, _ := system.CreateEntrainmentSession(
		ModeIsochronicTones,
		FrequencyAlpha,
		300.0,
		0.7,
		true,
	)

	system.StartEntrainment(sessionID)

	signal, err := system.GenerateIsochronicSignal(sessionID)

	if err != nil {
		t.Errorf("GenerateIsochronicSignal() error = %v", err)
	}

	if len(signal) == 0 {
		t.Error("GenerateIsochronicSignal() returned empty signal")
	}
}

// Test ProcessNeuroFeedback
func TestProcessNeuroFeedback(t *testing.T) {
	system := NewBrainWaveEntrainmentSystem()

	sessionID, _ := system.CreateEntrainmentSession(
		ModeNeurofeedback,
		FrequencyAlpha,
		300.0,
		0.7,
		true,
	)

	system.StartEntrainment(sessionID)

	eegData := make([]float64, 200)
	for i := 0; i < 200; i++ {
		eegData[i] = 0.1 * float64(i%10)
	}

	adaptations, err := system.ProcessNeuroFeedback(sessionID, eegData, nil)

	if err != nil {
		t.Errorf("ProcessNeuroFeedback() error = %v", err)
	}

	// With adaptive mode, should return adaptations
	if adaptations.Effectiveness < 0.0 || adaptations.Effectiveness > 1.0 {
		t.Errorf("Effectiveness = %f, want between 0.0 and 1.0", adaptations.Effectiveness)
	}
}

// Test CalculateQuantumNeuralAlignment
func TestCalculateQuantumNeuralAlignment(t *testing.T) {
	system := NewBrainWaveEntrainmentSystem()

	// Test with nil metrics
	alignment := system.CalculateQuantumNeuralAlignment(nil)
	if alignment != 0.5 {
		t.Errorf("Alignment with nil metrics = %f, want 0.5", alignment)
	}

	// Test with valid metrics
	quantumMetrics := map[string]float64{
		"quantum_coherence":   0.8,
		"consciousness_level": 0.9,
	}

	alignment = system.CalculateQuantumNeuralAlignment(quantumMetrics)

	if alignment < 0.0 || alignment > 1.0 {
		t.Errorf("Alignment = %f, want between 0.0 and 1.0", alignment)
	}
}

// Test GetEntrainmentStatus
func TestGetEntrainmentStatus(t *testing.T) {
	system := NewBrainWaveEntrainmentSystem()

	status := system.GetEntrainmentStatus()

	if status["state"] != "inactive" {
		t.Errorf("Status state = %v, want inactive", status["state"])
	}

	if status["active_sessions"].(int) != 0 {
		t.Errorf("Active sessions = %v, want 0", status["active_sessions"])
	}

	// Add a session
	system.CreateEntrainmentSession(ModeBinauralBeats, FrequencyAlpha, 300.0, 0.7, true)
	status = system.GetEntrainmentStatus()

	if status["active_sessions"].(int) != 1 {
		t.Errorf("Active sessions after adding = %v, want 1", status["active_sessions"])
	}
}

// Test GetEntrainmentRecommendations
func TestGetEntrainmentRecommendations(t *testing.T) {
	system := NewBrainWaveEntrainmentSystem()

	cognitiveState := map[string]float64{
		"attention":  0.3,
		"stress":    0.7,
		"creativity": 0.3,
		"energy":    0.3,
	}

	recommendations := system.GetEntrainmentRecommendations(cognitiveState)

	if len(recommendations) == 0 {
		t.Error("GetEntrainmentRecommendations() returned empty recommendations")
	}

	// Should have 4 recommendations (one for each condition)
	if len(recommendations) != 4 {
		t.Errorf("Recommendations count = %d, want 4", len(recommendations))
	}
}

// Test EntrainmentSystemState String()
func TestEntrainmentSystemStateString(t *testing.T) {
	tests := []struct {
		state   EntrainmentSystemState
		expected string
	}{
		{EntrainmentStateInactive, "inactive"},
		{EntrainmentStateActive, "active"},
		{EntrainmentStateAdaptive, "adaptive"},
		{EntrainmentSystemState(99), "unknown"},
	}

	for _, tt := range tests {
		result := tt.state.String()
		if result != tt.expected {
			t.Errorf("EntrainmentSystemState(%d).String() = %s, want %s", tt.state, result, tt.expected)
		}
	}
}

// Test ToJSON Serialization
func TestEntrainmentSystemToJSON(t *testing.T) {
	system := NewBrainWaveEntrainmentSystem()

	jsonData, err := system.ToJSON()
	if err != nil {
		t.Errorf("ToJSON() error = %v", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Errorf("JSON unmarshal error = %v", err)
	}

	if result["state"] != "inactive" {
		t.Errorf("JSON state = %v, want inactive", result["state"])
	}
}

// Test GenerateIsocronicSignal Error Cases
func TestGenerateIsochronicSignalErrorCases(t *testing.T) {
	generator := NewBrainWaveGenerator(44100)

	// Test with invalid frequency
	_, err := generator.GenerateIsochronicTones(0, 1.0, 0.5)
	if err != ErrInvalidFrequency {
		t.Errorf("GenerateIsochronicTones() with 0 frequency = %v, want ErrInvalidFrequency", err)
	}

	// Test with invalid intensity
	_, err = generator.GenerateIsochronicTones(10.0, 1.0, 1.1)
	if err != ErrInvalidIntensity {
		t.Errorf("GenerateIsochronicTones() with 1.1 intensity = %v, want ErrInvalidIntensity", err)
	}
}

// Test GeneratePhoticStimulation Error Cases
func TestGeneratePhoticStimulationErrorCases(t *testing.T) {
	generator := NewBrainWaveGenerator(44100)

	// Test with invalid frequency
	_, err := generator.GeneratePhoticStimulation(0, 1.0, 0.5)
	if err != ErrInvalidFrequency {
		t.Errorf("GeneratePhoticStimulation() with 0 frequency = %v, want ErrInvalidFrequency", err)
	}

	// Test with invalid intensity
	_, err = generator.GeneratePhoticStimulation(10.0, 1.0, 1.1)
	if err != ErrInvalidIntensity {
		t.Errorf("GeneratePhoticStimulation() with 1.1 intensity = %v, want ErrInvalidIntensity", err)
	}
}

// Test generateSessionID
func TestGenerateSessionID(t *testing.T) {
	sessionID1 := generateSessionID()
	sessionID2 := generateSessionID()

	// Should generate unique IDs
	if sessionID1 == sessionID2 {
		t.Error("Session IDs should be unique")
	}

	// Should start with "entrain_"
	if len(sessionID1) < 9 || sessionID1[:8] != "entrain_" {
		t.Errorf("Session ID = %s, want to start with 'entrain_'", sessionID1)
	}
}

// Test SmoothSignal Edge Cases
func TestSmoothSignalEdgeCases(t *testing.T) {
	generator := NewBrainWaveGenerator(44100)

	// Test with empty signal
	signal := generator.smoothSignal([]float64{}, 5)
	if len(signal) != 0 {
		t.Error("smoothSignal() with empty input should return empty")
	}

	// Test with short signal
	signal = generator.smoothSignal([]float64{1.0, 2.0}, 10)
	if len(signal) != 2 {
		t.Error("smoothSignal() with short signal should return original length")
	}
}

// Test EntrainmentSession Serialization
func TestEntrainmentSessionSerialization(t *testing.T) {
	session := &EntrainmentSession{
		SessionID:        "test_session",
		Mode:            ModeBinauralBeats,
		TargetFrequency: FrequencyAlpha,
		Duration:        300.0,
		Intensity:       0.7,
		CarrierFrequency: 200.0,
		ModulationDepth: 0.8,
		RampUpTime:       30.0,
		RampDownTime:     30.0,
		PhaseShift:       0.0,
		StereoSeparation: 1.0,
		AdaptiveMode:      true,
		EEGFeedbackWeight: 0.7,
		HRFeedbackWeight:  0.3,
		Timestamp:        float64(time.Now().UnixNano()) / 1e9,
		IsActive:         false,
	}

	jsonData, err := json.Marshal(session)
	if err != nil {
		t.Errorf("EntrainmentSession JSON marshal error = %v", err)
	}

	var result EntrainmentSession
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Errorf("EntrainmentSession JSON unmarshal error = %v", err)
	}

	if result.SessionID != session.SessionID {
		t.Errorf("SessionID = %s, want %s", result.SessionID, session.SessionID)
	}

	if result.Mode != session.Mode {
		t.Errorf("Mode = %d, want %d", result.Mode, session.Mode)
	}

	if result.Intensity != session.Intensity {
		t.Errorf("Intensity = %f, want %f", result.Intensity, session.Intensity)
	}
}
