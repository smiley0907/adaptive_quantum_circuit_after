# ============================================================
# CELL 28: STRUCTURAL COMPARISON
# ============================================================

final_structure_df = pd.merge(
    original_df[
        [
            "Qubits",
            "Gate_Count",
            "Circuit_Depth"
        ]
    ],
    optimized_df[
        [
            "Qubits",
            "Gate_Count",
            "Circuit_Depth"
        ]
    ],
    on="Qubits",
    suffixes=(
        "_Original",
        "_Optimized"
    )
)

final_structure_df[
    "Gate_Reduction_%"
] = (
    (
        final_structure_df["Gate_Count_Original"]
        -
        final_structure_df["Gate_Count_Optimized"]
    )
    /
    final_structure_df["Gate_Count_Original"]
) * 100

final_structure_df[
    "Depth_Reduction_%"
] = (
    (
        final_structure_df["Circuit_Depth_Original"]
        -
        final_structure_df["Circuit_Depth_Optimized"]
    )
    /
    final_structure_df["Circuit_Depth_Original"]
) * 100

display(final_structure_df)
