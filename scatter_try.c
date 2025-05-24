#include <mpi.h>
#include <stdio.h>

#define VECTOR_SIZE 5

int main(int argc, char *argv[]) {
    int rank, size;
    double send_buffer[20];  // Large enough to hold multiple copies
    double recv_buffer[VECTOR_SIZE];

    MPI_Init(&argc, &argv);
    MPI_Comm_rank(MPI_COMM_WORLD, &rank);
    MPI_Comm_size(MPI_COMM_WORLD, &size);

    if (rank == 0) {
        // Fill send_buffer with multiple copies of the same vector
        double temp_vector[VECTOR_SIZE] = {1, 2, 3, 4, 5};
        for (int i = 0; i < size; i++) {
            for (int j = 0; j < VECTOR_SIZE; j++) {
                send_buffer[i * VECTOR_SIZE + j] = temp_vector[j];
            }
        }
    }

    // Scatter same data to all processes
    MPI_Scatter(send_buffer, VECTOR_SIZE, MPI_DOUBLE, recv_buffer, VECTOR_SIZE, MPI_DOUBLE, 0, MPI_COMM_WORLD);

    // Print received vector
    printf("Process %d: Received vector = ", rank);
    for (int i = 0; i < VECTOR_SIZE; i++) {
        printf("%.1f ", recv_buffer[i]);
    }
    printf("\n");

    MPI_Finalize();
    return 0;
}
