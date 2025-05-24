#include "mpi.h"
#include <stdio.h>

#define FROM_MASTER 1

int main(int argc, char **argv){

char message='a';
char ack[10]="Recieved";
char ack1[10];
int messager;
MPI_Status status;

MPI_Init(&argc,&argv);
int world_size;
MPI_Comm_size(MPI_COMM_WORLD, &world_size);
int world_rank;
MPI_Comm_rank(MPI_COMM_WORLD, &world_rank);

if (world_rank==0){
MPI_Send(&message,1,MPI_CHAR,1,FROM_MASTER,MPI_COMM_WORLD);
MPI_Recv(&ack1,10,MPI_CHAR,1,FROM_MASTER,MPI_COMM_WORLD,&status);
printf("%s by 1\n",ack1);
}
else if(world_rank==1){
MPI_Recv(&messager,1,MPI_INT,0,FROM_MASTER,MPI_COMM_WORLD, &status);
printf("Process %d says : %d\n",world_rank,messager);
MPI_Send(&ack,10,MPI_CHAR,0,FROM_MASTER,MPI_COMM_WORLD);
}


MPI_Finalize();
}






