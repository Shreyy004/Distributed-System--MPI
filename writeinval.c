#include "mpi.h"
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[])
{
    int rank, value = 0, updated_value = 0, invalid_value = -1;
    int num_procs;

    MPI_Init(&argc, &argv);
    MPI_Comm_rank(MPI_COMM_WORLD, &rank);
    MPI_Comm_size(MPI_COMM_WORLD, &num_procs); // Get the total number of processes

    while (value >= 0) {
        if (rank == 0) {
            // Root process takes input from the user
            printf("Enter a value (negative to exit): ");
            scanf("%d", &value);
        }

        // Broadcast the value to all processes
        MPI_Bcast(&value, 1, MPI_INT, 0, MPI_COMM_WORLD);

        // Each process performs a "write-invalidate" (in this case, increments the value)
        updated_value = value + 1;
        printf("Process %d received value %d, updated to %d\n", rank, value, updated_value);

        // After the update, the process invalidates the value by broadcasting the invalid value
        MPI_Bcast(&invalid_value, 1, MPI_INT, 0, MPI_COMM_WORLD);

        if (rank == 0) {
            // Root process will invalidate the value and broadcast it to all processes
            printf("Root process (rank 0) invalidates the value: %d\n", invalid_value);
        }

        // Check if the value has been invalidated by the root process
        if (value == invalid_value) {
            printf("Process %d detected invalid value, updating...\n", rank);
            value = updated_value; // Update the value after invalidation
        }

        // Broadcast the new updated value to all processes after invalidation
        MPI_Bcast(&value, 1, MPI_INT, 0, MPI_COMM_WORLD);
        printf("Process %d received the updated value after invalidation: %d\n", rank, value);
    }

    MPI_Finalize();
    return 0;
}

