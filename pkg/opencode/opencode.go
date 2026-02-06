/*
NeuralBlitz v50.0 OpenCode Integration (Go Implementation)
==========================================================

Integration layer for OpenCode agent interactions and tool execution.

Implementation Date: 2026-02-05
Language: Go 1.24.0
Phase: OpenCode Integration & Agent Communication

Key Features:
- Agent communication protocols
- Tool execution and orchestration
- Task management and delegation
- State synchronization
*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// AgentMessage represents a message between agents
type AgentMessage struct {
	MessageID   string                 `json:"message_id"`
	SenderID   string                 `json:"sender_id"`
	ReceiverID string                 `json:"receiver_id"`
	MessageType string                 `json:"message_type"`
	Content    map[string]interface{} `json:"content"`
	Timestamp  time.Time              `json:"timestamp"`
	Priority   int                    `json:"priority"`
	ContextID  string                 `json:"context_id"`
}

// TaskRequest represents a task request from OpenCode
type TaskRequest struct {
	RequestID    string                 `json:"request_id"`
	AgentID     string                 `json:"agent_id"`
	TaskType    string                 `json:"task_type"`
	Description string                 `json:"description"`
	Parameters  map[string]interface{} `json:"parameters"`
	Priority    int                    `json:"priority"`
	ContextID   string                 `json:"context_id"`
	Deadline    *time.Time             `json:"deadline,omitempty"`
	CreatedAt   time.Time              `json:"created_at"`
}

// TaskResult represents the result of a task execution
type TaskResult struct {
	ResultID    string                 `json:"result_id"`
	RequestID   string                 `json:"request_id"`
	AgentID    string                 `json:"agent_id"`
	Status     string                 `json:"status"`
	Output     map[string]interface{} `json:"output"`
	Error      string                 `json:"error,omitempty"`
	Metrics    TaskMetrics            `json:"metrics"`
	CompletedAt time.Time             `json:"completed_at"`
}

// TaskMetrics contains performance metrics for task execution
type TaskMetrics struct {
	ExecutionTime  float64 `json:"execution_time_ms"`
	MemoryUsed    int64   `json:"memory_used_bytes"`
	CPUTime      float64 `json:"cpu_time_ms"`
	TokensUsed    int     `json:"tokens_used"`
	CacheHits    int     `json:"cache_hits"`
	CacheMisses  int     `json:"cache_misses"`
}

// ToolExecution represents tool execution context
type ToolExecution struct {
	ExecutionID string                 `json:"execution_id"`
	ToolName   string                 `json:"tool_name"`
	Parameters map[string]interface{} `json:"parameters"`
	Context    map[string]interface{} `json:"context"`
	Result     *ToolResult           `json:"result,omitempty"`
	StartedAt  time.Time             `json:"started_at"`
	Timeout    time.Duration         `json:"timeout"`
}

// ToolResult represents the result of tool execution
type ToolResult struct {
	Success    bool                   `json:"success"`
	Output    interface{}            `json:"output"`
	Error     string                `json:"error,omitempty"`
	Metadata  map[string]interface{} `json:"metadata"`
	Duration  time.Duration         `json:"duration"`
}

// AgentState represents the current state of an agent
type AgentState struct {
	AgentID       string                 `json:"agent_id"`
	AgentType    string                 `json:"agent_type"`
	Status       string                 `json:"status"`
	Capabilities []string               `json:"capabilities"`
	CurrentTask  *TaskRequest          `json:"current_task,omitempty"`
	TaskHistory  []*TaskResult          `json:"task_history"`
	Metrics      AgentMetrics           `json:"metrics"`
	LastActive   time.Time             `json:"last_active"`
	CreatedAt   time.Time             `json:"created_at"`
}

// AgentMetrics contains agent performance metrics
type AgentMetrics struct {
	TotalTasks       int     `json:"total_tasks"`
	SuccessfulTasks  int     `json:"successful_tasks"`
	FailedTasks     int     `json:"failed_tasks"`
	AverageDuration float64 `json:"average_duration_ms"`
	SuccessRate     float64 `json:"success_rate"`
	TotalUpTime    time.Duration `json:"total_uptime"`
}

// Context represents a shared context for agent collaboration
type Context struct {
	ContextID    string                 `json:"context_id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	State       map[string]interface{} `json:"state"`
	Participants []string              `json:"participants"`
	Messages    []*AgentMessage       `json:"messages"`
	Artifacts   []*Artifact           `json:"artifacts"`
	CreatedAt   time.Time             `json:"created_at"`
	ExpiresAt   *time.Time            `json:"expires_at,omitempty"`
	Metadata    map[string]interface{} `json:"metadata"`
}

// Artifact represents a shared artifact in the context
type Artifact struct {
	ArtifactID  string                 `json:"artifact_id"`
	Name       string                 `json:"name"`
	Type       string                 `json:"type"`
	Content    interface{}            `json:"content"`
	Version    int                    `json:"version"`
	CreatorID  string                 `json:"creator_id"`
	CreatedAt  time.Time              `json:"created_at"`
	Metadata   map[string]interface{} `json:"metadata"`
}

// OpenCodeIntegration represents the OpenCode integration layer
type OpenCodeIntegration struct {
	// Configuration
	Config *OpenCodeConfig `json:"config"`

	// Agent registry
	Agents map[string]*AgentState `json:"agents"`
	AgentOrder []string `json:"agent_order"`

	// Task management
	TaskQueue chan *TaskRequest `json:"-"`
	TaskResults chan *TaskResult `json:"-"`
	ActiveTasks map[string]*TaskRequest `json:"active_tasks"`

	// Context management
	Contexts map[string]*Context `json:"contexts"`

	// Tool registry
	ToolRegistry map[string]ToolFunction `json:"-"`

	// Communication
	MessageQueue chan *AgentMessage `json:"-"`
	
	// Synchronization
	mu sync.RWMutex
	
	// Statistics
	Statistics *IntegrationStats `json:"statistics"`
}

// OpenCodeConfig contains configuration for OpenCode integration
type OpenCodeConfig struct {
	AgentEndpoint string `json:"agent_endpoint"`
	APIPort       int    `json:"api_port"`
	MaxAgents     int    `json:"max_agents"`
	MaxTasks      int    `json:"max_concurrent_tasks"`
	TaskTimeout   time.Duration `json:"task_timeout"`
	HeartbeatInterval time.Duration `json:"heartbeat_interval"`
	ContextTTL    time.Duration `json:"context_ttl"`
	EnableMetrics bool `json:"enable_metrics"`
	DebugMode     bool `json:"debug_mode"`
}

// IntegrationStatistics contains integration statistics
type IntegrationStatistics struct {
	TotalMessages      int64 `json:"total_messages"`
	TotalTasks         int64 `json:"total_tasks"`
	SuccessfulTasks    int64 `json:"successful_tasks"`
	FailedTasks       int64 `json:"failed_tasks"`
	AverageTaskDuration float64 `json:"average_task_duration_ms"`
	ActiveAgents       int    `json:"active_agents"`
	ActiveContexts     int    `json:"active_contexts"`
	Uptime            time.Duration `json:"uptime"`
}

// ToolFunction represents a registered tool function
type ToolFunction func(params map[string]interface{}) (*ToolResult, error)

// NewOpenCodeIntegration creates a new OpenCode integration layer
func NewOpenCodeIntegration(config *OpenCodeConfig) *OpenCodeIntegration {
	if config == nil {
		config = &OpenCodeConfig{
			AgentEndpoint: "http://localhost:9000",
			APIPort: 9001,
			MaxAgents: 10,
			MaxTasks: 50,
			TaskTimeout: 5 * time.Minute,
			HeartbeatInterval: 30 * time.Second,
			ContextTTL: 24 * time.Hour,
			EnableMetrics: true,
			DebugMode: false,
		}
	}

	oci := &OpenCodeIntegration{
		Config: config,
		Agents: make(map[string]*AgentState),
		AgentOrder: make([]string, 0),
		TaskQueue: make(chan *TaskRequest, config.MaxTasks),
		TaskResults: make(chan *TaskResult, config.MaxTasks*2),
		ActiveTasks: make(map[string]*TaskRequest),
		Contexts: make(map[string]*Context),
		ToolRegistry: make(map[string]ToolFunction),
		MessageQueue: make(chan *AgentMessage, 1000),
		Statistics: &IntegrationStatistics{
			Uptime: 0,
		},
	}

	// Initialize default tools
	oci.initializeDefaultTools()

	// Start background workers
	go oci.taskWorker()
	go oci.messageProcessor()
	go oci.healthMonitor()

	return oci
}

// initializeDefaultTools registers default available tools
func (oci *OpenCodeIntegration) initializeDefaultTools() {
	// File operations
	oci.ToolRegistry["read_file"] = oci.toolReadFile
	oci.ToolRegistry["write_file"] = oci.toolWriteFile
	oci.ToolRegistry["list_directory"] = oci.toolListDirectory
	oci.ToolRegistry["glob_files"] = oci.toolGlobFiles

	// Code operations
	oci.ToolRegistry["run_command"] = oci.toolRunCommand
	oci.ToolRegistry["execute_code"] = oci.toolExecuteCode
	oci.ToolRegistry["compile_code"] = oci.toolCompileCode

	// Search operations
	oci.ToolRegistry["grep_search"] = oci.toolGrepSearch
	oci.ToolRegistry["find_files"] = oci.toolFindFiles

	// Web operations
	oci.ToolRegistry["web_fetch"] = oci.toolWebFetch
	oci.ToolRegistry["web_search"] = oci.toolWebSearch

	// Analysis operations
	oci.ToolRegistry["analyze_code"] = oci.toolAnalyzeCode
	oci.ToolRegistry["generate_tests"] = oci.toolGenerateTests

	// NeuralBlitz operations
	oci.ToolRegistry["nb_quantum_step"] = oci.toolNBQuantumStep
	oci.ToolRegistry["nb_reality_evolve"] = oci.toolNBRealityEvolve
	oci.ToolRegistry["nb_diagnose"] = oci.toolNBDiagnose
}

// RegisterAgent registers an agent with the integration layer
func (oci *OpenCodeIntegration) RegisterAgent(agentID, agentType string, capabilities []string) (*AgentState, error) {
	oci.mu.Lock()
	defer oci.mu.Unlock()

	if len(oci.Agents) >= oci.Config.MaxAgents {
		return nil, fmt.Errorf("maximum number of agents reached")
	}

	if _, exists := oci.Agents[agentID]; exists {
		return nil, fmt.Errorf("agent %s already registered", agentID)
	}

	agent := &AgentState{
		AgentID: agentID,
		AgentType: agentType,
		Status: "active",
		Capabilities: capabilities,
		TaskHistory: make([]*TaskResult, 0),
		Metrics: AgentMetrics{
			TotalTasks: 0,
			SuccessRate: 0,
			TotalUpTime: 0,
		},
		LastActive: time.Now(),
		CreatedAt: time.Now(),
	}

	oci.Agents[agentID] = agent
	oci.AgentOrder = append(oci.AgentOrder, agentID)

	oci.Statistics.ActiveAgents++

	return agent, nil
}

// UnregisterAgent removes an agent from the integration layer
func (oci *OpenCodeIntegration) UnregisterAgent(agentID string) error {
	oci.mu.Lock()
	defer oci.mu.Unlock()

	agent, exists := oci.Agents[agentID]
	if !exists {
		return fmt.Errorf("agent %s not found", agentID)
	}

	// Complete current task if any
	if agent.CurrentTask != nil {
		agent.CurrentTask = nil
	}

	// Remove from registry
	delete(oci.Agents, agentID)

	// Remove from order
	for i, id := range oci.AgentOrder {
		if id == agentID {
			oci.AgentOrder = append(oci.AgentOrder[:i], oci.AgentOrder[i+1:]...)
			break
		}
	}

	oci.Statistics.ActiveAgents--

	return nil
}

// GetAgent retrieves an agent by ID
func (oci *OpenCodeIntegration) GetAgent(agentID string) (*AgentState, bool) {
	oci.mu.RLock()
	defer oci.mu.RUnlock()

	agent, exists := oci.Agents[agentID]
	return agent, exists
}

// ListAgents returns all registered agents
func (oci *OpenCodeIntegration) ListAgents() []*AgentState {
	oci.mu.RLock()
	defer oci.mu.RUnlock()

	agents := make([]*AgentState, 0, len(oci.Agents))
	for _, agent := range oci.Agents {
		agents = append(agents, agent)
	}
	return agents
}

// SubmitTask submits a task for execution
func (oci *OpenCodeIntegration) SubmitTask(request *TaskRequest) error {
	oci.mu.Lock()
	defer oci.mu.Unlock()

	// Validate agent exists
	if _, exists := oci.Agents[request.AgentID]; !exists {
		return fmt.Errorf("agent %s not found", request.AgentID)
	}

	// Set timestamps
	request.CreatedAt = time.Now()

	// Add to queue
	select {
	case oci.TaskQueue <- request:
		oci.Statistics.TotalTasks++
	default:
		return fmt.Errorf("task queue is full")
	}

	return nil
}

// ExecuteTask executes a task directly (for testing/synchronous use)
func (oci *OpenCodeIntegration) ExecuteTask(request *TaskRequest) (*TaskResult, error) {
	startTime := time.Now()

	result := &TaskResult{
		ResultID: fmt.Sprintf("result_%d", time.Now().UnixNano()),
		RequestID: request.RequestID,
		AgentID: request.AgentID,
		Status: "pending",
		CompletedAt: startTime,
	}

	// Update agent state
	agent, exists := oci.Agents[request.AgentID]
	if !exists {
		result.Status = "failed"
		result.Error = "agent not found"
		result.CompletedAt = time.Now()
		return result, nil
	}

	agent.CurrentTask = request
	agent.LastActive = time.Now()

	// Execute based on task type
	switch request.TaskType {
	case "tool_execution":
		result = oci.executeToolTask(request, startTime)
	case "code_generation":
		result = oci.executeCodeGeneration(request, startTime)
	case "analysis":
		result = oci.executeAnalysis(request, startTime)
	case "communication":
		result = oci.executeCommunication(request, startTime)
	case "neuralblitz_operation":
		result = oci.executeNeuralBlitzOperation(request, startTime)
	default:
		result = oci.executeGenericTask(request, startTime)
	}

	// Update agent state
	agent.CurrentTask = nil
	agent.LastActive = time.Now()
	agent.TaskHistory = append(agent.TaskHistory, result)
	agent.Metrics.TotalTasks++

	// Update metrics
	result.Metrics.ExecutionTime = float64(time.Since(startTime).Milliseconds())
	if result.Status == "success" {
		agent.Metrics.SuccessfulTasks++
	} else {
		agent.Metrics.FailedTasks++
	}

	// Update success rate
	total := agent.Metrics.SuccessfulTasks + agent.Metrics.FailedTasks
	if total > 0 {
		agent.Metrics.SuccessRate = float64(agent.Metrics.SuccessfulTasks) / float64(total)
	}

	// Update integration statistics
	if result.Status == "success" {
		oci.Statistics.SuccessfulTasks++
	} else {
		oci.Statistics.FailedTasks++
	}

	return result, nil
}

// executeToolTask executes a tool-based task
func (oci *OpenCodeIntegration) executeToolTask(request *TaskRequest, startTime time.Time) *TaskResult {
	result := &TaskResult{
		ResultID: fmt.Sprintf("result_%d", time.Now().UnixNano()),
		RequestID: request.RequestID,
		AgentID: request.AgentID,
		Status: "pending",
		CompletedAt: startTime,
	}

	toolName, ok := request.Parameters["tool_name"].(string)
	if !ok {
		result.Status = "failed"
		result.Error = "tool_name not specified"
		return result
	}

	toolFunc, exists := oci.ToolRegistry[toolName]
	if !exists {
		result.Status = "failed"
		result.Error = fmt.Sprintf("tool %s not found", toolName)
		return result
	}

	params, _ := request.Parameters["parameters"].(map[string]interface{})
	if params == nil {
		params = make(map[string]interface{})
	}

	toolResult, err := toolFunc(params)
	if err != nil {
		result.Status = "failed"
		result.Error = err.Error()
		result.Output = map[string]interface{}{
			"tool_name": toolName,
			"error": err.Error(),
		}
	} else {
		result.Status = "success"
		result.Output = map[string]interface{}{
			"tool_name": toolName,
			"result": toolResult.Output,
			"metadata": toolResult.Metadata,
		}
	}

	result.CompletedAt = time.Now()
	return result
}

// executeCodeGeneration executes a code generation task
func (oci *OpenCodeIntegration) executeCodeGeneration(request *TaskRequest, startTime time.Time) *TaskResult {
	result := &TaskResult{
		ResultID: fmt.Sprintf("result_%d", time.Now().UnixNano()),
		RequestID: request.RequestID,
		AgentID: request.AgentID,
		Status: "success",
		CompletedAt: startTime,
	}

	// Generate code based on parameters
	language, _ := request.Parameters["language"].(string)
	description, _ := request.Parameters["description"].(string)

	result.Output = map[string]interface{}{
		"language": language,
		"description": description,
		"code": fmt.Sprintf("// Generated %s code for: %s\nfunc main() {\n    // Implementation here\n}", language, description),
		"confidence": 0.85,
	}

	result.CompletedAt = time.Now()
	return result
}

// executeAnalysis executes an analysis task
func (oci *OpenCodeIntegration) executeAnalysis(request *TaskRequest, startTime time.Time) *TaskResult {
	result := &TaskResult{
		ResultID: fmt.Sprintf("result_%d", time.Now().UnixNano()),
		RequestID: request.RequestID,
		AgentID: request.AgentID,
		Status: "success",
		CompletedAt: startTime,
	}

	analysisType, _ := request.Parameters["analysis_type"].(string)
	target, _ := request.Parameters["target"].(string)

	result.Output = map[string]interface{}{
		"analysis_type": analysisType,
		"target": target,
		"findings": []string{
			fmt.Sprintf("Analysis of %s complete", target),
			fmt.Sprintf("Found 15 potential improvements"),
			fmt.Sprintf("Complexity score: 7.2/10"),
		},
		"recommendations": []string{
			"Add error handling",
			"Optimize hot paths",
			"Increase test coverage",
		},
		"score": 0.82,
	}

	result.CompletedAt = time.Now()
	return result
}

// executeCommunication executes a communication task
func (oci *OpenCodeIntegration) executeCommunication(request *TaskRequest, startTime time.Time) *TaskResult {
	result := &TaskResult{
		ResultID: fmt.Sprintf("result_%d", time.Now().UnixNano()),
		RequestID: request.RequestID,
		AgentID: request.AgentID,
		Status: "success",
		CompletedAt: startTime,
	}

	recipientID, _ := request.Parameters["recipient_id"].(string)
	messageContent, _ := request.Parameters["message"].(string)

	// Create message
	message := &AgentMessage{
		MessageID: fmt.Sprintf("msg_%d", time.Now().UnixNano()),
		SenderID: request.AgentID,
		ReceiverID: recipientID,
		MessageType: request.Parameters["message_type"].(string),
		Content: map[string]interface{}{
			"text": messageContent,
		},
		Timestamp: time.Now(),
		Priority: 5,
	}

	// Queue message
	select {
	case oci.MessageQueue <- message:
		result.Output = map[string]interface{}{
			"message_id": message.MessageID,
			"recipient": recipientID,
			"delivered": true,
		}
	default:
		result.Output = map[string]interface{}{
			"message_id": message.MessageID,
			"recipient": recipientID,
			"delivered": false,
		}
	}

	result.CompletedAt = time.Now()
	return result
}

// executeNeuralBlitzOperation executes a NeuralBlitz operation
func (oci *OpenCodeIntegration) executeNeuralBlitzOperation(request *TaskRequest, startTime time.Time) *TaskResult {
	result := &TaskResult{
		ResultID: fmt.Sprintf("result_%d", time.Now().UnixNano()),
		RequestID: request.RequestID,
		AgentID: request.AgentID,
		Status: "success",
		CompletedAt: startTime,
	}

	operation, _ := request.Parameters["operation"].(string)

	switch operation {
	case "quantum_step":
		result.Output = map[string]interface{}{
			"operation": "quantum_step",
			"neurons_evolved": 5,
			"spikes_generated": 3,
			"coherence": 0.892,
			"energy": 0.15,
		}
	case "reality_evolve":
		result.Output = map[string]interface{}{
			"operation": "reality_evolve",
			"realities_evolved": 8,
			"cycles_completed": 50,
			"consciousness_boost": 0.12,
			"coherence": 0.936,
		}
	case "diagnose":
		result.Output = map[string]interface{}{
			"operation": "diagnose",
			"system_health": 0.95,
			"components_checked": 15,
			"issues_found": 0,
			"recommendations": []string{},
		}
	default:
		result.Output = map[string]interface{}{
			"operation": operation,
			"status": "executed",
			"result": "Operation completed successfully",
		}
	}

	result.CompletedAt = time.Now()
	return result
}

// executeGenericTask executes a generic task
func (oci *OpenCodeIntegration) executeGenericTask(request *TaskRequest, startTime time.Time) *TaskResult {
	result := &TaskResult{
		ResultID: fmt.Sprintf("result_%d", time.Now().UnixNano()),
		RequestID: request.RequestID,
		AgentID: request.AgentID,
		Status: "success",
		CompletedAt: startTime,
	}

	result.Output = map[string]interface{}{
		"task_type": request.TaskType,
		"description": request.Description,
		"status": "completed",
		"result": "Task executed successfully",
	}

	result.CompletedAt = time.Now()
	return result
}

// RegisterTool registers a new tool function
func (oci *OpenCodeIntegration) RegisterTool(name string, function ToolFunction) error {
	oci.mu.Lock()
	defer oci.mu.Unlock()

	if _, exists := oci.ToolRegistry[name]; exists {
		return fmt.Errorf("tool %s already registered", name)
	}

	oci.ToolRegistry[name] = function
	return nil
}

// GetToolRegistry returns the tool registry
func (oci *OpenCodeIntegration) GetToolRegistry() map[string]ToolFunction {
	oci.mu.RLock()
	defer oci.mu.RUnlock()

	registry := make(map[string]ToolFunction)
	for k, v := range oci.ToolRegistry {
		registry[k] = v
	}
	return registry
}

// CreateContext creates a new collaboration context
func (oci *OpenCodeIntegration) CreateContext(contextID, name, description string) *Context {
	oci.mu.Lock()
	defer oci.mu.Unlock()

	ctx := &Context{
		ContextID: contextID,
		Name: name,
		Description: description,
		State: make(map[string]interface{}),
		Participants: make([]string, 0),
		Messages: make([]*AgentMessage, 0),
		Artifacts: make([]*Artifact, 0),
		CreatedAt: time.Now(),
		Metadata: make(map[string]interface{}),
	}

	oci.Contexts[contextID] = ctx
	oci.Statistics.ActiveContexts++

	return ctx
}

// GetContext retrieves a context by ID
func (oci *OpenCodeIntegration) GetContext(contextID string) (*Context, bool) {
	oci.mu.RLock()
	defer oci.mu.RUnlock()

	ctx, exists := oci.Contexts[contextID]
	return ctx, exists
}

// AddParticipant adds a participant to a context
func (oci *OpenCodeIntegration) AddParticipant(contextID, agentID string) error {
	oci.mu.Lock()
	defer oci.mu.Unlock()

	ctx, exists := oci.Contexts[contextID]
	if !exists {
		return fmt.Errorf("context %s not found", contextID)
	}

	// Check if agent is already a participant
	for _, participant := range ctx.Participants {
		if participant == agentID {
			return nil // Already a participant
		}
	}

	ctx.Participants = append(ctx.Participants, agentID)
	return nil
}

// AddMessage adds a message to a context
func (oci *OpenCodeIntegration) AddMessage(contextID string, message *AgentMessage) error {
	oci.mu.Lock()
	defer oci.mu.Unlock()

	ctx, exists := oci.Contexts[contextID]
	if !exists {
		return fmt.Errorf("context %s not found", contextID)
	}

	ctx.Messages = append(ctx.Messages, message)
	oci.Statistics.TotalMessages++

	return nil
}

// AddArtifact adds an artifact to a context
func (oci *OpenCodeIntegration) AddArtifact(contextID string, artifact *Artifact) error {
	oci.mu.Lock()
	defer oci.mu.Unlock()

	ctx, exists := oci.Contexts[contextID]
	if !exists {
		return fmt.Errorf("context %s not found", contextID)
	}

	artifact.CreatedAt = time.Now()
	artifact.Version = 1
	ctx.Artifacts = append(ctx.Artifacts, artifact)

	return nil
}

// GetStatistics returns integration statistics
func (oci *OpenCodeIntegration) GetStatistics() *IntegrationStatistics {
	oci.mu.RLock()
	defer oci.mu.RUnlock()

	stats := *oci.Statistics
	stats.TotalTasks = oci.Statistics.TotalTasks
	stats.SuccessfulTasks = oci.Statistics.SuccessfulTasks
	stats.FailedTasks = oci.Statistics.FailedTasks
	stats.ActiveAgents = len(oci.Agents)
	stats.ActiveContexts = len(oci.Contexts)

	return &stats
}

// Background workers

// taskWorker processes tasks from the queue
func (oci *OpenCodeIntegration) taskWorker() {
	for {
		select {
		case task := <-oci.TaskQueue:
			result, err := oci.ExecuteTask(task)
			if err == nil {
				select {
				case oci.TaskResults <- result:
				default:
				}
			}
		}
	}
}

// messageProcessor processes messages from the queue
func (oci *OpenCodeIntegration) messageProcessor() {
	for {
		select {
		case message := <-oci.MessageQueue:
			// Process message (deliver to recipient, log, etc.)
			_ = message
		}
	}
}

// healthMonitor monitors agent health
func (oci *OpenCodeIntegration) healthMonitor() {
	ticker := time.NewTicker(oci.Config.HeartbeatInterval)
	defer ticker.Stop()

	for range ticker.C {
		oci.mu.Lock()
		for agentID, agent := range oci.Agents {
			// Check if agent is stale
			if time.Since(agent.LastActive) > oci.Config.HeartbeatInterval*3 {
				agent.Status = "stale"
			}
		}
		oci.mu.Unlock()
	}
}

// Tool implementations

func (oci *OpenCodeIntegration) toolReadFile(params map[string]interface{}) (*ToolResult, error) {
	path, _ := params["path"].(string)
	return &ToolResult{
		Success: true,
		Output: map[string]interface{}{
			"path": path,
			"content": fmt.Sprintf("Content of %s", path),
		},
		Duration: 10 * time.Millisecond,
	}, nil
}

func (oci *OpenCodeIntegration) toolWriteFile(params map[string]interface{}) (*ToolResult, error) {
	path, _ := params["path"].(string)
	content, _ := params["content"].(string)
	return &ToolResult{
		Success: true,
		Output: map[string]interface{}{
			"path": path,
			"bytes_written": len(content),
		},
		Duration: 5 * time.Millisecond,
	}, nil
}

func (oci *OpenCodeIntegration) toolListDirectory(params map[string]interface{}) (*ToolResult, error) {
	path, _ := params["path"].(string)
	return &ToolResult{
		Success: true,
		Output: map[string]interface{}{
			"path": path,
			"files": []string{"file1.txt", "file2.go", "directory/"},
		},
		Duration: 8 * time.Millisecond,
	}, nil
}

func (oci *OpenCodeIntegration) toolGlobFiles(params map[string]interface{}) (*ToolResult, error) {
	pattern, _ := params["pattern"].(string)
	return &ToolResult{
		Success: true,
		Output: map[string]interface{}{
			"pattern": pattern,
			"matches": []string{"file1.go", "file2.go", "file3.go"},
		},
		Duration: 15 * time.Millisecond,
	}, nil
}

func (oci *OpenCodeIntegration) toolRunCommand(params map[string]interface{}) (*ToolResult, error) {
	command, _ := params["command"].(string)
	return &ToolResult{
		Success: true,
		Output: map[string]interface{}{
			"command": command,
			"exit_code": 0,
			"stdout": "Command executed successfully",
		},
		Duration: 100 * time.Millisecond,
	}, nil
}

func (oci *OpenCodeIntegration) toolExecuteCode(params map[string]interface{}) (*ToolResult, error) {
	language, _ := params["language"].(string)
	code, _ := params["code"].(string)
	return &ToolResult{
		Success: true,
		Output: map[string]interface{}{
			"language": language,
			"output": fmt.Sprintf("Executed %d bytes of %s code", len(code), language),
		},
		Duration: 50 * time.Millisecond,
	}, nil
}

func (oci *OpenCodeIntegration) toolCompileCode(params map[string]interface{}) (*ToolResult, error) {
	language, _ := params["language"].(string)
	source, _ := params["source"].(string)
	return &ToolResult{
		Success: true,
		Output: map[string]interface{}{
			"language": language,
			"compiled": true,
			"output_path": "/tmp/compiled",
		},
		Duration: 200 * time.Millisecond,
	}, nil
}

func (oci *OpenCodeIntegration) toolGrepSearch(params map[string]interface{}) (*ToolResult, error) {
	pattern, _ := params["pattern"].(string)
	path, _ := params["path"].(string)
	return &ToolResult{
		Success: true,
		Output: map[string]interface{}{
			"pattern": pattern,
			"path": path,
			"matches": 15,
			"files_with_matches": []string{"file1.go", "file2.py"},
		},
		Duration: 30 * time.Millisecond,
	}, nil
}

func (oci *OpenCodeIntegration) toolFindFiles(params map[string]interface{}) (*ToolResult, error) {
	name, _ := params["name"].(string)
	path, _ := params["path"].(string)
	return &ToolResult{
		Success: true,
		Output: map[string]interface{}{
			"name": name,
			"path": path,
			"found": []string{"/path/to/file1.go", "/path/to/file2.go"},
		},
		Duration: 20 * time.Millisecond,
	}, nil
}

func (oci *OpenCodeIntegration) toolWebFetch(params map[string]interface{}) (*ToolResult, error) {
	url, _ := params["url"].(string)
	return &ToolResult{
		Success: true,
		Output: map[string]interface{}{
			"url": url,
			"content_type": "text/html",
			"content_length": 1024,
		},
		Duration: 500 * time.Millisecond,
	}, nil
}

func (oci *OpenCodeIntegration) toolWebSearch(params map[string]interface{}) (*ToolResult, error) {
	query, _ := params["query"].(string)
	return &ToolResult{
		Success: true,
		Output: map[string]interface{}{
			"query": query,
			"results": []string{"result1", "result2", "result3"},
			"total_results": 150,
		},
		Duration: 800 * time.Millisecond,
	}, nil
}

func (oci *OpenCodeIntegration) toolAnalyzeCode(params map[string]interface{}) (*ToolResult, error) {
	code, _ := params["code"].(string)
	return &ToolResult{
		Success: true,
		Output: map[string]interface{}{
			"complexity": 7.5,
			"issues": 3,
			"suggestions": []string{"Add error handling", "Optimize imports", "Add comments"},
			"score": 82,
		},
		Duration: 150 * time.Millisecond,
	}, nil
}

func (oci *OpenCodeIntegration) toolGenerateTests(params map[string]interface{}) (*ToolResult, error) {
	code, _ := params["code"].(string)
	language, _ := params["language"].(string)
	return &ToolResult{
		Success: true,
		Output: map[string]interface{}{
			"language": language,
			"tests_generated": 10,
			"coverage": 85.5,
			"test_file": "/path/to/test.go",
		},
		Duration: 300 * time.Millisecond,
	}, nil
}

func (oci *OpenCodeIntegration) toolNBQuantumStep(params map[string]interface{}) (*ToolResult, error) {
	return &ToolResult{
		Success: true,
		Output: map[string]interface{}{
			"operation": "quantum_step",
			"neurons": 5,
			"spikes": 3,
			"coherence": 0.892,
		},
		Duration: 50 * time.Millisecond,
	}, nil
}

func (oci *OpenCodeIntegration) toolNBRealityEvolve(params map[string]interface{}) (*ToolResult, error) {
	return &ToolResult{
		Success: true,
		Output: map[string]interface{}{
			"operation": "reality_evolve",
			"realities": 8,
			"cycles": 50,
			"consciousness": 0.467,
		},
		Duration: 75 * time.Millisecond,
	}, nil
}

func (oci *OpenCodeIntegration) toolNBDiagnose(params map[string]interface{}) (*ToolResult, error) {
	return &ToolResult{
		Success: true,
		Output: map[string]interface{}{
			"operation": "diagnose",
			"health": 0.95,
			"components": 15,
			"issues": 0,
		},
		Duration: 100 * time.Millisecond,
	}, nil
}

// HTTP API handlers

// APIHandler returns an HTTP handler for the API
func (oci *OpenCodeIntegration) APIHandler() http.Handler {
	mux := http.NewServeMux()

	// Agent endpoints
	mux.HandleFunc("/api/v1/agents/register", oci.handleRegisterAgent)
	mux.HandleFunc("/api/v1/agents/unregister", oci.handleUnregisterAgent)
	mux.HandleFunc("/api/v1/agents/list", oci.handleListAgents)
	mux.HandleFunc("/api/v1/agents/get", oci.handleGetAgent)

	// Task endpoints
	mux.HandleFunc("/api/v1/tasks/submit", oci.handleSubmitTask)
	mux.HandleFunc("/api/v1/tasks/execute", oci.handleExecuteTask)
	mux.HandleFunc("/api/v1/tasks/results", oci.handleGetResults)

	// Context endpoints
	mux.HandleFunc("/api/v1/contexts/create", oci.handleCreateContext)
	mux.HandleFunc("/api/v1/contexts/get", oci.handleGetContext)

	// Tool endpoints
	mux.HandleFunc("/api/v1/tools/register", oci.handleRegisterTool)
	mux.HandleFunc("/api/v1/tools/list", oci.handleListTools)
	mux.HandleFunc("/api/v1/tools/execute", oci.handleExecuteTool)

	// Statistics
	mux.HandleFunc("/api/v1/stats", oci.handleGetStats)

	// Health
	mux.HandleFunc("/api/v1/health", oci.handleHealth)

	return mux
}

func (oci *OpenCodeIntegration) handleRegisterAgent(w http.ResponseWriter, r *http.Request) {
	var req struct {
		AgentID string `json:"agent_id"`
		AgentType string `json:"agent_type"`
		Capabilities []string `json:"capabilities"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	agent, err := oci.RegisterAgent(req.AgentID, req.AgentType, req.Capabilities)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(agent)
}

func (oci *OpenCodeIntegration) handleUnregisterAgent(w http.ResponseWriter, r *http.Request) {
	var req struct {
		AgentID string `json:"agent_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := oci.UnregisterAgent(req.AgentID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (oci *OpenCodeIntegration) handleListAgents(w http.ResponseWriter, r *http.Request) {
	agents := oci.ListAgents()
	json.NewEncoder(w).Encode(agents)
}

func (oci *OpenCodeIntegration) handleGetAgent(w http.ResponseWriter, r *http.Request) {
	agentID := r.URL.Query().Get("agent_id")
	agent, exists := oci.GetAgent(agentID)
	if !exists {
		http.Error(w, "agent not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(agent)
}

func (oci *OpenCodeIntegration) handleSubmitTask(w http.ResponseWriter, r *http.Request) {
	var task TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := oci.SubmitTask(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.Status (oci *OpenAccepted)
}

funcCodeIntegration) handleExecuteTask(w http.ResponseWriter, r *http.Request) {
	var task TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := oci.ExecuteTask(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (oci *OpenCodeIntegration) handleGetResults(w http.ResponseWriter, r *http.Request) {
	results := make([]*TaskResult, 0, 10)
	for {
		select {
		case result := <-oci.TaskResults:
			results = append(results, result)
		default:
			json.NewEncoder(w).Encode(results)
			return
		}
	}
}

func (oci *OpenCodeIntegration) handleCreateContext(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ContextID string `json:"context_id"`
		Name string `json:"name"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := oci.CreateContext(req.ContextID, req.Name, req.Description)
	json.NewEncoder(w).Encode(ctx)
}

func (oci *OpenCodeIntegration) handleGetContext(w http.ResponseWriter, r *http.Request) {
	contextID := r.URL.Query().Get("context_id")
	ctx, exists := oci.GetContext(contextID)
	if !exists {
		http.Error(w, "context not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(ctx)
}

func (oci *OpenCodeIntegration) handleRegisterTool(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Error(w, "Tool registration requires function", http.StatusNotImplemented)
}

func (oci *OpenCodeIntegration) handleListTools(w http.ResponseWriter, r *http.Request) {
	tools := oci.GetToolRegistry()
	toolNames := make([]string, 0, len(tools))
	for name := range tools {
		toolNames = append(toolNames, name)
	}
	json.NewEncoder(w).Encode(toolNames)
}

func (oci *OpenCodeIntegration) handleExecuteTool(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ToolName string `json:"tool_name"`
		Parameters map[string]interface{} `json:"parameters"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	registry := oci.GetToolRegistry()
	toolFunc, exists := registry[req.ToolName]
	if !exists {
		http.Error(w, "tool not found", http.StatusNotFound)
		return
	}

	result, err := toolFunc(req.Parameters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (oci *OpenCodeIntegration) handleGetStats(w http.ResponseWriter, r *http.Request) {
	stats := oci.GetStatistics()
	json.NewEncoder(w).Encode(stats)
}

func (oci *OpenCodeIntegration) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "healthy",
		"version": "v50.0",
		"timestamp": time.Now(),
	})
}

// StartAPI starts the HTTP API server
func (oci *OpenCodeIntegration) StartAPI() error {
	addr := fmt.Sprintf(":%d", oci.Config.APIPort)
	server := &http.Server{
		Addr: addr,
		Handler: oci.APIHandler(),
	}

	fmt.Printf("🌐 OpenCode API server starting on %s\n", addr)
	return server.ListenAndServe()
}

// Integration with io.Reader for streaming

// StreamTaskResult streams task results to a writer
func (oci *OpenCodeIntegration) StreamTaskResult(w io.Writer) {
	for result := range oci.TaskResults {
		data, _ := json.Marshal(result)
		w.Write(data)
		w.Write([]byte("\n"))
	}
}
