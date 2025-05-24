#include <stdio.h>
#include <stdlib.h>
#include <mpi.h>

int main(int argc, char* argv[])
{
MPI_Init(&argc, &argv);
int my_rank;
MPI_Comm_rank(MPI_COMM_WORLD, &my_rank);
printf("[MPI Process %d]I started waiting for barrier\n",my_rank);
MPI_Barrier(MPI_COMM_WORLD);
printf("[MPI Process %d]I know all MPI processes have waited on the barrier\n",my_rank);
MPI_Finalize();
return EXIT_SUCCESS;
}
