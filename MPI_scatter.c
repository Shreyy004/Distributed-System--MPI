/*
-is a collective communiation fn in mpi that distributes an array of data from a single 
root process to all processes in the communicator.
- usually root - 0
- data is divided into equal-sized chunks and sent to each process.
- each process receives a single chunk.



syntax:
int MPI_Scatter(
    void *sendbuf,    // Address of send buffer (on root)
    int sendcount,    // Number of elements sent to each process
    MPI_Datatype sendtype, // Type of data being sent
    void *recvbuf,    // Address of receive buffer (on each process)
    int recvcount,    // Number of elements received by each process
    MPI_Datatype recvtype, // Type of data received
    int root,         // Rank of root process
    MPI_Comm comm     // Communicator
);

The sendbuf in the root process contains all the data to be distributed.
Each process (including root) receives sendcount elements.
The size of sendbuf must be sendcount × num_processes.
Each non-root process does not use sendbuf.
Data is distributed equally among all processes.

--> distributes data in rank order by default.
if procceses are more than buffer size then it sends 0.


Key Takeaways

✅ MPI_Scatter divides data from one process to many.
✅ Only root provides the data, and all processes receive a portion.
✅ Useful for work distribution, parallel processing, and data partitioning.


code:
#include <mpi.h>
#include <stdio.h>

int main(int argc, char* argv[]) {
    int rank, size;
    MPI_Init(&argc, &argv);
    MPI_Comm_rank(MPI_COMM_WORLD, &rank);
    MPI_Comm_size(MPI_COMM_WORLD, &size);

    int send_data[4] = {100, 200, 300, 400};  // Original order
    int recv_data;

    // Custom order mapping (root only modifies sendbuf)
    int custom_order[4] = {300, 100, 400, 200};  // New order

    // Root process should assign the reordered array before scattering
    if (rank == 0) {
        for (int i = 0; i < 4; i++) {
            send_data[i] = custom_order[i];
        }
    }

    // Scatter data with custom order
    MPI_Scatter(send_data, 1, MPI_INT, &recv_data, 1, MPI_INT, 0, MPI_COMM_WORLD);

    // Print received values
    printf("Process %d received %d\n", rank, recv_data);

    MPI_Finalize();
    return 0;
}



Key Takeaways

    MPI_Scatter sends data in order of ranks.
    If you want a different order, you must modify sendbuf manually before scattering.
    This is useful when:
        You need to distribute tasks based on priority.
        Some processes need specific data first.
        You want a non-sequential distribution.
*/


#include<stdio.h>
#include<mpi.h>


int main(int argc, char** argv) {

   int rank, size;
   MPI_Init(&argc, &argv);
   MPI_Comm_rank(MPI_COMM_WORLD, &rank);
   MPI_Comm_size(MPI_COMM_WORLD, &size);
   
   int data[8] = {10, 20, 30, 40, 50, 60, 70, 80}; //root's data
   int recv_data;
   
   MPI_Scatter(data, 1, MPI_INT, &recv_data, 1, MPI_INT, 0, MPI_COMM_WORLD);
   
   printf("Process %d received %d\n",rank,recv_data);
   
   MPI_Finalize();
   return 0;
  

}

















