#include <mpi.h>
#include <stdio.h>

#define VECTOR_SIZE 3  // We are working with 3D vectors

int main(int argc, char *argv[]) {
    int rank, size;
    double local_vector[VECTOR_SIZE], result_vector[VECTOR_SIZE];

    MPI_Init(&argc, &argv);
    MPI_Comm_rank(MPI_COMM_WORLD, &rank);
    MPI_Comm_size(MPI_COMM_WORLD, &size);

    double all_vectors[size][VECTOR_SIZE];  // Store all vectors at rank 0

    // Each process initializes its own vector
    for (int i = 0; i < VECTOR_SIZE; i++) {
        local_vector[i] = rank + i + 1;  // Example: Process 0 -> (1,2,3), Process 1 -> (2,3,4), etc.
    }

    // Root gathers all vectors from processes
    MPI_Gather(local_vector, VECTOR_SIZE, MPI_DOUBLE, all_vectors, VECTOR_SIZE, MPI_DOUBLE, 0, MPI_COMM_WORLD);

    // Root process computes sum of vectors
    if (rank == 0) {
        for (int i = 0; i < VECTOR_SIZE; i++) {
            result_vector[i] = 0.0;
            for (int j = 0; j < size; j++) {
                result_vector[i] += all_vectors[j][i];  // Summing element-wise
            }
        }

        printf("Final summed vector (using Gather + Manual Sum): (");
        for (int i = 0; i < VECTOR_SIZE; i++) {
            printf("%lf ", result_vector[i]);
        }
        printf(")\n");
    }

    // Broadcast the final summed vector to all processes
    MPI_Bcast(result_vector, VECTOR_SIZE, MPI_DOUBLE, 0, MPI_COMM_WORLD);

    // Print the received vector in all processes
    printf("Process %d received final summed vector: (", rank);
    for (int i = 0; i < VECTOR_SIZE; i++) {
        printf("%lf ", result_vector[i]);
    }
    printf(")\n");

    MPI_Finalize();
    return 0;
}
