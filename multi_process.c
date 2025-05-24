#include "mpi.h"
#include <stdio.h>

#define FROM_MASTER 1

int main(int argc, char **argv){

int num=100;
int n1,n2,n3;
char ack[10]="Recieved";
char ack1[10];
char messager[12];
MPI_Status status;

MPI_Init(&argc,&argv);
int world_size;
MPI_Comm_size(MPI_COMM_WORLD, &world_size);
int world_rank;
MPI_Comm_rank(MPI_COMM_WORLD, &world_rank);

if (world_rank==0){
MPI_Send(&num,1,MPI_INT,1,FROM_MASTER,MPI_COMM_WORLD);
MPI_Send(&num,1,MPI_INT,2,FROM_MASTER,MPI_COMM_WORLD);
MPI_Send(&num,1,MPI_INT,3,FROM_MASTER,MPI_COMM_WORLD);

MPI_Recv(&n1,1,MPI_INT,1,FROM_MASTER,MPI_COMM_WORLD, &status);
MPI_Recv(&n2,1,MPI_INT,2,FROM_MASTER,MPI_COMM_WORLD, &status);
MPI_Recv(&n3,1,MPI_INT,3,FROM_MASTER,MPI_COMM_WORLD, &status);

printf("Square:%d\n",n1);
printf("Cube:%d\n",n2);
printf("Square root:%d\n",n3);


}
else if(world_rank==1){
MPI_Recv(&n1,1,MPI_INT,0,FROM_MASTER,MPI_COMM_WORLD, &status);
int n11=n1*n1;
MPI_Send(&n11,1,MPI_INT,0,FROM_MASTER,MPI_COMM_WORLD);
}
else if(world_rank==2){
MPI_Recv(&n2,1,MPI_INT,0,FROM_MASTER,MPI_COMM_WORLD, &status);
int n22=n2*n2*n2;
MPI_Send(&n22,1,MPI_INT,0,FROM_MASTER,MPI_COMM_WORLD);
}
else if(world_rank==3){
MPI_Recv(&n3,4,MPI_INT,0,FROM_MASTER,MPI_COMM_WORLD, &status);
int n33=n3;
MPI_Send(&n33,1,MPI_INT,0,FROM_MASTER,MPI_COMM_WORLD);
}


MPI_Finalize();
}
