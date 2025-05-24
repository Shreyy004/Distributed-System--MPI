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
int arr1[5] = {1, 7, 3, 6, 4};
if (rank==0)
{
int temp, min_val;
for (int i=0; i<5; i++)
{
for (int j=0; j<5; j++)
{
if (arr1[i]>arr1[j])
{ temp = arr1[j]; arr1[j] = arr1[i]; arr1[i] = temp; } } }
MPI_Recv ((void *) &min_val, 1, MPI_INT, 1, 1, MPI_COMM_WORLD, &status);
printf("%d: sorted array: \n", rank);
for (int i=0; i<5; i++) { printf("%d ", arr1[i]); }
printf("\n");
printf("%d: minimum received=%d ", rank, min_val);
printf("0: Gddbye\n");
}
else
{
int min = arr1[0];
for (int i=1; i<5; i++)
{ if (min>arr1[i]) { min = arr1[i]; } }
MPI_Send((void *) &min, 1, MPI_INT, 0, 1, MPI_COMM_WORLD);
printf("1: Goodbye\n");
}
MPI_Finalize();
}
