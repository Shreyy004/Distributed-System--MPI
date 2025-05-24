#include "mpi.h"
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[])
{
    int rank, value = 0;
    MPI_Init(&argc, &argv);
    MPI_Comm_rank(MPI_COMM_WORLD, &rank);

    while (value >= 0)
    {
        if (rank == 0) {
            printf("Enter a value (negative to stop): ");
            scanf("%d", &value);
        }
        
        // Broadcast the value to all processes
        MPI_Bcast(&value, 1, MPI_INT, 0, MPI_COMM_WORLD);

        // Each process updates the value (for example, incrementing it)
        value++;

        // Print the updated value for each process
        printf("Process %d received value: %d\n", rank, value);
    }

    MPI_Finalize();
    return 0;
}

