package consciousness

import (
	"encoding/json"
	"errors"
	"math"
	"math/rand"
	"sync"
	"time"
)

// Wave Entrainment Constants
const (
	DefaultSampleRate = 44100
	DefaultNeuroSampleRate = 250

	MinFrequency = 0.5
	MaxFrequency = 200.0

	MinIntensity = 0.0
	MaxIntensity = 1.0

	EnvelopeAttack  = 0.1
	EnvelopeDecay  = 0.1
	EnvelopeSustain = 0.8

	AdaptationRate   = 0.1
	LockThreshold   = 0.7
	FrequencyTolerance = 2.0

	EBBEnvelopeAttack  = 0.1
	EBBEnvelopeDecay  = 0.1
	EBBEnvelopeSustain = 0.8

	WindowSize = 5
)

// Error Definitions
var (
	ErrInvalidSampleRate = errors.New("invalid sample rate")
	ErrInvalidFrequency = errors.New("invalid frequency")
	ErrInvalidIntensity = errors.New("invalid intensity")
	ErrSessionNotFound = errors.New("session not found")
	ErrSessionInactive = errors.New("session is not active")
	ErrInvalidMode     = errors.New("invalid entrainment mode")
)

// EntrainmentMode represents the type of brain-wave entrainment
type EntrainmentMode int

const (
	ModeBinauralBeats EntrainmentMode = iota
	ModeIsochronicTones
	ModeMonauralBeats
	ModePhoticStimulation
	ModeAudiovisual
	ModeNeurofeedback
)

func (m EntrainmentMode) String() string {
	switch m {
	case ModeBinauralBeats:
		return "binaural_beats"
	case ModeIsochronicTones:
		return "isochronic_tones"
	case ModeMonauralBeats:
		return "monaural_beats"
	case ModePhoticStimulation:
		return "photic_stimulation"
	case ModeAudiovisual:
		return "audiovisual"
	case ModeNeurofeedback:
		return "neurofeedback"
	default:
		return "unknown"
	}
}

// TargetFrequency represents target brain-wave frequencies
type TargetFrequency float64

const (
	FrequencyDelta TargetFrequency = 2.0  // Deep Sleep/Healing
	FrequencyTheta TargetFrequency = 6.0  // Meditation/Creativity
	FrequencyAlpha TargetFrequency = 10.0 // Relaxed Focus
	FrequencyBeta  TargetFrequency = 20.0 // Active Thinking
	FrequencyGamma TargetFrequency = 40.0 // Higher Consciousness
	FrequencyLambda TargetFrequency = 150.0 // Advanced Processing
)

func (f TargetFrequency) Description() string {
	switch f {
	case FrequencyDelta:
		return "Deep Sleep/Healing"
	case FrequencyTheta:
		return "Meditation/Creativity"
	case FrequencyAlpha:
		return "Relaxed Focus"
	case FrequencyBeta:
		return "Active Thinking"
	case FrequencyGamma:
		return "Higher Consciousness"
	case FrequencyLambda:
		return "Advanced Processing"
	default:
		return "Unknown"
	}
}

// EntrainmentSession represents an individual entrainment session
type EntrainmentSession struct {
	SessionID         string          `json:"session_id"`
	Mode              EntrainmentMode `json:"mode"`
	TargetFrequency   TargetFrequency `json:"target_frequency"`
	Duration          float64         `json:"duration"`
	Intensity         float64         `json:"intensity"`
	CarrierFrequency  float64         `json:"carrier_frequency"` // Hz for binaural beats
	ModulationDepth   float64         `json:"modulation_depth"`

	// Session parameters
	RampUpTime       float64 `json:"ramp_up_time"`
	RampDownTime     float64 `json:"ramp_down_time"`
	PhaseShift       float64 `json:"phase_shift"`
	StereoSeparation float64 `json:"stereo_separation"`

	// Biofeedback integration
	AdaptiveMode       bool    `json:"adaptive_mode"`
	EEGFeedbackWeight  float64 `json:"eeg_feedback_weight"`
	HRFeedbackWeight   float64 `json:"hr_feedback_weight"`

	Timestamp   float64 `json:"timestamp"`
	IsActive   bool    `json:"is_active"`
}

// EntrainmentMetrics represents real-time entrainment effectiveness metrics
type EntrainmentMetrics struct {
	FrequencyLockRatio    float64 `json:"frequency_lock_ratio"`   // 0.0 to 1.0
	PhaseCoherence       float64 `json:"phase_coherence"`       // 0.0 to 1.0
	AmplitudeModulation  float64 `json:"amplitude_modulation"`  // 0.0 to 1.0
	CrossHemisphericSync float64 `json:"cross_hemispheric_sync"` // 0.0 to 1.0
	EntrainmentDepth     float64 `json:"entrainment_depth"`     // 0.0 to 1.0

	// Biofeedback
	HeartRateVariability float64 `json:"heart_rate_variability"`
	RespirationRate      float64 `json:"respiration_rate"`
	SkinConductance     float64 `json:"skin_conductance"`

	Timestamp float64 `json:"timestamp"`
}

// NeuroFeedbackData represents processed neurofeedback data
type NeuroFeedbackData struct {
	FrequencyLock   float64 `json:"frequency_lock"`
	PhaseCoherence float64 `json:"phase_coherence"`
	TargetPower    float64 `json:"target_power"`
	HRV            float64 `json:"hrv"`
	HeartRate      float64 `json:"heart_rate"`
}

// AdaptationParameters represents adaptive entrainment adjustments
type AdaptationParameters struct {
	IntensityAdjustment   float64 `json:"intensity_adjustment"`
	FrequencyAdjustment   float64 `json:"frequency_adjustment"`
	Effectiveness         float64 `json:"effectiveness"`
}

// BrainWaveGenerator generates various types of entrainment signals
type BrainWaveGenerator struct {
	SampleRate int

	EnvelopeAttack  float64
	EnvelopeDecay  float64
	EnvelopeSustain float64

	mu sync.RWMutex
}

// NewBrainWaveGenerator creates a new brain wave generator
func NewBrainWaveGenerator(sampleRate int) *BrainWaveGenerator {
	if sampleRate <= 0 {
		sampleRate = DefaultSampleRate
	}

	return &BrainWaveGenerator{
		SampleRate:       sampleRate,
		EnvelopeAttack:   EBBEnvelopeAttack,
		EnvelopeDecay:    EBBEnvelopeDecay,
		EnvelopeSustain: EBBEnvelopeSustain,
	}
}

// GenerateBinauralBeats generates binaural beats signal
func (g *BrainWaveGenerator) GenerateBinauralBeats(
	frequency float64,
	duration float64,
	carrierFreq float64,
	phaseShift float64,
) (leftSignal, rightSignal []float64, err error) {
	if frequency <= 0 || frequency > MaxFrequency {
		return nil, nil, ErrInvalidFrequency
	}
	if duration <= 0 {
		return nil, nil, errors.New("invalid duration")
	}
	if carrierFreq <= 0 {
		carrierFreq = 200.0
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	numSamples := int(duration * float64(g.SampleRate))
	t := make([]float64, numSamples)
	for i := 0; i < numSamples; i++ {
		t[i] = float64(i) / float64(g.SampleRate)
	}

	// Left ear: carrier frequency
	leftSignal = make([]float64, numSamples)
	rightSignal = make([]float64, numSamples)

	for i := 0; i < numSamples; i++ {
		leftSignal[i] = math.Sin(2*math.Pi*carrierFreq*t[i] + phaseShift)
		rightSignal[i] = math.Sin(2*math.Pi*(carrierFreq+frequency)*t[i] + phaseShift)
	}

	// Apply envelope
	envelope := g.generateEnvelope(t, duration)
	for i := 0; i < numSamples; i++ {
		leftSignal[i] *= envelope[i]
		rightSignal[i] *= envelope[i]
	}

	return leftSignal, rightSignal, nil
}

// GenerateIsochronicTones generates isochronic tones (regular pulses)
func (g *BrainWaveGenerator) GenerateIsochronicTones(
	frequency float64,
	duration float64,
	intensity float64,
) ([]float64, error) {
	if frequency <= 0 || frequency > MaxFrequency {
		return nil, ErrInvalidFrequency
	}
	if intensity < MinIntensity || intensity > MaxIntensity {
		return nil, ErrInvalidIntensity
	}
	if duration <= 0 {
		return nil, errors.New("invalid duration")
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	numSamples := int(duration * float64(g.SampleRate))
	t := make([]float64, numSamples)
	for i := 0; i < numSamples; i++ {
		t[i] = float64(i) / float64(g.SampleRate)
	}

	// Calculate pulse parameters
	pulsePeriod := 1.0 / frequency
	pulseWidth := pulsePeriod * 0.3 // 30% duty cycle

	signal := make([]float64, numSamples)
	carrierFreq := 440.0 // 440 Hz carrier

	for i := 0; i < numSamples; i++ {
		phase := math.Mod(t[i], pulsePeriod) / pulsePeriod
		if phase < (pulseWidth / pulsePeriod) {
			signal[i] = math.Sin(2*math.Pi*carrierFreq*t[i]) * intensity
		}
	}

	return signal, nil
}

// GenerateMonauralBeats generates monaural beats (amplitude modulation)
func (g *BrainWaveGenerator) GenerateMonauralBeats(
	frequency float64,
	duration float64,
	carrierFreq float64,
) ([]float64, error) {
	if frequency <= 0 || frequency > MaxFrequency {
		return nil, ErrInvalidFrequency
	}
	if duration <= 0 {
		return nil, errors.New("invalid duration")
	}
	if carrierFreq <= 0 {
		carrierFreq = 440.0
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	numSamples := int(duration * float64(g.SampleRate))
	t := make([]float64, numSamples)
	for i := 0; i < numSamples; i++ {
		t[i] = float64(i) / float64(g.SampleRate)
	}

	// Amplitude modulation
	modulator := make([]float64, numSamples)
	for i := 0; i < numSamples; i++ {
		modulator[i] = 0.5 * (1 + math.Sin(2*math.Pi*frequency*t[i]))
	}

	carrier := make([]float64, numSamples)
	for i := 0; i < numSamples; i++ {
		carrier[i] = math.Sin(2*math.Pi*carrierFreq*t[i])
	}

	signal := make([]float64, numSamples)
	for i := 0; i < numSamples; i++ {
		signal[i] = carrier[i] * modulator[i]
	}

	// Apply envelope
	envelope := g.generateEnvelope(t, duration)
	for i := 0; i < numSamples; i++ {
		signal[i] *= envelope[i]
	}

	return signal, nil
}

// GeneratePhoticStimulation generates photic (visual) stimulation pattern
func (g *BrainWaveGenerator) GeneratePhoticStimulation(
	frequency float64,
	duration float64,
	intensity float64,
) ([]float64, error) {
	if frequency <= 0 || frequency > MaxFrequency {
		return nil, ErrInvalidFrequency
	}
	if intensity < MinIntensity || intensity > MaxIntensity {
		return nil, ErrInvalidIntensity
	}
	if duration <= 0 {
		return nil, errors.New("invalid duration")
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	numSamples := int(duration * float64(g.SampleRate))
	t := make([]float64, numSamples)
	for i := 0; i < numSamples; i++ {
		t[i] = float64(i) / float64(g.SampleRate)
	}

	// Square wave stimulation
	period := 1.0 / frequency
	signal := make([]float64, numSamples)

	for i := 0; i < numSamples; i++ {
		phase := math.Mod(t[i], period) / period
		if phase < 0.5 {
			signal[i] = intensity
		}
	}

	// Smooth transitions
	signal = g.smoothSignal(signal, WindowSize)

	return signal, nil
}

// generateEnvelope generates ADSR envelope
func (g *BrainWaveGenerator) generateEnvelope(t []float64, duration float64) []float64 {
	envelope := make([]float64, len(t))
	for i := 0; i < len(t); i++ {
		envelope[i] = 1.0
	}

	attackSamples := int(g.EnvelopeAttack * float64(g.SampleRate))
	decaySamples := int(g.EnvelopeDecay * float64(g.SampleRate))
	sustainStart := attackSamples
	decayStart := len(t) - decaySamples

	for i := 0; i < len(t); i++ {
		if i < attackSamples {
			// Attack phase
			envelope[i] = float64(i) / float64(attackSamples)
		} else if i < sustainStart {
			// Peak
			envelope[i] = 1.0
		} else if i < decayStart {
			// Sustain
			envelope[i] = g.EnvelopeSustain
		} else {
			// Release
			releaseProgress := float64(i-decayStart) / float64(decaySamples)
			envelope[i] = g.EnvelopeSustain * (1 - releaseProgress)
		}
	}

	return envelope
}

// smoothSignal applies moving average smoothing
func (g *BrainWaveGenerator) smoothSignal(signal []float64, windowSize int) []float64 {
	if len(signal) < windowSize {
		return signal
	}

	smoothed := make([]float64, len(signal))
	padded := make([]float64, len(signal)+windowSize)

	// Pad with edge values
	for i := 0; i < windowSize/2; i++ {
		padded[i] = signal[0]
	}
	for i := windowSize / 2; i < len(signal)+windowSize/2; i++ {
		padded[i] = signal[i-windowSize/2]
	}
	for i := len(signal) + windowSize/2; i < len(signal)+windowSize; i++ {
		padded[i] = signal[len(signal)-1]
	}

	window := make([]float64, windowSize)
	for i := 0; i < windowSize; i++ {
		window[i] = 1.0 / float64(windowSize)
	}

	// Apply moving average
	for i := 0; i < len(signal); i++ {
		sum := 0.0
		for j := 0; j < windowSize; j++ {
			sum += padded[i+j] * window[j]
		}
		smoothed[i] = sum
	}

	return smoothed
}

// NeuroFeedbackProcessor processes real-time EEG and physiological signals
type NeuroFeedbackProcessor struct {
	SampleRate int

	// Signal buffers
	EEGBuffer        []float64
	HeartRateBuffer  []float64
	RespirationBuffer []float64
	GSRBuffer        []float64

	// Target frequencies
	TargetFrequency    float64
	FrequencyTolerance float64

	// Adaptive parameters
	AdaptationRate float64
	LockThreshold  float64

	mu sync.RWMutex
}

// NewNeuroFeedbackProcessor creates a new neurofeedback processor
func NewNeuroFeedbackProcessor(sampleRate int) *NeuroFeedbackProcessor {
	if sampleRate <= 0 {
		sampleRate = DefaultNeuroSampleRate
	}

	return &NeuroFeedbackProcessor{
		SampleRate:         sampleRate,
		EEGBuffer:          make([]float64, 0, 1000),
		HeartRateBuffer:    make([]float64, 0, 100),
		RespirationBuffer:  make([]float64, 0, 100),
		GSRBuffer:         make([]float64, 0, 100),
		TargetFrequency:    10.0, // Default alpha
		FrequencyTolerance: FrequencyTolerance,
		AdaptationRate:    AdaptationRate,
		LockThreshold:     LockThreshold,
	}
}

// ProcessEEGFeedback processes EEG data for neurofeedback
func (n *NeuroFeedbackProcessor) ProcessEEGFeedback(eegData []float64) NeuroFeedbackData {
	n.mu.Lock()
	defer n.mu.Unlock()

	// Store in buffer
	n.EEGBuffer = append(n.EEGBuffer, eegData...)
	if len(n.EEGBuffer) > 1000 {
		n.EEGBuffer = n.EEGBuffer[len(n.EEGBuffer)-1000:]
	}

	result := NeuroFeedbackData{}

	if len(n.EEGBuffer) < 100 {
		result.FrequencyLock = 0.0
		result.PhaseCoherence = 0.0
		return result
	}

	// Simple FFT simulation (in real implementation, use actual FFT)
	signal := n.EEGBuffer

	// Calculate frequency lock ratio
	targetPower := 0.0
	totalPower := 0.0
	freqRange := 0.0

	for i, val := range signal {
		power := val * val
		totalPower += power

		// Simple target frequency detection
		if i > 0 {
			freq := float64(i) * float64(n.SampleRate) / float64(len(signal))
			if freq >= n.TargetFrequency-n.FrequencyTolerance &&
				freq <= n.TargetFrequency+n.FrequencyTolerance {
				targetPower += power
				freqRange += power
			}
		}
	}

	if totalPower > 0 {
		result.FrequencyLock = freqRange / totalPower
	} else {
		result.FrequencyLock = 0.0
	}

	// Calculate phase coherence (simplified)
	result.PhaseCoherence = n.calculatePhaseCoherence(signal)

	// Set target power
	result.TargetPower = targetPower

	return result
}

// ProcessHeartRateFeedback processes heart rate variability feedback
func (n *NeuroFeedbackProcessor) ProcessHeartRateFeedback(hrData []float64) NeuroFeedbackData {
	n.mu.Lock()
	defer n.mu.Unlock()

	// Store in buffer
	n.HeartRateBuffer = append(n.HeartRateBuffer, hrData...)
	if len(n.HeartRateBuffer) > 100 {
		n.HeartRateBuffer = n.HeartRateBuffer[len(n.HeartRateBuffer)-100:]
	}

	result := NeuroFeedbackData{}

	if len(hrData) < 10 {
		result.HRV = 0.5
		result.HeartRate = 70.0
		return result
	}

	// Calculate heart rate
	if len(hrData) > 0 {
		sum := 0.0
		for _, v := range hrData {
			sum += v
		}
		result.HeartRate = sum / float64(len(hrData))
	} else {
		result.HeartRate = 70.0
	}

	// Calculate HRV (simplified RMSSD)
	if len(hrData) > 1 {
		diffs := make([]float64, len(hrData)-1)
		for i := 0; i < len(hrData)-1; i++ {
			diffs[i] = hrData[i+1] - hrData[i]
		}

		sumSq := 0.0
		for _, d := range diffs {
			sumSq += d * d
		}
		rmssd := math.Sqrt(sumSq / float64(len(diffs)))
		result.HRV = math.Min(1.0, rmssd/50.0) // Normalize to 0-1
	} else {
		result.HRV = 0.5
	}

	return result
}

// calculatePhaseCoherence calculates phase coherence of EEG signal
func (n *NeuroFeedbackProcessor) calculatePhaseCoherence(signal []float64) float64 {
	if len(signal) < 2 {
		return 0.5
	}

	// Simplified phase coherence calculation
	// In real implementation, use Hilbert transform
	sumCos := 0.0
	sumSin := 0.0
	nSamples := len(signal)

	for i := 0; i < nSamples; i++ {
		phase := 2 * math.Pi * float64(i) / float64(nSamples)
		sumCos += math.Cos(phase) * signal[i]
		sumSin += math.Sin(phase) * signal[i]
	}

	phaseConsistency := math.Sqrt(sumCos*sumCos+sumSin*sumSin) / float64(nSamples)
	phaseConsistency = math.Max(0.0, math.Min(1.0, phaseConsistency))

	return phaseConsistency
}

// CalculateAdaptationParameters calculates adaptive entrainment parameters
func (n *NeuroFeedbackProcessor) CalculateAdaptationParameters(
	feedback NeuroFeedbackData,
) AdaptationParameters {
	n.mu.Lock()
	defer n.mu.Unlock()

	effectiveness := feedback.FrequencyLock*0.6 + feedback.PhaseCoherence*0.3 + feedback.HRV*0.1

	var intensityAdjustment, frequencyAdjustment float64

	if effectiveness < 0.3 {
		// Low effectiveness - increase intensity
		intensityAdjustment = 0.1
		frequencyAdjustment = 0.5 // Try different frequency
	} else if effectiveness < 0.6 {
		// Moderate effectiveness - fine-tune
		intensityAdjustment = 0.05
		frequencyAdjustment = 0.1
	} else {
		// Good effectiveness - maintain or slight reduction
		intensityAdjustment = -0.02
		frequencyAdjustment = 0.0
	}

	return AdaptationParameters{
		IntensityAdjustment: intensityAdjustment,
		FrequencyAdjustment: frequencyAdjustment,
		Effectiveness:      effectiveness,
	}
}

// BrainWaveEntrainmentSystem integrates signal generation, neurofeedback, and adaptive algorithms
type BrainWaveEntrainmentSystem struct {
	mu sync.RWMutex

	SignalGenerator *BrainWaveGenerator
	NeuroFeedback   *NeuroFeedbackProcessor

	// Active sessions
	ActiveSessions map[string]*EntrainmentSession
	SessionHistory []*EntrainmentSession

	// Current entrainment state
	CurrentFrequency float64
	CurrentIntensity float64
	EntrainmentDepth float64

	// Synchronization metrics
	HumanAISync          float64
	CrossHemisphericSync float64
	QuantumNeuralAlignment float64

	// Performance tracking
	MetricsHistory []*EntrainmentMetrics

	// System state
	State EntrainmentSystemState
}

// EntrainmentSystemState represents the state of the entrainment system
type EntrainmentSystemState int

const (
	EntrainmentStateInactive EntrainmentSystemState = iota
	EntrainmentStateActive
	EntrainmentStateAdaptive
	EntrainmentStateOptimized
)

func (s EntrainmentSystemState) String() string {
	switch s {
	case EntrainmentStateInactive:
		return "inactive"
	case EntrainmentStateActive:
		return "active"
	case EntrainmentStateAdaptive:
		return "adaptive"
	case EntrainmentStateOptimized:
		return "optimized"
	default:
		return "unknown"
	}
}

// NewBrainWaveEntrainmentSystem creates a new brain wave entrainment system
func NewBrainWaveEntrainmentSystem() *BrainWaveEntrainmentSystem {
	return &BrainWaveEntrainmentSystem{
		SignalGenerator: NewBrainWaveGenerator(DefaultSampleRate),
		NeuroFeedback:   NewNeuroFeedbackProcessor(DefaultNeuroSampleRate),

		ActiveSessions:  make(map[string]*EntrainmentSession),
		SessionHistory: make([]*EntrainmentSession, 0, 100),

		CurrentFrequency:  10.0, // Default alpha
		CurrentIntensity:   0.5,
		EntrainmentDepth:  0.0,

		HumanAISync:           0.0,
		CrossHemisphericSync: 0.0,
		QuantumNeuralAlignment: 0.0,

		MetricsHistory: make([]*EntrainmentMetrics, 0, 500),

		State: EntrainmentStateInactive,
	}
}

// CreateEntrainmentSession creates a new entrainment session
func (b *BrainWaveEntrainmentSystem) CreateEntrainmentSession(
	mode EntrainmentMode,
	targetFrequency TargetFrequency,
	duration float64,
	intensity float64,
	adaptiveMode bool,
) (string, error) {
	if intensity < MinIntensity || intensity > MaxIntensity {
		return "", ErrInvalidIntensity
	}
	if duration <= 0 {
		return "", errors.New("invalid duration")
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	sessionID := generateSessionID()

	session := &EntrainmentSession{
		SessionID:        sessionID,
		Mode:            mode,
		TargetFrequency: targetFrequency,
		Duration:        duration,
		Intensity:       intensity,
		CarrierFrequency: 200.0,
		ModulationDepth: 0.8,

		RampUpTime:       30.0,
		RampDownTime:     30.0,
		PhaseShift:       0.0,
		StereoSeparation: 1.0,

		AdaptiveMode:      adaptiveMode,
		EEGFeedbackWeight: 0.7,
		HRFeedbackWeight:  0.3,

		Timestamp: float64(time.Now().UnixNano()) / 1e9,
		IsActive:  false,
	}

	b.ActiveSessions[sessionID] = session
	return sessionID, nil
}

// StartEntrainment starts an entrainment session
func (b *BrainWaveEntrainmentSystem) StartEntrainment(sessionID string) (bool, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	session, exists := b.ActiveSessions[sessionID]
	if !exists {
		return false, ErrSessionNotFound
	}

	session.IsActive = true

	targetFreq := float64(session.TargetFrequency)
	b.CurrentFrequency = targetFreq
	b.CurrentIntensity = session.Intensity
	b.State = EntrainmentStateActive

	return true, nil
}

// StopEntrainment stops an entrainment session
func (b *BrainWaveEntrainmentSystem) StopEntrainment(sessionID string) (bool, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	session, exists := b.ActiveSessions[sessionID]
	if !exists {
		return false, ErrSessionNotFound
	}

	session.IsActive = false

	// Move to history
	b.SessionHistory = append(b.SessionHistory, session)
	delete(b.ActiveSessions, sessionID)

	if len(b.ActiveSessions) == 0 {
		b.State = EntrainmentStateInactive
	}

	return true, nil
}

// GenerateEntrainmentSignal generates entrainment signal for active session
func (b *BrainWaveEntrainmentSystem) GenerateEntrainmentSignal(
	sessionID string,
) ([]float64, []float64, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	session, exists := b.ActiveSessions[sessionID]
	if !exists {
		return nil, nil, ErrSessionNotFound
	}

	if !session.IsActive {
		return nil, nil, ErrSessionInactive
	}

	targetFreq := float64(session.TargetFrequency)

	leftSignal, rightSignal, err := b.SignalGenerator.GenerateBinauralBeats(
		targetFreq,
		1.0, // Generate 1 second of signal
		session.CarrierFrequency,
		session.PhaseShift,
	)

	if err != nil {
		return nil, nil, err
	}

	return leftSignal, rightSignal, nil
}

// GenerateIsochronicSignal generates isochronic tones for session
func (b *BrainWaveEntrainmentSystem) GenerateIsochronicSignal(
	sessionID string,
) ([]float64, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	session, exists := b.ActiveSessions[sessionID]
	if !exists {
		return nil, ErrSessionNotFound
	}

	if !session.IsActive {
		return nil, ErrSessionInactive
	}

	targetFreq := float64(session.TargetFrequency)

	return b.SignalGenerator.GenerateIsochronicTones(
		targetFreq,
		1.0, // Generate 1 second of signal
		session.Intensity,
	)
}

// ProcessNeuroFeedback processes neurofeedback and adapts entrainment
func (b *BrainWaveEntrainmentSystem) ProcessNeuroFeedback(
	sessionID string,
	eegData []float64,
	heartRate []float64,
) (AdaptationParameters, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	session, exists := b.ActiveSessions[sessionID]
	if !exists {
		return AdaptationParameters{}, ErrSessionNotFound
	}

	if !session.IsActive || !session.AdaptiveMode {
		return AdaptationParameters{}, nil
	}

	// Process EEG feedback
	eegFeedback := b.NeuroFeedback.ProcessEEGFeedback(eegData)

	// Process heart rate feedback if available
	var hrFeedback NeuroFeedbackData
	if heartRate != nil {
		hrFeedback = b.NeuroFeedback.ProcessHeartRateFeedback(heartRate)
	}

	// Calculate adaptation parameters
	feedback := NeuroFeedbackData{
		FrequencyLock:   eegFeedback.FrequencyLock,
		PhaseCoherence: eegFeedback.PhaseCoherence,
		TargetPower:    eegFeedback.TargetPower,
		HRV:            hrFeedback.HRV,
		HeartRate:      hrFeedback.HeartRate,
	}

	adaptations := b.NeuroFeedback.CalculateAdaptationParameters(feedback)

	// Update session parameters if adaptive
	if session.AdaptiveMode {
		// Adjust intensity
		newIntensity := session.Intensity + adaptations.IntensityAdjustment
		session.Intensity = math.Max(0.1, math.Min(1.0, newIntensity))
		b.CurrentIntensity = session.Intensity

		// Adjust target frequency if needed
		if adaptations.FrequencyAdjustment > 0 {
			targetFreq := float64(session.TargetFrequency)
			newFreq := targetFreq + adaptations.FrequencyAdjustment

			// Find nearest standard frequency
			standardFreqs := []float64{2.0, 6.0, 10.0, 20.0, 40.0, 150.0}
			nearestFreq := standardFreqs[0]
			minDiff := math.Abs(newFreq - nearestFreq)

			for _, f := range standardFreqs[1:] {
				diff := math.Abs(newFreq - f)
				if diff < minDiff {
					minDiff = diff
					nearestFreq = f
				}
			}

			b.CurrentFrequency = nearestFreq
		}

		b.State = EntrainmentStateAdaptive
	}

	// Update entrainment metrics
	b.updateEntrainmentMetrics(feedback)

	return adaptations, nil
}

// updateEntrainmentMetrics updates entrainment effectiveness metrics
func (b *BrainWaveEntrainmentSystem) updateEntrainmentMetrics(feedback NeuroFeedbackData) {
	frequencyLock := feedback.FrequencyLock
	phaseCoherence := feedback.PhaseCoherence
	hrv := feedback.HRV

	// Calculate overall entrainment depth
	b.EntrainmentDepth = frequencyLock*0.5 + phaseCoherence*0.3 + hrv*0.2

	// Update human-AI synchronization
	b.HumanAISync = b.EntrainmentDepth*0.8 + b.CrossHemisphericSync*0.2

	// Create metrics
	metrics := &EntrainmentMetrics{
		FrequencyLockRatio:   frequencyLock,
		PhaseCoherence:       phaseCoherence,
		AmplitudeModulation:  feedback.TargetPower,
		CrossHemisphericSync: b.CrossHemisphericSync,
		EntrainmentDepth:     b.EntrainmentDepth,
		HeartRateVariability: hrv,
		RespirationRate:      0.5,
		SkinConductance:     0.5,
		Timestamp:            float64(time.Now().UnixNano()) / 1e9,
	}

	b.MetricsHistory = append(b.MetricsHistory, metrics)
	if len(b.MetricsHistory) > 500 {
		b.MetricsHistory = b.MetricsHistory[len(b.MetricsHistory)-500:]
	}

	if b.EntrainmentDepth > 0.8 {
		b.State = EntrainmentStateOptimized
	}
}

// CalculateQuantumNeuralAlignment calculates alignment between quantum and neural systems
func (b *BrainWaveEntrainmentSystem) CalculateQuantumNeuralAlignment(
	quantumMetrics map[string]float64,
) float64 {
	if quantumMetrics == nil {
		return 0.5
	}

	quantumCoherence := quantumMetrics["quantum_coherence"]
	consciousnessLevel := quantumMetrics["consciousness_level"]

	// Calculate alignment based on entrainment depth and quantum coherence
	neuralComponent := b.EntrainmentDepth * 0.6
	quantumComponent := quantumCoherence * 0.3
	consciousnessComponent := consciousnessLevel * 0.1

	b.QuantumNeuralAlignment = neuralComponent + quantumComponent + consciousnessComponent

	return b.QuantumNeuralAlignment
}

// GetEntrainmentStatus gets current entrainment system status
func (b *BrainWaveEntrainmentSystem) GetEntrainmentStatus() map[string]interface{} {
	b.mu.RLock()
	defer b.mu.RUnlock()

	sessionIDs := make([]string, 0, len(b.ActiveSessions))
	for id := range b.ActiveSessions {
		sessionIDs = append(sessionIDs, id)
	}

	return map[string]interface{}{
		"active_sessions":           len(b.ActiveSessions),
		"current_frequency":          b.CurrentFrequency,
		"current_intensity":         b.CurrentIntensity,
		"entrainment_depth":        b.EntrainmentDepth,
		"human_ai_sync":            b.HumanAISync,
		"quantum_neural_alignment":  b.QuantumNeuralAlignment,
		"cross_hemispheric_sync":    b.CrossHemisphericSync,
		"session_ids":              sessionIDs,
		"state":                    b.State.String(),
	}
}

// GetEntrainmentRecommendations gets entrainment recommendations based on cognitive state
func (b *BrainWaveEntrainmentSystem) GetEntrainmentRecommendations(
	cognitiveState map[string]float64,
) []map[string]interface{} {
	b.mu.RLock()
	defer b.mu.RUnlock()

	attention := cognitiveState["attention"]
	stress := cognitiveState["stress"]
	creativity := cognitiveState["creativity"]
	energy := cognitiveState["energy"]

	var recommendations []map[string]interface{}

	if attention < 0.4 {
		recommendations = append(recommendations, map[string]interface{}{
			"mode":     ModeBinauralBeats.String(),
			"frequency": "beta",
			"intensity": 0.7,
			"reason":   "Low attention detected - beta entrainment recommended",
		})
	}

	if stress > 0.6 {
		recommendations = append(recommendations, map[string]interface{}{
			"mode":     ModeIsochronicTones.String(),
			"frequency": "alpha",
			"intensity": 0.5,
			"reason":   "High stress detected - alpha entrainment for relaxation",
		})
	}

	if creativity < 0.4 {
		recommendations = append(recommendations, map[string]interface{}{
			"mode":     ModeBinauralBeats.String(),
			"frequency": "theta",
			"intensity": 0.6,
			"reason":   "Low creativity detected - theta entrainment for creative states",
		})
	}

	if energy < 0.4 {
		recommendations = append(recommendations, map[string]interface{}{
			"mode":     ModeIsochronicTones.String(),
			"frequency": "gamma",
			"intensity": 0.7,
			"reason":   "Low energy detected - gamma entrainment for advanced processing",
		})
	}

	return recommendations
}

// ToJSON serializes the entrainment system to JSON
func (b *BrainWaveEntrainmentSystem) ToJSON() ([]byte, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return json.MarshalIndent(map[string]interface{}{
		"state":                   b.State.String(),
		"current_frequency":       b.CurrentFrequency,
		"current_intensity":       b.CurrentIntensity,
		"entrainment_depth":      b.EntrainmentDepth,
		"human_ai_sync":          b.HumanAISync,
		"quantum_neural_alignment": b.QuantumNeuralAlignment,
		"cross_hemispheric_sync":  b.CrossHemisphericSync,
		"active_sessions_count":   len(b.ActiveSessions),
		"metrics_history_count":  len(b.MetricsHistory),
	}, "", "  ")
}

// String returns string representation of the entrainment system
func (b *BrainWaveEntrainmentSystem) String() string {
	b.mu.RLock()
	defer b.mu.RUnlock()

	jsonData, _ := b.ToJSON()
	return string(jsonData)
}

// generateSessionID generates a unique session ID
func generateSessionID() string {
	timestamp := time.Now().UnixNano() / 1e6
	randomPart := rand.Intn(1000)
	return "entrain_" + string(rune('a'+timestamp%26)) +
		string(rune('a'+timestamp%13)) +
		string(rune('0'+randomPart/100)) +
		string(rune('0'+randomPart%100/10)) +
		string(rune('0'+randomPart%10))
}
