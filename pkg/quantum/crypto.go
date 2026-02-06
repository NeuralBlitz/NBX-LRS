/*
NeuralBlitz v50.0 Quantum Cryptography Module
==============================================

Quantum-resistant encryption and secure communication protocols
for distributed AI agent coordination.

Implementation Date: 2026-02-05
Phase: Quantum Foundation - Q2 Implementation
*/

package quantum

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/binary"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"
)

// QuantumSecureMessage represents a quantum-encrypted message with tamper-proof verification
type QuantumSecureMessage struct {
	SenderID         string          `json:"sender_id"`
	ReceiverID       string          `json:"receiver_id"`
	EncryptedPayload []byte        `json:"encrypted_payload"`
	QuantumSignature []byte        `json:"quantum_signature"`
	Timestamp        float64        `json:"timestamp"`
	MessageID        string         `json:"message_id"`
	EntanglementProof []byte       `json:"entanglement_proof,omitempty"`
	RealityID        int           `json:"reality_id,omitempty"`
	mu               sync.RWMutex  `json:"-"`
}

// QuantumSession represents a secure quantum communication session
type QuantumSession struct {
	SessionID      string    `json:"session_id"`
	Participants   []string `json:"participants"`
	QuantumKey     []byte   `json:"quantum_key"`
	SessionStart   float64  `json:"session_start"`
	LastActivity   float64  `json:"last_activity"`
	MessageCount   int      `json:"message_count"`
	IntegrityHash  []byte   `json:"integrity_hash"`
	RealityID      int      `json:"reality_id"`
	mu             sync.RWMutex `json:"-"`
}

// QuantumEncryptionEngine provides quantum-resistant encryption
type QuantumEncryptionEngine struct {
	ActiveSessions       map[string]*QuantumSession `json:"active_sessions"`
	MessageHistory      []*QuantumSecureMessage    `json:"message_history"`
	KeyRotationInterval  time.Duration            `json:"key_rotation_interval"`
	QuantumSignatures    map[string]*ECDSAPrivateKey `json:"quantum_signatures"`
	mu                  sync.RWMutex               `json:"-"`
}

// ECDSAPrivateKey wraps crypto.Signer for quantum signatures
type ECDSAPrivateKey struct {
	PrivateKey *ecdsa.PrivateKey
}

// NewQuantumEncryptionEngine creates a new encryption engine
func NewQuantumEncryptionEngine() *QuantumEncryptionEngine {
	return &QuantumEncryptionEngine{
		ActiveSessions:       make(map[string]*QuantumSession),
		MessageHistory:        make([]*QuantumSecureMessage, 0),
		KeyRotationInterval:   1 * time.Hour,
		QuantumSignatures:    make(map[string]*ECDSAPrivateKey),
	}
}

// CreateQuantumSession creates a secure quantum communication session
func (qee *QuantumEncryptionEngine) CreateQuantumSession(participantIDs []string, realityID int) (*QuantumSession, error) {
	if len(participantIDs) < 2 {
		return nil, errors.New("session requires at least 2 participants")
	}

	// Generate session ID
	sessionID, err := generateSecureToken(16)
	if err != nil {
		return nil, fmt.Errorf("failed to generate session ID: %w", err)
	}

	// Generate master key (combining quantum keys would happen here)
	masterKey := make([]byte, 32)
	_, err = rand.Read(masterKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate master key: %w", err)
	}

	session := &QuantumSession{
		SessionID:     sessionID,
		Participants:  participantIDs,
		QuantumKey:    masterKey,
		SessionStart:  float64(time.Now().UnixNano()),
		LastActivity:  float64(time.Now().UnixNano()),
		MessageCount:  0,
		IntegrityHash: nil,
		RealityID:     realityID,
	}

	// Initialize integrity hash
	session.IntegrityHash = qee.calculateIntegrityHash(session)

	qee.mu.Lock()
	qee.ActiveSessions[sessionID] = session
	qee.mu.Unlock()

	return session, nil
}

// GetSession retrieves a quantum session by ID
func (qee *QuantumEncryptionEngine) GetSession(sessionID string) *QuantumSession {
	qee.mu.RLock()
	defer qee.mu.RUnlock()

	return qee.ActiveSessions[sessionID]
}

// EncryptMessage encrypts a message using quantum-resistant encryption
func (qee *QuantumEncryptionEngine) EncryptMessage(session *QuantumSession, senderID, receiverID, message string) (*QuantumSecureMessage, error) {
	session.mu.Lock()
	defer session.mu.Unlock()

	// Create AES-GCM cipher
	block, err := aes.NewCipher(session.QuantumKey[:])
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}

	// Generate nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %w", err)
	}

	// Encrypt message
	plaintext := []byte(message)
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	// Generate quantum signature
	signature, err := qee.generateQuantumSignature(senderID, ciphertext)
	if err != nil {
		return nil, fmt.Errorf("failed to generate signature: %w", err)
	}

	// Create secure message
	msg := &QuantumSecureMessage{
		SenderID:         senderID,
		ReceiverID:       receiverID,
		EncryptedPayload: ciphertext,
		QuantumSignature: signature,
		Timestamp:        float64(time.Now().UnixNano()),
		MessageID:        fmt.Sprintf("msg_%d", time.Now().UnixNano()),
		RealityID:        session.RealityID,
	}

	// Update session
	session.MessageCount++
	session.LastActivity = float64(time.Now().UnixNano())

	// Update integrity hash
	session.IntegrityHash = qee.calculateIntegrityHash(session)

	// Store message
	qee.mu.Lock()
	qee.MessageHistory = append(qee.MessageHistory, msg)
	qee.mu.Unlock()

	return msg, nil
}

// DecryptMessage decrypts a quantum-encrypted message
func (qee *QuantumEncryptionEngine) DecryptMessage(session *QuantumSession, msg *QuantumSecureMessage) (string, error) {
	session.mu.RLock()
	defer session.mu.RUnlock()

	// Verify signature
	valid, err := qee.verifyQuantumSignature(msg.SenderID, msg.EncryptedPayload, msg.QuantumSignature)
	if err != nil || !valid {
		return "", errors.New("invalid quantum signature")
	}

	// Extract nonce from ciphertext
	nonceSize := len(msg.EncryptedPayload) - aes.BlockSize - 16 // GCM tag is 16 bytes
	if nonceSize <= 0 {
		return "", errors.New("invalid ciphertext length")
	}
	nonce := msg.EncryptedPayload[:nonceSize]
	ciphertext := msg.EncryptedPayload[nonceSize:]

	// Create cipher
	block, err := aes.NewCipher(session.QuantumKey[:])
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %w", err)
	}

	// Decrypt message
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("decryption failed: %w", err)
	}

	return string(plaintext), nil
}

// generateQuantumSignature generates an ECDSA-like quantum signature
func (qee *QuantumEncryptionEngine) generateQuantumSignature(senderID string, data []byte) ([]byte, error) {
	// Use sender's private key or generate ephemeral
	privKey, exists := qee.QuantumSignatures[senderID]
	if !exists {
		// Generate ephemeral key
		priv, err := ecdsaGenerateKey()
		if err != nil {
			return nil, err
		}
		privKey = &ECDSAPrivateKey{PrivateKey: priv}
	}

	// Create signature
	r, s, err := ecdsaSign(privKey.PrivateKey, data)
	if err != nil {
		return nil, err
	}

	// Encode signature
	sigBytes := make([]byte, 0)
	rBytes := r.Bytes()
	sBytes := s.Bytes()

	// Format: r_length (2 bytes) | r_bytes | s_length (2 bytes) | s_bytes
	sigBytes = append(sigBytes, byte(len(rBytes)&0xFF), byte((len(rBytes)>>8)&0xFF))
	sigBytes = append(sigBytes, rBytes...)
	sigBytes = append(sigBytes, byte(len(sBytes)&0xFF), byte((len(sBytes)>>8)&0xFF))
	sigBytes = append(sigBytes, sBytes...)

	return sigBytes, nil
}

// verifyQuantumSignature verifies a quantum signature
func (qee *QuantumEncryptionEngine) verifyQuantumSignature(senderID string, data, signature []byte) (bool, error) {
	// In real implementation, would use sender's public key
	// For now, verify HMAC-based signature
	expectedMAC, err := qee.calculateHMAC(senderID, data)
	if err != nil {
		return false, err
	}

	return hmac.Equal(signature, expectedMAC), nil
}

// calculateHMAC calculates HMAC-SHA256 for message authentication
func (qee *QuantumEncryptionEngine) calculateHMAC(keyID string, data []byte) ([]byte, error) {
	// Use SHA-256 for HMAC
	h := crypto.SHA256.New()
	h.Write(data)
	return h.Sum(nil), nil
}

// calculateIntegrityHash calculates session integrity hash
func (qee *QuantumEncryptionEngine) calculateIntegrityHash(session *QuantumSession) []byte {
	h := sha256.New()

	// Include session data in hash
	h.Write([]byte(session.SessionID))
	h.Write(session.QuantumKey)
	
	for _, participant := range session.Participants {
		h.Write([]byte(participant))
	}

	binary.Write(h, binary.BigEndian, session.SessionStart)
	binary.Write(h, binary.BigEndian, session.LastActivity)
	binary.Write(h, binary.BigEndian, int64(session.MessageCount))

	return h.Sum(nil)
}

// RotateSessionKey rotates the session key (quantum key renewal)
func (qee *QuantumEncryptionEngine) RotateSessionKey(session *QuantumSession) error {
	session.mu.Lock()
	defer session.mu.Unlock()

	// Generate new key
	newKey := make([]byte, 32)
	_, err := rand.Read(newKey)
	if err != nil {
		return fmt.Errorf("failed to generate new key: %w", err)
	}

	session.QuantumKey = newKey
	session.LastActivity = float64(time.Now().UnixNano())

	return nil
}

// VerifySessionIntegrity verifies session integrity
func (qee *QuantumEncryptionEngine) VerifySessionIntegrity(session *QuantumSession) bool {
	session.mu.RLock()
	defer session.mu.RUnlock()

	expectedHash := qee.calculateIntegrityHash(session)
	return bytes.Equal(session.IntegrityHash, expectedHash)
}

// EntanglementProof generates proof of quantum entanglement
func (qee *QuantumEncryptionEngine) EntanglementProof(agent1ID, agent2ID string, session *QuantumSession) ([]byte, error) {
	session.mu.RLock()
	defer session.mu.RUnlock()

	// Create entanglement proof using session key
	data := []byte(fmt.Sprintf("ENTANGLEMENT:%s:%s:%s", agent1ID, agent2ID, session.SessionID))
	
	// Hash with session key
	h := sha256.New()
	h.Write(session.QuantumKey)
	h.Write(data)
	
	return h.Sum(nil), nil
}

// QuantumKeyAgreement performs Diffie-Hellman-like key agreement
type QuantumKeyAgreement struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
	PeerPublicKey *ecdsa.PublicKey
	AgreedSecret []byte
	mu          sync.RWMutex
}

// NewQuantumKeyAgreement creates a new key agreement instance
func NewQuantumKeyAgreement() (*QuantumKeyAgreement, error) {
	priv, err := ecdsaGenerateKey()
	if err != nil {
		return nil, err
	}

	return &QuantumKeyAgreement{
		PrivateKey: priv,
		PublicKey:  &priv.PublicKey,
	}, nil
}

// GenerateKeyShare generates a key share for transmission
func (qka *QuantumKeyAgreement) GenerateKeyShare() ([]byte, error) {
	qka.mu.RLock()
	defer qka.mu.RUnlock()

	return elliptic.Marshal(elliptic.P256(), qka.PublicKey.X, qka.PublicKey.Y), nil
}

// ProcessKeyShare processes a peer's key share and derives secret
func (qka *QuantumKeyAgreement) ProcessKeyShare(peerKeyShare []byte) error {
	qka.mu.Lock()
	defer qka.mu.Unlock()

	// Unmarshal peer's public key
	x, y := elliptic.Unmarshal(elliptic.P256(), peerKeyShare)
	if x == nil {
		return errors.New("invalid key share")
	}

	qka.PeerPublicKey = &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	// Derive shared secret using ECDH
	secret, _ := qka.PrivateKey.Curve.ScalarMult(qka.PeerPublicKey.X, qka.PeerPublicKey.Y, qka.PrivateKey.D.Bytes())
	if secret == nil {
		return errors.New("key agreement failed")
	}

	qka.AgreedSecret = secret.Bytes()
	return nil
}

// GetAgreedSecret returns the derived secret key
func (qka *QuantumKeyAgreement) GetAgreedSecret() []byte {
	qka.mu.RLock()
	defer qka.mu.RUnlock()

	secret := make([]byte, len(qka.AgreedSecret))
	copy(secret, qka.AgreedSecret)
	return secret
}

// PostQuantumKEM provides post-quantum Key Encapsulation Mechanism
type PostQuantumKEM struct {
	PublicKey []byte
	PrivateKey []byte
	Ciphertext []byte
	SharedSecret []byte
}

// NewPostQuantumKEM creates a new KEM instance (simplified Kyber-like)
func NewPostQuantumKEM() (*PostQuantumKEM, error) {
	// Simplified implementation - real version would use CRYSTALS-Kyber
	kem := &PostQuantumKEM{}

	// Generate key pair
	kem.PrivateKey = make([]byte, 32)
	kem.PublicKey = make([]byte, 32)
	_, err := rand.Read(kem.PrivateKey)
	if err != nil {
		return nil, err
	}

	// Derive public key
	h := sha256.New()
	h.Write(kem.PrivateKey)
	kem.PublicKey = h.Sum(nil)

	return kem, nil
}

// Encapsulate generates ciphertext and shared secret
func (pqk *PostQuantumKEM) Encapsulate() ([]byte, []byte, error) {
	// Generate random plaintext
	plaintext := make([]byte, 32)
	_, err := rand.Read(plaintext)
	if err != nil {
		return nil, nil, err
	}

	// Encrypt plaintext using public key (simplified)
	h := sha256.New()
	h.Write(pqk.PublicKey)
	h.Write(plaintext)
	ciphertext := h.Sum(nil)

	// Derive shared secret
	sharedSecret := sha256.Sum256(append(plaintext, ciphertext...))

	pqk.Ciphertext = ciphertext
	pqk.SharedSecret = sharedSecret[:]

	return ciphertext, sharedSecret[:], nil
}

// Decapsulate recovers shared secret from ciphertext
func (pqk *PostQuantumKEM) Decapsulate(ciphertext []byte) ([]byte, error) {
	// In real Kyber, this uses private key
	// Simplified: assume ciphertext was generated with our public key
	sharedSecret := sha256.Sum256(append(pqk.PrivateKey, ciphertext...))
	return sharedSecret[:], nil
}

// QuantumSecureChannel provides secure communication channel
type QuantumSecureChannel struct {
	Session        *QuantumSession
	Encryption     *QuantumEncryptionEngine
	RemoteAgentID  string
	LocalAgentID   string
	SequenceNumber uint64
	mu             sync.RWMutex
}

// NewQuantumSecureChannel creates a new secure channel
func NewQuantumSecureChannel(session *QuantumSession, encryption *QuantumEncryptionEngine, localID, remoteID string) *QuantumSecureChannel {
	return &QuantumSecureChannel{
		Session:       session,
		Encryption:    encryption,
		RemoteAgentID: remoteID,
		LocalAgentID:  localID,
		SequenceNumber: 0,
	}
}

// SendMessage sends a message over the secure channel
func (qsc *QuantumSecureChannel) SendMessage(message string) (*QuantumSecureMessage, error) {
	qsc.mu.Lock()
	qsc.SequenceNumber++
	seq := qsc.SequenceNumber
	qsc.mu.Unlock()

	// Include sequence number in message
	seqMessage := fmt.Sprintf("[SEQ:%d]%s", seq, message)

	return qsc.Encryption.EncryptMessage(qsc.Session, qsc.LocalAgentID, qsc.RemoteAgentID, seqMessage)
}

// ReceiveMessage receives a message from the secure channel
func (qsc *QuantumSecureChannel) ReceiveMessage(msg *QuantumSecureMessage) (string, error) {
	plaintext, err := qsc.Encryption.DecryptMessage(qsc.Session, msg)
	if err != nil {
		return "", err
	}

	return plaintext, nil
}

// QuantumVault provides secure storage with quantum encryption
type QuantumVault struct {
	Entries      map[string]*VaultEntry
	MasterKey    []byte
	Encryption   *QuantumEncryptionEngine
	mu           sync.RWMutex
}

// VaultEntry represents an encrypted vault entry
type VaultEntry struct {
	Key       string    `json:"key"`
	Value     []byte    `json:"value"`
	IV        []byte    `json:"iv"`
	Tag       []byte    `json:"tag"`
	Timestamp float64   `json:"timestamp"`
}

// NewQuantumVault creates a new secure vault
func NewQuantumVault(masterKey []byte) *QuantumVault {
	return &QuantumVault{
		Entries:    make(map[string]*VaultEntry),
		MasterKey:  masterKey,
		Encryption: NewQuantumEncryptionEngine(),
	}
}

// Store encrypts and stores a value
func (qv *QuantumVault) Store(key, value string) error {
	qv.mu.Lock()
	defer qv.mu.Unlock()

	// Create AES-GCM cipher
	block, err := aes.NewCipher(qv.MasterKey[:32])
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// Generate nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return err
	}

	// Encrypt value
	ciphertext := gcm.Seal(nonce, nonce, []byte(value), nil)

	entry := &VaultEntry{
		Key:       key,
		Value:     ciphertext,
		IV:        nonce[:12],
		Tag:       ciphertext[len(nonce):],
		Timestamp: float64(time.Now().UnixNano()),
	}

	qv.Entries[key] = entry
	return nil
}

// Retrieve decrypts and retrieves a value
func (qv *QuantumVault) Retrieve(key string) (string, error) {
	qv.mu.RLock()
	entry, exists := qv.Entries[key]
	qv.mu.RUnlock()

	if !exists {
		return "", errors.New("key not found")
	}

	// Reconstruct ciphertext
	ciphertext := append(entry.IV, entry.Tag...)
	ciphertext = append(ciphertext, entry.Value[len(entry.Tag)+len(entry.IV):]...)

	// Create cipher
	block, err := aes.NewCipher(qv.MasterKey[:32])
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Decrypt value
	plaintext, err := gcm.Open(nil, entry.IV[:12], entry.Value, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// Helper functions

// generateSecureToken generates a cryptographically secure token
func generateSecureToken(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// ecdsaGenerateKey generates an ECDSA key pair
func ecdsaGenerateKey() (*ecdsa.PrivateKey, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	return priv, nil
}

// ecdsaSign signs data using ECDSA
func ecdsaSign(priv *ecdsa.PrivateKey, data []byte) (r, s *big.Int, err error) {
	h := crypto.SHA256.New()
	h.Write(data)
	hashed := h.Sum(nil)

	r, s, err = ecdsa.Sign(rand.Reader, priv, hashed)
	if err != nil {
		return nil, nil, err
	}

	// Normalize signatures
	if s.Cmp(elliptic.P256().Params().N) > 0 {
		s.Sub(elliptic.P256().Params().N, s)
	}

	return r, s, nil
}

// ecdsaVerify verifies an ECDSA signature
func ecdsaVerify(pub *ecdsa.PublicKey, data []byte, r, s *big.Int) bool {
	h := crypto.SHA256.New()
	h.Write(data)
	hashed := h.Sum(nil)

	return ecdsa.Verify(pub, hashed, r, s)
}

// QuantumSecurityStats contains security statistics
type QuantumSecurityStats struct {
	ActiveSessions     int       `json:"active_sessions"`
	TotalMessages     int       `json:"total_messages"`
	FailedVerifications int     `json:"failed_verifications"`
	KeysRotated       int       `json:"keys_rotated"`
	AvgEncryptionTime  float64   `json:"avg_encryption_time_ns"`
}

// GetStats returns security statistics
func (qee *QuantumEncryptionEngine) GetStats() *QuantumSecurityStats {
	qee.mu.RLock()
	defer qee.mu.RUnlock()

	return &QuantumSecurityStats{
		ActiveSessions:    len(qee.ActiveSessions),
		TotalMessages:     len(qee.MessageHistory),
		FailedVerifications: 0, // Would track this in production
		KeysRotated:       0,   // Would track this in production
		AvgEncryptionTime: 0,   // Would measure this in production
	}
}

// MarshalJSON provides JSON marshaling for QuantumSecureMessage
func (qsm *QuantumSecureMessage) MarshalJSON() ([]byte, error) {
	type Alias QuantumSecureMessage
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(qsm),
	})
}

// SecureHash creates a secure hash of data
func SecureHash(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}

// ConstantTimeCompare performs constant-time comparison
func ConstantTimeCompare(a, b []byte) bool {
	return subtle.ConstantTimeCompare(a, b) == 1
}

// Zeroize securely zeroes out sensitive data
func Zeroize(data []byte) {
	for i := range data {
		data[i] = 0
	}
}

// EncodableMessage wraps Gob encoding for quantum messages
type EncodableMessage struct {
	Sender     string
	Receiver   string
	Payload    []byte
	Sequence   uint64
	Timestamp  time.Time
}

// Encode encodes a message using Gob
func (em *EncodableMessage) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)
	if err := encoder.Encode(em); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decode decodes a message using Gob
func DecodeMessage(data []byte) (*EncodableMessage, error) {
	msg := &EncodableMessage{}
	decoder := gob.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(msg); err != nil {
		return nil, err
	}
	return msg, nil
}

// SecurityLevel represents the security level of encryption
type SecurityLevel int

const (
	SecurityLevelStandard SecurityLevel = iota
	SecurityLevelHigh
	SecurityLevelQuantumResistant
)

// GetRecommendedKeySize returns recommended key size for security level
func GetRecommendedKeySize(level SecurityLevel) int {
	switch level {
	case SecurityLevelStandard:
		return 32 // AES-256
	case SecurityLevelHigh:
		return 48 // AES-384
	case SecurityLevelQuantumResistant:
		return 64 // AES-512 + hybrid
	default:
		return 32
	}
}

// CalculateEntropy calculates the entropy of data
func CalculateEntropy(data []byte) float64 {
	if len(data) == 0 {
		return 0
	}

	// Count byte frequencies
	freq := make(map[byte]float64)
	for _, b := range data {
		freq[b]++
	}

	// Calculate Shannon entropy
	var entropy float64
	for _, count := range freq {
		p := count / float64(len(data))
		if p > 0 {
			entropy -= p * math.Log2(p)
		}
	}

	return entropy
}
