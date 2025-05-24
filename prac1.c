#include<stdio.h>
#include<mpi.h>

#define c 3
int main(int argc, char** argv) {

   int size, rank;
   MPI_Init(&argc,&argv);
   MPI_Comm_rank(MPI_COMM_WORLD,&rank);
   MPI_Comm_size(MPI_COMM_WORLD,&size);
   int N = size;
   int vector[N][c];
   int counter = 0;
   if (rank == 0) {
          
        for(int i=0;i<N;i++) {
    for(int j=0;j<c;j++) {
      vector[i][j] = ++counter;
      }
    }
    }
   
    int local_data[N];
    MPI_Scatter(vector,c,MPI_INT,local_data,c,MPI_INT,0,MPI_COMM_WORLD);
    
    printf("Process %d received data %ls \n",rank,local_data);
    int gather_data[N][3];
    MPI_Gather(local_data,c,MPI_INT,gather_data,c,MPI_INT,0,MPI_COMM_WORLD);
    int vector_sum[3];
    if (rank == 0) {
          
          
          for(int i=0;i<c;i++) {
             vector_sum[i] = 0;
             for(int j=0;j<N;j++) {
                    vector_sum[i] += gather_data[j][i]; 
      } 
    }
    printf("The total summed vector is: ");    
    for(int i=0;i<3;i++) {
      printf("%d ",vector_sum[i]);
    }
    printf("\n");
    
    }
    MPI_Bcast(vector_sum,c,MPI_INT,0,MPI_COMM_WORLD);
    
    printf("Process %d received the sum: ", rank);
for (int i = 0; i < c; i++) {
    printf("%d ", vector_sum[i]);
}
printf("\n");
    
    MPI_Finalize();
    
    
    

}
