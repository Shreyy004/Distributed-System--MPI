
#include <stdio.h>
#include <stdlib.h>
#include <mpi.h>

int main(int argc, char **argv)
{
int rank, num_procs;
MPI_Init (&argc, &argv);
MPI_Comm_size (MPI_COMM_WORLD, &num_procs);
MPI_Comm_rank (MPI_COMM_WORLD, &rank);
printf("%d\n", rank);
int startval, endval;
int res;
MPI_Status status;
if (rank==0)
{
int recv_val;
res=0;
startval = ((1000*rank)/num_procs)+1;
endval = ((1000*(rank+1))/num_procs);
printf("%d: before sum=%d\n", rank, res);
for (int j=startval; j<=endval; j++) { res=res+j; }
printf("%d: start=%d, end=%d\n", rank, startval, endval);
printf("%d: before sum=%d\n", rank, res);
for (int i=1; i<num_procs; i++)
{
MPI_Recv((void *) &recv_val, 1, MPI_INT, i, 1, MPI_COMM_WORLD, &status);
res = res + recv_val;

}
printf("%d: sum=%d", rank, res);
printf("%d: goodbye\n", rank);
}
else
{
int start1, end1;
start1 = ((1000*rank)/num_procs)+1;
end1 = ((1000*(rank+1))/num_procs);
int send_val=0;
printf("%d: start=%d, end=%d\n", rank, start1, end1);
for (int n=start1; n<=end1; n++) { send_val = send_val+n; }
MPI_Send((void *) &send_val, 1, MPI_INT, 0, 1, MPI_COMM_WORLD);
printf("%d: sending %d\n", rank, send_val);
printf("%d: goodbye\n", rank);

}

MPI_Finalize();
}
