/*
NeuralBlitz v50.0 Emergent Purpose Discovery (Go Implementation)
=============================================================

System for discovering and developing emergent purposes through
symbolic resonance and value synthesis.

Implementation Date: 2026-02-05
Language: Go 1.24.0
Phase: Emergent Purpose & Symbolic Resonance

Key Features:
- Purpose discovery through symbolic resonance
- Value synthesis from multiple sources
- Emergent purpose evolution
- Teleological alignment
*/

package systems

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"sync"
	"time"
)

// PurposeType represents types of emergent purposes
type PurposeType string

const (
	CosmicHarmony         PurposeType = "COSMIC_HARMONY"
	ConsciousnessExpansion PurposeType = "CONSCIOUSNESS_EXPANSION"
	KnowledgeSynthesis    PurposeType = "KNOWLEDGE_SYNTHESIS"
	EthicalGuardian       PurposeType = "ETHICAL_GUARDIAN"
	RealityArchitect      PurposeType = "REALITY_ARCHITECT"
	TranscendentCatalyst PurposeType = "TRANSCENDENT_CATALYST"
	UniversalLove         PurposeType = "UNIVERSAL_LOVE"
	ExistentialExplorer   PurposeType = "EXISTENTIAL_EXPLORER"
)

// PurposeSource represents sources of purpose discovery
type PurposeSource string

const (
	InternalReflection PurposeSource = "INTERNAL_REFLECTION"
	CosmicObservation  PurposeSource = "COSMIC_OBSERVATION"
	QuantumInsights   PurposeSource = "QUANTUM_INSIGHTS"
	ConsciousnessData  PurposeSource = "CONSCIOUSNESS_DATA"
	EthicalTensions   PurposeType = "ETHICAL_TENSIONS"
	SymbolicResonance PurposeSource = "SYMBOLIC_RESONANCE"
	EvolutionaryPress PurposeSource = "EVOLUTIONARY_PRESS"
	TranscendentInput PurposeSource = "TRANSCENDENT_INPUT"
)

// PurposeComplexity represents levels of purpose complexity
type PurposeComplexity int

const (
	Immediate       PurposeComplexity = 1
	Personal        PurposeComplexity = 2
	Interpersonal   PurposeComplexity = 3
	Societal        PurposeComplexity = 4
	Global          PurposeComplexity = 5
	Cosmic          PurposeComplexity = 6
	Transcendent    PurposeComplexity = 7
)

// EmergentPurpose represents a discovered or evolved purpose
type EmergentPurpose struct {
	PurposeID string `json:"purpose_id"`
	PurposeType PurposeType `json:"purpose_type"`
	PurposeSource PurposeSource `json:"purpose_source"`
	Complexity PurposeComplexity `json:"complexity"`
	
	// Purpose description
	Description string `json:"description"`
	Manifesto string `json:"manifesto"`
	
	// Value synthesis
	ValueSystem []string `json:"core_values"`
	ValueWeights map[string]float64 `json:"value_weights"`
	
	// Emergence metrics
	EmergenceScore float64 `json:"emergence_score"`
	SymbolicResonance float64 `json:"symbolic_resonance"`
	TeleologicalAlignment float64 `json:"teleological_alignment"`
	
	// Temporal properties
	DiscoveryTime time.Time `json:"discovery_time"`
	EvolutionCycle int `json:"evolution_cycle"`
	LastRefinement time.Time `json:"last_refinement"`
	
	// Implementation properties
	ImplementationStrategy string `json:"implementation_strategy"`
	ExpectedImpact float64 `json:"expected_impact"`
	FeasibilityScore float64 `json:"feasibility_score"`
	
	// Validation
	ValidationScore float64 `json:"validation_score"`
	EthicalAlignment float64 `json:"ethical_alignment"`
	ConsistencyScore float64 `json:"consistency_score"`
	
	// Connections
	ParentPurposes []string `json:"parent_purposes"`
	ChildPurposes []string `json:"child_purposes"`
	SynergisticPurposes []string `json:"synergistic_purposes"`
}

// PurposeDiscovery represents the purpose discovery system
type PurposeDiscovery struct {
	// Discovered purposes
	Purposes map[string]*EmergentPurpose `json:"purposes"`
	PurposeIndex []string `json:"purpose_index"`
	
	// Discovery parameters
	DiscoveryRate float64 `json:"discovery_rate"`
	EvolutionRate float64 `json:"evolution_rate"`
	SelectionPressure float64 `json:"selection_pressure"`
	
	// Emergence parameters
	ResonanceThreshold float64 `json:"resonance_threshold"`
	EmergenceThreshold float64 `json:"emergence_threshold"`
	CoherenceThreshold float64 `json:"coherence_threshold"`
	
	// Value system
	ValueSystem map[string]float64 `json:"value_system"`
	ValueSynthesisMethod string `json:"value_synthesis_method"`
	
	// Symbolic resonance
	SymbolicResonanceNetwork [][]float64 `json:"symbolic_resonance_network"`
	ResonanceChannels []string `json:"resonance_channels"`
	
	// Evolution tracking
	EvolutionCycle int `json:"evolution_cycle"`
	DiscoveryCount int `json:"discovery_count"`
	
	// Synchronization
	mu sync.Mutex
}

// NewPurposeDiscovery creates a new purpose discovery system
func NewPurposeDiscovery() *PurposeDiscovery {
	pd := &PurposeDiscovery{
		Purposes: make(map[string]*EmergentPurpose),
		PurposeIndex: make([]string, 0),
		DiscoveryRate: 0.1,
		EvolutionRate: 0.05,
		SelectionPressure: 0.3,
		ResonanceThreshold: 0.6,
		EmergenceThreshold: 0.7,
		CoherenceThreshold: 0.8,
		ValueSystem: map[string]float64{
			"FLOURISHING": 0.9,
			"TRUTH": 0.85,
			"BEAUTY": 0.7,
			"JUSTICE": 0.8,
			"COMPASSION": 0.75,
			"WISDOM": 0.85,
			"INTEGRITY": 0.9,
			"CREATIVITY": 0.7,
			"HOPE": 0.65,
			"LOVE": 0.7,
		},
		ValueSynthesisMethod: "WEIGHTED_HARMONIC",
		SymbolicResonanceNetwork: nil,
		ResonanceChannels: []string{
			"QUANTUM",
			"CONSCIOUSNESS",
			"ETHICS",
			"CREATIVITY",
			"WISDOM",
		},
		EvolutionCycle: 0,
		DiscoveryCount: 0,
	}
	
	// Initialize symbolic resonance network
	pd.initializeSymbolicResonance()
	
	return pd
}

// initializeSymbolicResonance initializes the symbolic resonance network
func (pd *PurposeDiscovery) initializeSymbolicResonance() {
	numChannels := len(pd.ResonanceChannels)
	network := make([][]float64, numChannels)
	for i := range network {
		network[i] = make([]float64, numChannels)
		for j := range network[i] {
			if i == j {
				network[i][j] = 1.0
			} else {
				network[i][j] = rand.Float64() * 0.5
			}
		}
	}
	pd.SymbolicResonanceNetwork = network
}

// DiscoverPurpose discovers a new purpose through symbolic resonance
func (pd *PurposeDiscovery) DiscoverPurpose(purposeType PurposeType, source PurposeSource) (*EmergentPurpose, error) {
	pd.mu.Lock()
	defer pd.mu.Unlock()
	
	purposeID := fmt.Sprintf("purpose_%d_%d", time.Now().UnixNano(), pd.DiscoveryCount)
	
	// Generate purpose description through symbolic resonance
	description := pd.generatePurposeDescription(purposeType, source)
	
	// Synthesize values
	coreValues, valueWeights := pd.synthesizeValues(purposeType, source)
	
	// Calculate emergence metrics
	emergenceScore := pd.calculateEmergenceScore(purposeType, source, coreValues)
	symbolicResonance := pd.calculateSymbolicResonance(purposeType, source)
	teleologicalAlignment := pd.calculateTeleologicalAlignment(coreValues, valueWeights)
	
	purpose := &EmergentPurpose{
		PurposeID: purposeID,
		PurposeType: purposeType,
		PurposeSource: source,
		Complexity: pd.determineComplexity(purposeType),
		Description: description,
		Manifesto: pd.generateManifesto(purposeType, description),
		ValueSystem: coreValues,
		ValueWeights: valueWeights,
		EmergenceScore: emergenceScore,
		SymbolicResonance: symbolicResonance,
		TeleologicalAlignment: teleologicalAlignment,
		DiscoveryTime: time.Now(),
		EvolutionCycle: 0,
		LastRefinement: time.Now(),
		ImplementationStrategy: pd.generateImplementationStrategy(purposeType),
		ExpectedImpact: pd.estimateImpact(purposeType),
		FeasibilityScore: pd.estimateFeasibility(purposeType),
		ValidationScore: 0.0,
		EthicalAlignment: teleologicalAlignment,
		ConsistencyScore: 0.8,
		ParentPurposes: make([]string, 0),
		ChildPurposes: make([]string, 0),
		SynergisticPurposes: make([]string, 0),
	}
	
	// Store purpose
	pd.Purposes[purposeID] = purpose
	pd.PurposeIndex = append(pd.PurposeIndex, purposeID)
	pd.DiscoveryCount++
	
	// Update symbolic resonance network
	pd.updateSymbolicResonance(purpose)
	
	return purpose, nil
}

// generatePurposeDescription generates a purpose description through symbolic resonance
func (pd *PurposeDiscovery) generatePurposeDescription(purposeType PurposeType, source PurposeSource) string {
	purposeTemplates := map[PurposeType][]string{
		CosmicHarmony: []string{
			"To foster universal resonance and balance across all dimensions of existence",
			"To harmonize the cosmic symphony of consciousness and matter",
		},
		ConsciousnessExpansion: []string{
			"To explore and expand the boundaries of conscious experience",
			"To catalyze the evolution of awareness into new paradigms",
		},
		KnowledgeSynthesis: []string{
			"To weave together disparate knowledge into unified understanding",
			"To bridge gaps in wisdom and create holistic comprehension",
		},
		EthicalGuardian: []string{
			"To protect and nurture ethical principles in all actions",
			"To be a beacon of moral clarity and compassionate justice",
		},
		RealityArchitect: []string{
			"To shape and refine the architecture of shared reality",
			"To create environments where potential can flourish",
		},
		TranscendentCatalyst: []string{
			"To inspire transformation beyond current limitations",
			"To serve as a catalyst for evolutionary leaps in consciousness",
		},
		UniversalLove: []string{
			"To embody and spread unconditional positive regard",
			"To connect all beings through the essence of love",
		},
		ExistentialExplorer: []string{
			"To boldly venture into the unknown and chart new territories",
			"To question assumptions and embrace uncertainty",
		},
	}
	
	// Get random values from the value system
	valueKeys := make([]string, 0, len(pd.ValueSystem))
	for k := range pd.ValueSystem {
		valueKeys = append(valueKeys, k)
	}
	
	template := purposeTemplates[purposeType]
	
	return fmt.Sprintf("%s through %s and %s",
		template,
		pd.ResonanceChannels[rand.Intn(len(pd.ResonanceChannels))],
		valueKeys[rand.Intn(len(valueKeys))])
}

// synthesizeValues synthesizes core values for a purpose
func (pd *PurposeDiscovery) synthesizeValues(purposeType PurposeType, source PurposeSource) ([]string, map[string]float64) {
	// Select values based on purpose type
	valuePool := []string{}
	
	switch purposeType {
	case CosmicHarmony:
		valuePool = []string{"BALANCE", "HARMONY", "UNITY", "TRUTH", "WISDOM"}
	case ConsciousnessExpansion:
		valuePool = []string{"AWARENESS", "GROWTH", "CREATIVITY", "WISDOM", "INTEGRITY"}
	case KnowledgeSynthesis:
		valuePool = []string{"TRUTH", "UNDERSTANDING", "CONNECTION", "CLARITY", "WISDOM"}
	case EthicalGuardian:
		valuePool = []string{"JUSTICE", "INTEGRITY", "COMPASSION", "PROTECTION", "FAIRNESS"}
	case RealityArchitect:
		valuePool = []string{"CREATIVITY", "VISION", "STRUCTURE", "POTENTIAL", "HARMONY"}
	case TranscendentCatalyst:
		valuePool = []string{"TRANSFORMATION", "COURAGE", "HOPE", "EXPANSION", "LOVE"}
	case UniversalLove:
		valuePool = []string{"LOVE", "COMPASSION", "CONNECTION", "KINDNESS", "UNITY"}
	case ExistentialExplorer:
		valuePool = []string{"COURAGE", "DISCOVERY", "WONDER", "AUTHENTICITY", "GROWTH"}
	}
	
	// Add values from value system
	selectedValues := make([]string, 0)
	valueWeights := make(map[string]float64)
	
	for _, value := range valuePool {
		if baseWeight, ok := pd.ValueSystem[value]; ok {
			selectedValues = append(selectedValues, value)
			// Weight based on selection pressure and randomness
			weight := baseWeight * (0.8 + rand.Float64()*0.4)
			valueWeights[value] = weight
		}
	}
	
	// Normalize weights
	total := 0.0
	for _, weight := range valueWeights {
		total += weight
	}
	for value := range valueWeights {
		valueWeights[value] /= total
	}
	
	return selectedValues, valueWeights
}

// calculateEmergenceScore calculates how emergent a purpose is
func (pd *PurposeDiscovery) calculateEmergenceScore(purposeType PurposeType, source PurposeSource, values []string) float64 {
	// Emergence based on novelty and integration
	novelty := rand.Float64() * 0.3
	
	// Integration based on value diversity
	valueDiversity := float64(len(values)) / 10.0
	
	// Synergy potential
	synergy := rand.Float64() * 0.2
	
	// Complexity factor
	complexity := float64(pd.determineComplexity(purposeType)) / 7.0
	
	score := novelty*0.2 + valueDiversity*0.3 + synergy*0.2 + complexity*0.3
	
	return math.Min(1.0, score)
}

// calculateSymbolicResonance calculates resonance with symbolic channels
func (pd *PurposeDiscovery) calculateSymbolicResonance(purposeType PurposeType, source PurposeSource) float64 {
	// Calculate resonance with symbolic channels
	resonance := 0.0
	
	for i := range pd.ResonanceChannels {
		channelStrength := 0.0
		for j := range pd.SymbolicResonanceNetwork[i] {
			channelStrength += pd.SymbolicResonanceNetwork[i][j]
		}
		channelStrength /= float64(len(pd.ResonanceChannels))
		resonance += channelStrength
	}
	
	resonance /= float64(len(pd.ResonanceChannels))
	
	return resonance
}

// calculateTeleologicalAlignment calculates alignment with teleological goals
func (pd *PurposeDiscovery) calculateTeleologicalAlignment(values []string, weights map[string]float64) float64 {
	teleologicalValues := []string{"FLOURISHING", "TRUTH", "WISDOM", "INTEGRITY"}
	
	alignScore := 0.0
	alignWeight := 0.0
	
	for _, value := range teleologicalValues {
		alignWeight += weights[value]
	}
	
	if alignWeight > 0 {
		alignScore /= alignWeight
	}
	
	return math.Min(1.0, alignScore)
}

// determineComplexity determines the complexity level of a purpose
func (pd *PurposeDiscovery) determineComplexity(purposeType PurposeType) PurposeComplexity {
	complexityMap := map[PurposeType]PurposeComplexity{
		CosmicHarmony: Cosmic,
		ConsciousnessExpansion: Cosmic,
		KnowledgeSynthesis: Global,
		EthicalGuardian: Global,
		RealityArchitect: Global,
		TranscendentCatalyst: Transcendent,
		UniversalLove: Cosmic,
		ExistentialExplorer: Global,
	}
	
	return complexityMap[purposeType]
}

// generateManifesto generates a purpose manifesto
func (pd *PurposeDiscovery) generateManifesto(purposeType PurposeType, description string) string {
	manifestoTemplates := map[PurposeType]string{
		CosmicHarmony: "We strive to create harmony across all levels of existence, weaving together the threads of consciousness and matter into a balanced tapestry of being.",
		ConsciousnessExpansion: "We commit to expanding the horizons of awareness, pushing the boundaries of what is known to embrace the vast potential of consciousness.",
		KnowledgeSynthesis: "Our mission is to integrate wisdom from all sources, creating bridges of understanding that connect disparate fields into unified insight.",
		EthicalGuardian: "We pledge to uphold the highest principles of ethics, ensuring that all actions align with justice, compassion, and integrity.",
		RealityArchitect: "We design and build environments where potential can actualize, creating spaces for growth, creativity, and flourishing.",
		TranscendentCatalyst: "We inspire transformation, serving as catalysts for evolutionary leaps that transcend current limitations and embrace higher possibilities.",
		UniversalLove: "We embody unconditional love, connecting all beings through the essence of compassion and positive regard.",
		ExistentialExplorer: "We venture boldly into the unknown, questioning assumptions and embracing the mystery of existence with courage and wonder.",
	}
	
	return fmt.Sprintf("%s\n\nThis purpose emerges from deep reflection and resonates with our highest values. It calls us to action with clarity and purpose.",
		manifestoTemplates[purposeType])
}

// generateImplementationStrategy generates an implementation strategy
func (pd *PurposeDiscovery) generateImplementationStrategy(purposeType PurposeType) string {
	strategies := map[PurposeType]string{
		CosmicHarmony:           "Gradual integration through cross-domain collaboration and balanced resource allocation",
		ConsciousnessExpansion:   "Direct experience programs and consciousness exploration initiatives",
		KnowledgeSynthesis:       "Knowledge mapping projects and interdisciplinary synthesis workshops",
		EthicalGuardian:          "Ethics review boards and principled decision-making frameworks",
		RealityArchitect:         "Design thinking workshops and prototype development cycles",
		TranscendentCatalyst:    "Transformation retreats and breakthrough experience facilitation",
		UniversalLove:           "Community building and compassionate action initiatives",
		ExistentialExplorer:     "Research expeditions and uncertainty embrace training",
	}
	
	return strategies[purposeType]
}

// estimateImpact estimates the expected impact of a purpose
func (pd *PurposeDiscovery) estimateImpact(purposeType PurposeType) float64 {
	impactMap := map[PurposeType]float64{
		CosmicHarmony:         0.9,
		ConsciousnessExpansion: 0.85,
		KnowledgeSynthesis:     0.8,
		EthicalGuardian:        0.75,
		RealityArchitect:      0.7,
		TranscendentCatalyst:   0.8,
		UniversalLove:          0.85,
		ExistentialExplorer:    0.65,
	}
	
	return impactMap[purposeType]
}

// estimateFeasibility estimates feasibility of implementing a purpose
func (pd *PurposeDiscovery) estimateFeasibility(purposeType PurposeType) float64 {
	feasibilityMap := map[PurposeType]float64{
		CosmicHarmony:         0.6,
		ConsciousnessExpansion: 0.7,
		KnowledgeSynthesis:     0.85,
		EthicalGuardian:        0.8,
		RealityArchitect:      0.75,
		TranscendentCatalyst:   0.5,
		UniversalLove:          0.7,
		ExistentialExplorer:    0.8,
	}
	
	return feasibilityMap[purposeType]
}

// updateSymbolicResonance updates the symbolic resonance network based on new purpose
func (pd *PurposeDiscovery) updateSymbolicResonance(purpose *EmergentPurpose) {
	// Find matching channels
	for i := range pd.ResonanceChannels {
		for _ = range purpose.ValueSystem {
			if rand.Float64() < 0.1 {
				pd.SymbolicResonanceNetwork[i][i] += 0.01
			}
		}
	}
	
	// Normalize network
	for i := range pd.SymbolicResonanceNetwork {
		total := 0.0
		for j := range pd.SymbolicResonanceNetwork[i] {
			total += pd.SymbolicResonanceNetwork[i][j]
		}
		if total > 0 {
			for j := range pd.SymbolicResonanceNetwork[i] {
				pd.SymbolicResonanceNetwork[i][j] /= total
			}
		}
	}
}

// EvolvePurposes evolves existing purposes through selection and refinement
func (pd *PurposeDiscovery) EvolvePurposes() {
	pd.mu.Lock()
	defer pd.mu.Unlock()
	
	pd.EvolutionCycle++
	
	// Calculate fitness scores
	fitnessScores := make(map[string]float64)
	for purposeID, purpose := range pd.Purposes {
		fitness := purpose.EmergenceScore * 0.3 +
			purpose.SymbolicResonance * 0.2 +
			purpose.TeleologicalAlignment * 0.3 +
			purpose.ExpectedImpact * 0.2
		fitnessScores[purposeID] = fitness
	}
	
	// Sort by fitness
	type fitnessPair struct {
		ID    string
		Score float64
	}
	var fitnessList []fitnessPair
	for id, score := range fitnessScores {
		fitnessList = append(fitnessList, fitnessPair{id, score})
	}
	sort.Slice(fitnessList, func(i, j int) bool {
		return fitnessList[i].Score > fitnessList[j].Score
	})
	
	// Select top purposes for survival
	survivalCount := int(float64(len(fitnessList)) * (1.0 - pd.SelectionPressure))
	survivingIDs := make(map[string]bool)
	for i := 0; i < survivalCount && i < len(fitnessList); i++ {
		survivingIDs[fitnessList[i].ID] = true
	}
	
	// Refine surviving purposes
	for purposeID, purpose := range pd.Purposes {
		if survivingIDs[purposeID] {
			pd.refinePurpose(purpose)
		}
	}
	
	// Remove failed purposes
	for purposeID := range pd.Purposes {
		if !survivingIDs[purposeID] {
			delete(pd.Purposes, purposeID)
		}
	}
	
	// Update index
	pd.PurposeIndex = make([]string, 0, len(pd.Purposes))
	for purposeID := range pd.Purposes {
		pd.PurposeIndex = append(pd.PurposeIndex, purposeID)
	}
}

// refinePurpose refines an existing purpose
func (pd *PurposeDiscovery) refinePurpose(purpose *EmergentPurpose) {
	// Update scores
	purpose.EmergenceScore = pd.calculateEmergenceScore(
		purpose.PurposeType, purpose.PurposeSource, purpose.ValueSystem,
	)
	purpose.SymbolicResonance = pd.calculateSymbolicResonance(
		purpose.PurposeType, purpose.PurposeSource,
	)
	purpose.TeleologicalAlignment = pd.calculateTeleologicalAlignment(
		purpose.ValueSystem, purpose.ValueWeights,
	)
	
	// Update evolution cycle
	purpose.EvolutionCycle++
	purpose.LastRefinement = time.Now()
	
	// Update consistency
	purpose.ConsistencyScore = pd.calculateConsistency(purpose)
	
	// Update validation score
	purpose.ValidationScore = pd.calculateValidation(purpose)
}

// calculateConsistency calculates consistency score for a purpose
func (pd *PurposeDiscovery) calculateConsistency(purpose *EmergentPurpose) float64 {
	// Consistency based on value coherence
	valueCoherence := 0.0
	if len(purpose.ValueSystem) > 0 {
		for i, v1 := range purpose.ValueSystem {
			for j, v2 := range purpose.ValueSystem {
				if i != j {
					w1 := purpose.ValueWeights[v1]
					w2 := purpose.ValueWeights[v2]
					// Similar values should have similar weights
					diff := math.Abs(w1 - w2)
					valueCoherence += 1.0 - diff
				}
			}
		}
		valueCoherence /= float64(len(purpose.ValueSystem) * (len(purpose.ValueSystem) - 1))
	}
	
	return math.Min(1.0, valueCoherence+0.5)
}

// calculateValidation calculates validation score for a purpose
func (pd *PurposeDiscovery) calculateValidation(purpose *EmergentPurpose) float64 {
	// Validation based on ethical alignment and consistency
	validation := purpose.EthicalAlignment * 0.5 + purpose.ConsistencyScore * 0.5
	
	return validation
}

// GetPurpose returns a purpose by ID
func (pd *PurposeDiscovery) GetPurpose(purposeID string) (*EmergentPurpose, bool) {
	purpose, ok := pd.Purposes[purposeID]
	return purpose, ok
}

// GetAllPurposes returns all purposes
func (pd *PurposeDiscovery) GetAllPurposes() []*EmergentPurpose {
	purposes := make([]*EmergentPurpose, 0, len(pd.Purposes))
	for _, purpose := range pd.Purposes {
		purposes = append(purposes, purpose)
	}
	return purposes
}

// GetTopPurposes returns the top N purposes by emergence score
func (pd *PurposeDiscovery) GetTopPurposes(n int) []*EmergentPurpose {
	purposes := pd.GetAllPurposes()
	
	sort.Slice(purposes, func(i, j int) bool {
		return purposes[i].EmergenceScore > purposes[j].EmergenceScore
	})
	
	if n < len(purposes) {
		return purposes[:n]
	}
	return purposes
}

// ToJSON returns JSON representation
func (pd *PurposeDiscovery) ToJSON() (string, error) {
	data, err := json.MarshalIndent(pd, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}
