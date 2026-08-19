# ============================================================
# CELL 21: CIRCUIT COMPLEXITY REDUCTION
# ============================================================

structure_comparison_df["Gate_Reduction_%"] = (
    (
        structure_comparison_df["Original_Gates"]
        -
        structure_comparison_df["Optimized_Gates"]
    )
    /
    structure_comparison_df["Original_Gates"]
) * 100

structure_comparison_df["Depth_Reduction_%"] = (
    (
        structure_comparison_df["Original_Depth"]
        -
        structure_comparison_df["Optimized_Depth"]
    )
    /
    structure_comparison_df["Original_Depth"]
) * 100

display(structure_comparison_df)
