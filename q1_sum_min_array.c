#include <stdio.h>
#include <stdlib.h>
#include <mpi.h>

int main (int argc, char **argv)
{
int rank, num_procs;
MPI_Init(&argc, &argv);
MPI_Comm_size (MPI_COMM_WORLD, &num_procs);
MPI_Comm_rank (MPI_COMM_WORLD, &rank);
printf("%d\n", rank);
MPI_Status status;
int min_val;
int arr1[10]= {6, 3, 4, 2, 7, 8, 1, 9, 10, 5};
if (rank ==0)
{
int sum=0;
for (int i=0; i<(sizeof(arr1)/sizeof(arr1[0])); i++) { sum = sum + arr1[i]; }
MPI_Recv ((void *) &min_val, 1, MPI_INT, 1, 1, MPI_COMM_WORLD, &status);
printf("%d: sum=%d ", rank, sum);
printf("received min: %d ", min_val);
printf("%d: goodbye\n", rank);
}
if (rank ==1)
{
int min=arr1[0];
for (int i=0; i<(sizeof(arr1)/sizeof(arr1[0])); i++)
{ if (min>arr1[i]){ min = arr1[i]; } }
MPI_Send ((void *) &min, 1, MPI_INT, 0, 1, MPI_COMM_WORLD);
printf("%d: goodbye\n", rank);
}

MPI_Finalize();
}


