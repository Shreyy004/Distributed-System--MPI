/*
- is used to collect data from multiple processes and assemble it into a single array at the root process.
- reverse of scatter

-each process send its local data
- roots tores all received data in a single array rank order.
- it is blocking , meaning a process must wait until it has sent its data.

Key Points

✅ Reverse of MPI_Scatter → Instead of sending, it receives data from all processes.
✅ Processes send in rank order → Data is automatically stored in the correct order at the root.
✅ Only root process stores the gathered data, others don’t need a recvbuf.

*/

#include<stdio.h>
#include<mpi.h>

int main(int argc, char** argv) {

  int rank, size;
  MPI_Init(&argc,&argv);
  MPI_Comm_rank(MPI_COMM_WORLD, &rank);
  MPI_Comm_size(MPI_COMM_WORLD, &size);
  
  int local_data = (rank + 1) * 10;
  int gathered_data[4]; 
  
  MPI_Gather(&local_data, 1, MPI_INT, gathered_data, 1, MPI_INT, 0,MPI_COMM_WORLD);
  
  if(rank ==0) {
    printf("Root process collected: ");
    for (int i=0;i<size;i++) {
      printf("%d",gathered_data[i]);
    }
    printf("\n");
    
  }
  
  MPI_Finalize();
  return 0;


}

