/*

-collects data from all processes, applies a reduction operation(like sum,min,max) and stores the result in the root process.
-non-root processes do not get the result, only the root does.


int MPI_Reduce(const void *sendbuf, void *recvbuf, int count, 
               MPI_Datatype datatype, MPI_Op op, int root, MPI_Comm comm);
parameters:
sendbuf	&local_value	Address of local data (each process sends this)
recvbuf	&global_sum	Address where root stores the result (ignored by other processes)
count	1	Number of elements per process
datatype	MPI_INT	Type of elements being reduced
op	MPI_SUM	Operation to apply (sum, max, min, etc.)
root	0	Process that receives the result
comm	MPI_COMM_WORLD	Group of processes involved


here summing all 
*/
#include <mpi.h>
#include <stdio.h>

int main(int argc, char* argv[]) {
    int rank, size;
    MPI_Init(&argc, &argv);
    MPI_Comm_rank(MPI_COMM_WORLD, &rank);
    MPI_Comm_size(MPI_COMM_WORLD, &size);

    int local_value = rank + 1;  // Each process has a value
    int global_sum;

    // Reduce sum to root process (rank 0)
    MPI_Reduce(&local_value, &global_sum, 1, MPI_INT, MPI_SUM, 0, MPI_COMM_WORLD);

    if (rank == 0) {  // Only root gets the result
        printf("Total sum of values from all processes: %d\n", global_sum);
    }

    MPI_Finalize();
    return 0;
}


