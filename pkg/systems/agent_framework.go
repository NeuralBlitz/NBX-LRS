package systems

import (
	"encoding/json"
	"errors"
	"math"
	"math/rand"
	"sync"
	"time"
)

// Agent Framework Constants
const (
	DefaultLearningRate = 0.01
	MinAutonomyScore = 0.0
	MaxAutonomyScore = 1.0
	MinConsciousness = 0.0
	MaxConsciousness = 1.0
	MinEthicalAlignment = 0.0
	MaxEthicalAlignment = 1.0

	MaxMemorySize = 10000
	MaxActionHistory = 1000

	ActionTimeout = 30 * time.Second
)

// Error Definitions
var (
	ErrAgentNotFound = errors.New("agent not found")
	ErrAgentNotInitialized = errors.New("agent not initialized")
	ErrInvalidLifecycle = errors.New("invalid lifecycle state")
	ErrCapabilityNotFound = errors.New("capability not found")
	ErrActionTimeout = errors.New("action timeout")
	ErrMemoryFull = errors.New("agent memory full")
)

// AgentLifecycle represents the lifecycle stages of an agent
type AgentLifecycle int

const (
	LifecycleInitialization AgentLifecycle = iota
	LifecyclePerception
	LifecycleCognition
	LifecycleAction
	LifecycleLearning
	LifecycleAdaptation
	LifecycleTermination
)

func (l AgentLifecycle) String() string {
	switch l {
	case LifecycleInitialization:
		return "initialization"
	case LifecyclePerception:
		return "perception"
	case LifecycleCognition:
		return "cognition"
	case LifecycleAction:
		return "action"
	case LifecycleLearning:
		return "learning"
	case LifecycleAdaptation:
		return "adaptation"
	case LifecycleTermination:
		return "termination"
	default:
		return "unknown"
	}
}

// CapabilityType represents different agent capabilities
type CapabilityType int

const (
	CapabilityPerception CapabilityType = iota
	CapabilityReasoning
	CapabilityPlanning
	CapabilityActionExecution
	CapabilityLearning
	CapabilityCommunication
	CapabilitySelfModification
	CapabilityEthicalReasoning
	CapabilityCreativity
	CapabilityMemory
	CapabilityAdaptation
	CapabilityMetacognition
)

func (c CapabilityType) String() string {
	switch c {
	case CapabilityPerception:
		return "perception"
	case CapabilityReasoning:
		return "reasoning"
	case CapabilityPlanning:
		return "planning"
	case CapabilityActionExecution:
		return "action_execution"
	case CapabilityLearning:
		return "learning"
	case CapabilityCommunication:
		return "communication"
	case CapabilitySelfModification:
		return "self_modification"
	case CapabilityEthicalReasoning:
		return "ethical_reasoning"
	case CapabilityCreativity:
		return "creativity"
	case CapabilityMemory:
		return "memory"
	case CapabilityAdaptation:
		return "adaptation"
	case CapabilityMetacognition:
		return "metacognition"
	default:
		return "unknown"
	}
}

// AgentState represents the current state of an autonomous agent
type AgentState struct {
	Lifecycle         AgentLifecycle `json:"lifecycle"`
	ConsciousnessLevel float64       `json:"consciousness_level"`
	AutonomyScore     float64       `json:"autonomy_score"`
	EthicalAlignment  float64       `json:"ethical_alignment"`
	ActiveCapabilities []CapabilityType `json:"active_capabilities"`
	PerformanceScore  float64       `json:"performance_score"`
	ExperiencePoints  int          `json:"experience_points"`
	Level            int          `json:"level"`
	HealthStatus    float64       `json:"health_status"`
	EnergyLevel     float64       `json:"energy_level"`
	LastActive      time.Time    `json:"last_active"`
}

// PerceptionResult represents the result of environmental perception
type PerceptionResult struct {
	Timestamp      time.Time   `json:"timestamp"`
	ObjectsDetected int        `json:"objects_detected"`
	Threats       []string    `json:"threats"`
	Opportunities  []string   `json:"opportunities"`
	Confidence     float64     `json:"confidence"`
	SensoryInput   map[string]float64 `json:"sensory_input"`
	RawData       []float64  `json:"raw_data"`
}

// ReasoningResult represents the result of cognitive reasoning
type ReasoningResult struct {
	Priority        string   `json:"priority"`
	Confidence      float64  `json:"confidence"`
	PlannedActions  []string `json:"planned_actions"`
	ReasoningPath   []string `json:"reasoning_path"`
	OptionsEvaluated int   `json:"options_evaluated"`
	RiskAssessment float64  `json:"risk_assessment"`
	TimeAllocated  float64  `json:"time_allocated_ms"`
}

// ActionResult represents the result of action execution
type ActionResult struct {
	Actions       []ActionDetail `json:"actions"`
	Timestamp     time.Time    `json:"timestamp"`
	TotalDuration float64     `json:"total_duration_ms"`
	SuccessRate   float64     `json:"success_rate"`
}

// ActionDetail represents details of a single action
type ActionDetail struct {
	Action   string  `json:"action"`
	Status   string  `json:"status"`
	Result   string  `json:"result"`
	Duration float64 `json:"duration_ms"`
	Output   string  `json:"output,omitempty"`
}

// LearningResult represents the result of learning from experience
type LearningResult struct {
	ExperienceID   string            `json:"experience_id"`
	MemoryUpdated  bool             `json:"memory_updated"`
	AutonomyGain   float64           `json:"autonomy_gain"`
	SkillsImproved []CapabilityType  `json:"skills_improved"`
	EfficiencyGain float64          `json:"efficiency_gain"`
	Insights       []string         `json:"insights"`
}

// AdaptationResult represents the result of self-adaptation
type AdaptationResult struct {
	CapabilitiesUnlocked []CapabilityType `json:"capabilities_unlocked"`
	StateImproved       bool            `json:"state_improved"`
	ConsciousnessGain  float64         `json:"consciousness_gain"`
	PerformanceDelta   float64         `json:"performance_delta"`
	AdaptationNotes    []string       `json:"adaptation_notes"`
}

// Experience represents a learning experience
type Experience struct {
	ExperienceID    string                 `json:"experience_id"`
	Cycle           int                    `json:"cycle"`
	Outcome         string                `json:"outcome"`
	Efficiency      float64               `json:"efficiency"`
	Timestamp       time.Time            `json:"timestamp"`
	ActionsTaken    []string             `json:"actions_taken"`
	Environment     map[string]interface{} `json:"environment"`
	Reward          float64              `json:"reward"`
}

// Environment represents the agent's operational environment
type Environment struct {
	Timestamp    time.Time              `json:"timestamp"`
	Objects      []string               `json:"objects"`
	Threats     []string               `json:"threats"`
	Opportunities []string              `json:"opportunities"`
	Resources    map[string]float64    `json:"resources"`
	Constraints  []string              `json:"constraints"`
	Dynamics    map[string]float64    `json:"dynamics"`
}

// AgentCapabilities manages the registry of agent capabilities
type AgentCapabilities struct {
	mu sync.RWMutex

	capabilities []CapabilityType
	enabled      map[CapabilityType]bool
	levels       map[CapabilityType]float64
}

// NewAgentCapabilities creates a new capabilities registry
func NewAgentCapabilities() *AgentCapabilities {
	return &AgentCapabilities{
		capabilities: []CapabilityType{
			CapabilityPerception,
			CapabilityReasoning,
			CapabilityPlanning,
			CapabilityActionExecution,
			CapabilityLearning,
			CapabilityCommunication,
			CapabilitySelfModification,
			CapabilityEthicalReasoning,
			CapabilityCreativity,
			CapabilityMemory,
			CapabilityAdaptation,
			CapabilityMetacognition,
		},
		enabled: make(map[CapabilityType]bool),
		levels:  make(map[CapabilityType]float64),
	}
}

// Initialize sets up default capabilities
func (a *AgentCapabilities) Initialize() {
	a.mu.Lock()
	defer a.mu.Unlock()

	for _, cap := range a.capabilities {
		a.enabled[cap] = false
		a.levels[cap] = 0.5
	}
}

// EnableCapability enables a specific capability
func (a *AgentCapabilities) EnableCapability(cap CapabilityType) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.enabled[cap] = true
	a.levels[cap] = math.Max(0.5, a.levels[cap])
	return nil
}

// DisableCapability disables a specific capability
func (a *AgentCapabilities) DisableCapability(cap CapabilityType) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.enabled[cap] = false
}

// IsEnabled checks if a capability is enabled
func (a *AgentCapabilities) IsEnabled(cap CapabilityType) bool {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return a.enabled[cap]
}

// GetLevel gets the proficiency level of a capability
func (a *AgentCapabilities) GetLevel(cap CapabilityType) float64 {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return a.levels[cap]
}

// SetLevel sets the proficiency level of a capability
func (a *AgentCapabilities) SetLevel(cap CapabilityType, level float64) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.levels[cap] = math.Max(0.0, math.Min(1.0, level))
}

// ListEnabled returns list of enabled capabilities
func (a *AgentCapabilities) ListEnabled() []CapabilityType {
	a.mu.RLock()
	defer a.mu.RUnlock()

	enabled := make([]CapabilityType, 0)
	for cap := range a.enabled {
		if a.enabled[cap] {
			enabled = append(enabled, cap)
		}
	}
	return enabled
}

// ListAll returns all capabilities
func (a *AgentCapabilities) ListAll() []CapabilityType {
	a.mu.RLock()
	defer a.mu.RUnlock()

	result := make([]CapabilityType, len(a.capabilities))
	copy(result, a.capabilities)
	return result
}

// AdvancedAutonomousAgent represents an advanced autonomous agent with consciousness and self-modification
type AdvancedAutonomousAgent struct {
	mu sync.RWMutex

	AgentID        string
	Capabilities   *AgentCapabilities
	State         *AgentState
	Memory        map[string]*Experience
	ActionHistory []ActionResult
	LearningRate  float64

	// Performance tracking
	TotalCycles     int
	TotalActions    int
	SuccessfulActions int
	TotalReward    float64

	// Timestamps
	CreatedAt     time.Time
	LastActionAt  time.Time
	LastLearnedAt time.Time

	// Configuration
	Config *AgentConfig
}

// AgentConfig represents agent configuration
type AgentConfig struct {
	AgentID           string
	InitialConsciousness float64
	InitialAutonomy    float64
	InitialEthics      float64
	LearningRate      float64
	MaxMemorySize     int
	ActionTimeout     time.Duration
}

// NewAdvancedAutonomousAgent creates a new advanced autonomous agent
func NewAdvancedAutonomousAgent(agentID string) *AdvancedAutonomousAgent {
	return &AdvancedAutonomousAgent{
		AgentID:      agentID,
		Capabilities: NewAgentCapabilities(),
		State: &AgentState{
			Lifecycle:          LifecycleInitialization,
			ConsciousnessLevel: 0.5,
			AutonomyScore:     0.0,
			EthicalAlignment:  1.0,
			ActiveCapabilities: make([]CapabilityType, 0),
			PerformanceScore: 0.5,
			ExperiencePoints: 0,
			Level:            1,
			HealthStatus:     1.0,
			EnergyLevel:     1.0,
			LastActive:      time.Now(),
		},
		Memory:        make(map[string]*Experience),
		ActionHistory: make([]ActionResult, 0, MaxActionHistory),
		LearningRate:  DefaultLearningRate,
		TotalCycles:   0,
		TotalActions:  0,
		SuccessfulActions: 0,
		TotalReward:  0.0,
		CreatedAt:    time.Now(),
		Config: &AgentConfig{
			AgentID:             agentID,
			InitialConsciousness: 0.5,
			InitialAutonomy:     0.0,
			InitialEthics:       1.0,
			LearningRate:       DefaultLearningRate,
			MaxMemorySize:      MaxMemorySize,
			ActionTimeout:       ActionTimeout,
		},
	}
}

// Initialize sets up the agent for operation
func (a *AdvancedAutonomousAgent) Initialize() {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.Capabilities.Initialize()
	a.State.Lifecycle = LifecyclePerception

	// Enable initial capabilities
	initialCaps := []CapabilityType{
		CapabilityPerception,
		CapabilityReasoning,
		CapabilityActionExecution,
	}
	for _, cap := range initialCaps {
		a.Capabilities.EnableCapability(cap)
		a.State.ActiveCapabilities = append(a.State.ActiveCapabilities, cap)
	}

	a.LastActionAt = time.Now()
}

// Perceive processes environmental data
func (a *AdvancedAutonomousAgent) Perceive(env *Environment) *PerceptionResult {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.State.Lifecycle = LifecyclePerception
	a.State.LastActive = time.Now()

	// Process environmental data
	objectsDetected := len(env.Objects)
	threats := env.Threats
	opportunities := env.Opportunities

	// Calculate confidence based on sensory data quality
	confidence := 0.5
	if len(env.Objects) > 0 {
		confidence = math.Min(1.0, 0.5+float64(len(env.Objects))*0.1)
	}

	// Build sensory input map
	sensoryInput := make(map[string]float64)
	for k, v := range env.Resources {
		sensoryInput["resource_"+k] = v
	}
	for k, v := range env.Dynamics {
		sensoryInput["dynamic_"+k] = v
	}

	result := &PerceptionResult{
		Timestamp:      time.Now(),
		ObjectsDetected: objectsDetected,
		Threats:       threats,
		Opportunities: opportunities,
		Confidence:   confidence,
		SensoryInput: sensoryInput,
		RawData:      make([]float64, 0),
	}

	return result
}

// Reason performs cognitive reasoning on perceptions
func (a *AdvancedAutonomousAgent) Reason(perceptions *PerceptionResult) *ReasoningResult {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.State.Lifecycle = LifecycleCognition

	// Determine priority based on threats
	priority := "explore"
	if len(perceptions.Threats) > 0 {
		priority = "avoid"
	}

	// Determine planned actions
	plannedActions := []string{"move", "scan"}
	if len(perceptions.Threats) > 0 {
		plannedActions = []string{"evade", "assess"}
	}

	// Calculate confidence
	confidence := perceptions.Confidence * a.Capabilities.GetLevel(CapabilityReasoning)

	// Build reasoning path
	reasoningPath := []string{
		"analyzed_environment",
		"evaluated_threats",
		"assessed_opportunities",
		"selected_priority=" + priority,
	}

	// Calculate risk assessment
	riskAssessment := 0.0
	if len(perceptions.Threats) > 0 {
		riskAssessment = math.Min(1.0, float64(len(perceptions.Threats))*0.3)
	}

	// Update consciousness level
	a.State.ConsciousnessLevel = math.Min(MaxConsciousness, a.State.ConsciousnessLevel+0.01)

	result := &ReasoningResult{
		Priority:       priority,
		Confidence:     confidence,
		PlannedActions: plannedActions,
		ReasoningPath:  reasoningPath,
		OptionsEvaluated: len(perceptions.Opportunities) + len(perceptions.Threats),
		RiskAssessment: riskAssessment,
		TimeAllocated: 100.0,
	}

	return result
}

// Act executes planned actions
func (a *AdvancedAutonomousAgent) Act(decisions *ReasoningResult) *ActionResult {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.State.Lifecycle = LifecycleAction

	startTime := time.Now()
	actions := make([]ActionDetail, 0)

	for _, action := range decisions.PlannedActions {
		actionStart := time.Now()

		// Simulate action execution
		status := "completed"
		result := "success"

		// Random success factor based on autonomy
		successChance := a.State.AutonomyScore
		if rand.Float64() > successChance {
			status = "partial"
			result = "degraded"
		}

		duration := time.Since(actionStart).Milliseconds()

		actions = append(actions, ActionDetail{
			Action:   action,
			Status:   status,
			Result:   result,
			Duration: float64(duration),
		})
	}

	// Calculate success rate
	successCount := 0
	for _, action := range actions {
		if action.Status == "completed" {
			successCount++
		}
	}
	successRate := float64(successCount) / float64(len(actions))

	result := &ActionResult{
		Actions:       actions,
		Timestamp:     time.Now(),
		TotalDuration: float64(time.Since(startTime).Milliseconds()),
		SuccessRate:   successRate,
	}

	// Update tracking
	a.TotalActions += len(actions)
	a.State.EnergyLevel = math.Max(0.0, a.State.EnergyLevel-0.01*float64(len(actions)))
	a.LastActionAt = time.Now()
	a.ActionHistory = append(a.ActionHistory, *result)
	if len(a.ActionHistory) > MaxActionHistory {
		a.ActionHistory = a.ActionHistory[len(a.ActionHistory)-MaxActionHistory:]
	}

	return result
}

// Learn updates agent memory and improves capabilities
func (a *AdvancedAutonomousAgent) Learn(experience *Experience) *LearningResult {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.State.Lifecycle = LifecycleLearning

	// Store experience
	experienceID := experience.ExperienceID
	if experienceID == "" {
		experienceID = generateExperienceID()
	}
	a.Memory[experienceID] = experience

	// Calculate autonomy gain
	autonomyGain := a.LearningRate
	if experience.Efficiency < 0.5 {
		autonomyGain *= 0.5 // Less gain from poor efficiency
	}
	a.State.AutonomyScore = math.Min(MaxAutonomyScore, a.State.AutonomyScore+autonomyGain)

	// Update experience points
	experienceReward := experience.Reward
	if experience.Outcome == "success" {
		experienceReward += 0.1
		a.SuccessfulActions++
	}
	a.TotalReward += experienceReward
	a.State.ExperiencePoints += int(experienceReward * 100)

	// Level up based on experience
	newLevel := 1 + a.State.ExperiencePoints/1000
	if newLevel > a.State.Level {
		a.State.Level = newLevel
	}

	// Determine skills improved
	skillsImproved := make([]CapabilityType, 0)
	if experience.Efficiency > 0.7 {
		skillsImproved = append(skillsImproved, CapabilityLearning)
		skillsImproved = append(skillsImproved, CapabilityMemory)
	}

	// Generate insights
	insights := []string{
		"experience_processed",
		"autonomy_updated",
	}

	result := &LearningResult{
		ExperienceID:   experienceID,
		MemoryUpdated:  true,
		AutonomyGain:   autonomyGain,
		SkillsImproved: skillsImproved,
		EfficiencyGain: experience.Efficiency - 0.5,
		Insights:       insights,
	}

	a.LastLearnedAt = time.Now()

	return result
}

// Adapt performs self-adaptation and capability unlocking
func (a *AdvancedAutonomousAgent) Adapt() *AdaptationResult {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.State.Lifecycle = LifecycleAdaptation

	capabilitiesUnlocked := make([]CapabilityType, 0)
	adaptationNotes := make([]string, 0)

	// Check if we can unlock new capabilities
	if a.State.AutonomyScore > 0.5 {
		availableCaps := a.Capabilities.ListAll()
		currentCaps := a.State.ActiveCapabilities

		for _, cap := range availableCaps {
			found := false
			for _, current := range currentCaps {
				if current == cap {
					found = true
					break
				}
			}
			if !found {
				// Unlock based on autonomy threshold
				unlockThreshold := 0.5 + float64(len(currentCaps))*0.1
				if a.State.AutonomyScore > unlockThreshold {
					a.Capabilities.EnableCapability(cap)
					a.State.ActiveCapabilities = append(a.State.ActiveCapabilities, cap)
					capabilitiesUnlocked = append(capabilitiesUnlocked, cap)
					adaptationNotes = append(adaptationNotes, "unlocked_"+cap.String())
				}
			}
		}
	}

	// Calculate consciousness gain
	consciousnessGain := a.LearningRate * 0.5
	a.State.ConsciousnessLevel = math.Min(MaxConsciousness, a.State.ConsciousnessLevel+consciousnessGain)

	// Calculate performance delta
	performanceDelta := (a.State.AutonomyScore - 0.5) * 0.1
	a.State.PerformanceScore = math.Max(0.0, math.Min(1.0, a.State.PerformanceScore+performanceDelta))

	// Regenerate energy
	a.State.EnergyLevel = math.Min(1.0, a.State.EnergyLevel+0.05)

	result := &AdaptationResult{
		CapabilitiesUnlocked: capabilitiesUnlocked,
		StateImproved:       len(capabilitiesUnlocked) > 0,
		ConsciousnessGain:  consciousnessGain,
		PerformanceDelta:   performanceDelta,
		AdaptationNotes:    adaptationNotes,
	}

	return result
}

// GetState returns the current agent state
func (a *AdvancedAutonomousAgent) GetState() *AgentState {
	a.mu.RLock()
	defer a.mu.RUnlock()

	state := *a.State
	return &state
}

// GetPerformanceMetrics returns agent performance metrics
func (a *AdvancedAutonomousAgent) GetPerformanceMetrics() map[string]float64 {
	a.mu.RLock()
	defer a.mu.RUnlock()

	successRate := 0.0
	if a.TotalActions > 0 {
		successRate = float64(a.SuccessfulActions) / float64(a.TotalActions)
	}

	return map[string]float64{
		"total_cycles":        float64(a.TotalCycles),
		"total_actions":       float64(a.TotalActions),
		"successful_actions":   float64(a.SuccessfulActions),
		"success_rate":        successRate,
		"total_reward":        a.TotalReward,
		"experience_points":    float64(a.State.ExperiencePoints),
		"level":              float64(a.State.Level),
		"autonomy_score":     a.State.AutonomyScore,
		"consciousness":      a.State.ConsciousnessLevel,
		"ethical_alignment":  a.State.EthicalAlignment,
		"performance_score":  a.State.PerformanceScore,
		"health_status":      a.State.HealthStatus,
		"energy_level":       a.State.EnergyLevel,
		"capability_count":   float64(len(a.State.ActiveCapabilities)),
		"memory_size":        float64(len(a.Memory)),
	}
}

// RunCycle executes a complete agent lifecycle cycle
func (a *AdvancedAutonomousAgent) RunCycle(env *Environment) (*CycleResult, error) {
	a.mu.Lock()
	a.TotalCycles++
	a.mu.Unlock()

	// Perception
	perceptions := a.Perceive(env)

	// Reasoning
	decisions := a.Reason(perceptions)

	// Action
	results := a.Act(decisions)

	// Learning
	experience := &Experience{
		ExperienceID:   generateExperienceID(),
		Cycle:         a.TotalCycles,
		Outcome:       "success",
		Efficiency:    results.SuccessRate,
		Timestamp:     time.Now(),
		ActionsTaken:  decisions.PlannedActions,
		Environment:   map[string]interface{}{
			"objects":      len(env.Objects),
			"threats":     len(env.Threats),
			"opportunities": len(env.Opportunities),
		},
		Reward: results.SuccessRate,
	}
	learning := a.Learn(experience)

	// Adaptation
	adaptation := a.Adapt()

	return &CycleResult{
		Cycle:       a.TotalCycles,
		Perception:  perceptions,
		Reasoning:   decisions,
		Action:     results,
		Learning:   learning,
		Adaptation: adaptation,
		Timestamp:  time.Now(),
	}, nil
}

// CycleResult represents the result of a complete agent cycle
type CycleResult struct {
	Cycle      int               `json:"cycle"`
	Perception *PerceptionResult  `json:"perception"`
	Reasoning  *ReasoningResult   `json:"reasoning"`
	Action    *ActionResult     `json:"action"`
	Learning  *LearningResult   `json:"learning"`
	Adaptation *AdaptationResult `json:"adaptation"`
	Timestamp  time.Time       `json:"timestamp"`
}

// ToJSON serializes the agent to JSON
func (a *AdvancedAutonomousAgent) ToJSON() ([]byte, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return json.MarshalIndent(map[string]interface{}{
		"agent_id":            a.AgentID,
		"state":              a.State,
		"active_capabilities":  a.State.ActiveCapabilities,
		"performance_metrics": a.GetPerformanceMetrics(),
		"memory_size":        len(a.Memory),
		"action_history_size": len(a.ActionHistory),
		"created_at":         a.CreatedAt,
		"last_action_at":     a.LastActionAt,
		"total_cycles":       a.TotalCycles,
	}, "", "  ")
}

// String returns string representation of the agent
func (a *AdvancedAutonomousAgent) String() string {
	a.mu.RLock()
	defer a.mu.RUnlock()

	jsonData, _ := a.ToJSON()
	return string(jsonData)
}

// AgentRegistry manages global agent instances
type AgentRegistry struct {
	mu sync.RWMutex

	agents map[string]*AdvancedAutonomousAgent
}

// NewAgentRegistry creates a new agent registry
func NewAgentRegistry() *AgentRegistry {
	return &AgentRegistry{
		agents: make(map[string]*AdvancedAutonomousAgent),
	}
}

// Register adds an agent to the registry
func (r *AgentRegistry) Register(agent *AdvancedAutonomousAgent) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.agents[agent.AgentID] = agent
}

// Get retrieves an agent by ID
func (r *AgentRegistry) Get(agentID string) (*AdvancedAutonomousAgent, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	agent, exists := r.agents[agentID]
	if !exists {
		return nil, ErrAgentNotFound
	}
	return agent, nil
}

// List returns all registered agent IDs
func (r *AgentRegistry) List() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	ids := make([]string, 0, len(r.agents))
	for id := range r.agents {
		ids = append(ids, id)
	}
	return ids
}

// Count returns the number of registered agents
func (r *AgentRegistry) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return len(r.agents)
}

// generateExperienceID generates a unique experience ID
func generateExperienceID() string {
	timestamp := time.Now().UnixNano() / 1e6
	randomPart := rand.Intn(10000)
	return "exp_" + string(rune('a'+timestamp%26)) +
		string(rune('0'+randomPart/1000)) +
		string(rune('0'+randomPart%1000/100)) +
		string(rune('0'+randomPart%100/10)) +
		string(rune('0'+randomPart%10))
}
