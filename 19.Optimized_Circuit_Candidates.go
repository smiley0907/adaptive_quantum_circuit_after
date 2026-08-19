# ============================================================
# CELL 19: GENERATE OPTIMIZED CIRCUIT CANDIDATES
# ============================================================

optimized_circuits = {}

for n in QUBIT_CONFIGS:

    original_circuit = original_circuits[n]

    optimized_circuit = transpile(
        original_circuit,
        simulator,
        optimization_level=3,
        seed_transpiler=RANDOM_SEED
    )

    optimized_circuits[n] = optimized_circuit

print("Optimized circuit candidates generated successfully.")
