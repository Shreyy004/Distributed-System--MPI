// Blocking send and receive


#include<stdio.h>
#include<mpi.h>

int main(int argc, char** argv) {

 MPI_Init(&argc,&argv);
 
 int rank,size;
 MPI_Comm_rank(MPI_COMM_WORLD, &rank);
 MPI_Comm_size(MPI_COMM_WORLD, &size);

if (size <2) {
printf("Run with atleast 2 processes\n");
MPI_Finalize();	
return 0;
}
int data;
if(rank == 0 ) {
  data = 100;
  printf("Process %d sending data %d to process 1\n",rank,data);
  MPI_Send(&data, 1, MPI_INT, 1, 0, MPI_COMM_WORLD);
}
else if(rank == 1) {
  MPI_Recv(&data, 1, MPI_INT, 0, 0, MPI_COMM_WORLD, MPI_STATUS_IGNORE);
  printf("Process %d received data: %d from process 0\n",rank,data);
}
MPI_Finalize();
return 0;

}
