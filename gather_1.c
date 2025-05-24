/*
it is possible! Instead of relying on MPI_Gather, each process can manually send its data to the root, and the root collects the data in an array. This can be done using MPI_Send and MPI_Recv instead of MPI_Gather.

Each process sends its data to rank 0 using MPI_Send.
Rank 0 receives all data using MPI_Recv and stores it in an array.
After collecting all data, rank 0 broadcasts the final array using MPI_Bcast (or MPI_Send to another process if needed).



*/

#include <mpi.h>
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char* argv[]) {
    int rank, size;
    MPI_Init(&argc, &argv);
    MPI_Comm_rank(MPI_COMM_WORLD, &rank);
    MPI_Comm_size(MPI_COMM_WORLD, &size);

    int local_data = (rank + 1) * 10;  // Each process has different data
    int* collected_data = NULL;  // Will store all received data at root

    if (rank == 0) {
        collected_data = (int*)malloc(size * sizeof(int));  // Allocate array at root
    }

    if (rank != 0) {
        
        MPI_Send(&local_data, 1, MPI_INT, 0, 0, MPI_COMM_WORLD);
    } else {
        
        collected_data[0] = local_data; // its own data

        for (int i = 1; i < size; i++) {
            MPI_Recv(&collected_data[i], 1, MPI_INT, i, 0, MPI_COMM_WORLD, MPI_STATUS_IGNORE);
        }

        
        printf("Root collected data: ");
        for (int i = 0; i < size; i++) {
            printf("%d ", collected_data[i]);
        }
        printf("\n");
}

    MPI_Finalize();
    return 0;
}

