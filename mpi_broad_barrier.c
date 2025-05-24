#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <mpi.h>

long broadcast (int rank, int num_procs);
void barrier (int rank);
int main(int argc, char **argv)
{
int num_procs;
int rank;
MPI_Init (&argc, &argv);
MPI_Comm_size (MPI_COMM_WORLD, &num_procs);
MPI_Comm_rank (MPI_COMM_WORLD, &rank);

printf("%d: hello (process count=%d)\n", rank, num_procs);
long random_value = broadcast (rank, num_procs);
barrier (rank);
printf("%d: goodbye\n", rank);
MPI_Finalize();

}

long broadcast (int rank, int num_procs)
{
long random_value;
int broadcaster_rank = num_procs - 1;
if (rank==broadcaster_rank)
{
srandom(time(NULL)+rank);
random_value = random() / (RAND_MAX/100);
printf("%d: broadcasting %ld\n", rank, random_value);
}
MPI_Bcast ((void *) &random_value, 1, MPI_LONG, broadcaster_rank, MPI_COMM_WORLD);
if (rank != broadcaster_rank)
{
printf("%d: received %ld\n", rank, random_value);
}
return random_value;
}

void barrier (int rank)
{
printf("%d: enter barrier\n", rank);
MPI_Barrier (MPI_COMM_WORLD);
printf("%d: exit barrier\n", rank);

}
