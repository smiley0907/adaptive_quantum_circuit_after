# ============================================================
# CELL 23: FEEDBACK OPTIMIZED CIRCUIT EXECUTION
# ============================================================

optimized_results = []

for n in QUBIT_CONFIGS:

    circuit = optimized_execution_circuits[n]

    print()
    print("-" * 60)
    print(f"Running Feedback Optimized Circuit: {n} qubits")
    print("-" * 60)

    execution_times = measure_original_circuit(
        circuit,
        shots=SHOTS,
        warmup_runs=WARMUP_RUNS,
        measurement_runs=MEASUREMENT_RUNS,
        seed=RANDOM_SEED + 100 + n
    )

    # --------------------------------------------------------
    # Statistical measurements
    # --------------------------------------------------------

    median_time = float(
        np.median(execution_times)
    )

    mean_time = float(
        np.mean(execution_times)
    )

    std_time = float(
        np.std(execution_times, ddof=1)
    )

    min_time = float(
        np.min(execution_times)
    )

    max_time = float(
        np.max(execution_times)
    )

    optimized_results.append({
        "Qubits": n,
        "Gate_Count": circuit.size(),
        "Circuit_Depth": circuit.depth(),
        "Shots": SHOTS,
        "Median_Time_sec": median_time,
        "Mean_Time_sec": mean_time,
        "Std_Time_sec": std_time,
        "Min_Time_sec": min_time,
        "Max_Time_sec": max_time
    })

    print(
        f"Median : {median_time:.6f} sec"
    )

    print(
        f"Mean   : {mean_time:.6f} sec"
    )

    print(
        f"Std    : {std_time:.6f} sec"
    )
