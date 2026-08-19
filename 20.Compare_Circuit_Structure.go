# ============================================================
# CELL 20: ORIGINAL VS OPTIMIZED CIRCUIT STRUCTURE
# ============================================================

structure_comparison = []

for n in QUBIT_CONFIGS:

    original = original_circuits[n]
    optimized = optimized_circuits[n]

    structure_comparison.append({
        "Qubits": n,

        "Original_Gates": original.size(),
        "Optimized_Gates": optimized.size(),

        "Original_Depth": original.depth(),
        "Optimized_Depth": optimized.depth()
    })

structure_comparison_df = pd.DataFrame(
    structure_comparison
)

display(structure_comparison_df)
