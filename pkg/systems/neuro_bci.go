/*
NeuralBlitz v50 Go Implementation - Neuro-BCI Interface
=======================================================

Direct brain-computer interface integration with biological neural systems
for real-time cognitive state monitoring and synchronization.

Implementation Date: 2026-02-06
Phase: Neuro-Symbiotic Integration - N1 Implementation

This package provides:
- Brain wave band definitions (Delta, Theta, Alpha, Beta, Gamma, Lambda)
- Neurochemical type definitions for emotion simulation
- Neural signal measurement structures
- Cognitive state assessment
- EEG signal simulation
- BCI backend for real-time neural monitoring
*/

package systems

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

// BrainWaveBand represents brain wave frequency bands
type BrainWaveBand int

const (
	BrainWaveBandDelta BrainWaveBand = iota
	BrainWaveBandTheta
	BrainWaveBandAlpha
	BrainWaveBandBeta
	BrainWaveBandGamma
	BrainWaveBandLambda
)

// BrainWaveBandInfo contains frequency range and name for a brain wave band
type BrainWaveBandInfo struct {
	MinFreq   float64
	MaxFreq   float64
	Name      string
}

var brainWaveBandInfo = []BrainWaveBandInfo{
	{0.5, 4.0, "delta"},   // Delta: Deep sleep
	{4.0, 8.0, "theta"},  // Theta: Drowsiness, meditation
	{8.0, 13.0, "alpha"}, // Alpha: Relaxed awareness
	{13.0, 30.0, "beta"}, // Beta: Active thinking
	{30.0, 100.0, "gamma"}, // Gamma: Higher cognition
	{100.0, 200.0, "lambda"}, // Lambda: Advanced processing
}

// String returns the name of the brain wave band
func (b BrainWaveBand) String() string {
	if b >= 0 && int(b) < len(brainWaveBandInfo) {
		return brainWaveBandInfo[b].Name
	}
	return "unknown"
}

// FrequencyRange returns the min and max frequencies for the band
func (b BrainWaveBand) FrequencyRange() (float64, float64) {
	if b >= 0 && int(b) < len(brainWaveBandInfo) {
		return brainWaveBandInfo[b].MinFreq, brainWaveBandInfo[b].MaxFreq
	}
	return 0, 0
}

// NeurochemicalType represents neurochemical systems for emotion simulation
type NeurochemicalType int

const (
	NeurochemicalDopamine NeurochemicalType = iota
	NeurochemicalSerotonin
	NeurochemicalNorepinephrine
	NeurochemicalGABA
	NeurochemicalAcetylcholine
	NeurochemicalEndorphin
	NeurochemicalCortisol
)

var neurochemicalNames = []string{
	"dopamine",     // Reward, motivation
	"serotonin",    // Mood, happiness
	"norepinephrine", // Alertness, stress
	"gaba",         // Inhibition, calm
	"acetylcholine", // Learning, memory
	"endorphin",    // Pain relief, euphoria
	"cortisol",     // Stress response
}

// String returns the name of the neurochemical
func (n NeurochemicalType) String() string {
	if n >= 0 && int(n) < len(neurochemicalNames) {
		return neurochemicalNames[n]
	}
	return "unknown"
}

// NeuralSignal represents a single neural signal measurement
type NeuralSignal struct {
	Timestamp      float64            `json:"timestamp"`
	ChannelID      string             `json:"channel_id"`
	Voltage        float64            `json:"voltage"`
	FrequencyBands map[string]float64 `json:"frequency_bands"`
	Phase          float64            `json:"phase"`
	Amplitude      float64            `json:"amplitude"`
	QualityScore   float64            `json:"quality_score"`
}

// ToJSON converts NeuralSignal to JSON string
func (n *NeuralSignal) ToJSON() (string, error) {
	data, err := json.Marshal(n)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// NeurochemicalState represents current neurochemical concentration state
type NeurochemicalState struct {
	Dopamine       float64 `json:"dopamine"`
	Serotonin     float64 `json:"serotonin"`
	Norepinephrine float64 `json:"norepinephrine"`
	GABA          float64 `json:"gaba"`
	Acetylcholine float64 `json:"acetylcholine"`
	Endorphin     float64 `json:"endorphin"`
	Cortisol      float64 `json:"cortisol"`
}

// NewNeurochemicalState creates a new neurochemical state with default values
func NewNeurochemicalState() *NeurochemicalState {
	return &NeurochemicalState{
		Dopamine:       0.5,
		Serotonin:     0.5,
		Norepinephrine: 0.5,
		GABA:          0.5,
		Acetylcholine: 0.5,
		Endorphin:     0.5,
		Cortisol:      0.5,
	}
}

// GetEmotionalProfile calculates emotional profile from neurochemical state
func (n *NeurochemicalState) GetEmotionalProfile() map[string]float64 {
	return map[string]float64{
		"happiness":  (n.Serotonin + n.Dopamine) / 2,
		"stress":     n.Cortisol*0.8 + n.Norepinephrine*0.2,
		"focus":      (n.Norepinephrine + n.Acetylcholine) / 2,
		"relaxation": (n.GABA + n.Serotonin) / 2,
		"motivation": n.Dopamine,
		"learning":   n.Acetylcholine,
		"pleasure":   (n.Endorphin + n.Dopamine) / 2,
	}
}

// ToJSON converts NeurochemicalState to JSON string
func (n *NeurochemicalState) ToJSON() (string, error) {
	data, err := json.Marshal(n)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// CognitiveState represents comprehensive cognitive state assessment
type CognitiveState struct {
	AttentionLevel   float64           `json:"attention_level"`
	MemoryLoad       float64           `json:"memory_load"`
	CognitiveFatigue float64           `json:"cognitive_fatigue"`
	EmotionalState   string           `json:"emotional_state"`
	ConsciousnessDepth float64         `json:"consciousness_depth"`
	NeuralCoherence  float64           `json:"neural_coherence"`
	ProcessingSpeed  float64           `json:"processing_speed"`
	CreativityLevel float64           `json:"creativity_level"`
	BrainWaveProfile map[string]float64 `json:"brain_wave_profile"`
	Neurochemicals   *NeurochemicalState `json:"neurochemicals"`
	QuantumAlignment float64          `json:"quantum_alignment"`
}

// NewCognitiveState creates a new cognitive state with default values
func NewCognitiveState() *CognitiveState {
	return &CognitiveState{
		AttentionLevel:    0.5,
		MemoryLoad:        0.5,
		CognitiveFatigue: 0.0,
		EmotionalState:   "neutral",
		ConsciousnessDepth: 0.5,
		NeuralCoherence:  0.5,
		ProcessingSpeed: 0.5,
		CreativityLevel: 0.5,
		BrainWaveProfile: make(map[string]float64),
		Neurochemicals:   NewNeurochemicalState(),
		QuantumAlignment: 0.5,
	}
}

// ToJSON converts CognitiveState to JSON string
func (c *CognitiveState) ToJSON() (string, error) {
	data, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// EEGSimulator generates realistic EEG signals for testing
type EEGSimulator struct {
	SamplingRate   int
	NumChannels    int
	ChannelNames   []string
	TimeOffset     float64
	BaseFrequencies map[string]float64
	NoiseLevel     float64
	ArtifactProb   float64
	mu             sync.RWMutex
}

// NewEEGSimulator creates a new EEG simulator
func NewEEGSimulator(samplingRate int, numChannels int) *EEGSimulator {
	channelNames := make([]string, numChannels)
	for i := 0; i < numChannels; i++ {
		channelNames[i] = fmt.Sprintf("EEG_%d", i+1)
	}

	return &EEGSimulator{
		SamplingRate:   samplingRate,
		NumChannels:    numChannels,
		ChannelNames:   channelNames,
		TimeOffset:     0.0,
		BaseFrequencies: map[string]float64{
			"delta": 2.0,
			"theta": 6.0,
			"alpha": 10.0,
			"beta":  20.0,
			"gamma": 40.0,
		},
		NoiseLevel:   0.1,
		ArtifactProb: 0.01,
	}
}

// GenerateSignal generates EEG signal based on cognitive state
func (e *EEGSimulator) GenerateSignal(duration float64, cognitiveState *CognitiveState) []*NeuralSignal {
	e.mu.Lock()
	defer e.mu.Unlock()

	numSamples := int(duration * float64(e.SamplingRate))
	signals := make([]*NeuralSignal, 0, numSamples*e.NumChannels)

	for sampleIdx := 0; sampleIdx < numSamples; sampleIdx++ {
		timestamp := e.TimeOffset + float64(sampleIdx)/float64(e.SamplingRate)

		for channelIdx := 0; channelIdx < e.NumChannels; channelIdx++ {
			channelID := e.ChannelNames[channelIdx]
			voltage := 0.0
			frequencyBands := make(map[string]float64)

			// Add brain wave components based on cognitive state
			for bandName, baseFreq := range e.BaseFrequencies {
				amplitude := e.getBandAmplitude(bandName, cognitiveState)
				freqVariation := rand.NormFloat64() * baseFreq * 0.1
				frequency := baseFreq + freqVariation
				phase := 2 * math.Pi * frequency * timestamp

				voltage += amplitude * math.Sin(phase)
				frequencyBands[bandName] = amplitude
			}

			// Add noise
			voltage += rand.NormFloat64() * e.NoiseLevel

			// Add artifacts occasionally
			if rand.Float64() < e.ArtifactProb {
				voltage += rand.NormFloat64() * 5.0
			}

			// Calculate signal properties
			amplitude := math.Abs(voltage)
			phase := math.Atan2(rand.NormFloat64()*0.01, voltage)
			quality := 0.98
			if rand.Float64() < 0.02 {
				quality = 0.5
			}

			signal := &NeuralSignal{
				Timestamp:      timestamp,
				ChannelID:      channelID,
				Voltage:        voltage,
				FrequencyBands: frequencyBands,
				Phase:          phase,
				Amplitude:      amplitude,
				QualityScore:   quality,
			}

			signals = append(signals, signal)
		}
	}

	e.TimeOffset += duration
	return signals
}

// getBandAmplitude gets amplitude for specific brain wave band based on cognitive state
func (e *EEGSimulator) getBandAmplitude(bandName string, cognitiveState *CognitiveState) float64 {
	baseAmplitudes := map[string]float64{
		"delta": 0.3,
		"theta": 0.5,
		"alpha": 0.7,
		"beta":  1.0,
		"gamma": 0.4,
	}

	amplitude := baseAmplitudes[bandName]

	// Modulate based on cognitive state
	switch bandName {
	case "delta":
		amplitude *= (1.0 - cognitiveState.AttentionLevel)
	case "theta":
		amplitude *= 1.0 - cognitiveState.AttentionLevel*0.5
	case "alpha":
		amplitude *= 1.0 - cognitiveState.CognitiveFatigue*0.3
	case "beta":
		amplitude *= cognitiveState.AttentionLevel
	case "gamma":
		amplitude *= (cognitiveState.CreativityLevel + cognitiveState.ProcessingSpeed) / 2
	}

	return amplitude
}

// BCIBackend interfaces with real EEG hardware or simulator for real-time neural monitoring
type BCIBackend struct {
	UseSimulator         bool
	SamplingRate         int
	NumChannels          int
	IsRecording          bool
	SignalBuffer         []*NeuralSignal
	CognitiveStateHistory []*CognitiveState
	EEGSimulator         *EEGSimulator
	CurrentCognitiveState *CognitiveState
	FFTWindowSize       int
	mu                  sync.RWMutex
	processingMu         sync.Mutex
	dataQueue           chan *NeuralSignal
	stopChan            chan struct{}
}

// NewBCIBackend creates a new BCI backend
func NewBCIBackend(useSimulator bool, samplingRate int) *BCIBackend {
	numChannels := 8
	if samplingRate <= 0 {
		samplingRate = 250
	}

	backend := &BCIBackend{
		UseSimulator:         useSimulator,
		SamplingRate:         samplingRate,
		NumChannels:          numChannels,
		IsRecording:          false,
		SignalBuffer:         make([]*NeuralSignal, 0, 1000),
		CognitiveStateHistory: make([]*CognitiveState, 0, 100),
		CurrentCognitiveState: NewCognitiveState(),
		FFTWindowSize:       256,
		dataQueue:           make(chan *NeuralSignal, 10000),
		stopChan:            make(chan struct{}),
	}

	if useSimulator {
	backend.EEGSimulator = NewEEGSimulator(samplingRate, numChannels)
	}

	return backend
}

// StartRecording starts real-time neural signal recording
func (b *BCIBackend) StartRecording() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.IsRecording {
		return false
	}

	b.IsRecording = true
	b.stopChan = make(chan struct{})

	if b.UseSimulator {
		go b.simulatorLoop()
	} else {
		fmt.Println("🔌 Initializing real EEG hardware...")
	}

	fmt.Println("🧠 BCI Recording Started")
	return true
}

// StopRecording stops neural signal recording
func (b *BCIBackend) StopRecording() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	if !b.IsRecording {
		return false
	}

	b.IsRecording = false
	close(b.stopChan)

	fmt.Println("🧠 BCI Recording Stopped")
	return true
}

// simulatorLoop generates simulated neural signals
func (b *BCIBackend) simulatorLoop() {
	lastTime := time.Now()

	for {
		select {
		case <-b.stopChan:
			return
		default:
			currentTime := time.Now()
			dt := currentTime.Sub(lastTime).Seconds()

			if dt >= 0.1 {
				b.processingMu.Lock()
				if b.EEGSimulator != nil {
					signals := b.EEGSimulator.GenerateSignal(dt, b.CurrentCognitiveState)
					for _, signal := range signals {
						select {
						case b.dataQueue <- signal:
						default:
						}
					}
				}
				b.processingMu.Unlock()
				lastTime = currentTime
			}

			time.Sleep(10 * time.Millisecond)
		}
	}
}

// ProcessSignals processes neural signals and updates cognitive state
func (b *BCIBackend) ProcessSignals(batchSize int) []*CognitiveState {
	processedStates := make([]*CognitiveState, 0)

	// Collect batch of signals
	signalsBatch := make([]*NeuralSignal, 0, batchSize)

	for i := 0; i < batchSize; i++ {
		select {
		case signal := <-b.dataQueue:
			signalsBatch = append(signalsBatch, signal)
			b.SignalBuffer = append(b.SignalBuffer, signal)
			if len(b.SignalBuffer) > 1000 {
				b.SignalBuffer = b.SignalBuffer[len(b.SignalBuffer)-1000:]
			}
		default:
			break
		}
	}

	if len(signalsBatch) > 0 {
		// Process signals and update cognitive state
		newCognitiveState := b.analyzeNeuralSignals(signalsBatch)

		if newCognitiveState != nil {
			b.processingMu.Lock()
			b.CurrentCognitiveState = newCognitiveState
			b.CognitiveStateHistory = append(b.CognitiveStateHistory, newCognitiveState)
			if len(b.CognitiveStateHistory) > 100 {
				b.CognitiveStateHistory = b.CognitiveStateHistory[len(b.CognitiveStateHistory)-100:]
			}
			b.processingMu.Unlock()
			processedStates = append(processedStates, newCognitiveState)
		}
	}

	return processedStates
}

// analyzeNeuralSignals analyzes neural signals to determine cognitive state
func (b *BCIBackend) analyzeNeuralSignals(signals []*NeuralSignal) *CognitiveState {
	if len(signals) == 0 {
		return b.CurrentCognitiveState
	}

	// Group signals by channel
	channelSignals := make(map[string][]*NeuralSignal)
	for _, signal := range signals {
		channelSignals[signal.ChannelID] = append(channelSignals[signal.ChannelID], signal)
	}

	// Calculate brain wave power for each channel
	brainWavePowers := make(map[string]float64)

	for _, channelSignalList := range channelSignals {
		voltages := make([]float64, len(channelSignalList))
		for i, s := range channelSignalList {
			voltages[i] = s.Voltage
		}

		// Apply FFT
		if len(voltages) >= b.FFTWindowSize {
			voltages = voltages[:b.FFTWindowSize]
		} else {
			// Pad with zeros
			padded := make([]float64, b.FFTWindowSize)
			copy(padded, voltages)
			voltages = padded
		}

		// Compute FFT (simplified - using real FFT approximation)
		fftResult := computeFFT(voltages)
		frequencies := computeFFTFrequencies(len(voltages), float64(b.SamplingRate))

		// Calculate power in each frequency band
		for band := BrainWaveBandDelta; band <= BrainWaveBandLambda; band++ {
			minFreq, maxFreq := band.FrequencyRange()
			power := 0.0
			for i, freq := range frequencies {
				if freq >= minFreq && freq <= maxFreq {
					power += fftResult[i] * fftResult[i]
				}
			}
			bandName := band.String()
			if existing, ok := brainWavePowers[bandName]; ok {
				brainWavePowers[bandName] = (existing + power) / 2
			} else {
				brainWavePowers[bandName] = power
			}
		}
	}

	// Update cognitive state based on brain wave analysis
	newState := *b.CurrentCognitiveState

	// Calculate attention from beta/alpha ratio
	if alpha, ok := brainWavePowers["alpha"]; ok {
		if beta, ok := brainWavePowers["beta"]; ok {
			if alpha > 0 {
				newState.AttentionLevel = math.Min(1.0, beta/(alpha+0.001))
			}
		}
	}

	// Calculate cognitive fatigue from theta/beta ratio
	if beta, ok := brainWavePowers["beta"]; ok {
		if theta, ok := brainWavePowers["theta"]; ok {
			if beta > 0 {
				newState.CognitiveFatigue = math.Min(1.0, theta/(beta+0.001))
			}
		}
	}

	// Update brain wave profile
	newState.BrainWaveProfile = brainWavePowers

	// Calculate consciousness depth from gamma coherence
	if gamma, ok := brainWavePowers["gamma"]; ok {
		newState.ConsciousnessDepth = math.Min(1.0, gamma/100.0)
	}

	// Neural coherence based on signal quality
	totalQuality := 0.0
	for _, signal := range signals {
		totalQuality += signal.QualityScore
	}
	newState.NeuralCoherence = totalQuality / float64(len(signals))

	// Calculate emotional state from neurochemical profile
	emotionalProfile := newState.Neurochemicals.GetEmotionalProfile()
	if stress, ok := emotionalProfile["stress"]; ok {
		if happiness, ok := emotionalProfile["happiness"]; ok {
			if stress > 0.6 {
				newState.EmotionalState = "stressed"
			} else if happiness > 0.7 {
				newState.EmotionalState = "happy"
			} else {
				newState.EmotionalState = "neutral"
			}
		}
	}

	return &newState
}

// computeFFT computes a simplified FFT (real-valued approximation)
func computeFFT(signal []float64) []float64 {
	n := len(signal)
	result := make([]complex128, n)

	for k := 0; k < n; k++ {
		var sum complex128
		for t := 0; t < n; t++ {
			angle := 2 * math.Pi * float64(k*t) / float64(n)
			sum += complex(signal[t], 0) * complex(math.Cos(angle), -math.Sin(angle))
		}
		result[k] = sum
	}

	// Return magnitude
	magnitude := make([]float64, n)
	for i, c := range result {
		magnitude[i] = math.Sqrt(real(c)*real(c) + imag(c)*imag(c))
	}

	return magnitude
}

// computeFFTFrequencies computes frequency bins for FFT
func computeFFTFrequencies(n int, samplingRate float64) []float64 {
	frequencies := make([]float64, n)
	for i := 0; i < n; i++ {
		if i < n/2 {
			frequencies[i] = float64(i) * samplingRate / float64(n)
		} else {
			frequencies[i] = float64(i-n) * samplingRate / float64(n)
		}
	}
	return frequencies
}

// GetCurrentCognitiveState returns the current cognitive state
func (b *BCIBackend) GetCurrentCognitiveState() *CognitiveState {
	b.processingMu.Lock()
	defer b.processingMu.Unlock()
	return b.CurrentCognitiveState
}

// GetSignalBuffer returns the current signal buffer
func (b *BCIBackend) GetSignalBuffer() []*NeuralSignal {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.SignalBuffer
}

// GetCognitiveStateHistory returns the cognitive state history
func (b *BCIBackend) GetCognitiveStateHistory() []*CognitiveState {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.CognitiveStateHistory
}

// ToJSON converts BCIBackend status to JSON string
func (b *BCIBackend) ToJSON() (string, error) {
	status := map[string]interface{}{
		"use_simulator":           b.UseSimulator,
		"sampling_rate":           b.SamplingRate,
		"num_channels":           b.NumChannels,
		"is_recording":            b.IsRecording,
		"signal_buffer_size":      len(b.SignalBuffer),
		"cognitive_history_size":  len(b.CognitiveStateHistory),
		"fft_window_size":        b.FFTWindowSize,
	}
	data, err := json.Marshal(status)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Error definitions
var (
	ErrNotRecording      = fmt.Errorf("not currently recording")
	ErrAlreadyRecording  = fmt.Errorf("already recording")
	ErrNoData            = fmt.Errorf("no data available")
)
