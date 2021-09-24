package srp

/*
func TestFullHandshake(t *testing.T) {
	// First, register a client.
	// I is defined.
	identifier := []byte("testuser")
	passphrase := []byte("Password123!")
	s, v, err := NewClient(identifier, passphrase)
	if err != nil {
		t.Errorf("Error in NewClient()")
	}

	// *************************************************************************
	// Client sends I, s, v to the server.
	// *************************************************************************

	// The client now initiates a handshake to create a session
	A, a, err := InitiateHandshake()
	if err != nil {
		t.Errorf("Error in InitiateHandshake()")
	}

	// *************************************************************************
	// Client sends I, A to the server...
	// *************************************************************************

	// Lookup "v" and "s" from "I"

	B, S, serverK, err := Handshake(A, v)
	if err != nil {
		t.Errorf("Error in Handshake()")
	}
	if testing.Verbose() {
		fmt.Println("Server K: " + formatBytes(serverK))
	}

	// *************************************************************************
	// Server sends B, s to the client...
	// *************************************************************************

	clientK, err := CompleteHandshake(A, a, identifier, passphrase, s, B)
	if err != nil {
		t.Errorf("Error in CompleteHandshake()")
	}
	if testing.Verbose() {
		fmt.Println("Client K: " + formatBytes(clientK))
	}

	// *************************************************************************
	// Client and server MIGHT have a shared K.
	// *************************************************************************

	// These proofs will almost certainly not fail, but we invoke the functions
	// just to be diligent.
	clientProof := ClientProof(identifier, A, B, S)
	ServerProof(A, clientProof, serverK)

	// *************************************************************************
	// Client and server SHOULD have a shared K.
	// *************************************************************************

	if !bytes.Equal(serverK, clientK) {
		t.Errorf("Server K does not match Client K. Proofs failed!")
	}
}

//
// Helper method that takes a byte slice, and returns a pretty-printable
// base64 string.
//
func formatBytes(x []byte) string {
	// Convert bytes to a base64 string
	str := base64.StdEncoding.EncodeToString(x)

	// Return up to 40 characters of the base64 string
	if len(str) > 40 {
		return fmt.Sprintf("%v...", str[:40])
	} else {
		return str
	}
}
*/
