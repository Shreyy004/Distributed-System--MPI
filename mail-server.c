#include <mpi.h>
#include <stdio.h>
#include <string.h>

#define MAX_MSG_LEN 256

int main(int argc, char *argv[]) {
    int rank, size;
    MPI_Init(&argc, &argv);
    MPI_Comm_rank(MPI_COMM_WORLD, &rank);
    MPI_Comm_size(MPI_COMM_WORLD, &size);

    if (rank == 0) {
        printf("Mail Server Started. Waiting for messages...\n");

        // Synchronize before receiving messages
        MPI_Barrier(MPI_COMM_WORLD);

        for (int i = 1; i < size; i++) {
            char message[MAX_MSG_LEN];

            // Receive message from client i
            MPI_Recv(message, MAX_MSG_LEN, MPI_CHAR, i, 0, MPI_COMM_WORLD, MPI_STATUS_IGNORE);

            // Print message after receiving it
            printf("Received email from client %d: %s\n", i, message);

            // Send acknowledgment to client
            MPI_Send("ACK", 4, MPI_CHAR, i, 1, MPI_COMM_WORLD);
        }
    } else {
        // Synchronize all clients before sending messages
        MPI_Barrier(MPI_COMM_WORLD);

        // Client prepares and sends message
        char email[MAX_MSG_LEN];
        snprintf(email, MAX_MSG_LEN, "Hello from Client %d!", rank);
        MPI_Send(email, strlen(email) + 1, MPI_CHAR, 0, 0, MPI_COMM_WORLD);

        // Wait for acknowledgment
        char ack[4];
        MPI_Recv(ack, 4, MPI_CHAR, 0, 1, MPI_COMM_WORLD, MPI_STATUS_IGNORE);
    }

    // Final barrier to synchronize before exit
    MPI_Barrier(MPI_COMM_WORLD);

    MPI_Finalize();
    return 0;
}
