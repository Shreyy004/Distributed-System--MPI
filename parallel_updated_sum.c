#include <stdio.h>
#include <stdlib.h>
#include <mpi.h>

int main (int argc, char **argv)
{
int rank, num_procs;
MPI_Init(&argc, &argv);
MPI_Comm_size (MPI_COMM_WORLD, &num_procs);
MPI_Comm_rank (MPI_COMM_WORLD, &rank);
printf("%d: hi\n", rank);
int arr1[10] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
int startval, endval;
MPI_Status status;
if (rank==0)
{
int res[10/num_procs];
for (int i=1; i<num_procs; i++)
{
startval = ((10*i)/num_procs);
endval = ((10*(i+1))/num_procs)-1;
MPI_Recv ((void *) &res, endval-startval+1, MPI_INT, i, 1, MPI_COMM_WORLD, &status);
for (int j=startval; j<=endval; j++) { arr1[j]=res[j-startval]; }
}
printf("updated array: ");
for (int n=0; n<10; n++) { printf("%d ", arr1[n]);  }
printf("%d: goodbye\n", rank);

}
else
{
int start1, end1;
start1 = ((10*rank)/num_procs);	
end1 = ((10*(rank+1))/num_procs) - 1;
for (int m=start1; m<=end1; m++) { arr1[m] = arr1[m]*2; }
MPI_Send ((void *) &arr1[start1], end1-start1+1, MPI_INT, 0, 1, MPI_COMM_WORLD);
printf("%d: goodbye\n", rank);
}

MPI_Finalize();
}

