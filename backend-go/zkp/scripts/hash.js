import { buildPoseidon } from "circomlibjs";

async function main() {
  const poseidon = await buildPoseidon();   // ✅ build it first
  const F = poseidon.F;                    // ✅ field helper

  const inputs = [
    123456789n,   // degreeHash
    85n,          // cgpa
    987654321n,   // issuerDidHash
    111222333n    // issuerSignatureHash
  ];

  const hash = poseidon(inputs);

  // ✅ print as decimal string (what you store/compare typically)
  console.log(F.toString(hash));
}

main().catch(console.error);
