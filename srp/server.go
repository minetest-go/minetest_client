package srp

import (
	"errors"
)

// Handshake -
//
// Params:
//  A ([]byte) - a client's generated public key
//  v ([]byte) - a client's stored verifer
//
// Return:
//  []byte - the generated public key "B", to be sent to the client
//  []byte - the computed pre-session key "S", to be kept secret
//  []byte - the computed session key "K", to be kept secret
//  error
//
//  NOTE: Only the returned "B" value should be sent to the client. "S" and "K"
//		  are very secret.
//
func Handshake(A, v []byte) ([]byte, []byte, []byte, error) {
	// "A" cannot be zero
	if isZero(A) {
		return nil, nil, nil, errors.New("Server found \"A\" to be zero. Aborting handshake")
	}

	// Create a random secret "b"
	b, err := randomBytes(32)
	if err != nil {
		return nil, nil, nil, err
	}

	// Calculate the SRP-6a version of the multiplier parameter "k"
	k := Hash(dGrp.pad(dGrp.N), dGrp.pad(dGrp.g))

	// Compute a value "B" based on "b"
	//   B = (v + g^b) % N
	B := dGrp.add(dGrp.mul(k, v), dGrp.exp(dGrp.g, b))

	// Calculate "u"
	u := Hash(dGrp.pad(A), dGrp.pad(B))

	// Compute the pseudo-session key, "S"
	//  S = (Av^u) ^ b
	S := dGrp.exp(dGrp.mul(A, dGrp.exp(v, u)), b)

	// The actual session key is the hash of the pseudo-session key "S"
	K := Hash(S)

	return B, S, K, nil
}

// ServerProof -
//
// Params:
//  A ([]byte) - the client's session public key
//  M ([]byte) - the client's proof as computed with ClientProof()
//  K ([]byte) - the computed session secret
//
// Return:
//  []byte - the server's proof of knowing K
//
func ServerProof(A, M, K []byte) []byte {
	return Hash(A, M, K)
}
