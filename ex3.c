#include<stdio.h>
#include <mpi.h>

int main(int argc , char** argv) {
     
    int size,rank;
    MPI_Init(&argc,&argv);
    MPI_Comm_rank(MPI_COMM_WORLD,&rank);
    MPI_Comm_size(MPI_COMM_WORLD,&size);
    
    int n = 1000/size;
    int l_s;
    int s , e;
    
    s = (rank * n) + 1;
    e = (rank+1) * n;
    if (rank == size -1) {
     e = 1000;
    }
    for (int i=s;i<=e;i++) {
     l_s += i;
    }
    printf("The local sum of process %d is %d\n",rank,l_s);
    if (rank == 0) {

      int t_s = l_s;
      int temp;
      for (int i=1;i<size;i++) {
       MPI_Recv(&temp,1,MPI_INT,i,0,MPI_COMM_WORLD,MPI_STATUS_IGNORE);
       t_s += temp;
      }
      
      printf("The total sum of numbers from 0 to 1000 is: %d\n",t_s);
    
    } else {
      
       MPI_Send(&l_s,1,MPI_INT,0,0,MPI_COMM_WORLD);
    }
    
    MPI_Finalize();
}
