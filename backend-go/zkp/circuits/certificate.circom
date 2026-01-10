pragma circom 2.1.6;

include "circomlib/circuits/poseidon.circom";
include "circomlib/circuits/comparators.circom";

template CertificateProof() {
    signal input onChainHash;
    signal input MIN_CGPA;

    signal input degreeHash;
    signal input cgpa;
    signal input issuerDidHash;
    signal input issuerSignatureHash;

    component h = Poseidon(4);
    h.inputs[0] <== degreeHash;
    h.inputs[1] <== cgpa;
    h.inputs[2] <== issuerDidHash;
    h.inputs[3] <== issuerSignatureHash;

    h.out === onChainHash;

    component gte = GreaterEqThan(32);
    gte.in[0] <== cgpa;
    gte.in[1] <== MIN_CGPA;
    gte.out === 1;
}

component main = CertificateProof();
