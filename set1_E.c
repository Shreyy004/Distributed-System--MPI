#include<stdio.h>
#include<mpi.h>


int main(int argc,char** argv) {
   
   MPI_Init(&argc,&argv);
   int rank , size;
   MPI_Comm_rank(MPI_COMM_WORLD,&rank);
   MPI_Comm_size(MPI_COMM_WORLD,&size);
   
   int N = size * 4;
   int a[N];
   if (rank ==0 ) {
   for(int i=0;i<N;i++) {
     a[i] = i+1;
   }
   }
   int local[N/size];
   MPI_Scatter(a,N/size,MPI_INT,local,N/size,MPI_INT,0,MPI_COMM_WORLD);
   
   int local_sum = 0;
   for(int i=0;i<N/size;i++) {
    local_sum += local[i];
   }
   printf("Process %d partial sum is %d\n",rank,local_sum);
   int gather_sum[size];
   MPI_Gather(&local_sum,1,MPI_INT,gather_sum,1,MPI_INT,0,MPI_COMM_WORLD);
   
   int sum=0;
   MPI_Reduce(&local_sum,&sum,1,MPI_INT,MPI_SUM,0,MPI_COMM_WORLD);
   
   MPI_Bcast(&sum,1,MPI_INT,0,MPI_COMM_WORLD);
    MPI_Finalize();

}
