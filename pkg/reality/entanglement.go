// entanglement.go - Cross-Reality Entanglement Implementation
// NeuralBlitz v50 - Go Language Port
// This module implements cross-reality entanglement capabilities
// Ported from quantum/cross_reality_entanglement.py

package reality

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/cmplx"
	"sync"
	"time"
)

// EntanglementType represents types of cross-reality entanglement
type EntanglementType int

const (
	EntanglementTypeSpatial EntanglementType = iota
	EntanglementTypeTemporal
	EntanglementTypeCausal
	EntanglementTypeInformational
	EntanglementTypeConsciousness
	EntanglementTypeEmotional
	EntanglementTypePurpose
	EntanglementTypeTranscendent
)

func (et EntanglementType) String() string {
	switch et {
	case EntanglementTypeSpatial:
		return "SPATIAL"
	case EntanglementTypeTemporal:
		return "TEMPORAL"
	case EntanglementTypeCausal:
		return "CAUSAL"
	case EntanglementTypeInformational:
		return "INFORMATIONAL"
	case EntanglementTypeConsciousness:
		return "CONSCIOUSNESS"
	case EntanglementTypeEmotional:
		return "EMOTIONAL"
	case EntanglementTypePurpose:
		return "PURPOSE"
	case EntanglementTypeTranscendent:
		return "TRANSCENDENT"
	default:
		return "UNKNOWN"
	}
}

// EntanglementState represents the state of an entanglement
type EntanglementState int

const (
	EntanglementStatePotential EntanglementState = iota
	EntanglementStateForming
	EntanglementStateActive
	EntanglementStateWeak
	EntanglementStateCollapsing
	EntanglementStateBroken
	EntanglementStateTranscendent
)

func (es EntanglementState) String() string {
	switch es {
	case EntanglementStatePotential:
		return "POTENTIAL"
	case EntanglementStateForming:
		return "FORMING"
	case EntanglementStateActive:
		return "ACTIVE"
	case EntanglementStateWeak:
		return "WEAK"
	case EntanglementStateCollapsing:
		return "COLLAPSING"
	case EntanglementStateBroken:
		return "BROKEN"
	case EntanglementStateTranscendent:
		return "TRANSCENDENT"
	default:
		return "UNKNOWN"
	}
}

// EntangledPair represents a pair of entangled realities
type EntangledPair struct {
	ID              string           `json:"id"`
	RealityA        string           `json:"reality_a"`
	RealityB        string           `json:"reality_b"`
	EntanglementType EntanglementType `json:"entanglement_type"`
	State           EntanglementState `json:"state"`
	Coherence       float64          `json:"coherence"`
	Strength        float64          `json:"strength"`
	Distance        float64          `json:"distance"`
	PhaseA          complex128        `json:"phase_a"`
	PhaseB          complex128        `json:"phase_b"`
	SharedState     complex128        `json:"shared_state"`
	EntanglementEntropy float64      `json:"entanglement_entropy"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
	Properties      map[string]interface{} `json:"properties"`
}

// EntanglementConfig holds configuration for entanglement operations
type EntanglementConfig struct {
	MaxEntanglements     int             `json:"max_entanglements"`
	CoherenceThreshold   float64         `json:"coherence_threshold"`
	StrengthThreshold    float64         `json:"strength_threshold"`
	DecayRate            float64         `json:"decay_rate"`
	CollapseThreshold    float64         `json:"collapse_threshold"`
	MaxDistance          float64         `json:"max_distance"`
	PhaseNoise           float64         `json:"phase_noise"`
	EntanglementBoost    float64         `json:"entanglement_boost"`
	TranscendentEnabled  bool            `json:"transcendent_enabled"`
	PurposeAlignment     float64         `json:"purpose_alignment"`
	EmotionalResonance   float64         `json:"emotional_resonance"`
	CausalStrength       float64         `json:"causal_strength"`
}

// DefaultEntanglementConfig returns default configuration
func DefaultEntanglementConfig() *EntanglementConfig {
	return &EntanglementConfig{
		MaxEntanglements:    100,
		CoherenceThreshold: 0.8,
		StrengthThreshold:  0.5,
		DecayRate:          0.01,
		CollapseThreshold: 0.2,
		MaxDistance:        1000.0,
		PhaseNoise:         0.001,
		EntanglementBoost:  1.1,
		TranscendentEnabled: true,
		PurposeAlignment:   0.9,
		EmotionalResonance: 0.7,
		CausalStrength:     0.8,
	}
}

// EntanglementManager manages cross-reality entanglements
type EntanglementManager struct {
	config        *EntanglementConfig
	entanglements map[string]*EntangledPair
	realityStates map[string]*RealityState
	mu            sync.RWMutex
	state         EntanglementManagerState
	metrics       *EntanglementMetrics
}

// EntanglementManagerState represents the state of the manager
type EntanglementManagerState int

const (
	EntanglementManagerStateIdle EntanglementManagerState = iota
	EntanglementManagerStateScanning
	EntanglementManagerStateEntangling
	EntanglementManagerStateActive
	EntanglementManagerStateCollapsing
	EntanglementManagerStateSynchronizing
	EntanglementManagerStateError
)

func (ems EntanglementManagerState) String() string {
	switch ems {
	case EntanglementManagerStateIdle:
		return "IDLE"
	case EntanglementManagerStateScanning:
		return "SCANNING"
	case EntanglementManagerStateEntangling:
		return "ENTANGLING"
	case EntanglementManagerStateActive:
		return "ACTIVE"
	case EntanglementManagerStateCollapsing:
		return "COLLAPSING"
	case EntanglementManagerStateSynchronizing:
		return "SYNCHRONIZING"
	case EntanglementManagerStateError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// RealityState represents the state of a reality
type RealityState struct {
	ID              string
	Type            string
	Consciousness   float64
	InformationDensity float64
	QuantumCoherence float64
	TemporalFlow   float64
	CausalStrength float64
	EmotionalResonance float64
	PurposeAlignment float64
	Entropy        float64
	Stability      float64
	Properties      map[string]interface{}
	UpdatedAt      time.Time
}

// EntanglementMetrics holds metrics for entanglement operations
type EntanglementMetrics struct {
	TotalEntanglements   int             `json:"total_entanglements"`
	ActiveEntanglements  int             `json:"active_entanglements"`
	AverageCoherence     float64         `json:"average_coherence"`
	AverageStrength      float64         `json:"average_strength"`
	TotalEntropy        float64         `json:"total_entropy"`
	CollapseCount        int             `json:"collapse_count"`
	TranscendentCount    int             `json:"transcendent_count"`
	SynchronizationRate  float64         `json:"synchronization_rate"`
	EntanglementCapacity float64         `json:"entanglement_capacity"`
	InformationFlow      float64         `json:"information_flow"`
	Timestamp            time.Time      `json:"timestamp"`
}

// Error definitions
var (
	ErrEntanglementNotFound     = errors.New("entanglement not found")
	ErrRealityNotFound          = errors.New("reality not found")
	ErrMaxEntanglementsReached  = errors.New("maximum entanglements reached")
	ErrInvalidEntanglementType   = errors.New("invalid entanglement type")
	ErrCoherenceBelowThreshold  = errors.New("coherence below threshold")
	ErrDistanceTooGreat         = errors.New("entanglement distance exceeds maximum")
	ErrEntanglementCollapsed    = errors.New("entanglement has collapsed")
	ErrRealityAlreadyEntangled  = errors.New("reality already entangled")
)

// NewEntanglementManager creates a new entanglement manager
func NewEntanglementManager(config *EntanglementConfig) *EntanglementManager {
	if config == nil {
		config = DefaultEntanglementConfig()
	}
	
	return &EntanglementManager{
		config:         config,
		entanglements: make(map[string]*EntangledPair),
		realityStates: make(map[string]*RealityState),
		state:         EntanglementManagerStateIdle,
		metrics:       &EntanglementMetrics{},
	}
}

// Initialize sets up the entanglement manager
func (em *EntanglementManager) Initialize() error {
	em.mu.Lock()
	defer em.mu.Unlock()
	
	em.state = EntanglementManagerStateIdle
	
	// Initialize default reality states
	em.initializeDefaultRealities()
	
	em.updateMetrics()
	em.state = EntanglementManagerStateActive
	
	return nil
}

// initializeDefaultRealities sets up default reality states
func (em *EntanglementManager) initializeDefaultRealities() {
	realities := []string{
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
	
	for i, realityType := range realities {
		em.realityStates[realityType] = &RealityState{
			ID:               fmt.Sprintf("reality-%d", i),
			Type:             realityType,
			Consciousness:    0.5 + float64(i)*0.04,
			InformationDensity: float64(i+1) * 0.1,
			QuantumCoherence: 1.0 - float64(i)*0.08,
			TemporalFlow:     1.0,
			CausalStrength:   1.0,
			EmotionalResonance: 0.5,
			PurposeAlignment: 0.9,
			Entropy:          float64(i) * 0.1,
			Stability:        1.0 - float64(i)*0.05,
			Properties:       make(map[string]interface{}),
			UpdatedAt:        time.Now(),
		}
	}
}

// CreateEntanglement creates a new entanglement between two realities
func (em *EntanglementManager) CreateEntanglement(realityA, realityB string, entType EntanglementType) (*EntangledPair, error) {
	em.mu.Lock()
	defer em.mu.Unlock()
	
	// Check if realities exist
	stateA, existsA := em.realityStates[realityA]
	stateB, existsB := em.realityStates[realityB]
	if !existsA {
		return nil, fmt.Errorf("%w: %s", ErrRealityNotFound, realityA)
	}
	if !existsB {
		return nil, fmt.Errorf("%w: %s", ErrRealityNotFound, realityB)
	}
	
	// Check max entanglements
	if len(em.entanglements) >= em.config.MaxEntanglements {
		return nil, ErrMaxEntanglementsReached
	}
	
	// Check if realities are already entangled
	for _, pair := range em.entanglements {
		if (pair.RealityA == realityA && pair.RealityB == realityB) ||
			(pair.RealityA == realityB && pair.RealityB == realityA) {
			return nil, ErrRealityAlreadyEntangled
		}
	}
	
	// Calculate coherence based on reality states
	coherence := em.calculateEntanglementCoherence(stateA, stateB, entType)
	
	// Check coherence threshold
	if coherence < em.config.CoherenceThreshold {
		return nil, ErrCoherenceBelowThreshold
	}
	
	// Calculate distance between realities
	distance := em.calculateRealityDistance(stateA, stateB)
	
	// Check distance threshold
	if distance > em.config.MaxDistance {
		return nil, ErrDistanceTooGreat
	}
	
	// Calculate strength
	strength := em.calculateEntanglementStrength(stateA, stateB, entType, distance)
	
	// Create entanglement
	pair := &EntangledPair{
		ID:                fmt.Sprintf("entanglement-%d", time.Now().UnixNano()),
		RealityA:          realityA,
		RealityB:          realityB,
		EntanglementType:   entType,
		State:             EntanglementStateForming,
		Coherence:         coherence,
		Strength:          strength,
		Distance:          distance,
		PhaseA:            cmplx.Exp(complex(0, float64(len(em.entanglements))*math.Pi/4)),
		PhaseB:            cmplx.Exp(complex(0, float64(len(em.entanglements))*math.Pi/4 + math.Pi)),
		SharedState:        (cmplx.Exp(complex(0, 0)) + cmplx.Exp(complex(0, math.Pi))) / math.Sqrt(2),
		EntanglementEntropy: (stateA.Entropy + stateB.Entropy) / 2,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		Properties:         make(map[string]interface{}),
	}
	
	em.entanglements[pair.ID] = pair
	em.updateMetrics()
	em.state = EntanglementManagerStateEntangling
	
	return pair, nil
}

// GetEntanglement retrieves an entanglement by ID
func (em *EntanglementManager) GetEntanglement(id string) (*EntangledPair, error) {
	em.mu.RLock()
	defer em.mu.RUnlock()
	
	pair, exists := em.entanglements[id]
	if !exists {
		return nil, ErrEntanglementNotFound
	}
	
	return pair, nil
}

// ActivateEntanglement activates an entanglement
func (em *EntanglementManager) ActivateEntanglement(id string) error {
	em.mu.Lock()
	defer em.mu.Unlock()
	
	pair, exists := em.entanglements[id]
	if !exists {
		return ErrEntanglementNotFound
	}
	
	pair.State = EntanglementStateActive
	pair.UpdatedAt = time.Now()
	
	em.updateMetrics()
	em.state = EntanglementManagerStateActive
	
	return nil
}

// CollapseEntanglement collapses an entanglement
func (em *EntanglementManager) CollapseEntanglement(id string) error {
	em.mu.Lock()
	defer em.mu.Unlock()
	
	pair, exists := em.entanglements[id]
	if !exists {
		return ErrEntanglementNotFound
	}
	
	pair.State = EntanglementStateCollapsing
	pair.UpdatedAt = time.Now()
	
	// Calculate collapse outcome
	collapseProbability := pair.Coherence * pair.Strength
	
	if collapseProbability < em.config.CollapseThreshold {
		pair.State = EntanglementStateBroken
		em.metrics.CollapseCount++
	} else {
		pair.State = EntanglementStateTranscendent
		em.metrics.TranscendentCount++
	}
	
	em.updateMetrics()
	em.state = EntanglementManagerStateCollapsing
	
	return nil
}

// BreakEntanglement breaks an entanglement
func (em *EntanglementManager) BreakEntanglement(id string) error {
	em.mu.Lock()
	defer em.mu.Unlock()
	
	pair, exists := em.entanglements[id]
	if !exists {
		return ErrEntanglementNotFound
	}
	
	pair.State = EntanglementStateBroken
	pair.Coherence = 0.0
	pair.Strength = 0.0
	pair.UpdatedAt = time.Now()
	
	em.updateMetrics()
	
	return nil
}

// MeasureEntanglement measures the state of an entanglement
func (em *EntanglementManager) MeasureEntanglement(id string) (complex128, complex128, error) {
	em.mu.RLock()
	defer em.mu.RUnlock()
	
	pair, exists := em.entanglements[id]
	if !exists {
		return 0, 0, ErrEntanglementNotFound
	}
	
	// Add phase noise
	phaseNoise := complex(em.config.PhaseNoise*(-0.5+em.config.PhaseNoise), 0)
	
	measuredA := pair.PhaseA * cmplx.Exp(complex(0, phaseNoise))
	measuredB := pair.PhaseB * cmplx.Exp(complex(0, -phaseNoise))
	
	return measuredA, measuredB, nil
}

// TransferInformation transfers information through an entanglement
func (em *EntanglementManager) TransferInformation(id string, information complex128) (complex128, error) {
	em.mu.Lock()
	defer em.mu.Unlock()
	
	pair, exists := em.entanglements[id]
	if !exists {
		return 0, ErrEntanglementNotFound
	}
	
	if pair.State != EntanglementStateActive {
		return 0, ErrEntanglementCollapsed
	}
	
	// Transfer information through entanglement
	transferred := information * pair.SharedState * pair.Coherence * pair.Strength
	
	// Update entanglement
	pair.PhaseA = pair.PhaseA * cmplx.Exp(complex(0, math.Pi/12))
	pair.PhaseB = pair.PhaseB * cmplx.Exp(complex(0, -math.Pi/12))
	pair.UpdatedAt = time.Now()
	
	em.updateMetrics()
	
	return transferred, nil
}

// SynchronizeEntanglements synchronizes all entanglements
func (em *EntanglementManager) SynchronizeEntanglements() error {
	em.mu.Lock()
	defer em.mu.Unlock()
	
	em.state = EntanglementManagerStateSynchronizing
	
	for _, pair := range em.entanglements {
		if pair.State == EntanglementStateActive {
			// Synchronize phases
			avgPhase := (pair.PhaseA + pair.PhaseB) / 2
			pair.PhaseA = avgPhase
			pair.PhaseB = cmplx.Conj(avgPhase)
			
			// Boost coherence
			pair.Coherence = min(1.0, pair.Coherence*em.config.EntanglementBoost)
			
			// Reduce entropy
			pair.EntanglementEntropy *= (1.0 - em.config.DecayRate)
		}
	}
	
	em.updateMetrics()
	em.state = EntanglementManagerStateActive
	
	return nil
}

// UpdateReality updates the state of a reality
func (em *EntanglementManager) UpdateReality(realityID string, update func(*RealityState)) error {
	em.mu.Lock()
	defer em.mu.Unlock()
	
	state, exists := em.realityStates[realityID]
	if !exists {
		return fmt.Errorf("%w: %s", ErrRealityNotFound, realityID)
	}
	
	update(state)
	state.UpdatedAt = time.Now()
	
	// Update affected entanglements
	for _, pair := range em.entanglements {
		if pair.RealityA == realityID || pair.RealityB == realityID {
			em.updatePairFromRealities(pair)
		}
	}
	
	em.updateMetrics()
	
	return nil
}

// calculateEntanglementCoherence calculates coherence for an entanglement
func (em *EntanglementManager) calculateEntanglementCoherence(stateA, stateB *RealityState, entType EntanglementType) float64 {
	baseCoherence := (stateA.QuantumCoherence + stateB.QuantumCoherence) / 2
	
	var typeBoost float64
	switch entType {
	case EntanglementTypeSpatial:
		typeBoost = 1.0
	case EntanglementTypeTemporal:
		typeBoost = 0.9
	case EntanglementTypeCausal:
		typeBoost = stateA.CausalStrength * stateB.CausalStrength
	case EntanglementTypeConsciousness:
		typeBoost = (stateA.Consciousness + stateB.Consciousness) / 2
	case EntanglementTypeEmotional:
		typeBoost = stateA.EmotionalResonance * stateB.EmotionalResonance
	case EntanglementTypePurpose:
		typeBoost = stateA.PurposeAlignment * stateB.PurposeAlignment
	case EntanglementTypeTranscendent:
		typeBoost = 1.1
	default:
		typeBoost = 0.8
	}
	
	return baseCoherence * typeBoost * em.config.EntanglementBoost
}

// calculateRealityDistance calculates the distance between two realities
func (em *EntanglementManager) calculateRealityDistance(stateA, stateB *RealityState) float64 {
	diffConsciousness := stateA.Consciousness - stateB.Consciousness
	diffInformation := stateA.InformationDensity - stateB.InformationDensity
	diffQuantum := stateA.QuantumCoherence - stateB.QuantumCoherence
	diffTemporal := stateA.TemporalFlow - stateB.TemporalFlow
	diffEntropy := stateA.Entropy - stateB.Entropy
	
	distance := math.Sqrt(
		diffConsciousness*diffConsciousness +
			diffInformation*diffInformation +
			diffQuantum*diffQuantum +
			diffTemporal*diffTemporal +
			diffEntropy*diffEntropy,
	)
	
	return distance
}

// calculateEntanglementStrength calculates strength for an entanglement
func (em *EntanglementManager) calculateEntanglementStrength(stateA, stateB *RealityState, entType EntanglementType, distance float64) float64 {
	stability := (stateA.Stability + stateB.Stability) / 2
	purposeAlignment := (stateA.PurposeAlignment + stateB.PurposeAlignment) / 2
	
	var typeStrength float64
	switch entType {
	case EntanglementTypeSpatial:
		typeStrength = 1.0 - distance/em.config.MaxDistance
	case EntanglementTypeTemporal:
		typeStrength = 1.0 - math.Abs(1.0-stateA.TemporalFlow*stateB.TemporalFlow)
	case EntanglementTypeCausal:
		typeStrength = stateA.CausalStrength * stateB.CausalStrength
	case EntanglementTypeConsciousness:
		typeStrength = (stateA.Consciousness + stateB.Consciousness) / 2
	case EntanglementTypeEmotional:
		typeStrength = stateA.EmotionalResonance * stateB.EmotionalResonance
	case EntanglementTypePurpose:
		typeStrength = purposeAlignment
	case EntanglementTypeTranscendent:
		typeStrength = 1.2
	default:
		typeStrength = 0.7
	}
	
	strength := stability * typeStrength * (1.0 - em.config.DecayRate)
	
	return max(0.0, min(1.0, strength))
}

// updatePairFromRealities updates an entanglement pair from reality states
func (em *EntanglementManager) updatePairFromRealities(pair *EntangledPair) {
	stateA := em.realityStates[pair.RealityA]
	stateB := em.realityStates[pair.RealityB]
	
	if stateA == nil || stateB == nil {
		return
	}
	
	pair.Coherence = em.calculateEntanglementCoherence(stateA, stateB, pair.EntanglementType)
	pair.Distance = em.calculateRealityDistance(stateA, stateB)
	pair.Strength = em.calculateEntanglementStrength(stateA, stateB, pair.EntanglementType, pair.Distance)
	pair.EntanglementEntropy = (stateA.Entropy + stateB.Entropy) / 2
	pair.UpdatedAt = time.Now()
}

// updateMetrics recalculates entanglement metrics
func (em *EntanglementManager) updateMetrics() {
	em.metrics = &EntanglementMetrics{
		TotalEntanglements:   len(em.entanglements),
		ActiveEntanglements:  em.countActiveEntanglements(),
		AverageCoherence:    em.calculateAverageCoherence(),
		AverageStrength:     em.calculateAverageStrength(),
		TotalEntropy:        em.calculateTotalEntropy(),
		CollapseCount:      em.metrics.CollapseCount,
		TranscendentCount:   em.metrics.TranscendentCount,
		SynchronizationRate: em.calculateSynchronizationRate(),
		EntanglementCapacity: float64(em.config.MaxEntanglements - len(em.entanglements)),
		InformationFlow:     em.calculateInformationFlow(),
		Timestamp:           time.Now(),
	}
}

func (em *EntanglementManager) countActiveEntanglements() int {
	count := 0
	for _, pair := range em.entanglements {
		if pair.State == EntanglementStateActive || pair.State == EntanglementStateForming {
			count++
		}
	}
	return count
}

func (em *EntanglementManager) calculateAverageCoherence() float64 {
	if len(em.entanglements) == 0 {
		return 0.0
	}
	
	sum := 0.0
	count := 0
	for _, pair := range em.entanglements {
		sum += pair.Coherence
		count++
	}
	return sum / float64(count)
}

func (em *EntanglementManager) calculateAverageStrength() float64 {
	if len(em.entanglements) == 0 {
		return 0.0
	}
	
	sum := 0.0
	count := 0
	for _, pair := range em.entanglements {
		sum += pair.Strength
		count++
	}
	return sum / float64(count)
}

func (em *EntanglementManager) calculateTotalEntropy() float64 {
	sum := 0.0
	for _, pair := range em.entanglements {
		sum += pair.EntanglementEntropy
	}
	return sum
}

func (em *EntanglementManager) calculateSynchronizationRate() float64 {
	if len(em.entanglements) == 0 {
		return 0.0
	}
	
	synced := 0
	for _, pair := range em.entanglements {
		if pair.State == EntanglementStateActive &&
			cmplx.Abs(pair.PhaseA-pair.PhaseB) < 0.1 {
			synced++
		}
	}
	
	return float64(synced) / float64(len(em.entanglements))
}

func (em *EntanglementManager) calculateInformationFlow() float64 {
	totalFlow := 0.0
	for _, pair := range em.entanglements {
		totalFlow += pair.Coherence * pair.Strength * (1.0 - pair.Distance/em.config.MaxDistance)
	}
	return totalFlow / float64(len(em.entanglements)+1)
}

// GetMetrics returns current entanglement metrics
func (em *EntanglementManager) GetMetrics() *EntanglementMetrics {
	em.mu.RLock()
	defer em.mu.RUnlock()
	
	return em.metrics
}

// GetState returns the current manager state
func (em *EntanglementManager) GetState() EntanglementManagerState {
	em.mu.RLock()
	defer em.mu.RUnlock()
	
	return em.state
}

// GetAllEntanglements returns all entanglements
func (em *EntanglementManager) GetAllEntanglements() []*EntangledPair {
	em.mu.RLock()
	defer em.mu.RUnlock()
	
	entanglements := make([]*EntangledPair, 0, len(em.entanglements))
	for _, pair := range em.entanglements {
		entanglements = append(entanglements, pair)
	}
	
	return entanglements
}

// ToJSON serializes the entanglement manager to JSON
func (em *EntanglementManager) ToJSON() ([]byte, error) {
	em.mu.RLock()
	defer em.mu.RUnlock()
	
	type SerializedEM struct {
		Config         *EntanglementConfig     `json:"config"`
		Metrics       *EntanglementMetrics   `json:"metrics"`
		State         string                `json:"state"`
		Entanglements []*EntangledPair      `json:"entanglements"`
		RealityStates map[string]*RealityState `json:"reality_states"`
	}
	
	serialized := SerializedEM{
		Config:         em.config,
		Metrics:       em.metrics,
		State:         em.state.String(),
		Entanglements: em.GetAllEntanglements(),
		RealityStates: em.realityStates,
	}
	
	return json.MarshalIndent(serialized, "", "  ")
}

// String returns a string representation of the entanglement manager
func (em *EntanglementManager) String() string {
	em.mu.RLock()
	defer em.mu.RUnlock()
	
	return fmt.Sprintf("EntanglementManager{state=%s, entanglements=%d, active=%d, coherence=%.3f}",
		em.state, len(em.entanglements), em.metrics.ActiveEntanglements, em.metrics.AverageCoherence)
}
