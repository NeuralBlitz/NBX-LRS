/*
NeuralBlitz v50.0 Quantum Spiking Neuron (Go Implementation)
=============================================================

Advanced quantum-classical hybrid neurons that integrate quantum mechanical
effects with classical spiking neuron dynamics.

Implementation Date: 2026-02-05
Language: Go 1.24.0
Phase: Quantum-Classical Hybrid Neurons

Key Features:
- Schrödinger equation integration
- Quantum measurement collapse
- Hamiltonian dynamics
- Spike generation with refractory periods
- Quantum coherence tracking
*/

package systems

import (
	"encoding/json"
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"sync"
	"time"
)

// QuantumSpikingError represents errors in quantum spiking neuron operations
type QuantumSpikingError struct {
	Message string
}

func (e *QuantumSpikingError) Error() string {
	return e.Message
}

// NeuronState represents the state of a quantum spiking neuron
type NeuronState string

const (
	Resting NeuronState = "resting"
	Active  NeuronState = "active"
	Refractory NeuronState = "refractory"
)

// QuantumState represents the quantum state of a neuron
type QuantumState struct {
	Alpha complex128  // |0⟩ coefficient
	Beta  complex128  // |1⟩ coefficient
}

// QuantumSpikingNeuron represents a quantum-classical hybrid neuron
type QuantumSpikingNeuron struct {
	// Neuron parameters
	NeuronID string `json:"neuron_id"`
	
	// Classical parameters
	RestingPotential float64 `json:"resting_potential"`
	ThresholdPotential float64 `json:"threshold_potential"`
	MembraneTimeConstant float64 `json:"membrane_time_constant"`
	MembraneResistance float64 `json:"membrane_resistance"`
	
	// Quantum parameters
	QuantumTunneling float64 `json:"quantum_tunneling"`
	CoherenceTime float64 `json:"coherence_time"`
	
	// State
	MembranePotential float64 `json:"membrane_potential"`
	QuantumState QuantumState `json:"quantum_state"`
	State NeuronState `json:"state"`
	SpikeCount int `json:"spike_count"`
	
	// Timing
	LastSpikeTime float64 `json:"last_spike_time"`
	RefractoryPeriod float64 `json:"refractory_period"`
	TimeSinceLastSpike float64 `json:"time_since_last_spike"`
	
	// Integration parameters
	IntegrationStep float64 `json:"integration_step"`
	MaxHistory int `json:"max_history"`
	
	// History
	SpikeHistory []float64 `json:"spike_history"`
	PotentialHistory []float64 `json:"potential_history"`
	CoherenceHistory []float64 `json:"coherence_history"`
	
	// Quantum measurement results
	LastMeasurementResult int `json:"last_measurement_result"`
	MeasurementTimes []float64 `json:"measurement_times"`
	
	// Synchronization
	mu sync.Mutex
}

// NewQuantumSpikingNeuron creates a new quantum spiking neuron
func NewQuantumSpikingNeuron(neuronID string) *QuantumSpikingNeuron {
	return &QuantumSpikingNeuron{
		NeuronID: neuronID,
		
		// Classical parameters (in mV, ms, MΩ)
		RestingPotential: -70.0,
		ThresholdPotential: -55.0,
		MembraneTimeConstant: 20.0,
		MembraneResistance: 1.0,
		
		// Quantum parameters
		QuantumTunneling: 0.1,
		CoherenceTime: 100.0,
		
		// State
		MembranePotential: -70.0,
		QuantumState: QuantumState{
			Alpha: complex(1.0, 0.0),  // Start in |0⟩
			Beta: complex(0.0, 0.0),
		},
		State: Resting,
		SpikeCount: 0,
		
		// Timing
		LastSpikeTime: -1000.0,
		RefractoryPeriod: 3.0,
		TimeSinceLastSpike: 1000.0,
		
		// Integration
		IntegrationStep: 0.1,
		MaxHistory: 10000,
		
		// History
		SpikeHistory: make([]float64, 0),
		PotentialHistory: make([]float64, 0),
		CoherenceHistory: make([]float64, 0),
		
		// Measurement results
		LastMeasurementResult: 0,
		MeasurementTimes: make([]float64, 0),
	}
}

// CalculateHamiltonian computes the Hamiltonian for the neuron
func (qsn *QuantumSpikingNeuron) CalculateHamiltonian() (complex128, complex128, complex128, complex128) {
	// Normalize membrane potential to [0, 1] range
	V_norm := (qsn.MembranePotential - qsn.RestingPotential) / 
		(qsn.ThresholdPotential - qsn.RestingPotential)
	
	// Pauli matrices
	sigma_z := complex128(complex(1, 0))  // |0⟩⟨0| - |1⟩⟨1|
	sigma_x := complex128(complex(0, 1)) // |0⟩⟨1| + |1⟩⟨0|
	
	// Hamiltonian: H = V(t)σz + Δσx
	H := complex128(complex(V_norm, 0))*sigma_z + 
		complex128(complex(qsn.QuantumTunneling, 0))*sigma_x
	
	return H, sigma_z, sigma_x, complex128(0)
}

// EvolveQuantumState evolves the quantum state using Schrödinger equation
// iℏ d|ψ⟩/dt = Ĥ|ψ⟩
func (qsn *QuantumSpikingNeuron) EvolveQuantumState(dt float64) error {
	qsn.mu.Lock()
	defer qsn.mu.Unlock()
	
	H, _, _, _ := qsn.CalculateHamiltonian()
	
	// Time evolution operator: U(t) = exp(-iHt/ℏ)
	// Using: exp(iAt) where A is anti-Hermitian
	// U = I + At + (At)²/2! + ...
	
	// For small dt, use first-order approximation
	// U ≈ I - iHt (ℏ = 1)
	H_scaled := complex128(complex(dt, 0)) * H
	
	// Matrix exponential using Taylor series (fallback)
	// For better accuracy, use scipy.linalg.expm in Python version
	U := qsn.matrixExponential(H_scaled)
	
	// Apply evolution operator
	newAlpha := U[0][0]*qsn.QuantumState.Alpha + U[0][1]*qsn.QuantumState.Beta
	newBeta := U[1][0]*qsn.QuantumState.Alpha + U[1][1]*qsn.QuantumState.Beta
	
	// Normalize
	norm := cmplx.Abs(newAlpha)*cmplx.Abs(newAlpha) + cmplx.Abs(newBeta)*cmplx.Abs(newBeta)
	if norm > 0 {
		norm = math.Sqrt(norm)
		newAlpha /= complex(norm, 0)
		newBeta /= complex(norm, 0)
	}
	
	qsn.QuantumState.Alpha = newAlpha
	qsn.QuantumState.Beta = newBeta
	
	// Apply decoherence
	qsn.applyDecoherence(dt)
	
	return nil
}

// matrixExponential computes matrix exponential using Taylor series
func (qsn *QuantumSpikingNeuron) matrixExponential(H complex128) [][]complex128 {
	// For 2x2 matrix: exp(H) = I + H + H²/2! + H³/3! + ...
	result := [][]complex128{
		{complex(1, 0), complex(0, 0)},
		{complex(0, 0), complex(1, 0)},
	}
	
	H_power := [][]complex128{
		{complex(1, 0), complex(0, 0)},
		{complex(0, 0), complex(1, 0)},
	}
	
	// Compute exp(H) ≈ I + H + H²/2 + ... (5 terms for accuracy)
	for n := 1; n <= 5; n++ {
		// H_power = H_power * H
		H_matrix := [][]complex128{
			{complex(real(H), imag(H)), complex(imag(H), -real(H))},
			{complex(-imag(H), real(H)), complex(real(H), -imag(H))},
		}
		H_power = qsn.multiplyMatrices(H_power, H_matrix)
		
		// Add H^n/n! to result
		factorial := float64(qsn.factorial(n))
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				result[i][j] += H_power[i][j] / complex(factorial, 0)
			}
		}
	}
	
	return result
}

// multiplyMatrices multiplies two 2x2 matrices
func (qsn *QuantumSpikingNeuron) multiplyMatrices(A, B [][]complex128) [][]complex128 {
	result := [][]complex128{
		{complex(0, 0), complex(0, 0)},
		{complex(0, 0), complex(0, 0)},
	}
	
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				result[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return result
}

// factorial computes n!
func (qsn *QuantumSpikingNeuron) factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// applyDecoherence applies decoherence to quantum state
func (qsn *QuantumSpikingNeuron) applyDecoherence(dt float64) {
	// Decoherence factor: exp(-dt/T₂)
	decoherenceFactor := math.Exp(-dt / qsn.CoherenceTime)
	
	// Off-diagonal elements decay
	// ρ = |α|²|0⟩⟨0| + αβ*|0⟩⟨1| + α*β|1⟩⟨0| + |β|²|1⟩⟨1|
	// Off-diagonal: αβ* and α*β decay
	
	// Apply decoherence to off-diagonal elements
	// This is a simplified model - more sophisticated models exist
	alphaMagSq := real(qsn.QuantumState.Alpha * cmplx.Conj(qsn.QuantumState.Alpha))
	betaMagSq := real(qsn.QuantumState.Beta * cmplx.Conj(qsn.QuantumState.Beta))
	_ = qsn.QuantumState.Alpha * cmplx.Conj(qsn.QuantumState.Beta) * complex(decoherenceFactor, 0)
	
	// Renormalize
	norm := alphaMagSq + betaMagSq
	if norm > 0 {
		alphaMagSq /= norm
		betaMagSq /= norm
	}
	
	// Keep diagonal elements, decay coherence
	// For simplicity, just ensure normalization
	newAlpha := complex(math.Sqrt(alphaMagSq), 0)
	newBeta := complex(math.Sqrt(betaMagSq), 0)
	
	qsn.QuantumState.Alpha = newAlpha
	qsn.QuantumState.Beta = newBeta
}

// Measure performs quantum measurement on the neuron state
func (qsn *QuantumSpikingNeuron) Measure() (int, error) {
	qsn.mu.Lock()
	defer qsn.mu.Unlock()
	
	// Probability of measuring |1⟩
	prob1 := real(qsn.QuantumState.Beta * cmplx.Conj(qsn.QuantumState.Beta))
	
	// Random measurement
	var result int
	if rand.Float64() < prob1 {
		result = 1
	} else {
		result = 0
	}
	
	// Collapse state based on measurement
	if result == 1 {
		qsn.QuantumState.Alpha = complex(0, 0)
		qsn.QuantumState.Beta = complex(1, 0)
	} else {
		qsn.QuantumState.Alpha = complex(1, 0)
		qsn.QuantumState.Beta = complex(0, 0)
	}
	
	qsn.LastMeasurementResult = result
	qsn.MeasurementTimes = append(qsn.MeasurementTimes, float64(time.Now().UnixNano())/1e9)
	
	return result, nil
}

// EvolveMembranePotential evolves membrane potential using exponential Euler
func (qsn *QuantumSpikingNeuron) EvolveMembranePotential(inputCurrent float64, dt float64) error {
	qsn.mu.Lock()
	defer qsn.mu.Unlock()
	
	// Exponential Euler integration
	// dV/dt = -(V - V_rest)/τ + R*I(t)
	
	// Steady-state value for input current
	V_ss := qsn.RestingPotential + qsn.MembraneResistance*inputCurrent
	
	// Decay factor
	decay := math.Exp(-dt / qsn.MembraneTimeConstant)
	
	// Update membrane potential
	qsn.MembranePotential = qsn.MembranePotential*decay + V_ss*(1.0-decay)
	
	return nil
}

// CheckForSpike checks if neuron should generate a spike
func (qsn *QuantumSpikingNeuron) CheckForSpike() bool {
	qsn.mu.Lock()
	defer qsn.mu.Unlock()
	
	// Check if in refractory period
	if qsn.State == Refractory {
		qsn.TimeSinceLastSpike += qsn.IntegrationStep
		
		if qsn.TimeSinceLastSpike >= qsn.RefractoryPeriod {
			qsn.State = Resting
			qsn.TimeSinceLastSpike = 0
		}
		return false
	}
	
	// Check threshold crossing
	if qsn.MembranePotential >= qsn.ThresholdPotential {
		// Generate spike
		qsn.generateSpike()
		return true
	}
	
	return false
}

// generateSpike handles spike generation
func (qsn *QuantumSpikingNeuron) generateSpike() {
	qsn.State = Refractory
	qsn.SpikeCount++
	qsn.LastSpikeTime = float64(time.Now().UnixNano()) / 1e9
	qsn.TimeSinceLastSpike = 0.0
	qsn.SpikeHistory = append(qsn.SpikeHistory, qsn.LastSpikeTime)
	
	// Keep history bounded
	if len(qsn.SpikeHistory) > qsn.MaxHistory {
		qsn.SpikeHistory = qsn.SpikeHistory[len(qsn.SpikeHistory)-qsn.MaxHistory:]
	}
	
	// Reset membrane potential (simplified)
	qsn.MembranePotential = qsn.RestingPotential
}

// Step performs one integration step of the neuron
func (qsn *QuantumSpikingNeuron) Step(inputCurrent float64) (bool, error) {
	// Evolve classical membrane dynamics
	if err := qsn.EvolveMembranePotential(inputCurrent, qsn.IntegrationStep); err != nil {
		return false, err
	}
	
	// Evolve quantum state
	if err := qsn.EvolveQuantumState(qsn.IntegrationStep); err != nil {
		return false, err
	}
	
	// Check for spike
	spiked := qsn.CheckForSpike()
	
	// Record history
	qsn.PotentialHistory = append(qsn.PotentialHistory, qsn.MembranePotential)
	
	// Calculate and record coherence
	coherence := qsn.CalculateCoherence()
	qsn.CoherenceHistory = append(qsn.CoherenceHistory, coherence)
	
	// Keep history bounded
	if len(qsn.PotentialHistory) > qsn.MaxHistory {
		qsn.PotentialHistory = qsn.PotentialHistory[len(qsn.PotentialHistory)-qsn.MaxHistory:]
		qsn.CoherenceHistory = qsn.CoherenceHistory[len(qsn.CoherenceHistory)-qsn.MaxHistory:]
	}
	
	return spiked, nil
}

// CalculateCoherence calculates quantum coherence of the state
func (qsn *QuantumSpikingNeuron) CalculateCoherence() float64 {
	// Coherence = |ρ_01| where ρ is density matrix
	// ρ_01 = αβ*
	rho01 := qsn.QuantumState.Alpha * cmplx.Conj(qsn.QuantumState.Beta)
	return real(rho01 * cmplx.Conj(rho01))
}

// GetState returns comprehensive neuron state
func (qsn *QuantumSpikingNeuron) GetState() map[string]interface{} {
	qsn.mu.Lock()
	defer qsn.mu.Unlock()
	
	return map[string]interface{}{
		"neuron_id": qsn.NeuronID,
		"state": string(qsn.State),
		"membrane_potential": qsn.MembranePotential,
		"spike_count": qsn.SpikeCount,
		"coherence": qsn.CalculateCoherence(),
		"quantum_alpha": qsn.QuantumState.Alpha,
		"quantum_beta": qsn.QuantumState.Beta,
		"last_spike_time": qsn.LastSpikeTime,
		"integration_step": qsn.IntegrationStep,
	}
}

// GetSpikeRate calculates spike rate over recent history
func (qsn *QuantumSpikingNeuron) GetSpikeRate(timeWindow float64) float64 {
	qsn.mu.Lock()
	defer qsn.mu.Unlock()
	
	now := float64(time.Now().UnixNano()) / 1e9
	recentSpikes := 0
	
	for _, spikeTime := range qsn.SpikeHistory {
		if now-spikeTime <= timeWindow {
			recentSpikes++
		}
	}
	
	return float64(recentSpikes) / timeWindow
}

// ToJSON returns JSON representation
func (qsn *QuantumSpikingNeuron) ToJSON() (string, error) {
	state := qsn.GetState()
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// QuantumSpikingNetwork represents a network of quantum spiking neurons
type QuantumSpikingNetwork struct {
	Neurons map[string]*QuantumSpikingNeuron `json:"neurons"`
	Connections map[string][]string `json:"connections"`
	Weights map[string]float64 `json:"weights"`
	
	// Network parameters
	NumNeurons int `json:"num_neurons"`
	InputCurrent float64 `json:"input_current"`
	
	// Synchronization
	mu sync.Mutex
}

// NewQuantumSpikingNetwork creates a new network of quantum spiking neurons
func NewQuantumSpikingNetwork(numNeurons int, inputCurrent float64) *QuantumSpikingNetwork {
	network := &QuantumSpikingNetwork{
		Neurons: make(map[string]*QuantumSpikingNeuron),
		Connections: make(map[string][]string),
		Weights: make(map[string]float64),
		NumNeurons: numNeurons,
		InputCurrent: inputCurrent,
	}
	
	// Create neurons
	for i := 0; i < numNeurons; i++ {
		neuronID := fmt.Sprintf("neuron_%d", i)
		network.Neurons[neuronID] = NewQuantumSpikingNeuron(neuronID)
		network.Connections[neuronID] = make([]string, 0)
	}
	
	// Create random connections (small-world network)
	connectionProbability := 0.1
	for i := 0; i < numNeurons; i++ {
		neuronID := fmt.Sprintf("neuron_%d", i)
		for j := 0; j < numNeurons; j++ {
			if i != j && rand.Float64() < connectionProbability {
				targetID := fmt.Sprintf("neuron_%d", j)
				network.Connections[neuronID] = append(network.Connections[neuronID], targetID)
				network.Weights[neuronID+targetID] = rand.Float64()*0.5 + 0.1
			}
		}
	}
	
	return network
}

// Step advances the network by one integration step
func (qsn *QuantumSpikingNetwork) Step() (int, error) {
	qsn.mu.Lock()
	defer qsn.mu.Lock()
	
	totalSpikes := 0
	
	// Record current spikes
	currentSpikes := make(map[string]bool)
	for neuronID, neuron := range qsn.Neurons {
		spiked, err := neuron.Step(qsn.InputCurrent)
		if err != nil {
			return 0, err
		}
		if spiked {
			currentSpikes[neuronID] = true
			totalSpikes++
		}
	}
	
	// Apply synaptic connections after evolution
	for neuronID, neuron := range qsn.Neurons {
		if currentSpikes[neuronID] {
			for _, targetID := range qsn.Connections[neuronID] {
				weight := qsn.Weights[neuronID+targetID]
				if targetNeuron, ok := qsn.Neurons[targetID]; ok {
					targetNeuron.EvolveMembranePotential(weight*10.0, neuron.IntegrationStep)
				}
			}
		}
	}
	
	return totalSpikes, nil
}

// GetNetworkState returns the state of the entire network
func (qsn *QuantumSpikingNetwork) GetNetworkState() map[string]interface{} {
	qsn.mu.Lock()
	defer qsn.mu.Lock()
	
	neuronStates := make(map[string]interface{})
	for neuronID, neuron := range qsn.Neurons {
		neuronStates[neuronID] = neuron.GetState()
	}
	
	return map[string]interface{}{
		"num_neurons": qsn.NumNeurons,
		"total_spikes": qsn.GetTotalSpikes(),
		"neuron_states": neuronStates,
	}
}

// GetTotalSpikes returns total spikes in the network
func (qsn *QuantumSpikingNetwork) GetTotalSpikes() int {
	total := 0
	for _, neuron := range qsn.Neurons {
		total += neuron.SpikeCount
	}
	return total
}
