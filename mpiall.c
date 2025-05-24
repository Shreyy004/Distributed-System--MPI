#include "mpi.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MESSAGE_SIZE 6 // "Hello" is 5 characters + null terminator

int main(int argc, char *argv[]) {
    int rank, size;
    char *sbuf, *rbuf;

    MPI_Init(&argc, &argv);
    MPI_Comm_rank(MPI_COMM_WORLD, &rank);
    MPI_Comm_size(MPI_COMM_WORLD, &size);

    // Allocate send and receive buffers (one buffer per process)
    sbuf = (char*)malloc(size * MESSAGE_SIZE * sizeof(char)); // Send buffer
    if (!sbuf) {
        perror("can't allocate send buffer");
        MPI_Abort(MPI_COMM_WORLD, EXIT_FAILURE);
    }

    rbuf = (char*)malloc(size * MESSAGE_SIZE * sizeof(char)); // Receive buffer
    if (!rbuf) {
        perror("can't allocate receive buffer");
        free(sbuf);
        MPI_Abort(MPI_COMM_WORLD, EXIT_FAILURE);
    }

    // Set the send buffer data for each process (sending "Hello")
    for (int i = 0; i < size; i++) {
        snprintf(&sbuf[i * MESSAGE_SIZE], MESSAGE_SIZE, "Hello");
        printf("Process %d sending message to process %d: %s\n", rank, i, &sbuf[i * MESSAGE_SIZE]);
    }

    // Perform the MPI_Alltoall communication
    MPI_Alltoall(sbuf, MESSAGE_SIZE, MPI_CHAR, rbuf, MESSAGE_SIZE, MPI_CHAR, MPI_COMM_WORLD);

    // Print the data received by the process from all other processes
    for (int i = 0; i < size; i++) {
        printf("Process %d received message from process %d: %s\n", rank, i, &rbuf[i * MESSAGE_SIZE]);
    }

    // Free the buffers
    free(rbuf);
    free(sbuf);

    MPI_Finalize();
    return 0;
}

