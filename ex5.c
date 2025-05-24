#include <mpi.h>
#include <stdio.h>

#define VECTOR_SIZE 10  // Not necessarily divisible by size

int main(int argc, char *argv[]) {
    int rank, size;
    double vector[VECTOR_SIZE];  
    double local_vector[VECTOR_SIZE];  // Large enough to fit any portion
    double local_sum = 0.0, global_sum = 0.0;
    double gathered_sums[4];  

    MPI_Init(&argc, &argv);
    MPI_Comm_rank(MPI_COMM_WORLD, &rank);
    MPI_Comm_size(MPI_COMM_WORLD, &size);

    int base_size = VECTOR_SIZE / size;
    int last_size = VECTOR_SIZE - base_size * (size - 1);  // Remaining elements for last process
    int local_size = (rank == size - 1) ? last_size : base_size;

    if (rank == 0) {
        for (int i = 0; i < VECTOR_SIZE; i++) {
            vector[i] = i + 1;
        }
    }

    int counts[size], displs[size];  // For variable-sized Scatter
    for (int i = 0; i < size; i++) {
        counts[i] = (i == size - 1) ? last_size : base_size; //Number of elements each process receives.
        displs[i] = (i == 0) ? 0 : displs[i - 1] + counts[i - 1]; //The start index of the previous process (displs[i - 1])+ the number of elements the previous process received (counts[i - 1])
    }

    MPI_Scatterv(vector, counts, displs, MPI_DOUBLE, local_vector, local_size, MPI_DOUBLE, 0, MPI_COMM_WORLD);

    for (int i = 0; i < local_size; i++) {
        local_sum += local_vector[i];
    }

    MPI_Gather(&local_sum, 1, MPI_DOUBLE, gathered_sums, 1, MPI_DOUBLE, 0, MPI_COMM_WORLD);
    if (rank == 0) {
        global_sum = 0.0;
        for (int i = 0; i < size; i++) {
            global_sum += gathered_sums[i];
        }
        printf("Global sum using MPI_Gather: %f\n", global_sum);
    }

    MPI_Reduce(&local_sum, &global_sum, 1, MPI_DOUBLE, MPI_SUM, 0, MPI_COMM_WORLD);
    if (rank == 0) {
        printf("Global sum using MPI_Reduce: %f\n", global_sum);
    }

    MPI_Allreduce(&local_sum, &global_sum, 1, MPI_DOUBLE, MPI_SUM, MPI_COMM_WORLD);
    printf("Process %d: Global sum using MPI_Allreduce: %f\n", rank, global_sum);

    MPI_Finalize();
    return 0;
}
