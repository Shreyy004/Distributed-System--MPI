#include <stdio.h>
#include <stdlib.h>
#include <mpi.h>

int main(int argc, char **argv)
{
int rank, num_procs;
int recvcount;
int recv[100], send[3]={-3, 10, -5};
MPI_Init (&argc, &argv);
MPI_Comm_size (MPI_COMM_WORLD, &num_procs);
MPI_Comm_rank (MPI_COMM_WORLD, &rank);
MPI_Status status;
printf("rank: %d\n", rank);
if (rank==0)
{
for (int i=1; i<num_procs; i++)
{
MPI_Recv((void *) &recv[(i-1)*3], 3, MPI_INT, MPI_ANY_SOURCE, MPI_ANY_TAG, MPI_COMM_WORLD, &status);
printf("0: received from: %d ", status.MPI_SOURCE);
printf("0: tag: %d ", status.MPI_TAG);
printf("0: error: %d ", status.MPI_ERROR);
printf("0: values: \n");
MPI_Get_count (&status, MPI_INT, &recvcount);
for (int j=0; j<recvcount; j++)
{
printf("%d ", recv[((i-1)*3)+j]);
}
printf("0: Goodbye\n");
}

}
else
{
MPI_Send((void *) &send, 3, MPI_INT, 0, 3, MPI_COMM_WORLD);
printf("%d: sending: ", rank);
for (int i=0; i<3; i++)
{
printf("%d ", send[i]);
}
printf("\n");
printf("%d: Goodbye\n", rank);
}
MPI_Finalize();

}
