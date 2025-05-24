#include <stdio.h>
#include <stdlib.h>
#include <mpi.h>

int main(int argc, char **argv)
{
    int rank, num_procs;
    MPI_Init (&argc, &argv);
    MPI_Comm_size (MPI_COMM_WORLD, &num_procs);
    MPI_Comm_rank (MPI_COMM_WORLD, &rank);
    
    int startval, endval;
    int res = 0;
    MPI_Status status;
    
    // Determine the range for each process
    startval = (1000 * rank) / num_procs + 1;
    endval = (1000 * (rank + 1)) / num_procs;
    
    // Debug print for ranges
    printf("%d: start=%d, end=%d\n", rank, startval, endval);

    // Local sum for each process
    for (int j = startval; j <= endval; j++) {
        res += j;
    }

    if (rank == 0)
    {
        // Process 0 sums its own range
        printf("%d: before sum=%d\n", rank, res);

        // Receive sum from other processes and add it to local sum
        int recv_val;
        for (int i = 1; i < num_procs; i++)
        {
            MPI_Recv(&recv_val, 1, MPI_INT, i, 1, MPI_COMM_WORLD, &status);
            res += recv_val;
        }
        
        printf("%d: total sum=%d\n", rank, res);
    }
    else
    {
        // Send the result from other processes to process 0
        MPI_Send(&res, 1, MPI_INT, 0, 1, MPI_COMM_WORLD);
    }

    printf("%d: goodbye\n", rank);
    MPI_Finalize();
}

