package reality

import (
	"testing"
	"time"
)

func TestNewDimensionalComputing(t *testing.T) {
	config := DefaultDimensionalConfig()
	dc := NewDimensionalComputing(config)
	
	if dc == nil {
		t.Fatal("NewDimensionalComputing returned nil")
	}
	
	if dc.config != config {
		t.Error("Config was not set correctly")
	}
	
	if dc.state != DimensionalComputingStateInitializing {
		t.Errorf("Expected state INITIALIZING, got %s", dc.state)
	}
}

func TestDimensionalComputingInitialize(t *testing.T) {
	dc := NewDimensionalComputing(nil)
	
	err := dc.Initialize()
	if err != nil {
		t.Fatalf("Initialize failed: %v", err)
	}
	
	if dc.state != DimensionalComputingStateActive {
		t.Errorf("Expected state ACTIVE, got %s", dc.state)
	}
	
	if len(dc.dimensions) == 0 {
		t.Error("No dimensions were created during initialization")
	}
	
	// Should create 9 default dimensions + 1 hyper dimension = 10
	expectedDims := 10
	if len(dc.dimensions) != expectedDims {
		t.Errorf("Expected %d dimensions, got %d", expectedDims, len(dc.dimensions))
	}
}

func TestAddDimension(t *testing.T) {
	dc := NewDimensionalComputing(nil)
	dc.Initialize()
	
	initialCount := len(dc.dimensions)
	
	dim, err := dc.AddDimension(DimensionTypeSpatial, 1.0, 1.5)
	if err != nil {
		t.Fatalf("AddDimension failed: %v", err)
	}
	
	if dim == nil {
		t.Fatal("AddDimension returned nil")
	}
	
	if dim.Type != DimensionTypeSpatial {
		t.Errorf("Expected type SPATIAL, got %s", dim.Type)
	}
	
	if dim.Coordinate != 1.0 {
		t.Errorf("Expected coordinate 1.0, got %f", dim.Coordinate)
	}
	
	if len(dc.dimensions) != initialCount+1 {
		t.Errorf("Expected %d dimensions, got %d", initialCount+1, len(dc.dimensions))
	}
}

func TestGetDimension(t *testing.T) {
	dc := NewDimensionalComputing(nil)
	dc.Initialize()
	
	// Get first dimension
	for id, dim := range dc.dimensions {
		retrieved, err := dc.GetDimension(id)
		if err != nil {
			t.Fatalf("GetDimension failed: %v", err)
		}
		
		if retrieved.ID != dim.ID {
			t.Errorf("Expected dimension %s, got %s", dim.ID, retrieved.ID)
		}
		break
	}
	
	// Try to get non-existent dimension
	_, err := dc.GetDimension("non-existent")
	if err == nil {
		t.Error("Expected error for non-existent dimension")
	}
}

func TestGetDimensionByType(t *testing.T) {
	dc := NewDimensionalComputing(nil)
	dc.Initialize()
	
	dim, err := dc.GetDimensionByType(DimensionTypeSpatial)
	if err != nil {
		t.Fatalf("GetDimensionByType failed: %v", err)
	}
	
	if dim.Type != DimensionTypeSpatial {
		t.Errorf("Expected type SPATIAL, got %s", dim.Type)
	}
	
	// Try non-existent type
	_, err = dc.GetDimensionByType(DimensionTypeHyper)
	if err != nil {
		// Hyper dimension may not exist if hyper is disabled
		t.Logf("Hyper dimension not found (may be disabled): %v", err)
	}
}

func TestCreateVector(t *testing.T) {
	dc := NewDimensionalComputing(nil)
	dc.Initialize()
	
	coords := map[DimensionType]float64{
		DimensionTypeSpatial: 1.0,
		DimensionTypeTemporal: 0.5,
		DimensionTypeQuantum: 0.8,
	}
	
	vector, err := dc.CreateVector("test-vector", coords)
	if err != nil {
		t.Fatalf("CreateVector failed: %v", err)
	}
	
	if vector == nil {
		t.Fatal("CreateVector returned nil")
	}
	
	if vector.Dimensions[DimensionTypeSpatial] != 1.0 {
		t.Errorf("Expected spatial coordinate 1.0, got %f", vector.Dimensions[DimensionTypeSpatial])
	}
	
	if vector.Magnitude <= 0 {
		t.Error("Expected positive magnitude")
	}
}

func TestProjectVector(t *testing.T) {
	dc := NewDimensionalComputing(nil)
	dc.Initialize()
	
	coords := map[DimensionType]float64{
		DimensionTypeSpatial: 2.0,
		DimensionTypeTemporal: 1.5,
	}
	
	input, err := dc.CreateVector("input-vector", coords)
	if err != nil {
		t.Fatalf("CreateVector failed: %v", err)
	}
	
	projected, err := dc.ProjectVector(input, DimensionTypeSpatial)
	if err != nil {
		t.Fatalf("ProjectVector failed: %v", err)
	}
	
	if projected == nil {
		t.Fatal("ProjectVector returned nil")
	}
	
	if projected.Dimensions[DimensionTypeSpatial] <= 0 {
		t.Error("Expected positive projected dimension")
	}
}

func TestComputeProjection(t *testing.T) {
	dc := NewDimensionalComputing(nil)
	dc.Initialize()
	
	coords := map[DimensionType]float64{
		DimensionTypeSpatial: 3.0,
		DimensionTypeTemporal: 2.0,
	}
	
	input, err := dc.CreateVector("input-vector", coords)
	if err != nil {
		t.Fatalf("CreateVector failed: %v", err)
	}
	
	computation, err := dc.Compute(ComputationTypeProjection, input)
	if err != nil {
		t.Fatalf("Compute failed: %v", err)
	}
	
	if computation == nil {
		t.Fatal("Compute returned nil")
	}
	
	if computation.Type != ComputationTypeProjection {
		t.Errorf("Expected computation type PROJECTION, got %s", computation.Type)
	}
	
	if computation.OutputVector == nil {
		t.Error("Expected output vector in computation")
	}
}

func TestComputeEntanglement(t *testing.T) {
	dc := NewDimensionalComputing(nil)
	dc.Initialize()
	
	coords := map[DimensionType]float64{
		DimensionTypeQuantum: 1.0,
	}
	
	input, err := dc.CreateVector("input-vector", coords)
	if err != nil {
		t.Fatalf("CreateVector failed: %v", err)
	}
	
	computation, err := dc.Compute(ComputationTypeEntanglement, input)
	if err != nil {
		t.Fatalf("Compute failed: %v", err)
	}
	
	if computation == nil {
		t.Fatal("Compute returned nil")
	}
	
	// Entanglement should increase magnitude
	if computation.OutputVector.Magnitude < input.Magnitude {
		t.Errorf("Expected entangled magnitude >= input magnitude")
	}
}

func TestComputeCollapse(t *testing.T) {
	dc := NewDimensionalComputing(nil)
	dc.Initialize()
	
	coords := map[DimensionType]float64{
		DimensionTypeSpatial: 2.0,
		DimensionTypeTemporal: 2.0,
	}
	
	input, err := dc.CreateVector("input-vector", coords)
	if err != nil {
		t.Fatalf("CreateVector failed: %v", err)
	}
	
	computation, err := dc.Compute(ComputationTypeCollapse, input)
	if err != nil {
		t.Fatalf("Compute failed: %v", err)
	}
	
	if computation == nil {
		t.Fatal("Compute returned nil")
	}
	
	// Collapse should reduce magnitude
	if computation.OutputVector.Magnitude > input.Magnitude {
		t.Errorf("Expected collapsed magnitude <= input magnitude")
	}
}

func TestComputeExpansion(t *testing.T) {
	dc := NewDimensionalComputing(nil)
	dc.Initialize()
	
	coords := map[DimensionType]float64{
		DimensionTypeInformation: 1.0,
	}
	
	input, err := dc.CreateVector("input-vector", coords)
	if err != nil {
		t.Fatalf("CreateVector failed: %v", err)
	}
	
	// Increase expansion rate
	dc.config.ExpansionRate = 2.0
	
	computation, err := dc.Compute(ComputationTypeExpansion, input)
	if err != nil {
		t.Fatalf("Compute failed: %v", err)
	}
	
	if computation == nil {
		t.Fatal("Compute returned nil")
	}
	
	// Expansion should increase magnitude
	if computation.OutputVector.Magnitude <= input.Magnitude {
		t.Error("Expected expanded magnitude > input magnitude")
	}
}

func TestCreateRealityBranch(t *testing.T) {
	dc := NewDimensionalComputing(nil)
	dc.Initialize()
	
	modification := map[DimensionType]float64{
		DimensionTypeTemporal: 2.0,
	}
	
	branch, err := dc.CreateRealityBranch("parent-reality", modification)
	if err != nil {
		t.Fatalf("CreateRealityBranch failed: %v", err)
	}
	
	if branch == nil {
		t.Fatal("CreateRealityBranch returned nil")
	}
	
	if branch.Modifications[DimensionTypeTemporal] != 2.0 {
		t.Errorf("Expected modification 2.0, got %f", branch.Modifications[DimensionTypeTemporal])
	}
}

func TestBridgeRealities(t *testing.T) {
	dc := NewDimensionalComputing(nil)
	dc.Initialize()
	
	bridge, err := dc.BridgeReality("reality-a", "reality-b", RealityLinkTypeQuantumEntanglement)
	if err != nil {
		t.Fatalf("BridgeReality failed: %v", err)
	}
	
	if bridge == nil {
		t.Fatal("BridgeReality returned nil")
	}
	
	if bridge.LinkType != RealityLinkTypeQuantumEntanglement {
		t.Errorf("Expected link type QUANTUM_ENTANGLEMENT, got %s", bridge.LinkType)
	}
}

func TestGetMetrics(t *testing.T) {
	dc := NewDimensionalComputing(nil)
	dc.Initialize()
	
	metrics := dc.GetMetrics()
	
	if metrics == nil {
		t.Fatal("GetMetrics returned nil")
	}
	
	if metrics.TotalDimensions == 0 {
		t.Error("Expected non-zero total dimensions")
	}
	
	if metrics.AverageCoherence <= 0 {
		t.Error("Expected positive average coherence")
	}
}

func TestDimensionalComputingString(t *testing.T) {
	dc := NewDimensionalComputing(nil)
	dc.Initialize()
	
	str := dc.String()
	
	if len(str) == 0 {
		t.Error("String() returned empty string")
	}
}

func TestDimensionalComputingToJSON(t *testing.T) {
	dc := NewDimensionalComputing(nil)
	dc.Initialize()
	
	jsonData, err := dc.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON failed: %v", err)
	}
	
	if len(jsonData) == 0 {
		t.Error("ToJSON returned empty data")
	}
}

func TestDimensionTypes(t *testing.T) {
	tests := []struct {
		dt       DimensionType
		expected string
	}{
		{DimensionTypeSpatial, "SPATIAL"},
		{DimensionTypeTemporal, "TEMPORAL"},
		{DimensionTypeQuantum, "QUANTUM"},
		{DimensionTypeInformation, "INFORMATION"},
		{DimensionTypeConsciousness, "CONSCIOUSNESS"},
		{DimensionTypeCausal, "CAUSAL"},
		{DimensionTypeEntropic, "ENTROPIC"},
		{DimensionTypeRealityBranch, "REALITY_BRANCH"},
		{DimensionTypeSemantic, "SEMANTIC"},
		{DimensionTypeHyper, "HYPER"},
	}
	
	for _, test := range tests {
		if test.dt.String() != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.dt.String())
		}
	}
}

func TestDimensionalStates(t *testing.T) {
	tests := []struct {
		ds       DimensionalState
		expected string
	}{
		{DimensionalStateStable, "STABLE"},
		{DimensionalStateUnstable, "UNSTABLE"},
		{DimensionalStateCollapsing, "COLLAPSING"},
		{DimensionalStateExpanding, "EXPANDING"},
		{DimensionalStateOscillating, "OSCILLATING"},
		{DimensionalStateSingular, "SINGULAR"},
	}
	
	for _, test := range tests {
		if test.ds.String() != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.ds.String())
		}
	}
}

func TestComputationTypes(t *testing.T) {
	tests := []struct {
		ct       ComputationType
		expected string
	}{
		{ComputationTypeProjection, "PROJECTION"},
		{ComputationTypeTransformation, "TRANSFORMATION"},
		{ComputationTypeEntanglement, "ENTANGLEMENT"},
		{ComputationTypeCollapse, "COLLAPSE"},
		{ComputationTypeExpansion, "EXPANSION"},
		{ComputationTypeInterference, "INTERFERENCE"},
		{ComputationTypeResonance, "RESONANCE"},
		{ComputationTypeTunneling, "TUNNELING"},
		{ComputationTypeSuperposition, "SUPERPOSITION"},
		{ComputationTypeHolographic, "HOLOGRAPHIC"},
	}
	
	for _, test := range tests {
		if test.ct.String() != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.ct.String())
		}
	}
}

func TestRealityLinkTypes(t *testing.T) {
	tests := []struct {
		rt       RealityLinkType
		expected string
	}{
		{RealityLinkTypeQuantumEntanglement, "QUANTUM_ENTANGLEMENT"},
		{RealityLinkTypeWormhole, "WORMHOLE"},
		{RealityLinkTypeCausalTunnel, "CAUSAL_TUNNEL"},
		{RealityLinkTypeInformationChannel, "INFORMATION_CHANNEL"},
		{RealityLinkTypeConsciousnessLink, "CONSCIOUSNESS_LINK"},
		{RealityLinkTypeDimensionalGateway, "DIMENSIONAL_GATEWAY"},
	}
	
	for _, test := range tests {
		if test.rt.String() != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.rt.String())
		}
	}
}

func TestIdentityMatrix(t *testing.T) {
	matrix := NewIdentityMatrix()
	
	if matrix == nil {
		t.Fatal("NewIdentityMatrix returned nil")
	}
	
	// Check diagonal elements
	for i := 0; i < 4; i++ {
		if matrix.Data[i][i] != 1.0 {
			t.Errorf("Expected diagonal element 1.0 at [%d][%d], got %f", i, i, matrix.Data[i][i])
		}
	}
	
	// Check off-diagonal elements
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i != j && matrix.Data[i][j] != 0.0 {
				t.Errorf("Expected off-diagonal element 0.0 at [%d][%d], got %f", i, j, matrix.Data[i][j])
			}
		}
	}
}

func TestMatrixApply(t *testing.T) {
	matrix := NewIdentityMatrix()
	
	result := matrix.Apply(DimensionTypeSpatial, 5.0)
	
	if result != 5.0 {
		t.Errorf("Expected 5.0, got %f", result)
	}
}

func TestConcurrency(t *testing.T) {
	dc := NewDimensionalComputing(nil)
	dc.Initialize()
	
	done := make(chan bool)
	
	// Test concurrent access
	for i := 0; i < 10; i++ {
		go func() {
			defer func() { done <- true }()
			
			// Multiple concurrent operations
			_, _ = dc.GetMetrics()
			_ = dc.GetState()
			
			coords := map[DimensionType]float64{
				DimensionTypeInformation: float64(i),
			}
			_, _ = dc.CreateVector("test-vector", coords)
		}()
	}
	
	// Wait for all goroutines
	for i := 0; i < 10; i++ {
		<-done
	}
	
	// Verify no crashes
	t.Log("Concurrent operations completed successfully")
}

func TestMetricsUpdate(t *testing.T) {
	dc := NewDimensionalComputing(nil)
	dc.Initialize()
	
	initialMetrics := dc.GetMetrics()
	
	// Add new dimension
	dc.AddDimension(DimensionTypeSemantic, 1.0, 1.0)
	
	updatedMetrics := dc.GetMetrics()
	
	if updatedMetrics.TotalDimensions <= initialMetrics.TotalDimensions {
		t.Error("Metrics were not updated after adding dimension")
	}
}

func TestDefaultConfig(t *testing.T) {
	config := DefaultDimensionalConfig()
	
	if config == nil {
		t.Fatal("DefaultDimensionalConfig returned nil")
	}
	
	if config.MaxDimensions != 11 {
		t.Errorf("Expected MaxDimensions 11, got %d", config.MaxDimensions)
	}
	
	if config.DimensionResolution != 0.001 {
		t.Errorf("Expected DimensionResolution 0.001, got %f", config.DimensionResolution)
	}
	
	if !config.HyperDimensionEnabled {
		t.Error("Expected HyperDimensionEnabled to be true")
	}
}

func TestDimensionalComputingState(t *testing.T) {
	tests := []struct {
		state    DimensionalComputingState
		expected string
	}{
		{DimensionalComputingStateInitializing, "INITIALIZING"},
		{DimensionalComputingStateActive, "ACTIVE"},
		{DimensionalComputingStateComputing, "COMPUTING"},
		{DimensionalComputingStateCollapsing, "COLLAPSING"},
		{DimensionalComputingStateExpanding, "EXPANDING"},
		{DimensionalComputingStateSynchronizing, "SYNCHRONIZING"},
		{DimensionalComputingStateError, "ERROR"},
	}
	
	for _, test := range tests {
		if test.state.String() != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.state.String())
		}
	}
}

func TestVectorTimestamp(t *testing.T) {
	dc := NewDimensionalComputing(nil)
	dc.Initialize()
	
	before := time.Now()
	
	coords := map[DimensionType]float64{
		DimensionTypeSpatial: 1.0,
	}
	vector, _ := dc.CreateVector("test-vector", coords)
	
	after := time.Now()
	
	if vector.Timestamp.Before(before) || vector.Timestamp.After(after) {
		t.Error("Vector timestamp is not within expected range")
	}
}
