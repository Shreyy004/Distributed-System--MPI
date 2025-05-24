#include "mpi.h"
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
    int rank, size;
    int *sendbuf = NULL, *recvbuf;
    int total_elements = 12;  // Total number of elements in the array
    int elements_per_process;

    MPI_Init(&argc, &argv);
    MPI_Comm_rank(MPI_COMM_WORLD, &rank);
    MPI_Comm_size(MPI_COMM_WORLD, &size);

    // Check if the number of elements is divisible by the number of processes
    if (total_elements % size != 0) {
        if (rank == 0) {
            printf("Error: The total number of elements is not divisible by the number of processes.\n");
        }
        MPI_Abort(MPI_COMM_WORLD, EXIT_FAILURE);
    }

    // Number of elements each process will handle
    elements_per_process = total_elements / size;

    // Allocate memory for receive buffer
    recvbuf = (int*)malloc(elements_per_process * sizeof(int));
    if (!recvbuf) {
        perror("Unable to allocate memory for recvbuf");
        MPI_Abort(MPI_COMM_WORLD, EXIT_FAILURE);
    }

    if (rank == 0) {
        // Root process initializes the send buffer with data
        sendbuf = (int*)malloc(total_elements * sizeof(int));
        if (!sendbuf) {
            perror("Unable to allocate memory for sendbuf");
            MPI_Abort(MPI_COMM_WORLD, EXIT_FAILURE);
        }

        // Fill the send buffer with numbers (1 to total_elements)
        for (int i = 0; i < total_elements; i++) {
            sendbuf[i] = i + 1;  // Fill with numbers 1, 2, 3, ...
        }

        // Print the data being sent by the root process
        printf("Process %d (Root) sending: ", rank);
        for (int i = 0; i < total_elements; i++) {
            printf("%d ", sendbuf[i]);
        }
        printf("\n");
    }

    // Scatter the data from root to all processes
    MPI_Scatter(sendbuf, elements_per_process, MPI_INT, recvbuf, elements_per_process, MPI_INT, 0, MPI_COMM_WORLD);

    // Each process prints its received data
    printf("Process %d received: ", rank);
    for (int i = 0; i < elements_per_process; i++) {
        printf("%d ", recvbuf[i]);
    }
    printf("\n");

    // Gather all data at the root process
    MPI_Gather(recvbuf, elements_per_process, MPI_INT, sendbuf, elements_per_process, MPI_INT, 0, MPI_COMM_WORLD);

    // Only the root process will print the gathered data
    if (rank == 0) {
        printf("Process %d (Root) gathered: ", rank);
        for (int i = 0; i < total_elements; i++) {
            printf("%d ", sendbuf[i]);
        }
        printf("\n");

        // Free the send buffer as root
        free(sendbuf);
    }

    // Free the receive buffer in all processes
    free(recvbuf);

    MPI_Finalize();
    return 0;
}

