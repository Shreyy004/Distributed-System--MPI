#include <stdio.h>
#include <stdlib.h>
#include <mpi.h>

int main(int argc, char* argv[])
{
MPI_Init(&argc, &argv);
int token,world_rank,world_size;
MPI_Comm_rank(MPI_COMM_WORLD,&world_rank);
MPI_Comm_size(MPI_COMM_WORLD,&world_size);
if (world_rank != 0){
MPI_Recv(&token,1,MPI_INT,world_rank-1,0,MPI_COMM_WORLD,MPI_STATUS_IGNORE);
printf("Process %d recieved token %d from process %d\n",world_rank,token,world_rank-1);
}
else{
token=-1;
}
MPI_Send(&token,1,MPI_INT,(world_rank+1)%world_size,0,MPI_COMM_WORLD);
if (world_rank == 0){ 
MPI_Recv(&token,1,MPI_INT,world_rank-1,0,MPI_COMM_WORLD,MPI_STATUS_IGNORE);
printf("Process %d recieved token %d from process %d\n",world_rank,token,world_size-1);
}
MPI_Finalize();
return EXIT_SUCCESS;
}

