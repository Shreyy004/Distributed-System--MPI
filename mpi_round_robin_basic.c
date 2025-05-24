#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <mpi.h>

void round_robin (int rank, int num_procs);
int main (int argc, char **argv)
{
int num_procs;
int rank;

MPI_Init (&argc, &argv);
MPI_Comm_size (MPI_COMM_WORLD, &num_procs);
MPI_Comm_rank (MPI_COMM_WORLD, &rank);
printf("%d: hi (number of processes=%d)\n", rank, num_procs);
round_robin (rank, num_procs);
printf("%d: goodbye\n", rank);
MPI_Finalize();
}

void round_robin (int rank, int num_procs)
{
long int rand_mine, rand_prev;
int rank_next = (rank+1)%num_procs;
int rank_prev;
if (rank==0) { rank_prev = num_procs-1; }
else { rank_prev = rank-1;  }
MPI_Status status;
srandom (time(NULL)+rank);
rand_mine = random() / (RAND_MAX/100);
printf("%d: random is %ld\n", rank, rand_mine);
if (rank%2==0)
{
printf ("%d: sending %ld to %d\n", rank, rand_mine, rank_next);
MPI_Send((void *) &rand_mine, 1, MPI_LONG, rank_next, 1, MPI_COMM_WORLD);
printf ("%d: receiving from %d\n", rank, rank_prev);
MPI_Recv((void *) &rand_prev, 1, MPI_LONG, rank_prev, 1, MPI_COMM_WORLD, &status);
}
else
{
printf("%d: receiving from %d\n", rank, rank_prev);
MPI_Recv ((void *) &rand_prev, 1, MPI_LONG, rank_prev, 1, MPI_COMM_WORLD, &status);
printf("%d: sending %ld to %d\n", rank, rand_mine, rank_next);
MPI_Send ((void *) &rand_mine, 1, MPI_LONG, rank_next, 1, MPI_COMM_WORLD);
}
printf("%d: I had %ld; %d had %ld\n", rank, rand_mine, rank_prev, rand_prev);
}

