# ============================================================
# CELL 22: PREPARE OPTIMIZED CIRCUITS FOR EXECUTION
# ============================================================

optimized_execution_circuits = {}

for n in QUBIT_CONFIGS:

    optimized_execution_circuits[n] = optimized_circuits[n]

print("Optimized circuits prepared for execution.")
