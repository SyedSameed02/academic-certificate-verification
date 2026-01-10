package zkp

import (
	"bytes"
	"os/exec"
)

func VerifyProof(proofPath, publicPath string) (bool, error) {

	cmd := exec.Command(
		"snarkjs",
		"groth16",
		"verify",
		"zkp/keys/verification_key.json",
		publicPath,
		proofPath,
	)

	var out bytes.Buffer
	var errBuf bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errBuf

	err := cmd.Run()
	if err != nil {
		return false, err
	}

	if bytes.Contains(out.Bytes(), []byte("OK")) {
		return true, nil
	}

	return false, nil
}
