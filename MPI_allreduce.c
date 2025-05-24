/*
MPI_Allreduce has the same parameters as MPI_Reduce, except that every process gets the result instead of just the root.
No separate root process → The result is available everywhere.


int MPI_Allreduce(const void *sendbuf, void *recvbuf, int count, 
                  MPI_Datatype datatype, MPI_Op op, MPI_Comm comm);
 sendbuf	Data sent by each process
recvbuf	Address where every process stores the result
count	Number of elements per process
datatype	Type of data (MPI_INT, MPI_FLOAT, etc.)
op	Reduction operation (SUM, MAX, MIN, etc.)
comm	Communicator (group of processes)

*/

#include <mpi.h>
#include <stdio.h>

int main(int argc, char* argv[]) {
    int rank, size;
    MPI_Init(&argc, &argv);
    MPI_Comm_rank(MPI_COMM_WORLD, &rank);
    MPI_Comm_size(MPI_COMM_WORLD, &size);

    int local_value = (rank + 1) * 10;  // Different values per process
    int max_value;

    // Compute maximum value across all processes
    MPI_Allreduce(&local_value, &max_value, 1, MPI_INT, MPI_MAX, MPI_COMM_WORLD);

    // Each process prints the same result
    printf("Process %d: Max value among all processes is %d\n", rank, max_value);

    MPI_Finalize();
    return 0;
}

