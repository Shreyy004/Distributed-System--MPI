#include<mpi.h>
#include<stdio.h>
void main(int argc, char **argv){
int myrank;
MPI_Status status;
double a[100],b[100];
MPI_Request req;
MPI_Init(&argc, &argv);
MPI_Comm_rank(MPI_COMM_WORLD,&myrank);
if(myrank==0){
MPI_Isend(a,100,MPI_DOUBLE,1,17,MPI_COMM_WORLD,&req);
MPI_Irecv(b,100,MPI_DOUBLE,1,18,MPI_COMM_WORLD,&req);
printf("Done");
}
else if(myrank==1){
MPI_Isend(a,100,MPI_DOUBLE,0,18,MPI_COMM_WORLD,&req);
MPI_Irecv(b,100,MPI_DOUBLE,0,17,MPI_COMM_WORLD,&req);
printf("Complete");
}
MPI_Wait(&req,&status);
MPI_Finalize();
}

