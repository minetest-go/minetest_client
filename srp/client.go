//
// In SRP, the client is responsible for quite a bit of the cryptographic
// processing. This has advantages and disadvantages, but a noteworthy advantage
// is the inherent trust that comes with generating random values on one's own
// hardware.
//
// At a high level, a client is responsible for the following tasks:
//  1. Generating their salt and verifier
//  2. Initiating an authentication handshake
//  3. Generating a key based on the server's response to (2)
//  4. Providing a proof of K
//  5. Verifying a server's proof of K
//

package srp

import (
	"errors"
)

// NewClient -
//
// Params:
//  I ([]byte) - a client's identifier
//  p ([]byte) - a client's passphrase
//
// Return:
//  []byte - a client's salt
//  []byte - a client's verifer
//  error
//
func NewClient(I, p []byte) ([]byte, []byte, error) {
	// Generate a random salt. Default to 32 bytes.
	//   <salt> = random()
	s, err := randomBytes(32)
	if err != nil {
		return nil, nil, err
	}

	// Compute the secret "x" value
	//   x = SHA(<salt> | SHA(<username> | ":" | <raw password>))
	x := Hash(s, Hash(I, []byte(":"), p))

	// Calculate the verifer.
	//   <password verifier> = v = g^x % N
	v := dGrp.exp(dGrp.g, x)

	return s, v, nil
}

// InitiateHandshake -
//
// Params:
//  None
//
// Return:
//  []byte - the client's public key "A" to be sent to the server
//  []byte - the client's private key "a" to complete the handshake
//  error
//
func InitiateHandshake() ([]byte, []byte, error) {
	// Create a random secret "a" value
	//   a = random()
	a, err := randomBytes(32)
	if err != nil {
		return nil, nil, err
	}

	// Calculate "A" based on "a"
	//   A = g^a % N
	A := dGrp.exp(dGrp.g, a)

	return A, a, nil
}

// CompleteHandshake -
//
// Params:
//  A ([]byte) - the client's session public key
//  a ([]byte) - the client's session private key
//  I ([]byte) - the client's identifier
//  p ([]byte) - the client's secret passphrase
//  s ([]byte) - the client's salt looked up by the server
//  B ([]byte) - the server's public key for this session
//
// Return:
//  []byte - the client's computed session key
//  error
//
func CompleteHandshake(A, a, I, p, s, B []byte) ([]byte, error) {
	// "B" cannot be zero
	if isZero(B) {
		return nil, errors.New("\"B\" value is zero. Aborting handshake")
	}

	// Calculate "u"
	u := Hash(dGrp.pad(A), dGrp.pad(B))

	// "u" cannot be zero
	if isZero(u) {
		return nil, errors.New("\"u\" value is zero. Aborting handshake")
	}

	// Compute the secret "x" value
	//   x = SHA(<salt> | SHA(<username> | ":" | <raw password>))
	x := Hash(s, Hash(I, []byte(":"), p))

	// Calculate the SRP-6a version of the multiplier parameter "k"
	k := Hash(dGrp.N, dGrp.pad(dGrp.g))

	// Compute the pseudo-session key, "S"
	//   S = (B - kg^x) ^ (a + ux)
	//
	//    let l = (B - kg^x),
	//        r = (a + ux)
	//
	//    ... so that S = l ^ r
	l := dGrp.sub(B, dGrp.mul(k, dGrp.exp(dGrp.g, x)))
	r := dGrp.add(a, dGrp.mul(u, x))
	S := dGrp.exp(l, r)

	// The actual session key is the hash of the pseudo-session key "S"
	K := Hash(S)

	// Return K
	return K, nil
}

// ClientProof -
//
// Params:
//  I ([]byte) - the client's identifier
//  s ([]byte) - the client's salt looked up by the server
//  A ([]byte) - the client's session public key
//  B ([]byte) - the server's public key for this session
//  K ([]byte) - the client's computed session key
//
// Return:
//  []byte - the client's proof of knowing K
//
func ClientProof(I, s, A, B, K []byte) []byte {
	h_N := Hash(dGrp.N)
	h_g := Hash(dGrp.g)
	h_xor := make([]byte, 32)
	for i := 0; i < 32; i++ {
		h_xor[i] = h_N[i] ^ h_g[i]
	}
	
	return Hash(h_xor, Hash(I), s, A, B, K)
}
