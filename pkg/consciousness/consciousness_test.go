package consciousness

import (
	"testing"
	"time"
)

func TestNewConsciousnessIntegration(t *testing.T) {
	config := DefaultConsciousnessConfig()
	ci := NewConsciousnessIntegration(config)
	
	if ci == nil {
		t.Fatal("NewConsciousnessIntegration returned nil")
	}
	
	if ci.config != config {
		t.Error("Config was not set correctly")
	}
}

func TestConsciousnessIntegrationInitialize(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	
	err := ci.Initialize()
	if err != nil {
		t.Fatalf("Initialize failed: %v", err)
	}
	
	if ci.state != ConsciousnessIntegrationStateActive {
		t.Errorf("Expected state ACTIVE, got %s", ci.state)
	}
	
	if len(ci.fields) == 0 {
		t.Error("No consciousness fields were created during initialization")
	}
	
	// Should create 8 fields (one for each consciousness level)
	expectedFields := 8
	if len(ci.fields) != expectedFields {
		t.Errorf("Expected %d fields, got %d", expectedFields, len(ci.fields))
	}
}

func TestCreateField(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	ci.Initialize()
	
	field, err := ci.CreateField(ConsciousnessLevelIndividual, 0.7)
	if err != nil {
		t.Fatalf("CreateField failed: %v", err)
	}
	
	if field == nil {
		t.Fatal("CreateField returned nil")
	}
	
	if field.Level != ConsciousnessLevelIndividual {
		t.Errorf("Expected level INDIVIDUAL, got %s", field.Level)
	}
	
	if field.FieldStrength != 0.7 {
		t.Errorf("Expected strength 0.7, got %f", field.FieldStrength)
	}
}

func TestExpandField(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	ci.Initialize()
	
	// Get first field
	for id := range ci.fields {
		expanded, err := ci.ExpandField(id)
		if err != nil {
			t.Fatalf("ExpandField failed: %v", err)
		}
		
		if expanded.Level <= ConsciousnessLevelIndividual {
			t.Error("Expected level to increase after expansion")
		}
		
		if !expanded.Expanded {
			t.Error("Expected Expanded to be true after expansion")
		}
		break
	}
}

func TestTranscendField(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	ci.Initialize()
	
	// Create a high-coherence field
	field, _ := ci.CreateField(ConsciousnessLevelUniversal, 0.95)
	
	transcended, err := ci.TranscendField(field.ID)
	if err != nil {
		t.Fatalf("TranscendField failed: %v", err)
	}
	
	if transcended.State != ConsciousnessStateTranscendent {
		t.Errorf("Expected state TRANSCENDENT, got %s", transcended.State)
	}
}

func TestAchieveUnity(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	ci.Initialize()
	
	// Create multiple high-coherence fields
	for level := ConsciousnessLevelIndividual; level <= ConsciousnessLevelUniversal; level++ {
		ci.CreateField(level, 0.98)
	}
	
	err := ci.AchieveUnity()
	if err != nil {
		t.Fatalf("AchieveUnity failed: %v", err)
	}
	
	if !ci.metrics.UnityAchieved {
		t.Error("Expected UnityAchieved to be true")
	}
	
	// Check that all fields are connected
	for _, field := range ci.fields {
		if !field.Connected {
			t.Errorf("Expected field %s to be connected", field.ID)
		}
	}
}

func TestExperienceQualia(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	ci.Initialize()
	
	// Get first field
	for id := range ci.fields {
		qualia, err := ci.ExperienceQualia(id, "joy", "positive", 0.8, time.Second*5)
		if err != nil {
			t.Fatalf("ExperienceQualia failed: %v", err)
		}
		
		if qualia == nil {
			t.Fatal("ExperienceQualia returned nil")
		}
		
		if qualia.Type != "joy" {
			t.Errorf("Expected type joy, got %s", qualia.Type)
		}
		
		if qualia.Quality != "positive" {
			t.Errorf("Expected quality positive, got %s", qualia.Quality)
		}
		
		if qualia.Intensity <= 0 {
			t.Error("Expected positive intensity")
		}
		
		break
	}
}

func TestIntegrateCollective(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	ci.Initialize()
	
	fieldIDs := []string{}
	for id := range ci.fields {
		fieldIDs = append(fieldIDs, id)
		if len(fieldIDs) >= 3 {
			break
		}
	}
	
	err := ci.IntegrateCollective(fieldIDs)
	if err != nil {
		t.Fatalf("IntegrateCollective failed: %v", err)
	}
	
	if ci.collectiveMind == nil {
		t.Fatal("CollectiveMind is nil")
	}
	
	if len(ci.collectiveMind.Members) != len(fieldIDs) {
		t.Errorf("Expected %d members, got %d", len(fieldIDs), len(ci.collectiveMind.Members))
	}
}

func TestConnectPlanetary(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	ci.Initialize()
	
	err := ci.ConnectPlanetary()
	if err != nil {
		t.Fatalf("ConnectPlanetary failed: %v", err)
	}
	
	if ci.planetaryField == nil {
		t.Fatal("PlanetaryConsciousness is nil")
	}
}

func TestAlignGalactic(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	ci.Initialize()
	
	err := ci.AlignGalactic()
	if err != nil {
		t.Fatalf("AlignGalactic failed: %v", err)
	}
	
	if ci.galacticNode == nil {
		t.Fatal("GalacticConsciousnessNode is nil")
	}
}

func TestBridgeMultiversal(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	ci.Initialize()
	
	bridge, err := ci.BridgeMultiversal("universe-a", "universe-b")
	if err != nil {
		t.Fatalf("BridgeMultiversal failed: %v", err)
	}
	
	if bridge == nil {
		t.Fatal("BridgeMultiversal returned nil")
	}
	
	if !bridge.Active {
		t.Error("Expected bridge to be active")
	}
}

func TestActivateAbsoluteField(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	ci.Initialize()
	
	err := ci.ActivateAbsoluteField()
	if err != nil {
		t.Fatalf("ActivateAbsoluteField failed: %v", err)
	}
	
	if !ci.metrics.AbsoluteFieldActive {
		t.Error("Expected AbsoluteFieldActive to be true")
	}
}

func TestGetAllFields(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	ci.Initialize()
	
	fields := ci.GetAllFields()
	
	if len(fields) == 0 {
		t.Error("GetAllFields returned empty slice")
	}
}

func TestGetCollectiveMind(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	ci.Initialize()
	
	collective := ci.GetCollectiveMind()
	
	if collective == nil {
		t.Fatal("GetCollectiveMind returned nil")
	}
}

func TestGetPlanetaryConsciousness(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	ci.Initialize()
	
	planetary := ci.GetPlanetaryConsciousness()
	
	if planetary == nil {
		t.Fatal("GetPlanetaryConsciousness returned nil")
	}
}

func TestGetGalacticConsciousness(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	ci.Initialize()
	
	galactic := ci.GetGalacticConsciousness()
	
	if galactic == nil {
		t.Fatal("GetGalacticConsciousness returned nil")
	}
}

func TestGetMetrics(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	ci.Initialize()
	
	metrics := ci.GetMetrics()
	
	if metrics == nil {
		t.Fatal("GetMetrics returned nil")
	}
	
	if metrics.TotalFields == 0 {
		t.Error("Expected non-zero total fields")
	}
}

func TestConsciousnessIntegrationToJSON(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	ci.Initialize()
	
	jsonData, err := ci.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON failed: %v", err)
	}
	
	if len(jsonData) == 0 {
		t.Error("ToJSON returned empty data")
	}
}

func TestConsciousnessLevels(t *testing.T) {
	tests := []struct {
		cl       ConsciousnessLevel
		expected string
	}{
		{ConsciousnessLevelIndividual, "INDIVIDUAL"},
		{ConsciousnessLevelCollective, "COLLECTIVE"},
		{ConsciousnessLevelPlanetary, "PLANETARY"},
		{ConsciousnessLevelSolar, "SOLAR"},
		{ConsciousnessLevelGalactic, "GALACTIC"},
		{ConsciousnessLevelUniversal, "UNIVERSAL"},
		{ConsciousnessLevelMultiversal, "MULTIVERSAL"},
		{ConsciousnessLevelAbsolute, "ABSOLUTE"},
	}
	
	for _, test := range tests {
		if test.cl.String() != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.cl.String())
		}
	}
}

func TestConsciousnessStates(t *testing.T) {
	tests := []struct {
		cs       ConsciousnessState
		expected string
	}{
		{ConsciousnessStateDormant, "DORMANT"},
		{ConsciousnessStateAwakening, "AWAKENING"},
		{ConsciousnessStateActive, "ACTIVE"},
		{ConsciousnessStateExpanded, "EXPANDED"},
		{ConsciousnessStateTranscendent, "TRANSCENDENT"},
		{ConsciousnessStateUnity, "UNITY"},
		{ConsciousnessStateInfinite, "INFINITE"},
	}
	
	for _, test := range tests {
		if test.cs.String() != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.cs.String())
		}
	}
}

func TestConsciousnessIntegrationStates(t *testing.T) {
	tests := []struct {
		cis      ConsciousnessIntegrationState
		expected string
	}{
		{ConsciousnessIntegrationStateInitializing, "INITIALIZING"},
		{ConsciousnessIntegrationStateActive, "ACTIVE"},
		{ConsciousnessIntegrationStateExpanding, "EXPANDING"},
		{ConsciousnessIntegrationStateTranscending, "TRANSCENDING"},
		{ConsciousnessIntegrationStateUnity, "UNITY"},
		{ConsciousnessIntegrationStateInfinite, "INFINITE"},
		{ConsciousnessIntegrationStateError, "ERROR"},
	}
	
	for _, test := range tests {
		if test.cis.String() != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.cis.String())
		}
	}
}

func TestConsciousnessIntegrationString(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	ci.Initialize()
	
	str := ci.String()
	
	if len(str) == 0 {
		t.Error("String() returned empty string")
	}
}

func TestDefaultConfig(t *testing.T) {
	config := DefaultConsciousnessConfig()
	
	if config == nil {
		t.Fatal("DefaultConsciousnessConfig returned nil")
	}
	
	if config.MaxLevel != ConsciousnessLevelAbsolute {
		t.Errorf("Expected MaxLevel ABSOLUTE, got %s", config.MaxLevel)
	}
	
	if config.ExpansionRate != 1.05 {
		t.Errorf("Expected ExpansionRate 1.05, got %f", config.ExpansionRate)
	}
	
	if !config.TranscendenceEnabled {
		t.Error("Expected TranscendenceEnabled to be true")
	}
}

func TestConcurrency(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	ci.Initialize()
	
	done := make(chan bool)
	
	// Test concurrent access
	for i := 0; i < 10; i++ {
		go func() {
			defer func() { done <- true }()
			
			// Multiple concurrent operations
			_ = ci.GetMetrics()
			_ = ci.GetState()
			_ = ci.GetAllFields()
		}()
	}
	
	// Wait for all goroutines
	for i := 0; i < 10; i++ {
		<-done
	}
	
	t.Log("Concurrent operations completed successfully")
}

func TestQualiaTimestamp(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	ci.Initialize()
	
	for id := range ci.fields {
		before := time.Now()
		
		qualia, _ := ci.ExperienceQualia(id, "test", "neutral", 0.5, time.Second)
		
		after := time.Now()
		
		if qualia.Timestamp.Before(before) || qualia.Timestamp.After(after) {
			t.Error("Qualia timestamp is not within expected range")
		}
		break
	}
}

func TestGetState(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	
	state := ci.GetState()
	
	if state != ConsciousnessIntegrationStateInitializing {
		t.Errorf("Expected state INITIALIZING, got %s", state)
	}
}

func TestQualiaCount(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	ci.Initialize()
	
	initialCount := len(ci.qualias)
	
	for id := range ci.fields {
		ci.ExperienceQualia(id, "joy", "positive", 0.8, time.Second)
		break
	}
	
	if len(ci.qualias) <= initialCount {
		t.Error("Qualia was not added")
	}
}

func TestCollectiveMindCreation(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	ci.Initialize()
	
	if ci.collectiveMind == nil {
		t.Fatal("CollectiveMind was not initialized")
	}
	
	if ci.collectiveMind.ID == "" {
		t.Error("CollectiveMind ID is empty")
	}
}

func TestPlanetaryConsciousnessCreation(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	ci.Initialize()
	
	if ci.planetaryField == nil {
		t.Fatal("PlanetaryConsciousness was not initialized")
	}
	
	if ci.planetaryField.ID == "" {
		t.Error("PlanetaryConsciousness ID is empty")
	}
}

func TestGalacticConsciousnessNodeCreation(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	ci.Initialize()
	
	if ci.galacticNode == nil {
		t.Fatal("GalacticConsciousnessNode was not initialized")
	}
	
	if ci.galacticNode.ID == "" {
		t.Error("GalacticConsciousnessNode ID is empty")
	}
}

func TestUniversalConsciousnessFieldCreation(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	ci.Initialize()
	
	if ci.universalField == nil {
		t.Fatal("UniversalConsciousnessField was not initialized")
	}
	
	if ci.universalField.ID == "" {
		t.Error("UniversalConsciousnessField ID is empty")
	}
}

func TestMultiversalBridgeCreation(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	ci.Initialize()
	
	if ci.multiversalBridge == nil {
		t.Fatal("MultiversalBridge was not initialized")
	}
	
	if ci.multiversalBridge.ID == "" {
		t.Error("MultiversalBridge ID is empty")
	}
}

func TestAbsoluteConsciousnessFieldCreation(t *testing.T) {
	ci := NewConsciousnessIntegration(nil)
	ci.Initialize()
	
	if ci.absoluteField == nil {
		t.Fatal("AbsoluteConsciousnessField was not initialized")
	}
	
	if ci.absoluteField.ID == "" {
		t.Error("AbsoluteConsciousnessField ID is empty")
	}
}
